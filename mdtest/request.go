package mdtest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdrequest"
)

type TransportHandleFunc func(req *http.Request) (*http.Response, error)

type TransportFake struct {
	handle TransportHandleFunc
}

func (r TransportFake) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.handle(req)
}

func NewGraphQLRequestFake(handleFunc TransportHandleFunc) mdrequest.GraphQL {
	client := http.Client{
		Transport: TransportFake{
			handle: handleFunc,
		}}
	return mdrequest.NewGraphQL(mdrequest.NewHTTP(client))
}

func JSONResponse(jsonObj map[string]interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(jsonObj)
	if err != nil {
		return nil, err
	}

	body := ioutil.NopCloser(strings.NewReader(string(jsonStr)))
	return &http.Response{
		StatusCode:    http.StatusOK,
		Body:          body,
		ContentLength: int64(len(jsonStr)),
	}, nil
}

var _ fw.HTTPRequest = (*HTTPRequestFake)(nil)

type HTTPRequestFake struct {
	err error
}

func (h HTTPRequestFake) JSON(method string, url string, headers map[string]string, body string, v interface{}) error {
	return h.err
}

func NewHTTPRequestFake(err error) HTTPRequestFake {
	return HTTPRequestFake{
		err: err,
	}
}
