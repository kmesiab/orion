package github

import (
	"context"

	"github.com/google/go-github/v57/github"
	gh "github.com/kmesiab/go-github-diff"
)

type ClientInterface interface {
	ParsePullRequestURL(pullRequestURL string) (*gh.PullRequestURL, error)
	GetPullRequest(ctx context.Context, url string) ([]*gh.GitDiff, error)
	ParseGitDiff(diff string, ignoreList []string) []*gh.GitDiff
}

type Client struct{}

func (g *Client) ParsePullRequestURL(pullRequestURL string) (*gh.PullRequestURL, error) {
	return gh.ParsePullRequestURL(pullRequestURL)
}

func (g *Client) ParseGitDiff(diff string, ignoreList []string) []*gh.GitDiff {
	return gh.ParseGitDiff(diff, ignoreList)
}

func (g *Client) GetPullRequest(ctx context.Context, url string) ([]*gh.GitDiff, error) {
	var (
		err        error
		diffString string
		pr         *gh.PullRequestURL
	)

	if pr, err = gh.ParsePullRequestURL(url); err != nil {
		return nil, err
	}

	ghClient := github.NewClient(nil)

	if diffString, err = gh.GetPullRequest(ctx, pr, ghClient); err != nil {
		return nil, err
	}

	return gh.ParseGitDiff(diffString, []string{}), nil
}
