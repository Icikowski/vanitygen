package formatters

import (
	"git.sr.ht/~icikowski/vanitygen/config"
)

type Formatter interface {
	EncodeRoot(config.Packages) ([]byte, error)
	EncodePackage(config.Package) ([]byte, map[string][]byte, error)
}

type formatterGenerator func(globals) (Formatter, error)
