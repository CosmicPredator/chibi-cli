package viewmodel

import (
	"context"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleMediaSearch(searchQuery string, mediaType string, perPage int) error {
	mediaType = internal.MediaTypeEnumMapper(mediaType)

	var err error
	var searchResult *responses.MediaSearch

	err = ui.ActionSpinner("Searching...", func(ctx context.Context) error {
		searchResult, err = api.SearchMedia(searchQuery, perPage, mediaType)
		return err
	})
	if err != nil {
		return err
	}

	mediaSearchUI := ui.MediaSearchUI{
		MediaList: &searchResult.Data.Page.Media,
	}
	err = mediaSearchUI.Render()
	return err
}
