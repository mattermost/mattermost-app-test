package postaction

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func getError() []apps.Binding {
	base := []apps.Binding{}

	base = append(base, getWithError())
	base = append(base, getWithEmptyError())
	base = append(base, getWithInvalidForm())
	base = append(base, getWithNavigateInvalid())
	base = append(base, getWith404())
	base = append(base, getWithHTML())
	base = append(base, getWithManifest())
	base = append(base, getWithUnknownResponse())
	base = append(base, getWithNoIcon())
	base = append(base, getWithSVG())

	return base
}

func getWithError() apps.Binding {
	return apps.Binding{
		Location: "with_error",
		Label:    "with_error",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.Error,
		},
	}
}

func getWithEmptyError() apps.Binding {
	return apps.Binding{
		Location: "with_empty_error",
		Label:    "with_empty_error",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.ErrorEmpty,
		},
	}
}
func getWithInvalidForm() apps.Binding {
	return apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.FormInvalid,
		},
	}
}

func getWithNavigateInvalid() apps.Binding {
	return apps.Binding{
		Location: "with_navigate_invalid",
		Label:    "with_navigate_invalid",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.NavigateInvalid,
		},
	}
}

func getWith404() apps.Binding {
	return apps.Binding{
		Location: "with_404",
		Label:    "with_404",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.NotFoundPath,
		},
	}
}

func getWithHTML() apps.Binding {
	return apps.Binding{
		Location: "with_html",
		Label:    "with_html",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.HTMLPath,
		},
	}
}

func getWithManifest() apps.Binding {
	return apps.Binding{
		Location: "with_manifest",
		Label:    "with_manifest",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse() apps.Binding {
	return apps.Binding{
		Location: "with_unkwon_response",
		Label:    "with_unkwon_response",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}

func getWithNoIcon() apps.Binding {
	return apps.Binding{
		Location: "with_no_icon",
		Label:    "with_no_icon",
		Icon:     "foo",
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}

func getWithSVG() apps.Binding {
	return apps.Binding{
		Location: "with_svg",
		Label:    "with_svg",
		Icon:     "icon.svg",
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}
