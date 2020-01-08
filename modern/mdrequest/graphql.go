package mdrequest

import (
	"encoding/json"
	"net/http"

	"github.com/short-d/app/fw"
)

type graphQlResponse struct {
	Data interface{} `json:"data"`
}

var _ fw.GraphQlRequest = (*GraphQL)(nil)

type GraphQL struct {
	http fw.HTTPRequest
	root string
}

func (g GraphQL) Query(query fw.GraphQlQuery, headers map[string]string, response interface{}) error {
	var res graphQlResponse

	reqBuf, err := json.Marshal(query)
	if err != nil {
		return err
	}

	headers["Content-Type"] = "application/json"
	err = g.http.JSON(http.MethodPost, g.root, headers, string(reqBuf), &res)
	if err != nil {
		return err
	}

	resBuf, err := json.Marshal(res.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBuf, &response)
	return err
}

func (g GraphQL) RootUrl(root string) fw.GraphQlRequest {
	g.root = root
	return g
}

func NewGraphQL(http fw.HTTPRequest) GraphQL {
	return GraphQL{
		http: http,
	}
}
