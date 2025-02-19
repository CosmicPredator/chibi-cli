package cmd

import (
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleLoginCmd(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleLogin()
	if err != nil {
		ui.ErrorText(err)
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with anilist",
	Run: handleLoginCmd,
}
