package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

var invalidResponseBinding = apps.Binding{
	Label: "invalid-response",
	Bindings: []apps.Binding{
		newBinding("invalid-navigate", path.InvalidNavigate),
		newBinding("invalid-form", path.InvalidForm),
		newBinding("unknown-type", path.InvalidUnknownType),
		newBinding("HTML-random", path.InvalidHTML),
		newBinding("JSON-random", path.Manifest),
	},
}

var invalidBindingBinding = apps.Binding{
	Label: "invalid-input-binding",
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
					Submit:   apps.NewCall(path.ErrorDefault),
				},
			},
		},
		{
			Label:       "space-in-label",
			Description: "`Command with space` is not visible.",
			Bindings: []apps.Binding{
				{
					Label:  "Command with space",
					Submit: apps.NewCall(path.ErrorDefault),
				},
				{
					Label:  "Command-with-no-space",
					Submit: callOK,
				},
			},
		},
	},
}

var invalidFormBinding = apps.Binding{
	Label: "invalid-input-form",
	Bindings: []apps.Binding{
		{
			Label:       "unsubmittable",
			Description: "Form is included in the binding does not have submit",
			Form: &apps.Form{
				Title: "unsubmittable form",
			},
		},
		{
			Label:       "conflicting-fields",
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
			Label:       "conflicting-options",
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
			Label:       "empty-option",
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
		{
			Label:       "empty-lookup",
			Description: "a form field is a dynamic select with an empty lookup response.",
			Form: &apps.Form{
				Submit: apps.NewCall(path.ErrorDefault),
				Fields: []apps.Field{
					{
						Type:                apps.FieldTypeDynamicSelect,
						IsRequired:          true,
						Name:                "field1",
						AutocompletePosition: 1,
						SelectDynamicLookup: apps.NewCall(path.LookupEmpty),
					},
				},
			},
		},
	},
}
