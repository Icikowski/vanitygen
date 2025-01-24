package providers

import (
	"git.sr.ht/~icikowski/vanitygen/config"
)

type Provider interface {
	GetGoImportTag() string
	GetGoSourceTag() string
}

type providerGenerator func(domain string, pkg config.Package) Provider
