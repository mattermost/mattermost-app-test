package main

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"

	"github.com/mattermost/mattermost-app-test/path"
)

var callOK = apps.NewCall(path.OK)

var responseOK = apps.CallResponse{
	Type: apps.CallResponseTypeOK,
	Data: "any data is OK",
	Text: "OK",
}

func initHTTPOK(r *mux.Router) {
	handleCall(r, path.OK, handleOK)
	handleCall(r, path.OKEmpty, handleOKEmpty)
}

func handleOK(creq *apps.CallRequest) apps.CallResponse {
	return apps.NewTextResponse("```\n%s\n```\n", utils.Pretty(creq))
}

func handleOKEmpty(_ *apps.CallRequest) apps.CallResponse {
	return apps.NewTextResponse("")
}
