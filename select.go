package main

var drmShows DrmShows

func FindDrmShows(t TvShows) (serverResponse ServerResponse) {

	for _, show := range t {
		if show.Drm == true && show.EpisodeCount > 0 {
			dShow := new(DrmShow)
			dShow.Image = show.Image.ShowImage
			dShow.Slug = show.Slug
			dShow.Title = show.Title
			serverResponse.Response = append(serverResponse.Response, *dShow)
		}
	}

	return serverResponse
}
