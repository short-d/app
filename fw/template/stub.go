package template

var _ Template = (*Stub)(nil)

type Stub struct {
	content string
}

func (s Stub) Render(renderTemplate string, includeTemplates []string, data interface{}) (string, error) {
	return s.content, nil
}

func NewTemplateFake(content string) Stub {
	return Stub{
		content: content,
	}
}
