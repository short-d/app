package mdtest

import "github.com/short-d/app/fw"

var _ fw.Template = (*TemplateFake)(nil)

type TemplateFake struct {
	content string
}

func (t TemplateFake) Render(renderTemplate string, includeTemplates []string, data interface{}) (string, error) {
	return t.content, nil
}

func NewTemplateFake(content string) TemplateFake {
	return TemplateFake{
		content: content,
	}
}
