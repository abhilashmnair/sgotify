package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Response received from the Spotify API endpoint upon authorisation
// AccessToken  string	An access token that can be provided in subsequent calls, for example to Spotify Web API services.
// TokenType    string	How the access token may be used: always "Bearer"
// ExpiresIn        int	    The time period (in seconds) for which the access token is valid.
type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Function to get the access token. Accepts the clientID and clientSecret as parameters.
// Returns a Response object
func (r *Response) GetToken(clientID, clientSecret string) (*Response, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if r.AccessToken == "" {
		return nil, fmt.Errorf("could not get access token")
	}

	return r, nil
}
