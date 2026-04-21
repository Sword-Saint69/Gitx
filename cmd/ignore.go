package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var appendFlag bool

var ignoreCmd = &cobra.Command{
	Use:   "ignore [language]",
	Short: "Generate .gitignore from templates (Go, Node, Python, Rust)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := git.GenerateIgnore(args[0], appendFlag)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("%s .gitignore updated for %s.\n", ui.SecondaryStyle.Render(ui.IconCheck), args[0])
	},
}

func init() {
	ignoreCmd.Flags().BoolVarP(&appendFlag, "append", "a", false, "Append to existing .gitignore")
	rootCmd.AddCommand(ignoreCmd)
}
