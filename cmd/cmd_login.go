package cmd

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

func handleLoginCmd(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleLogin()
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with anilist",
	Run:   handleLoginCmd,
}
