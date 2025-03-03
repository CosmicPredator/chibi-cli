package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

var mediaAddStatus string

func handleMediaAdd(cmd *cobra.Command, args []string) {
	mediaId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(
			ui.ErrorText(errors.New("invalid media ID. Please provide a valid one")),
		)
		os.Exit(0)
	}
	err = viewmodel.HandleMediaUpdate(
		viewmodel.MediaUpdateParams{
			IsNewAddition: true,
			MediaId:       mediaId,
			Status:        mediaAddStatus,
		},
	)
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaAddCmd = &cobra.Command{
	Use:   "add [id]",
	Short: "Add a media to your list",
	Args:  cobra.ExactArgs(1),
	Run:   handleMediaAdd,
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
