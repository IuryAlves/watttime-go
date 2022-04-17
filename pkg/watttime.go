package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"encoding/json"
	"net/http"

	"github.com/IuryAlves/watttime-go/internal"
)

const (
	BaseEndpoint     = "https://api2.watttime.org"
	LoginEndpoint    = BaseEndpoint + "/login"
	IndexEndpoint    = BaseEndpoint + "/index"
	RegisterEndpoint = BaseEndpoint + "/register"
)

type WattTime struct {
	Client internal.HTTPClient
}

// New Instantiates a WattTime Client
func New() *WattTime {
	return &WattTime{Client: &http.Client{}}
}

// Login Authenticates towards the WattTime API
//
// Login returns a token of type string. Use the token for further requests towards the API
func (w WattTime) Login(username, password string) (string, error) {
	var token Token
	req, _ := http.NewRequest("GET", LoginEndpoint, nil)
	req.SetBasicAuth(username, password)
	resp, err := w.Client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 599 {
		return "", fmt.Errorf("login failed: Status Code %d", resp.StatusCode)
	}
	body, _ := internal.ReadBody(resp)

	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Println("Cannot unmarshall JSON:", err.Error())
		return "", err
	}
	return token.Value, nil
}

// Register creates a new user for the WattTime API
//
// Returns (true, nil) if the registration was successful and (false, err) otherwise
func (w WattTime) Register(username, password, email, org string) (bool, error){
	type registerPayload struct {
		Username string
		Password string
		Email    string
		Org      string
	}

	payload, err := json.Marshal(&registerPayload{
		Username: username,
		Password: password,
		Email: email,
		Org: org,
	})

	if err != nil {
		return false, err
	}
	req, _ := http.NewRequest(
		"POST",
		RegisterEndpoint,
		ioutil.NopCloser(bytes.NewReader(payload),
		))
	resp, err := w.Client.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 599 {
		return false, fmt.Errorf("register failed: Status Code %d", resp.StatusCode)
	}
	return true, nil
}

// Index Provides a real-time signal indicating the marginal carbon intensity for
// the local grid for the current time (updated every 5 minutes).
func (w WattTime) Index(token string, options IndexOptions) (RealTimeEmissionsIndex, error) {
	err := validateIndexOptions(options)
	if err != nil {
		return RealTimeEmissionsIndex{}, err
	}
	req, _ := http.NewRequest("GET", IndexEndpoint, nil)
	q := internal.QueryFromOptions(options)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := w.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return RealTimeEmissionsIndex{}, err
	}

	body, _ := internal.ReadBody(resp)
	var rtei RealTimeEmissionsIndex
	if err := json.Unmarshal(body, &rtei); err != nil {
		fmt.Println("cannot unmarshall JSON: ", err.Error())
		return RealTimeEmissionsIndex{}, err
	}
	return rtei, nil
}

// validateIndexOptions checks that 'latitude' and 'longitude' are not provided
// in case 'ba' is provided
func validateIndexOptions(options IndexOptions) error {
	if len(options.Ba) > 0 && (options.Latitude != 0 || options.Longitude != 0) {
		return fmt.Errorf("provide ba OR provide latitude+longitude, not all three")
	}
	return nil
}
