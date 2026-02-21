package viewmodel

import (
	"context"
	"fmt"
	
	"github.com/CosmicPredator/chibi/internal/kvdb"
	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func to log user out from AniList
// this is achieved by just deleting the config/chibi folder (for now)

// TODO: Implement proper logout operations
func HandleLogout() error {
	db, err := kvdb.Open()
	if err != nil {
		return err
	}
	defer db.Close()
	
	err = db.Delete(context.TODO(), "auth_token")
	err = db.Delete(context.TODO(), "user_id")
	err = db.Delete(context.TODO(), "user_name")
	fmt.Println(ui.SuccessText("Logged out successfully!"))
	return nil
}