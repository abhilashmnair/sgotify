package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PlaylistTrack struct {
	AddedAt string `json:"added_at"`
	AddedBy struct {
		ExternalURLs ExternalURL `json:"external_urls"`
		Followers    Followers   `json:"followers"`
		Hyperlink    string      `json:"href"`
		Id           string      `json:"id"`
		Type         string      `json:"type"`
		URI          string      `json:"uri"`
	} `json:"added_by"`
	IsLocal bool  `json:"is_local"`
	Track   Track `json:"track"`
}

type Playlist struct {
	Collaborative bool          `json:"collaborative"`
	Description   string        `json:"description"`
	ExternalURLs  ExternalURL   `json:"external_urls"`
	Followers     Followers     `json:"followers"`
	Hyperlink     string        `json:"href"`
	Id            string        `json:"id"`
	Image         []ImageObject `json:"images"`
	Name          string        `json:"name"`
	Owner         struct {
		ExternalURLs ExternalURL `json:"external_urls"`
		Followers    Followers   `json:"followers"`
		Hyperlink    string      `json:"href"`
		Id           string      `json:"id"`
		Type         string      `json:"type"`
		URI          string      `json:"uri"`
		Name         string      `json:"display_name"`
	} `json:"owner"`
	Public     bool   `json:"public"`
	SnapshotId string `json:"snapshot_id"`
	Tracks     struct {
		Hyperlink   string          `json:"href"`
		Limit       int             `json:"limit"`
		Next        string          `json:"next"`
		Offset      int             `json:"offset"`
		Previous    string          `json:"previous"`
		TotalTracks int             `json:"total"`
		Items       []PlaylistTrack `json:"items"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

func (playlist *Playlist) GetPlaylist(id, token string) (*Playlist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", id), nil)
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

	if err := json.Unmarshal(body, &playlist); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return playlist, nil
}
