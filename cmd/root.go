package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitx",
	Short: "gitx is the ultimate git power toolkit",
	Long: `A single CLI that replaces dozens of git aliases, scripts, and third-party tools 
with one unified, beautiful command-line experience.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Root flags if needed
}
