package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var contributorsCmd = &cobra.Command{
	Use:   "contributors",
	Short: "Ranked contributor list",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		contributors, err := git.GetContributors(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewContributorsModel(contributors))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(contributorsCmd)
}
