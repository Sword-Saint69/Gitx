package git

import (
	"bufio"
	"os/exec"
	"strings"
)

type ReviewIssue struct {
	File     string
	Line     int
	Message  string
	Severity string // LOW, MEDIUM, HIGH
}

func PerformReview(path string) ([]ReviewIssue, error) {
	// 1. Get current staged diff
	cmd := exec.Command("git", "diff", "--staged")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var issues []ReviewIssue
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	
	currentFile := ""
	lineNum := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		
		if strings.HasPrefix(line, "+++ b/") {
			currentFile = strings.TrimPrefix(line, "+++ b/")
			lineNum = 0
			continue
		}
		
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			lineNum++
			content := strings.TrimPrefix(line, "+")
			
			// Rule 1: Forgotten TODOs
			if strings.Contains(strings.ToUpper(content), "TODO") || strings.Contains(strings.ToUpper(content), "FIXME") {
				issues = append(issues, ReviewIssue{
					File:     currentFile,
					Line:     lineNum,
					Message:  "Found a forgotten TODO/FIXME.",
					Severity: "MEDIUM",
				})
			}
			
			// Rule 2: Large lines (potential secrets or minified code)
			if len(content) > 200 {
				issues = append(issues, ReviewIssue{
					File:     currentFile,
					Line:     lineNum,
					Message:  "Extremely long line detected (>200 chars).",
					Severity: "LOW",
				})
			}

			// Rule 3: Debugging statements
			if strings.Contains(content, "fmt.Println") || strings.Contains(content, "console.log") {
				issues = append(issues, ReviewIssue{
					File:     currentFile,
					Line:     lineNum,
					Message:  "Possible debug print found.",
					Severity: "MEDIUM",
				})
			}
		}
	}

	return issues, nil
}
