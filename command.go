package main

import (
	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func commandBindings(cc apps.Context) []apps.Binding {
	b := apps.Binding{
		Label: CommandTrigger,
		Icon:  "icon.png",
		Bindings: []apps.Binding{
			embeddedCommandBinding(cc),
			formCommandBinding(cc),
			otherCommandBinding(cc),
			subscribtionCommandBinding("subscribe", path.Subscribe),
			subscribtionCommandBinding("unsubscribe", path.Unsubscribe),
			testCommandBinding(cc),
		},
	}

	return []apps.Binding{b}
}

// var channelHeaderBindings = []apps.Binding{
// 	// valid bindings and responses
// 	newBinding("submit return ok", constants.Submit),
// 	newBinding("submit return ok empty", constants.SubmitEmpty),
// 	newBinding("submit return simple form", constants.Form),
// 	newBinding("submit return full form", constants.FormFull),
// 	newBinding("submit return dynamic form", constants.FormDynamic),
// 	newBinding("submit return internal navigate", constants.NavigateInternal),
// 	newBinding("submit return external navigate", constants.NavigateExternal),
// 	newFormBinding("attached simple form", simpleForm),
// 	newFormBinding("attached form source simple", simpleSourceForm),
// 	newFormBinding("attached full form", fullForm),
// 	newFormBinding("attached form source full", fullSourceForm),

// 	// error bindings
// 	newBinding("submit error", constants.Error),
// 	newBinding("submit error empty ", constants.ErrorEmpty),
// 	newBinding("submit error invalid navigate", constants.NavigateInvalid),
// 	newBinding("submit error 404 response", constants.NotFoundPath),
// 	newBinding("submit error HTML response", constants.HTMLPath),
// 	newBinding("submit error JSON response", constants.ManifestPath),
// 	newBinding("submit error unknown response type", constants.UnknownPath),
// 	apps.Binding{
// 		Label:  "icon not found",
// 		Icon:   "foo",
// 		Submit: apps.NewCall(constants.Submit),
// 	},
// 	apps.Binding{
// 		Label:  "icon is SVG",
// 		Icon:   "icon.svg",
// 		Submit: apps.NewCall(constants.Submit),
// 	},

// 	// invalid bindings
// 	newFormBinding("invalid attached form", apps.Form{}),
// 	newFormBinding("invalid form in source", invalidSourceForm),
// 	apps.Binding{
// 		Label:  "invalid no icon",
// 		Submit: apps.NewCall(constants.Submit),
// 	},
// 	apps.Binding{
// 		Location: "invalid no label",
// 		Icon:     "icon.png",
// 		Submit:   apps.NewCall(constants.Submit),
// 	},
// 	apps.Binding{
// 		Label: "invalid neither submit nor form",
// 		Icon:  "icon.png",
// 	},
// 	apps.Binding{
// 		Location: "invalid whitespace label",
// 		Label:    " ",
// 		Icon:     "icon.png",
// 		Submit:   apps.NewCall(constants.Submit),
// 	},
// }
