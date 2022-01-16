package main

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

func testCommandBinding(cc apps.Context) apps.Binding {
	var out []apps.Binding
	if cc.Channel != nil && cc.Channel.Name == "town-square" {
		out = append(out, newBinding("town-square-channel-specific", path.OK))
	}

	out = append(out, apps.Binding{
		Label: "valid-response",
		Bindings: []apps.Binding{
			newBinding("OK", path.OK),
			newBinding("OK-empty", path.OKEmpty),
			newBinding("form", path.FormSimple),
			newBinding("form-source", path.FormSimpleSource), // TODO <>/<> does not work, move to invalid?
			newBinding("navigate-external", path.NavigateExternal),
			newBinding("navigate-internal", path.NavigateInternal),
		},
	})

	out = append(out, apps.Binding{
		Label: "error-response",
		Bindings: []apps.Binding{
			newBinding("error", path.ErrorDefault),
			newBinding("error-empty", path.ErrorEmpty),
			newBinding("error-404", path.Error404),
			newBinding("error-500", path.Error500),
		},
	})

	out = append(out, apps.Binding{
		Label: "with-form",
		Bindings: []apps.Binding{
			{
				Label: "empty",
				Form: &apps.Form{
					Submit: callOK,
				},
			},
			{
				Label: "simple",
				Form:  &simpleForm,
			},
			{
				// TODO <>/<> does not work, move to invalid?
				Label: "simple-source",
				Form:  &simpleFormSource,
			},
		},
	})

	return apps.Binding{
		Label:    "test-command",
		Bindings: out,
	}
}
