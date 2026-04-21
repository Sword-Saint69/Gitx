package git

import (
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type ChurnEntry struct {
	File  string
	Count int
}

type BusFactorEntry struct {
	File    string
	Authors int
}

func GetChurn(path string, top int) ([]ChurnEntry, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	err = cIter.ForEach(func(c *object.Commit) error {
		// Get changes in this commit
		stats, err := c.Stats()
		if err != nil {
			return nil
		}
		for _, s := range stats {
			counts[s.Name]++
		}
		return nil
	})

	var entries []ChurnEntry
	for f, c := range counts {
		if strings.Contains(f, ".git") {
			continue
		}
		entries = append(entries, ChurnEntry{f, c})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Count > entries[j].Count
	})

	if len(entries) > top {
		entries = entries[:top]
	}

	return entries, nil
}

func GetBusFactor(path string) ([]BusFactorEntry, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	fileAuthors := make(map[string]map[string]bool)
	err = cIter.ForEach(func(c *object.Commit) error {
		stats, err := c.Stats()
		if err != nil {
			return nil
		}
		for _, s := range stats {
			if fileAuthors[s.Name] == nil {
				fileAuthors[s.Name] = make(map[string]bool)
			}
			fileAuthors[s.Name][c.Author.Email] = true
		}
		return nil
	})

	var entries []BusFactorEntry
	for f, authors := range fileAuthors {
		if strings.Contains(f, ".git") {
			continue
		}
		entries = append(entries, BusFactorEntry{f, len(authors)})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Authors < entries[j].Authors
	})

	return entries, nil
}
