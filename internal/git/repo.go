package git

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RepoStats struct {
	Name         string
	Commits      int
	Contributors int
	Branches     int
	Age          string
	LastCommit   time.Time
}

func GetStats(path string) (*RepoStats, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	// Commits and Contributors
	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	commitCount := 0
	authors := make(map[string]bool)
	var firstCommit, lastCommit time.Time

	err = cIter.ForEach(func(c *object.Commit) error {
		commitCount++
		authors[c.Author.Email] = true
		if firstCommit.IsZero() || c.Author.When.Before(firstCommit) {
			firstCommit = c.Author.When
		}
		if c.Author.When.After(lastCommit) {
			lastCommit = c.Author.When
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Branches
	bIter, err := r.Branches()
	if err != nil {
		return nil, err
	}
	branchCount := 0
	err = bIter.ForEach(func(ref *plumbing.Reference) error {
		branchCount++
		return nil
	})

	age := time.Since(firstCommit)
	ageStr := formatDuration(age)

	return &RepoStats{
		Name:         "Current Repo", // Could get from config or remote
		Commits:      commitCount,
		Contributors: len(authors),
		Branches:     branchCount,
		Age:          ageStr,
		LastCommit:   lastCommit,
	}, nil
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	if days < 30 {
		return time.Now().Add(-d).Format("Jan 2006")
	}
	months := days / 30
	years := months / 12
	if years > 0 {
		return time.Unix(0, 0).Add(d).Format("2006-01-02") // Simplified
	}
	return time.Unix(0, 0).Add(d).Format("Jan 2006")
}
