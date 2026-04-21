package git

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type LogEntry struct {
	Hash    string
	Author  string
	Date    string
	Subject string
}

func GetLog(path string, count int) ([]LogEntry, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	var entries []LogEntry
	i := 0
	err = cIter.ForEach(func(c *object.Commit) error {
		if i >= count {
			return nil
		}
		entries = append(entries, LogEntry{
			Hash:    c.Hash.String(),
			Author:  c.Author.Name,
			Date:    c.Author.When.Format("2006-01-02"),
			Subject: strings.Split(c.Message, "\n")[0],
		})
		i++
		return nil
	})

	return entries, err
}
