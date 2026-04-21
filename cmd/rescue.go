package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var rescueCmd = &cobra.Command{
	Use:   "rescue",
	Short: "Find and recover 'lost' commits floating in the object database",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		
		fmt.Printf("%s Scanning for ghost commits (this may take a moment)...\n\n", ui.AccentStyle.Render("[!]"))
		
		ghosts, err := git.FindLostCommits(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" GHOST COMMIT RECOVERY "))
		fmt.Println()

		if len(ghosts) == 0 {
			fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " No dangling commits found. Your history is clean."))
			return
		}

		fmt.Printf("%s Found %d floating commits. These are NOT in any branch.\n\n", ui.WarningStyle.Render("[i]"), len(ghosts))

		for _, g := range ghosts {
			fmt.Printf("  %s %s\n", ui.PrimaryStyle.Render(g.Hash[:7]), ui.ValueStyle.Render(g.Subject))
		}

		fmt.Println()
		fmt.Printf("%s To recover one, run: %s\n", ui.AccentStyle.Render(">>>"), ui.SubtleStyle.Render("git merge <hash>"))
	},
}

func init() {
	rootCmd.AddCommand(rescueCmd)
}
