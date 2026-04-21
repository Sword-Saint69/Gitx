package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var refactorCmd = &cobra.Command{
	Use:   "refactor [file]",
	Short: "Analyze code for refactoring opportunities",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		
		var files []string
		if len(args) > 0 {
			files = []string{args[0]}
		} else {
			// Scan all .go files in current dir recursively
			filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor") {
					files = append(files, path)
				}
				return nil
			})
		}

		fmt.Println(ui.HeaderStyle.Render(" REFACTORING OPPORTUNITIES "))
		fmt.Println()

		found := false
		for _, f := range files {
			relPath, _ := filepath.Rel(cwd, f)
			suggestions, _ := git.AnalyzeRefactor(f)
			if len(suggestions) > 0 {
				found = true
				fmt.Println(ui.BlockHeader(" " + relPath + " "))
				for _, s := range suggestions {
					priorityStyle := ui.SecondaryStyle
					if s.Priority == "HIGH" {
						priorityStyle = ui.WarningStyle
					}
					fmt.Printf("  %s %s: %s\n", priorityStyle.Render(s.Priority), ui.PrimaryStyle.Render(s.Context), ui.ValueStyle.Render(s.Reason))
				}
				fmt.Println()
			}
		}

		if !found {
			fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " Code looks clean! No obvious refactoring needed."))
		}
	},
}

func init() {
	rootCmd.AddCommand(refactorCmd)
}
