package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/x/term"
	"github.com/spf13/cobra"
)

var listMediaType string
var listStatus string

func handleLs() {
	CheckIfTokenExists()

	var mediaType string

	switch listMediaType {
	case "anime", "a":
		mediaType = "ANIME"
	case "manga", "m":
		mediaType = "MANGA"
	}

	mediaList := internal.NewMediaList()
	err := mediaList.Get(mediaType, listStatus)
	if err != nil {
		ErrorMessage(err.Error())
	}

	if len(mediaList.Data.MediaListCollection.Lists) == 0 {
		fmt.Println(
			ERROR_MESSAGE_TEMPLATE.Render("No entires found!"),
		)
		os.Exit(0)
	}

	rows := [][]string{}

	for _, lists := range mediaList.Data.MediaListCollection.Lists {
		for _, entry := range lists.Entries {
			var progress string
			if mediaType == "ANIME" {
				progress = fmt.Sprintf("%d/%d", entry.Progress, entry.Media.Episodes)
			} else {
				progress = fmt.Sprintf("%d/%d", entry.Progress, entry.Media.Chapters)
			}

			if lists.Status == "REPEATING" {
				entry.Media.Title.UserPreferred = "(R) " + entry.Media.Title.UserPreferred
			}

			rows = append(rows, []string{
				strconv.Itoa(entry.Media.Id),
				entry.Media.Title.UserPreferred,
				progress,
			})
		}
	}

	for _, i := range mediaList.Data.MediaListCollection.Lists[0].Entries {
		var progress string
		if mediaType == "ANIME" {
			progress = fmt.Sprintf("%d/%d", i.Progress, i.Media.Episodes)
		} else {
			progress = fmt.Sprintf("%d/%d", i.Progress, i.Media.Chapters)
		}

		rows = append(rows, []string{
			strconv.Itoa(i.Media.Id),
			i.Media.Title.UserPreferred,
			progress,
		})
	}

	// get size of terminal
	tw, _, err := term.GetSize((os.Stdout.Fd()))
	if err != nil {
		ErrorMessage(err.Error())
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			// style for table header row
			if row == -1 {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).Align(lipgloss.Center)
			}

			// force title column to wrap by specifying terminal width
			if col == 1 {
				return lipgloss.NewStyle().Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2).Width(tw)
			}

			return lipgloss.NewStyle().Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)
		}).
		Headers("ID", "TITLE", "PROGRESS").
		Rows(rows...).Width(tw)

	fmt.Println(t)
}

var mediaListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your current anime/manga list",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		handleLs()
	},
}

func init() {
	mediaListCmd.Flags().StringVarP(
		&listMediaType, "type", "t", "anime", "Type of media. for anime, pass 'anime' or 'a', for manga, use 'manga' or 'm'",
	)
	mediaListCmd.Flags().StringVarP(
		&listStatus, "status", "s", "watching", "Status of the media. Can be 'watching/w or reading/r', 'planning/p', 'completed/c', 'dropped/d', 'paused/ps'",
	)
}
