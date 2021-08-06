package constants

const (
	ManifestPath = "/manifest"
	InstallPath  = "/install"
	BindingsPath = "/bindings"

	AppSecret = "1234"

	BindingPathOK      = "/oks/ok"
	BindingPathOKEmpty = "/oks/empty"

	BindingPathError                         = "/errors/error"
	BindingPathErrorEmpty                    = "/errors/empty"
	BindingPathMarkdownFormError             = "/errors/markdownform"
	BindingPathMarkdownFormErrorMissingField = "/errors/markdownformMissingField"

	BindingPathFormOK                       = "/forms/ok"
	BindingPathFullFormOK                   = "/forms/full_ok"
	BindingPathFullDisabledOK               = "/forms/full_disabled_ok"
	BindingPathDynamicFormOK                = "/forms/dynamic_form_ok"
	BindingPathFormInvalid                  = "/forms/invalid"
	BindingPathMultiselectForm              = "/forms/multiselect"
	BindingPathWithButtonsOK                = "/forms/buttons"
	BindingPathMarkdownForm                 = "/forms/markdown"
	BindingPathMarkdownFormWithMissingError = "/forms/markdownWitMissingError"
	BindingPathRedefineFormOK               = "/forms/redefine_form_ok"
	BindingPathEmbeddedFormOK               = "/forms/embedded_form_ok"

	BindingPathLookupOK        = "/lookups/ok"
	BindingPathLookupMultiword = "/lookups/multiword"
	BindingPathLookupEmpty     = "/lookups/empty"
	BindingPathLookupInvalid   = "/lookups/invalid"

	BindingPathNavigateInternal = "/nav/internal"
	BindingPathNavigateExternal = "/nav/external"
	BindingPathNavigateInvalid  = "/nav/invalid"

	BindingPath404     = "/foo"
	BindingPathHTML    = "/html"
	BindingPathUnknown = "/unknown"

	StaticAssetPath = "/static"

	CommandTrigger = "test"

	OtherPathOpenDialog              = "/other/open"
	OtherOpenDialogNoResponse        = "/NoResponse"
	OtherOpenDialogEmptyResponse     = "/EmptyResponse"
	OtherOpenDialogEphemeralResponse = "/EphemeralResponse"
	OtherOpenDialogUpdateResponse    = "/UpdateResponse"
	OtherOpenDialogBadResponse       = "/BadResponse"

	SubscribeBotMention = "/subscribe/bot_mention"
	NotifyBotMention    = "/notify/bot_mention"
)
