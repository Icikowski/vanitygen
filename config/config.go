package config

import (
	"errors"

	ghErrors "github.com/pkg/errors"
)

type Config struct {
	Domain   string   `json:"domain" yaml:"domain"`
	SiteName string   `json:"siteName" yaml:"siteName"`
	Author   string   `json:"author" yaml:"author"`
	DocsURL  *string  `json:"docs" yaml:"docs"`
	Packages Packages `json:"pkgs" yaml:"pkgs"`
}

func (c Config) Validate() error {
	var errs error

	if c.Domain == "" {
		errs = errors.Join(errs, errors.New("domain is not set"))
	}

	if c.SiteName == "" {
		errs = errors.Join(errs, errors.New("site name is not set"))
	}

	if c.Author == "" {
		errs = errors.Join(errs, errors.New("author is not set"))
	}

	if c.DocsURL != nil && *c.DocsURL == "" {
		errs = errors.Join(errs, errors.New("docs URL is present but not set"))
	}

	if err := c.Packages.Validate(); err != nil {
		errs = errors.Join(ghErrors.Wrap(err, "some packages are invalid"))
	}

	return errs
}
