package formatters

import (
	"embed"
	"fmt"

	"github.com/Icikowski/vanitygen/config"
)

var (
	//go:embed plain.tmpl
	srcPlain embed.FS
	//go:embed bootstrap.tmpl
	srcBootstrap embed.FS
)

var all map[string]formatterGenerator = map[string]formatterGenerator{
	"plain":     newFormatter(srcPlain, "plain.tmpl"),
	"bootstrap": newFormatter(srcBootstrap, "bootstrap.tmpl"),
}

func New(name string, cfg config.Config) (Formatter, error) {
	global := globalsFromConfig(cfg)

	if formatter, ok := all[name]; ok {
		return formatter(global)
	}
	return nil, fmt.Errorf("unknown formatter '%s", name)
}
