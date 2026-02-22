package kvdb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CosmicPredator/chibi/internal"
)

func resolveDataPath() (string, error) {
	if customPath := strings.TrimSpace(os.Getenv(internal.DATA_PATH_ENV)); customPath != "" {
		return filepath.Clean(customPath), nil
	}
	if legacyPath := strings.TrimSpace(os.Getenv(internal.LEGACY_PATH_ENV)); legacyPath != "" {
		return filepath.Clean(legacyPath), nil
	}

	path, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user config path: %w", err)
	}
	return filepath.Join(path, internal.APP_DIR_NAME), nil
}
