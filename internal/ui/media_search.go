package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/charmbracelet/lipgloss"
)

type MediaSearchUI struct {
	MediaList *[]responses.MediaSearchList
}

type MediaSearchResult struct {
	Id string
	Title string
	Format string
	Score string
}

func (l *MediaSearchUI) renderColumn(entries ...*MediaSearchResult) string {
	col := func(w int) lipgloss.Style {
		return lipgloss.NewStyle().Width(w).MarginRight(2).Align(lipgloss.Right)
	}

	styles := []lipgloss.Style{
		col(7),
		col(8),
		col(8),
		col(0),
	}

	headerStyle := func (style lipgloss.Style) lipgloss.Style {
		return style.MarginBottom(1).Underline(true).Bold(true).Foreground(lipgloss.ANSIColor(5))
	}

	var sb strings.Builder
	header := []string{
		headerStyle(styles[0]).Render("ID"),
		headerStyle(styles[1]).Render("FORMAT"),
		headerStyle(styles[2]).Render("SCORE"),
		headerStyle(styles[3]).Render("TITLE"),
	}

	sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, header...) + "\n")
	for _, entry := range entries {
		row := []string{
			styles[0].Foreground(lipgloss.ANSIColor(6)).Render(entry.Id),
			styles[1].Foreground(lipgloss.ANSIColor(2)).Render(entry.Format),
			styles[2].Foreground(lipgloss.ANSIColor(3)).Render(entry.Score),
			styles[3].Render(entry.Title),
		}
		sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...) + "\n")
	}

	return sb.String()
}

// render UI string
func (ms *MediaSearchUI) Render() error {
	var rows []*MediaSearchResult = make([]*MediaSearchResult, 0)

	for _, media := range *ms.MediaList {
		var averageScore string
		if media.AverageScore == nil {
			averageScore = "?"
		} else {
			averageScore = fmt.Sprintf("%.0f%%", *media.AverageScore)
		}

		rows = append(rows, &MediaSearchResult{
			Id: strconv.Itoa(media.Id),
			Title: media.Title.UserPreferred,
			Format: internal.MediaFormatFormatter(media.MediaFormat),
			Score: averageScore,
		})
	}

	fmt.Println(ms.renderColumn(rows...))
	return nil
}
