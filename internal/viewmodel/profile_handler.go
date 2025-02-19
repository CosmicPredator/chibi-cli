package viewmodel

import (
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleProfile() error {
	profile, err := api.GetUserProfile()
	if err != nil {
		return err
	}

	profileUI := ui.ProfileUI{
		Id: profile.Data.Viewer.Id,
		Name: profile.Data.Viewer.Name,
		TotalAnime: profile.Data.Viewer.Statistics.Anime.Count,
		TotalManga: profile.Data.Viewer.Statistics.Manga.Count,
		MinutesWatched: profile.Data.Viewer.Statistics.Anime.MinutesWatched,
		ChaptersRead: profile.Data.Viewer.Statistics.Manga.ChaptersRead,
		SiteUrl: profile.Data.Viewer.SiteUrl,
	}

	err = profileUI.Render()
	return err
}