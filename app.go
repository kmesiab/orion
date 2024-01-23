package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kmesiab/orion/github"
)

// App represents the main application with its dependencies.
// It includes the application context and the GitHub service
// which is used for interactions with the GitHub API.
type App struct {
	ctx           context.Context        // Application context for managing lifecycle and cancellation.
	GithubService github.ClientInterface // Interface to the GitHub client for API interactions.
}

// Startup is called when the app starts.
// The context is saved, so we can call the runtime method.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// ProcessPullRequest processes a given GitHub Pull Request URL and returns a combined
// string representation of all Git diffs in the pull request. This function is designed
// to be bound to a TypeScript function, enabling the frontend to interact with the
// GitHub API through this Go method.
//
// Parameters:
//   - url: A string representing the full URL of a GitHub pull request.
//
// Returns:
//   - A string containing the combined diffs of the pull request, or an error message
//     if the URL is invalid or if there's an error in fetching the pull request.
//
// The function first attempts to parse the provided URL using the GithubService. If the
// URL is invalid, it returns an error message. Then, it fetches the pull request data
// from GitHub and combines the diffs of each file in the pull request into a single
// string, which is then returned.
//
// If there's an error in fetching the pull request, an error message is returned.
// This function is particularly useful in applications that need to process or display
// the changes in a pull request, such as in code review tools or development dashboards.
func (a *App) ProcessPullRequest(url string) string {
	// Parse the pull request URL to validate it

	var (
		err   error
		diffs string
	)

	if _, err = a.GithubService.ParsePullRequestURL(url); err != nil {
		return fmt.Errorf("invalid GitHub pull request url. error: %s", err).Error()
	}

	// Fetch the pull request details using the GitHub service
	if diffs, err = a.GithubService.GetPullRequest(a.ctx, url); err != nil {
		return fmt.Errorf("error getting pull request. error: %s", err).Error()
	}

	return diffs
}

// GetCodeReviewFromAPI takes a string containing the combined diffs of a pull request
// and returns a string containing the code review. This function is designed to be
// bound to a TypeScript function, enabling the frontend to interact with the
// GitHub API through this Go method.
func (a *App) GetCodeReviewFromAPI(diffs string) string {
	//This function is not implemented yet.
	//It currently starts a timer (sleeps for 2 seconds) and returns a string.
	//Start a timer for 5 seconds
	time.Sleep(2 * time.Second)
	return `
	I've reviewed your recent pull request, and overall, the changes look great! I appreciate your efforts. However, I have a few suggestions for improvement:

1. **Variable Naming:**
   - Consider using more descriptive variable names to enhance code readability.

2. **Comments:**
   - Add comments to explain complex logic or any non-trivial parts of the code.

3. **Error Handling:**
   - Ensure proper error handling for critical sections of the code.

4. **Code Duplication:**
   - Watch out for code duplication. It's always good to refactor common logic into reusable functions.

5. **Testing:**
   - Don't forget to include test cases for the new functionality.

Please address these points, and once done, feel free to merge the changes. Great job overall!`
}
