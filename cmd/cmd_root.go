package cmd

import (
	"context"
	"fmt"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "chibi",
	Long: "Chibi for AniList - A lightweight anime & manga tracker CLI app powered by AniList.\nRead the documentation at https://chibi-cli.pages.dev/",
}

func Execute(version string) {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		loginCmd,
		logoutCmd,
		profileCmd,
		mediaSearchCmd,
		mediaListCmd,
		mediaUpdateCmd,
		mediaAddCmd,
		mediaInfoCmd,
	)
	if err := fang.Execute(
		context.TODO(), 
		rootCmd, 
		fang.WithVersion(version),
		); err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}
