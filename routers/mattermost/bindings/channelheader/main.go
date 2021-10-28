package channelheader

import "github.com/mattermost/mattermost-plugin-apps/apps"

func Get(context apps.Context) apps.Binding {
	siteURL := context.MattermostSiteURL
	appID := string(context.AppID)
	out := apps.Binding{
		Location: apps.LocationChannelHeader,
		Bindings: []apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getInvalid(siteURL, appID)...)
	out.Bindings = append(out.Bindings, getError(siteURL, appID)...)

	return out
}
