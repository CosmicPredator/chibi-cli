package cmd

import (
	"github.com/CosmicPredator/chibi/internal/ui"
	"github.com/CosmicPredator/chibi/internal/viewmodel"
	"github.com/spf13/cobra"
)

// func getUserProfile() {
// 	CheckIfTokenExists()
// 	profile := internal.NewProfile()
// 	err := profile.Get()
// 	if err != nil {
// 		ErrorMessage(err.Error())
// 	}

// 	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6"))
// 	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8BE9FD"))

// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("ID"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Id)))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("Name"), valueStyle.Render(profile.Data.Viewer.Name))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Anime"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Anime.Count)))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Manga"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Manga.Count)))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Days Watched"), valueStyle.Render(fmt.Sprintf("%.2f", float32(profile.Data.Viewer.Statistics.Anime.MinutesWatched)/1440)))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("Total Chapters Read"), valueStyle.Render(strconv.Itoa(profile.Data.Viewer.Statistics.Manga.ChaptersRead)))
// 	fmt.Printf("%-20s : %s\n", keyStyle.Render("URL"), valueStyle.Render(profile.Data.Viewer.SiteUrl))
// }

func handleProfile(cmd *cobra.Command, args []string) {
	err := viewmodel.HandleProfile()
	if err != nil {
		ui.ErrorText(err)
	}
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get's your AniList profile (requires login)",
	Run: handleProfile,
}
