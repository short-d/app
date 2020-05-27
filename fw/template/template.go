package template

type Template interface {
	Render(renderTemplate string, includeTemplates []string, data interface{}) (string, error)
}
