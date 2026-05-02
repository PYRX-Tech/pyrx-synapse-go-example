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

	slug := fmt.Sprintf("tpl-create-%d", time.Now().Unix())

	template, err := client.Templates.Create(synapse.TemplateCreateParams{
		Slug:       slug,
		Name:       "Create Test",
		Subject:    "Order confirmed",
		BodyHTML:   "<h1>Hi</h1><p>Your order is confirmed.</p>",
		SenderName: "Synapse",
		FromEmail:  "noreply@example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created: %+v\n", template)

	// Cleanup
	_ = client.Templates.Delete(slug)
}
