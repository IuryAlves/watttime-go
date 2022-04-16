package internal

import (
	"io/ioutil"
	"net/http"
)

var (
	Client HTTPClient
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func ReadBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
