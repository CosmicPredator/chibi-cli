package viewmodel

import (
	"context"
	"errors"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleMediaNextDate(id int) error {
	var mediaNextDate *responses.MediaNextAiring
	var err error

	err = ui.ActionSpinner("Fetching next episode date...", func(ctx context.Context) error {
		mediaNextDate, err = api.GetMediaNextAiringDate(id)
		return err
	})
	if err != nil {
		return err
	}

	if mediaNextDate.Data.Media.Type != string(internal.ANIME) {
		return errors.New("next episode date is only available for anime entries")
	}

	mediaNextDateUI := &ui.MediaNextDateUI{
		ID:     mediaNextDate.Data.Media.ID,
		Title:  mediaNextDate.Data.Media.Title.UserPreferred,
		Status: mediaNextDate.Data.Media.Status,
	}
	if mediaNextDate.Data.Media.NextAiringEpisode != nil {
		nextEpisode := mediaNextDate.Data.Media.NextAiringEpisode.Episode
		airingAt := mediaNextDate.Data.Media.NextAiringEpisode.AiringAt
		timeUntilAiring := mediaNextDate.Data.Media.NextAiringEpisode.TimeUntilAiring

		mediaNextDateUI.NextEpisode = &nextEpisode
		mediaNextDateUI.AiringAtUnix = &airingAt
		mediaNextDateUI.TimeUntilAiring = &timeUntilAiring
	}

	return mediaNextDateUI.Render()
}
