package providers

import (
	"github.com/Icikowski/vanitygen/config"
)

type Provider interface {
	GetGoImportTag() string
	GetGoSourceTag() string
}

type providerGenerator func(domain string, pkg config.Package) Provider
