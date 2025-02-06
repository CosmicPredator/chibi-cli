package internal

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/CosmicPredator/chibi/types"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type MediaListUpdate struct {
	Data struct {
		AnimeListCollection ListCollection `json:"AnimeListCollection"`
		MangaListCollection ListCollection `json:"MangaListCollection"`
	} `json:"data"`
}

func getTotalCurrent(mediaId int) (int, int, error) {
	query := `query ($id: Int) {
		AnimeListCollection: MediaListCollection(userId: $id, type: ANIME, status_in:[CURRENT, REPEATING]){
			lists {
				status
				entries {
					progress
					media {
						id
						title {
							userPreferred
						}
						episodes
						chapters
					}
				}
			}
		}
		MangaListCollection: MediaListCollection(userId: $id, type: MANGA, status_in:[CURRENT, REPEATING]){
			lists {
				status
				entries {
					progress
					media {
						id
						title {
							userPreferred
						}
						episodes
						chapters
					}
				}
			}
		}
	}`

	tokenConfig := types.NewTokenConfig()
	err := tokenConfig.ReadFromJsonFile()

	if err != nil {
		return 0, 0, err
	}

	var responseObj MediaListUpdate

	anilistClient := NewAnilistClient()
	err = anilistClient.ExecuteGraqhQL(
		query,
		map[string]interface{}{
			"id": tokenConfig.UserId,
		},
		&responseObj,
	)

	if err != nil {
		return 0, 0, err
	}

	for _, list := range responseObj.Data.AnimeListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				return entry.Progress, entry.Media.Episodes, nil
			}
		}
	}

	for _, list := range responseObj.Data.MangaListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Media.Id == mediaId {
				return entry.Progress, entry.Media.Chapters, nil
			}
		}
	}

	return 0, 0, errors.New("list empty")
}

type updateMediaFields struct {
	CompletedAtDate  int
	CompletedAtMonth int
	CompletedAtYear  int
	Notes            string
	Score            float32
}

func updateMediaEntry() (*updateMediaFields, error) {
	mediaFields := &updateMediaFields{}
	currDate := fmt.Sprintf("%d/%d/%d\n", time.Now().Day(), time.Now().Month(), time.Now().Year())
	var scoreString string

	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Date of completion").
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
				Value(&mediaFields.Notes),
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
		return nil, err
	}
	scoreFloat, err := strconv.ParseFloat(scoreString, 32)
	if err != nil {
		return nil, err
	}

	mediaFields.CompletedAtDate = completedDate.Day()
	mediaFields.CompletedAtMonth = int(completedDate.Month())
	mediaFields.CompletedAtYear = completedDate.Year()
	mediaFields.Score = float32(scoreFloat)

	return mediaFields, nil
}

type MediaUpdate struct {
	Data struct {
		SaveMediaListEntry struct {
			MediaId int `json:"mediaId"`
		} `json:"SaveMediaListEntry"`
	} `json:"data"`
}

func (mu *MediaUpdate) Get(isMediaAdd bool, mediaId int, progress string, status string, startDate string) error {
	if status == "" {
		status = "COMPLETED"
	}
	var accumulatedProgress int
	current, total, err := getTotalCurrent(mediaId)

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
	}

	if err != nil {
		return err
	}

	if total != 0 && accumulatedProgress > total {
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(
			fmt.Sprintf("Entered value is greater than total episodes / chapters, which is %d", total),
		))
		os.Exit(0)
	}

	mutation :=
		`mutation(
        $id: Int, 
        $progress: Int,
        $score: Float,
        $notes: String,
        $cDate: Int,
        $cMonth: Int,
        $cYear: Int,
        $sDate: Int,
        $sMonth: Int,
        $sYear: Int,
        $status: MediaListStatus
    ) {
        SaveMediaListEntry(
            mediaId: $id, 
            progress: $progress,
            status: $status,
            score: $score,
            notes: $notes,
            completedAt: {
                day: $cDate,
                month: $cMonth,
                year: $cYear
            },
            startedAt: {
                day: $sDate,
                month: $sMonth,
                year: $sYear
            }
        ) {
            mediaId
        }
    }`

	if isMediaAdd {
		variables := map[string]interface{}{
			"id":     mediaId,
			"status": status,
		}

		if startDate != "" {
			startDateRaw, err := time.Parse("02/01/2006", startDate)
			if err != nil {
				return err
			}

			if status == "CURRENT" {
				variables["sDate"] = startDateRaw.Day()
				variables["sMonth"] = int(startDateRaw.Month())
				variables["sYear"] = startDateRaw.Year()
			}
		}

		err = NewAnilistClient().ExecuteGraqhQL(
			mutation,
			variables,
			&mu,
		)
		return err
	}

	var canEditList bool = false

	if accumulatedProgress == total {
		huh.NewConfirm().
			Title("Seems like you completed the anime/manga. Do you want to mark this as completed?").
			Affirmative("Yes!").
			Negative("No").
			Value(&canEditList).
			Run()
	}

	if canEditList {
		mediaFields, err := updateMediaEntry()
		if err != nil {
			return err
		}

		err = NewAnilistClient().ExecuteGraqhQL(
			mutation,
			map[string]interface{}{
				"id":       mediaId,
				"progress": accumulatedProgress,
				"score":    mediaFields.Score,
				"notes":    mediaFields.Notes,
				"cDate":    mediaFields.CompletedAtDate,
				"cMonth":   mediaFields.CompletedAtMonth,
				"cYear":    mediaFields.CompletedAtYear,
			},
			&mu,
		)
		return err
	}

	err = NewAnilistClient().ExecuteGraqhQL(
		mutation,
		map[string]interface{}{
			"id":       mediaId,
			"progress": accumulatedProgress,
		},
		&mu,
	)
	return err
}

func NewMediaUpdate() *MediaUpdate {
	return &MediaUpdate{}
}
