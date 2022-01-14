package postaction

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func Get(context apps.Context) apps.Binding {
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
				Path: constants.SubmitOK,
			},
		})
	}

	out.Bindings = append(out.Bindings, getValid()...)
	out.Bindings = append(out.Bindings, getInvalid()...)
	out.Bindings = append(out.Bindings, getError()...)

	return out
}
