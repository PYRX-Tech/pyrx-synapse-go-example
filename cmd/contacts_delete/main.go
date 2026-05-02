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

	extID := fmt.Sprintf("del_test_%d", time.Now().Unix())

	// Create contact first
	_, err = client.Identify(synapse.IdentifyParams{
		ExternalID: extID,
		Email:      extID + "@test.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error identifying: %v\n", err)
		os.Exit(1)
	}

	err = client.Contacts.Delete(extID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Contact deleted")
}
