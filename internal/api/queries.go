package api

var searchMediaQuery = `query($searchQuery: String, $perPage: Int, $mediaType: MediaType) {
    Page(perPage: $perPage) {
        media(search: $searchQuery, type: $mediaType) {
            id
            title {
                userPreferred
            }
			type
            averageScore
        }
    }
}`

var mediaListQuery = `query($userId: Int, $type: MediaType, $status: [MediaListStatus]) {
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
                    type
                    chapters
                    volumes
                    episodes
                }
            }
        }
    }
}`

var viewerQuery = `query {
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