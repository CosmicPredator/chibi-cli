### ✨ What's New
- Chibi is now available in Snap Store for Linux users. To install chibi, run
    ```shell
    sudo snap install chibi
- Added Loading indicator to all API Requests.
- Added `chibi logout` command to log you out from AniList.
- Start Date is automatically added when `chibi add` command in invoked and `status` flag is set to `watching`.
- `chibi update` command now supports 2 new flags,
    - `-n "<note>"` for entry notes.
    - `-r <score>` for entry score.

### ✨ Other Changes
- Migrated from JSON file storage to Sqlite3 storage for configurations.
- Complete Architecture change for faster response.
- Added Code Comments in missing areas

### 🐛 Bug Fixes
- Fixed a bad `if` check when handling start date. (085fff237d04fde01e6b19d96078af2030ab1bbb)
- Fixed app exiting with weird error messages.
- Fixed "WinGet not detecting app version". This change will take effect when installing upcoming versions of chibi.

Thanks @mist8kengas for the contributions ☺️.
