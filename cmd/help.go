package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/ui"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show styled help for gitx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.HeaderStyle.Render(" COMMAND REFERENCE "))
		fmt.Println()

		// Categorize commands (Manually for better logic)
		categories := map[string][]string{
			"AI & Intel":   {"gen-msg", "chat", "pr-summary", "review", "refactor"},
			"Insights":     {"stats", "loc", "contributors", "timeline"},
			"Analysis":     {"churn", "bus-factor", "who", "hotspots"},
			"Visualization": {"weather", "pulse"},
			"Productivity": {"wip", "ignore", "changelog"},
			"Power Tools":  {"speed", "rescue", "leaderboard"},
			"Maintenance":  {"clean", "undo", "doctor", "secrets", "snapshot"},
		}

		// Sort keys to maintain stable order
		catOrder := []string{"AI & Intel", "Insights", "Analysis", "Visualization", "Productivity", "Power Tools", "Maintenance"}

		for _, cat := range catOrder {
			fmt.Printf("%s\n", ui.PrimaryStyle.Render(strings.ToUpper(cat)))
			for _, cName := range categories[cat] {
				c, _, _ := rootCmd.Find(strings.Split(cName, " "))
				if c != nil {
					fmt.Printf("  %-15s %s\n", ui.SecondaryStyle.Render(c.Name()), ui.SubtleStyle.Render(c.Short))
				}
			}
			fmt.Println()
		}

		fmt.Println(ui.SubtleStyle.Render("Usage: gitx <command> [options]"))
	},
}

func init() {
	rootCmd.SetHelpCommand(helpCmd)
}
