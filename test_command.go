package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

func testCommandBinding(cc apps.Context) apps.Binding {
	var out []apps.Binding
	if cc.Channel != nil && cc.Channel.Name == "town-square" {
		out = append(out, newBinding("town-square-channel-specific", path.OK))
	}

	out = append(out, apps.Binding{
		Label: "valid-response",
		Bindings: []apps.Binding{
			newBinding("OK", path.OK),
			newBinding("OK-empty", path.OKEmpty),
			newBinding("form", path.FormSimple),
			newBinding("form-source", path.FormSimpleSource), // TODO <>/<> does not work, move to invalid?
			newBinding("navigate-external", path.NavigateExternal),
			newBinding("navigate-internal", path.NavigateInternal),
		},
	})

	out = append(out, apps.Binding{
		Label: "error-response",
		Bindings: []apps.Binding{
			newBinding("error", path.ErrorDefault),
			newBinding("error-empty", path.ErrorEmpty),
			newBinding("error-404", path.Error404),
			newBinding("error-500", path.Error500),
		},
	})

	out = append(out, apps.Binding{
		Label: "invalid-response",
		Bindings: []apps.Binding{
			newBinding("navigate", path.InvalidNavigate),
			newBinding("form", path.InvalidForm),
			newBinding("unknown-type", path.InvalidUnknownType),
			newBinding("HTML-random", path.InvalidHTML),
			newBinding("JSON-random", path.Manifest),
		},
	})

	out = append(out, apps.Binding{
		Label: "valid-definition-form",
		Bindings: []apps.Binding{
			{
				Label:       "empty",
				Description: "Empty submittable form is included in the binding, no flags",
				Form: &apps.Form{
					Submit: callOK,
				},
			},
			{
				Label:       "simple",
				Description: "Simple form is included in the binding",
				Form:        &simpleForm,
			},
			{
				// TODO <>/<> does not work, move to invalid?
				Label:       "simple-source",
				Description: "Simple form is referenced (`source=`) in the binding, DOES NOT WORK",
				Form:        &simpleFormSource,
			},
			{
				Label:       "full",
				Form:        &fullForm,
				Description: "Full form is included in the binding",
			},
		},
	})

	out = append(out, apps.Binding{
		Label: "invalid-definition-binding",
		Bindings: []apps.Binding{
			{
				Label:       "valid",
				Description: "the only valid binding in this submenu",
				Submit:      callOK,
			},
			{
				Label:       "conflicting-label",
				Description: "2 sub-bindings have label `Command`, only one should appear",
				Bindings: []apps.Binding{
					{
						Location: "command1",
						Label:    "Command",
						Submit:   callOK,
					},
					{
						Location: "command2",
						Label:    "Command",
						Submit:   callOK,
					},
				},
			},
			{
				Label:       "space-in-label",
				Description: "`Command with space` is not visible.",
				Bindings: []apps.Binding{
					{
						Label:  "Command with space",
						Submit: callOK,
					},
					{
						Label:  "Command-with-no-space",
						Submit: callOK,
					},
				},
			},
		},
	})

	out = append(out, apps.Binding{
		Label: "invalid-definition-form",
		Bindings: []apps.Binding{
			{
				Label:       "valid",
				Description: "the only valid binding in this submenu",
				Submit:      callOK,
			},
			{
				Label:       "form-unsubmittable",
				Description: "Form is included in the binding does not have submit",
				Form: &apps.Form{
					Title: "unsubmittable form",
				},
			},
			{
				Label:       "form-conflicting-fields",
				Description: "2 fields have label `field`, only `field1` should appear",
				Form: &apps.Form{
					Submit: callOK,
					Fields: []apps.Field{
						{
							Type:  apps.FieldTypeText,
							Name:  "field1",
							Label: "field",
						},
						{
							Type:  apps.FieldTypeText,
							Name:  "field2",
							Label: "field",
						},
					},
				},
			},
			{
				Label:       "form-conflicting-options",
				Description: "2 select options have value `opt`, only `opt1` should appear",
				Form: &apps.Form{
					Submit: callOK,
					Fields: []apps.Field{
						{
							Type: apps.FieldTypeStaticSelect,
							Name: "field",
							SelectStaticOptions: []apps.SelectOption{
								{
									Label: "opt1",
									Value: "opt",
								},
								{
									Label: "opt2",
									Value: "opt",
								},
							},
						},
					},
				},
			},
			{
				Label:       "form-empty-option",
				Description: "a select option has no name/value; only `opt1` should appear",
				Form: &apps.Form{
					Submit: callOK,
					Fields: []apps.Field{
						{
							Type: apps.FieldTypeStaticSelect,
							Name: "field",
							SelectStaticOptions: []apps.SelectOption{
								{
									Label: "opt1",
									Value: "opt",
								},
								{
									Label: "",
									Value: "",
								},
							},
						},
					},
				},
			},
			{
				Label:       "form-space-in-field-label",
				Description: "a form field has a label with a space, only `field-with-no-space` should appear.",
				Form: &apps.Form{
					Submit: callOK,
					Fields: []apps.Field{
						{
							Type:  apps.FieldTypeText,
							Name:  "field1",
							Label: "field with space",
						},
						{
							Type:  apps.FieldTypeText,
							Name:  "field2",
							Label: "field-with-no-space",
						},
					},
				},
			},
		},
	})

	return apps.Binding{
		Label:    "test-command",
		Bindings: out,
	}
}
