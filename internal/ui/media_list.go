package ui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/x/term"
)

type MediaListUI struct {
	MediaType string
	MediaList *responses.MediaList
}

// table renderer for media list
func (l *MediaListUI) renderTable(rows ...[]string) (*table.Table, error) {
	// get size of terminal
	tw, _, err := term.GetSize((os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	tw = int(float32(tw) * float32(0.9))

	table := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			// style for table header row
			if row == -1 {
				return lipgloss.
					NewStyle().
					Foreground(lipgloss.Color("#FF79C6")).
					Bold(true).
					Align(lipgloss.Center)
			}

			if row%2 == 0 && (col == 0 || col == 2 || col == 3) {
				return lipgloss.NewStyle().Align(lipgloss.Center).Faint(true)
			}

			// force title column to wrap by specifying terminal width
			if col == 1 {
				colStyle := lipgloss.
					NewStyle().
					Align(lipgloss.Center).
					PaddingLeft(1).
					PaddingRight(1).
					Width((tw - 6) / 3).Inline(true)
				if row%2 == 0 {
					colStyle = colStyle.Faint(true)
				}
				return colStyle
			}

			return lipgloss.NewStyle().Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)
		}).
		Headers("ID", "TITLE", "FORMAT", "PROGRESS").
		Rows(rows...).Width(tw)

	return table, nil
}

func (l *MediaListUI) Render() error {
	rows := [][]string{}

	var selectedList responses.ListCollection

	if internal.MediaType(l.MediaType) == internal.ANIME {
		selectedList = l.MediaList.Data.AnimeListCollection
	} else {
		selectedList = l.MediaList.Data.MangaListCollection
	}

	for _, list := range selectedList.Lists {
		for _, entry := range list.Entries {
			var progress string

			if l.MediaType == string(internal.ANIME) {
				var total string
				if entry.Media.Episodes == nil {
					total = "?"
				} else {
					total = strconv.Itoa(*entry.Media.Episodes)
				}

				progress = fmt.Sprintf("%v/%v", entry.Progress, total)
			} else {
				var total string
				if entry.Media.Chapters == nil {
					total = "?"
				} else {
					total = strconv.Itoa(*entry.Media.Chapters)
				}
				progress = fmt.Sprintf("%v/%v", entry.Progress, total)
			}

			if list.Status == "REPEATING" {
				entry.Media.Title.UserPreferred = "(R) " + entry.Media.Title.UserPreferred
			}

			rows = append(rows, []string{
				strconv.Itoa(entry.Media.Id),
				entry.Media.Title.UserPreferred,
				internal.MediaFormatFormatter(entry.Media.MediaFormat),
				progress,
			})
		}
	}

	table, err := l.renderTable(rows...)
	if err != nil {
		return err
	}

	fmt.Println(table)

	return nil
}
