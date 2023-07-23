package providers

import "fmt"

type baseProvider struct {
	domain string
	name   string
	repo   string
	branch string
}

// GetGoImportTag implements Provider.
func (p *baseProvider) GetGoImportTag() string {
	return fmt.Sprintf(
		"%s/%s git %s",
		p.domain, p.name,
		p.repo,
	)
}
