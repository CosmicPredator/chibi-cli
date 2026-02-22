package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

type MediaListUI struct {
	MediaType string
	MediaList *responses.MediaList
}

type MediaListEntry struct {
	Id          string
	Title       string
	Format      string
	Progress    string
	NextEpEpoch int64
}

func (l *MediaListUI) renderColumn(mediaType internal.MediaType, entries ...*MediaListEntry) string {
	palette := theme.Current()
	col := func(w int) lipgloss.Style {
		return lipgloss.NewStyle().Width(w).MarginRight(2).Align(lipgloss.Right)
	}

	epFormatCol := func() string {
		if mediaType == internal.ANIME {
			return "NEXT EP IN"
		} else {
			return "FORMAT"
		}
	}()

	epFormatValue := func(entry *MediaListEntry) string {
		if mediaType == internal.ANIME {
			return internal.FormatAiringTs(entry.NextEpEpoch)
		} else {
			return entry.Format
		}
	}

	styles := []lipgloss.Style{
		col(7),
		col(10),
		col(8),
		col(0),
	}

	headerStyle := func(style lipgloss.Style) lipgloss.Style {
		return style.MarginBottom(1).Underline(true).Bold(true).Foreground(lipgloss.Color(palette.TableHeader))
	}

	var sb strings.Builder
	header := []string{
		headerStyle(styles[0]).Render("ID"),
		headerStyle(styles[1]).Render(epFormatCol),
		headerStyle(styles[2]).Render("PROGRESS"),
		headerStyle(styles[3]).Render("TITLE"),
	}

	sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, header...) + "\n")
	for _, entry := range entries {
		row := []string{
			styles[0].Foreground(lipgloss.Color(palette.TableID)).Render(entry.Id),
			styles[1].Foreground(lipgloss.Color(palette.TableFormat)).Render(epFormatValue(entry)),
			styles[2].Foreground(lipgloss.Color(palette.TableMetric)).Render(entry.Progress),
			styles[3].Render(entry.Title),
		}
		sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...) + "\n")
	}

	return sb.String()
}

func (l *MediaListUI) Render() error {
	palette := theme.Current()
	var rows []*MediaListEntry = make([]*MediaListEntry, 0)
	var selectedList responses.ListCollection

	if internal.MediaType(l.MediaType) == internal.ANIME {
		selectedList = l.MediaList.Data.AnimeListCollection
	} else {
		selectedList = l.MediaList.Data.MangaListCollection
	}

	for _, list := range selectedList.Lists {
		if list.Status != "CURRENT" && list.Status != "REPEATING" {
			continue
		}

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
					Foreground(lipgloss.Color(palette.TableRepeating)).
					Render("(R) ") + entry.Media.Title.UserPreferred
			}

			rows = append(rows, &MediaListEntry{
				Id:          strconv.Itoa(entry.Media.Id),
				Title:       entry.Media.Title.UserPreferred,
				Format:      internal.MediaFormatFormatter(entry.Media.MediaFormat),
				Progress:    progress,
				NextEpEpoch: entry.Media.NextAiringEpisode.AiringAt,
			})
		}
	}

	fmt.Println(l.renderColumn(internal.MediaType(l.MediaType), rows...))
	return nil
}
