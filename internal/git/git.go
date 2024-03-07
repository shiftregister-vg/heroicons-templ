package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func Clone(repo string) *git.Repository {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repo,
	})
	if err != nil {
		panic(err)
	}
	return r
}
