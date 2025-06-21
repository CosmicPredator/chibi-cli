package viewmodel

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/credstore"
	"github.com/CosmicPredator/chibi/internal/ui"
)

type MediaUpdateParams struct {
	IsNewAddition bool
	MediaId       int
	Progress      string
	Status        string
	StartDate     string
	Notes         string
	Score         float32
}

// Gets current and total progress (episode/chapter) for given
// Media ID and returns it
func getCurrentProgress(userId int, mediaId int) (current int, total *int, err error) {
	var mediaList *responses.MediaList

	// load medialist collection
	err = ui.ActionSpinner("Getting your list...", func(ctx context.Context) error {
		mediaList, err = api.GetMediaList(
			userId,
			[]string{"CURRENT", "REPEATING"},
		)
		return err
	})

	if err != nil {
		return
	}

	// iterate over anime list collection and see if media ID matches
	for _, list := range mediaList.Data.AnimeListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				current = entry.Progress
				total = entry.Media.Episodes
				return
			}
		}
	}

	// iterate over manga list and see if media ID matches
	for _, list := range mediaList.Data.MangaListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				current = entry.Progress
				total = entry.Media.Chapters
				return
			}
		}
	}

	// if media id not in list, return default values
	return
}

// this func gets incvoked when "chibi add" command is invoked
func handleNewAdditionAction(params MediaUpdateParams) error {
	payload := map[string]any{
		"id":     params.MediaId,
		"status": internal.MediaStatusEnumMapper(params.Status),
	}

	// if passed status is watching and start date is empty,
	// fill the startData field with current date (today)
	if params.StartDate != "" {
		startDateRaw, err := time.Parse("02/01/2006", params.StartDate)
		if err != nil {
			return err
		}

		if payload["status"] == "CURRENT" {
			payload["sDate"] = startDateRaw.Day()
			payload["sMonth"] = int(startDateRaw.Month())
			payload["sYear"] = startDateRaw.Year()
		}
	} else {
		startDate := time.Now()

		if payload["status"] == "CURRENT" {
			payload["sDate"] = startDate.Day()
			payload["sMonth"] = int(startDate.Month())
			payload["sYear"] = startDate.Year()
		}
	}

	// perform API mutate request
	var response *responses.MediaUpdateResponse
	var err error
	err = ui.ActionSpinner("Adding entry...", func(ctx context.Context) error {
		response, err = api.UpdateMediaEntry(payload)
		return err
	})
	if err != nil {
		return err
	}

	// humanize strings for clear output
	var statusString string
	if internal.MediaStatusEnumMapper(params.Status) == "CURRENT" {
		statusString = "watching"
	} else {
		statusString = strings.ToLower(internal.MediaStatusEnumMapper(params.Status))
	}

	fmt.Println(
		ui.SuccessText(
			fmt.Sprintf(
				"Added %s to %s",
				response.Data.SaveMediaListEntry.Media.Title.UserPreferred,
				statusString,
			),
		),
	)

	return nil
}

