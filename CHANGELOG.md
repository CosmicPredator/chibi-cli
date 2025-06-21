### ✨ What's New

#### The `chibi profile` command now supports **Rendering Pixel Perfect Avatar** in the terminal. See [supported](https://sw.kovidgoyal.net/kitty/graphics-protocol/#:~:text=Other%20terminals%20that%20have%20implemented%20the%20graphics%20protocol%3A) terminals.
<img src="https://github.com/user-attachments/assets/0eb56d2b-1b13-4f54-b71f-f953b012bc93" width="500">

--- 

#### New colorful help text (thanks to [charmbracelet/fang](https://github.com/charmbracelet/fang))
<img src="https://github.com/user-attachments/assets/6884eed5-97f7-45a5-b639-1b2c6f8dc860" width="500"/>

---

#### The Table Layout for `chibi ls` and `chibi search` commands has been redesigned to be a compact [`eza`](https://github.com/eza-community/eza) like layout.
<img src="https://github.com/user-attachments/assets/baba0b6c-d2f8-4063-805a-6972883827ec" width="500"/>


### ✨ Other Changes
- Removed SQLite Dependencies.
- Used keyring store instead of SQLite DB.
- Removed [`charmbracelet/huh`](https://github.com/charmbracelet/huh) dependencies (forms, spinners) and implemented those manually.

> [!NOTE]  
> By reducing those dependencies, I was able to save approximately 3MB 😁.

### 🐛 Bug Fixes
- Fixed an early token check in `chibi ls` command, which restricts from calling the AniList API.
