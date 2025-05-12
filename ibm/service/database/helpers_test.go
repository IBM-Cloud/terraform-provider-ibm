// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
)

func TestIsMoreThan24Hours(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	testcases := []struct {
		description string
		duration    time.Duration
		expected    bool
	}{
		{
			description: "When duration is EXACTLY 24 hours, Expect false",
			duration:    24 * time.Hour,
			expected:    false,
		},
		{
			description: "When duration is MORE than 24 hours, Expect true",
			duration:    25 * time.Hour,
			expected:    true,
		},
		{
			description: "When duration is LESS than 24 hours, Expect false",
			duration:    23 * time.Hour,
			expected:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := helper.isMoreThan24Hours(tc.duration)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestFutureTimeToISO(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	result := helper.futureTimeToISO(30 * time.Minute)
	expected := strfmt.DateTime(result)

	require.Equal(t, expected, result)
}

func TestCalculateExpirationDatetime(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	expected24Hours := strfmt.DateTime(helper.futureTimeToISO(24 * time.Hour))
	expected20minutes := strfmt.DateTime(helper.futureTimeToISO(20 * time.Minute))

	testcases := []struct {
		description string
		duration    time.Duration
		expected    strfmt.DateTime
	}{
		{
			description: "When duration is EXACTLY 24 hours, Expect an ISO 24 hrs from now",
			duration:    24 * time.Hour,
			expected:    expected24Hours,
		},
		{
			description: "When duration is MORE than 24 hours, Expect an ISO 24 hrs from now as that is the maximum",
			duration:    25 * time.Hour,
			expected:    expected24Hours,
		},
		{
			description: "When duration is LESS than 24 hours, Expect an ISO of now + duration",
			duration:    20 * time.Minute,
			expected:    expected20minutes,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := helper.calculateExpirationDatetime(tc.duration)
			require.Equal(t, tc.expected, result)
		})
	}
}

type MockTaskClient struct {
	Tasks []clouddatabasesv5.Task
	Err   error
}

func (m *MockTaskClient) ListDeploymentTasks(opts *clouddatabasesv5.ListDeploymentTasksOptions) (*clouddatabasesv5.Tasks, *core.DetailedResponse, error) {
	if m.Err != nil {
		return nil, nil, m.Err
	}
	return &clouddatabasesv5.Tasks{
		Tasks: m.Tasks,
	}, &core.DetailedResponse{}, nil
}

func TestMatchingTaskInProgress(t *testing.T) {
	str := "2025-05-12T10:00:00Z"
	parsedTime, _ := time.Parse(time.RFC3339, str)
	mockCreatedAt := strfmt.DateTime(parsedTime)

	testcases := []struct {
		description        string
		mockTasks          []clouddatabasesv5.Task
		mockError          error
		instanceID         string
		matchResourceType  string
		expectedInProgress bool
		expectedTask       clouddatabasesv5.Task
		expectError        bool
	}{
		{
			description: "When matching task is running, Expect true and matching task",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("123"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
				{
					ID:              core.StringPtr("1234"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
			},
			instanceID:         "inst-1",
			matchResourceType:  taskUpgrade,
			expectedInProgress: true,
			expectedTask: clouddatabasesv5.Task{
				ID:              core.StringPtr("1234"),
				Status:          core.StringPtr(databaseTaskRunningStatus),
				ResourceType:    core.StringPtr(taskUpgrade),
				CreatedAt:       &mockCreatedAt,
				ProgressPercent: core.Int64Ptr(74),
				Description:     core.StringPtr("Upgrade running"),
			},
		},
		{
			description: "When matching task is queued, Expect true and matching task",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("123"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
				{
					ID:              core.StringPtr("234"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
			},
			instanceID:         "inst-2",
			matchResourceType:  taskUpgrade,
			expectedInProgress: true,
			expectedTask: clouddatabasesv5.Task{
				ID:              core.StringPtr("234"),
				Status:          core.StringPtr(databaseTaskQueuedStatus),
				ResourceType:    core.StringPtr(taskUpgrade),
				CreatedAt:       &mockCreatedAt,
				ProgressPercent: core.Int64Ptr(74),
				Description:     core.StringPtr("Upgrade running"),
			},
		},
		{
			description: "When matching task is completed, Expect false",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("101"),
					Status:          core.StringPtr(databaseTaskCompletedStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
				{
					ID:              core.StringPtr("102"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr("backup"),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("backup running"),
				},
			},
			instanceID:         "inst-4",
			matchResourceType:  taskUpgrade,
			expectedInProgress: false,
		},
		{
			description: "When matching task is NOT the running task, Expect false",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("789"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
			},
			instanceID:         "inst-3",
			matchResourceType:  taskUpgrade,
			expectedInProgress: false,
		},
		{
			description:        "When there is an error getting tasks, Expect error",
			mockError:          fmt.Errorf("API error"),
			instanceID:         "inst-5",
			matchResourceType:  taskUpgrade,
			expectError:        true,
			expectedInProgress: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			mockClient := &MockTaskClient{
				Tasks: tc.mockTasks,
				Err:   tc.mockError,
			}

			tm := &TaskManager{
				Client:     mockClient,
				InstanceID: tc.instanceID,
			}

			inProgress, task, err := tm.matchingTaskInProgress(tc.matchResourceType)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedInProgress, inProgress)

				if tc.expectedInProgress {
					require.NotNil(t, task)
					require.Equal(t, tc.expectedTask, *task)
				} else {
					require.Nil(t, task)
				}
			}
		})
	}
}
