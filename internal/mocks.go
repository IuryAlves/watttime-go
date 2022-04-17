package internal

import "net/http"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
	Request *http.Request
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.Request = req
	return GetDoFunc(req)
}
