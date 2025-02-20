package cmd

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleProfile(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleProfile()
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Displays your AniList profile (requires login)",
	Run:   handleProfile,
}
