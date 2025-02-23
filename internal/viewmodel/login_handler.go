package viewmodel

import (
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/db"
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

	dbConn, err := db.NewDbConn(true)
	if err != nil {
		return err
	}
	defer dbConn.Close()

	// write access token to db
	err = dbConn.SetConfig("auth_token", loginUI.GetAuthToken())
	if err != nil {
		return err
	}

	// gets user profile details from api and saves 
	// the username and ID to db
	profile, err := api.GetUserProfile()
	if err != nil {
		return err
	}

	err = dbConn.SetConfig("user_id", strconv.Itoa(profile.Data.Viewer.Id))
	if err != nil {
		return err
	}

	err = dbConn.SetConfig("user_name", profile.Data.Viewer.Name)
	if err != nil {
		return err
	}

	// display success message
	fmt.Println(
		ui.SuccessText(fmt.Sprintf("Logged in as %s", profile.Data.Viewer.Name)),
	)

	return nil
}
