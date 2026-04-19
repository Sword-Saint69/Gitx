package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type TimelineStats struct {
	Days map[string]int // Format "2006-01-02"
}

func GetTimeline(path string) (*TimelineStats, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{All: true})
	if err != nil {
		return nil, err
	}

	days := make(map[string]int)
	err = cIter.ForEach(func(c *object.Commit) error {
		day := c.Author.When.Format("2006-01-02")
		days[day]++
		return nil
	})

	return &TimelineStats{Days: days}, err
}
