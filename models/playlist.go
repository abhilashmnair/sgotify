package models

type Playlist struct {
	Collaborative bool          `json:"collaborative"`
	Description   string        `json:"description"`
	ExternalURLs  ExternalURL   `json:"external_urls"`
	Followers     Followers     `json:"followers"`
	Hyperlink     string        `json:"href"`
	Id            string        `json:"id"`
	Image         []ImageObject `json:"images"`
	Name          string        `json:"name"`
	Owner         Owner         `json:"owner"`
	Public        bool          `json:"public"`
	SnapshotId    string        `json:"snapshot_id"`
	Tracks        struct {
		Hyperlink   string          `json:"href"`
		Limit       int             `json:"limit"`
		Next        string          `json:"next"`
		Offset      int             `json:"offset"`
		Previous    string          `json:"previous"`
		TotalTracks int             `json:"total"`
		Items       []PlaylistTrack `json:"items"`
	} `json:"tracks"`
}
