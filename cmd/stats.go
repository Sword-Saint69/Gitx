package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Repo-wide analytics dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		stats, err := git.GetStats(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewStatsModel(stats))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
