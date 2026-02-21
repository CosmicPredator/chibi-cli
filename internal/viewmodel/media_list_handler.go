package viewmodel

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/kvdb"
	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func for "chibi ls" command
func HandleMediaList(mediaType, mediaStatus string) error {
	mediaType = internal.MediaTypeEnumMapper(mediaType)
	mediaStatus = internal.MediaStatusEnumMapper(mediaStatus)

	// get user id
	db, err := kvdb.Open()
	if err != nil {
		return fmt.Errorf("unable to open databse: %w", err)
	}
	defer db.Close()
	
	userId, err := db.Get(context.TODO(), "user_id")
	if err != nil {
		return errors.New("not logged in. Please use \"chibi login\" to continue")
	}

	userIdInt, err := strconv.Atoi(string(userId))
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
