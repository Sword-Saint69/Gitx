package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var locCmd = &cobra.Command{
	Use:   "loc",
	Short: "Lines of code breakdown by language",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		stats, err := git.GetLOC(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(stats) == 0 {
			fmt.Println("No supported files found.")
			return
		}

		p := tea.NewProgram(ui.NewLOCModel(stats))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(locCmd)
}
