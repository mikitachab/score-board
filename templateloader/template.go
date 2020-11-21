package templateloader

import (
	"html/template"
	"io"
)

// TemplateInterface define interface for template
type TemplateInterface interface {
	Render(w io.Writer, data interface{}) error
}

// Template represent compiled template
type Template struct {
	Name     string
	template *template.Template
}

// Render render temaplate
func (t *Template) Render(wr io.Writer, data interface{}) error {
	err := t.template.ExecuteTemplate(wr, "base", ViewData{data})
	return err
}
