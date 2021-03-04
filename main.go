package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/routers/mattermost"
)

func main() {
	// Init routers
	r := mux.NewRouter()
	mattermost.Init(r)

	http.Handle("/", r)
	_ = http.ListenAndServe(":3000", nil)
}
