package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Full-text search across all commits",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		query := args[0]
		results, err := git.SearchHistory(cwd, query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewSearchModel(query, results))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
