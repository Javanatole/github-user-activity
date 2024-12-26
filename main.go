package main

import (
	"fmt"
	"github-user-activitiy/github"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: go run main.go Javanatole")
		return
	}

	events, err := github.GetGitHubEventFromUsername(os.Args[1])
	if err != nil {
		fmt.Println("Can't fetch github events, please check the username or the network")
		return
	}
	repositories, err := github.GetEventsByRepositories(events)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output:")
	for repo, events := range repositories {
		if len(events.PushEvents) > 0 {
			fmt.Printf("- Push %d commits to %s\n", len(events.PushEvents), repo)
		}
		if events.Issues > 0 {
			fmt.Printf("- Open %d issues in %s\n", events.Issues, repo)
		}
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
