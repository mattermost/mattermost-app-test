package postaction

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getValid(siteURL string) []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithOK(siteURL))
	base = append(base, getWithEmptyOK(siteURL))
	base = append(base, getWithForm(siteURL))
	base = append(base, getWithFullForm(siteURL))
	base = append(base, getWithNavigateExternal(siteURL))
	base = append(base, getWithNavigateInternal(siteURL))
	base = append(base, getWithoutIcon(siteURL))

	return base
}

func getWithOK(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_ok",
		Label:    "with_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithEmptyOK(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_empty_ok",
		Label:    "with_empty_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOKEmpty,
		},
	}
}

func getWithForm(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithFullForm(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_full_form",
		Label:    "with_full_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFullFormOK,
		},
	}
}

func getWithNavigateExternal(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_navigate_external",
		Label:    "with_navigate_external",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithNavigateInternal(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_naviate_internal",
		Label:    "with_naviate_internal",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}

func getWithoutIcon(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "without_icon",
		Label:    "without_icon",
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}
