package ui

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// displays text in green with a check mark on the left
func SuccessText(msg string) string {
	palette := theme.Current()
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(palette.SuccessText)).
		Render("\u2713 " + msg)
}

// displays text in red with a cross on the left
func ErrorText(err error) string {
	palette := theme.Current()
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(palette.ErrorText)).
		Render("\u2717 Someting went wrong! Reason: ", err.Error())
}

// displays text in cyan foreground
func HighlightedText(msg string) string {
	palette := theme.Current()
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(palette.HighlightText)).
		PaddingLeft(0).
		PaddingRight(0).
		Render(msg)
}

// displays spinner while the supplied action func is getting executed
func ActionSpinner(title string, action func(context.Context) error) error {
	palette := theme.Current()
	done := make(chan bool)
	go func() {
		spinnerStyle := lipgloss.
			NewStyle().
			Foreground(lipgloss.Color(palette.Spinner))
		dots := []string{"\u280b", "\u2819", "\u2839", "\u2838", "\u283c", "\u2834", "\u2826", "\u2827", "\u2807", "\u280f"}
		styledDots := make([]string, 0, len(dots))
		for _, character := range dots {
			styledDots = append(styledDots, spinnerStyle.Render(character))
		}
		i := 0
		for {
			select {
			case <-done:
				fmt.Fprintf(os.Stderr, "\r\033[K")
				return
			default:
				fmt.Fprintf(os.Stderr, "\r%s %s", styledDots[i], title)
				time.Sleep(150 * time.Millisecond)
				i = (i + 1) % len(dots)
			}
		}
	}()
	err := action(context.TODO())
	done <- true
	return err
}

func PrettyInput(title, defaultVal string, validatorFunc func(s string) error) (string, error) {
	palette := theme.Current()
	if defaultVal == "" {
		defaultVal = "none"
	}
	formattedTitle := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(palette.PromptTitle)).
		Bold(true).
		Render(title)

	formattedDefaultVal := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color(palette.PromptDefault)).
		Bold(true).
		Render(fmt.Sprintf("Default: %s", defaultVal))
	prompt := fmt.Sprintf("%s (%s): ", formattedTitle, formattedDefaultVal)
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	err := validatorFunc(value)
	return value, err
}
