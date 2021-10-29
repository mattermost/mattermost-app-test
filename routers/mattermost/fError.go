package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/utils"
)

func fError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallErrorResponse(w, "Error")
}

func fEmptyError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallErrorResponse(w, "")
}

func fMarkdownFormError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallResponse(w, apps.CallResponse{
		Type:      apps.CallResponseTypeError,
		ErrorText: "## This is a very **BIG** error.\nYou should probably take a look at it.",
		Data: map[string]map[string]string{
			"errors": {
				"text":    "These are not the emojis you are looking for :sweat_smile:",
				"boolean": "Are you sure you should _mark_ this field?",
				"static":  "## Careful\nThis is an error.",
				"missing": "Some missing field.",
			},
		},
	})
}

func fMarkdownFormErrorMissingField(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeError,
		Data: map[string]map[string]string{
			"errors": {
				"missing": "Some missing field.",
			},
		},
	})
}
