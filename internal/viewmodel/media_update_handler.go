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
	"github.com/CosmicPredator/chibi/internal/db"
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type MediaUpdateParams struct {
	IsNewAddition bool
	MediaId       int
	Progress      string
	Status        string
	StartDate     string
	Notes         string
}

func getCurrentProgress(userId int, mediaId int) (current int, total *int, err error) {
	var mediaList *responses.MediaList
	err = spinner.New().Title("Getting your list...").Action(func() {
		mediaList, err = api.GetMediaList(
			userId,
			[]string{"CURRENT", "REPEATING"},
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

func handleNewAddition(params MediaUpdateParams) error {
	payload := map[string]any{
		"id":     params.MediaId,
		"status": internal.MediaStatusEnumMapper(params.Status),
	}

	if params.StartDate != "" {
		startDateRaw, err := time.Parse("02/01/2006", params.StartDate)
		if err != nil {
			return err
		}

		if params.Status == "CURRENT" {
			payload["sDate"] = startDateRaw.Day()
			payload["sMonth"] = int(startDateRaw.Month())
			payload["sYear"] = startDateRaw.Year()
		}
	}

	var response *responses.MediaUpdateResponse
	var err error
	err = ui.ActionSpinner("Adding entry...", func(ctx context.Context) error {
		response, err = api.UpdateMediaEntry(payload)
		return err
	})
	if err != nil {
		return err
	}

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

func handleMediaCompletedAction(params MediaUpdateParams, progress int) error {
	currDate := fmt.Sprintf("%d/%02d/%d\n", time.Now().Day(), time.Now().Month(), time.Now().Year())
	var scoreString string
	var notes string

	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Completed Date").
				Value(&currDate).
				Description("Date should be in format DD/MM/YYYY").
				Validate(func(s string) error {
					layout := "02/01/2006"
					_, err := time.Parse(layout, strings.TrimSpace(s))
					return err
				}),
		),
		huh.NewGroup(
			huh.NewText().
				Title("Notes").
				Description("Note: you can add multiple lines").
				Value(&notes),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Score").
				Description("If your score is in emoji, type 1 for 😞, 2 for 😐 and 3 for 😊").
				Prompt("> ").
				Validate(func(s string) error {
					_, err := strconv.ParseFloat(s, 64)
					return err
				}).
				Value(&scoreString),
		),
	).Run()
	completedDate, err := time.Parse("02/01/2006", strings.TrimSpace(currDate))
	if err != nil {
		return err
	}
	scoreFloat, err := strconv.ParseFloat(scoreString, 32)
	if err != nil {
		return err
	}

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

	fmt.Println(
		ui.SuccessText(
			fmt.Sprintf(
				"Marked %s as completed",
				response.Data.SaveMediaListEntry.Media.Title.UserPreferred),
		),
	)

	return nil
}

func HandleMediaUpdate(params MediaUpdateParams) error {
	if params.IsNewAddition {
		handleNewAddition(params)
		return nil
	}

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

	var response *responses.MediaUpdateResponse
	err = ui.ActionSpinner("Updating entry...", func(ctx context.Context) error {
		response, err = api.UpdateMediaEntry(map[string]any{
			"id":       params.MediaId,
			"progress": accumulatedProgress,
			"status":   status,
			"notes":    notes,
		})
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
