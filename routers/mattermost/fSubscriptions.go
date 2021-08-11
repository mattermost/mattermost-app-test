package mattermost

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

type SubscriptionsCommandFormValues struct {
	Subject struct {
		Label apps.Subject `json:"label"`
		Value apps.Subject `json:"value"`
	} `json:"subject"`

	Subscribe struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"subscribe"`
}

func unmarshalSubscriptionsCommandFormValues(form map[string]interface{}) (*SubscriptionsCommandFormValues, error) {
	b, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}

	values := &SubscriptionsCommandFormValues{}

	err = json.Unmarshal(b, values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func fSubscriptionsCommand(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		context := c.Context
		client := mmclient.AsAdmin(context)

		values, err := unmarshalSubscriptionsCommandFormValues(c.Values)
		if err != nil {
			utils.WriteCallErrorResponse(w, errors.Wrap(err, "failed to unmarshal command submission into form values").Error())
			return
		}

		path := "/notify/" + string(values.Subject.Value)
		sub := &apps.Subscription{
			Subject: values.Subject.Value,
			AppID:   context.AppID,
			Call:    apps.NewCall(path),
		}

		message := ""
		subscribe := values.Subscribe.Value == "subscribe"

		if subscribe {
			_, err := client.Subscribe(sub)
			if err != nil {
				err = errors.Wrapf(err, "failed to subscribe to %v notifications.", values.Subject)
				b, _ := json.Marshal(c)

				err = errors.Wrap(err, string(b))
				cr := apps.NewErrorCallResponse(err)
				json.NewEncoder(w).Encode(cr)

				return
			}

			message = fmt.Sprintf("Successfully subscribed to %v notifications.", values.Subject)
		} else {
			_, err := client.Unsubscribe(sub)
			if err != nil {
				err = errors.Wrapf(err, "failed to unsubscribe from %v notifications.", values.Subject)
				cr := apps.NewErrorCallResponse(err)
				json.NewEncoder(w).Encode(cr)

				return
			}

			message = fmt.Sprintf("Successfully unsubscribed from %v notifications.", values.Subject)
		}

		if values.Subject.Value == apps.SubjectBotMentioned {
			member, _ := client.GetChannelMember(context.ChannelID, context.BotUserID, "")
			if member == nil {
				message += " I'm not a member of this channel so I won't be notified of posts mentioning me here."
			} else {
				message += " I'm a member of this channel so I will be notified of posts mentioning me here."
			}
		}

		s := fmt.Sprintf("```\n%v\n```\n%v", c.RawCommand, message)
		utils.WriteCallStandardResponse(w, s)
	}
}

func fSubscriptionsBotMention(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := "Received `bot_mentioned` notification"
		client := mmclient.AsBot(c.Context)

		rootID := c.Context.PostID
		if c.Context.RootPostID != "" {
			rootID = c.Context.RootPostID
		}

		_, err := client.CreatePost(&model.Post{
			ChannelId: c.Context.ChannelID,
			RootId:    rootID,
			Message:   message,
		})

		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotJoinedChannel(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := "Received `bot_joined_channel` notification"
		client := mmclient.AsBot(c.Context)

		_, err := client.CreatePost(&model.Post{
			ChannelId: c.Context.ChannelID,
			Message:   message,
		})

		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotLeftChannel(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := "Received `bot_left_channel` notification"
		client := mmclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUserID, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotJoinedTeam(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := "Received `bot_joined_team` notification"
		client := mmclient.AsBot(c.Context)

		channel, resp := client.GetChannelByName("town-square", c.Context.TeamID, "")
		if resp.Error != nil {
			log.Println(resp.Error.Error())
			message += ", but failed to get Town Square channel: " + resp.Error.Error()

			_, err := client.DM(c.Context.ActingUserID, message)
			if err != nil {
				log.Println(err.Error())
			}

			utils.WriteCallStandardResponse(w, message)

			return
		}

		_, err := client.CreatePost(&model.Post{
			ChannelId: channel.Id,
			Message:   message,
		})

		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotLeftTeam(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := "Received `bot_left_team` notification"
		client := mmclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUserID, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}
