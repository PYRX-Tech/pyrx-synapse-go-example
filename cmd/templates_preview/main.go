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

	slug := fmt.Sprintf("sdk-preview-test-%d", time.Now().Unix())

	// Create a template to preview
	_, err = client.Templates.Create(synapse.TemplateCreateParams{
		Slug:       slug,
		Name:       "Preview Test",
		Subject:    "Hi {{first_name}}",
		BodyHTML:   "<p>Hello {{first_name}}</p>",
		SenderName: "Test",
		FromEmail:  "test@example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating: %v\n", err)
		os.Exit(1)
	}

	preview, err := client.Templates.Preview(slug, synapse.TemplatePreviewParams{
		Contact: map[string]interface{}{"email": "jane@example.com", "first_name": "Jane"},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Preview: %+v\n", preview)

	// Cleanup
	_ = client.Templates.Delete(slug)
}
