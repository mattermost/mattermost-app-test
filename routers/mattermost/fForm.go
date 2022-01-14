package mattermost

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
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

var simpleSourceForm = apps.Form{
	Source: apps.NewCall(constants.Form),
}

var fullSourceForm = apps.Form{
	Source: apps.NewCall(constants.FormFull),
}

var invalidSourceForm = apps.Form{
	Source: apps.NewCall(constants.FormInvalid),
}

var simpleForm = apps.Form{
	Title:  "Test",
	Header: "Test header",
	Submit: apps.NewCall(constants.Submit).WithExpand(apps.Expand{
		Team: apps.ExpandSummary,
	}),
	Fields: []apps.Field{
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
}

var fullForm = apps.Form{
	Title:  "Test Full Form",
	Header: "Test header",
	Fields: []apps.Field{
		{
			Name: "lookup",
			Type: apps.FieldTypeDynamicSelect,

			SelectDynamicLookup: apps.NewCall(constants.Lookup),
		},
		{
			Name: "text",
			Type: apps.FieldTypeText,
		},
		{
			Type: "markdown",
			Name: "markdown",

			Description: "***\n## User information\nRemember to fill all these fields with the **user** information, not the general information.",
		},
		{
			Name: "boolean",
			Type: apps.FieldTypeBool,
		},
		{
			Name: "channel",
			Type: apps.FieldTypeChannel,
		},
		{
			Name: "user",
			Type: apps.FieldTypeUser,
		},
		{
			Name: "static",
			Type: apps.FieldTypeStaticSelect,

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
			Name: "multi",
			Type: apps.FieldTypeStaticSelect,

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
			Name: "user_readonly",
			Type: apps.FieldTypeUser,

			ReadOnly: true,
		},
		{
			Name: "static_readonly",
			Type: apps.FieldTypeStaticSelect,

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
}

func fFormOK(w http.ResponseWriter, _ *http.Request, _ *apps.CallRequest) {
	utils.WriteCallResponse(w, apps.NewFormResponse(simpleForm))
}

func fFullFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallResponse(w, apps.NewFormResponse(fullForm))
}

func fDynamicFormOK(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	numFields := len(c.Values)
	fields := []apps.Field{}

	for i := 0; i < numFields+5; i++ {
		fields = append(fields, apps.Field{
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
			Submit: apps.NewCall(constants.FormDynamic),
			Source: apps.NewCall(constants.FormDynamic),
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

	fields := []apps.Field{}

	for i := 0; i < iters; i++ {
		name := fmt.Sprintf("text%v", i)

		fields = append(fields, apps.Field{
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
			Submit: apps.NewCall(constants.Submit),
			Source: apps.NewCall(constants.FormDynamic),
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
			Submit: apps.NewCall(constants.Submit),
			Fields: []apps.Field{
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
	client := appclient.AsBot(c.Context)
	p := &model.Post{
		ChannelId: c.Context.ChannelID,
	}

	p.AddProp(apps.PropAppBindings, []apps.Binding{
		{
			Location: "embedded",
			Form: &apps.Form{
				Title:  "Test",
				Header: "Test header",
				Submit: apps.NewCall(constants.Submit),
				Fields: []apps.Field{},
			},
			AppID:       c.Context.AppID,
			Description: "Please fill out this form so we can get it fixed  :hammer_and_wrench:",
			Bindings: []apps.Binding{
				{
					Location: "problem",
					Submit:   apps.NewCall(constants.Submit),
					Bindings: []apps.Binding{
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
					Submit:   apps.NewCall(constants.Submit),
					Bindings: []apps.Binding{
						{
							Location: "work",
							Label:    "Cell Phone",
						},
					},
				},
				{
					Location: "button",
					Label:    "Submit",
					Submit:   apps.NewCall(constants.Submit),
				},
			},
		},
	})

	_, _ = client.CreatePost(p)

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
			utils.WriteCallResponse(w, apps.NewErrorResponse(errors.New("you caused an error :)")))
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
			Source: &apps.Call{
				Path:  constants.FormWithButtons,
				State: numButtons,
			},
			Submit: apps.NewCall(constants.Submit),
			Fields: []apps.Field{
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
			Submit: &apps.Call{
				Path: constants.MarkdownFormError,
			},
			Fields: []apps.Field{
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
					Description: fullMarkdown, // "Go [here](www.google.com) for more information.",
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
			Submit: &apps.Call{
				Path: constants.MarkdownFormErrorMissingField,
			},
			Fields: []apps.Field{
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
					Description: fullMarkdown, // "Go [here](www.google.com) for more information.",
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
