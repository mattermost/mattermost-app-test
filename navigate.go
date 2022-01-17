package main

import (
	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

func initHTTPNavigate(r *mux.Router) {
	handleCall(r, path.NavigateExternal, handleNavigateExternal)
	handleCall(r, path.NavigateInternal, handleNavigateInternal)
	handleCall(r, path.InvalidNavigate, handleNavigateInvalid)
}

func handleNavigateInternal(creq *apps.CallRequest) apps.CallResponse {
	return apps.CallResponse{
		Type:          apps.CallResponseTypeNavigate,
		NavigateToURL: creq.Context.MattermostSiteURL + "/ad-1/channels/town-square",
	}
}

func handleNavigateExternal(creq *apps.CallRequest) apps.CallResponse {
	return apps.CallResponse{
		Type:          apps.CallResponseTypeNavigate,
		NavigateToURL: "http://www.google.com",
	}
}

func handleNavigateInvalid(creq *apps.CallRequest) apps.CallResponse {
	return apps.CallResponse{
		Type: apps.CallResponseTypeNavigate,
	}
}
