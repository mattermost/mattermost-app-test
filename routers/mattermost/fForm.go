package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
	"github.com/mattermost/mattermost-server/v5/model"
)

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
