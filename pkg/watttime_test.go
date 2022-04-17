package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"io/ioutil"
	"net/http"

	"github.com/IuryAlves/watttime-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	wattTime := &WattTime{Client: &internal.MockClient{}}
	r, _ := json.Marshal(Token{Value: "token"})
	internal.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(r)),
		}, nil
	}
	token, _ := wattTime.Login("test", "test")

	assert.EqualValues(t, "token", token)
}

func TestLoginFailed(t *testing.T) {
	wattTime := &WattTime{Client: &internal.MockClient{}}
	internal.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 401,
			Body:       nil,
		}, nil
	}
	_, err := wattTime.Login("test", "test")
	assert.EqualError(t, err, errors.New("login failed: Status Code 401").Error())
}
