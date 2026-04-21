package git

import (
	"os/exec"
	"strings"
)

func GenerateCommitMsg(path string) (string, error) {
	// 1. Get staged changes
	cmd := exec.Command("git", "diff", "--staged")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	diff := string(out)
	if diff == "" {
		return "No changes staged for commit.", nil
	}

	// 2. Simple Heuristic Analysis
	lines := strings.Split(diff, "\n")
	modifiedFiles := []string{}
	isFeature := false
	isFix := false
	
	for _, line := range lines {
		if strings.HasPrefix(line, "+++ b/") {
			file := strings.TrimPrefix(line, "+++ b/")
			modifiedFiles = append(modifiedFiles, file)
		}
		lower := strings.ToLower(line)
		if strings.Contains(lower, "error") || strings.Contains(lower, "bug") || strings.Contains(lower, "fix") {
			isFix = true
		}
		if strings.Contains(lower, "func ") || strings.Contains(lower, "type ") {
			isFeature = true
		}
	}

	// 3. Construct message
	prefix := "chore"
	if isFeature {
		prefix = "feat"
	}
	if isFix {
		prefix = "fix"
	}

	scope := "other"
	if len(modifiedFiles) > 0 {
		scope = strings.Split(modifiedFiles[0], ".")[0]
		if strings.Contains(scope, "/") {
			parts := strings.Split(scope, "/")
			scope = parts[len(parts)-1]
		}
	}

	msg := prefix + "(" + scope + "): work on " + strings.Join(modifiedFiles, ", ")
	if len(msg) > 70 {
		msg = msg[:67] + "..."
	}

	return msg, nil
}
