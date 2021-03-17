package command

import (
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(siteURL string) *apps.Binding {
	base := &apps.Binding{
		Label:       "com.mattermost.test",
		Description: "Test app",
		Location:    "com.mattermost.test",
		Icon:        utils.GetIconURL(siteURL, "icon.png"),
		Bindings:    []*apps.Binding{},
	}
	out := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{
			base,
		},
	}

	base.Bindings = append(base.Bindings, getValid(siteURL))
	base.Bindings = append(base.Bindings, getInvalid(siteURL))
	base.Bindings = append(base.Bindings, getError(siteURL))

	return out
}
