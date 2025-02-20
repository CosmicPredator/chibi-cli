package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type ProfileUI struct {
	Id             int
	Name           string
	TotalAnime     int
	TotalManga     int
	MinutesWatched int
	ChaptersRead   int
	SiteUrl        string
}

// Simple Key Value pair used for ordered output
// since map produces unordered kv pairs while iterating
type KV struct {
	Key   string
	Value string
}

func (p *ProfileUI) Render() error {
	// convert minutes to days
	daysWatched := float32(p.MinutesWatched) / 1440

	// populating []KV with data
	var dataSlice = []KV{
		{"ID", strconv.Itoa(p.Id)},
		{"Name", p.Name},
		{"Total Anime", strconv.Itoa(p.TotalAnime)},
		{"Total Manga", strconv.Itoa(p.TotalManga)},
		{"Total Days Watched", fmt.Sprintf("%.2f", daysWatched)},
		{"Total Chapters Read", strconv.Itoa(p.ChaptersRead)},
		{"Site URL", p.SiteUrl},
	}

	// finding the max key length for padded output
	maxKeyLen := 0
	for _, kv := range dataSlice {
		if len(kv.Key) > maxKeyLen {
			maxKeyLen = len(kv.Key)
		}
	}

	// define styles for both key and value string
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD"))

	// iterating over dataSlice and printing the KV pairs
	// with appropriate padding
	for _, kv := range dataSlice {
		fmt.Printf(
			"%s : %s\n",
			keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
			valueStyle.Render(kv.Value),
		)
	}

	return nil
}
