package api

import (
	"strings"
	"testing"
)

func TestMediaEntryUpdateMutation_OmitsUnsetOptionalFields(t *testing.T) {
	query := mediaEntryUpdateMutation(map[string]any{
		"id":       1,
		"progress": 5,
	})

	if !strings.Contains(query, "progress: $progress") {
		t.Fatalf("expected progress arg in mutation, got: %s", query)
	}
	if strings.Contains(query, "startedAt:") {
		t.Fatalf("did not expect startedAt in mutation when start date is not provided")
	}
	if strings.Contains(query, "completedAt:") {
		t.Fatalf("did not expect completedAt in mutation when completed date is not provided")
	}
	if strings.Contains(query, "notes: $notes") {
		t.Fatalf("did not expect notes arg in mutation when notes are not provided")
	}
	if strings.Contains(query, "status: $status") {
		t.Fatalf("did not expect status arg in mutation when status is not provided")
	}
}

func TestMediaEntryUpdateMutation_IncludesProvidedDates(t *testing.T) {
	query := mediaEntryUpdateMutation(map[string]any{
		"id":     1,
		"cDate":  1,
		"cMonth": 2,
		"cYear":  2026,
		"sDate":  3,
		"sMonth": 4,
		"sYear":  2025,
	})

	if !strings.Contains(query, "completedAt: { day: $cDate, month: $cMonth, year: $cYear }") {
		t.Fatalf("expected completedAt in mutation, got: %s", query)
	}
	if !strings.Contains(query, "startedAt: { day: $sDate, month: $sMonth, year: $sYear }") {
		t.Fatalf("expected startedAt in mutation, got: %s", query)
	}
}
