package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func getValid() apps.Binding {
	base := apps.Binding{
		Location: "valid",
		Label:    "valid",
		Bindings: []apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getWithEmptyForm())
	base.Bindings = append(base.Bindings, getWithForm())
	base.Bindings = append(base.Bindings, getOpenFullFormModal())
	base.Bindings = append(base.Bindings, getOpenDynamicFormModal())
	base.Bindings = append(base.Bindings, getRedefineForm())
	base.Bindings = append(base.Bindings, getEmbeddedForm())
	base.Bindings = append(base.Bindings, getWithoutForm())
	base.Bindings = append(base.Bindings, getWithLookup())
	base.Bindings = append(base.Bindings, getWithEmptyLookup())
	base.Bindings = append(base.Bindings, getWithFullForm())
	base.Bindings = append(base.Bindings, getWithNoOptionStatic())
	base.Bindings = append(base.Bindings, getWithInternalNavResponse())
	base.Bindings = append(base.Bindings, getWithExternalNavResponse())
	base.Bindings = append(base.Bindings, getWithFormResponse())
	base.Bindings = append(base.Bindings, getWithRequiredFields())
	base.Bindings = append(base.Bindings, getWithMultiwordOption())
	base.Bindings = append(base.Bindings, getWithMultiwordDynamicOption())
	base.Bindings = append(base.Bindings, getOpenMultiselectForm())
	base.Bindings = append(base.Bindings, getMultiselectCommand())
	base.Bindings = append(base.Bindings, getOpenFullFormDisabledModal())
	base.Bindings = append(base.Bindings, getOpenFormWithButtons())
	base.Bindings = append(base.Bindings, getOpenMarkdownForm())
	base.Bindings = append(base.Bindings, getOpenMarkdownFormWithMissingFieldError())
	base.Bindings = append(base.Bindings, getWithCallInForm())

	return base
}

