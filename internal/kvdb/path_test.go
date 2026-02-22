package kvdb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CosmicPredator/chibi/internal"
)

func TestResolveDataPathUsesCustomEnvPath(t *testing.T) {
	customPath := filepath.Join("tmp", "chibi-custom")
	legacyPath := filepath.Join("tmp", "chibi-legacy")

	t.Setenv(internal.DATA_PATH_ENV, customPath)
	t.Setenv(internal.LEGACY_PATH_ENV, legacyPath)

	path, err := resolveDataPath()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if path != filepath.Clean(customPath) {
		t.Fatalf("expected path %q, got %q", filepath.Clean(customPath), path)
	}
}

func TestResolveDataPathUsesLegacyEnvPath(t *testing.T) {
	legacyPath := filepath.Join("tmp", "chibi-legacy")

	t.Setenv(internal.DATA_PATH_ENV, "")
	t.Setenv(internal.LEGACY_PATH_ENV, legacyPath)

	path, err := resolveDataPath()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if path != filepath.Clean(legacyPath) {
		t.Fatalf("expected path %q, got %q", filepath.Clean(legacyPath), path)
	}
}

func TestResolveDataPathFallsBackToUserConfigDir(t *testing.T) {
	t.Setenv(internal.DATA_PATH_ENV, "")
	t.Setenv(internal.LEGACY_PATH_ENV, "")

	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("unable to get user config path: %v", err)
	}

	path, err := resolveDataPath()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedPath := filepath.Join(userConfigPath, internal.APP_DIR_NAME)
	if path != expectedPath {
		t.Fatalf("expected path %q, got %q", expectedPath, path)
	}
}
