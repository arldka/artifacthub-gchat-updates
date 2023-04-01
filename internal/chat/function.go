package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"os"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type ah_repository struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Publisher string `json:"publisher"`
}

type ah_package struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Url string `json:"url"`
	Changes []string `json:"changes"`
	Repository ah_repository `json:"repository"`
}

type ah_payload struct {
	Pkg ah_package `json:"package"`
}

func sendMessage(cards gc_cards) {
	var body, err = json.Marshal(cards)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sending message to: %s\n", os.Getenv("WEBHOOK_URL"))
	resp, err := http.Post(os.Getenv("WEBHOOK_URL"),
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Response: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", body)
}

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received notification: %s\n", r.URL.Path)
	event :=  cloudevents.NewEvent()
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		_ = fmt.Errorf("request is not in the cloud events format: %s", err)
		return
	}
	data := &ah_payload{}
	err = event.DataAs(data)
	if err != nil {
		_ = fmt.Errorf("failed to unmarshal data: %s", err)
		return
	}
	fmt.Printf("Received package: %s\n", data.Pkg.Name)
	sendMessage(gcMessageGenerator(data))
	return
}
