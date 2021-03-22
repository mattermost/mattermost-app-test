package channelheader

import "github.com/mattermost/mattermost-plugin-apps/apps"

func Get(siteURL, appID string) *apps.Binding {
	out := &apps.Binding{
		Location: apps.LocationChannelHeader,
		Bindings: []*apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getInvalid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getError(siteURL, appID)...)

	return out
}
