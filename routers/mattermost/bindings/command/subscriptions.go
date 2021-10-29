package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func makeSubscriptionOption(subject apps.Subject) apps.SelectOption {
	return apps.SelectOption{
		Label: string(subject),
		Value: string(subject),
	}
}

func getSubscribeCommand(_ *apps.Context) *apps.Binding {
	return &apps.Binding{
		Location: "subscriptions",
		Label:    "subscriptions",
		Form: &apps.Form{
			Call: &apps.Call{
				Path:  constants.SubscribeCommand,
				State: true,
				Expand: &apps.Expand{
					AdminAccessToken: apps.ExpandAll,
				},
			},
			Fields: []*apps.Field{
				{
					Name:                 "subject",
					Label:                "subject",
					IsRequired:           true,
					AutocompletePosition: 1,
					Type:                 apps.FieldTypeStaticSelect,
					SelectStaticOptions: []apps.SelectOption{
						makeSubscriptionOption(apps.SubjectBotMentioned),
						makeSubscriptionOption(apps.SubjectBotJoinedChannel),
						makeSubscriptionOption(apps.SubjectBotLeftChannel),
						makeSubscriptionOption(apps.SubjectBotJoinedTeam),
						makeSubscriptionOption(apps.SubjectBotLeftTeam),
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
