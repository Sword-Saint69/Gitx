package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type PRSummary struct {
	Title       string
	Description string
	Stats       string
}

func GeneratePRSummary(path string) (*PRSummary, error) {
	// 1. Get current branch
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchCmd.Dir = path
	branchOut, _ := branchCmd.Output()
	branch := strings.TrimSpace(string(branchOut))

	// 2. Identify base branch (try main then master)
	base := "main"
	checkCmd := exec.Command("git", "rev-parse", "--verify", "main")
	checkCmd.Dir = path
	if err := checkCmd.Run(); err != nil {
		base = "master"
	}

	// 3. Get commit logs
	logCmd := exec.Command("git", "log", base+".."+branch, "--oneline")
	logCmd.Dir = path
	logOut, _ := logCmd.Output()

	// 4. Get diff stats
	statCmd := exec.Command("git", "diff", base+".."+branch, "--stat")
	statCmd.Dir = path
	statOut, _ := statCmd.Output()

	summary := &PRSummary{
		Title:       fmt.Sprintf("PR: %s", strings.ReplaceAll(branch, "-", " ")),
		Description: string(logOut),
		Stats:       string(statOut),
	}

	return summary, nil
}
