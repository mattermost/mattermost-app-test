package path

import "github.com/mattermost/mattermost-plugin-apps/apps"

const (
	Manifest = "/manifest.json"
	Install  = "/install"
	Bindings = "/bindings"

	ProfileView          = "/profile/view"
	ProfileCommand       = "/profile/command"
	ProfileChannelHeader = "/profile/channel_header"

	OK      = "/ok"
	OKEmpty = "/empty"

	ErrorDefault                  = "/errors/default"
	ErrorEmpty                    = "/errors/empty"
	ErrorMarkdownForm             = "/errors/markdownform"
	ErrorMarkdownFormMissingField = "/errors/markdownformMissingField"

	FormSimple                   = "/forms/simple"
	FormFull                     = "/forms/full"
	FormFullDisabled             = "/forms/full_disabled"
	FormDynamic                  = "/forms/dynamic"
	FormInvalid                  = "/forms/invalid"
	FormMultiselect              = "/forms/multiselect"
	FormWithButtons              = "/forms/buttons"
	FormMarkdown                 = "/forms/markdown"
	FormMarkdownWithMissingError = "/forms/markdownWitMissingError"
	FormRedefine                 = "/forms/redefine"
	FormEmbedded                 = "/forms/embedded"

	Lookup          = "/lookups/ok"
	LookupMultiword = "/lookups/multiword"
	LookupEmpty     = "/lookups/empty"
	LookupInvalid   = "/lookups/invalid"

	NavigateInternal = "/nav/internal"
	NavigateExternal = "/nav/external"
	NavigateInvalid  = "/nav/invalid"

	NotFound    = "/foo"
	HTML        = "/html"
	UnknownType = "/unknown"

	Static = "/static"

	OtherOpenDialog                  = "/other/open"
	OtherOpenDialogNoResponse        = "/NoResponse"
	OtherOpenDialogEmptyResponse     = "/EmptyResponse"
	OtherOpenDialogEphemeralResponse = "/EphemeralResponse"
	OtherOpenDialogUpdateResponse    = "/UpdateResponse"
	OtherOpenDialogBadResponse       = "/BadResponse"

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
