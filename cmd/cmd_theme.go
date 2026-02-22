package cmd

import (
	"fmt"
	"strings"

	"github.com/CosmicPredator/chibi/internal/theme"
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/spf13/cobra"
)

func handleTheme(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Current theme: %s\n", theme.CurrentName())
		fmt.Printf("Available themes: %s\n", strings.Join(theme.Available(), ", "))
		fmt.Println("Set a theme with: chibi theme <name>")
		if themesPath, err := theme.ThemesPath(); err == nil {
			fmt.Printf("Add your own theme: create a .toml file in %s\n", themesPath)
		} else {
			fmt.Println("Add your own theme: create a .toml file in DATA_DIR/themes")
		}
		return
	}

	if err := theme.Save(args[0]); err != nil {
		fmt.Println(ui.ErrorText(err))
		return
	}

	applyThemeToMessageTemplates()
	fmt.Println(ui.SuccessText(fmt.Sprintf("Theme set to %s", theme.CurrentName())))
}

var themeCmd = &cobra.Command{
	Use:   "theme [name]",
	Short: "Show or set the active color theme",
	Args:  cobra.MaximumNArgs(1),
	Run:   handleTheme,
}
