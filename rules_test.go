package authsignal

import (
	"encoding/json"
	"testing"
)

// The API soft-deletes rules, so getting a deleted rule still succeeds. DeletedAt is how a caller
// tells it apart from a live one.
func TestRuleResponseReadsDeletedAt(t *testing.T) {
	var deleted RuleResponse

	err := json.Unmarshal([]byte(`{"ruleId":"abc","deletedAt":"2026-10-13T00:00:00.000Z"}`), &deleted)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if deleted.DeletedAt == nil || *deleted.DeletedAt != "2026-10-13T00:00:00.000Z" {
		t.Fatalf("expected deletedAt to be set")
	}

	var live RuleResponse

	err = json.Unmarshal([]byte(`{"ruleId":"abc"}`), &live)
	if err != nil {
		t.Fatalf("failed to unmarshal json")
	}

	if live.DeletedAt != nil {
		t.Fatalf("expected deletedAt to be nil for a live rule")
	}
}
