package formatters

import (
	"github.com/Icikowski/vanitygen/config"
)

type Formatter interface {
	EncodeRoot(config.Packages) ([]byte, error)
	EncodePackage(config.Package) ([]byte, error)
}

type formatterGenerator func(globals) (Formatter, error)
