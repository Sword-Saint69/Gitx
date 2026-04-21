package git

import (
	"os/exec"
	"strings"
)

type GhostCommit struct {
	Hash    string
	Subject string
}

func FindLostCommits(path string) ([]GhostCommit, error) {
	// 1. Run fsck to find dangling commits
	cmd := exec.Command("git", "fsck", "--lost-found")
	cmd.Dir = path
	out, _ := cmd.CombinedOutput()

	var ghosts []GhostCommit
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "dangling commit") {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				hash := parts[2]
				// Get subject for the ghost commit
				subCmd := exec.Command("git", "log", "-1", "--format=%s", hash)
				subCmd.Dir = path
				subOut, _ := subCmd.Output()
				
				ghosts = append(ghosts, GhostCommit{
					Hash:    hash,
					Subject: strings.TrimSpace(string(subOut)),
				})
			}
		}
	}

	return ghosts, nil
}
