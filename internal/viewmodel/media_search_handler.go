package viewmodel

import (
	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleMediaSearch(searchQuery string, mediaType string, perPage int) error {
	mediaType = internal.MediaTypeEnumMapper(mediaType)

	searchResult, err := api.SearchMedia(searchQuery, perPage, mediaType)
	if err != nil {
		return err
	}

	mediaSearchUI := ui.MediaSearchUI{
		MediaList: &searchResult.Data.Page.Media,
	}
	err = mediaSearchUI.Render()
	return err
}