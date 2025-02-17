package responses

type MediaSearch struct {
	Data struct {
		Page struct {
			Media []struct {
				Id    int `json:"id"`
				Title struct {
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
				AverageScore float64 `json:"averageScore"`
				MediaType    string  `json:"type"`
			} `json:"media"`
		} `json:"page"`
	} `json:"data"`
}
