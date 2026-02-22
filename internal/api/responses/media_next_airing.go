package responses

type MediaNextAiring struct {
	Data struct {
		Media *struct {
			ID     int    `json:"id"`
			Type   string `json:"type"`
			Status string `json:"status"`
			Title  struct {
				UserPreferred string `json:"userPreferred"`
			} `json:"title"`
			NextAiringEpisode *struct {
				Episode         int `json:"episode"`
				AiringAt        int `json:"airingAt"`
				TimeUntilAiring int `json:"timeUntilAiring"`
			} `json:"nextAiringEpisode"`
		} `json:"Media"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}
