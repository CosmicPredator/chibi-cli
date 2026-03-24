---
title: Info
lang: en-US
---

# The `info` Command

The `chibi info` command provides detailed information about a specific anime or manga using its AniList ID.

## Usage
```bash
$ chibi info [id] [flags]
```

## Flags

### `-h, --help`
Provides help information for the `chibi info` command.

### `--json`
Outputs the media information in a structured JSON format. This includes IDs, titles (romaji, english, native), synonyms, scores, progress, and full descriptions.

#### Example:
```bash
$ chibi info 21 --json
```

## Output
By default, the command displays a formatted list of metadata, including:
- ID and MAL ID
- Titles (English, Romaji, Native)
- Format (TV, Movie, Manga, etc.)
- Average Score
- Episode/Chapter count
- Genres and Tags
- Studios
- Detailed Description
