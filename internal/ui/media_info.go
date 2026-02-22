package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type MediaInfoUI struct {
	Id             int
	MalId          int
	EnglishTitle   string
	RomajiTitle    string
	NativeTitle    string
	Score          int
	IsAnime        bool
	ChapterEpisode int
	VolumeDuration int
	CoverImage     string
	Genres         string
	Tags           string
	Studios        string
	Description    string
	Format         string
}

func (m *MediaInfoUI) Render() error {
	var dataSlice = []KV{
		{"ID", strconv.Itoa(m.Id)},
		{"MAL ID", strconv.Itoa(m.MalId)},
		{"English Title", m.EnglishTitle},
		{"Romaji Title", m.RomajiTitle},
		{"Native Title", m.NativeTitle},
		{"Format", m.Format},
		{"Score", strconv.Itoa(m.Score)},
		{"Chapters/Episodes", strconv.Itoa(m.ChapterEpisode)},
		{"Volumes/Duration", strconv.Itoa(m.VolumeDuration)},
		{"Genres", m.Genres},
		{"Tags", m.Tags},
		{"Studios", m.Studios},
		{"Description", m.Description},
	}

	maxKeyLen := 0
	for _, kv := range dataSlice {
		if len(kv.Key) > maxKeyLen {
			maxKeyLen = len(kv.Key)
		}
	}

	sep := " : "
	valueIndent := strings.Repeat(" ", maxKeyLen+len(sep))

	palette := theme.Current()
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(palette.KeyText))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(palette.ValueText))

	var sb strings.Builder
	for _, kv := range dataSlice {
		key := keyStyle.Width(maxKeyLen).Render(kv.Key)
		val := valueStyle.Render(kv.Value)
		wrapWidth := 100 - (maxKeyLen + len(sep))
		wrapped := wordwrap.String(val, wrapWidth)

		lines := strings.Split(wrapped, "\n")

		fmt.Fprintf(&sb, "%s%s%s\n",
			key,
			sep,
			lines[0],
		)

		for _, line := range lines[1:] {
			fmt.Fprintf(&sb, "%s%s\n", valueIndent, line)
		}

		if kv.Key == "Volumes/Duration" || kv.Key == "Studios" {
			sb.WriteString("\n")
		}
	}

	// Display the output
	fmt.Println(sb.String())
	return nil
}
