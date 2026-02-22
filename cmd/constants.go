package cmd

import (
	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

var ERROR_MESSAGE_TEMPLATE lipgloss.Style

var SUCCESS_MESSAGE_TEMPLATE lipgloss.Style

var OTHER_MESSAGE_TEMPLATE lipgloss.Style

func applyThemeToMessageTemplates() {
	palette := theme.Current()

	ERROR_MESSAGE_TEMPLATE = lipgloss.
		NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(palette.MessageError)).
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(palette.MessageError))

	SUCCESS_MESSAGE_TEMPLATE = lipgloss.
		NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(palette.MessageSuccess)).
		Foreground(lipgloss.Color(palette.MessageSuccess)).
		PaddingLeft(1).
		PaddingRight(1)

	OTHER_MESSAGE_TEMPLATE = lipgloss.
		NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(palette.MessageOther)).
		Foreground(lipgloss.Color(palette.MessageOther)).
		PaddingLeft(1).
		PaddingRight(1)
}

func init() {
	applyThemeToMessageTemplates()
}
