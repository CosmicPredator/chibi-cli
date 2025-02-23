package internal

// very very very important constants

// Base URL is not gonna get changed for a while.
// So, keeping it as constant is not gonna hurt anyone.
const(
	API_ENDPOINT = "https://graphql.anilist.co"
	AUTH_URL = "https://anilist.co/api/v2/oauth/authorize?client_id=4593&response_type=token"
	BOLT_BUCKET_NAME = "ChibiConfig"
	BOLT_DB_NAME = "chibi_config.db"
)

type MediaType string 

const (
	ANIME MediaType = "ANIME"
	MANGA MediaType = "MANGA"
)