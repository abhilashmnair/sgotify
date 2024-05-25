# sGOtify
![Version](https://img.shields.io/badge/version-v1.0-blue)

sGOtify is a Go package that provides an easy-to-use wrapper for working with the Spotify API.
This repository follows the [Spotify Web API Documentation](https://developer.spotify.com/documentation/web-api)

This project is still under `development`

## Installation
To install the library
```
go get github.com/abhilashmnair/sgotify
```

## Usage
The usage of the package is relative simple. The following instructions go through the steps.

#### Authorize
To use the package, you need to authorize using the clientID and clientSecret obtain from [Spotify API](https://developer.spotify.com/)

```
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