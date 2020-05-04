package graphql

import (
	"encoding/json"
	"net/http"

	"github.com/short-d/app/fw/webreq"
)

type graphQLResponse struct {
	Data interface{} `json:"data"`
}

type Client struct {
	httpClient webreq.HTTP
	root       string
}

func (c Client) Query(query Query, headers map[string]string, response interface{}) error {
	var res graphQLResponse

	reqBuf, err := json.Marshal(query)
	if err != nil {
		return err
	}

	headers["Content-Type"] = "application/json"
	err = c.httpClient.JSON(http.MethodPost, c.root, headers, string(reqBuf), &res)
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

func (c Client) SetRootUrl(root string) {
	c.root = root
}

func NewClient(httpClient webreq.HTTP) Client {
	return Client{
		httpClient: httpClient,
	}
}
