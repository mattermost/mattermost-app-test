package channel_header

import "github.com/mattermost/mattermost-plugin-apps/apps"

const icon = "https://icons.iconarchive.com/icons/icons8/ios7/128/Science-Test-Tube-icon.png"
const svgIcon = "https://upload.wikimedia.org/wikipedia/commons/2/21/Speaker_Icon.svg"

func Get() *apps.Binding {
	out := &apps.Binding{
		Location: apps.LocationChannelHeader,
		Bindings: []*apps.Binding{},
	}

	out.Bindings = append(out.Bindings, getValid()...)
	out.Bindings = append(out.Bindings, getInvalid()...)
	out.Bindings = append(out.Bindings, getError()...)
	return out
}
