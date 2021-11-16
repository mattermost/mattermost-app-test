package channelheader

import "github.com/mattermost/mattermost-plugin-apps/apps"

func Get(context apps.Context) apps.Binding {
	out := apps.Binding{
		Location: apps.LocationChannelHeader,
		Bindings: []apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid()...)
	out.Bindings = append(out.Bindings, getInvalid()...)
	out.Bindings = append(out.Bindings, getError()...)

	return out
}
