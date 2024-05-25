package sgotify

import (
	"fmt"

	"github.com/abhilashmnair/sgotify/auth"
	"github.com/abhilashmnair/sgotify/models"
)

func Authorize(ClientID, ClientSecret string) (*auth.Response, error) {
	auth := &auth.Response{}
	data, err := auth.GetToken(ClientID, ClientSecret)
	if err != nil {
		fmt.Printf("Error getting access token: %s\n", err)
		return nil, err
	}
	return data, nil
}

func FetchTrack(token, trackID string) (*models.Track, error) {
	track := &models.Track{}
	data, err := track.GetTrack(trackID, token)
	if err != nil {
		fmt.Println("Error fetching track data:", err)
		return nil, err
	}
	return data, nil
}

func FetchAlbum(token, albumID string) (*models.Album, error) {
	album := &models.Album{}
	data, err := album.GetAlbum(albumID, token)
	if err != nil {
		fmt.Println("Error fetching album data:", err)
		return nil, err
	}
	return data, nil
}

func FetchArtist(token, artistID string) (*models.Artist, error) {
	artist := &models.Artist{}
	data, err := artist.GetArtist(artistID, token)
	if err != nil {
		fmt.Println("Error fetching artist data:", err)
		return nil, err
	}
	return data, nil
}
