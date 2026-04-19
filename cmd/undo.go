package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo recent git operations",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetReflog(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(entries) == 0 {
			fmt.Println("No reflog entries found.")
			return
		}

		p := tea.NewProgram(ui.NewUndoModel(entries, cwd), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
