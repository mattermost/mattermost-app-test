package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fBindings(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	out := bindings.Get(c.Context.MattermostSiteURL, string(c.Context.AppID))
	utils.WriteBindings(w, out)
}
