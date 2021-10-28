package postaction

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(context apps.Context) apps.Binding {
	siteURL := context.MattermostSiteURL
	appID := string(context.AppID)
	out := apps.Binding{
		Location: apps.LocationPostMenu,
		Bindings: []apps.Binding{},
	}

	if context.Channel.Name == "town-square" {
		out.Bindings = append(out.Bindings, apps.Binding{
			Location: "town_square",
			Label:    "town_square",
			Form: &apps.Form{
				Fields: []apps.Field{},
			},
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
		})
	}

	out.Bindings = append(out.Bindings, getValid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getInvalid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getError(siteURL, appID)...)

	return out
}
