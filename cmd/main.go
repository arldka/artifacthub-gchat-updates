package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/arthur-laurentdka/artifacthub-gchat-updates/internal/chat"
)

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	http.HandleFunc("/", chat.NotificationHandler)
	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
