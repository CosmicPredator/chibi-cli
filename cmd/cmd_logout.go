package cmd

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleLogoutCommand(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleLogout()
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var logoutCmd = &cobra.Command{
	Use: "logout",
	Short: "logs you out from anilist",
	Run: handleLogoutCommand,
}