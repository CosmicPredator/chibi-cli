package internal

import (
	"os"
	"path"
)

// creates a directory "chibi" in os specific config folder.
// in unix-like systems, the path is /home/user/.config/chibi.
// in windows, the path is %AppData%\chibi
func CreateConfigDir() {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")
	_, err := os.Stat(configDir)

	if err == nil {
		os.RemoveAll(configDir)
	}
	os.MkdirAll(configDir, 0755)
} 

// maps "type" command line argument string to valid
// MediaType enum required by AniList API
func MediaTypeEnumMapper(mediaType string) string {
	switch mediaType {
	case "manga", "m":
		return "MANGA"
	default:
		return "ANIME"
	}
}

// maps "status" command line argument string to valid
// MediaType enum required by AniList API
func MediaStatusEnumMapper(mediaStatus string) string {
	switch mediaStatus {
	case "watching", "reading", "w", "r":
		return "CURRENT"
	case "planning", "p":
		return "PLANNING"
	case "completed", "c":
		return "COMPLETED"
	case "dropped", "d":
		return "DROPPED"
	case "paused", "ps":
		return "PAUSED"
	default:
		return "CURRENT"
	}
}

func MediaFormatFormatter(mediaFormat string) string {
	switch mediaFormat {
	case "TV":
		return "Tv"
	case "TV_SHORT":
		return "Tv Short"
	case "MOVIE":
		return "Movie"
	case "SPECIAL":
		return "Special"
	case "OVA":
		return "Ova"
	case "ONA":
		return "Ona"
	case "MUSIC":
		return "Music"
	case "MANGA":
		return "Manga"
	case "NOVEL":
		return "Novel"
	case "ONE_SHOT":
		return "One Shot"
	default:
		return ""
	}
}