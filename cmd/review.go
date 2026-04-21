package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Perform a pre-commit heuristic code review",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		issues, err := git.PerformReview(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" CODE REVIEW REPORT "))
		fmt.Println()

		if len(issues) == 0 {
			fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " No issues found! Your diff looks clean."))
			return
		}

		for _, issue := range issues {
			severityStyle := ui.SecondaryStyle
			if issue.Severity == "HIGH" {
				severityStyle = ui.ErrorStyle
			} else if issue.Severity == "MEDIUM" {
				severityStyle = ui.WarningStyle
			}

			fmt.Printf("%s %s\n", severityStyle.Render("["+issue.Severity+"]"), ui.PrimaryStyle.Render(issue.File))
			fmt.Printf("  %s %s\n\n", ui.SubtleStyle.Render("Line:"), ui.ValueStyle.Render(fmt.Sprintf("%d - %s", issue.Line, issue.Message)))
		}

		fmt.Printf("\n%s Foundations laid for LLM-based deep review.\n", ui.AccentStyle.Render("[i]"))
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}
