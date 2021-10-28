package command

import (
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func getError(siteURL, appID string) apps.Binding {
	base := apps.Binding{
		Location: "error",
		Label:    "error",
		Bindings: []apps.Binding{},
	}

	base.Bindings = append(base.Bindings, getWithError(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithEmptyError(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithInvalidNavigate(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithInvalidForm(siteURL, appID))
	base.Bindings = append(base.Bindings, getWith404Error(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithHTMLSite(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithArbitraryJSON(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithUnknownResponse(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormInvalid(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormError(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormErrorEmpty(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormNavigate(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormOK(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormHTMLSite(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormArbitraryJSON(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithFormUnknownResponse(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookupError(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookupForm(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookupNavigate(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookup404(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookupHTML(siteURL, appID))
	base.Bindings = append(base.Bindings, getWithLookupUnknown(siteURL, appID))

	return base
}

func getWithError(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_error",
		Label:    "with_error",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithEmptyError(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_empty_error",
		Label:    "with_empty_error",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathErrorEmpty,
		},
	}
}

func getWithInvalidNavigate(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_invalid_navigate",
		Label:    "with_invalid_navigate",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathNavigateInvalid,
		},
	}
}

func getWithInvalidForm(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_invalid_form",
		Label:    "with_invalid_form",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathFormInvalid,
		},
	}
}

func getWith404Error(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_404_error",
		Label:    "with_404_error",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPath404,
		},
	}
}

func getWithHTMLSite(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_html_site",
		Label:    "with_html_site",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithArbitraryJSON(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_arbitrary_json",
		Label:    "with_arbitrary_json",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithUnknownResponse(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_unknown_response",
		Label:    "with_unknown_response",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithFormInvalid(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_invalid",
		Label:    "with_form_invalid",
		Call: &apps.Call{
			Path: constants.BindingPathFormInvalid,
		},
	}
}

func getWithFormError(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_error",
		Label:    "with_form_error",
		Call: &apps.Call{
			Path: constants.BindingPathError,
		},
	}
}

func getWithFormErrorEmpty(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_error_empty",
		Label:    "with_form_error_empty",
		Call: &apps.Call{
			Path: constants.BindingPathErrorEmpty,
		},
	}
}

func getWithFormNavigate(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_navigate",
		Label:    "with_form_navigate",
		Call: &apps.Call{
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithFormOK(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_ok",
		Label:    "with_form_ok",
		Call: &apps.Call{
			Path: constants.BindingPathOK,
		},
	}
}

func getWithFormHTMLSite(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_html_site",
		Label:    "with_form_html_site",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithFormArbitraryJSON(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_arbitrary_json",
		Label:    "with_form_arbitrary_json",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.ManifestPath,
		},
	}
}

func getWithFormUnknownResponse(_, _ string) apps.Binding {
	return apps.Binding{
		Location: "with_form_unknown_response",
		Label:    "with_form_unknown_response",
		Form:     &apps.Form{},
		Call: &apps.Call{
			Path: constants.BindingPathUnknown,
		},
	}
}

func getWithLookupError(_, _ string) apps.Binding {
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
			Path: constants.BindingPathError,
		},
	}
}

func getWithLookupForm(_, _ string) apps.Binding {
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
			Path: constants.BindingPathFormOK,
		},
	}
}

func getWithLookupNavigate(_, _ string) apps.Binding {
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
			Path: constants.BindingPathNavigateExternal,
		},
	}
}

func getWithLookup404(_, _ string) apps.Binding {
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
			Path: constants.BindingPath404,
		},
	}
}

func getWithLookupHTML(_, _ string) apps.Binding {
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
			Path: constants.BindingPathHTML,
		},
	}
}

func getWithLookupUnknown(_, _ string) apps.Binding {
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
			Path: constants.BindingPathUnknown,
		},
	}
}
