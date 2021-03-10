package main

import (
	"embed"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/routers/mattermost"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const (
	baseURLPosition = 1
	addressPosition = 2
)

//go:embed manifest.json
var manifestSource []byte //nolint: gochecknoglobals

//go:embed static
var staticAssets embed.FS //nolint: gochecknoglobals

func main() {
	var manifest apps.Manifest

	err := json.Unmarshal(manifestSource, &manifest)
	if err != nil {
		panic("failed to load manfest: " + err.Error())
	}

	// Init routers
	r := mux.NewRouter()
	mattermost.Init(r, &manifest, staticAssets)

	http.Handle("/", r)

	if os.Getenv("LOCAL") == "true" {
		baseURL := "http://localhost:3000"
		if len(os.Args) > baseURLPosition {
			baseURL = os.Args[baseURLPosition]
		}

		addr := ":3000"
		if len(os.Args) > addressPosition {
			addr = os.Args[addressPosition]
		}

		manifest.HTTPRootURL = baseURL
		manifest.Type = apps.AppTypeHTTP

		_ = http.ListenAndServe(addr, nil)

		return
	}

	lambda.Start(httpadapter.New(r).Proxy)
}
