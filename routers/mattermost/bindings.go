package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings"
	"github.com/mattermost/mattermost-app-test/utils"
)

func fBindings(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	out := bindings.Get(c.Context)
	utils.WriteBindings(w, out)
}
