package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

func fManifest(m *apps.Manifest) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteManifest(w, *m)
	}
}
