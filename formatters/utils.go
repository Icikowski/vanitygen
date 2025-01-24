package formatters

import (
	"time"

	"git.sr.ht/~icikowski/vanitygen/config"
	"git.sr.ht/~icikowski/vanitygen/constants"
	"git.sr.ht/~icikowski/vanitygen/providers"
	"pkg.icikowski.pl/generics"
)

const (
	templateHome    string = "home"
	templatePackage string = "pkg"
)

type globals struct {
	Domain   string
	SiteName string
	Author   string
	DocsURL  string

	Now time.Time
}

func globalsFromConfig(cfg config.Config) globals {
	return globals{
		Domain:   cfg.Domain,
		SiteName: cfg.SiteName,
		Author:   cfg.Author,
		DocsURL:  generics.Fallback(cfg.DocsURL, constants.DocsURL),
		Now:      time.Now(),
	}
}

type pkgAttrs struct {
	Global  globals
	Package config.Package
	Meta    providers.Provider
}

type homeAttrs struct {
	Global   globals
	Packages config.Packages
}
