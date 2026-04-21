package git

import (
	"os/exec"
	"strings"
	"time"
)

type PerfReport struct {
	StatusTime   time.Duration
	LogTime      time.Duration
	FileCount    int
	Suggestions  []string
}

func ProfileGit(path string) (*PerfReport, error) {
	// 1. Measure Status
	start := time.Now()
	cmd := exec.Command("git", "status")
	cmd.Dir = path
	_ = cmd.Run()
	statusTime := time.Since(start)

	// 2. Measure Log walking
	start = time.Now()
	cmd = exec.Command("git", "rev-list", "--count", "--all")
	cmd.Dir = path
	_ = cmd.Run()
	logTime := time.Since(start)

	// 3. Count files
	cmd = exec.Command("git", "ls-files")
	cmd.Dir = path
	out, _ := cmd.Output()
	files := len(strings.Split(string(out), "\n"))

	var suggestions []string
	if statusTime > 300*time.Millisecond {
		suggestions = append(suggestions, "Enable FSMonitor: git config core.fsmonitor true")
	}
	if files > 1000 {
		suggestions = append(suggestions, "Enable CommitGraph: git config core.commitGraph true")
		suggestions = append(suggestions, "Enable FileMode Cache: git config core.filemode false (if on Win)")
	}

	return &PerfReport{
		StatusTime:  statusTime,
		LogTime:     logTime,
		FileCount:   files,
		Suggestions: suggestions,
	}, nil
}
