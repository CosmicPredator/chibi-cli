package internal

import "github.com/CosmicPredator/chibi/types"

type MediaList struct {
	Data struct {
		MediaListCollection struct {
			Lists []struct {
				Status string `json:"status"`
				Entries []struct {
					Progress        int `json:"progress"`
					ProgressVolumes int `json:"progressVolumes"`
					Media           struct {
						Id    int `json:"id"`
						Title struct {
							UserPreferred string `json:"userPreferred"`
						} `json:"title"`
						Chapters int `json:"chapters"`
						Volumes  int `json:"volumes"`
						Episodes int `json:"episodes"`
					} `json:"media"`
				} `json:"entries"`
			} `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

func parseMediaStatus(status string) string {
	switch status {
	case "watching", "reading", "w", "r":
		return "CURRENT"
	case "planning", "p":
		return "PLANNING"
	case "completed", "c":
		return "COMPLETED"
	case "dropped", "d":
		return "DROPPED"
	case "paused", "ps":
		return "PAUSED"
	default:
		return "CURRENT"
	}
}

func (ml *MediaList) Get(mediaType string, status string) error {
	anilistClient := NewAnilistClient()
	tokenConfig := types.NewTokenConfig()
	err := tokenConfig.ReadFromJsonFile()

	if err != nil {
		return err
	}

	query :=
		`query($userId: Int, $type: MediaType, $status: [MediaListStatus]) {
        MediaListCollection(userId: $userId, type: $type, status_in: $status) {
            lists {
				status
                entries {
                    progress
                    progressVolumes
                    media {
                        id
                        title {
                            userPreferred
                        }
                        chapters
                        volumes
                        episodes
                    }
                }
            }
        }
    }`

	parsedStatus := parseMediaStatus(status)

	var parsedStatusSlice []string = make([]string, 0)

	if parsedStatus == "CURRENT" {
		parsedStatusSlice = append(parsedStatusSlice, parsedStatus, "REPEATING")
	} else {
		parsedStatusSlice = append(parsedStatusSlice, parsedStatus)
	}

	err = anilistClient.ExecuteGraqhQL(
		query,
		map[string]interface{}{
			"type":   mediaType,
			"userId": tokenConfig.UserId,
			"status": parsedStatusSlice,
		},
		&ml,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewMediaList() *MediaList {
	return &MediaList{}
}
