package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create or list lightweight named snapshots",
}

var snapshotCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new snapshot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		err := git.CreateSnapshot(cwd, args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("✅ Snapshot '%s' created.\n", args[0])
	},
}

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		snapshots, err := git.ListSnapshots(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(snapshots) == 0 {
			fmt.Println("No gitx snapshots found.")
			return
		}

		fmt.Println(ui.HeaderStyle.Render(" 📸 Gitx Snapshots "))
		for _, s := range snapshots {
			fmt.Printf("%s %s\n", ui.AccentStyle.Render("["+s.Index+"]"), ui.ValueStyle.Render(s.Message))
		}
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotCreateCmd)
	snapshotCmd.AddCommand(snapshotListCmd)
	rootCmd.AddCommand(snapshotCmd)
}
