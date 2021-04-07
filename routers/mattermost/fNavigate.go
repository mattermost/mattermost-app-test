package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func fNavigateInternal(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type:          apps.CallResponseTypeNavigate,
		NavigateToURL: c.Context.MattermostSiteURL + "/ad-1/channels/town-square",
	}
	utils.WriteCallResponse(w, resp)
}

func fNavigateExternal(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type:          apps.CallResponseTypeNavigate,
		NavigateToURL: "http://www.google.com",
	}
	utils.WriteCallResponse(w, resp)
}

func fNavigateInvalid(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeNavigate,
	}
	utils.WriteCallResponse(w, resp)
}
