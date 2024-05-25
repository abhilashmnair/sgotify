package models

type ImageObject struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type LinkedFrom struct {
	ExternalURLs struct {
		URL string `json:"spotify"`
	} `json:"external_urls"`
	Hyperlink string `json:"href"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	URI       string `json:"uri"`
}

type Copyright struct {
	Description string `json:"text"`
	Type        string `json:"type"`
}

type ExternalID struct {
	UPC  string `json:"upc"`
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
}

type ExternalURL struct {
	URL string `json:"spotify"`
}

type Followers struct {
	Hyperlink string
	Count     int
}

type Owner struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	Followers    Followers   `json:"followers"`
	Hyperlink    string      `json:"href"`
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
	Name         string      `json:"display_name"`
}
