package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

func channelHeaderBindings(cc apps.Context) []apps.Binding {
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
