package command

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-test/constants"
)

func getError() apps.Binding {
	base := apps.Binding{
		Location: "error",
		Label:    "error",
		Bindings: []apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getWithError())
	base.Bindings = append(base.Bindings, getWithEmptyError())
	base.Bindings = append(base.Bindings, getWithInvalidNavigate())
	base.Bindings = append(base.Bindings, getWithInvalidForm())
	base.Bindings = append(base.Bindings, getWith404Error())
	base.Bindings = append(base.Bindings, getWithHTMLSite())
	base.Bindings = append(base.Bindings, getWithArbitraryJSON())
	base.Bindings = append(base.Bindings, getWithUnknownResponse())
	base.Bindings = append(base.Bindings, getWithFormInvalid())
	base.Bindings = append(base.Bindings, getWithFormError())
	base.Bindings = append(base.Bindings, getWithFormErrorEmpty())
	base.Bindings = append(base.Bindings, getWithFormNavigate())
	base.Bindings = append(base.Bindings, getWithFormOK())
	base.Bindings = append(base.Bindings, getWithFormHTMLSite())
	base.Bindings = append(base.Bindings, getWithFormArbitraryJSON())
	base.Bindings = append(base.Bindings, getWithFormUnknownResponse())
	base.Bindings = append(base.Bindings, getWithLookupError())
	base.Bindings = append(base.Bindings, getWithLookupForm())
	base.Bindings = append(base.Bindings, getWithLookupNavigate())
	base.Bindings = append(base.Bindings, getWithLookup404())
	base.Bindings = append(base.Bindings, getWithLookupHTML())
	base.Bindings = append(base.Bindings, getWithLookupUnknown())

	return base
}


func getWithInvalidNavigate() apps.Binding {
	return apps.Binding{
		Location: "with_invalid_navigate",
		Label:    "with_invalid_navigate",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.NavigateInvalid,
		},
	}
}

func getWithInvalidForm() apps.Binding {
	return apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.FormInvalid,
		},
	}
}


func getWithHTMLSite() apps.Binding {
	return apps.Binding{
		Location: "with_html_site",
		Label:    "with_html_site",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.HTMLPath,
		},
	}
}

func getWithArbitraryJSON() apps.Binding {
	return apps.Binding{
		Location: "with_arbitrary_json",
		Label:    "with_arbitrary_json",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse() apps.Binding {
	return apps.Binding{
		Location: "with_unknown_response",
		Label:    "with_unknown_response",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}

func getWithFormInvalid() apps.Binding {
	return apps.Binding{
		Location: "with_form_invalid",
		Label:    "with_form_invalid",
		Call: &apps.Call{
			Path: constants.FormInvalid,
		},
	}
}

func getWithFormError() apps.Binding {
	return apps.Binding{
		Location: "with_form_error",
		Label:    "with_form_error",
		Call: &apps.Call{
			Path: constants.Error,
		},
	}
}

func getWithFormErrorEmpty() apps.Binding {
	return apps.Binding{
		Location: "with_form_error_empty",
		Label:    "with_form_error_empty",
		Call: &apps.Call{
			Path: constants.ErrorEmpty,
		},
	}
}

func getWithFormNavigate() apps.Binding {
	return apps.Binding{
		Location: "with_form_navigate",
		Label:    "with_form_navigate",
		Call: &apps.Call{
			Path: constants.NavigateExternal,
		},
	}
}

func getWithFormOK() apps.Binding {
	return apps.Binding{
		Location: "with_form_ok",
		Label:    "with_form_ok",
		Call: &apps.Call{
			Path: constants.SubmitOK,
		},
	}
}

func getWithFormHTMLSite() apps.Binding {
	return apps.Binding{
		Location: "with_form_html_site",
		Label:    "with_form_html_site",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.HTMLPath,
		},
	}
}

func getWithFormArbitraryJSON() apps.Binding {
	return apps.Binding{
		Location: "with_form_arbitrary_json",
		Label:    "with_form_arbitrary_json",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithFormUnknownResponse() apps.Binding {
	return apps.Binding{
		Location: "with_form_unknown_response",
		Label:    "with_form_unknown_response",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}

func getWithLookupError() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_error",
		Label:    "with_lookup_error",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.Error,
		},
	}
}

func getWithLookupForm() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_form",
		Label:    "with_lookup_form",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.Form,
		},
	}
}

func getWithLookupNavigate() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_navigate",
		Label:    "with_lookup_navigate",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.NavigateExternal,
		},
	}
}

func getWithLookup404() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_404",
		Label:    "with_lookup_404",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.NotFoundPath,
		},
	}
}

func getWithLookupHTML() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_html",
		Label:    "with_lookup_html",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.HTMLPath,
		},
	}
}

func getWithLookupUnknown() apps.Binding {
	return apps.Binding{
		Location: "with_lookup_unknown",
		Label:    "with_lookup_unknown",
		Form: &apps.Form{
			Fields: []apps.Field{
				{
					Name:  "lookup",
					Type:  apps.FieldTypeDynamicSelect,
					Label: "lookup",
				},
			},
		},
		Call: &apps.Call{
			Path: constants.UnknownPath,
		},
	}
}
