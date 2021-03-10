package bindings

import (
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/channelheader"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/command"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/postaction"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(siteURL string) []*apps.Binding {
	out := []*apps.Binding{}

	out = append(out, command.Get(siteURL))
	out = append(out, channelheader.Get(siteURL))
	out = append(out, postaction.Get(siteURL))

	return out
}
