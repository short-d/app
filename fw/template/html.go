package template

import (
	"bytes"
	"html/template"
	"path"
)

var _ Template = (*HTML)(nil)

type HTML struct {
	templateRootDir string
}

func (t HTML) Render(renderTemplate string, includeTemplates []string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(t.getFilePaths(includeTemplates)...)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, renderTemplate, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (t HTML) getFilePaths(filenames []string) []string {
	var filePaths []string
	for _, filename := range filenames {
		filePaths = append(filePaths, path.Join(t.templateRootDir, filename))
	}
	return filePaths
}

func NewHTML(templateRootDir string) HTML {
	return HTML{
		templateRootDir: templateRootDir,
	}
}
