package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Auto-generate changelog from conventional commits",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		changes, err := git.GetChangelog(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" REPOSITORY CHANGELOG "))
		fmt.Println()

		sections := []string{"Features", "Fixes", "Other"}
		for _, section := range sections {
			items := changes[section]
			if len(items) == 0 {
				continue
			}

			fmt.Println(ui.BlockHeader(" " + section + " "))
			// Limit to 10 per section for the display
			for i, item := range items {
				if i > 5 {
					fmt.Printf("  %s %d more...\n", ui.SubtleStyle.Render("-"), len(items)-5)
					break
				}
				fmt.Printf("  %s %s\n", ui.SecondaryStyle.Render("-"), ui.ValueStyle.Render(item))
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)
}
