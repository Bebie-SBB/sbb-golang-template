package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := os.Getenv("HEALTH_URL")
	if url == "" {
		url = "http://localhost:9090/health"
	}

	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Health check failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println("Health check succeeded")
		os.Exit(0)
	}
	fmt.Println("Health check failed, status:", resp.StatusCode)
	os.Exit(1)
}
