package channelheader

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getValid() []apps.Binding {
	base := []apps.Binding{}

	base = append(base, getWithOK())
	base = append(base, getWithEmptyOK())
	base = append(base, getWithForm())
	base = append(base, getWithFullForm())
	base = append(base, getWithDynamicForm())
	base = append(base, getWithNavigateExternal())
	base = append(base, getWithNavigateInternal())
	base = append(base, getWithFormInBinding())

	return base
}

func getWithOK() apps.Binding {
	return apps.Binding{
		Location: "with_ok",
		Label:    "with_ok",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithEmptyOK() apps.Binding {
	return apps.Binding{
		Location: "with_empty_ok",
		Label:    "with_empty_ok",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathOKEmpty,
		},
	}
}

func getWithForm() apps.Binding {
	return apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithFullForm() apps.Binding {
	return apps.Binding{
		Location: "with_full_form",
		Label:    "with_full_form",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathFullFormOK,
		},
	}
}

func getWithDynamicForm() apps.Binding {
	return apps.Binding{
		Location: "with_dynamic_form",
		Label:    "with_dynamic_form",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathDynamicFormOK,
		},
	}
}

func getWithNavigateExternal() apps.Binding {
	return apps.Binding{
		Location: "with_navigate_external",
		Label:    "with_navigate_external",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithNavigateInternal() apps.Binding {
	return apps.Binding{
		Location: "with_naviate_internal",
		Label:    "with_naviate_internal",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}

func getWithFormInBinding() apps.Binding {
	return apps.Binding{
		Location: "formInBinding",
		Label:    "formInBinding",
		Icon:     "icon.png",
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
		Form: &apps.Form{
			Title: "Some form",
			Fields: []apps.Field{
				{
					Name:  "foo",
					Type:  apps.FieldTypeText,
					Label: "foo",
				},
			},
		},
	}
}
