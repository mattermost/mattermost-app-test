package channelheader

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getValid(siteURL, appID string) []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithOK(siteURL, appID))
	base = append(base, getWithEmptyOK(siteURL, appID))
	base = append(base, getWithForm(siteURL, appID))
	base = append(base, getWithFullForm(siteURL, appID))
	base = append(base, getWithDynamicForm(siteURL, appID))
	base = append(base, getWithNavigateExternal(siteURL, appID))
	base = append(base, getWithNavigateInternal(siteURL, appID))

	return base
}

func getWithOK(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_ok",
		Label:    "with_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithEmptyOK(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_empty_ok",
		Label:    "with_empty_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOKEmpty,
		},
	}
}

func getWithForm(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithFullForm(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_full_form",
		Label:    "with_full_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFullFormOK,
		},
	}
}

func getWithDynamicForm(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_dynamic_form",
		Label:    "with_dynamic_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathDynamicFormOK,
		},
	}
}

func getWithNavigateExternal(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_navigate_external",
		Label:    "with_navigate_external",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithNavigateInternal(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_naviate_internal",
		Label:    "with_naviate_internal",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}
