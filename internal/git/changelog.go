package git

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type ChangelogEntry struct {
	Type    string
	Subject string
}

func GetChangelog(path string) (map[string][]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	changes := map[string][]string{
		"Features": []string{},
		"Fixes":    []string{},
		"Other":    []string{},
	}

	err = cIter.ForEach(func(c *object.Commit) error {
		msg := strings.TrimSpace(c.Message)
		lower := strings.ToLower(msg)

		if strings.HasPrefix(lower, "feat") {
			changes["Features"] = append(changes["Features"], msg)
		} else if strings.HasPrefix(lower, "fix") {
			changes["Fixes"] = append(changes["Fixes"], msg)
		} else {
			changes["Other"] = append(changes["Other"], msg)
		}
		return nil
	})

	return changes, err
}
