// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestIsMoreThan24Hours(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	tests := []struct {
		name     string
		duration time.Duration
		expected bool
	}{
		{"Exactly 24 hours", 24 * time.Hour, false},
		{"More than 24 hours", 25 * time.Hour, true},
		{"Less than 24 hours", 23 * time.Hour, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helper.isMoreThan24Hours(tt.duration)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateExpirationDatetime(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	testcases := []struct {
		name            string
		timeoutDuration time.Duration
		expected        string
	}{
		{"Exactly 24 hours", 24 * time.Hour, "2025-01-02T15:00:00Z"},
		{"More than 24 hours", 25 * time.Hour, "2025-01-02T15:00:00Z"}, // Should cap at 24h
		{"Less than 24 hours", 20 * time.Minute, "2025-01-01T15:20:00Z"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := helper.calculateExpirationDatetime(tc.timeoutDuration)

			assert.Equal(t, tc.expected, result)
		})
	}
}
