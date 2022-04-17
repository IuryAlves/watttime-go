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

func TestIndexBaSuccess(t *testing.T) {
	client := &internal.MockClient{}
	wattTime := &WattTime{Client: client}
	r, _ := json.Marshal(RealTimeEmissionsIndex{Ba: "SE", Freq: "10", Percent: "10", Moer: "10", PointTime: "10"})
	internal.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(r)),
			Request:	req,
		}, nil
	}
	options := &IndexOptions{Ba: "SE"}
	rtei, _ := wattTime.Index("token", *options)

	assert.EqualValues(t, rtei.Ba, "SE")
	assert.EqualValues(t, "SE", client.Request.URL.Query().Get("ba"))
}

func TestIndexLongitudeLatitudeSuccess(t *testing.T) {
	client := &internal.MockClient{}
	wattTime := &WattTime{Client: client}
	r, _ := json.Marshal(
		RealTimeEmissionsIndex{
			Ba: "CAISO_NORTH",
			Freq: "300",
			Percent: "53",
			Moer: "850.743982",
			PointTime: "2019-01-29T14:55:00.00Z"},
		)
	internal.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(r)),
			Request:	req,
		}, nil
	}
	options := &IndexOptions{Latitude: 42.372, Longitude: -72.519}
	rtei, _ := wattTime.Index("token", *options)

	assert.EqualValues(t, rtei.Ba, "CAISO_NORTH")
	assert.EqualValues(t, "42.372", client.Request.URL.Query().Get("latitude"))
	assert.EqualValues(t, "-72.519", client.Request.URL.Query().Get("longitude"))
}