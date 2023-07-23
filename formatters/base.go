package formatters

import (
	"bytes"
	"html/template"
	"io/fs"

	"github.com/Icikowski/vanitygen/config"
	"github.com/Icikowski/vanitygen/providers"
	"github.com/pkg/errors"
)

type baseFormatter struct {
	global globals
	tmpl   *template.Template
}

// EncodePackage implements Formatter.
func (f *baseFormatter) EncodePackage(pkg config.Package) ([]byte, error) {
	provider, err := providers.New(f.global.Domain, pkg)
	if err != nil {
		return nil, errors.Wrap(err, "providers.New")
	}

	params := pkgAttrs{
		Global:  f.global,
		Package: pkg,
		Meta:    provider,
	}

	buf := bytes.NewBuffer([]byte{})
	if err := f.tmpl.ExecuteTemplate(buf, templatePackage, params); err != nil {
		return nil, errors.Wrap(err, "f.tmpl.ExecuteTemplate")
	}

	return buf.Bytes(), nil
}

// EncodeRoot implements Formatter.
func (f *baseFormatter) EncodeRoot(pkgs config.Packages) ([]byte, error) {
	params := homeAttrs{
		Global:   f.global,
		Packages: pkgs,
	}

	buf := bytes.NewBuffer([]byte{})
	if err := f.tmpl.ExecuteTemplate(buf, templateHome, params); err != nil {
		return nil, errors.Wrap(err, "f.tmpl.ExecuteTemplate")
	}

	return buf.Bytes(), nil
}

var _ Formatter = &baseFormatter{}

func newFormatter(fs fs.FS, filename string) formatterGenerator {
	return func(global globals) (Formatter, error) {
		tmpl, err := template.ParseFS(fs, filename)
		if err != nil {
			return nil, errors.Wrap(err, "template.ParseFS")
		}

		return &baseFormatter{
			global: global,
			tmpl:   tmpl,
		}, nil
	}
}
