package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getValid(siteURL string) *apps.Binding {
	base := &apps.Binding{
		Location: "valid",
		Label:    "valid",
		Bindings: []*apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getWithEmptyForm(siteURL))
	base.Bindings = append(base.Bindings, getWithForm(siteURL))
	base.Bindings = append(base.Bindings, getWithoutForm(siteURL))
	base.Bindings = append(base.Bindings, getWithLookup(siteURL))
	base.Bindings = append(base.Bindings, getWithEmptyLookup(siteURL))
	base.Bindings = append(base.Bindings, getWithFullForm(siteURL))
	base.Bindings = append(base.Bindings, getWithNoOptionStatic(siteURL))
	base.Bindings = append(base.Bindings, getWithInternalNavResponse(siteURL))
	base.Bindings = append(base.Bindings, getWithExternalNavResponse(siteURL))
	base.Bindings = append(base.Bindings, getWithFormResponse(siteURL))
	base.Bindings = append(base.Bindings, getWithRequiredFields(siteURL))
	base.Bindings = append(base.Bindings, getWithMultiwordOption(siteURL))
	base.Bindings = append(base.Bindings, getWithMultiwordDynamicOption(siteURL))

	return base
}

func getWithEmptyForm(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "empty_form",
		Label:    "empty_form",
		Form: &apps.Form{
			Fields: []*apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithForm(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithoutForm(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "no_form",
		Label:    "no_form",
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithLookup(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "lookup",
		Label:    "lookup",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathLookupOK,
		},
	}
}

func getWithEmptyLookup(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "empty_lookup",
		Label:    "empty_lookup",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathLookupEmpty,
		},
	}
}

func getWithFullForm(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "full_form",
		Label:    "full_form",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
				},
				{
					Name:  "channel",
					Type:  apps.FieldTypeChannel,
					Label: "channel",
				},
				{
					Name:  "user",
					Type:  apps.FieldTypeUser,
					Label: "user",
				},
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathLookupOK,
		},
	}
}

func getWithNoOptionStatic(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "no_option_static",
		Label:    "no_option_static",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:                "static",
					Type:                apps.FieldTypeStaticSelect,
					Label:               "static",
					SelectStaticOptions: []apps.SelectOption{},
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithInternalNavResponse(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "external_nav",
		Label:    "external_nav",
		Form: &apps.Form{
			Fields: []*apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithExternalNavResponse(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "internal_nav",
		Label:    "internal_nav",
		Form: &apps.Form{
			Fields: []*apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}

func getWithFormResponse(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_form_response",
		Label:    "with_form_response",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithRequiredFields(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_required_fields",
		Label:    "with_required_fields",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
				{
					Name:       "text2",
					Type:       apps.FieldTypeText,
					Label:      "text2",
					IsRequired: true,
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithMultiwordOption(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_multiword_option",
		Label:    "with_multiword_option",
		Form: &apps.Form{
			Fields: []*apps.Field{
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "multiword option 1",
							Value: "multiword option 1",
						},
						{
							Label: "multiword option 2",
							Value: "multiword option 2",
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

func getWithMultiwordDynamicOption(_ string) *apps.Binding {
	return &apps.Binding{
		Location: "with_multiword_dynamic_option",
		Label:    "with_multiword_dynamic_option",
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
			Path: constants.BindingPathLookupMultiword,
		},
	}
}
