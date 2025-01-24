package providers

import (
	"fmt"

	"git.sr.ht/~icikowski/vanitygen/config"
)

type sourcehutProvider struct {
	baseProvider
}

// GetGoSourceTag implements Provider.
func (p *sourcehutProvider) GetGoSourceTag() string {
	return fmt.Sprintf(
		"%s/%s %s %s/tree/%s/item{/dir} %s/tree/%s/item{/dir}/{file}#L{line}",
		p.domain, p.name,
		p.repo,
		p.repo, p.branch,
		p.repo, p.branch,
	)
}

var _ Provider = &gitlabProvider{}

func newSourcehutProvider(domain string, pkg config.Package) Provider {
	return &sourcehutProvider{
		baseProvider: baseProvider{
			domain: domain,
			name:   pkg.Name,
			repo:   pkg.RepositoryURL,
			branch: pkg.Branch,
		},
	}
}
