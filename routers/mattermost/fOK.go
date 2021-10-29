package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

func fOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if c.Values != nil && c.Values["navigate"] != nil {
		if navigate, _ := c.Values["navigate"].(bool); navigate {
			handleNavigate(w, c)
			return
		}
	}

	utils.WriteCallStandardResponse(w, "OK")
}

func fEmptyOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallStandardResponse(w, "")
}

func handleNavigate(w http.ResponseWriter, c *apps.CallRequest) {
	team := c.Context.Team

	teamName := team.Name
	resp := apps.CallResponse{
		Type:          apps.CallResponseTypeNavigate,
		NavigateToURL: fmt.Sprintf("%s/%s/channels/town-square", c.Context.MattermostSiteURL, teamName),
	}
	utils.WriteCallResponse(w, resp)
}
