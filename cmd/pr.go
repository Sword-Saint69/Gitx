package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var prCmd = &cobra.Command{
	Use:   "pr-summary",
	Short: "Generate a markdown PR summary for the current branch",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		summary, err := git.GeneratePRSummary(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" PULL REQUEST SUMMARY "))
		fmt.Println()
		
		fmt.Println(ui.BlockHeader(" Title "))
		fmt.Printf("  %s\n\n", ui.ValueStyle.Render(summary.Title))

		fmt.Println(ui.BlockHeader(" Changes "))
		if summary.Description == "" {
			fmt.Println("  No commits found between branch and base branch.")
		} else {
			fmt.Printf("  %s\n", ui.SubtleStyle.Render(summary.Description))
		}

		fmt.Println(ui.BlockHeader(" File Stats "))
		fmt.Printf("  %s\n", ui.SubtleStyle.Render(summary.Stats))

		fmt.Println()
		fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " PR Summary generated successfully."))
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}
