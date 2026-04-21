package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/ui"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Ask questions about the repository",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		p := tea.NewProgram(ui.NewChatModel(cwd))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
