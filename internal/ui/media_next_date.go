package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type MediaNextDateUI struct {
	ID              int
	Title           string
	Status          string
	NextEpisode     *int
	AiringAtUnix    *int
	TimeUntilAiring *int
}

func formatMediaStatus(status string) string {
	if status == "" {
		return "Unknown"
	}

	rawParts := strings.Split(strings.ToLower(status), "_")
	parts := make([]string, 0, len(rawParts))
	for _, part := range rawParts {
		if part == "" {
			continue
		}
		parts = append(parts, strings.ToUpper(part[:1])+part[1:])
	}
	if len(parts) == 0 {
		return "Unknown"
	}
	return strings.Join(parts, " ")
}

func formatCountdown(seconds int) string {
	if seconds <= 0 {
		return "Now"
	}

	d := time.Duration(seconds) * time.Second
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	if len(parts) == 0 {
		return "Less than a minute"
	}
	return strings.Join(parts, " ")
}

func (m *MediaNextDateUI) Render() error {
	nextEpisode := "Not available"
	airingAtLocal := "Not available"
	airingAtUTC := "Not available"
	countdown := "Not available"

	if m.NextEpisode != nil {
		nextEpisode = strconv.Itoa(*m.NextEpisode)
	}
	if m.AiringAtUnix != nil {
		airingAt := time.Unix(int64(*m.AiringAtUnix), 0)
		airingAtLocal = airingAt.Local().Format(time.RFC1123)
		airingAtUTC = airingAt.UTC().Format(time.RFC1123)
	}
	if m.TimeUntilAiring != nil {
		countdown = formatCountdown(*m.TimeUntilAiring)
	}

	dataSlice := []KV{
		{"ID", strconv.Itoa(m.ID)},
		{"Title", m.Title},
		{"Status", formatMediaStatus(m.Status)},
		{"Next Episode", nextEpisode},
		{"Airs At (Local)", airingAtLocal},
		{"Airs At (UTC)", airingAtUTC},
		{"Countdown", countdown},
	}

	maxKeyLen := 0
	for _, kv := range dataSlice {
		if len(kv.Key) > maxKeyLen {
			maxKeyLen = len(kv.Key)
		}
	}

	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD"))

	var sb strings.Builder
	for _, kv := range dataSlice {
		fmt.Fprintf(
			&sb,
			"%s : %s\n",
			keyStyle.MarginRight(maxKeyLen-len(kv.Key)).Render(kv.Key),
			valueStyle.Render(kv.Value),
		)
	}

	fmt.Println(sb.String())
	return nil
}
