package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/charmbracelet/lipgloss"
)

type MediaListUI struct {
	MediaType string
	MediaList *responses.MediaList
}

type MediaListEntry struct {
	Id string
	Title string
	Format string
	Progress string
}

func (l *MediaListUI) renderColumn(entries ...*MediaListEntry) string {
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
		headerStyle(styles[2]).Render("PROGRESS"),
		headerStyle(styles[3]).Render("TITLE"),
	}

	sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, header...) + "\n")
	for _, entry := range entries {
		row := []string{
			styles[0].Foreground(lipgloss.ANSIColor(6)).Render(entry.Id),
			styles[1].Foreground(lipgloss.ANSIColor(2)).Render(entry.Format),
			styles[2].Foreground(lipgloss.ANSIColor(3)).Render(entry.Progress),
			styles[3].Render(entry.Title),
		}
		sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...) + "\n")
	}

	return sb.String()
}

func (l *MediaListUI) Render() error {
	var rows []*MediaListEntry = make([]*MediaListEntry, 0)
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
				entry.Media.Title.UserPreferred = lipgloss.
					NewStyle().
					Foreground(lipgloss.ANSIColor(4)).
					Render("(R) ") + entry.Media.Title.UserPreferred
			}

			rows = append(rows, &MediaListEntry{
				Id: strconv.Itoa(entry.Media.Id),
				Title: entry.Media.Title.UserPreferred,
				Format: internal.MediaFormatFormatter(entry.Media.MediaFormat),
				Progress: progress,
			})
		}
	}

	fmt.Println(l.renderColumn(rows...))
	return nil
}