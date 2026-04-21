package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var useAI bool

var genMsgCmd = &cobra.Command{
	Use:   "gen-msg",
	Short: "Suggest a commit message based on staged changes",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		
		var msg string
		var err error

		if useAI {
			fmt.Printf("%s Consultig Gemini AI for a better message...\n", ui.AccentStyle.Render("[*]"))
			// Get staged diff
			diffCmd := exec.Command("git", "diff", "--staged")
			diffCmd.Dir = cwd
			diffOut, _ := diffCmd.Output()
			
			msg, err = git.GenerateAICommitMsg(string(diffOut))
		} else {
			msg, err = git.GenerateCommitMsg(cwd)
		}

		if err != nil {
			fmt.Printf("%s %v\n", ui.ErrorStyle.Render("Error:"), err)
			if useAI {
				fmt.Println(ui.SubtleStyle.Render("Note: Ensure GEMINI_API_KEY is set in your environment."))
			}
			return
		}

		p := tea.NewProgram(ui.NewGenMsgModel(msg))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	genMsgCmd.Flags().BoolVarP(&useAI, "ai", "a", false, "Use Gemini AI to generate the message")
	rootCmd.AddCommand(genMsgCmd)
}
