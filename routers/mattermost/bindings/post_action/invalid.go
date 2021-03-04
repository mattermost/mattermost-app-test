package post_action

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getInvalid() []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithNoLabel())
	base = append(base, getWithNoCall())
	base = append(base, getWithWhitespaceLabel())
	return base
}

func getWithNoLabel() *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_no_label",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithNoCall() *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_no_call",
		Label:    "ERROR_with_no_call",
		Icon:     icon,
	}
}

func getWithWhitespaceLabel() *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_whitespace_label",
		Label:    " ",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}
