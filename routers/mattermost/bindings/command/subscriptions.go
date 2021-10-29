package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func makeSubscriptionOption(subject apps.Subject) apps.SelectOption {
	return apps.SelectOption{
		Label: string(subject),
		Value: string(subject),
	}
}

func getSubscribeCommand(context apps.Context) apps.Binding {
	return apps.Binding{
		Location: "subscriptions",
		Label:    "subscriptions",
		Call: &apps.Call{
			Path:  constants.SubscribeCommand,
			State: true,
			Expand: &apps.Expand{
				ActingUserAccessToken: apps.ExpandAll,
			},
		},
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:                 "subject",
					Label:                "subject",
					IsRequired:           true,
					AutocompletePosition: 1,
					Type:                 apps.FieldTypeStaticSelect,
					SelectStaticOptions: []apps.SelectOption{
						makeSubscriptionOption(apps.SubjectUserCreated),
						makeSubscriptionOption(apps.SubjectBotJoinedChannel),
						makeSubscriptionOption(apps.SubjectBotLeftChannel),
						makeSubscriptionOption(apps.SubjectBotJoinedTeam),
						makeSubscriptionOption(apps.SubjectBotLeftTeam),
						makeSubscriptionOption(apps.SubjectBotMentioned),
						makeSubscriptionOption(apps.SubjectUserJoinedChannel),
						makeSubscriptionOption(apps.SubjectUserLeftChannel),
						makeSubscriptionOption(apps.SubjectPostCreated),
						makeSubscriptionOption(apps.SubjectUserJoinedTeam),
						makeSubscriptionOption(apps.SubjectUserLeftTeam),
						makeSubscriptionOption(apps.SubjectChannelCreated),
					},
				},
				{
					Name:                 "subscribe",
					Label:                "subscribe",
					IsRequired:           true,
					AutocompletePosition: 2,
					Type:                 apps.FieldTypeStaticSelect,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "subscribe",
							Value: "subscribe",
						},
						{
							Label: "unsubscribe",
							Value: "unsubscribe",
						},
					},
				},
			},
		},
	}
}
