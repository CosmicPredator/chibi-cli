package cmd

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

var listMediaType string
var listStatus string

func handleLs(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleMediaList(listMediaType, listStatus)
	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your current anime/manga list",
	Aliases: []string{"ls"},
	Run:     handleLs,
}

func init() {
	mediaListCmd.Flags().StringVarP(
		&listMediaType, "type", "t", "anime", "Type of media. for anime, pass 'anime' or 'a', for manga, use 'manga' or 'm'",
	)
	mediaListCmd.Flags().StringVarP(
		&listStatus, "status", "s", "watching", "Status of the media. Can be 'watching/w or reading/r', 'planning/p', 'completed/c', 'dropped/d', 'paused/ps'",
	)
}
