package constants

import "pkg.icikowski.pl/sets"

// Formatters' names
const (
	FormatterPlain     string = "plain"
	FormatterBootstrap string = "bootstrap"
)

// AllFormatters is a set of all formatters' names
var AllFormatters *sets.Set[string] = sets.New(
	FormatterPlain,
	FormatterBootstrap,
)
