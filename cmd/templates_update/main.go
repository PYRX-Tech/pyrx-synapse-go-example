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

	slug := fmt.Sprintf("tpl-update-%d", time.Now().Unix())

	// Create first
	_, err = client.Templates.Create(synapse.TemplateCreateParams{
		Slug:       slug,
		Name:       "Update Test",
		Subject:    "Original subject",
		BodyHTML:   "<h1>Hi</h1>",
		SenderName: "Synapse",
		FromEmail:  "noreply@example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating: %v\n", err)
		os.Exit(1)
	}

	updated, err := client.Templates.Update(slug, synapse.TemplateUpdateParams{
		Subject: "Your order is confirmed!",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Updated: %+v\n", updated)

	// Cleanup
	_ = client.Templates.Delete(slug)
}
