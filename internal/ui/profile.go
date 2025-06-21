package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/charmbracelet/lipgloss"
)

type ProfileUI struct {
	Id             int
	Name           string
	TotalAnime     int
	TotalManga     int
	MinutesWatched int
	ChaptersRead   int
	AvatarUrl      string
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

	var sb strings.Builder

	// iterating over dataSlice and adding the KV pairs to String Builder
	// with appropriate padding
	for _, kv := range dataSlice {
		if internal.CanSupportKittyGP() {
			sb.WriteString(
				fmt.Sprintf(
					"%*s%s : %s\n",
					20, "",
					keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
					valueStyle.Render(kv.Value),
				),
			)
		} else {
			sb.WriteString(
				fmt.Sprintf(
					"%s : %s\n",
					keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
					valueStyle.Render(kv.Value),
				),
			)
		}
	}

	// Display the output
	err := RenderWithImage(
		p.AvatarUrl,
		sb.String(),
		KGPParams{
			R: "7",
			C: "15",
		}, 8)
	return err
}
