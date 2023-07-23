package providers

import (
	"fmt"

	"github.com/Icikowski/vanitygen/config"
)

type gitlabProvider struct {
	baseProvider
}

// GetGoSourceTag implements Provider.
func (p *gitlabProvider) GetGoSourceTag() string {
	return fmt.Sprintf(
		"%s/%s %s %s/-/tree/%s{/dir} %s/-/blob/%s{/dir}/{file}#L{line}",
		p.domain, p.name,
		p.repo,
		p.repo, p.branch,
		p.repo, p.branch,
	)
}

var _ Provider = &gitlabProvider{}

func newGitlabProvider(domain string, pkg config.Package) Provider {
	return &gitlabProvider{
		baseProvider: baseProvider{
			domain: domain,
			name:   pkg.Name,
			repo:   pkg.RepositoryURL,
			branch: pkg.Branch,
		},
	}
}
