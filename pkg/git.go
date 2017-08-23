package bumpversion

import git "srcd.works/go-git.v4"

type gitRepo struct {
	path string
	repo *git.Repository
}

func newGit(path string) (*gitRepo, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &gitRepo{
		path: path,
		repo: repo,
	}, nil
}
