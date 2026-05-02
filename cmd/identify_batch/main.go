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

	resp, err := client.IdentifyBatch(synapse.IdentifyBatchParams{
		Contacts: []synapse.IdentifyParams{
			{ExternalID: "user_1", Email: "alice@example.com", Properties: map[string]interface{}{"plan": "starter"}},
			{ExternalID: "user_2", Email: "bob@example.com", Properties: map[string]interface{}{"plan": "growth"}},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Batch identified: %+v\n", resp)
}
