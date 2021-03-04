package channel_header

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getError() []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithError())
	base = append(base, getWithEmptyError())
	base = append(base, getWithInvalidForm())
	base = append(base, getWithNavigateInvalid())
	base = append(base, getWith404())
	base = append(base, getWithHTML())
	base = append(base, getWithManifest())
	base = append(base, getWithUnknownResponse())
	base = append(base, getWith404Icon())
	base = append(base, getWithSVG())
	return base
}

func getWithError() *apps.Binding {
	return &apps.Binding{
		Location: "with_error",
		Label:    "with_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithEmptyError() *apps.Binding {
	return &apps.Binding{
		Location: "with_empty_error",
		Label:    "with_empty_error",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathErrorEmpty,
		},
	}
}
func getWithInvalidForm() *apps.Binding {
	return &apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormInvalid,
		},
	}
}

func getWithNavigateInvalid() *apps.Binding {
	return &apps.Binding{
		Location: "with_navigate_invalid",
		Label:    "with_navigate_invalid",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInvalid,
		},
	}
}

func getWith404() *apps.Binding {
	return &apps.Binding{
		Location: "with_404",
		Label:    "with_404",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPath404,
		},
	}
}

func getWithHTML() *apps.Binding {
	return &apps.Binding{
		Location: "with_html",
		Label:    "with_html",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithManifest() *apps.Binding {
	return &apps.Binding{
		Location: "with_manifest",
		Label:    "with_manifest",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse() *apps.Binding {
	return &apps.Binding{
		Location: "with_unkwon_response",
		Label:    "with_unkwon_response",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWith404Icon() *apps.Binding {
	return &apps.Binding{
		Location: "with_404_icon",
		Label:    "with_404_icon",
		Icon:     constants.BindingPath404,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithSVG() *apps.Binding {
	return &apps.Binding{
		Location: "with_svg",
		Label:    "with_svg",
		Icon:     svgIcon,
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}
