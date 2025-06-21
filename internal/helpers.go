package internal

import (
	"os"
	"strings"
)

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
		return "?"
	}
}

func CanSupportKittyGP() bool {
    term := os.Getenv("TERM")
    termProgram := os.Getenv("TERM_PROGRAM")
    konsoleVersion := os.Getenv("KONSOLE_VERSION")

    if strings.Contains(strings.ToLower(term), "ghostty") {
        return true
    }
    if konsoleVersion != "" {
        return true
    }
    if strings.HasPrefix(strings.ToLower(term), "xterm-kitty") {
        return true
    }
    if strings.Contains(strings.ToLower(termProgram), "warp") {
        return true
    }
    if strings.Contains(strings.ToLower(term), "wayst") {
        return true
    }
    if strings.Contains(strings.ToLower(termProgram), "wezterm") ||
        strings.Contains(strings.ToLower(term), "wezterm") {
        return true
    }

    return false
}