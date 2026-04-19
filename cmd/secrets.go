package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/scanner"
	"github.com/user/gitx/internal/ui"
)

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Security & Secrets tools",
}

var secretsScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan commit history and files for leaked secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		matches, err := scanner.ScanSecrets(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		p := tea.NewProgram(ui.NewSecretsModel(matches))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	secretsCmd.AddCommand(secretsScanCmd)
	rootCmd.AddCommand(secretsCmd)
}
