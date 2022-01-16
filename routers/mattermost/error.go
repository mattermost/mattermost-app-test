package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

func fError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallErrorResponse(w, "Error")
}

func fEmptyError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallErrorResponse(w, "")
}
