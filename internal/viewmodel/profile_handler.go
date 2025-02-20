package viewmodel

import (
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleProfile() error {
	// get profile info from API
	profile, err := api.GetUserProfile()
	if err != nil {
		return err
	}

	// populate ProfileUI struct fields with the data from API
	profileUI := ui.ProfileUI{
		Id:             profile.Data.Viewer.Id,
		Name:           profile.Data.Viewer.Name,
		TotalAnime:     profile.Data.Viewer.Statistics.Anime.Count,
		TotalManga:     profile.Data.Viewer.Statistics.Manga.Count,
		MinutesWatched: profile.Data.Viewer.Statistics.Anime.MinutesWatched,
		ChaptersRead:   profile.Data.Viewer.Statistics.Manga.ChaptersRead,
		SiteUrl:        profile.Data.Viewer.SiteUrl,
	}

	// display profile UI
	err = profileUI.Render()
	return err
}
