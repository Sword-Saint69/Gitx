package git

import (
	"os/exec"
	"strings"
)

type ReflogEntry struct {
	Hash      string
	Selector  string
	Operation string
	Subject   string
}

func GetReflog(path string) ([]ReflogEntry, error) {
	// go-git has limited reflog support, so we'll shell out for this specific high-level command
	// to get the human-readable operations accurately.
	cmd := exec.Command("git", "reflog", "-n", "20")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var entries []ReflogEntry
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 3 {
			continue
		}

		// Line format: hash selector operation: subject
		// Example: abc1234 HEAD@{0} commit: add feature
		hash := parts[0]
		selector := parts[1]
		rest := parts[2]

		opParts := strings.SplitN(rest, ": ", 2)
		op := opParts[0]
		subject := ""
		if len(opParts) > 1 {
			subject = opParts[1]
		}

		entries = append(entries, ReflogEntry{
			Hash:      hash,
			Selector:  selector,
			Operation: op,
			Subject:   subject,
		})
	}

	return entries, nil
}

func UndoTo(path, hash string) error {
	cmd := exec.Command("git", "reset", "--hard", hash)
	cmd.Dir = path
	return cmd.Run()
}
