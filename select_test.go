package main

import (
	"testing"
)

func TestFindDrmShows(t *testing.T) {

	tvshow := []TvShow{
		{
			Drm:          true,
			EpisodeCount: 0,
			Slug:         "slug1",
			Title:        "title1",
		},
		{
			Drm:          true,
			EpisodeCount: 1,
			Slug:         "slug2",
			Title:        "title2",
		},
	}
	resp := FindDrmShows(tvshow)
	if resp.Response[0].Title != "title2" {
		t.Errorf("FindDrmShows response was incorrect, got: %s, want: %s.", resp.Response[0].Title, "title2")
	}
}
