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

type MediaSearchUI struct {
	MediaList *[]responses.MediaSearchList
}

func (ms *MediaSearchUI) renderTable(rows ...[]string) (*table.Table, error) {
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
					PaddingLeft(2).
					PaddingRight(2).
					Width((tw - 6) / 3).Inline(true)
				if row%2 == 0 {
					colStyle = colStyle.Faint(true)
				}
				return colStyle
			}

			return lipgloss.
				NewStyle().
				Align(lipgloss.Center).
				PaddingLeft(2).
				PaddingRight(2)
		}).
		Headers("ID", "TITLE", "FORMAT", "SCORE").
		Rows(rows...).Width(tw)

	return table, nil
}

func (ms *MediaSearchUI) Render() error {
	rows := [][]string{}

	for _, media := range *ms.MediaList {
		var averageScore string
		if media.AverageScore == nil {
			averageScore = "?"
		} else {
			averageScore = fmt.Sprintf("%.2f", *media.AverageScore)
		}

		rows = append(rows, []string{
			strconv.Itoa(media.Id),
			media.Title.UserPreferred,
			internal.MediaFormatFormatter(media.MediaFormat),
			averageScore,
		})
	}

	table, err := ms.renderTable(rows...)
	if err != nil {
		return err
	}

	fmt.Println(table)
	return nil
}
