package channel_header

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getValid() []*apps.Binding {
	base := []*apps.Binding{}

	base = append(base, getWithOK())
	base = append(base, getWithEmptyOK())
	base = append(base, getWithForm())
	base = append(base, getWithNavigateExternal())
	base = append(base, getWithNavigateInternal())
	return base
}

func getWithOK() *apps.Binding {
	return &apps.Binding{
		Location: "with_ok",
		Label:    "with_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithEmptyOK() *apps.Binding {
	return &apps.Binding{
		Location: "with_empty_ok",
		Label:    "with_empty_ok",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathOKEmpty,
		},
	}
}

func getWithForm() *apps.Binding {
	return &apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithNavigateExternal() *apps.Binding {
	return &apps.Binding{
		Location: "with_navigate_external",
		Label:    "with_navigate_external",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithNavigateInternal() *apps.Binding {
	return &apps.Binding{
		Location: "with_naviate_internal",
		Label:    "with_naviate_internal",
		Icon:     icon,
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}
