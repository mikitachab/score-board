package templateloader

import (
	"html/template"
	"path/filepath"
)

// Interface define interface of TemplateLoader
type Interface interface {
	LoadTemplate(string) (*Template, error)
}

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

// ViewData is a struct which represent data
// that should be passed to view template
type ViewData struct {
	ViewData interface{}
}

// LoadTemplate load and return compiled template
func (tl *TemplateLoader) LoadTemplate(templateName string) (*Template, error) {
	template, err := tl.compileTemplates(templateName)
	if err != nil {
		return nil, err
	}
	return &Template{templateName, template}, nil
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
