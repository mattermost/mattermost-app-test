package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func Get(siteURL, appID string, args ...string) *apps.Binding {
	base := &apps.Binding{
		Label:       constants.CommandTrigger,
		Description: "Test commands",
		Location:    constants.CommandTrigger,
		Icon:        utils.GetIconURL(siteURL, "icon.png", appID),
		Bindings:    []*apps.Binding{},
	}
	out := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{
			base,
		},
	}

	if len(args) > 0 && args[0] == "town-square" {
		base.Bindings = append(base.Bindings, &apps.Binding{
			Location: "town_square",
			Label:    "town_square",
			Form: &apps.Form{
				Fields: []*apps.Field{},
			},
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
		})
	}

	base.Bindings = append(base.Bindings, getValid(siteURL, appID))
	base.Bindings = append(base.Bindings, getInvalid(siteURL, appID))
	base.Bindings = append(base.Bindings, getError(siteURL, appID))

	return out
}

const fullMarkdown = "## Markdown title" +
	"\nHello world" +
	"\nText styles: _italics_ **bold** **_bold-italic_** ~~strikethrough~~ `code`" +
	"\nUsers and channels: @sysadmin ~town-square" +
	"\n```" +
	"\nCode block" +
	"\n```" +
	"\n:+1: :banana_dance:" +
	"\n***" +
	"\n> Quote\n" +
	"\nLink: [here](www.google.com)" +
	"\nImage: ![img](https://gdm-catalog-fmapi-prod.imgix.net/ProductLogo/4acbc64f-552d-4944-8474-b44a13a7bd3e.png?auto=format&q=50&fit=fill)" +
	"\nList:" +
	"\n- this" +
	"\n- is" +
	"\n- a" +
	"\n- list" +
	"\nNumbered list" +
	"\n1. this" +
	"\n2. is" +
	"\n3. a" +
	"\n4. list" +
	"\nItems" +
	"\n- [ ] Item one" +
	"\n- [ ] Item two" +
	"\n- [x] Completed item"

const markdownTable = "\n| Left-Aligned  | Center Aligned  | Right Aligned |" +
	"\n| :------------ |:---------------:| -----:|" +
	"\n| Left column 1 | this text       |  $100 |" +
	"\n| Left column 2 | is              |   $10 |" +
	"\n| Left column 3 | centered        |    $1 |"
