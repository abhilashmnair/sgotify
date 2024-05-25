package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	Duration         int                `json:"duration_ms"`
	IsExplicit       bool               `json:"explicit"`
	ExternalURLs     ExternalURL        `json:"external_urls"`
	Hyperlink        string             `json:"href"`
	Id               string             `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       LinkedFrom         `json:"linked_from"`
	Restrictions     Restrictions       `json:"restrictions"`
	Name             string             `json:"name"`
	PreviewURL       string             `json:"preview_url"`
	TrackNumber      int                `json:"track_number"`
	Type             string             `json:"type"`
	URI              string             `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type Track struct {
	SimplifiedAlbum  `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	Duration         int          `json:"duration_ms"`
	IsExplicit       bool         `json:"explicit"`
	ExternalIDs      ExternalID   `json:"external_ids"`
	ExternalURLs     ExternalURL  `json:"external_urls"`
	Hyperlink        string       `json:"href"`
	Id               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type PlaylistTrack struct {
}

func (track *Track) GetTrack(id, token string) (*Track, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", id), nil)
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

	if err := json.Unmarshal(body, &track); err != nil {
		return nil, err
	}

	return track, nil

}
