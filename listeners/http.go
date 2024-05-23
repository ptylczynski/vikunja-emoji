package listeners

import (
	"fmt"
	"net/http"
	"vikunja-emoji/vikunja"
)

func Register() {
	http.HandleFunc("/task-events/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("New Event detected!")
		vikunja.AddEmoji(request)
		fmt.Fprintf(writer, "OK")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Port error " + err.Error())
		return
	}
}
