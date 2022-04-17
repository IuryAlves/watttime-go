package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func QueryFromOptions(options any) url.Values {
	q := make(url.Values)
	optionValues := reflect.ValueOf(options)
	optionType := optionValues.Type()
	var fieldName string
	for i := 0; i < optionValues.NumField(); i++ {
		fieldName = strings.ToLower(optionType.Field(i).Name)
		q.Add(fieldName, fmt.Sprintf("%v", optionValues.Field(i).Interface()))
	}
	return q
}
