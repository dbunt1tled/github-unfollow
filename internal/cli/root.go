package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CLI struct {
	rootCmd *cobra.Command
}

func NewCLI() *CLI {
	cli := &CLI{
		rootCmd: &cobra.Command{
			Use:   "github-unfollow",
			Short: "Follow & Unfollow GitHub users",
			Long:  "Manager GitHub users in your account (Follow/Unfollow)",
		},
	}

	cli.rootCmd.AddCommand(NewFollowCommand())
	cli.rootCmd.AddCommand(NewUnFollowCommand())

	return cli
}

func (c *CLI) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err) //nolint:forbidigo // print error
		os.Exit(1)
	}
}
