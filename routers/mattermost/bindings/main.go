package bindings

import (
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/channelheader"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/command"
	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings/postaction"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(siteURL, appID string, args ...string) []*apps.Binding {
	out := []*apps.Binding{}

	out = append(out, command.Get(siteURL, appID, args...))
	out = append(out, channelheader.Get(siteURL, appID))
	out = append(out, postaction.Get(siteURL, appID))

	return out
}
