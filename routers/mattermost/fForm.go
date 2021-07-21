package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

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

func fFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
				Expand: &apps.Expand{
					Team: apps.ExpandSummary,
				},
			},
			Fields: []*apps.Field{
				{
					Name:       "text",
					Type:       apps.FieldTypeText,
					Label:      "text",
					ModalLabel: "text",
				},
				{
					Name:  "navigate",
					Type:  apps.FieldTypeBool,
					Label: "navigate",
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFullFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Full Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathLookupOK,
			},
			Fields: []*apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
				{
					Name:  "text",
					Type:  apps.FieldTypeText,
					Label: "text",
				},
				{
					Type:        "markdown",
					Name:        "markdown",
					Description: "***\n## User information\nRemember to fill all these fields with the **user** information, not the general information.",
				},
				// {
				// 	Name: "mk1",
				// 	Type: "markdown",
				// 	Description: "## Markdown title" +
				// 		"\nHello world" +
				// 		"\nText styles: _italics_ **bold** **_bold-italic_** ~~strikethrough~~ `code`" +
				// 		"\nUsers and channels: @sysadmin ~town-square" +
				// 		"\n```" +
				// 		"\nCode block" +
				// 		"\n```" +
				// 		"\n:+1: :banana_dance:" +
				// 		"\n***" +
				// 		"\n> Quote\n" +
				// 		"\nLink: [here](www.google.com)" +
				// 		"\nImage: ![img](https://gdm-catalog-fmapi-prod.imgix.net/ProductLogo/4acbc64f-552d-4944-8474-b44a13a7bd3e.png?auto=format&q=50&fit=fill)" +
				// 		"\nList:" +
				// 		"\n- this" +
				// 		"\n- is" +
				// 		"\n- a" +
				// 		"\n- list" +
				// 		"\nNumbered list" +
				// 		"\n1. this" +
				// 		"\n2. is" +
				// 		"\n3. a" +
				// 		"\n4. list" +
				// 		"\nItems" +
				// 		"\n- [ ] Item one" +
				// 		"\n- [ ] Item two" +
				// 		"\n- [x] Completed item",
				// },
				// {
				// 	Name: "mk2",
				// 	Type: "markdown",
				// 	Description: "\n| Left-Aligned  | Center Aligned  | Right Aligned |" +
				// 		"\n| :------------ |:---------------:| -----:|" +
				// 		"\n| Left column 1 | this text       |  $100 |" +
				// 		"\n| Left column 2 | is              |   $10 |" +
				// 		"\n| Left column 3 | centered        |    $1 |",
				// },
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
				},
				{
					Name:  "channel",
					Type:  apps.FieldTypeChannel,
					Label: "channel",
				},
				{
					Name:  "user",
					Type:  apps.FieldTypeUser,
					Label: "user",
				},
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
				},
				{
					Name:          "multi",
					Type:          apps.FieldTypeStaticSelect,
					Label:         "multi",
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
						{
							Label: "static value 3",
							Value: "sv3",
						},
						{
							Label: "1",
							Value: "1",
						},
						{
							Label: "2",
							Value: "2",
						},
						{
							Label: "3",
							Value: "3",
						},
						{
							Label: "4",
							Value: "4",
						},
						{
							Label: "5",
							Value: "5",
						},
						{
							Label: "6",
							Value: "6",
						},
						{
							Label: "7",
							Value: "7",
						},
						{
							Label: "8",
							Value: "8",
						},
						{
							Label: "9",
							Value: "9",
						},
						{
							Label: "10",
							Value: "10",
						},
					},
				},
				{
					Name:     "user_readonly",
					Type:     apps.FieldTypeUser,
					Label:    "user_readonly",
					ReadOnly: true,
				},
				{
					Name:     "static_readonly",
					Type:     apps.FieldTypeStaticSelect,
					Label:    "static_readonly",
					ReadOnly: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFullFormDisabledOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Full Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathLookupOK,
			},
			Fields: []*apps.Field{
				{
					Name:     "lookup",
					Type:     apps.FieldTypeDynamicSelect,
					Label:    "lookup",
					ReadOnly: true,
				},
				{
					Name:     "text",
					Type:     apps.FieldTypeText,
					Label:    "text",
					ReadOnly: true,
					Value:    "Hello world",
				},
				{
					Type:  "markdown",
					Name:  "markdown",
					Value: "Hello ~~world~~",
				},
				{
					Name:     "boolean",
					Type:     apps.FieldTypeBool,
					Label:    "boolean",
					ReadOnly: true,
				},
				{
					Name:     "channel",
					Type:     apps.FieldTypeChannel,
					Label:    "channel",
					ReadOnly: true,
				},
				{
					Name:     "user",
					Type:     apps.FieldTypeUser,
					Label:    "user",
					ReadOnly: true,
				},
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
					},
					ReadOnly: true,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fDynamicFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numFields := len(c.Values)
	fields := []*apps.Field{}

	for i := 0; i < numFields+5; i++ {
		fields = append(fields, &apps.Field{
			Name:          fmt.Sprintf("static%v", i),
			Type:          apps.FieldTypeStaticSelect,
			Label:         fmt.Sprintf("static%v", i),
			SelectRefresh: true,
			SelectStaticOptions: []apps.SelectOption{
				{
					Label: "static value 1",
					Value: "sv1",
				},
				{
					Label: "static value 2",
					Value: "sv2",
				},
			},
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Dynamic Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathDynamicFormOK,
			},
			Fields: fields,
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormInvalid(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
	}
	utils.WriteCallResponse(w, resp)
}

var iterationsPerChannelID = map[string]int{}

const maxIterations = 5

func fFormRedefine(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	cid := c.Context.ChannelID
	iters := iterationsPerChannelID[cid]
	iters = (iters + 1) % maxIterations
	iterationsPerChannelID[cid] = iters

	fields := []*apps.Field{}

	for i := 0; i < iters; i++ {
		name := fmt.Sprintf("text%v", i)

		fields = append(fields, &apps.Field{
			Name:       name,
			Type:       apps.FieldTypeText,
			Label:      name,
			ModalLabel: name,
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: fields,
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormMultiselect(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test Multiselect Form",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathOK,
			},
			Fields: []*apps.Field{
				{
					Name:          "static",
					Type:          apps.FieldTypeStaticSelect,
					Label:         "static",
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "static value 1",
							Value: "sv1",
						},
						{
							Label: "static value 2",
							Value: "sv2",
						},
						{
							Label: "static value 3",
							Value: "sv3",
						},
						{
							Label: "static value 4",
							Value: "sv4",
						},
					},
				},
				{
					Name:          "user",
					Type:          apps.FieldTypeUser,
					Label:         "user",
					SelectIsMulti: true,
				},
				{
					Name:          "channel",
					Type:          apps.FieldTypeChannel,
					Label:         "channel",
					SelectIsMulti: true,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormEmbedded(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	client := mmclient.AsBot(c.Context)
	p := &model.Post{
		ChannelId: c.Context.ChannelID,
	}

	p.AddProp(apps.PropAppBindings, []*apps.Binding{
		{
			Location: "embedded",
			Form: &apps.Form{
				Title:  "Test",
				Header: "Test header",
				Call: &apps.Call{
					Path: constants.BindingPathOK,
				},
				Fields: []*apps.Field{},
			},
			AppID:       c.Context.AppID,
			Description: "Please fill out this form so we can get it fixed  :hammer_and_wrench:",
			Bindings: []*apps.Binding{
				{
					Location: "problem",
					Call: &apps.Call{
						Path: constants.BindingPathOK,
					},
					Bindings: []*apps.Binding{
						{
							Location: "hardware",
							Label:    "Hardware Failure",
						},
						{
							Location: "software",
							Label:    "Software Error",
						},
						{
							Location: "wrong",
							Label:    "Wrong Product",
						},
					},
				},
				{
					Location: "provider",
					Call: &apps.Call{
						Path: constants.BindingPathOK,
					},
					Bindings: []*apps.Binding{
						{
							Location: "work",
							Label:    "Cell Phone",
						},
					},
				},
				{
					Location: "button",
					Label:    "Submit",
					Call: &apps.Call{
						Path: constants.BindingPathOK,
					},
				},
			},
		},
	})

	client.CreatePost(p)

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithButtonsOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numButtonsFloat, _ := c.State.(float64)
	numButtons := int(numButtonsFloat)

	if v, ok := c.Values["submit"].(string); ok {
		switch v {
		case "add_buttons":
			numButtons++
		case "error":
			utils.WriteCallResponse(w, *apps.NewErrorCallResponse(errors.New("You caused an error :)")))
		}
	}

	buttons := []apps.SelectOption{
		{
			Label: "add buttons",
			Value: "add_buttons",
		},
		{
			Label: "error",
			Value: "error",
		},
	}

	for i := 0; i < numButtons; i++ {
		buttons = append(buttons, apps.SelectOption{
			Label: fmt.Sprintf("button%v", i),
			Value: fmt.Sprintf("button%v", i),
		})
	}

	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:         "Test multiple buttons Form",
			Header:        "Test header",
			SubmitButtons: "submit",
			Call: &apps.Call{
				Path:  constants.BindingPathWithButtonsOK,
				State: numButtons,
			},
			Fields: []*apps.Field{
				{
					Name:                "submit",
					Type:                apps.FieldTypeStaticSelect,
					Label:               "static",
					SelectIsMulti:       true,
					SelectStaticOptions: buttons,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithMarkdownError(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test markdown descriptions and errors",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathMarkdownFormError,
			},
			Fields: []*apps.Field{
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					Description: `| Option | Message  | Image |
| :------------ |:---------------:| -----:|
| Opt1 | You are good     |  :smile: |
| Opt2 | You are awesome              | :+1: |
| Opt3| You are great       |    :smirk:  |`,
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "button1",
							Value: "button1",
						},
						{
							Label: "button2",
							Value: "button2",
						},
						{
							Label: "button3",
							Value: "button3",
						},
						{
							Label: "button4",
							Value: "button4",
						},
					},
				},
				{
					Name:        "text",
					Type:        apps.FieldTypeText,
					Label:       "text",
					Description: fullMarkdown, //"Go [here](www.google.com) for more information.",
				},
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
					Description: `Mark this field only if:
					1. You want
					2. You need
					3. You should`,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}

func fFormWithMarkdownErrorMissingField(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	resp := apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Test markdown descriptions and errors",
			Header: "Test header",
			Call: &apps.Call{
				Path: constants.BindingPathMarkdownFormErrorMissingField,
			},
			Fields: []*apps.Field{
				{
					Name:  "static",
					Type:  apps.FieldTypeStaticSelect,
					Label: "static",
					Description: `| Option | Message  | Image |
| :------------ |:---------------:| -----:|
| Opt1 | You are good     |  :smile: |
| Opt2 | You are awesome              | :+1: |
| Opt3| You are great       |    :smirk:  |`,
					SelectIsMulti: true,
					SelectStaticOptions: []apps.SelectOption{
						{
							Label: "button1",
							Value: "button1",
						},
						{
							Label: "button2",
							Value: "button2",
						},
						{
							Label: "button3",
							Value: "button3",
						},
						{
							Label: "button4",
							Value: "button4",
						},
					},
				},
				{
					Name:        "text",
					Type:        apps.FieldTypeText,
					Label:       "text",
					Description: fullMarkdown, //"Go [here](www.google.com) for more information.",
				},
				{
					Name:  "boolean",
					Type:  apps.FieldTypeBool,
					Label: "boolean",
					Description: `Mark this field only if:
					1. You want
					2. You need
					3. You should`,
				},
			},
		},
	}
	utils.WriteCallResponse(w, resp)
}
