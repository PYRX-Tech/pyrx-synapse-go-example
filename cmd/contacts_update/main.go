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

	// Identify first to ensure contact exists
	_, err = client.Identify(synapse.IdentifyParams{
		ExternalID: "sdk_update_test",
		Email:      "update@example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error identifying: %v\n", err)
		os.Exit(1)
	}

	updated, err := client.Contacts.Update("sdk_update_test", synapse.ContactUpdateParams{
		Properties: map[string]interface{}{"plan": "growth"},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Updated: %+v\n", updated)
}
