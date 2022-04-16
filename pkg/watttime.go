package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IuryAlves/watttime-go/internal"
	"net/http"
)

const (
	BASE_ENDPOINT  = "https://api2.watttime.org"
	LOGIN_ENDPOINT = BASE_ENDPOINT + "/login"
	INDEX_ENDPOINT = BASE_ENDPOINT + "/index"
)

type WattTime struct {
	client internal.HTTPClient
}


func (w WattTime) Login(username, password string) (string, error) {
	var token Token
	req, _ := http.NewRequest("GET", LOGIN_ENDPOINT, nil)
	req.SetBasicAuth(username, password)
	resp, err := w.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 599 {
		return "", errors.New(fmt.Sprintf("login failed: Status Code %d", resp.StatusCode))
	}
	body, _ := internal.ReadBody(resp)

	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Println("Cannot unmarshall JSON:", err.Error())
		return "", err
	}
	return token.Value, nil
}

func (w WattTime) Index(token string, ba string) (RealTimeEmissionsIndex, error) {
	req, _ := http.NewRequest("GET", INDEX_ENDPOINT, nil)
	q := req.URL.Query()
	q.Add("ba", ba)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := w.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return RealTimeEmissionsIndex{}, err
	}

	body, err := internal.ReadBody(resp)
	var rtei RealTimeEmissionsIndex
	if err := json.Unmarshal(body, &rtei); err != nil {
		fmt.Println("cannot unmarshall JSON: ", err.Error())
		return RealTimeEmissionsIndex{}, err
	}
	return rtei, nil
}
