package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

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

// displays spinner while the supplied action
// func is getting executed
// func ActionSpinner(title string, action func(context.Context) error) error {
// 	return spinner.
// 		New().
// 		Title(title).
// 		ActionWithErr(action).
// 		Run()
// }

func ActionSpinner(title string, action func(context.Context) error) error {
    done := make(chan bool)
    go func() {
		spinnerStyle := lipgloss.
			NewStyle().
			Foreground(lipgloss.ANSIColor(5))
        dots := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		var styledDots []string = make([]string, 0) 
		for _, character := range dots {
			styledDots = append(styledDots, spinnerStyle.Render(character))
		}
        i := 0
        for {
            select {
            case <-done:
                fmt.Print("\r\033[K")
                return
            default:
                fmt.Printf("\r%s %s", styledDots[i], title)
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
	if defaultVal == "" {
		defaultVal = "none"
	}
	formattedTitle := lipgloss.
		NewStyle().
		Foreground(lipgloss.ANSIColor(5)).
		Bold(true).
		Render(title)
	
	formattedDefaultVal := lipgloss.
		NewStyle().
		Foreground(lipgloss.ANSIColor(1)).
		Bold(true).
		Render(fmt.Sprintf("Default: %s", defaultVal))
	prompt := fmt.Sprintf("%s (%s): ", formattedTitle, formattedDefaultVal)
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	err := validatorFunc(value)
	return value, err
}