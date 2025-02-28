package viewmodel

import (
	"fmt"
	"os"
	"path"

	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func to log user out from AniList
// this is achieved by just deleting the config/chibi folder (for now)

// TODO: Implement proper logout operations
func HandleLogout() error {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")

	_, err := os.Stat(configDir)
	if err != nil {
		ui.HighlightedText("Already logged out.")
		return nil
	}

	err = os.RemoveAll(configDir)
	if err != nil {
		return err
	}

	fmt.Println(ui.SuccessText("Logged out successfully!"))
	return nil
}