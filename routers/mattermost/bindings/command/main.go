package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(siteURL, appID string) *apps.Binding {
	base := &apps.Binding{
		Label:       constants.CommandTrigger,
		Description: "Test app",
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

	base.Bindings = append(base.Bindings, getValid(siteURL, appID))
	base.Bindings = append(base.Bindings, getInvalid(siteURL, appID))
	base.Bindings = append(base.Bindings, getError(siteURL, appID))

	return out
}
