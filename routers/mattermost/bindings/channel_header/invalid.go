package channel_header

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getInvalid(siteURL string) []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithNoIcon(siteURL))
	base = append(base, getWithNoLabel(siteURL))
	base = append(base, getWithNoCall(siteURL))
	base = append(base, getWithWhitespaceLabel(siteURL))
	return base
}

func getWithNoIcon(siteURL string) *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_no_icon",
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithNoLabel(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")
	return &apps.Binding{
		Location: "ERROR_with_no_label",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithNoCall(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")
	return &apps.Binding{
		Location: "ERROR_with_no_call",
		Label:    "ERROR_with_no_call",
		Icon:     icon,
	}
}

func getWithWhitespaceLabel(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")
	return &apps.Binding{
		Location: "ERROR_with_whitespace_label",
		Label:    " ",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}
