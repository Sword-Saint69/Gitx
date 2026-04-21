package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/ui"
)

var wipCmd = &cobra.Command{
	Use:   "wip",
	Short: "Quick save work-in-progress",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		r, err := git.PlainOpen(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		w, err := r.Worktree()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err = w.AddWithOptions(&git.AddOptions{All: true})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		msg := fmt.Sprintf("WIP: %s", time.Now().Format("2006-01-02 15:04:05"))
		_, err = w.Commit(msg, &git.CommitOptions{})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " Work saved as WIP commit."))
		fmt.Println(ui.SubtleStyle.Render(msg))
	},
}

func init() {
	rootCmd.AddCommand(wipCmd)
}
