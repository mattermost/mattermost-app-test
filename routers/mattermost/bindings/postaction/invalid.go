package postaction

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getInvalid(siteURL, appID string) []apps.Binding {
	base := []apps.Binding{}

	base = append(base, getWithNoLabel(siteURL, appID))
	base = append(base, getWithNoCall(siteURL, appID))
	base = append(base, getWithWhitespaceLabel(siteURL, appID))

	return base
}

func getWithNoLabel(siteURL, appID string) apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return apps.Binding{
		Location: "ERROR_with_no_label",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithNoCall(siteURL, appID string) apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return apps.Binding{
		Location: "ERROR_with_no_call",
		Label:    "ERROR_with_no_call",
		Icon:     icon,
	}
}

func getWithWhitespaceLabel(siteURL, appID string) apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return apps.Binding{
		Location: "ERROR_with_whitespace_label",
		Label:    " ",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}
