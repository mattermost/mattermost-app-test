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
	BindingPathFormInvalid = "/forms/invalid"

	BindingPathLookupOK    = "/lookups/ok"
	BindingPathLookupEmpty = "/lookups/empty"

	BindingPathNavigateInternal = "/nav/internal"
	BindingPathNavigateExternal = "/nav/external"
	BindingPathNavigateInvalid  = "/nav/invalid"

	BindingPath404     = "/foo"
	BindingPathHTML    = "/html"
	BindingPathUnknown = "/unkown"
)
