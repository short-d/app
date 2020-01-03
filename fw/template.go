package fw

type Template interface {
	Render(renderTemplate string, includeTemplates []string, data interface{}) (string, error)
}
