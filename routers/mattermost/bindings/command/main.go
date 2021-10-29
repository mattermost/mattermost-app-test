package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func Get(context apps.Context) apps.Binding {
	base := apps.Binding{
		Label:       constants.CommandTrigger,
		Description: "Test commands",
		Location:    constants.CommandTrigger,
		Icon:        "icon.png",
		Bindings:    []apps.Binding{},
	}

	if context.Channel.Name == "town-square" {
		base.Bindings = append(base.Bindings, apps.Binding{
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

	base.Bindings = append(base.Bindings, getValid())
	base.Bindings = append(base.Bindings, getInvalid())
	base.Bindings = append(base.Bindings, getError())
	base.Bindings = append(base.Bindings, getOthers(context))
	base.Bindings = append(base.Bindings, getSubscribeCommand(context))

	out := apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []apps.Binding{
			base,
		},
	}

	return out
}
