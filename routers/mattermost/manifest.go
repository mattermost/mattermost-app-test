package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const (
	displayName = "Test App"
	description = "Test App"
)

func fManifest(w http.ResponseWriter, r *http.Request) {
	baseURL := "http://localhost:3000"
	manifest := apps.Manifest{
		AppID:       "com.mattermost.test",
		DisplayName: displayName,
		Description: description,
		HTTPRootURL: baseURL,
		HomepageURL: baseURL,
		Type:        apps.AppTypeHTTP,
		RequestedLocations: apps.Locations{
			apps.LocationPostMenu,
			apps.LocationCommand,
			apps.LocationChannelHeader,
		},
	}

	utils.WriteManifest(w, manifest)
}
