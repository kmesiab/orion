package github

import (
	"context"
	"testing"

	gh "github.com/kmesiab/go-github-diff"
	"github.com/stretchr/testify/assert" // You might need to add testify to your project
)

// TestParsePullRequestURL tests the ParsePullRequestURL method.  This is only
// to detect potential regressions in this library.
func TestParsePullRequestURL(t *testing.T) {
	// Setup test cases
	testCases := []struct {
		name     string
		url      string
		expected *gh.PullRequestURL
		isErr    bool
	}{
		{
			name: "Valid URL",
			url:  "https://github.com/owner/repo/pull/123",
			expected: &gh.PullRequestURL{
				Owner:    "owner",
				Repo:     "repo",
				PRNumber: 123,
			},
			isErr: false,
		},
		{
			name:  "Invalid URL",
			url:   "https://invalid.com",
			isErr: true,
		},
	}

	client := &Client{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := client.ParsePullRequestURL(tc.url)
			if tc.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

// TestGetPullRequest tests the GetPullRequest method
// [INTEGRATION_TEST]
func TestGetPullRequest(t *testing.T) {
	const prURL = "https://github.com/kmesiab/orion/pull/4"

	client := &Client{}

	// Test the GetPullRequest method
	gitDiffs, err := client.GetPullRequest(context.Background(), prURL)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, gitDiffs)
}
