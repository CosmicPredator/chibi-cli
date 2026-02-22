<div align=center>
    <img width=150
        src=https://i.imgur.com/qQ6Qs0v.png/>
</div>

<div align="center">
<h1>Chibi for AniList</h1>
<h2>A lightweight anime & manga tracker CLI app powered by AniList</h2>

<div align="center">
<a href="https://snapcraft.io/chibi">
    <img alt="Get it from the Snap Store" src=https://snapcraft.io/en/dark/install.svg />
</a>
<a href="#windows-via-winget">
    <img alt="Install for windows via winget" src=https://i.imgur.com/ENKa9Lv.png/>
</a>
</div>

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)
[![Build (Multiple Arch)](https://github.com/CosmicPredator/chibi-cli/actions/workflows/build.yml/badge.svg)](https://github.com/CosmicPredator/chibi-cli/actions/workflows/build.yml)
[![Release (Multiple Arch)](https://github.com/CosmicPredator/chibi-cli/actions/workflows/create_release.yml/badge.svg?branch=prod)](https://github.com/CosmicPredator/chibi-cli/actions/workflows/create_release.yml)

<!-- ![Made with VHS](https://vhs.charm.sh/vhs-4o1iqUYYSVr7QIO5m9Q5nX.gif) -->

</div>

## Features
- 😊 Easily manage your anime and manga lists without even opening your browser.
- 🪶 Lightweight and easy on your keyboard.
- 🌈 Colorful and structured outputs.
- 🗔 Supports most terminals and shells.
- 🔄 Changes are synced directly with AniList. No local saving BS.
- 🚀 Faster by design.

## Getting Started
This section provides the quickest way to get started with chibi-cli. For detailed tutorial, refer to [Documentation](#documentation)

### Optional Pre-Requisites
- Make sure you use any one of the [Nerd Fonts](https://www.nerdfonts.com/) for a proper output.
- Make sure your terminal supports 24 bit ANSI color profile.
- Most modern terminals like **Windows Terminal**, **Gnome Terminal**, **Kitty** or **Alacritty** etc., should work.

### Quick Installation

#### Linux (via snap store)
```bash
$ sudo snap install chibi
```

#### Windows (via winget)
```pwsh
PS C:\> winget install CosmicPredator.Chibi
```

#### Manual Installation
- Download the binary for your OS from the [releases](https://github.com/CosmicPredator/chibi-cli/releases) page.

- Open your favourite terminal in the directory where you downloaded chibi.

- Type in `./chibi` and you are in!

> [!NOTE]
> For windows, you may type `./chibi.exe` (in powershell).

> [!NOTE]
> For linux, make the binary executable by the following command,
>    ```sh
>    $ chmod +x ./chibi
>    ```

### Custom Data Directory
By default, Chibi stores local data in your OS config directory under `chibi`.

You can override this path with environment variables:
- `CHIBI_DATA_PATH` (preferred)
- `CHIBI_PATH` (legacy fallback)

Examples:
```bash
export CHIBI_DATA_PATH="$HOME/.local/share/chibi"
chibi login
```

```pwsh
$env:CHIBI_DATA_PATH = "D:\apps\chibi-data"
chibi login
```

### Themes
Chibi loads themes from `DATA_DIR/themes` using TOML files.

`DATA_DIR` is resolved the same way as local config storage:
- `CHIBI_DATA_PATH` (preferred)
- `CHIBI_PATH` (legacy fallback)
- OS config directory fallback

- Show active theme and available themes:
  ```bash
  chibi theme
  ```
- Persist a theme:
  ```bash
  chibi theme nord
  ```
- Temporarily override via environment variable:
  ```bash
  CHIBI_THEME=sunset chibi profile
  ```

Each theme file in `DATA_DIR/themes/*.toml` should include all color keys.
Example themes are available in the repository `themes/` directory.

## Documentation
You can check the docs [here](https://chibi-cli.pages.dev/).

## Contributing
Contributions are heartily welcomed...!

Please refer to the [pull request guide](https://github.com/CosmicPredator/chibi-cli/blob/develop/.github/PULL_REQUEST_TEMPLATE.md) 
before creating a pull request. 

## Special Thanks
This project is not possible without the following libraries,

- AniList - [Website](https://anilist.co)
- spf13/cobra - [GitHub](https://github.com/spf13/cobra)
- charmbracelet/huh - [GitHub](https://github.com/charmbracelet/huh)
- charmbracelet/lipgloss - [GitHub](https://github.com/charmbracelet/lipgloss)
