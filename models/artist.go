package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SimplifiedArtist struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Hyperlink    string      `json:"href"`
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

type Artist struct {
	ExternalURLs ExternalURL   `json:"external_urls"`
	Hyperlink    string        `json:"href"`
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
	Followers    Followers     `json:"followers"`
	Genres       []string      `json:"genres"`
	Image        []ImageObject `json:"images"`
	Popularity   int           `json:"popularity"`
}

func (artist *Artist) GetArtist(id, token string) (*Artist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/artists/%s", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

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

	if err := json.Unmarshal(body, &artist); err != nil {
		return nil, err
	}

	return artist, nil

}
