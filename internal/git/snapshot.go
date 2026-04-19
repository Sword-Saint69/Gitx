package git

import (
	"os/exec"
	"strings"
)

type Snapshot struct {
	Index   string
	Message string
}

func CreateSnapshot(path, name string) error {
	msg := "gitx snapshot: " + name
	cmd := exec.Command("git", "stash", "push", "-m", msg)
	cmd.Dir = path
	return cmd.Run()
}

func ListSnapshots(path string) ([]Snapshot, error) {
	cmd := exec.Command("git", "stash", "list")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var snapshots []Snapshot
	for _, line := range lines {
		if strings.Contains(line, "gitx snapshot: ") {
			parts := strings.SplitN(line, ": ", 3)
			if len(parts) < 3 {
				continue
			}
			index := strings.Trim(parts[0], "stash@{}")
			msg := strings.TrimPrefix(parts[2], "gitx snapshot: ")
			snapshots = append(snapshots, Snapshot{
				Index:   index,
				Message: msg,
			})
		}
	}
	return snapshots, nil
}

func RestoreSnapshot(path, index string) error {
	cmd := exec.Command("git", "stash", "apply", "stash@{"+index+"}")
	cmd.Dir = path
	return cmd.Run()
}
