package viewmodel

import (
	"context"
	"fmt"
	"strconv"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/credstore"
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
	err = credstore.SetCredential("auth_token", loginUI.GetAuthToken())
	if err != nil {
		return err
	}

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

	err = credstore.SetCredential("user_id", strconv.Itoa(profile.Data.Viewer.Id))
	if err != nil {
		return err
	}

	err = credstore.SetCredential("user_name", profile.Data.Viewer.Name)
	if err != nil {
		return err
	}

	// display success message
	fmt.Println(
		ui.SuccessText(fmt.Sprintf("Logged in as %s", profile.Data.Viewer.Name)),
	)

	return nil
}
