package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleMediaInfo(_ *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(ui.ErrorText(errors.New("invalid media id provided")))
	}
	
	err = viewmodel.HandleMediaInfo(id)
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaInfoCmd = &cobra.Command{
	Use: "info [id]",
	Short: "Displays general information of media",
	Run: handleMediaInfo,
	Args: cobra.MinimumNArgs(1),
}