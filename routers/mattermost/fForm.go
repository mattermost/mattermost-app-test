package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fFormOK(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
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

func fFullFormOK(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
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

func fFormInvalid(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
	}
	utils.WriteCallResponse(w, resp)
}
