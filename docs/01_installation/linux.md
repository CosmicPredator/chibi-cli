---
title: Linux
lang: en-US
---

# Linux

## Install via Script
You can install chibi via a bash script for any distros. 

To install chibi,
```shell
$ curl -sSL https://raw.githubusercontent.com/CosmicPredator/chibi-cli/refs/heads/develop/install.sh | sudo bash
```

To uninstall chibi,
```shell
$ curl -sSL https://raw.githubusercontent.com/CosmicPredator/chibi-cli/refs/heads/develop/install.sh | sudo bash -s -- --uninstall
```
you can find the script [here](https://github.com/CosmicPredator/chibi-cli/blob/develop/install.sh).


## Ubuntu / Debian
For Ubuntu/Debian-based distros, you can install Chibi directly from the `.deb` package provided on the [Releases](https://github.com/CosmicPredator/chibi-cli/releases/latest) page. After downloading, you can either double-click on it to install via the software manager or use the following command in a terminal:
```shell
$ sudo dpkg -i chibi_1.0.0_amd64.deb
```

## Other Distros
For other distros, you can download the latest binary (`chibi_1.0.0_linux_amd64`) from the [Releases](https://github.com/CosmicPredator/chibi-cli/releases/latest). After downloading, open a terminal in the directory where you downloaded Chibi and follow these instructions,

1. Rename the binary to `chibi`,
    ```shell
    $ mv ./chibi_1.0.0_linux_amd64 ./chibi
    ```

2. Make the binary executable:
    ```shell
    $ chmod +x ./chibi
    ```

3. Move the binary to `/usr/local/bin` (requires sudo privileges):
    ```shell
    $ sudo mv ./chibi /usr/local/bin/chibi
    ```

4. Verify the installation by checking the version:
    ```shell
    $ chibi version
    # 1.0.0
    ```

::: warning
`1.0.0` in the name `chibi_1.0.0_linux_amd64` and `chibi_1.0.0_amd64.deb` refers to the version of Chibi you are downloading. In your case, the latest binary name will be different depending on the latest version of chibi. 
:::

## Next Steps
With Chibi installed, you can proceed to [CLI Guide](../02_cli_guide/index) or explore available commands by typing:
```shell
$ chibi --help
```
