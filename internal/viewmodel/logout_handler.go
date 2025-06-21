package viewmodel

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal/credstore"
	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func to log user out from AniList
// this is achieved by just deleting the config/chibi folder (for now)

// TODO: Implement proper logout operations
func HandleLogout() error {
	err := credstore.DeleteCredentials()
	if err != nil {
		return err
	}
	fmt.Println(ui.SuccessText("Logged out successfully!"))
	return nil
}