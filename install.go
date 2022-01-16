package main

import (
	"log"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func handleInstall(c *apps.CallRequest) apps.CallResponse {
	log.Println("handleInstall called")
	return apps.NewTextResponse("handleInstall called")
}
