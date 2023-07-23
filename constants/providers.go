package constants

import "pkg.icikowski.pl/sets"

// Providers' names
const (
	ProviderGitea     string = "gitea"
	ProviderGitHub    string = "github"
	ProviderGitLab    string = "gitlab"
	ProviderSourcehut string = "sourcehut"
)

// AllProviders is a set of all providers' names
var AllProviders *sets.Set[string] = sets.New(
	ProviderGitea,
	ProviderGitHub,
	ProviderGitLab,
	ProviderSourcehut,
)
