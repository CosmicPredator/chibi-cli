package viewmodel

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/db"
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/charmbracelet/huh/spinner"
)


type MediaUpdateParams struct {
	IsNewAddition bool
	MediaId int
	Progress string
	Status string
	StartDate string
}

func getCurrentProgress(userId int, mediaId int) (current int, total *int, err error) {
	var mediaList *responses.MediaList
	err = spinner.New().Title("Getting your list...").Action(func() {
		mediaList, err = api.GetMediaList(
			userId,
			[]string{ "CURRENT", "REPEATING" },
		)
	}).Run()

	if err != nil {
		return
	}

	for _, list := range mediaList.Data.AnimeListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				current = entry.Progress
				total = entry.Media.Episodes
				return
			}
		}
	}

	for _, list := range mediaList.Data.MangaListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				current = entry.Progress
				total = entry.Media.Chapters
				return
			}
		}
	}

	return
}

// func handleNewAddition(params MediaUpdateParams) error {
// 	return nil
// }

func HandleMediaUpdate(params MediaUpdateParams) error {
	// if params.IsNewAddition {
	// 	handleNewAddition(params)
	// }

	dbCtx, err := db.NewDbConn(false)
	if err != nil {
		return err
	}
	userId, err := dbCtx.GetConfig("user_id")
	if err != nil {
		return err
	}

	userIdInt, _ := strconv.Atoi(*userId)
	current, total, err := getCurrentProgress(userIdInt, params.MediaId)
	if err != nil {
		return err
	}

	accumulatedProgress, err := parseRelativeProgress(params.Progress, current)
	if err != nil {
		return err
	}

	if total != nil && *total != 0 && accumulatedProgress > *total {
		return fmt.Errorf("entered value is greater than total episodes / chapters, which is %d", total)
	}

	var response *responses.MediaUpdateResponse
	err = spinner.New().Title("Updating entry...").Action(func() {
		response, err = api.UpdateMediaEntry(map[string]any{
			"id": params.MediaId,
			"progress": accumulatedProgress,
		})
	}).Run()
	if err != nil {
		return err
	}

	fmt.Println(
		ui.SuccessText(
			fmt.Sprintf(
				"Progress updated for %s (%d -> %d)\n", 
				response.Data.SaveMediaListEntry.Media.Title.UserPreferred, 
				current, accumulatedProgress),
		),
	)

	return nil
}

func parseRelativeProgress(progress string, current int) (int, error) {
	var accumulatedProgress int
	if strings.Contains(progress, "+") || strings.Contains(progress, "-") {
		if progress[:1] == "+" {
			prgInt, _ := strconv.Atoi(progress[1:])
			accumulatedProgress = current + prgInt
		} else {
			if current == 0 {
				accumulatedProgress = 0
			} else {
				prgInt, _ := strconv.Atoi(progress[1:])
				accumulatedProgress = current - prgInt
			}
		}
	} else {
		pgrInt, err := strconv.Atoi(progress)
		if err != nil {
			return 0, err
		}
		accumulatedProgress = pgrInt
	}
	return accumulatedProgress, nil
}