package main

import (
	"fmt"
	"github-user-activitiy/github"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	events, err := github.GetGitHubEventFromUsername("Javanatole")
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		fmt.Println(event.Type)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
