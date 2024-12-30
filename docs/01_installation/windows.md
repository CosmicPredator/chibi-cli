---
title: Windows
lang: en-US
---

# Windows
This guide walks you through the steps to install and set up **chibi** on a Windows machine. Ensure you have administrative privileges to execute the necessary commands.

## Install using winget
You can install chibi via `winget` using the below command:
```shell
PS C:\> winget install CosmicPredator.Chibi
```

## Install via windows installer
You can also download the chibi installer from [here](https://github.com/CosmicPredator/chibi-cli/releases/latest) and run it.


## Manual Installation
1. **Download the Executable**

    Get the latest `.exe` file from the [Releases](https://github.com/CosmicPredator/chibi-cli/releases/latest) page.

2. **Rename the File**

    Rename the downloaded file (e.g., `chibi_1.0.0_windows_amd64.exe`) to `chibi.exe`.

3. **Move and Set Path**

    - Move the `chibi.exe` file to your preferred directory.  
    - Copy the directory path where you placed the file.

4. **Update Environment Variables** 

   - Open **"Edit Environment Variables for this System"** from the Windows search bar.  
   - Navigate to **Environment Variables > System Variables > Path**.  
   - Click **"Add"** and paste the copied directory path.  
   - Confirm by clicking **"OK"**, then close any open terminals.

## Verifying the Installation

1. Open a new terminal (Command Prompt or PowerShell).
2. Run the command `chibi version`.

   If the installation is successful, the version number will be displayed:
   ```shell
   PS C:\> chibi version
   # 1.0.0


## Next Steps
With Chibi installed, you can proceed to [CLI Guide](../02_cli_guide/index) or explore available commands by typing:
```shell
PS C:\> chibi --help
```