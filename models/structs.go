package models

type ExplicitContentStruct struct {
	Filter_enabled bool `json:"filter_enabled"`
	Filter_locked  bool `json:"filter_locked"`
}

type ExternalUrlsStruct struct {
	Spotify string `json:"spotify"`
}

type FollowersStruct struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ImagesStruct struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Profile struct {
	Country          string `json:"country"`
	Display_name     string `json:"display_name"`
	Email            string `json:"email"`
	Explicit_content ExplicitContentStruct
	External_urls    ExternalUrlsStruct
	Followers        FollowersStruct
	Href             string `json:"href"`
	Id               string `json:"id"`
	Images           []ImagesStruct
	Product          string `json:"product"`
	Type             string `json:"type"`
	Uri              string `json:"uri"`
}

type Device struct {
	Id                 string `json:"id"`
	Is_active          bool   `json:"is_active"`
	Is_private_session bool   `json:"is_private_session"`
	Is_restricted      bool   `json:"is_restricted"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Volume_percent     int    `json:"volume_percent"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}

type Context struct {
	Type          string             `json:"type"`
	Href          string             `json:"href"`
	External_urls ExternalUrlsStruct `json:"external_urls"`
	Uri           string             `json:"uri"`
}

type TrackRestriction struct {
	Reason string `json:"reason"`
}

type TrackCopyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type TrackExternalIDs struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type AlbumArtist struct {
	External_urls ExternalUrlsStruct `json:"external_urls"`
	Href          string             `json:"href"`
	Id            string             `json:"id"`
	Name          string             `json:"name"`
	Type          string             `json:"type"`
	Uri           string             `json:"uri"`
}

type TrackArtist struct {
	External_urls ExternalUrlsStruct `json:"external_urls"`
	Href          string             `json:"href"`
	Id            string             `json:"id"`
	Name          string             `json:"name"`
	Type          string             `json:"type"`
	Uri           string             `json:"uri"`
	Followers     FollowersStruct    `json:"followers"`
	Genres        []string           `json:"genres"`
	Images        []ImagesStruct     `json:"images"`
	Popularity    int                `json:"popularity"`
}

type TrackItem struct {
	Album             TrackAlbum         `json:"album"`
	Artists           []TrackArtist      `json:"artists"`
	Available_markets []string           `json:"available_markets"`
	Disc_number       int                `json:"disc_number"`
	Duration_ms       int                `json:"duration_ms"`
	Explicit          bool               `json:"explicit"`
	External_ids      TrackExternalIDs   `json:"external_ids"`
	External_urls     ExternalUrlsStruct `json:"external_urls"`
	Href              string             `json:"href"`
	Id                string             `json:"id"`
	Is_playable       bool               `json:"is_playable"`
	Restrictions      TrackRestriction   `json:"restrictions"`
	Name              string             `json:"name"`
	Popularity        int                `json:"popularity"`
	Preview_url       string             `json:"preview_url"`
	Track_number      int                `json:"track_number"`
	Type              string             `json:"type"`
	Uri               string             `json:"uri"`
	Is_local          bool               `json:"is_local"`
}

type TrackAlbum struct {
	Album_type             string             `json:"album_type"`
	Total_tracks           int                `json:"total_tracks"`
	Available_markets      []string           `json:"available_markets"`
	External_urls          ExternalUrlsStruct `json:"external_urls"`
	Href                   string             `json:"href"`
	Id                     string             `json:"id"`
	Images                 ImagesStruct       `json:"images"`
	Name                   string             `json:"name"`
	Release_date           string             `json:"release_date"`
	Release_date_precision string             `json:"release_date_precision"`
	Restrictions           TrackRestriction   `json:"restriction"`
	Type                   string             `json:"type"`
	Uri                    string             `json:"uri"`
	Copyrights             []TrackCopyright   `json:"copyrights"`
	External_ids           TrackExternalIDs   `json:"external_ids"`
	Genres                 []string           `json:"genres"`
	Label                  string             `json:"label"`
	Popularity             int                `json:"popularity"`
	Album_group            string             `json:"album_group"`
	Artists                []AlbumArtist      `json:"artists"`
}

type TrackActions struct {
	Interrupting_playback   bool `json:"interrupting_playback"`
	Pausing                 bool `json:"pausing"`
	Resuming                bool `json:"resuming"`
	Seeking                 bool `json:"seeking"`
	Skipping_next           bool `json:"skipping_next"`
	Skipping_prev           bool `json:"skipping_prev"`
	Toggling_repeat_context bool `json:"toggling_repeat_context"`
	Toggling_shuffle        bool `json:"toggling_shuffle"`
	Toggling_repeat_track   bool `json:"toggling_repeat_track"`
	Transferring_playback   bool `json:"transferring_playback"`
}

type GetCurrentlyPlayingResponse struct {
	Device                 Device       `json:"device"`
	Repeat_state           string       `json:"repeat_state"`
	Shuffle_state          string       `json:"shuffle_state"`
	Context                Context      `json:"context"`
	Timestamp              int          `json:"timestamp"`
	Progress_ms            int          `json:"progress_ms"`
	Is_playing             bool         `json:"is_playing"`
	Item                   TrackItem    `json:"item"`
	Currently_playing_type string       `json:"currently_playing_type"`
	Actions                TrackActions `json:"actions"`
}
