package main

import (
	"github.com/mattermost/mattermost-app-test/path"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

var lookupForm = apps.Form{
	Title:  "Test Lookup Form",
	Submit: callOK,
	Fields: []apps.Field{
		{
			Name:                "simple",
			Type:                apps.FieldTypeDynamicSelect,
			SelectDynamicLookup: apps.NewCall(path.Lookup),
		},
		{
			Name:                "multiword",
			Type:                apps.FieldTypeDynamicSelect,
			SelectDynamicLookup: apps.NewCall(path.LookupMultiword),
		},
		{
			Name:                "empty",
			Type:                apps.FieldTypeDynamicSelect,
			SelectDynamicLookup: apps.NewCall(path.LookupEmpty),
		},
		{
			Name:                "invalid",
			Type:                apps.FieldTypeDynamicSelect,
			SelectDynamicLookup: apps.NewCall(path.LookupInvalid),
		},
	},
}
