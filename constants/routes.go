package constants

import "github.com/mattermost/mattermost-plugin-apps/apps"

const (
	ManifestPath = "/manifest.json"
	InstallPath  = "/install"
	BindingsPath = "/bindings"

	AppSecret = "1234"

	Submit      = "/submit/ok"
	SubmitEmpty = "/submit/okEmpty"

	Error                         = "/error/error"
	ErrorEmpty                    = "/errors/empty"
	MarkdownFormError             = "/errors/markdownform"
	MarkdownFormErrorMissingField = "/errors/markdownformMissingField"

	Form                         = "/forms/ok"
	FormFull                     = "/forms/full"
	FormFullDisabled             = "/forms/full_disabled"
	FormDynamic                  = "/forms/dynamic"
	FormInvalid                  = "/forms/invalid"
	FormMultiselect              = "/forms/multiselect"
	FormWithButtons              = "/forms/buttons"
	FormMarkdown                 = "/forms/markdown"
	FormMarkdownWithMissingError = "/forms/markdownWitMissingError"
	FormRedefine                 = "/forms/redefine_form"
	FormEmbedded                 = "/forms/embedded"

	Lookup          = "/lookups/ok"
	LookupMultiword = "/lookups/multiword"
	LookupEmpty     = "/lookups/empty"
	LookupInvalid   = "/lookups/invalid"

	NavigateInternal = "/nav/internal"
	NavigateExternal = "/nav/external"
	NavigateInvalid  = "/nav/invalid"

	NotFoundPath     = "/foo"
	HTMLPath    = "/html"
	UnknownPath = "/unknown"

	StaticPath = "/static"

	CommandTrigger = "test"

	OtherPathOpenDialog              = "/other/open"
	OtherOpenDialogNoResponse        = "/NoResponse"
	OtherOpenDialogEmptyResponse     = "/EmptyResponse"
	OtherOpenDialogEphemeralResponse = "/EphemeralResponse"
	OtherOpenDialogUpdateResponse    = "/UpdateResponse"
	OtherOpenDialogBadResponse       = "/BadResponse"

	SubscribeCommand = "/subscribe"

	// Global

	NotifyUserCreated      = "/notify/" + string(apps.SubjectUserCreated)
	NotifyBotMention       = "/notify/" + string(apps.SubjectBotMentioned)
	NotifyBotJoinedChannel = "/notify/" + string(apps.SubjectBotJoinedChannel)
	NotifyBotLeftChannel   = "/notify/" + string(apps.SubjectBotLeftChannel)
	NotifyBotJoinedTeam    = "/notify/" + string(apps.SubjectBotJoinedTeam)
	NotifyBotLeftTeam      = "/notify/" + string(apps.SubjectBotLeftTeam)

	// Channel

	NotifyUserJoinedChannel = "/notify/" + string(apps.SubjectUserJoinedChannel)
	NotifyUserLeftChannel   = "/notify/" + string(apps.SubjectUserLeftChannel)
	NotifyPostCreated       = "/notify/" + string(apps.SubjectPostCreated)

	// Team

	NotifyUserJoinedTeam = "/notify/" + string(apps.SubjectUserJoinedTeam)
	NotifyUserLeftTeam   = "/notify/" + string(apps.SubjectUserLeftTeam)
	NotifyChannelCreated = "/notify/" + string(apps.SubjectChannelCreated)
)
