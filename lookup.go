package main

import (
	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/path"
)

func initHTTPLookup(r *mux.Router) {
	handleCall(r, path.Lookup, handleLookup(simpleLookup))
	handleCall(r, path.LookupMultiword, handleLookup(multiwordLookup))
	handleCall(r, path.LookupEmpty, handleLookup(emptyLookup))
	handleCall(r, path.LookupInvalid, handleLookup(invalidLookup))
}

var simpleLookup = []apps.SelectOption{
	{
		Label: "dynamic value 1 label",
		Value: "sv1",
	},
	{
		Label: "dynamic value 2 label",
		Value: "sv2",
	},
}

var multiwordLookup = []apps.SelectOption{
	{
		Label: "dynamic value 1 label",
		Value: "dynamic value 1",
	},
	{
		Label: "dynamic value 2 label",
		Value: "dynamic value 2",
	},
}

var emptyLookup = []apps.SelectOption{}

var invalidLookup = []apps.SelectOption{
	{
		Label: "Valid",
		Value: "Valid",
	},
	{
		Label: "invalid",
	},
	{
		Value: "invalid",
	},
}
