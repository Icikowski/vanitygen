package providers

import (
	"fmt"

	"git.sr.ht/~icikowski/vanitygen/config"
)

type githubProvider struct {
	baseProvider
}

// GetGoSourceTag implements Provider.
func (p *githubProvider) GetGoSourceTag() string {
	return fmt.Sprintf(
		"%s/%s %s %s/tree/%s{/dir} %s/blob/%s{/dir}/{file}#L{line}",
		p.domain, p.name,
		p.repo,
		p.repo, p.branch,
		p.repo, p.branch,
	)
}

var _ Provider = &githubProvider{}

func newGithubProvider(domain string, pkg config.Package) Provider {
	return &githubProvider{
		baseProvider: baseProvider{
			domain: domain,
			name:   pkg.Name,
			repo:   pkg.RepositoryURL,
			branch: pkg.Branch,
		},
	}
}
