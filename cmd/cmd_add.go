package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/spf13/cobra"
)

var mediaAddStatus string

func handleMediaAdd(mediaId int) {
	CheckIfTokenExists()

	switch mediaAddStatus {
	case "watching", "reading", "w", "r":
		mediaAddStatus = "CURRENT"
	case "planning", "p":
		mediaAddStatus = "PLANNING"
	case "completed", "c":
		mediaAddStatus = "COMPLETED"
	case "dropped", "d":
		mediaAddStatus = "DROPPED"
	case "paused", "ps":
		mediaAddStatus = "PAUSED"
	case "repeating", "rp":
		mediaAddStatus = "REPEATING"
	default:
		mediaAddStatus = "CURRENT"
	}

	mediaUpdate := internal.NewMediaUpdate()
	if mediaAddStatus == "CURRENT" {
		currDate := fmt.Sprintf("%d/%d/%d", time.Now().Day(), time.Now().Month(), time.Now().Year())
		err := mediaUpdate.Get(true, mediaId, "0", mediaAddStatus, currDate)
		if err != nil {
			ErrorMessage(err.Error())
		}
	} else {
		err := mediaUpdate.Get(true, mediaId, "0", mediaAddStatus, "")
		if err != nil {
			ErrorMessage(err.Error())
		}
	}

	fmt.Println(
		SUCCESS_MESSAGE_TEMPLATE.Render("Done ✅"),
	)
}

var mediaAddCmd = &cobra.Command{
	Use:   "add [id]",
	Short: "Add a media to your list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mediaId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ERROR_MESSAGE_TEMPLATE.Render(
				"Invalid media ID. Please provide a valid one",
			))
			os.Exit(0)
		}
		handleMediaAdd(mediaId)
	},
}

func init() {
	mediaAddCmd.Flags().StringVarP(
		&mediaAddStatus,
		"status",
		"s",
		"planning",
		"Status of the media. Can be 'watching/w or reading/r', 'planning/p', 'completed/c', 'dropped/d', 'paused/ps', 'repeating/rp'",
	)
}
