---
title: Nextdate
lang: en-US
---

# The `nextdate` Command

The `chibi nextdate` command shows the next episode release date for an anime by AniList ID.

## Usage
```bash
$ chibi nextdate [id]
```

### Alias
You can also use:
```bash
$ chibi next [id]
```

## Parameters

### `[id]`
The AniList media ID of an anime entry.

Example:
```bash
$ chibi nextdate 16498
```

## Output
The command prints:
- ID and title
- Current airing status
- Next episode number (if available)
- Airing timestamp in local time and UTC
- Time remaining until the next episode

If AniList does not have a scheduled next episode yet, the date fields are shown as `Not available`.
