package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/dbunt1tled/github-unfollow/internal/config"
	"github.com/dbunt1tled/github-unfollow/internal/git_hub_manager"
	"github.com/dbunt1tled/github-unfollow/internal/helper"
	"github.com/dbunt1tled/github-unfollow/internal/worker"
	"github.com/spf13/cobra"
)

func NewUnFollowCommand() *cobra.Command {
	cmd := unFollowCommand()
	cmd.Flags().BoolP("force", "f", false, "Force unfollow")

	return cmd
}

func unFollowCommand() *cobra.Command {
	return &cobra.Command{ //nolint:gochecknoglobals // need for init command
		Use:   "unfollow",
		Short: "Unfollow users",
		Long:  "Unfollow users who not follow you",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				err                              error
				force                            bool
				cfg                              *config.Config
				followers, toUnfollow, following []string
			)
			cfg, err = config.Load()
			if err != nil {
				return err
			}
			force, err = cmd.Flags().GetBool("force")
			if err != nil {
				return fmt.Errorf("failed to get force flag: %w", err)
			}
			gm := git_hub_manager.NewGitHubManager(cfg.GitHubToken, cfg.GitHubUsername)
			followers, err = gm.GetFollowers(nil)
			if err != nil {
				return err
			}
			following, err = gm.GetFollowing()
			if err != nil {
				return err
			}

			toUnfollow = gm.DiffUsernames(followers, following)

			for i, user := range toUnfollow {
				fmt.Printf("%d. %s\n", i+1, user)
			}
			if len(toUnfollow) == 0 {
				fmt.Println("No users to unfollow")
				return nil
			}
			if force == false {
				fmt.Printf("\nYou are shure that you want to unfollow %d users?", len(toUnfollow))
				confirm := helper.GetInput("")
				if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
					fmt.Println("Operation canceled")
					return nil
				}
			}

			wp := worker.NewWorker(cfg.WorkerCount, cfg.QueueSize)
			wp.Start()

			for _, user := range toUnfollow {
				wp.AddTask(func() {
					err = gm.UnfollowUser(user, cfg.TimeDelay)
					if err != nil {
						fmt.Printf("Error unfollow %s: %v", user, err)
						os.Exit(1)
					}
					fmt.Printf("UnFollowed %s\n", user)
				})

			}
			wp.Wait()
			wp.Stop()

			return nil
		},
	}
}
