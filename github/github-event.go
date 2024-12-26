package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetGitHubEventFromUsername get GitHub events from username
func GetGitHubEventFromUsername(username string) ([]GitHubEvent, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return []GitHubEvent{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Can't close response body")
		}
	}(response.Body)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []GitHubEvent{}, err
	}
	var events []GitHubEvent
	err = json.Unmarshal(data, &events)
	if err != nil {
		return []GitHubEvent{}, err
	}
	return events, nil
}

func GetEventsByRepositories(events []GitHubEvent) (map[string]*EventsByRepo, error) {
	eventsByRepositories := make(map[string]*EventsByRepo)
	for _, event := range events {
		eventsByRepo, ok := eventsByRepositories[event.Repo.Name]
		if !ok {
			eventsByRepositories[event.Repo.Name] = &EventsByRepo{
				PushEvents: make([]Payload, 0),
			}
			eventsByRepo = eventsByRepositories[event.Repo.Name]
		}
		if event.Type == "PushEvent" {
			eventsByRepo.PushEvents = append(eventsByRepo.PushEvents, event.Payload)
		}
		if event.Type == "IssueCommentEvent" {
			eventsByRepo.Issues += 1
		}
	}
	return eventsByRepositories, nil
}
