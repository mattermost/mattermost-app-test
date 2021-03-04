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

func fFormInvalid(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
	}
	utils.WriteCallResponse(w, resp)
}
