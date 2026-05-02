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

	// List contacts first to get an ID
	listResp, err := client.Contacts.List(synapse.ContactListParams{Page: 1, PerPage: 1})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing contacts: %v\n", err)
		os.Exit(1)
	}

	if listResp.Data == nil || len(listResp.Data) == 0 {
		fmt.Println("No contacts found")
		os.Exit(0)
	}

	contactID := listResp.Data[0].ID
	contact, err := client.Contacts.Get(contactID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Contact: %+v\n", contact)
}
