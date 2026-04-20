package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var churnCmd = &cobra.Command{
	Use:   "churn",
	Short: "Files with the most changes over time",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetChurn(cwd, 15)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewChurnModel(entries))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

var busFactorCmd = &cobra.Command{
	Use:   "bus-factor",
	Short: "Calculate bus factor - files with only 1 contributor",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		entries, err := git.GetBusFactor(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewBusFactorModel(entries))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(churnCmd)
	rootCmd.AddCommand(busFactorCmd)
}
