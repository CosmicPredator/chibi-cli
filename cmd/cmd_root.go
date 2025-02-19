package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var isVersionCmd bool
var appVersion string = "develop"

var rootCmd = &cobra.Command{
	Use:  "chibi",
	Long: "Chibi for AniList - A lightweight anime & manga tracker CLI app powered by AniList.",
	Run: func(cmd *cobra.Command, args []string) {
		if isVersionCmd {
			fmt.Println(SUCCESS_MESSAGE_TEMPLATE.Render(appVersion))
		} else {
			cmd.Help()
		}
	},
	Version: appVersion,
}

func Execute(version string) {
	appVersion = version
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		loginCmd,
		profileCmd,
		mediaSearchCmd,
		mediaListCmd,
		mediaUpdateCmd,
		mediaAddCmd,
	)
	rootCmd.Execute()
}
