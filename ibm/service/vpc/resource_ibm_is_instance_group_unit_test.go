// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestInstanceGroupLifecycleReasons(t *testing.T) {
	tests := []struct {
		name     string
		reasons  []vpcv1.InstanceGroupLifecycleReason
		expected string
	}{
		{
			name:     "no reasons reported",
			expected: "no lifecycle reasons reported",
		},
		{
			name: "single reason",
			reasons: []vpcv1.InstanceGroupLifecycleReason{
				{
					Code:    core.StringPtr("resource_suspended_by_provider"),
					Message: core.StringPtr("The resource has been suspended. Contact IBM support with the CRN for next steps."),
				},
			},
			expected: "resource_suspended_by_provider: The resource has been suspended. Contact IBM support with the CRN for next steps.",
		},
		{
			name: "multiple reasons are joined",
			reasons: []vpcv1.InstanceGroupLifecycleReason{
				{
					Code:    core.StringPtr("internal_error"),
					Message: core.StringPtr("Internal error. Contact IBM support."),
				},
				{
					Code:    core.StringPtr("resource_suspended_by_provider"),
					Message: core.StringPtr("The resource has been suspended."),
				},
			},
			expected: "internal_error: Internal error. Contact IBM support.; resource_suspended_by_provider: The resource has been suspended.",
		},
		{
			// A reason missing either half is skipped, so a list of nothing but
			// incomplete reasons must still produce a non-empty explanation.
			name: "reasons missing a code or a message are skipped",
			reasons: []vpcv1.InstanceGroupLifecycleReason{
				{Message: core.StringPtr("no code")},
				{Code: core.StringPtr("no_message")},
			},
			expected: "no lifecycle reasons reported",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			instanceGroup := &vpcv1.InstanceGroup{LifecycleReasons: tc.reasons}
			assert.Equal(t, tc.expected, instanceGroupLifecycleReasons(instanceGroup))
		})
	}
}
