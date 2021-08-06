package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getSubscribeCommands(context *apps.Context) *apps.Binding {
	base := &apps.Binding{
		Location: "subscriptions",
		Label:    "subscriptions",
		Bindings: []*apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getBotMentionCommand())

	return base
}

func getBotMentionCommand() *apps.Binding {
	return &apps.Binding{
		Location: "bot_mention",
		Label:    "bot_mention",
		Bindings: []*apps.Binding{
			{
				Location:    "subscribe",
				Label:       "subscribe",
				Description: "Subscribe to bot mention subscriptions",
				Form:        &apps.Form{},
				Call: &apps.Call{
					Path:  constants.SubscribeBotMention,
					State: true,
					Expand: &apps.Expand{
						AdminAccessToken: apps.ExpandAll,
					},
				},
			},
			{
				Location:    "unsubscribe",
				Label:       "unsubscribe",
				Description: "Unsubscribe from bot mention subscriptions",
				Form:        &apps.Form{},
				Call: &apps.Call{
					Path:  constants.SubscribeBotMention,
					State: false,
					Expand: &apps.Expand{
						AdminAccessToken: apps.ExpandAll,
					},
				},
			},
		},
	}
}
