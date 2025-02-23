package responses

type MediaUpdateResponse struct {
	Data struct {
		SaveMediaListEntry struct {
			Media struct {
				Id int `json:"id"`
				Title struct {
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
			} `json:"media"`
		} `json:"SaveMediaListEntry"`
	} `json:"data"`
}