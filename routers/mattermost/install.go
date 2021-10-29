package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

func fInstall(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallStandardResponse(w, "Welcome to the test App.")
}
