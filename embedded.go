package main

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-server/v6/model"
)

func embeddedCommandBinding(cc apps.Context) apps.Binding {
	return apps.Binding{
		Label: "embedded",
		Bindings: []apps.Binding{
			newBinding("create", path.EmbeddedCreate),
		},
	}
}

func initHTTPEmbedded(r *mux.Router) {
	handleCall(r, path.EmbeddedCreate, handleCreateEmbedded)
}

func handleCreateEmbedded(creq *apps.CallRequest) apps.CallResponse {
	client := appclient.AsBot(creq.Context)
	p := &model.Post{
		ChannelId: creq.Context.ChannelID,
	}

	p.AddProp(apps.PropAppBindings, []apps.Binding{
		{
			Location: "embedded",
			Form: &apps.Form{
				Title:  "Test",
				Header: "Test header",
				Submit: callOK,
				Fields: []apps.Field{},
			},
			AppID:       creq.Context.AppID,
			Description: "Please fill out this form so we can get it fixed  :hammer_and_wrench:",
			Bindings: []apps.Binding{
				{
					Location: "problem",
					Submit:   callOK,
					Bindings: []apps.Binding{
						{
							Location: "hardware",
							Label:    "Hardware Failure",
						},
						{
							Location: "software",
							Label:    "Software Error",
						},
						{
							Location: "wrong",
							Label:    "Wrong Product",
						},
					},
				},
				{
					Location: "provider",
					Submit:   callOK,
					Bindings: []apps.Binding{
						{
							Location: "work",
							Label:    "Cell Phone",
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

	_, _ = client.CreatePost(p)

	return apps.NewTextResponse("")
}
