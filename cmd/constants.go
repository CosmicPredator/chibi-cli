package cmd

import "github.com/charmbracelet/lipgloss"

var ERROR_MESSAGE_TEMPLATE = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#FF0000")).
	PaddingLeft(1).
	PaddingRight(1).
	Foreground(lipgloss.Color("#FF0000"))

var SUCCESS_MESSAGE_TEMPLATE = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00FF00")).
	Foreground(lipgloss.Color("#00FF00")).
	PaddingLeft(1).
	PaddingRight(1)

var OTHER_MESSAGE_TEMPLATE = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#00FFFF")).
	Foreground(lipgloss.Color("#00FFFF")).
	PaddingLeft(1).
	PaddingRight(1)
