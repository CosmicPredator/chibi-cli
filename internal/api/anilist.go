package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/CosmicPredator/chibi/internal/api/responses"
	"github.com/CosmicPredator/chibi/internal/db"
)

// Helper function to parse query string and variable map
// and performs HTTP POST request to the AniList API.
// The response will be returned in []byte
func queryAnilist(query string, variables map[string]any) ([]byte, error) {
	payload := map[string]any{
		"query":     query,
		"variables": variables,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", internal.API_ENDPOINT, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	dbConn, err := db.NewDbConn(false)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	token, err := dbConn.GetConfig("auth_token")
	if err != nil {
		return nil, errors.New("not logged in. Please use \"chibi login\" to continue")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+*token)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Request Failed. Status Code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Helper function to perform media search.
// Requires rough title as string, number of results
// to be returned (perPage) and mediaType.
// mediaType should be either "ANIME" or "MANGA"
func SearchAnime(title string, perPage int, mediaType string) (*responses.MediaSearch, error) {
	payload := map[string]any{
		"title":     title,
		"perPage":   perPage,
		"mediaType": mediaType,
	}

	response, err := queryAnilist(searchMediaQuery, payload)
	if err != nil {
		return nil, err
	}

	var responseStruct responses.MediaSearch
	err = json.Unmarshal(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

// Helper function to perform media list.
// Required mediaType as string and mediaStatus as string
func GetMediaList(userId int, mediaType string, mediaStatusIn []string) (*responses.MediaList, error) {
	payload := map[string]any{
		"id": userId,
		"statusIn": mediaStatusIn,
	}

	response, err := queryAnilist(mediaListQuery, payload)
	if err != nil {
		return nil, err
	}

	var responseStruct responses.MediaList
	err = json.Unmarshal(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

// Herlper function to get details about the
// logged user
func GetUserProfile() (*responses.Profile, error) {
	response, err := queryAnilist(viewerQuery, nil)
	if err != nil {
		return nil, err
	}

	var responseStruct responses.Profile
	err = json.Unmarshal(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}
