package internal

// very very very important constants

// Base URL is not gonna get changed for a while.
// So, keeping it as constant is not gonna hurt anyone.
const (
	API_ENDPOINT    = "https://graphql.anilist.co"
	AUTH_URL        = "https://anilist.co/api/v2/oauth/authorize?client_id=4593&response_type=token"
	APP_DIR_NAME    = "chibi"
	DB_PATH         = "chibi_config.db"
	DATA_PATH_ENV   = "CHIBI_DATA_PATH"
	LEGACY_PATH_ENV = "CHIBI_PATH"
)

type MediaType string

const (
	ANIME MediaType = "ANIME"
	MANGA MediaType = "MANGA"
)
