package path

const (
	Manifest = "/manifest.json"
	Static   = "/static"
	Bindings = "/bindings"
	Install  = "/install"
	Notify   = "/notify"

	// Commands
	CreateEmbedded      = "/create-embedded"
	CreateDialogMessage = "/create-dialog"
	Subscribe           = "/subscribe"
	Unsubscribe         = "/unsubscribe"

	// Submit responses
	OK               = "/ok"
	OKEmpty          = "/empty"
	NavigateInternal = "/nav/internal"
	NavigateExternal = "/nav/external"

	// Form responses
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

	// Lookup responses
	Lookup          = "/lookups/ok"
	LookupMultiword = "/lookups/multiword"
	LookupEmpty     = "/lookups/empty"

	// Error responses
	ErrorDefault                  = "/errors/default"
	ErrorEmpty                    = "/errors/empty"
	ErrorMarkdownForm             = "/errors/markdownform"
	ErrorMarkdownFormMissingField = "/errors/markdownformMissingField"
	Error404                      = "/errors/foo"
	Error500                      = "/errors/internal"

	// Invalid responses
	InvalidHTML        = "/invalid/html"
	InvalidUnknownType = "/invalid/unknown-type"
	InvalidLookup      = "/invalid/lookup"
	InvalidForm        = "/invalid/form"
	InvalidNavigate    = "/invalid/nav"

	// Responders for slack attachment buttons, open interactive dialogs.
	DialogNoResponse        = "/NoResponse"
	DialogEmptyResponse     = "/EmptyResponse"
	DialogEphemeralResponse = "/EphemeralResponse"
	DialogUpdateResponse    = "/UpdateResponse"
	DialogBadResponse       = "/BadResponse"
)
