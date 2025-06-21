package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

var pageSize int
var searchMediaType string

func handleMediaSearch(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("No seach queries provided")
		os.Exit(0)
	}

	combinedQuery := strings.Join(args, "")
	err := viewmodel.HandleMediaSearch(
		combinedQuery,
		searchMediaType,
		pageSize,
	)

	if err != nil {
		fmt.Println(ui.ErrorText(err))
	}
}

var mediaSearchCmd = &cobra.Command{
	Use:   "search [query...]",
	Short: "Search for anime and manga",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleMediaSearch,
}

func init() {
	mediaSearchCmd.Flags().StringVarP(
		&searchMediaType, "type", "t", "anime", "Type of media. for anime, pass 'anime' or 'a', for manga, use 'manga' or 'm'")
	mediaSearchCmd.Flags().IntVarP(&pageSize, "page", "p", 10, "The number of results to be returned")
}
