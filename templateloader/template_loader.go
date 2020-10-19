package templateloader

import (
	"html/template"
	"io"
	"path/filepath"
)

type TemplateLoader struct {
	templatesPath string
}

func NewTemplateLoader() *TemplateLoader {
	return &TemplateLoader{
		templatesPath: "template",
	}
}

type RenderTemplateFunc func(wr io.Writer, data interface{}) error

type ViewData struct {
	ViewData interface{}
}

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
