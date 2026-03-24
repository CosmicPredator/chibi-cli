---
title: Search
lang: en-US
---

# The `search` Command Documentation

The `chibi search` command allows you to search for anime and manga by providing a query. You can specify the type of media and control the number of results returned.

## Usage
```bash
$ chibi search [query...] [flags]
```

## Flags

### `-h, --help`
Provides help information for the `chibi search` command. Use this flag to display details about how the command works and the options available.

#### Example:
```bash
$ chibi search --help
```

### `-p, --page int`
Specifies the number of results to be returned. By default, the command will return up to 10 results. You can increase or decrease this value to suit your needs.

#### Example:
```bash
$ chibi search "Attack on Titan" --page 5
```
This command will return up to 5 results for the query `"Attack on Titan"`.

### `-t, --type string`
Filters the search results by the type of media. Use the full name or shorthand notation:

- `anime` or `a`: Filters for anime.
- `manga` or `m`: Filters for manga.

If no type is specified, the default type is `anime`.

#### Example:
```bash
$ chibi search "One Piece" --type manga
```
This command will search for manga titled `"One Piece"`.

### Combined Example
You can combine the `--page` and `--type` flags to fine-tune your search results:

```bash
$ chibi search "Naruto" --type anime --page 3
```
This command will return up to 3 anime results for the query `"Naruto"`.

### `-j, --json`
Outputs search results in a structured JSON format.

The JSON output includes:
- `id`: The AniList Media ID
- `titles`: A nested object containing `romaji`, `english`, and `native` titles

#### Example:
```bash
$ chibi search "One Piece" --json
```

## Defaults
- The default `type` is `anime`.
- The default `page` is 10 results.

### Example of Default Usage
```bash
$ chibi search My Hero Academia
```
This command will search for up to 10 anime results for the query `"My Hero Academia"`.

## Output
The output will display the search results, including the title, score and id of the media.

Example Output:
```bash
┌─────────┬─────────────────────────────────────┬────────┐
│   ID    │                TITLE                │ SCORE  │
├─────────┼─────────────────────────────────────┼────────┤
│  21459  │          My Hero Academia           │  77.0  │
│  18289  │    My Hero Academia FINAL SEASON    │  0.00  │
│  10072  │    My Hero Academia: Two Heroes     │  74.0  │
│  10427  │      My Hero Academia Season 4      │  78.0  │
│  13963  │      My Hero Academia Season 6      │  82.0  │
│  11719  │      My Hero Academia Season 5      │  73.0  │
│  10016  │      My Hero Academia Season 3      │  79.0  │
│  16313  │      My Hero Academia Season 7      │  81.0  │
│  21856  │      My Hero Academia Season 2      │  79.0  │
│  18573  │    My Hero Academia: Vigilantes     │  00.0  │
└─────────┴─────────────────────────────────────┴────────┘
```