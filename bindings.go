package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func handleBindings(c *apps.CallRequest) apps.CallResponse {
	return apps.NewDataResponse([]apps.Binding{
		// {
		// 	Location: apps.LocationChannelHeader,
		// 	Bindings: channelHeaderBindings,
		// },
		{
			Location: apps.LocationCommand,
			Bindings: commandBindings(c.Context),
		},
	})
}

func newBareBinding(label string) apps.Binding {
	return apps.Binding{
		Label: label,
		Icon:  "icon.png",
	}
}

func newBinding(label, submitPath string) apps.Binding {
	return apps.Binding{
		Label:  label,
		Icon:   "icon.png",
		Submit: apps.NewCall(submitPath),
	}
}

func newFormBinding(label string, form apps.Form) apps.Binding {
	return apps.Binding{
		Label: label,
		Icon:  "icon.png",
		Form:  &form,
	}
}