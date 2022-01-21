package main

import (
	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func commandBindings(cc apps.Context) []apps.Binding {
	b := apps.Binding{
		Label: CommandTrigger,
		Icon:  "icon.png",
		Bindings: []apps.Binding{
			embeddedCommandBinding(cc),
			formCommandBinding(cc),
			otherCommandBinding(cc),
			subscribtionCommandBinding("subscribe", path.Subscribe),
			subscribtionCommandBinding("unsubscribe", path.Unsubscribe),

			testCommandBinding(cc),
		},
	}

	return []apps.Binding{b}
}

func testCommandBinding(cc apps.Context) apps.Binding {
	out := []apps.Binding{
		validResponseBinding,
		errorResponseBinding,
		validInputBinding,
	}

	if includeInvalid {
		out = append(out,
			invalidResponseBinding,
			invalidBindingBinding,
			invalidFormBinding,
		)
	}

	if cc.Channel != nil && cc.Channel.Name == "town-square" {
		out = append([]apps.Binding{
			newBinding("town-square-channel-specific", path.OK),
		}, out...)
	}

	return apps.Binding{
		Label:    "test-command",
		Bindings: out,
	}
}
