package mattermost

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
	"github.com/mattermost/mattermost-server/v5/model"
)

func fUnknown(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: "unknown",
	}
	utils.WriteCallResponse(w, resp)
}

func fHTML(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	html := `
		<!DOCTYPE html>
		<html>
			<head>
			</head>
			<body>
				<p>HTML example</p>
			</body>
		</html>
		`

	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(html))
}

func postOpenDialogTest(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		url := m.HTTPRootURL
		context := c.Context
		client := mmclient.AsActingUser(context)
		post := &model.Post{
			ChannelId: context.ChannelID,
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
							URL: url + constants.OtherPathOpenDialog + constants.OtherOpenDialogNoResponse,
							Context: map[string]interface{}{
								"url": context.MattermostSiteURL,
							},
						},
					},
					{
						Type: "button",
						Name: "empty response",
						Integration: &model.PostActionIntegration{
							URL: url + constants.OtherPathOpenDialog + constants.OtherOpenDialogEmptyResponse,
							Context: map[string]interface{}{
								"url": context.MattermostSiteURL,
							},
						},
					},
					{
						Type: "button",
						Name: "ephemeral response",
						Integration: &model.PostActionIntegration{
							URL: url + constants.OtherPathOpenDialog + constants.OtherOpenDialogEphemeralResponse,
							Context: map[string]interface{}{
								"url": context.MattermostSiteURL,
							},
						},
					},
					{
						Type: "button",
						Name: "update response",
						Integration: &model.PostActionIntegration{
							URL: url + constants.OtherPathOpenDialog + constants.OtherOpenDialogUpdateResponse,
							Context: map[string]interface{}{
								"url": context.MattermostSiteURL,
							},
						},
					},
					{
						Type: "button",
						Name: "not recognized response (doesn't open dialog)",
						Integration: &model.PostActionIntegration{
							URL: url + constants.OtherPathOpenDialog + constants.OtherOpenDialogBadResponse,
							Context: map[string]interface{}{
								"url": context.MattermostSiteURL,
							},
						},
					},
				},
			},
		})
		client.CreatePost(post)

		utils.WriteCallStandardResponse(w, "")
	}
}

func postOpenDialogTestNoResponse(w http.ResponseWriter, r *http.Request) {
	req := model.PostActionIntegrationRequestFromJson(r.Body)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, resp := client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
	utils.DumpObject(resp)
}

func postOpenDialogTestEmptyResponse(w http.ResponseWriter, r *http.Request) {
	req := model.PostActionIntegrationRequestFromJson(r.Body)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, resp := client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
	utils.DumpObject(resp)
	w.Write([]byte("{}"))
}

func postOpenDialogTestEphemeralResponse(w http.ResponseWriter, r *http.Request) {
	req := model.PostActionIntegrationRequestFromJson(r.Body)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, resp := client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
	utils.DumpObject(resp)
	b, _ := json.Marshal(model.PostActionIntegrationResponse{
		EphemeralText: "Test ephemeral",
	})
	w.Write(b)
}

func postOpenDialogTestUpdateResponse(w http.ResponseWriter, r *http.Request) {
	req := model.PostActionIntegrationRequestFromJson(r.Body)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, resp := client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
	utils.DumpObject(resp)
	b, _ := json.Marshal(model.PostActionIntegrationResponse{
		Update: &model.Post{
			Message: "Updated!",
		},
	})
	w.Write(b)
}

func postOpenDialogTestBadResponse(w http.ResponseWriter, r *http.Request) {
	req := model.PostActionIntegrationRequestFromJson(r.Body)
	url := req.Context["url"].(string)
	client := model.NewAPIv4Client(url)
	_, resp := client.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: req.TriggerId,
		URL:       url,
		Dialog: model.Dialog{
			Title:            "Test",
			IntroductionText: "Do not submit this dialog, it will fail.",
		},
	})
	utils.DumpObject(resp)
	w.Write([]byte("ABCDEFG"))
}
