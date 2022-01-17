package path

import "github.com/mattermost/mattermost-plugin-apps/apps"

const (
	Manifest = "/manifest.json"
	Static   = "/static"
	Bindings = "/bindings"
	Install  = "/install"

	OK               = "/ok"
	OKEmpty          = "/empty"
	NavigateInternal = "/nav/internal"
	NavigateExternal = "/nav/external"

	ErrorDefault                  = "/errors/default"
	ErrorEmpty                    = "/errors/empty"
	ErrorMarkdownForm             = "/errors/markdownform"
	ErrorMarkdownFormMissingField = "/errors/markdownformMissingField"
	Error404                      = "/errors/foo"
	Error500                      = "/errors/internal"

	InvalidHTML        = "/invalid/html"
	InvalidUnknownType = "/invalid/unknown-type"
	InvalidLookup      = "/invalid/lookup"
	InvalidForm        = "/invalid/form"
	InvalidNavigate    = "/invalid/nav"

	FormSimple                    = "/forms/simple"
	FormSimpleSource              = "/forms/simpleSource"
	FormMarkdownError             = "/forms/markdownError"
	FormMarkdownErrorMissingField = "/forms/markdownMissingError"
	FormRefresh                   = "/forms/refresh"
	FormFull                      = "/forms/full"
	FormLookup                    = "/forms/lookup"
	FormFullSource                = "/forms/fullSource"
	FormFullReadonly              = "/forms/fullDisabled"
	FormMultiselect               = "/forms/multiselect"
	FormButtons                   = "/forms/buttons"
	FormRedefine                  = "/forms/redefine"

	Lookup          = "/lookups/ok"
	LookupMultiword = "/lookups/multiword"
	LookupEmpty     = "/lookups/empty"

	CreateEmbedded = "/create-embedded"

	CreateDialogMessage     = "/open-dialog"
	DialogNoResponse        = "/NoResponse"
	DialogEmptyResponse     = "/EmptyResponse"
	DialogEphemeralResponse = "/EphemeralResponse"
	DialogUpdateResponse    = "/UpdateResponse"
	DialogBadResponse       = "/BadResponse"

	NotifyBotJoinedChannel  = "/notify/" + string(apps.SubjectBotJoinedChannel)
	NotifyBotJoinedTeam     = "/notify/" + string(apps.SubjectBotJoinedTeam)
	NotifyBotLeftChannel    = "/notify/" + string(apps.SubjectBotLeftChannel)
	NotifyBotLeftTeam       = "/notify/" + string(apps.SubjectBotLeftTeam)
	NotifyBotMention        = "/notify/" + string(apps.SubjectBotMentioned)
	NotifyChannelCreated    = "/notify/" + string(apps.SubjectChannelCreated)
	NotifyPostCreated       = "/notify/" + string(apps.SubjectPostCreated)
	NotifyUserCreated       = "/notify/" + string(apps.SubjectUserCreated)
	NotifyUserJoinedChannel = "/notify/" + string(apps.SubjectUserJoinedChannel)
	NotifyUserJoinedTeam    = "/notify/" + string(apps.SubjectUserJoinedTeam)
	NotifyUserLeftChannel   = "/notify/" + string(apps.SubjectUserLeftChannel)
	NotifyUserLeftTeam      = "/notify/" + string(apps.SubjectUserLeftTeam)
)
