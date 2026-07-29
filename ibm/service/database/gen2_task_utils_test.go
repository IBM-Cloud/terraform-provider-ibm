// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"strings"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// ---------------------------------------------------------------------------
// gen2GetOperationDescription
// ---------------------------------------------------------------------------

func TestGen2GetOperationDescription_LastOperationDescription(t *testing.T) {
	desc := "Provisioning complete"
	instance := &rc.ResourceInstance{
		LastOperation: &rc.ResourceInstanceLastOperation{
			Description: core.StringPtr(desc),
			Type:        core.StringPtr("provision"),
		},
	}
	got := gen2GetOperationDescription(instance)
	if got != desc {
		t.Errorf("expected %q, got %q", desc, got)
	}
}

func TestGen2GetOperationDescription_LastOperationTypeFallback(t *testing.T) {
	instance := &rc.ResourceInstance{
		LastOperation: &rc.ResourceInstanceLastOperation{
			Description: core.StringPtr(""),
			Type:        core.StringPtr("update"),
		},
	}
	got := gen2GetOperationDescription(instance)
	want := "Operation: update"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGen2GetOperationDescription_StateFallback(t *testing.T) {
	instance := &rc.ResourceInstance{
		State: core.StringPtr("active"),
	}
	got := gen2GetOperationDescription(instance)
	want := "Instance state: active"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGen2GetOperationDescription_DefaultFallback(t *testing.T) {
	instance := &rc.ResourceInstance{}
	got := gen2GetOperationDescription(instance)
	want := "Gen2 database instance operation"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

// ---------------------------------------------------------------------------
// gen2MapStateToStatus
// ---------------------------------------------------------------------------

func TestGen2MapStateToStatus(t *testing.T) {
	tests := []struct {
		state *string
		want  string
	}{
		{nil, "unknown"},
		{core.StringPtr("active"), "completed"},
		{core.StringPtr("provisioning"), "running"},
		{core.StringPtr("in progress"), "running"},
		{core.StringPtr("removed"), "completed"},
		{core.StringPtr("failed"), "failed"},
		{core.StringPtr("inactive"), "inactive"},
		{core.StringPtr("unknown-state"), "unknown-state"},
	}

	for _, tc := range tests {
		label := "nil"
		if tc.state != nil {
			label = *tc.state
		}
		t.Run(label, func(t *testing.T) {
			instance := &rc.ResourceInstance{State: tc.state}
			got := gen2MapStateToStatus(instance)
			if got != tc.want {
				t.Errorf("state=%q: expected %q, got %q", label, tc.want, got)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// gen2CalculateProgress
// ---------------------------------------------------------------------------

func TestGen2CalculateProgress(t *testing.T) {
	tests := []struct {
		state *string
		want  int
	}{
		{nil, 0},
		{core.StringPtr("active"), 100},
		{core.StringPtr("provisioning"), 50},
		{core.StringPtr("in progress"), 75},
		{core.StringPtr("failed"), 100},
		{core.StringPtr("removed"), 100},
		{core.StringPtr("inactive"), 0},
		{core.StringPtr("unknown-state"), 0},
	}

	for _, tc := range tests {
		label := "nil"
		if tc.state != nil {
			label = *tc.state
		}
		t.Run(label, func(t *testing.T) {
			instance := &rc.ResourceInstance{State: tc.state}
			got := gen2CalculateProgress(instance)
			if got != tc.want {
				t.Errorf("state=%q: expected %d, got %d", label, tc.want, got)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// gen2GetOperationTime
// ---------------------------------------------------------------------------

func TestGen2GetOperationTime_UpdatedAt(t *testing.T) {
	updatedAt := strfmt.DateTime(time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC))
	createdAt := strfmt.DateTime(time.Date(2024, 6, 1, 11, 0, 0, 0, time.UTC))
	instance := &rc.ResourceInstance{
		UpdatedAt: &updatedAt,
		CreatedAt: &createdAt,
	}
	got := gen2GetOperationTime(instance)
	// UpdatedAt should be preferred over CreatedAt
	if !strings.Contains(got, "2024") {
		t.Errorf("unexpected time string: %q", got)
	}
}

func TestGen2GetOperationTime_CreatedAtFallback(t *testing.T) {
	createdAt := strfmt.DateTime(time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC))
	instance := &rc.ResourceInstance{
		CreatedAt: &createdAt,
	}
	got := gen2GetOperationTime(instance)
	if !strings.Contains(got, "2024") {
		t.Errorf("unexpected time string: %q", got)
	}
}

func TestGen2GetOperationTime_NowFallback(t *testing.T) {
	before := time.Now().UTC().Add(-2 * time.Second)
	instance := &rc.ResourceInstance{}
	got := gen2GetOperationTime(instance)
	after := time.Now().UTC().Add(2 * time.Second)

	parsed, err := time.Parse(time.RFC3339, got)
	if err != nil {
		t.Fatalf("returned string is not RFC3339: %q", got)
	}
	if parsed.Before(before) || parsed.After(after) {
		t.Errorf("expected time near now, got %q", got)
	}
}
