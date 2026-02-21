import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Chibi CLI Docs",
  description: "Docs and Guides for Chibi for AniList",
  markdown: {
    theme: {
      light: "catppuccin-mocha",
      dark: "catppuccin-mocha"
    }
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting Started', link: '/00_getting_started/index' },
      { text: 'CLI Guide', link: '/02_cli_guide/index' }
    ],

    sidebar: [
      {
        text: "Getting Started",
        items: [
          { text: "Pre Requisites", link: "/00_getting_started/prereq" },
        ]
      },
      {
        text: "Installing Chibi",
        items: [
          { text: "Linux", link: "/01_installation/linux" },
          { text: "Windows", link: "/01_installation/windows" },
          { text: "Mac OS", link: "/01_installation/macos" }
        ]
      },
      {
        text: "CLI Guide",
        items: [
          { text: "help", link: "/02_cli_guide/help" },
          { text: "login", link: "/02_cli_guide/login" },
          { text: "profile", link: "/02_cli_guide/profile" },
          { text: "ls", link: "/02_cli_guide/ls" },
          { text: "search", link: "/02_cli_guide/search" },
          { text: "add", link: "/02_cli_guide/add" },
          { text: "update", link: "/02_cli_guide/update" }
        ]
      }
    ],

    footer: {
      message: 'Released under GNU GPL-3.0 license',
      copyright: 'Copyright © 2024-present Cosmic Predator'
    },

    search: {
      provider: 'local'
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/CosmicPredator/chibi-cli' }
    ]
  }
})
