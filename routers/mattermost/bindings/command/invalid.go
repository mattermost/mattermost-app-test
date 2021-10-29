package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func getInvalid(siteURL, appID string) *apps.Binding {
	base := &apps.Binding{
		Location: "invalid",
		Label:    "invalid",
		Bindings: []*apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getWithSameCommand(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithSameFieldNames(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithSameOptionNames(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithMultiwordCommand(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithMultiwordField(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithWhitespaceLabel(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithNoLabel(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithNoCall(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithInvalidLookup(siteURL, appID))
	base.Bindings = append(base.Bindings, getNoNameField(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithNoNameOption(siteURL, appID))

	return base
}

func getWithSameCommand(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_same_command",
		Label:    "with_same_command",
		Bindings: []*apps.Binding{
			{
				Location: "command1",
				Label:    "Command",
				Form:     &apps.Form{},
				Call: &apps.Call{
					Path: constants.BindingPathOK,
				},
			},
			{
				Location: "command2",
				Label:    "Command",
				Form:     &apps.Form{},
				Call: &apps.Call{
					Path: constants.BindingPathError,
				},
			},
		},
	}
}

func getWithSameFieldNames(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_same_field_names",
		Label:    "with_same_field_names",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "text1",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
				{
					Name:  "text2",
					Type:  apps.FieldTypeBool,
					Label: "text",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithSameOptionNames(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_same_option_names",
		Label:    "with_same_option_names",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "select",
					Type:  apps.FieldTypeStaticSelect,
					Label: "select",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 1",
							Value: "sv2",
						},
					},
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithNoNameOption(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_no_name_option",
		Label:    "with_no_name_option",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "select",
					Type:  apps.FieldTypeStaticSelect,
					Label: "select",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "",
							Value: "",
						},
						{
							Label: "static value 1",
							Value: "sv2",
						},
					},
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithMultiwordCommand(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_multiword_command",
		Label:    "with_multiword command",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithMultiwordField(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_multiword_field",
		Label:    "with_multiword_field",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "multiword field",
					Type:  apps.FieldTypeText,
					Label: "multiword field",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getNoNameField(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_noname_field",
		Label:    "with_noname_field",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "",
					Type:  apps.FieldTypeText,
					Label: "multiword field",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithWhitespaceLabel(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_whitespace_label",
		Label:    " ",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithNoLabel(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_no_label",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithNoCall(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "ERROR_with_no_call",
		Label:    "ERROR_with_no_call",
		Form:     &apps.Form{},
	}
}

func getWithInvalidLookup(_, _ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_invalid_lookup",
		Label:    "with_invalid_lookup",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "dynamic",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "dynamic",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathLookupInvalid,
		},
	}
}
