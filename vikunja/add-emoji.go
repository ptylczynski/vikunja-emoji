package vikunja

import (
	"encoding/json"
	"fmt"
	"github.com/forPelevin/gomoji"
	"net/http"
	"vikunja-emoji/gpt"
)

func AddEmoji(req *http.Request) {
	task := getTask(req)
	err := Validate(task)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = getNewTitle(task)
	if err != nil {
		fmt.Println(err)
		return
	}
	PostNewTitle(task)
}

func getTask(req *http.Request) *Task {
	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		fmt.Println("Request body cannot be parsed, because of " + err.Error())
		return nil
	}
	title := request.Data.Task.Title
	fmt.Println("Received Title is " + title)
	return &request.Data.Task
}

func getNewTitle(task *Task) error {
	title := removeOldEmoji(task.Title)
	emoji := gpt.GetEmoji(title)
	if emoji != "" {
		title = emoji + " " + title
	} else {
		fmt.Println("No emoji found for " + title)
		return fmt.Errorf("No emoji found for " + title)
	}
	task.Title = title
	return nil
}

func removeOldEmoji(s string) string {
	return gomoji.RemoveEmojis(s)
}
