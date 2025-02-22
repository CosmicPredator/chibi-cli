package internal

import (
	"os"
	"path"
)

// creates a directory "chibi" in os specific config folder.
// in unix-like systems, the path is /home/user/.config/chibi.
// in windows, the path is %AppData%\chibi
func CreateConfigDir() {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")
	_, err := os.Stat(configDir)

	if err == nil {
		os.RemoveAll(configDir)
	}
	os.MkdirAll(configDir, 0755)
} 