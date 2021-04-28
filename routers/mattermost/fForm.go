package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func fFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: []*apps.Field{
				{
					Name:       "text",
					Type:       apps.FieldTypeText,
					Label:      "text",
					ModalLabel: "text",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFullFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Full Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathLookupOK,
			},
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
	}
	utils.WriteCallResponse(w, resp)
}

func fDynamicFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numFields := len(c.Values)
	fields := []*apps.Field{}

	for i := 0; i < numFields+5; i++ {
		fields = append(fields, &apps.Field{
			Name:          fmt.Sprintf("static%v", i),
			Type:          apps.FieldTypeStaticSelect,
			Label:         fmt.Sprintf("static%v", i),
			SelectRefresh: true,
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
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Dynamic Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathDynamicFormOK,
			},
			Fields: fields,
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormInvalid(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
	}
	utils.WriteCallResponse(w, resp)
}

var numElementsDefined = 0

func fFormRedefine(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numElementsDefined++
	fields := []*apps.Field{}
	for i := 0; i < numElementsDefined; i++ {
		name := fmt.Sprintf("text%v", i)
		fields = append(fields, &apps.Field{
			Name:       name,
			Type:       apps.FieldTypeText,
			Label:      name,
			ModalLabel: name,
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: fields,
		},
	}
	utils.WriteCallResponse(w, resp)
}
