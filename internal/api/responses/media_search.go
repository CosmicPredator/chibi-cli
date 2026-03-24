package responses

type MediaSearchList struct {
	Id    int `json:"id"`
	Title struct {
		UserPreferred string `json:"userPreferred"`
		Romaji        string `json:"romaji"`
		English       string `json:"english"`
		Native        string `json:"native"`
	} `json:"title"`
	AverageScore *float64 `json:"averageScore"`
	MediaType    string   `json:"type"`
	MediaFormat  string   `json:"format"`
}

type MediaSearch struct {
	Data struct {
		Page struct {
			Media []MediaSearchList `json:"media"`
		} `json:"page"`
	} `json:"data"`
}
