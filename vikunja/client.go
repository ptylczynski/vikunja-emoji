package vikunja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Task struct {
	Id        int64
	Title     string
	ProjectId int64
}

type EventData struct {
	Task Task
}

// tk_9cc767b98806b95e0cdb3b207ab3873a42da24a9
type Request struct {
	EventName string
	Time      string
	Data      EventData
}

func PostNewTitle(task *Task) {
	fmt.Println("Sending updated Title " + task.Title + " to vikunja")
	client := &http.Client{}
	request, err := createRequest(task)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	setHeaders(request)
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func createRequest(task *Task) (*http.Request, error) {
	url := fmt.Sprintf("%s/api/v1/tasks/%d", os.Getenv("VIKUNJA_SERVICE_URL"), task.ProjectId)
	body := createBody(task.Title, task.Id)
	request, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		fmt.Println("Request cannot be created, because of " + err.Error())
		return nil, err
	}
	setHeaders(request)
	return request, nil
}

func setHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("VIKUNJA_SERVICE_TOKEN"))
}

func createBody(title string, id int64) []byte {
	newTask := Task{Title: title, Id: id}
	body, _ := json.Marshal(newTask)
	return body
}
