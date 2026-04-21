package git

import (
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Hotspot struct {
	File  string
	Churn int
	LOC   int
	Score float64
}

func GetHotspots(path string) ([]Hotspot, error) {
	// 1. Get Churn
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	churnCounts := make(map[string]int)
	err = cIter.ForEach(func(c *object.Commit) error {
		stats, err := c.Stats()
		if err != nil {
			return nil
		}
		for _, s := range stats {
			churnCounts[s.Name]++
		}
		return nil
	})

	// 2. Get LOC for those files
	var hotspots []Hotspot
	for file, churn := range churnCounts {
		if strings.Contains(file, ".git") || strings.Contains(file, "node_modules") {
			continue
		}

		lines, _ := countLines(file) // Reusing existing countLines logic
		if lines == 0 {
			continue
		}

		hotspots = append(hotspots, Hotspot{
			File:  file,
			Churn: churn,
			LOC:   lines,
			Score: float64(churn * lines),
		})
	}

	sort.Slice(hotspots, func(i, j int) bool {
		return hotspots[i].Score > hotspots[j].Score
	})

	if len(hotspots) > 10 {
		hotspots = hotspots[:10]
	}

	return hotspots, nil
}
