package graphql

import (
	"bytes"
	"html/template"
)

// GraphiQLConfig contains all the options which can be used to customize
// GraphiQL editor.
// You can find more configurations here:
// https://github.com/graphql/graphiql/blob/main/packages/graphiql/README.md#props
type GraphiQLConfig struct {
	APIEndpoint        string
	DefaultQuery       string
	VariableEditorOpen bool
	DocEditorOpen      bool
}

type WebUI interface {
	RenderHTML() string
}

var _ WebUI = (*GraphiQL)(nil)

type GraphiQL struct {
	html string
}

func (g GraphiQL) RenderHTML() string {
	return g.html
}

func NewGraphiQL(endpoint string, defaultQuery string) GraphiQL {
	config := GraphiQLConfig{
		APIEndpoint:        endpoint,
		DefaultQuery:       defaultQuery,
		VariableEditorOpen: true,
		DocEditorOpen:      true,
	}

	tmpl := template.Must(template.New("graphiQL").Parse(graphiQLTemplate))
	var buf bytes.Buffer

	err := tmpl.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	return GraphiQL{
		html: buf.String(),
	}
}

const graphiQLTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        height: 100%;
        margin: 0;
        width: 100%;
        overflow: hidden;
      }

      #graphiql {
        height: 100vh;
      }
    </style>
    <script
      crossorigin
      src="https://unpkg.com/react@16/umd/react.production.min.js"
    ></script>
    <script
      crossorigin
      src="https://unpkg.com/react-dom@16/umd/react-dom.production.min.js"
    ></script>
    <link rel="stylesheet" href="https://unpkg.com/graphiql/graphiql.min.css" />
  </head>

  <body>
    <div id="graphiql">Loading...</div>
    <script
      src="https://unpkg.com/graphiql/graphiql.min.js"
      type="application/javascript"
    ></script>
    <script>
      function graphQLFetcher(graphQLParams) {
        return fetch(
          '{{.APIEndpoint}}',
          {
            method: 'post',
            headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(graphQLParams),
            credentials: 'omit',
          },
        ).then(function (response) {
          return response.json().catch(function () {
            return response.text();
          });
        });
      }

      ReactDOM.render(
        React.createElement(GraphiQL, {
          fetcher: graphQLFetcher,
          defaultQuery: {{.DefaultQuery}},
          defaultVariableEditorOpen: {{.VariableEditorOpen}},
          docExplorerOpen: {{.DocEditorOpen}}
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </body>
</html>
`
