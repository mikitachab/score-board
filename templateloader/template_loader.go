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

func (tl *TemplateLoader) GetRenderTemplateFunc(name string) (RenderTemplateFunc, error) {
	templates, err := tl.compileTemplates(name)
	if err != nil {
		return nil, err
	}
	return func(wr io.Writer, data interface{}) error {
		err := templates.ExecuteTemplate(wr, "base", data)
		return err
	}, nil
}

func (tl *TemplateLoader) compileTemplates(viewTemplate string) (*template.Template, error) {
	paths := []string{
		tl.getTemplatePath("base.html"),
		tl.getTemplatePath("head.html"),
		tl.getTemplatePath("navbar.html"),
		tl.getTemplatePath(viewTemplate),
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
