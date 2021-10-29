package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
)

func Get(context *apps.Context) *apps.Binding {
	siteURL := context.MattermostSiteURL
	appID := string(context.AppID)
	base := &apps.Binding{
		Label:       constants.CommandTrigger,
		Description: "Test commands",
		Location:    constants.CommandTrigger,
		Icon:        utils.GetIconURL(siteURL, "icon.png", appID),
		Bindings:    []*apps.Binding{},
	}
	out := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{
			base,
		},
	}

	if context.Channel.Name == "town-square" {
		base.Bindings = append(base.Bindings, &apps.Binding{
			Location: "town_square",
			Label:    "town_square",
			Form: &apps.Form{
				Fields: []*apps.Field{},
			},
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
		})
	}

	base.Bindings = append(base.Bindings, getValid(siteURL, appID))
	base.Bindings = append(base.Bindings, getInvalid(siteURL, appID))
	base.Bindings = append(base.Bindings, getError(siteURL, appID))
	base.Bindings = append(base.Bindings, getOthers(context))
	base.Bindings = append(base.Bindings, getSubscribeCommand(context))

	return out
}
