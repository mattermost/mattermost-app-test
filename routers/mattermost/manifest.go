package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
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
		Bindings: &apps.Call{
			Path: constants.BindingsPath,
			Expand: &apps.Expand{
				App:                   apps.ExpandAll,
				ActingUser:            apps.ExpandAll,
				ActingUserAccessToken: apps.ExpandAll,
				AdminAccessToken:      apps.ExpandAll,
				Channel:               apps.ExpandAll,
				Mentioned:             apps.ExpandAll,
				ParentPost:            apps.ExpandAll,
				Post:                  apps.ExpandAll,
				RootPost:              apps.ExpandAll,
				Team:                  apps.ExpandAll,
				User:                  apps.ExpandAll,
			},
		},
	}

	utils.WriteManifest(w, manifest)
}
