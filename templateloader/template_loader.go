package templateloader

import (
	"html/template"
	"io"
	"path/filepath"
)

// TemplateLoader is abstraction for dealing with
// compiling and rendering view templates
type TemplateLoader struct {
	templatesPath string
}

// NewTemplateLoader create default TemplateLoader
func NewTemplateLoader() *TemplateLoader {
	return &TemplateLoader{
		templatesPath: "template",
	}
}

// RenderTemplateFunc is a callback for rendering template in views
type RenderTemplateFunc func(wr io.Writer, data interface{}) error

// ViewData is a struct which represent data
// that should be passed to view template
type ViewData struct {
	ViewData interface{}
}

// GetRenderTemplateFunc compile view templates and return callback to render it
func (tl *TemplateLoader) GetRenderTemplateFunc(templateName string) (RenderTemplateFunc, error) {
	templates, err := tl.compileTemplates(templateName)
	if err != nil {
		return nil, err
	}
	return func(wr io.Writer, data interface{}) error {
		err := templates.ExecuteTemplate(wr, "base", ViewData{data})
		return err
	}, nil
}

func (tl *TemplateLoader) compileTemplates(templateName string) (*template.Template, error) {
	paths := []string{
		tl.getTemplatePath("base.html"),
		tl.getTemplatePath("head.html"),
		tl.getTemplatePath("navbar.html"),
		tl.getTemplatePath(templateName),
	}

	templates, err := template.ParseFiles(paths...)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (tl *TemplateLoader) getTemplatePath(templateName string) string {
	return filepath.Join(tl.templatesPath, templateName)
}
