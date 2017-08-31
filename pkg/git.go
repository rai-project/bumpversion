package bumpversion

import (
	"github.com/k0kubun/pp"
	git "srcd.works/go-git.v4"
)

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

func (g *gitRepo) Status() (git.Status, error) {
	repo := g.repo
	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}
	return worktree.Status()
}

func (g *gitRepo) Commit() {
}

func (g *gitRepo) LatestTagInfo() {
  repo := g.repo
  repo.
	pp.Println(repo.Tags())
}
