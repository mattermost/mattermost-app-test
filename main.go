package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const (
	AppSecret      = "1234"
	CommandTrigger = "test"
)

var localMode bool
var includeInvalid bool

//go:embed manifest.json
var manifestSource []byte //nolint: gochecknoglobals
var manifest apps.Manifest

//go:embed static
var staticAssets embed.FS //nolint: gochecknoglobals

func main() {
	err := json.Unmarshal(manifestSource, &manifest)
	if err != nil {
		panic("failed to load manfest: " + err.Error())
	}

	localMode = os.Getenv("LOCAL") == "true"
	includeInvalid = os.Getenv("INCLUDE_INVALID") == "true"

	r := mux.NewRouter()
	initHTTP(r)
	http.Handle("/", r)

	if localMode {
		baseURL := "http://localhost:3000"
		if len(os.Args) > 1 {
			baseURL = os.Args[1]
		}

		manifest.HTTP.RootURL = baseURL

		addr := ":3000"
		if len(os.Args) > 2 {
			addr = os.Args[2]
		}

		fmt.Println("Listening on", addr)
		fmt.Println("Use '/apps install http " + baseURL + "/manifest.json' to install the app")
		fmt.Printf("Use %q as the app's JWT secret\n", AppSecret)
		panic(http.ListenAndServe(addr, nil))
	}

	lambda.Start(httpadapter.New(r).Proxy)
}
