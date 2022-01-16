package main

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

func handleError(text string) http.HandlerFunc {
	return httputils.DoHandleJSON(apps.CallResponse{
		Type: apps.CallResponseTypeError,
		Text: text,
	})
}

func handleErrorWithData(text string, data interface{}) http.HandlerFunc {
	return httputils.DoHandleJSON(apps.CallResponse{
		Type: apps.CallResponseTypeError,
		Text: text,
		Data: data,
	})
}
