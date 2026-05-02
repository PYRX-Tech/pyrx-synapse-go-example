package main

import (
	"fmt"
	"os"
	"time"

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

	slug := fmt.Sprintf("sdk-get-test-%d", time.Now().Unix())

	// Create a template first
	_, err = client.Templates.Create(synapse.TemplateCreateParams{
		Slug:       slug,
		Name:       "Get Test",
		Subject:    "Test",
		BodyHTML:   "<p>Hi</p>",
		SenderName: "Test",
		FromEmail:  "test@example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating: %v\n", err)
		os.Exit(1)
	}

	template, err := client.Templates.Get(slug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Template: %+v\n", template)

	// Cleanup
	_ = client.Templates.Delete(slug)
}
