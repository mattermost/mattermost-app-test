package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

var responseOK = apps.CallResponse{
	Type: apps.CallResponseTypeOK,
	Data: "any data OK",
	Text: "OK",
}

var responseOKEmpty = apps.CallResponse{
	Type: apps.CallResponseTypeOK,
}

func handleOK(_ *apps.CallRequest) apps.CallResponse {
	return responseOK
}
func handleOKEmpty(_ *apps.CallRequest) apps.CallResponse {
	return responseOKEmpty
}
