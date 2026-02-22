package theme

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/CosmicPredator/chibi/internal/kvdb"
)

const (
	ThemeEnvName = "CHIBI_THEME"
	themeKey     = "theme"
)

type Palette struct {
	SuccessText   string
	ErrorText     string
	HighlightText string

	Spinner       string
	PromptTitle   string
	PromptDefault string

	KeyText   string
	ValueText string

	TableHeader    string
	TableID        string
	TableFormat    string
	TableMetric    string
	TableRepeating string

	MessageError   string
	MessageSuccess string
	MessageOther   string
}

var palettes = map[string]Palette{
	"default": {
		SuccessText:    "#00FF00",
		ErrorText:      "#CC0000",
		HighlightText:  "#00FFFF",
		Spinner:        "5",
		PromptTitle:    "5",
		PromptDefault:  "1",
		KeyText:        "#FF79C6",
		ValueText:      "#8BE9FD",
		TableHeader:    "5",
		TableID:        "6",
		TableFormat:    "2",
		TableMetric:    "3",
		TableRepeating: "4",
		MessageError:   "#FF0000",
		MessageSuccess: "#00FF00",
		MessageOther:   "#00FFFF",
	},
	"nord": {
		SuccessText:    "#A3BE8C",
		ErrorText:      "#BF616A",
		HighlightText:  "#88C0D0",
		Spinner:        "#81A1C1",
		PromptTitle:    "#81A1C1",
		PromptDefault:  "#D08770",
		KeyText:        "#B48EAD",
		ValueText:      "#8FBCBB",
		TableHeader:    "#81A1C1",
		TableID:        "#88C0D0",
		TableFormat:    "#A3BE8C",
		TableMetric:    "#EBCB8B",
		TableRepeating: "#5E81AC",
		MessageError:   "#BF616A",
		MessageSuccess: "#A3BE8C",
		MessageOther:   "#88C0D0",
	},
	"sunset": {
		SuccessText:    "#43A047",
		ErrorText:      "#E53935",
		HighlightText:  "#00ACC1",
		Spinner:        "#FB8C00",
		PromptTitle:    "#FB8C00",
		PromptDefault:  "#F4511E",
		KeyText:        "#8E24AA",
		ValueText:      "#039BE5",
		TableHeader:    "#FB8C00",
		TableID:        "#1E88E5",
		TableFormat:    "#43A047",
		TableMetric:    "#FDD835",
		TableRepeating: "#3949AB",
		MessageError:   "#E53935",
		MessageSuccess: "#43A047",
		MessageOther:   "#00ACC1",
	},
}

var (
	mu      sync.RWMutex
	current = "default"
)

func normalizeThemeName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func Current() Palette {
	mu.RLock()
	defer mu.RUnlock()
	return palettes[current]
}

func CurrentName() string {
	mu.RLock()
	defer mu.RUnlock()
	return current
}

func Available() []string {
	out := make([]string, 0, len(palettes))
	for name := range palettes {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

func SetCurrent(name string) error {
	normalized := normalizeThemeName(name)
	if normalized == "" {
		return errors.New("theme name cannot be empty")
	}
	if _, ok := palettes[normalized]; !ok {
		return fmt.Errorf(
			"unknown theme %q (available: %s)",
			name,
			strings.Join(Available(), ", "),
		)
	}

	mu.Lock()
	current = normalized
	mu.Unlock()
	return nil
}

func Save(name string) error {
	if err := SetCurrent(name); err != nil {
		return err
	}

	db, err := kvdb.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Set(context.TODO(), themeKey, []byte(CurrentName()))
}

func Load() error {
	if envTheme := strings.TrimSpace(os.Getenv(ThemeEnvName)); envTheme != "" {
		return SetCurrent(envTheme)
	}

	db, err := kvdb.Open()
	if err != nil {
		return nil
	}
	defer db.Close()

	rawTheme, err := db.Get(context.TODO(), themeKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return nil
	}

	name := strings.TrimSpace(string(rawTheme))
	if name == "" {
		return nil
	}

	// Ignore stale values from old/invalid config and fallback to default.
	if err := SetCurrent(name); err != nil {
		return nil
	}
	return nil
}
