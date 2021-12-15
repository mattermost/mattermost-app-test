package mattermost

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-app-test/utils"
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
		client := appclient.AsActingUser(context)

		values, err := unmarshalSubscriptionsCommandFormValues(c.Values)
		if err != nil {
			utils.WriteCallErrorResponse(w, errors.Wrap(err, "failed to unmarshal command submission into form values").Error())
			return
		}

		subject := values.Subject.Value

		path := "/notify/" + string(subject)
		sub := &apps.Subscription{
			Subject: subject,
			AppID:   context.AppID,
			Call:    apps.NewCall(path),
		}

		switch subject {
		case apps.SubjectUserJoinedChannel,
			apps.SubjectUserLeftChannel,
			apps.SubjectPostCreated:
			sub.ChannelID = context.ChannelID
		case apps.SubjectUserJoinedTeam,
			apps.SubjectUserLeftTeam,
			apps.SubjectChannelCreated:
			sub.TeamID = context.TeamID
		}

		message := ""
		subscribe := values.Subscribe.Value == "subscribe"

		if subscribe {
			err := client.Subscribe(sub)
			if err != nil {
				err = errors.Wrapf(err, "failed to subscribe to `%v` notifications.", subject)
				b, _ := json.Marshal(c)

				err = errors.Wrap(err, string(b))
				cr := apps.NewErrorResponse(err)
				_ = json.NewEncoder(w).Encode(cr)

				return
			}

			message = fmt.Sprintf("Successfully subscribed to `%v` notifications.", subject)
		} else {
			err := client.Unsubscribe(sub)
			if err != nil {
				err = errors.Wrapf(err, "failed to unsubscribe from `%v` notifications.", subject)
				cr := apps.NewErrorResponse(err)
				_ = json.NewEncoder(w).Encode(cr)

				return
			}

			message = fmt.Sprintf("Successfully unsubscribed from `%v` notifications.", subject)
		}

		s := fmt.Sprintf("```\n%v\n```\n%v", c.RawCommand, message)
		utils.WriteCallStandardResponse(w, s)
	}
}

func fSubscriptionsUserCreated(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotMention(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

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
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

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
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsBotJoinedTeam(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		channel, _, err := client.GetChannelByName("town-square", c.Context.TeamID, "")
		if err != nil {
			log.Println(err.Error())
			message += ", but failed to get Town Square channel: " + err.Error()

			_, err = client.DM(c.Context.ActingUser.Id, message)
			if err != nil {
				log.Println(err.Error())
			}

			utils.WriteCallStandardResponse(w, message)

			return
		}

		_, err = client.CreatePost(&model.Post{
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
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsPostCreated(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

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

func fSubscriptionsUserJoinedChannel(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

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

func fSubscriptionsUserLeftChannel(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

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

func fSubscriptionsUserJoinedTeam(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsUserLeftTeam(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}

func fSubscriptionsChannelCreated(m *apps.Manifest) func(http.ResponseWriter, *http.Request, *apps.CallRequest) {
	return func(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
		message := fmt.Sprintf("Received `%s` notification", c.Context.Subject)
		client := appclient.AsBot(c.Context)

		_, err := client.DM(c.Context.ActingUser.Id, message)
		if err != nil {
			log.Println(err.Error())
		}

		utils.WriteCallStandardResponse(w, message)
	}
}
