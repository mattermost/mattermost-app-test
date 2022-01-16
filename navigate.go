package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

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
