package viewmodel

import (
	"fmt"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api"
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/types"
)

func HandleLogin() error {
	loginUI := ui.LoginUI{}
	loginUI.SetLoginURL(internal.AUTH_URL)
	
	err := loginUI.Render()
	if err != nil {
		return err
	}

	tokenConfig := types.NewTokenConfig()
	tokenConfig.AccessToken = loginUI.GetAuthToken()

	err = tokenConfig.FlushToJsonFile()
	if err != nil {
		return err
	}

	internal.ACCESS_TOKEN = loginUI.GetAuthToken()
	profile, err := api.GetUserProfile()
	if err != nil {
		return err
	}

	tokenConfig.UserId = profile.Data.Viewer.Id
	tokenConfig.Username = profile.Data.Viewer.Name

	err = tokenConfig.FlushToJsonFile()
	if err != nil {
		return err
	}

	fmt.Println(
		ui.SuccessText(fmt.Sprintf("Logged in as %s", profile.Data.Viewer.Name)),
	)

	return nil
}