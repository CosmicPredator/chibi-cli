package ui

import "github.com/charmbracelet/lipgloss"

func SuccessText(msg string) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Render("✓ " + msg)
}

func ErrorText(err error) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#CC0000")).
		Render("✘ Someting went wrong! Reason: ", err.Error())
}

func HighlightedText(msg string) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		PaddingLeft(0).
		PaddingRight(0).
		Render(msg)
}