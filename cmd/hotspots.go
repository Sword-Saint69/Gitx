package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var hotspotsCmd = &cobra.Command{
	Use:   "hotspots",
	Short: "Combine churn + complexity to find risky code areas",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		hotspots, err := git.GetHotspots(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewHotspotsModel(hotspots))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(hotspotsCmd)
}
