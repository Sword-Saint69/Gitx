package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "Contribution activity heatmap",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		stats, err := git.GetTimeline(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewTimelineModel(stats))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(timelineCmd)
}
