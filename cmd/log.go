package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Interactive TUI log viewer",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetLog(cwd, 100)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewLogModel(entries), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
