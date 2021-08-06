package mattermost

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

func fSubscriptionsCommandBotMention(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		context := c.Context
		client := mmclient.AsAdmin(context)

		subscribe, _ := c.State.(bool)

		message := ""
		sub := &apps.Subscription{
			Subject: apps.SubjectBotMentioned,
			AppID:   context.AppID,
			Call:    apps.NewCall("/notify/bot_mention"),
		}

		if subscribe {
			_, err := client.Subscribe(sub)
			if err != nil {
				err = errors.Wrap(err, "failed to subscribe to bot_mention notifications.")
				b, _ := json.Marshal(c)

				err = errors.Wrap(err, string(b))
				cr := apps.NewErrorCallResponse(err)
				json.NewEncoder(w).Encode(cr)

				return
			}

			message = "Successfully subscribed to bot_mention notifications."
		} else {
			_, err := client.Unsubscribe(sub)
			if err != nil {
				err = errors.Wrap(err, "failed to unsubscribe from bot_mention notifications.")
				cr := apps.NewErrorCallResponse(err)
				json.NewEncoder(w).Encode(cr)

				return
			}

			message = "Successfully unsubscribed from bot_mention notifications."
		}

		resp := fmt.Sprintf("```\n%s\n```\n%s", c.RawCommand, message)
		utils.WriteCallStandardResponse(w, resp)
	}
}

func fSubscriptionsBotMention(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		publicResponse(c.Context, "Received bot mention")

		utils.WriteCallStandardResponse(w, "Notify response")
	}
}

func publicResponse(context *apps.Context, message string) {
	client := mmclient.AsBot(context)
	client.CreatePost(&model.Post{
		ChannelId: context.ChannelID,
		Message:   message,
	})
}
