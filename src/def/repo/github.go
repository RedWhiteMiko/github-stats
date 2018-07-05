package repo

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	client *github.Client
	repo   string
	owner  string
}

// Init - Initialize
func Init(owner string, repo string) *Github {
	context := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "29b3f7177ceb0c04bbd71a8d04bc8ff71382cadd"},
	)
	token := oauth2.NewClient(context, tokenSource)

	git := new(Github)
	git.repo = repo
	git.owner = owner
	git.client = github.NewClient(token)
	return git
}

func (git *Github) GetCloneURL() (string, error) {
	info, _, err := git.client.Repositories.Get(context.Background(), git.owner, git.repo)
	if err != nil {
		return "", err
	}
	return info.GetCloneURL(), nil
}

func (git *Github) GetLastCommit() (*github.RepositoryCommit, error) {
	commits, err := git.ListCommits()
	if err != nil {
		return nil, err
	}
	return commits[0], nil
}

func (git *Github) ListCommits() ([]*github.RepositoryCommit, error) {
	commits, _, err := git.client.Repositories.ListCommits(context.Background(), git.owner, git.repo, nil)
	if err != nil {
		return nil, err
	}
	return commits, nil
}
