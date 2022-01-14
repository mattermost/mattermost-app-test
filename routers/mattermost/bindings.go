package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
)

func commandBindings(context apps.Context) []apps.Binding {
	var out []apps.Binding
	if context.Channel.Name == "town-square" {
		out = append(out, new
			Label:    "town_square",
Submit: apps. constants.SubmitOK,
			},
		})
	}

	base.Bindings = append(base.Bindings, getValid())
	base.Bindings = append(base.Bindings, getInvalid())
	base.Bindings = append(base.Bindings, getError())
	base.Bindings = append(base.Bindings, getOthers(context))
	base.Bindings = append(base.Bindings, getSubscribeCommand(context))


	return out
}

var channelHeaderBindings = []apps.Binding{
	// valid bindings and responses
	newBinding("submit return ok", constants.Submit),
	newBinding("submit return ok empty", constants.SubmitEmpty),
	newBinding("submit return simple form", constants.Form),
	newBinding("submit return full form", constants.FormFull),
	newBinding("submit return dynamic form", constants.FormDynamic),
	newBinding("submit return internal navigate", constants.NavigateInternal),
	newBinding("submit return external navigate", constants.NavigateExternal),
	newFormBinding("attached simple form", simpleForm),
	newFormBinding("attached form source simple", simpleSourceForm),
	newFormBinding("attached full form", fullForm),
	newFormBinding("attached form source full", fullSourceForm),

	// error bindings
	newBinding("submit error", constants.Error),
	newBinding("submit error empty ", constants.ErrorEmpty),
	newBinding("submit error invalid navigate", constants.NavigateInvalid),
	newBinding("submit error 404 response", constants.NotFoundPath),
	newBinding("submit error HTML response", constants.HTMLPath),
	newBinding("submit error JSON response", constants.ManifestPath),
	newBinding("submit error unknown response type", constants.UnknownPath),
	apps.Binding{
		Label:  "icon not found",
		Icon:   "foo",
		Submit: apps.NewCall(constants.Submit),
	},
	apps.Binding{
		Label:  "icon is SVG",
		Icon:   "icon.svg",
		Submit: apps.NewCall(constants.Submit),
	},

	// invalid bindings
	newFormBinding("invalid attached form", apps.Form{}),
	newFormBinding("invalid form in source", invalidSourceForm),
	apps.Binding{
		Label:  "invalid no icon",
		Submit: apps.NewCall(constants.Submit),
	},
	apps.Binding{
		Location: "invalid no label",
		Icon:     "icon.png",
		Submit:   apps.NewCall(constants.Submit),
	},
	apps.Binding{
		Label: "invalid neither submit nor form",
		Icon:  "icon.png",
	},
	apps.Binding{
		Location: "invalid whitespace label",
		Label:    " ",
		Icon:     "icon.png",
		Submit:   apps.NewCall(constants.Submit),
	},
}

func fBindings(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	out := append([]apps.Binding{
		{
			Location: apps.LocationChannelHeader,
			Bindings: channelHeaderBindings,
		},
		{
			Location: apps.LocationCommand,
			Bindings: commandBindings(c.Context),
		},
	})
	utils.WriteBindings(w, out)
}

func newBareBinding(label string) apps.Binding {
	return apps.Binding{
		Label: label,
		Icon:  "icon.png",
	}
}

func newBinding(label, submitPath string) apps.Binding {
	return apps.Binding{
		Label:  label,
		Icon:   "icon.png",
		Submit: apps.NewCall(submitPath),
	}
}

func newFormBinding(label string, form apps.Form) apps.Binding {
	return apps.Binding{
		Label: label,
		Icon:  "icon.png",
		Form:  &form,
	}
}
