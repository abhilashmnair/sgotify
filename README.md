![Spotify_Icon_RGB_Black](https://raw.githubusercontent.com/abhilashmnair/sgotify/main/logo.png)
# sGOtify - Spotify API Wrapper for GO
[![Go Reference](https://pkg.go.dev/badge/godoc.org/github.com/abhilashmnair/sgotify.svg)](https://pkg.go.dev/github.com/abhilashmnair/sgotify)

sGOtify is a Go package that provides an easy-to-use wrapper for working with the Spotify API.
This repository follows the [Spotify Web API Documentation](https://developer.spotify.com/documentation/web-api)

`NOTE: This project is still under development`

## Installation
To install the library
```
go get github.com/abhilashmnair/sgotify
```

## Usage
The usage of the package is relatively simple. The following instructions go through the steps.

#### Authorize
To use the package, you need to authorize using the clientID and clientSecret obtained from [Spotify API](https://developer.spotify.com/)

```
ClientID = "your_client_id"
ClientSecret = "your_client_secret"

auth, err := sgotify.Authorize(ClientID, ClientSecret)
if err != nil {
	fmt.Println("Authorization failed", err)
	return
}

token := auth.AccessToken
```

Example: Fetch a track with given ID
```
trackID = 28x6WpVq1ty4dPDeqK0zPO
track, err := sgotify.FetchTrack(token, trackID)
if err != nil {
	fmt.Println("Fetching failed!", err)
	return
}

fmt.Printf("%s - %s\n", track.Artists[0].Name, track.Name) \\ The Weeknd - Blinding Lights
```

## LICENSE
This repository is licensed under [MIT License](https://github.com/abhilashmnair/sgotify/tree/main?tab=MIT-1-ov-file). See LICENSE for full licensing text.
