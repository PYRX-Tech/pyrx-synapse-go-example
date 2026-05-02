# Synapse Go Example

All 15 SDK endpoints with [pyrx-synapse-go](https://github.com/pyrx-tech/pyrx-synapse-go).

## Setup

1. Ensure Go 1.21+ is installed
2. Copy `.env.example` to `.env` and fill in your credentials
3. Export environment variables: `export $(cat .env | xargs)`

## Examples

### Core
```bash
go run ./cmd/track_event         # Track event
go run ./cmd/track_batch         # Batch track
go run ./cmd/identify_contact    # Identify contact
go run ./cmd/identify_batch      # Batch identify
go run ./cmd/send_email          # Send email
```

### Contacts
```bash
go run ./cmd/contacts_list       go run ./cmd/contacts_get
go run ./cmd/contacts_update     go run ./cmd/contacts_delete
```

### Templates
```bash
go run ./cmd/templates_list      go run ./cmd/templates_get
go run ./cmd/templates_create    go run ./cmd/templates_update
go run ./cmd/templates_delete    go run ./cmd/templates_preview
```

- [Synapse Docs](https://synapse.pyrx.tech/developers)
- [Go SDK](https://synapse.pyrx.tech/developers/sdks/go)