func getWithEmptyForm() apps.Binding {
	return apps.Binding{
		Location: "empty_form",
		Label:    "empty_form",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithForm() apps.Binding {
	return apps.Binding{
		Location: "with_form",
		Label:    "with_form",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithoutForm() apps.Binding {
	return apps.Binding{
		Location: "no_form",
		Label:    "no_form",
		Call: &apps.Call{
			Path: constants.BindingPathFormOK,
		},
	}
}

func getRedefineForm() apps.Binding {
	return apps.Binding{
		Location: "redefine_form",
		Label:    "redefine_form",
		Call: &apps.Call{
			Path: constants.BindingPathRedefineFormOK,
		},
	}
}

func getEmbeddedForm() apps.Binding {
	return apps.Binding{
		Location: "embedded_form",
		Label:    "embedded_form",
		Form:     &apps.Form{Fields: []apps.Field{}},
		Call: &apps.Call{
			Path: constants.BindingPathEmbeddedFormOK,
		},
	}
}

func getWithLookup() apps.Binding {
	return apps.Binding{
		Location: "lookup",
		Label:    "lookup",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithEmptyLookup() apps.Binding {
	return apps.Binding{
		Location: "empty_lookup",
		Label:    "empty_lookup",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithFullForm() apps.Binding {
	return apps.Binding{
		Location: "full_form",
		Label:    "full_form",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
				{
					Name:                 "mybeatifulflag",
					Type:                 apps.FieldTypeText,
					Label:                "mybeatifulflag",
					AutocompletePosition: -1,
				},
				{
					Name:                 "otherpotitional",
					Type:                 apps.FieldTypeText,
					Label:                "mybeatifulpositional",
					AutocompletePosition: 1,
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
				{
					Name:     "textReadOnly",
					Type:     apps.FieldTypeText,
					Label:    "textReadOnly",
					ReadOnly: true,
					Value:    "Readonly value",
				},
				{
					Name: "mk1",
					Type: "markdown",
					Description: "## Markdown title" +
						"\nHello world" +
						"\nText styles: _italics_ **bold** **_bold-italic_** ~~strikethrough~~ `code`" +
						"\nUsers and channels: @sysadmin ~town-square" +
						"\n```" +
						"\nCode block" +
						"\n```" +
						"\n:+1: :banana_dance:" +
						"\n***" +
						"\n> Quote\n" +
						"\nLink: [here](www.google.com)" +
						"\nImage: ![img](https://gdm-catalog-fmapi-prod.imgix.net/ProductLogo/4acbc64f-552d-4944-8474-b44a13a7bd3e.png?auto=format&q=50&fit=fill)" +
						"\nList:" +
						"\n- this" +
						"\n- is" +
						"\n- a" +
						"\n- list" +
						"\nNumbered list" +
						"\n1. this" +
						"\n2. is" +
						"\n3. a" +
						"\n4. list" +
						"\nItems" +
						"\n- [ ] Item one" +
						"\n- [ ] Item two" +
						"\n- [x] Completed item",
				},
				{
					Name: "mk2",
					Type: "markdown",
					Description: "\n| Left-Aligned  | Center Aligned  | Right Aligned |" +
						"\n| :------------ |:---------------:| -----:|" +
						"\n| Left column 1 | this text       |  $100 |" +
						"\n| Left column 2 | is              |   $10 |" +
						"\n| Left column 3 | centered        |    $1 |",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathLookupOK,
		},
	}
}

func getWithNoOptionStatic() apps.Binding {
	return apps.Binding{
		Location: "no_option_static",
		Label:    "no_option_static",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithInternalNavResponse() apps.Binding {
	return apps.Binding{
		Location: "external_nav",
		Label:    "external_nav",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithExternalNavResponse() apps.Binding {
	return apps.Binding{
		Location: "internal_nav",
		Label:    "internal_nav",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInternal,
		},
	}
}

func getWithFormResponse() apps.Binding {
	return apps.Binding{
		Location: "with_form_response",
		Label:    "with_form_response",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithRequiredFields() apps.Binding {
	return apps.Binding{
		Location: "with_required_fields",
		Label:    "with_required_fields",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithMultiwordOption() apps.Binding {
	return apps.Binding{
		Location: "with_multiword_option",
		Label:    "with_multiword_option",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getWithMultiwordDynamicOption() apps.Binding {
	return apps.Binding{
		Location: "with_multiword_dynamic_option",
		Label:    "with_multiword_dynamic_option",
		Form: &apps.Form{
			Fields: []apps.Field{
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

func getOpenFullFormModal() apps.Binding {
	return apps.Binding{
		Location: "open_full_form_modal",
		Label:    "open_full_form_modal",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathFullFormOK,
		},
	}
}

func getOpenFullFormDisabledModal() apps.Binding {
	return apps.Binding{
		Location: "open_full_form_disabled_modal",
		Label:    "open_full_form_disabled_modal",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathFullDisabledOK,
		},
	}
}

func getOpenDynamicFormModal() apps.Binding {
	return apps.Binding{
		Location: "open_dynamic_form_modal",
		Label:    "open_dynamic_form_modal",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathDynamicFormOK,
		},
	}
}

func getOpenFormWithButtons() apps.Binding {
	return apps.Binding{
		Location: "open_form_with_buttons",
		Label:    "open_form_with_buttons",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathWithButtonsOK,
		},
	}
}

func getOpenMultiselectForm() apps.Binding {
	return apps.Binding{
		Location: "open_multiselect_form_modal",
		Label:    "open_multiselect_form_modal",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathMultiselectForm,
		},
	}
}

func getOpenMarkdownForm() apps.Binding {
	return apps.Binding{
		Location: "open_markdown_form_modal",
		Label:    "open_markdown_form_modal",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathMarkdownForm,
		},
	}
}

func getOpenMarkdownFormWithMissingFieldError() apps.Binding {
	return apps.Binding{
		Location: "open_markdown_form_modal_missing_error",
		Label:    "open_markdown_form_modal_missing_error",
		Form: &apps.Form{
			Fields: []apps.Field{},
		},
		Call: &apps.Call{
			Path: constants.BindingPathMarkdownFormWithMissingError,
		},
	}
}

func getMultiselectCommand() apps.Binding {
	return apps.Binding{
		Location: "multiselect_command",
		Label:    "multiselect_command",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:          "static",
					Type:          apps.FieldTypeStaticSelect,
					Label:         "static",
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
						{
							Label: "static value 3",
							Value: "sv3",
						},
						{
							Label: "static value 4",
							Value: "sv4",
						},
					},
				},
				{
					Name:          "user",
					Type:          apps.FieldTypeUser,
					Label:         "user",
					SelectIsMulti: true,
				},
				{
					Name:          "channel",
					Type:          apps.FieldTypeChannel,
					Label:         "channel",
					SelectIsMulti: true,
				},
			},
		},
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithCallInForm() apps.Binding {
	return apps.Binding{
		Location: "callInForm",
		Label:    "callInForm",
		Form: &apps.Form{
			Title: "Some form",
			Fields: []apps.Field{
				{
					Name:  "foo",
					Type:  apps.FieldTypeText,
					Label: "foo",
				},
			},
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
		},
	}
}
