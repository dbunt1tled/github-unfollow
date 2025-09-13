package cmd

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

var follow = &cobra.Command{ //nolint:gochecknoglobals // need for init command
	Use:   "follow <username>",
	Short: "Follow users",
	Long:  "Follow <username> followers who not follow you",
	Args:  cobra.ExactArgs(1), //nolint:mnd // args count
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err                            error
			force                          bool
			cfg                            *config.Config
			username                       string
			followers, toFollow, following []string
		)
		username = args[0]
		cfg, err = config.Load()
		if err != nil {
			return err
		}
		force, err = cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("failed to get force flag: %w", err)
		}
		gm := git_hub_manager.NewGitHubManager(cfg.GitHubToken, cfg.GitHubUsername)
		followers, err = gm.GetFollowers(&username)
		if err != nil {
			return err
		}
		following, err = gm.GetFollowing()
		if err != nil {
			return err
		}
		following = append(following, cfg.GitHubUsername)
		toFollow = gm.DiffUsernames(following, followers)

		if len(toFollow) == 0 {
			fmt.Println("No users to follow")
			return nil
		}

		for i, user := range toFollow {
			fmt.Printf("%d. %s\n", i+1, user)
		}

		if force == false {
			fmt.Printf("\nYou are shure that you want to unfollow %d users?", len(toFollow))
			confirm := helper.GetInput("")
			if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
				fmt.Println("Operation canceled")
				return nil
			}
		}

		wp := worker.NewWorker(cfg.WorkerCount, cfg.QueueSize)
		wp.Start()

		for _, user := range toFollow {
			wp.AddTask(func() {
				err = gm.FollowUser(user, cfg.TimeDelay)
				if err != nil {
					fmt.Printf("Error follow %s: %v", user, err)
					os.Exit(1)
				}
				fmt.Printf("Followed %s\n", user)
			})

		}
		wp.Wait()
		wp.Stop()

		return nil
	},
}

//nolint:gochecknoinits // need for init command
func init() {
	rootCmd.AddCommand(follow)
	follow.Flags().BoolP("force", "f", false, "Force follow")
}
