package git

import (
	"os/exec"
	"strings"
)

type SearchResult struct {
	Hash    string
	File    string
	Line    string
	Content string
}

func SearchHistory(path, query string) ([]SearchResult, error) {
	// git grep -i <query> $(git rev-list --all --max-count=100)
	// We limit to 100 commits for performance in this demo
	revCmd := exec.Command("git", "rev-list", "--all", "--max-count", "100")
	revCmd.Dir = path
	revsOut, err := revCmd.Output()
	if err != nil {
		return nil, err
	}
	revs := strings.Fields(string(revsOut))

	if len(revs) == 0 {
		return nil, nil
	}

	args := append([]string{"grep", "-i", "-n", query}, revs...)
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	out, _ := cmd.CombinedOutput() // Grep returns error code 1 if no match

	lines := strings.Split(string(out), "\n")
	var results []SearchResult
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}
		// Format: hash:file:line:content
		results = append(results, SearchResult{
			Hash:    parts[0],
			File:    parts[1],
			Line:    parts[2],
			Content: strings.TrimSpace(strings.Join(parts[3:], ":")), // Might have more colons in content
		})
	}

	return results, nil
}
