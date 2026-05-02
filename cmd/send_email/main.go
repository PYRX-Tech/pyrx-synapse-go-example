package main

import (
	"fmt"
	"os"

	synapse "github.com/pyrx-tech/pyrx-synapse-go"
)

func main() {
	baseURL := os.Getenv("SYNAPSE_API_URL")
	if baseURL == "" {
		baseURL = "https://synapse-api.pyrx.tech"
	}
	client, err := synapse.NewClient(synapse.Config{
		APIKey:      os.Getenv("SYNAPSE_API_KEY"),
		WorkspaceID: os.Getenv("SYNAPSE_WORKSPACE_ID"),
		BaseURL:     baseURL,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Send(synapse.SendParams{
		TemplateSlug: "welcome-email",
		To:           map[string]interface{}{"user_id": "user_123", "email": "jane@example.com"},
		Attributes:   map[string]interface{}{"first_name": "Jane"},
	})
	if err != nil {
		fmt.Printf("Send failed (expected if template doesn't exist): %v\n", err)
		os.Exit(0)
	}
	fmt.Printf("Email sent: %+v\n", resp)
}
