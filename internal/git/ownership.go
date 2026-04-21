package git

import (
	"os/exec"
	"sort"
	"strings"
)

type Ownership struct {
	Author string
	Lines  int
	Pct    float64
}

func GetOwnership(path, file string) ([]Ownership, error) {
	// Use git blame --line-porcelain to get detailed author info for every line
	cmd := exec.Command("git", "blame", "--line-porcelain", file)
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	total := 0
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "author ") {
			author := strings.TrimPrefix(line, "author ")
			counts[author]++
			total++
		}
	}

	var ownership []Ownership
	for auth, count := range counts {
		ownership = append(ownership, Ownership{
			Author: auth,
			Lines:  count,
			Pct:    float64(count) / float64(total) * 100,
		})
	}

	sort.Slice(ownership, func(i, j int) bool {
		return ownership[i].Lines > ownership[j].Lines
	})

	return ownership, nil
}
