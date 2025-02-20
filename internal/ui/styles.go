package ui

import "github.com/charmbracelet/lipgloss"

// displays text in green with ✓ on the left
func SuccessText(msg string) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Render("✓ " + msg)
}

// displays text in red with ✘ on the left
func ErrorText(err error) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#CC0000")).
		Render("✘ Someting went wrong! Reason: ", err.Error())
}

// displays text in cyan foreground
func HighlightedText(msg string) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		PaddingLeft(0).
		PaddingRight(0).
		Render(msg)
}
