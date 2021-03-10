package postaction

import "github.com/mattermost/mattermost-plugin-apps/apps"

func Get(siteURL string) *apps.Binding {
	out := &apps.Binding{
		Location: apps.LocationPostMenu,
		Bindings: []*apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid(siteURL)...)
	out.Bindings = append(out.Bindings, getInvalid(siteURL)...)
	out.Bindings = append(out.Bindings, getError(siteURL)...)

	return out
}
