package git

import (
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type StandupEntry struct {
	Repo    string
	Commits []string
}

func GetStandup(path string) ([]StandupEntry, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	since := time.Now().AddDate(0, 0, -1)
	var commits []string
	err = cIter.ForEach(func(c *object.Commit) error {
		if c.Author.When.After(since) {
			commits = append(commits, strings.Split(c.Message, "\n")[0])
		}
		return nil
	})

	return []StandupEntry{{Repo: "Current", Commits: commits}}, err
}
