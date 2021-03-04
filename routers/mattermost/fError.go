package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fError(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	utils.WriteCallErrorResponse(w, "Error")
}

func fEmptyError(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	utils.WriteCallErrorResponse(w, "")
}
