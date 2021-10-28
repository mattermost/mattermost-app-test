package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getOthers(context apps.Context) apps.Binding {
	base := apps.Binding{
		Location: "others",
		Label:    "others",
		Bindings: []apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getPostOpenDialogTest())

	return base
}

func getPostOpenDialogTest() apps.Binding {
	return apps.Binding{
		Location: "open_dialog_from_interactive_message",
		Label:    "open_dialog_from_interactive_message",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.OtherPathOpenDialog,
			Expand: &apps.Expand{
				ActingUserAccessToken: apps.ExpandAll,
				App:                   apps.ExpandAll,
			},
		},
	}
}
