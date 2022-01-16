package main

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

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

func handleForm(f apps.Form) http.HandlerFunc {
	return httputils.DoHandleJSON(apps.NewFormResponse(f))
}
