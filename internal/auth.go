package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CosmicPredator/chibi/types"
)

type AuthRequest struct{}

func (a AuthRequest) GetAuthURL() string {
	return AUTH_URL
}

func (a AuthRequest) Login(authCode string) error {
	if err := saveAccessToken(authCode); err != nil {
		return err
	}

	return nil
}

func saveAccessToken(authCode string) error {
	client := NewAnilistClient()
	tokenConfig := types.NewTokenConfig()
	tokenConfig.AccessToken = authCode

	query :=
		`query {
        Viewer {
            id
            name
        }
    }`
	requestData, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", client.ApiUrl, bytes.NewBuffer(requestData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authCode))

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData struct {
		Data struct {
			Viewer struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"Viewer"`
		} `json:"data"`
	}

	err = json.Unmarshal(responseBytes, &responseData)
	if err != nil {
		return err
	}

	tokenConfig.Username = responseData.Data.Viewer.Name
	tokenConfig.UserId = responseData.Data.Viewer.Id
	tokenConfig.FlushToJsonFile()

	return nil
}

func NewAuthRequest() *AuthRequest {
	return &AuthRequest{}
}
