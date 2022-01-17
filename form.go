package main

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func formCommandBinding(cc apps.Context) apps.Binding {
	return apps.Binding{
		Label: "form",
		Bindings: []apps.Binding{
			newBinding("buttons", path.FormButtons),
			newBinding("full-readonly", path.FormFullReadonly),
			newBinding("full", path.FormFull),
			newBinding("lookup", path.FormLookup),
			newBinding("markdown-error-missing-field", path.FormMarkdownErrorMissingField),
			newBinding("markdown-error", path.FormMarkdownError),
			newBinding("multiselect", path.FormMultiselect),
			newBinding("refresh", path.FormRefresh),
			newBinding("simple", path.FormSimple),
		},
	}
}

func initHTTPForms(r *mux.Router) {
	handleCall(r, path.FormButtons, handleFormButtons)
	handleCall(r, path.FormFull, handleForm(fullForm))
	handleCall(r, path.FormFullReadonly, handleForm(fullFormReadonly()))
	handleCall(r, path.FormFullSource, handleForm(simpleFormSource))
	handleCall(r, path.InvalidForm, handleForm(apps.Form{}))
	handleCall(r, path.FormLookup, handleForm(lookupForm))
	handleCall(r, path.FormMarkdownError, handleForm(formWithMarkdownError))
	handleCall(r, path.FormMarkdownErrorMissingField, handleForm(formWithMarkdownErrorMissingField))
	handleCall(r, path.FormMultiselect, handleForm(formMultiselect))
	handleCall(r, path.FormRefresh, handleFormRefresh)
	handleCall(r, path.FormSimple, handleForm(simpleForm))
	handleCall(r, path.FormSimpleSource, handleForm(simpleFormSource))
}

var simpleForm = apps.Form{
	Title:  "Simple Form",
	Submit: callOK,
	Fields: []apps.Field{
		{
			Type: apps.FieldTypeText,
			Name: "test_field",
		},
	},
}

var simpleFormSource = apps.Form{
	Source: apps.NewCall(path.FormSimple),
}

var invalidFormSource = apps.Form{
	Source: apps.NewCall(path.InvalidForm),
}

var fullFormSource = apps.Form{
	Source: apps.NewCall(path.FormFull),
}
