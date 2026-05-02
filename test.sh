#!/usr/bin/env bash
set -uo pipefail
RED='\033[0;31m'; GREEN='\033[0;32m'; NC='\033[0m'
PASS=0; FAIL=0

[[ -z "${SYNAPSE_API_KEY:-}" ]] && echo "Set SYNAPSE_API_KEY" && exit 1

cd "$(dirname "$0")"

run_script() {
  local name="$1" cmd="$2"
  if eval "$cmd" > /dev/null 2>&1; then
    echo -e "  ${GREEN}✓${NC} $name"
    ((PASS++)) || true
  else
    echo -e "  ${RED}✗${NC} $name"
    ((FAIL++)) || true
  fi
}

echo "Running all scripts..."

echo "── Core ──"
run_script "track_event" "go run ./cmd/track_event"
run_script "track_batch" "go run ./cmd/track_batch"
run_script "identify_contact" "go run ./cmd/identify_contact"
run_script "identify_batch" "go run ./cmd/identify_batch"
run_script "send_email" "go run ./cmd/send_email"

echo "── Contacts ──"
run_script "contacts_list" "go run ./cmd/contacts_list"
run_script "contacts_get" "go run ./cmd/contacts_get"
run_script "contacts_update" "go run ./cmd/contacts_update"
run_script "contacts_delete" "go run ./cmd/contacts_delete"

echo "── Templates ──"
run_script "templates_list" "go run ./cmd/templates_list"
run_script "templates_get" "go run ./cmd/templates_get"
run_script "templates_create" "go run ./cmd/templates_create"
run_script "templates_update" "go run ./cmd/templates_update"
run_script "templates_preview" "go run ./cmd/templates_preview"
run_script "templates_delete" "go run ./cmd/templates_delete"

echo ""
echo "Results: $PASS passed, $FAIL failed"
[[ $FAIL -eq 0 ]] && echo -e "${GREEN}All passed!${NC}" || echo -e "${RED}$FAIL failed${NC}"
exit $FAIL
