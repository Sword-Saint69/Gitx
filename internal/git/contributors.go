package git

import (
	"sort"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Contributor struct {
	Name    string
	Email   string
	Commits int
}

func GetContributors(path string) ([]Contributor, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	counts := make(map[string]*Contributor)
	err = cIter.ForEach(func(c *object.Commit) error {
		email := c.Author.Email
		if _, ok := counts[email]; !ok {
			counts[email] = &Contributor{
				Name:  c.Author.Name,
				Email: email,
			}
		}
		counts[email].Commits++
		return nil
	})
	if err != nil {
		return nil, err
	}

	var contributors []Contributor
	for _, c := range counts {
		contributors = append(contributors, *c)
	}

	sort.Slice(contributors, func(i, j int) bool {
		return contributors[i].Commits > contributors[j].Commits
	})

	return contributors, nil
}
