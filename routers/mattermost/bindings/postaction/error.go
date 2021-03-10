package postaction

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getError(siteURL string) []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithError(siteURL))
	base = append(base, getWithEmptyError(siteURL))
	base = append(base, getWithInvalidForm(siteURL))
	base = append(base, getWithNavigateInvalid(siteURL))
	base = append(base, getWith404(siteURL))
	base = append(base, getWithHTML(siteURL))
	base = append(base, getWithManifest(siteURL))
	base = append(base, getWithUnknownResponse(siteURL))
	base = append(base, getWithNoIcon(siteURL))
	base = append(base, getWithSVG(siteURL))

	return base
}

func getWithError(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_error",
		Label:    "with_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithEmptyError(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_empty_error",
		Label:    "with_empty_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathErrorEmpty,
		},
	}
}
func getWithInvalidForm(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormInvalid,
		},
	}
}

func getWithNavigateInvalid(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_navigate_invalid",
		Label:    "with_navigate_invalid",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInvalid,
		},
	}
}

func getWith404(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_404",
		Label:    "with_404",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPath404,
		},
	}
}

func getWithHTML(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_html",
		Label:    "with_html",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithManifest(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_manifest",
		Label:    "with_manifest",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse(siteURL string) *apps.Binding {
	icon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_unkwon_response",
		Label:    "with_unkwon_response",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithNoIcon(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_no_icon",
		Label:    "with_no_icon",
		Icon:     constants.BindingPath404,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithSVG(siteURL string) *apps.Binding {
	svgIcon := utils.GetIconURL(siteURL, "icon.png")

	return &apps.Binding{
		Location: "with_svg",
		Label:    "with_svg",
		Icon:     svgIcon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}
