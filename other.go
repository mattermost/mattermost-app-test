package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-app-test/path"
)

func otherCommandBinding(_ apps.Context) apps.Binding {
	b := newBinding("dialog_from_interactive_message", path.CreateDialogMessage)
	b.Submit = b.Submit.WithExpand(apps.Expand{ActingUserAccessToken: apps.ExpandAll})

	return apps.Binding{
		Label: "other",
		Bindings: []apps.Binding{
			b,
		},
	}
}

func initHTTPOther(r *mux.Router) {
	handleCall(r, path.CreateDialogMessage, handleCreateDialogMessage)

	r.HandleFunc(path.CreateDialogMessage+path.DialogNoResponse, postOpenDialogTestNoResponse)
	r.HandleFunc(path.CreateDialogMessage+path.DialogEmptyResponse, postOpenDialogTestEmptyResponse)
	r.HandleFunc(path.CreateDialogMessage+path.DialogEphemeralResponse, postOpenDialogTestEphemeralResponse)
	r.HandleFunc(path.CreateDialogMessage+path.DialogUpdateResponse, postOpenDialogTestUpdateResponse)
	r.HandleFunc(path.CreateDialogMessage+path.DialogBadResponse, postOpenDialogTestBadResponse)
}

func handleCreateDialogMessage(creq *apps.CallRequest) apps.CallResponse {
	url := manifest.HTTP.RootURL
	cc := creq.Context
	client := appclient.AsActingUser(cc)
	post := &model.Post{
		ChannelId: cc.ChannelID,
		Message:   "TEST",
	}

	model.ParseSlackAttachment(post, []*model.SlackAttachment{
		{
			Text: "Select an action, and a dialog should open",
			Actions: []*model.PostAction{
				{
					Type: "button",
					Name: "no response",
					Integration: &model.PostActionIntegration{
						URL: url + path.CreateDialogMessage + path.DialogNoResponse,
						Context: map[string]interface{}{
							"url": cc.MattermostSiteURL,
						},
					},
				},
				{
					Type: "button",
					Name: "empty response",
					Integration: &model.PostActionIntegration{
						URL: url + path.CreateDialogMessage + path.DialogEmptyResponse,
						Context: map[string]interface{}{
							"url": cc.MattermostSiteURL,
						},
					},
				},
				{
					Type: "button",
					Name: "ephemeral response",
					Integration: &model.PostActionIntegration{
						URL: url + path.CreateDialogMessage + path.DialogEphemeralResponse,
						Context: map[string]interface{}{
							"url": cc.MattermostSiteURL,
						},
					},
				},
				{
					Type: "button",
					Name: "update response",
					Integration: &model.PostActionIntegration{
						URL: url + path.CreateDialogMessage + path.DialogUpdateResponse,
						Context: map[string]interface{}{
							"url": cc.MattermostSiteURL,
						},
					},
				},
				{
					Type: "button",
					Name: "not recognized response (doesn't open dialog)",
					Integration: &model.PostActionIntegration{
						URL: url + path.CreateDialogMessage + path.DialogBadResponse,
						Context: map[string]interface{}{
							"url": cc.MattermostSiteURL,
						},
					},
				},
			},
		},
	})

	post, err := client.CreatePost(post)
	if err != nil {
		return apps.NewErrorResponse(err)
	}

	return apps.NewTextResponse(utils.Pretty(post))
}

func postOpenDialogTestNoResponse(w http.ResponseWriter, r *http.Request) {
	var req model.PostActionIntegrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, _ = client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
}

func postOpenDialogTestEmptyResponse(w http.ResponseWriter, r *http.Request) {
	var req model.PostActionIntegrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, _ = client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
			Elements: []model.DialogElement{
				{
					DisplayName: "foo",
					Name:        "foo",
					Type:        "text",
				},
			},
		},
	})

	// utils.DumpObject(resp)

	_, _ = w.Write([]byte("{}"))
}

func postOpenDialogTestEphemeralResponse(w http.ResponseWriter, r *http.Request) {
	var req model.PostActionIntegrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, _ = client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})

	// utils.DumpObject(resp)

	b, _ := json.Marshal(model.PostActionIntegrationResponse{
		EphemeralText: "Test ephemeral",
	})
	_, _ = w.Write(b)
}

func postOpenDialogTestUpdateResponse(w http.ResponseWriter, r *http.Request) {
	var req model.PostActionIntegrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, _ = client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})

	// utils.DumpObject(resp)

	b, _ := json.Marshal(model.PostActionIntegrationResponse{
		Update: &model.Post{
			Message: "Updated!",
		},
	})
	_, _ = w.Write(b)
}

func postOpenDialogTestBadResponse(w http.ResponseWriter, r *http.Request) {
	var req model.PostActionIntegrationRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, _ = client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})

	// utils.DumpObject(resp)

	_, _ = w.Write([]byte("ABCDEFG"))
}
