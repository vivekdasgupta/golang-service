package main

type ClientRequest struct {
	Payload      []TvShow `json:"payload"`
	Skip         uint64   `json:"skip"`
	Take         uint64   `json:"take"`
	TotalRecords uint64   `json:"totalRecords"`
}

type TvShow struct {
	Drm          bool   `json:"drm"`
	EpisodeCount uint64 `json:"episodeCount"`
	Image        struct {
		ShowImage string `json:"showImage"`
	} `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type TvShows []TvShow

type DrmShow struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type DrmShows []DrmShow

type ServerResponse struct {
	Response []DrmShow `json:"response"`
}

type ErrorResponse struct {
	ErrorMesg string `json:"error"`
}
