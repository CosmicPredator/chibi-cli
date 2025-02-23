package cmd

import (
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

// // TODO: Update progress relatively. For example "+2", "-10" etc.,
var progress string

// func handleUpdate(mediaId int) {
// 	CheckIfTokenExists()

// 	progressInt, err := strconv.Atoi(progress)
// 	if err == nil {
// 		if progressInt == 0 {
// 			fmt.Println(
// 				ERROR_MESSAGE_TEMPLATE.Render("The flag 'progress' should be greater than 0."),
// 			)
// 		}
// 	}

// 	mediaUpdate := internal.NewMediaUpdate()
// 	err = mediaUpdate.Get(false, mediaId, progress, "", "")

// 	if err != nil {
// 		ErrorMessage(err.Error())
// 	}
// 	fmt.Println(
// 		SUCCESS_MESSAGE_TEMPLATE.Render(
// 			"Done ✅",
// 		),
// 	)
// }

func handleUpdate(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		progress = args[1]
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(
			ERROR_MESSAGE_TEMPLATE.Render("Invalid media id. please provide a valid one..."),
		)
	}

	err = viewmodel.HandleMediaUpdate(
		viewmodel.MediaUpdateParams{
			IsNewAddition: false,
			MediaId: id,
			Progress: progress,
			Status: "none",
			StartDate: "none",
		},
	)

	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaUpdateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a list entry",
	Args:  cobra.MinimumNArgs(1),
	Run: handleUpdate,
}

func init() {
	mediaUpdateCmd.Flags().StringVarP(
		&progress,
		"progress",
		"p",
		"0",
		"The number of episodes/chapter to update",
	)
}
