package command

import "github.com/mattermost/mattermost-plugin-apps/apps"

func Get() *apps.Binding {
	out := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid())
	out.Bindings = append(out.Bindings, getInvalid())
	out.Bindings = append(out.Bindings, getError())
	return out
}
