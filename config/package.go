package config

import (
	"errors"
	"fmt"

	ghErrors "github.com/pkg/errors"

	"github.com/Icikowski/vanitygen/constants"
)

type Package struct {
	Name          string   `json:"name" yaml:"name"`
	Provider      string   `json:"provider" yaml:"provider"`
	RepositoryURL string   `json:"repoUrl" yaml:"repoUrl"`
	Branch        string   `json:"branch" yaml:"branch"`
	Website       *string  `json:"website,omitempty" yaml:"website,omitempty"`
	Subpackages   []string `json:"subpackages,omitempty" yaml:"subpackages,omitempty"`
}

func (p Package) Validate() error {
	var errs error

	if p.Name == "" {
		errs = errors.Join(errs, errors.New("name is not set"))
	}

	if !constants.AllProviders.Contains(p.Provider) {
		errs = errors.Join(errs, fmt.Errorf("provider '%s' is not valid", p.Provider))
	}

	if p.RepositoryURL == "" {
		errs = errors.Join(errs, errors.New("repository URL is not set"))
	}

	if p.Branch == "" {
		errs = errors.Join(errs, errors.New("branch is not set"))
	}

	if p.Website != nil && *p.Website == "" {
		errs = errors.Join(errs, errors.New("website is present but not set"))
	}

	return errs
}

type Packages []Package

func (p Packages) Validate() error {
	var errs error

	for i := range p {
		errs = errors.Join(errs, ghErrors.Wrapf(p[i].Validate(), "package [%d] is invalid", i))
	}

	return errs
}
