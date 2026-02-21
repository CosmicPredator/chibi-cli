package viewmodel

import (
	"context"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/kvdb"
	"github.com/CosmicPredator/chibi/internal/ui"
)

func HandleLogin() error {
	loginUI := ui.LoginUI{}
	loginUI.SetLoginURL(internal.AUTH_URL)

	// display login URL
	err := loginUI.Render()
	if err != nil {
		return err
	}

	// write access token to db
	db, err := kvdb.Open()
	if err != nil {
		return err
	}
	
	err = db.Set(context.TODO(), "auth_token", []byte(loginUI.GetAuthToken()))
	if err != nil {
		return err
	}
	defer db.Close()

	// gets user profile details from api and saves
	// the username and ID to db
	var profile *responses.Profile
	err = ui.ActionSpinner("Logging In...", func(ctx context.Context) error {
		profile, err = api.GetUserProfile()
		return err
	})
	if err != nil {
		return err
	}

	err = db.Set(context.TODO(), "user_id", []byte(strconv.Itoa(profile.Data.Viewer.Id)))
	if err != nil {
		return err
	}

	err = db.Set(context.TODO(), "user_name", []byte(profile.Data.Viewer.Name))
	if err != nil {
		return err
	}

	// display success message
	fmt.Println(
		ui.SuccessText(fmt.Sprintf("Logged in as %s", profile.Data.Viewer.Name)),
	)

	return nil
}
