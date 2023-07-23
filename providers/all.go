package providers

import (
	"fmt"

	"github.com/Icikowski/vanitygen/config"
	"github.com/Icikowski/vanitygen/constants"
)

var all map[string]providerGenerator = map[string]providerGenerator{
	constants.ProviderGitea:  newGiteaProvider,
	constants.ProviderGitHub: newGithubProvider,
	constants.ProviderGitLab: newGitlabProvider,
}

func New(domain string, pkg config.Package) (Provider, error) {
	if generator, ok := all[pkg.Provider]; ok {
		return generator(domain, pkg), nil
	}
	return nil, fmt.Errorf("unknown provider '%s'", pkg.Provider)
}
