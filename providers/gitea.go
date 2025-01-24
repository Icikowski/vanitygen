package providers

import (
	"fmt"

	"git.sr.ht/~icikowski/vanitygen/config"
)

type giteaProvider struct {
	baseProvider
}

// GetGoSourceTag implements Provider.
func (p *giteaProvider) GetGoSourceTag() string {
	return fmt.Sprintf(
		"%s/%s %s %s/src/branch/%s{/dir} %s/src/branch/%s{/dir}/{file}#L{line}",
		p.domain, p.name,
		p.repo,
		p.repo, p.branch,
		p.repo, p.branch,
	)
}

var _ Provider = &giteaProvider{}

func newGiteaProvider(domain string, pkg config.Package) Provider {
	return &giteaProvider{
		baseProvider: baseProvider{
			domain: domain,
			name:   pkg.Name,
			repo:   pkg.RepositoryURL,
			branch: pkg.Branch,
		},
	}
}
