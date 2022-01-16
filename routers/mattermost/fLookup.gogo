package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

type lookupResponse struct {
	Items []apps.SelectOption `json:"items"`
}

func fLookupOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{
				{
					Label: "dynamic value 1 label",
					Value: "sv1",
				},
				{
					Label: "dynamic value 2 label",
					Value: "sv2",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupMultiword(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{
				{
					Label: "dynamic value 2 label",
					Value: "dynamic value 1",
				},
				{
					Label: "dynamic value 2 label",
					Value: "dynamic value 2",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupEmpty(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: lookupResponse{
			Items: []apps.SelectOption{},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fLookupInvalid(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
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
