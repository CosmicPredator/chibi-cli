package api

const searchMediaQuery = `query($searchQuery: String, $perPage: Int, $mediaType: MediaType) {
    Page(perPage: $perPage) {
        media(search: $searchQuery, type: $mediaType) {
            id
            title {
                userPreferred
            }
			type
            averageScore
			format
        }
    }
}`

const mediaListQuery = `query ($id: Int, $statusIn: [MediaListStatus]) {
	AnimeListCollection: MediaListCollection(userId: $id, type: ANIME, status_in:$statusIn){
		lists {
			status
			entries {
				progress
				media {
					id
					title {
						userPreferred
					}
					episodes
					chapters
					format
				}
			}
		}
	}
	MangaListCollection: MediaListCollection(userId: $id, type: MANGA, status_in:$statusIn){
		lists {
			status
			entries {
				progress
				media {
					id
					title {
						userPreferred
					}
					episodes
					chapters
					format
				}
			}
		}
	}
}`

const viewerQuery = `query {
    Viewer {
        id
        name
        statistics {
            anime {
                count
                minutesWatched
            }
            manga {
                count
                chaptersRead
            }
        }
        siteUrl
    }
}`
