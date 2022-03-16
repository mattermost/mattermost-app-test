package main

import (
	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"

	"github.com/mattermost/mattermost-app-test/path"
)

func embeddedCommandBinding(_ apps.Context) apps.Binding {
	return apps.Binding{
		Label: "embedded",
		Bindings: []apps.Binding{
			newBinding("create", path.CreateEmbedded),
		},
	}
}

func initHTTPEmbedded(r *mux.Router) {
	handleCall(r, path.CreateEmbedded, handleCreateEmbedded)
}

func handleCreateEmbedded(creq *apps.CallRequest) apps.CallResponse {
	client := appclient.AsBot(creq.Context)
	p := &model.Post{
		ChannelId: creq.Context.ChannelID,
	}

	p.AddProp(apps.PropAppBindings, []apps.Binding{
		{
			Location:    "embedded",
			AppID:       creq.Context.AppID,
			Description: "Please fill out this form so we can get it fixed  :hammer_and_wrench:",
			Bindings: []apps.Binding{
				{
					Location: "problem",
					Bindings: []apps.Binding{
						{
							Location: "hardware",
							Submit:   callOK,
							Label:    "Hardware Failure",
						},
						{
							Location: "software",
							Label:    "Software Error",
							Submit:   callOK,
						},
						{
							Location: "wrong",
							Label:    "Wrong Product",
							Submit:   callOK,
						},
					},
				},
				{
					Location: "provider",
					Bindings: []apps.Binding{
						{
							Location: "work",
							Label:    "Cell Phone",
							Submit:   callOK,
						},
					},
				},
				{
					Location: "button",
					Label:    "Submit",
					Submit:   callOK,
				},
			},
		},
	})

	_, err := client.CreatePost(p)
	if err != nil {
		return apps.NewErrorResponse(err)
	}

	return apps.NewTextResponse("")
}
