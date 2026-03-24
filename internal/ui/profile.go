package ui

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

type ProfileUI struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	TotalAnime     int    `json:"totalAnime"`
	TotalManga     int    `json:"totalManga"`
	MinutesWatched int    `json:"minutesWatched"`
	ChaptersRead   int    `json:"chaptersRead"`
	AvatarUrl      string `json:"avatarUrl"`
	SiteUrl        string `json:"siteUrl"`
	JSON           bool   `json:"-"`
}

func (p *ProfileUI) Render() error {
	if p.JSON {
		jsonData, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))
		return nil
	}
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
	palette := theme.Current()
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(palette.KeyText))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(palette.ValueText))

	var sb strings.Builder

	// iterating over dataSlice and adding the KV pairs to String Builder
	// with appropriate padding
	for _, kv := range dataSlice {
		if internal.CanSupportKittyGP() {
			fmt.Fprintf(&sb, "%*s%s : %s\n",
				20, "",
				keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
				valueStyle.Render(kv.Value))
		} else {
			fmt.Fprintf(&sb, "%s : %s\n",
				keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
				valueStyle.Render(kv.Value))
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
