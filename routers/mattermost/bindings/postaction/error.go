package postaction

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getError(siteURL, appID string) []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithError(siteURL, appID))
	base = append(base, getWithEmptyError(siteURL, appID))
	base = append(base, getWithInvalidForm(siteURL, appID))
	base = append(base, getWithNavigateInvalid(siteURL, appID))
	base = append(base, getWith404(siteURL, appID))
	base = append(base, getWithHTML(siteURL, appID))
	base = append(base, getWithManifest(siteURL, appID))
	base = append(base, getWithUnknownResponse(siteURL, appID))
	base = append(base, getWithNoIcon(siteURL, appID))
	base = append(base, getWithSVG(siteURL, appID))

	return base
}

func getWithError(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_error",
		Label:    "with_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithEmptyError(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_empty_error",
		Label:    "with_empty_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathErrorEmpty,
		},
	}
}
func getWithInvalidForm(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormInvalid,
		},
	}
}

func getWithNavigateInvalid(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_navigate_invalid",
		Label:    "with_navigate_invalid",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInvalid,
		},
	}
}

func getWith404(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_404",
		Label:    "with_404",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPath404,
		},
	}
}

func getWithHTML(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_html",
		Label:    "with_html",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithManifest(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_manifest",
		Label:    "with_manifest",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse(siteURL, appID string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png", appID)

	return &apps.Binding{
		Location: "with_unkwon_response",
		Label:    "with_unkwon_response",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithNoIcon(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_no_icon",
		Label:    "with_no_icon",
		Icon:     constants.BindingPath404,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithSVG(siteURL, appID string) *apps.Binding {
	svgIcon := utils.GetIconURL(siteURL, "icon.svg", appID)

	return &apps.Binding{
		Location: "with_svg",
		Label:    "with_svg",
		Icon:     svgIcon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}
