package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleanup & Hygiene tools",
}

var cleanBranchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "Interactively delete branches",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		branches, err := git.GetBranches(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(branches) == 0 {
			fmt.Println("No branches found.")
			return
		}

		p := tea.NewProgram(ui.NewCleanupModel(branches, cwd), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	cleanCmd.AddCommand(cleanBranchesCmd)
	rootCmd.AddCommand(cleanCmd)
}
