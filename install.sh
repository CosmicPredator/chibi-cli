#!/bin/bash

if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root or with sudo."
    exit 1
fi


PLATFORM=$(uname | tr "[:upper:]" "[:lower:]")
ARCHITECTURE=$(uname -m | sed -e "s/^x86_64$/amd64/" -e "s/^aarch64$/arm64/")
OS_ARCH=$(echo "$PLATFORM"_"$ARCHITECTURE")
REPO_URL="https://api.github.com/repos/CosmicPredator/chibi-cli/releases/latest"
BINARY_NAME="chibi"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"

install_binary() {

    echo "Platform: $PLATFORM Arch: $ARCHITECTURE"

    if [ "$PLATFORM" == "darwin" ] && ! command -v jq > /dev/null 2>&1; then
        tools_path=$(xcode-select -p 2>/dev/null)
        if [ -z "$tools_path" ]; then
            echo "Error: Either install jq or enable command line tools. To enable command line tools run  'xcode-select --install'"
            exit 1
        fi
    fi

    echo "Fetching the latest release information..."

    if command -v curl > /dev/null 2>&1; then
        RELEASE_INFO=$(curl -s "$REPO_URL")
    elif command -v wget > /dev/null 2>&1; then
        RELEASE_INFO=$(wget -q -O- "$REPO_URL")
    else
        echo "Error: Neither curl nor wget is installed."
        exit 1
    fi

    if command -v jq > /dev/null 2>&1; then
        DOWNLOAD_URL=$(echo "$RELEASE_INFO" | jq -r ".assets[] | select(.name | test(\"$OS_ARCH\")) | .browser_download_url")
    elif command -v python3 > /dev/null 2>&1; then
        DOWNLOAD_URL=$(echo "$RELEASE_INFO" | python3 -c "import sys, json; data = json.load(sys.stdin); print(next(asset['browser_download_url'] for asset in data['assets'] if '$OS_ARCH' in asset['name']))")
    else
        echo "Error: Neither jq or python3 is installed"
    fi

    if [ -z "$DOWNLOAD_URL" ]; then
        echo "Error: No binary found for $OS_ARCH in the latest release."
        exit 1
    fi

    echo "Downloading $BINARY_NAME from $DOWNLOAD_URL..."

    if command -v wget > /dev/null 2>&1; then
        wget -O "$BINARY_NAME" "$DOWNLOAD_URL"
    else
        curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"
    fi

    if [ ! -f "$BINARY_NAME" ]; then
        echo "Error: Failed to download $BINARY_NAME."
        exit 1
    fi

    echo "Installing $BINARY_NAME to $INSTALL_PATH..."

    mkdir -p "$(dirname "$INSTALL_PATH")"
    mv "$BINARY_NAME" "$INSTALL_PATH"

    chmod +x "$INSTALL_PATH"

    if [ -x "$INSTALL_PATH" ]; then
        echo "$BINARY_NAME installed successfully at $INSTALL_PATH"
    else
        echo "Error: Failed to install $BINARY_NAME."
        exit 1
    fi
}

uninstall_binary() {
    if [ -f "$INSTALL_PATH" ]; then
        echo "Uninstalling $BINARY_NAME from $INSTALL_PATH..."
        rm -f "$INSTALL_PATH"

        if [ ! -f "$INSTALL_PATH" ]; then
            echo "$BINARY_NAME uninstalled successfully."
        else
            echo "Error: Failed to uninstall $BINARY_NAME."
            exit 1
        fi
    else
        echo "Error: $BINARY_NAME is not installed."
        exit 1
    fi

    if [ "$PLATFORM" == "darwin" ]; then
        CONFIG_DIR="$HOME/Library/Application Support/chibi"
    else
        CONFIG_DIR="/home/$SUDO_USER/.config/chibi"
    fi

    if [ -d "$CONFIG_DIR" ]; then
        echo "Removing $CONFIG_DIR directory..."
        rm -rf "$CONFIG_DIR"
        
        if [ ! -d "$CONFIG_DIR" ]; then
            echo "$CONFIG_DIR removed successfully."
        else
            echo "Error: Failed to remove $CONFIG_DIR."
            exit 1
        fi
    else
        echo "No configuration directory found at $CONFIG_DIR."
    fi
}



if [ "$1" == "--uninstall" ]; then
    uninstall_binary
else
    install_binary
fi

exit 0
