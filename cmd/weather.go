package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/gitx/internal/git"
	"github.com/user/gitx/internal/ui"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Show a metaphorical weather report for the repository",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		report, err := git.GetWeather(cwd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(ui.RenderWeather(report))
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
