package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dbunt1tled/github-unfollow/internal/git_hub_manager"
	"github.com/dbunt1tled/github-unfollow/internal/helper"
	"github.com/dbunt1tled/github-unfollow/internal/worker"

	"github.com/joho/godotenv"
)

func main() {
	var (
		err    error
		wc, qs int
	)
	err = godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	username := os.Getenv("GITHUB_USER")
	token := os.Getenv("GITHUB_TOKEN")

	wc, err = strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		wc = 10
	}
	qs, err = strconv.Atoi(os.Getenv("QUEUE_SIZE"))
	if err != nil {
		qs = 10
	}

	if username == "" || token == "" {
		fmt.Println("Please set GITHUB_USER and GITHUB_TOKEN environment variables")
		return
	}
	gm := git_hub_manager.NewGitHubManager(token, username)
	followers, err := gm.GetFollowers()
	if err != nil {
		log.Fatal(err)
	}
	following, err := gm.GetFollowing()
	if err != nil {
		log.Fatal(err)
	}
	toUnfollow := gm.FindUserToUnfollow(followers, following)
	for i, user := range toUnfollow {
		fmt.Printf("%d. %s\n", i+1, user)
	}
	if len(toUnfollow) == 0 {
		fmt.Println("No users to unfollow")
		return
	}
	force := flag.Bool("force", false, "Force the operation")
	flag.Parse()
	if *force == false {
		fmt.Printf("\nYou are shure that you want to unfollow %d users?", len(toUnfollow))
		confirm := helper.GetInput("")
		if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
			fmt.Println("Operation canceled")
			return
		}
	}

	wp := worker.NewWorker(wc, qs)
	wp.Start()

	for _, user := range toUnfollow {
		wp.AddTask(func() {
			err = gm.UnfollowUser(user)
			if err != nil {
				log.Fatal(err)
			}
		})

	}
	wp.Wait()
	wp.Stop()
}
