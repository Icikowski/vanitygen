package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/Icikowski/vanitygen/config"
	"github.com/Icikowski/vanitygen/formatters"
	"github.com/caarlos0/env/v9"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	indexFile string = "index.html"
)

var args config.Arguments

func init() {
	if err := env.Parse(&args); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&args.Output, "out", args.Output, "output directory")
	flag.StringVar(&args.Config, "config", args.Config, "configuration file")
	flag.StringVar(&args.Formatter, "format", args.Formatter, "output formatter [plain/bootstrap]")
	flag.Parse()
}

func getDataLoader(filename string) (func([]byte, any) error, error) {
	switch path.Ext(filename) {
	case ".json":
		return json.Unmarshal, nil
	case ".yaml", ".yml":
		return yaml.Unmarshal, nil
	}

	return nil, fmt.Errorf("unknown extension '%s'", path.Ext(filename))
}

func main() {
	data, err := os.ReadFile(args.Config)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "unable to read file '%s'", args.Config))
	}

	dataLoader, err := getDataLoader(args.Config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to determine data loader"))
	}

	var cfg config.Config
	if err := dataLoader(data, &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "unable to unmarshal config"))
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(errors.Wrap(err, "config is invalid"))
	}

	formatter, err := formatters.New(args.Formatter, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to instantiate formatter"))
	}

	if err := os.Mkdir(args.Output, 0666); err != nil && !os.IsExist(err) {
		log.Fatal(errors.Wrap(err, "unable to create output directory"))
	}

	sort.SliceStable(cfg.Packages, func(i, j int) bool {
		return strings.Compare(cfg.Packages[i].Name, cfg.Packages[j].Name) < 0
	})

	for _, pkg := range cfg.Packages {
		target := path.Join(args.Output, pkg.Name, indexFile)

		log.Printf("building index for package '%s' in '%s'", pkg.Name, target)

		if err := os.Mkdir(path.Dir(target), 0666); err != nil && !os.IsExist(err) {
			log.Fatal(errors.Wrap(err, "unable to create package output directory"))
		}

		out, err := formatter.EncodePackage(pkg)
		if err != nil {
			log.Fatal(errors.Wrap(err, "unable to encode package"))
		}

		f, err := os.Create(target)
		if err != nil {
			log.Panic(errors.Wrapf(err, "unable to create package index '%s'", target))
		}
		defer f.Close()

		if _, err := f.Write(out); err != nil {
			log.Panic(errors.Wrapf(err, "unable to write package index '%s'", target))
		}
	}

	target := path.Join(args.Output, indexFile)

	log.Printf("building index list in '%s'", target)

	out, err := formatter.EncodeRoot(cfg.Packages)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to encode index"))
	}

	f, err := os.Create(target)
	if err != nil {
		log.Panic(errors.Wrapf(err, "unable to create index '%s'", target))
	}
	defer f.Close()

	if _, err := f.Write(out); err != nil {
		log.Panic(errors.Wrapf(err, "unable to write index '%s'", target))
	}
}
