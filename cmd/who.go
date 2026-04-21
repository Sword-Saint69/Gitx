package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var whoCmd = &cobra.Command{
	Use:   "who [file]",
	Short: "Show who owns each section of a file (enhanced blame)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		ownership, err := git.GetOwnership(cwd, args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" FILE OWNERSHIP "))
		fmt.Printf("%s %s\n\n", ui.SubtleStyle.Render("File:"), ui.ValueStyle.Render(args[0]))

		for _, o := range ownership {
			bar := ui.SecondaryStyle.Render(strings.Repeat("█", int(o.Pct/2)))
			fmt.Printf("%-25s %s %5.1f%%\n", ui.PrimaryStyle.Render(o.Author), bar, o.Pct)
		}
	},
}

func init() {
	rootCmd.AddCommand(whoCmd)
}
