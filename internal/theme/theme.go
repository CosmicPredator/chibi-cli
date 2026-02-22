package theme

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/kvdb"
)

const (
	themeKey         = "theme"
	themesDirName    = "themes"
	defaultThemeName = "default"
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

type paletteFile struct {
	Name string `toml:"name"`

	SuccessText   string `toml:"success_text"`
	ErrorText     string `toml:"error_text"`
	HighlightText string `toml:"highlight_text"`

	Spinner       string `toml:"spinner"`
	PromptTitle   string `toml:"prompt_title"`
	PromptDefault string `toml:"prompt_default"`

	KeyText   string `toml:"key_text"`
	ValueText string `toml:"value_text"`

	TableHeader    string `toml:"table_header"`
	TableID        string `toml:"table_id"`
	TableFormat    string `toml:"table_format"`
	TableMetric    string `toml:"table_metric"`
	TableRepeating string `toml:"table_repeating"`

	MessageError   string `toml:"message_error"`
	MessageSuccess string `toml:"message_success"`
	MessageOther   string `toml:"message_other"`
}

var defaultPalettes = map[string]Palette{
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
	mu       sync.RWMutex
	palettes = clonePalettes(defaultPalettes)
	current  = defaultThemeName
)

func normalizeThemeName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func clonePalettes(source map[string]Palette) map[string]Palette {
	out := make(map[string]Palette, len(source))
	for name, palette := range source {
		out[name] = palette
	}
	return out
}

func (file paletteFile) palette() (Palette, error) {
	palette := Palette{
		SuccessText:    strings.TrimSpace(file.SuccessText),
		ErrorText:      strings.TrimSpace(file.ErrorText),
		HighlightText:  strings.TrimSpace(file.HighlightText),
		Spinner:        strings.TrimSpace(file.Spinner),
		PromptTitle:    strings.TrimSpace(file.PromptTitle),
		PromptDefault:  strings.TrimSpace(file.PromptDefault),
		KeyText:        strings.TrimSpace(file.KeyText),
		ValueText:      strings.TrimSpace(file.ValueText),
		TableHeader:    strings.TrimSpace(file.TableHeader),
		TableID:        strings.TrimSpace(file.TableID),
		TableFormat:    strings.TrimSpace(file.TableFormat),
		TableMetric:    strings.TrimSpace(file.TableMetric),
		TableRepeating: strings.TrimSpace(file.TableRepeating),
		MessageError:   strings.TrimSpace(file.MessageError),
		MessageSuccess: strings.TrimSpace(file.MessageSuccess),
		MessageOther:   strings.TrimSpace(file.MessageOther),
	}

	required := []struct {
		name  string
		value string
	}{
		{name: "success_text", value: palette.SuccessText},
		{name: "error_text", value: palette.ErrorText},
		{name: "highlight_text", value: palette.HighlightText},
		{name: "spinner", value: palette.Spinner},
		{name: "prompt_title", value: palette.PromptTitle},
		{name: "prompt_default", value: palette.PromptDefault},
		{name: "key_text", value: palette.KeyText},
		{name: "value_text", value: palette.ValueText},
		{name: "table_header", value: palette.TableHeader},
		{name: "table_id", value: palette.TableID},
		{name: "table_format", value: palette.TableFormat},
		{name: "table_metric", value: palette.TableMetric},
		{name: "table_repeating", value: palette.TableRepeating},
		{name: "message_error", value: palette.MessageError},
		{name: "message_success", value: palette.MessageSuccess},
		{name: "message_other", value: palette.MessageOther},
	}

	for _, field := range required {
		if field.value == "" {
			return Palette{}, fmt.Errorf("missing %q", field.name)
		}
	}

	return palette, nil
}

func loadPalettesFromFiles() (map[string]Palette, error) {
	loadedPalettes := clonePalettes(defaultPalettes)

	themesPath, err := ThemesPath()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(themesPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return loadedPalettes, nil
		}
		return nil, fmt.Errorf("unable to read themes directory %q: %w", themesPath, err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.EqualFold(filepath.Ext(entry.Name()), ".toml") {
			continue
		}

		themePath := filepath.Join(themesPath, entry.Name())
		var file paletteFile
		if _, err := toml.DecodeFile(themePath, &file); err != nil {
			return nil, fmt.Errorf("unable to parse theme file %q: %w", themePath, err)
		}

		themeName := normalizeThemeName(file.Name)
		if themeName == "" {
			themeName = normalizeThemeName(strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())))
		}
		if themeName == "" {
			return nil, fmt.Errorf("theme file %q has an empty name", themePath)
		}

		palette, err := file.palette()
		if err != nil {
			return nil, fmt.Errorf("invalid theme %q in %q: %w", themeName, themePath, err)
		}

		loadedPalettes[themeName] = palette
	}

	return loadedPalettes, nil
}

func ThemesPath() (string, error) {
	dataPath, err := kvdb.DataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataPath, themesDirName), nil
}

func Current() Palette {
	mu.RLock()
	defer mu.RUnlock()

	if palette, ok := palettes[current]; ok {
		return palette
	}
	if defaultPalette, ok := palettes[defaultThemeName]; ok {
		return defaultPalette
	}
	return defaultPalettes[defaultThemeName]
}

func CurrentName() string {
	mu.RLock()
	defer mu.RUnlock()
	return current
}

func Available() []string {
	mu.RLock()
	defer mu.RUnlock()

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

	mu.RLock()
	_, ok := palettes[normalized]
	mu.RUnlock()
	if !ok {
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
	loadedPalettes, err := loadPalettesFromFiles()
	if err != nil {
		return err
	}

	mu.Lock()
	palettes = loadedPalettes
	current = defaultThemeName
	mu.Unlock()

	if envTheme := strings.TrimSpace(os.Getenv(internal.THEME_ENV)); envTheme != "" {
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
