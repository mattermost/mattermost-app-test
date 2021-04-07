package constants

const (
	ManifestPath = "/manifest"
	InstallPath  = "/install"
	BindingsPath = "/bindings"

	AppSecret = "1234"

	BindingPathOK      = "/oks/ok"
	BindingPathOKEmpty = "/oks/empty"

	BindingPathError      = "/errors/error"
	BindingPathErrorEmpty = "/errors/empty"

	BindingPathFormOK      = "/forms/ok"
	BindingPathFullFormOK  = "/forms/full_ok"
	BindingPathFormInvalid = "/forms/invalid"

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
)
