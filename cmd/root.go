package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint:gochecknoglobals // need for init commands
	Use:   "github-unfollow",
	Short: "Follow & Unfollow GitHub users",
	Long:  "Manager GitHub users in your account (Follow/Unfollow)",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err) //nolint:forbidigo // print error
		os.Exit(1)
	}
}
