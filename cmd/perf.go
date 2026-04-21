package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Profile and optimize git performance for the current repo",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		
		fmt.Printf("%s Profiling git performance...\n\n", ui.AccentStyle.Render("[*]"))
		
		report, err := git.ProfileGit(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" PERFORMANCE AUDIT "))
		fmt.Println()

		fmt.Printf("%-20s %s\n", ui.PrimaryStyle.Render("Git Status:"), ui.ValueStyle.Render(fmt.Sprintf("%v", report.StatusTime)))
		fmt.Printf("%-20s %s\n", ui.PrimaryStyle.Render("Log Walk:"), ui.ValueStyle.Render(fmt.Sprintf("%v", report.LogTime)))
		fmt.Printf("%-20s %s\n", ui.PrimaryStyle.Render("File Count:"), ui.ValueStyle.Render(fmt.Sprintf("%d", report.FileCount)))
		
		if len(report.Suggestions) > 0 {
			fmt.Println("\n" + ui.BlockHeader(" Optimization Suggestions "))
			for _, s := range report.Suggestions {
				fmt.Printf("  %s %s\n", ui.SecondaryStyle.Render(">"), s)
			}
		} else {
			fmt.Println("\n" + ui.SecondaryStyle.Render(ui.IconCheck+" Performance is optimal for this project size."))
		}
	},
}

func init() {
	rootCmd.AddCommand(speedCmd)
}
