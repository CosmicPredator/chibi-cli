package theme

import (
	"testing"

	"github.com/CosmicPredator/chibi/internal"
)

func resetDefaultTheme(t *testing.T) {
	t.Helper()
	if err := SetCurrent("default"); err != nil {
		t.Fatalf("unable to reset theme: %v", err)
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
	t.Setenv(ThemeEnvName, "")

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
	t.Setenv(ThemeEnvName, "nord")

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
