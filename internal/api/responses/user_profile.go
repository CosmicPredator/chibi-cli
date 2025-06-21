package responses

type Profile struct {
	Data struct {
		Viewer struct {
			Name    string `json:"name"`
			SiteUrl string `json:"siteUrl"`
			Id      int    `json:"id"`
			Avatar  struct {
				Large string `json:"large"`
			} `json:"avatar"`
			Statistics struct {
				Anime struct {
					Count          int `json:"count"`
					MinutesWatched int `json:"minutesWatched"`
				} `json:"anime"`
				Manga struct {
					Count        int `json:"count"`
					ChaptersRead int `json:"chaptersRead"`
				} `json:"manga"`
			} `json:"statistics"`
		} `json:"Viewer"`
	} `json:"data"`
}
