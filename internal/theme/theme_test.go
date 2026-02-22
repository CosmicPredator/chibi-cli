package theme

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CosmicPredator/chibi/internal"
)

func resetDefaultTheme(t *testing.T) {
	t.Helper()
	if err := SetCurrent("default"); err != nil {
		t.Fatalf("unable to reset theme: %v", err)
	}
}

func writeThemeFile(t *testing.T, dataPath, name, content string) {
	t.Helper()

	themesPath := filepath.Join(dataPath, "themes")
	if err := os.MkdirAll(themesPath, 0o755); err != nil {
		t.Fatalf("unable to create themes directory: %v", err)
	}

	themePath := filepath.Join(themesPath, name)
	if err := os.WriteFile(themePath, []byte(content), 0o644); err != nil {
		t.Fatalf("unable to write theme file: %v", err)
	}
}

func TestSetCurrentCaseInsensitive(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	if err := SetCurrent("NoRd"); err != nil {
		t.Fatalf("SetCurrent returned error: %v", err)
	}

	if got, want := CurrentName(), "nord"; got != want {
		t.Fatalf("CurrentName() = %q, want %q", got, want)
	}
}

func TestSetCurrentUnknownTheme(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	if err := SetCurrent("unknown"); err == nil {
		t.Fatalf("SetCurrent should fail for unknown theme")
	}
}

func TestSaveAndLoadPersistedTheme(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	t.Setenv(internal.DATA_PATH_ENV, t.TempDir())
	t.Setenv(internal.THEME_ENV, "")

	if err := Save("sunset"); err != nil {
		t.Fatalf("Save returned error: %v", err)
	}
	if err := SetCurrent("default"); err != nil {
		t.Fatalf("SetCurrent returned error: %v", err)
	}

	if err := Load(); err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if got, want := CurrentName(), "sunset"; got != want {
		t.Fatalf("CurrentName() = %q, want %q", got, want)
	}
}

func TestEnvThemeOverridesPersistedTheme(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	t.Setenv(internal.DATA_PATH_ENV, t.TempDir())
	t.Setenv(internal.THEME_ENV, "nord")

	if err := Save("sunset"); err != nil {
		t.Fatalf("Save returned error: %v", err)
	}
	if err := SetCurrent("default"); err != nil {
		t.Fatalf("SetCurrent returned error: %v", err)
	}

	if err := Load(); err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if got, want := CurrentName(), "nord"; got != want {
		t.Fatalf("CurrentName() = %q, want %q", got, want)
	}
}

func TestLoadReadsThemeFilesFromDataPath(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	dataPath := t.TempDir()
	t.Setenv(internal.DATA_PATH_ENV, dataPath)
	t.Setenv(internal.THEME_ENV, "")

	writeThemeFile(t, dataPath, "ocean.toml", `
name = "ocean"
success_text = "#65D6CE"
error_text = "#FF6B6B"
highlight_text = "#6EC6FF"
spinner = "#4FC3F7"
prompt_title = "#4FC3F7"
prompt_default = "#FFD166"
key_text = "#B39DDB"
value_text = "#80CBC4"
table_header = "#4FC3F7"
table_id = "#6EC6FF"
table_format = "#65D6CE"
table_metric = "#FFD166"
table_repeating = "#5C6BC0"
message_error = "#FF6B6B"
message_success = "#65D6CE"
message_other = "#6EC6FF"
`)

	if err := Load(); err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if err := SetCurrent("ocean"); err != nil {
		t.Fatalf("SetCurrent returned error: %v", err)
	}

	if got, want := CurrentName(), "ocean"; got != want {
		t.Fatalf("CurrentName() = %q, want %q", got, want)
	}

	if got, want := Current().TableHeader, "#4FC3F7"; got != want {
		t.Fatalf("Current().TableHeader = %q, want %q", got, want)
	}
}

func TestLoadThemeNameFallsBackToFileName(t *testing.T) {
	resetDefaultTheme(t)
	defer resetDefaultTheme(t)

	dataPath := t.TempDir()
	t.Setenv(internal.DATA_PATH_ENV, dataPath)
	t.Setenv(internal.THEME_ENV, "")

	writeThemeFile(t, dataPath, "retro.toml", `
success_text = "#A6E22E"
error_text = "#F92672"
highlight_text = "#66D9EF"
spinner = "#FD971F"
prompt_title = "#FD971F"
prompt_default = "#E6DB74"
key_text = "#AE81FF"
value_text = "#66D9EF"
table_header = "#FD971F"
table_id = "#66D9EF"
table_format = "#A6E22E"
table_metric = "#E6DB74"
table_repeating = "#AE81FF"
message_error = "#F92672"
message_success = "#A6E22E"
message_other = "#66D9EF"
`)

	if err := Load(); err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if err := SetCurrent("retro"); err != nil {
		t.Fatalf("SetCurrent returned error: %v", err)
	}

	if got, want := CurrentName(), "retro"; got != want {
		t.Fatalf("CurrentName() = %q, want %q", got, want)
	}
}
