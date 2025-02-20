package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

type LoginUI struct {
	loginURL string
	token    string
}

// setter only method on loginURL field
func (l *LoginUI) SetLoginURL(loginUrl string) {
	l.loginURL = loginUrl
}

// getter only method on token field
func (l LoginUI) GetAuthToken() string {
	return l.token
}

func (l *LoginUI) Render() error {
	// display login url
	var sb strings.Builder
	sb.WriteString("Open the below link in browser to login with anilist:")
	sb.WriteString("\n")
	sb.WriteString(HighlightedText(l.loginURL))
	sb.WriteString("\n\n")
	fmt.Print(sb.String())

	// display token entry form
	err := huh.NewText().
		Title("Paste your token here").
		Value(&l.token).
		CharLimit(2000).
		Validate(func(s string) error {
			if s == "" {
				return errors.New("please provide a valid token")
			}
			return nil
		}).
		Run()
	return err
}
