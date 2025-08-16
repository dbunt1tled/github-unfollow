package git_hub_manager

import (
	"context"
	"fmt"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

type GitHubManager struct {
	client   *github.Client
	username string
}

func NewGitHubManager(token string, username string) *GitHubManager {
	var ctx context.Context = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GitHubManager{
		client:   client,
		username: username,
	}
}

func (gm *GitHubManager) GetFollowers() ([]string, error) {
	var allFollowers []string
	ctx := context.Background()
	opts := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := gm.client.Users.ListFollowers(ctx, gm.username, opts)
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
	opts := &github.ListOptions{PerPage: 100}
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

func (gm *GitHubManager) FindUserToUnfollow(followers []string, following []string) []string {
	followersSet := make(map[string]bool)
	for _, follower := range followers {
		followersSet[follower] = true
	}
	var toUnfollow []string
	for _, followingUser := range following {
		if !followersSet[followingUser] {
			toUnfollow = append(toUnfollow, followingUser)
		}
	}
	return toUnfollow
}

func (gm *GitHubManager) UnfollowUser(username string) error {
	ctx := context.Background()
	_, err := gm.client.Users.Unfollow(ctx, username)
	if err != nil {
		return fmt.Errorf("ошибка отписки от %s: %v", username, err)
	}
	return nil
}
