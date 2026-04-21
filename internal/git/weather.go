package git

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type WeatherReport struct {
	Condition   string // Sunny, Cloudy, Stormy, etc.
	Temperature int    // Author Diversity (Cold/Hot)
	WindSpeed   int    // Churn rate (commits/week)
	Humidity    int    // Staged changes (%)
	Summary     string
}

func GetWeather(path string) (*WeatherReport, error) {
	r, gitErr := git.PlainOpen(path)
	if gitErr != nil {
		return nil, gitErr
	}

	// 1. Calculate Wind (Churn in last 7 days)
	cIter, _ := r.Log(&git.LogOptions{All: true})
	wind := 0
	authors := make(map[string]bool)
	since := time.Now().AddDate(0, 0, -7)
	
	_ = cIter.ForEach(func(c *object.Commit) error {
		authors[c.Author.Email] = true
		if c.Author.When.After(since) {
			wind++
		}
		return nil
	})

	// 2. Humidity (Staged vs Working)
	w, _ := r.Worktree()
	status, _ := w.Status()
	staged := 0
	for _, s := range status {
		if s.Staging != git.Unmodified {
			staged++
		}
	}
	humidity := staged * 10 // Heuristic: 10 files = 100% humidity

	// 3. Condition (Based on Hotspots)
	hotspots, _ := GetHotspots(path)
	condition := "Sunny"
	summary := "Conditions are ideal for stable development."
	
	if len(hotspots) > 0 && hotspots[0].Score > 5000 {
		condition = "Thunderstorm"
		summary = "High-risk code found. Expect turbulence during deployment."
	} else if wind > 20 {
		condition = "Windy"
		summary = "Rapid changes detected. Ensure code reviews are thorough."
	} else if humidity > 50 {
		condition = "Foggy"
		summary = "Heavy uncommitted changes. Visibility is low."
	}

	return &WeatherReport{
		Condition:   condition,
		Temperature: len(authors) * 5 + 60, // 60F base + authors
		WindSpeed:   wind,
		Humidity:    humidity,
		Summary:     summary,
	}, nil
}
