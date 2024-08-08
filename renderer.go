// Package echorenderer provides an implementation of [echo.Renderer] using [html/template].
package echorenderer

import (
	"html/template"
	"io"
	"io/fs"

	"github.com/labstack/echo/v4"
)

// Creates a [template.FuncMap] using an [echo.Context].
type FuncMapper func(c echo.Context) template.FuncMap

// Provides [echo.Renderer] implementation through [template.Template].
type Renderer struct {
	*template.Template

	fsys  fs.FS
	funcs FuncMapper
}

// Provides configuration for [Renderer].
type Options struct {
	// An [fs.FS] that contains the templates.
	FS fs.FS
	// The glob patterns used in [template.Parse].
	Include []string
	// An optional [FuncMapper].
	Funcs FuncMapper
}

// Render implements [echo.Renderer].
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	tmpl, err := r.Template.Clone()
	if err != nil {
		return err
	}

	_ = tmpl.Funcs(r.funcs(c))
	_, err = tmpl.ParseFS(r.fsys, name)

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

// Creates a new [Renderer].
func New(opts *Options) (*Renderer, error) {
	tmpl := template.New("").Funcs(opts.Funcs(nil))
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
