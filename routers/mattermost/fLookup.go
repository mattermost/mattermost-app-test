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
					Label: "static value 1",
					Value: "sv1",
				},
				{
					Label: "static value 2",
					Value: "sv2",
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
