package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SimplifiedAlbum struct {
	AlbumType            string             `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	Hyperlink            string             `json:"href"`
	Id                   string             `json:"id"`
	Image                []ImageObject      `json:"images"`
	Name                 string             `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restrictions       `json:"restrictions"`
	Type                 string             `json:"type"`
	URI                  string             `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
}

type Album struct {
	AlbumType            string             `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []string           `json:"available_markets"`
	ExternalURLs         ExternalURL        `json:"external_urls"`
	Hyperlink            string             `json:"href"`
	Id                   string             `json:"id"`
	Image                []ImageObject      `json:"images"`
	Name                 string             `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restrictions       `json:"restrictions"`
	Type                 string             `json:"type"`
	URI                  string             `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
	Tracks               struct {
		Hyperlink   string            `json:"href"`
		Limit       int               `json:"limit"`
		Next        string            `json:"next"`
		Offset      int               `json:"offset"`
		Previous    string            `json:"previous"`
		TotalTracks int               `json:"total"`
		Items       []SimplifiedTrack `json:"items"`
	} `json:"tracks"`
	Copyrights  []Copyright `json:"copyrights"`
	ExternalIDs ExternalID  `json:"external_ids"`
	Genres      []string    `json:"genres"`
	Label       string      `json:"label"`
	Popularity  int         `json:"popularity"`
}

func (album *Album) GetAlbum(id, token string) (*Album, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/albums/%s", id), nil)
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

	if err := json.Unmarshal(body, &album); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return album, nil
}
