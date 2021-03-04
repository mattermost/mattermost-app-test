package bindings

import (
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/channel_header"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/command"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/post_action"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get() []*apps.Binding {
	out := []*apps.Binding{}

	out = append(out, command.Get())
	out = append(out, channel_header.Get())
	out = append(out, post_action.Get())
	return out
}
