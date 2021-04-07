package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func fOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallStandardResponse(w, "OK")
}

func fEmptyOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallStandardResponse(w, "")
}
