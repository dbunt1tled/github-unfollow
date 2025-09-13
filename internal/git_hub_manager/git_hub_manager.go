package git_hub_manager

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

type GitHubManager struct {
	client   *github.Client
	username string
}

func NewGitHubManager(token string, username string) *GitHubManager {
	var ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GitHubManager{
		client:   client,
		username: username,
	}
}

func (gm *GitHubManager) GetFollowers(name *string) ([]string, error) {
	var allFollowers []string
	username := gm.username
	if name != nil {
		username = *name
	}
	ctx := context.Background()
	opts := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := gm.client.Users.ListFollowers(ctx, username, opts)
		if err != nil {
			return nil, err
		}
		for _, follower := range followers {
			allFollowers = append(allFollowers, follower.GetLogin())
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return allFollowers, nil
}

func (gm *GitHubManager) GetFollowing() ([]string, error) {
	var allFollowing []string
	ctx := context.Background()
	opts := &github.ListOptions{PerPage: 200}
	for {
		following, resp, err := gm.client.Users.ListFollowing(ctx, gm.username, opts)
		if err != nil {
			return nil, err
		}
		for _, follower := range following {
			allFollowing = append(allFollowing, follower.GetLogin())
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return allFollowing, nil
}

func (gm *GitHubManager) DiffUsernames(users1 []string, users2 []string) []string {
	users1Set := make(map[string]bool)
	for _, u := range users1 {
		users1Set[u] = true
	}
	var usersDiff []string
	for _, u := range users2 {
		if !users1Set[u] {
			usersDiff = append(usersDiff, u)
		}
	}
	return usersDiff
}

func (gm *GitHubManager) UnfollowUser(username string, delay int) error {
	ctx := context.Background()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	_, err := gm.client.Users.Unfollow(ctx, username)
	if err != nil {
		return fmt.Errorf("error to unfollow %s: %v", username, err)
	}
	return nil
}

func (gm *GitHubManager) FollowUser(username string, delay int) error {
	ctx := context.Background()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	_, err := gm.client.Users.Follow(ctx, username)
	if err != nil {
		return fmt.Errorf("error to follow %s: %v", username, err)
	}
	return nil
}
