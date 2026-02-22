package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleNextDate(_ *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(ui.ErrorText(errors.New("invalid media id provided")))
		return
	}

	err = viewmodel.HandleMediaNextDate(id)
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaNextDateCmd = &cobra.Command{
	Use:     "nextdate [id]",
	Short:   "Shows next episode release date for an anime",
	Aliases: []string{"next"},
	Run:     handleNextDate,
	Args:    cobra.ExactArgs(1),
}
