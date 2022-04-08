package pkg

import (
	"IuryAlves/watttime-go.git/internal"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BASE_ENDPOINT  = "https://api2.watttime.org"
	LOGIN_ENDPOINT = BASE_ENDPOINT + "/login"
	INDEX_ENDPOINT = BASE_ENDPOINT + "/index"
)

var client = &http.Client{}

func Login(username, password string) (string, error) {
	req, _ := http.NewRequest("GET", LOGIN_ENDPOINT, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body, err := internal.ReadBody(resp)
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Println("Cannot unmarshall JSON")
		return "", err
	}
	return token.Value, nil
}

func Index(token string, ba string) (RealTimeEmissionsIndex, error) {
	req, _ := http.NewRequest("GET", INDEX_ENDPOINT, nil)
	q := req.URL.Query()
	q.Add("ba", ba)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return RealTimeEmissionsIndex{}, err
	}

	body, err := internal.ReadBody(resp)
	var rtei RealTimeEmissionsIndex
	if err := json.Unmarshal(body, &rtei); err != nil {
		fmt.Println("Cannot unmarshall JSON")
		return RealTimeEmissionsIndex{}, err
	}
	return rtei, nil
}
