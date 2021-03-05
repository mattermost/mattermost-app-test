package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

type lookupResponse struct {
	Items []apps.SelectOption `json:"items"`
}

func fLookupOK(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{
				{
					Label: "static value 1 label",
					Value: "sv1",
				},
				{
					Label: "static value 2 label",
					Value: "sv2",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupMultiword(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{
				{
					Label: "static value 2 label",
					Value: "static value 1",
				},
				{
					Label: "static value 2 label",
					Value: "static value 2",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupEmpty(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupInvalid(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{
				{
					Label: "Valid",
					Value: "Valid",
				},
				{
					Label: "invalid",
				},
				{
					Value: "invalid",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}
