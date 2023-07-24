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
func (f *baseFormatter) EncodePackage(pkg config.Package) ([]byte, map[string][]byte, error) {
	provider, err := providers.New(f.global.Domain, pkg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "providers.New")
	}

	params := pkgAttrs{
		Global:  f.global,
		Package: pkg,
		Meta:    provider,
	}

	buf := bytes.NewBuffer([]byte{})
	if err := f.tmpl.ExecuteTemplate(buf, templatePackage, params); err != nil {
		return nil, nil, errors.Wrap(err, "f.tmpl.ExecuteTemplate")
	}

	var subs map[string][]byte = nil
	if pkg.Subpackages != nil && len(pkg.Subpackages) != 0 {
		subs = map[string][]byte{}
		for _, sub := range pkg.Subpackages {
			subPkg := config.Package{
				Name:          pkg.Name + "/" + sub,
				Provider:      pkg.Provider,
				RepositoryURL: pkg.RepositoryURL,
				Branch:        pkg.Branch,
				Website:       pkg.Website,
			}

			subProvider, err := providers.New(f.global.Domain, subPkg)
			if err != nil {
				return nil, nil, errors.Wrap(err, "providers.New")
			}

			subParams := pkgAttrs{
				Global:  f.global,
				Package: subPkg,
				Meta:    subProvider,
			}

			subBuf := bytes.NewBuffer([]byte{})
			if err := f.tmpl.ExecuteTemplate(subBuf, templatePackage, subParams); err != nil {
				return nil, nil, errors.Wrap(err, "f.tmpl.ExecuteTemplate")
			}
			subs[sub] = subBuf.Bytes()
		}
	}

	return buf.Bytes(), subs, nil
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
