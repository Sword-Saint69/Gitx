package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/ui"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system and repository for issues",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.HeaderStyle.Render(" SYSTEM DIAGNOSTICS "))
		fmt.Println()

		check("Git Installation", "git version")
		check("User Config-Name", "git config user.name")
		check("User Config-Email", "git config user.email")
		check("Remote Configuration", "git remote -v")
		
		fmt.Println()
		fmt.Println(ui.SubtleStyle.Render("Diagnostics complete."))
	},
}

func check(label, command string) {
	fmt.Printf("%-25s ", ui.LabelStyle.Render(label))
	
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()
	
	if err != nil {
		fmt.Println(ui.ErrorStyle.Render(ui.IconWarning + " FAILED"))
	} else {
		val := strings.TrimSpace(string(out))
		if val == "" {
			fmt.Println(ui.WarningStyle.Render("MISSING"))
		} else {
			fmt.Println(ui.SecondaryStyle.Render(ui.IconCheck + " ") + ui.ValueStyle.Render(val))
		}
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
