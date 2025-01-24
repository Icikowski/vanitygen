package providers

import (
	"fmt"

	"git.sr.ht/~icikowski/vanitygen/config"
	"git.sr.ht/~icikowski/vanitygen/constants"
)

var all map[string]providerGenerator = map[string]providerGenerator{
	constants.ProviderGitea:     newGiteaProvider,
	constants.ProviderGitHub:    newGithubProvider,
	constants.ProviderGitLab:    newGitlabProvider,
	constants.ProviderSourcehut: newSourcehutProvider,
}

func New(domain string, pkg config.Package) (Provider, error) {
	if generator, ok := all[pkg.Provider]; ok {
		return generator(domain, pkg), nil
	}
	return nil, fmt.Errorf("unknown provider '%s'", pkg.Provider)
}
