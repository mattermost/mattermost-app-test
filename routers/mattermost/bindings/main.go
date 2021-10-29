package bindings

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/channelheader"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/command"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/postaction"
)

func Get(context apps.Context) []apps.Binding {
	out := []apps.Binding{}

	out = append(out, command.Get(context))
	out = append(out, channelheader.Get(context))
	out = append(out, postaction.Get(context))

	return out
}
