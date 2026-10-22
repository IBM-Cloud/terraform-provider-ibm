// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import "testing"

func TestJSONEqual(t *testing.T) {
	if !jsonEqual(map[string]any{"a": 1, "b": []any{true}}, map[string]any{"b": []any{true}, "a": 1.0}) {
		t.Error("equal values with different numeric types should match")
	}
	if jsonEqual(map[string]any{"enable_cspm": true}, map[string]any{"enable_cspm": false}) {
		t.Error("different values must not match")
	}
}
