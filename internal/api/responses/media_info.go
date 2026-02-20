package responses

type MediaInfo struct {
	Data struct {
		Media struct {
			ID    int `json:"id"`
			IDMal int `json:"idMal"`
			Title struct {
				English string `json:"english"`
				Romaji  string `json:"romaji"`
				Native  string `json:"native"`
			} `json:"title"`
			MeanScore  int `json:"meanScore"`
			CoverImage struct {
				ExtraLarge string `json:"extraLarge"`
			} `json:"coverImage"`
			Genres []string `json:"genres"`
			Tags   []struct {
				Name string `json:"name"`
			} `json:"tags"`
			Studios struct {
				Nodes []struct {
					Name string `json:"name"`
				} `json:"nodes"`
			} `json:"studios"`
			Description string `json:"description"`
			Format      string `json:"format"`
			Episodes    int    `json:"episodes"`
			Duration    int    `json:"duration"`
			Chapters    int    `json:"chapters"`
			Volumes     int    `json:"volumes"`
			Type        string `json:"type"`
		} `json:"Media"`
	} `json:"data"`
}