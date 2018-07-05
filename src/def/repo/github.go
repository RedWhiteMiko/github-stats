package repo

import (
	"context"
	"os"

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
	git := new(Github)
	git.repo = repo
	git.owner = owner

	// Use Token if defined
	githubToken := os.Getenv("GITHUB_TOKEN")
	if len(githubToken) == 0 {
		git.client = github.NewClient(nil)
	} else {
		context := context.Background()
		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		token := oauth2.NewClient(context, tokenSource)
		git.client = github.NewClient(token)
	}
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
