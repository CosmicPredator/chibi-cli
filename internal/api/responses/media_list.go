package responses

type ListCollection struct {
	Lists []struct {
		Status  string `json:"status"`
		Entries []struct {
			Progress        int `json:"progress"`
			ProgressVolumes int `json:"progressVolumes"`
			Media           struct {
				Id    int `json:"id"`
				Title struct {
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
				Chapters *int `json:"chapters"`
				Volumes  *int `json:"volumes"`
				Episodes *int `json:"episodes"`
				MediaFormat string `json:"format"`
			} `json:"media"`
		} `json:"entries"`
	} `json:"lists"`
}

type MediaList struct {
	Data struct {
		AnimeListCollection ListCollection `json:"AnimeListCollection"`
		MangaListCollection ListCollection `json:"MangaListCollection"`
	} `json:"data"`
}
