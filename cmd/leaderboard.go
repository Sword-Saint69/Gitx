package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var leaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "Show repository contributor leaderboards",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetLeaderboards(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" REPOSITORY LEADERBOARDS "))
		fmt.Println()

		// 1. Category: The Historian (Most Commits)
		sort.Slice(entries, func(i, j int) bool { return entries[i].Commits > entries[j].Commits })
		printCategory("THE HISTORIAN (Most Commits)", entries, func(e git.LeaderboardEntry) string {
			return fmt.Sprintf("%d commits", e.Commits)
		})

		// 2. Category: The Cleaner (Most Deletions)
		sort.Slice(entries, func(i, j int) bool { return entries[i].Deletions > entries[j].Deletions })
		printCategory("THE CLEANER (Most Deletions)", entries, func(e git.LeaderboardEntry) string {
			return fmt.Sprintf("%d deletions", e.Deletions)
		})

		// 3. Category: The Nomad (Most Files Touched)
		sort.Slice(entries, func(i, j int) bool { return len(entries[i].Files) > len(entries[j].Files) })
		printCategory("THE NOMAD (Most Files Moved)", entries, func(e git.LeaderboardEntry) string {
			return fmt.Sprintf("%d files", len(e.Files))
		})
	},
}

func printCategory(title string, entries []git.LeaderboardEntry, valFn func(git.LeaderboardEntry) string) {
	fmt.Println(ui.BlockHeader(" " + title + " "))
	for i, e := range entries {
		if i >= 3 { break }
		rank := fmt.Sprintf("#%d", i+1)
		fmt.Printf("  %s %-20s %s\n", ui.AccentStyle.Render(rank), ui.PrimaryStyle.Render(e.Author), ui.ValueStyle.Render(valFn(e)))
	}
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(leaderboardCmd)
}
