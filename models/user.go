package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name         string        `json:"display_name"`
	ExternalURLs ExternalURL   `json:"external_urls"`
	Followers    Followers     `json:"followers"`
	Hyperlink    string        `json:"href"`
	Id           string        `json:"id"`
	Image        []ImageObject `json:"images"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
}

func (user *User) GetUser(id, token string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/users/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return user, nil

}
