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

	fsys    fs.FS
	pattern string
	mapper  FuncMapper
}

// Provides configuration for [Renderer].
type Options struct {
	// An [fs.FS] that contains the templates.
	FS fs.FS
	// The glob pattern used in [template.Parse].
	Pattern string
	// An optional [FuncMapper].
	FuncMapper FuncMapper
}

// Render implements [echo.Renderer].
func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	tmpl, err := r.Template.Clone()
	if err != nil {
		return err
	}

	_, err = tmpl.
		Funcs(r.mapper(c)).
		ParseFS(r.fsys, name)

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

// Creates a new [Renderer].
func New(opts *Options) (*Renderer, error) {
	tmpl, err := template.New("").
		Funcs(opts.FuncMapper(nil)).
		ParseFS(opts.FS, opts.Pattern)

	if err != nil {
		return nil, err
	}

	return &Renderer{
		Template: tmpl,
		fsys:     opts.FS,
		pattern:  opts.Pattern,
		mapper:   opts.FuncMapper,
	}, nil

}

var _ echo.Renderer = (*Renderer)(nil)
