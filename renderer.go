// Package echorenderer provides an implementation of [echo.Renderer] using [html/template].
package echorenderer

import (
	"errors"
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
)

type (
	// Creates a [template.FuncMap] using an [echo.Context].
	FuncMapper func(c echo.Context) template.FuncMap

	// Provides configuration for [Renderer].
	Options struct {
		// An [fs.FS] that contains the templates.
		FS fs.FS
		// The glob patterns used in [template.Parse].
		Include []string
		// An optional [FuncMapper].
		Funcs FuncMapper
	}

	// Provides [echo.Renderer] implementation through [template.Template].
	Renderer struct {
        *template.Template

		fsys  fs.FS
		funcs FuncMapper
	}
)

var (
    ErrNoInclude = errors.New("echorenderer: must have at least one include pattern")
    ErrNoFS = errors.New("echorenderer: must have filesystem")
)

// Render implements [echo.Renderer].
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	tmpl, err := r.Clone()
	if err != nil {
		return err
	}

	if r.funcs != nil {
		_ = tmpl.Funcs(r.funcs(c))
	}

	_, err = tmpl.ParseFS(r.fsys, name)

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

// Creates a new [Renderer].
func New(opts *Options) (*Renderer, error) {
    if opts.FS == nil {
        return nil, ErrNoFS
    }

	if opts.Include == nil || len(opts.Include) == 0 {
		return nil, ErrNoInclude
	}

	tmpl := template.New("")
	if opts.Funcs != nil {
		_ = tmpl.Funcs(opts.Funcs(nil))
	}

	for _, pattern := range opts.Include {
		_, err := tmpl.ParseFS(opts.FS, pattern)
		if err != nil {
			return nil, err
		}
	}

	return &Renderer{
		Template: tmpl,
		fsys:     opts.FS,
		funcs:    opts.Funcs,
	}, nil

}

var _ echo.Renderer = (*Renderer)(nil)
