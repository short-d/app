package graphql

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/webreq"
)

type graphQLResponse struct {
	Data   interface{}   `json:"data"`
	Errors []interface{} `json:"errors"`
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

	if len(res.Errors) > 0 {
		return errors.New(fmt.Sprintf("%v", res.Errors[0]))
	}

	resBuf, err := json.Marshal(res.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBuf, &response)
	return err
}

type ClientFactory struct {
	httpClient webreq.HTTP
}

func (c ClientFactory) NewClient(root string) Client {
	return Client{
		httpClient: c.httpClient,
		root:       root,
	}
}

func NewClientFactory(httpClient webreq.HTTP) ClientFactory {
	return ClientFactory{httpClient: httpClient}
}
