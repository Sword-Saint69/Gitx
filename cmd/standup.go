package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var standupCmd = &cobra.Command{
	Use:   "standup",
	Short: "Show your commits from the last 24 hours",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetStandup(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" STANDUP REPORT "))
		fmt.Println()

		for _, entry := range entries {
			fmt.Println(ui.BlockHeader(" Repo: " + entry.Repo + " "))
			if len(entry.Commits) == 0 {
				fmt.Println("  No commits found in the last 24 hours.")
			}
			for _, c := range entry.Commits {
				fmt.Printf("  %s %s\n", ui.SecondaryStyle.Render(ui.IconDot), ui.ValueStyle.Render(c))
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(standupCmd)
}
