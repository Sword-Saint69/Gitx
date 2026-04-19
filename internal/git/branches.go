package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Branch struct {
	Name     string
	IsMerged bool
	IsRemote bool
}

func GetBranches(path string) ([]Branch, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	bIter, err := r.Branches()
	if err != nil {
		return nil, err
	}

	var branches []Branch
	// In a real implementation, we'd check against 'main' or 'master' for merged status.
	// For now, let's just list them.
	err = bIter.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, Branch{
			Name:     ref.Name().Short(),
			IsMerged: false, // Placeholder
			IsRemote: false,
		})
		return nil
	})

	return branches, err
}

func DeleteBranch(path string, name string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	return r.Storer.RemoveReference(plumbing.NewBranchReferenceName(name))
}
