package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type LeaderboardEntry struct {
	Author    string
	Commits   int
	Additions int
	Deletions int
	Files     map[string]bool
}

func GetLeaderboards(path string) ([]LeaderboardEntry, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	stats := make(map[string]*LeaderboardEntry)
	err = cIter.ForEach(func(c *object.Commit) error {
		if stats[c.Author.Email] == nil {
			stats[c.Author.Email] = &LeaderboardEntry{
				Author: c.Author.Name,
				Files:  make(map[string]bool),
			}
		}
		
		e := stats[c.Author.Email]
		e.Commits++
		
		s, err := c.Stats()
		if err == nil {
			for _, cs := range s {
				e.Additions += cs.Addition
				e.Deletions += cs.Deletion
				e.Files[cs.Name] = true
			}
		}
		return nil
	})

	var result []LeaderboardEntry
	for _, v := range stats {
		result = append(result, *v)
	}

	return result, nil
}
