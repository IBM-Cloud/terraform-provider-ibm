// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/core"
	"gotest.tools/assert"
)

type MockMeta struct{}

func MockListDeploymentTaskRunning(instanceID string, meta interface{}) (*clouddatabasesv5.Tasks, error) {
	return &clouddatabasesv5.Tasks{
		Tasks: []clouddatabasesv5.Task{
			{
				Status:      core.StringPtr("running"),
				Description: core.StringPtr("Upgrading database"),
			},
		},
	}, nil
}

func MockListDeploymentTaskCompleted(instanceID string, meta interface{}) (*clouddatabasesv5.Tasks, error) {
	return &clouddatabasesv5.Tasks{
		Tasks: []clouddatabasesv5.Task{
			{
				Status:      core.StringPtr("completed"),
				Description: core.StringPtr("Upgrading database"),
			},
		},
	}, nil
}

func MockGetDeploymentCapability(capability, instanceID, deploymentType, region string, meta interface{}) (*clouddatabasesv5.Capability, error) {
	return &clouddatabasesv5.Capability{
		Versions: []clouddatabasesv5.VersionsCapabilityItem{
			{
				Version: core.StringPtr("5"),
				Status:  core.StringPtr("deprecated"),
				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{
					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("6")},
					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("7")},
					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("8")},
				},
			},
		},
	}, nil
}

func MockGetDeploymentCapabilityNoTransitions(capability, instanceID, deploymentType, region string, meta interface{}) (*clouddatabasesv5.Capability, error) {
	return &clouddatabasesv5.Capability{
		Versions: []clouddatabasesv5.VersionsCapabilityItem{
			{
				Version:     core.StringPtr("8"),
				Status:      core.StringPtr("preview"),
				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{},
			},
		},
	}, nil
}

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		name               string
		instanceID         string
		newVersion         string
		skipBackup         bool
		mockListFunc       func(instanceID string, meta interface{}) (*clouddatabasesv5.Tasks, error)
		mockCapabilityFunc func(capability string, instanceID string, deploymentType string, region string, meta interface{}) (*clouddatabasesv5.Capability, error)
		expectedError      string
	}{
		{
			name:               "Upgrade task already running",
			instanceID:         "test-instance",
			newVersion:         "6",
			skipBackup:         false,
			mockListFunc:       MockListDeploymentTaskRunning,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "There is already an upgrade task running. Please wait for this to complete",
		},
		{
			name:               "No upgrade paths",
			instanceID:         "test-instance",
			newVersion:         "9",
			skipBackup:         false,
			mockListFunc:       MockListDeploymentTaskCompleted,
			mockCapabilityFunc: MockGetDeploymentCapabilityNoTransitions,
			expectedError:      "You are not allowed to upgrade version, there are no approved upgrade paths for your current version, please look at our docs here",
		},
		{
			name:               "Upgrade version not allowed",
			instanceID:         "test-instance",
			newVersion:         "10",
			skipBackup:         false,
			mockListFunc:       MockListDeploymentTaskCompleted,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "Version 10 is not a valid upgrade version. Allowed versions [6 7 8]",
		},
		{
			name:               "Valid version, no error",
			instanceID:         "test-instance",
			newVersion:         "7",
			skipBackup:         false,
			mockListFunc:       MockListDeploymentTaskCompleted,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "",
		},
	}

	for _, tc := range tests {
		listDeploymentTasksFunc = tc.mockListFunc
		getDeploymentCapabilityFunc = tc.mockCapabilityFunc

		err := validateVersion(tc.instanceID, tc.newVersion, tc.skipBackup, &MockMeta{})

		// TODO fix actual versus expected
		// Log the actual result for debugging
		if err != nil {
			t.Logf("Test %s: received error: %v", tc.name, err)
		} else {
			t.Logf("Test %s: received no error", tc.name)
		}
		// TODO check nil error message
		if tc.expectedError != "" {
			assert.Error(t, err, tc.expectedError)
		} else {
			assert.NilError(t, err)
		}
	}
}
