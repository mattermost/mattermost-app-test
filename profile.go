package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"

	"github.com/mattermost/mattermost-app-test/path"
)

var (
	profileViewCall    = apps.NewCall(path.ProfileView).WithExpand(apps.Expand{ActingUserAccessToken: apps.ExpandAll})
	profileCommandCall = apps.NewCall(path.ProfileCommand).WithExpand(apps.Expand{ActingUserAccessToken: apps.ExpandAll})
)

func profileCommandBinding(context apps.Context) apps.Binding {
	return apps.Binding{
		Label: "profile",
		Bindings: []apps.Binding{
			{
				Label:  "view",
				Submit: profileViewCall,
			},
			{
				Label: "command",
				Form: &apps.Form{
					Fields: []apps.Field{
						{
							Name:                 "name",
							AutocompletePosition: 1,
							IsRequired:           true,
							Type:                 apps.FieldTypeStaticSelect,
							SelectStaticOptions: []apps.SelectOption{
								{
									Label: "simple",
									Value: "simple",
								},
							},
						},
					},
					Submit: profileCommandCall,
				},
			},
		},
	}
}

func handleProfileView(creq *apps.CallRequest) apps.CallResponse {
	commandProfile := ""
	client := appclient.AsBot(creq.Context)
	err := client.KVGet("test_command_profile", "", &commandProfile)
	if err != nil {
		return apps.NewErrorResponse(err)
	}
	if commandProfile == "" {
		commandProfile = "none"
	} else {
		commandProfile = "`" + commandProfile + "`"
	}
	return apps.NewTextResponse("Selected profiles:\n - Command: %s\n", commandProfile)
}

func handleProfileCommand(creq *apps.CallRequest) apps.CallResponse {
	name := creq.GetValue("name", "")
	client := appclient.AsBot(creq.Context)
	_, err := client.KVSet("test_command_profile", "", name)
	if err != nil {
		return apps.NewErrorResponse(err)
	}

	return apps.NewTextResponse("Selected profile `%s` for test command, please refresh bindings", name)
}
