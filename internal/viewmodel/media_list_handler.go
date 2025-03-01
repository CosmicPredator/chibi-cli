package viewmodel

import (
	"context"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/db"
	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func for "chibi ls" command
func HandleMediaList(mediaType, mediaStatus string) error {
	mediaType = internal.MediaTypeEnumMapper(mediaType)
	mediaStatus = internal.MediaStatusEnumMapper(mediaStatus)

	// get user id
	dbCtx, err := db.NewDbConn(false)
	if err != nil {
		return err
	}
	userId, err := dbCtx.GetConfig("user_id")
	if err != nil {
		return err
	}

	userIdInt, err := strconv.Atoi(*userId)
	if err != nil {
		return err
	}

	// if status arg is "watching", the include both
	// current and repeating
	var mediaStatuIn []string
	if mediaStatus == "CURRENT" {
		mediaStatuIn = []string{mediaStatus, "REPEATING"}
	} else {
		mediaStatuIn = []string{mediaStatus}
	}

	// perform media list API request
	var mediaList *responses.MediaList
	err = ui.ActionSpinner("Fetching lists...", func(ctx context.Context) error {
		mediaList, err = api.GetMediaList(
			userIdInt, mediaStatuIn,
		)
		return err
	})
	if err != nil {
		return err
	}

	// display the result
	mediaListUI := ui.MediaListUI{
		MediaType: mediaType,
		MediaList: mediaList,
	}

	err = mediaListUI.Render()
	return err
}