// this func gets invoked when the current progress
// matches the total progress
func handleMediaCompletedAction(params MediaUpdateParams, progress int) error {
	defaultDate := fmt.Sprintf("%d/%02d/%d", time.Now().Day(), time.Now().Month(), time.Now().Year())
	var currDate string 
	var scoreString string
	var notes string

	// display a series of forms
	// 1. Completed Data
	for {
		input, err := ui.PrettyInput("Completed Date", defaultDate, func(s string) error {
			if strings.TrimSpace(s) == "" {
				return nil
			}
			layout := "02/01/2006"
			_, err := time.Parse(layout, strings.TrimSpace(s))
			return err
		})
		if err != nil {
			fmt.Println(ui.ErrorText(err))
			continue
		}
		currDate = strings.TrimSpace(input)
		if currDate == "" {
			currDate = defaultDate
		}
		break
	}

	// 2. Notes
	for {
		input, err := ui.PrettyInput("Notes", "", func(s string) error {
			return nil
		})
		if err != nil {
			fmt.Println(ui.ErrorText(err))
			continue
		}
		notes = input
		break
	}

	// 3. Score
	for {
		input, err := ui.PrettyInput("Score (use 1, 2, 3 for emojis)", "", func(s string) error {
			_, err := strconv.ParseFloat(s, 64)
			return err
		})
		if err != nil {
			fmt.Println(ui.ErrorText(err))
			continue
		}
		scoreString = input
		break
	}

	// parse form strings to API required data type
	completedDate, err := time.Parse("02/01/2006", strings.TrimSpace(currDate))
	if err != nil {
		return err
	}
	scoreFloat, err := strconv.ParseFloat(scoreString, 32)
	if err != nil {
		return err
	}

	// perform API mutation request
	var response *responses.MediaUpdateResponse
	err = ui.ActionSpinner("Marking as completed...", func(ctx context.Context) error {
		response, err = api.UpdateMediaEntry(map[string]any{
			"id":       params.MediaId,
			"progress": progress,
			"score":    scoreFloat,
			"notes":    notes,
			"cDate":    completedDate.Day(),
			"cMonth":   int(completedDate.Month()),
			"cYear":    completedDate.Year(),
		})
		return err
	})
	if err != nil {
		return err
	}

	// display success text
	fmt.Println(
		ui.SuccessText(
			fmt.Sprintf(
				"Marked %s as completed",
				response.Data.SaveMediaListEntry.Media.Title.UserPreferred),
		),
	)

	return nil
}

// handles media update logic and functionalities
// This func has 3 scenarios/routes
// 1. Invoke handleNewAdditionAction() when MediaUpdateParams.IsNewAddition is true
// 2. Invoke handleMediaCompletedAction() when current/accumulated progress == total progress
// 3. else go on with the flow (just progress update)
func HandleMediaUpdate(params MediaUpdateParams) error {
	// route 1
	if params.IsNewAddition {
		handleNewAdditionAction(params)
		return nil
	}

	userId, err := credstore.GetCredential("user_id")
	if err != nil {
		return err
	}

	userIdInt, _ := strconv.Atoi(*userId)
	current, total, err := getCurrentProgress(userIdInt, params.MediaId)
	if err != nil {
		return err
	}

	// parses relative progress (+2, -4) to current + relative progress
	accumulatedProgress, err := parseRelativeProgress(params.Progress, current)
	if err != nil {
		return err
	}

	status := internal.MediaStatusEnumMapper(params.Status)
	if status == "COMPLETED" {
		if *total != 0 && accumulatedProgress < *total {
			var markAsCompleted string
			fmt.Print("Accumulated progress is less than total episodes / chapters. Mark as media completed (y/N)? ")
			fmt.Scan(&markAsCompleted)

			if strings.ToLower(markAsCompleted) != "y" {
				return nil
			}
		}
		err = handleMediaCompletedAction(params, accumulatedProgress)
		return err
	}

	if total != nil {
		if *total != 0 && accumulatedProgress > *total {
			return fmt.Errorf("entered value is greater than total episodes / chapters, which is %d", *total)
		}

		// route 2
		if accumulatedProgress == *total {
			var markAsCompleted string
			fmt.Print("Mark as media completed (y/N)? ")
			fmt.Scan(&markAsCompleted)

			if strings.ToLower(markAsCompleted) == "y" {
				err = handleMediaCompletedAction(params, accumulatedProgress)
				return err
			}
			return nil
		}
	}

	var notes string
	if len(params.Notes) > 0 {
		notes = strings.ReplaceAll(params.Notes, `\n`, "\n")
	}

	// route 3
	var response *responses.MediaUpdateResponse
	err = ui.ActionSpinner("Updating entry...", func(ctx context.Context) error {
		payload := map[string]any{
			"id":       params.MediaId,
			"progress": accumulatedProgress,
			"status":   status,
			"notes":    notes,
		}
		if params.Score > 0 {
			payload["score"] = params.Score
		}
		response, err = api.UpdateMediaEntry(payload)
		return err
	})

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

// helper func to create absolute progress from relative progress
func parseRelativeProgress(progress string, current int) (int, error) {
	var accumulatedProgress int
	if len(progress) == 0 {
		return current, nil
	}
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
