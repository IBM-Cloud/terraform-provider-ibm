// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestIsVersionUpgradeAllowed(t *testing.T) {
	testcases := []struct {
		description    string
		version        Version
		oldVersion     string
		upgradeVersion string
		expectedResult bool
	}{
		{
			description: "When transition for upgrade version is in-place, Expect true",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: inPlace, ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(false)},
					{Method: inPlace, ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			upgradeVersion: "7.0",
			oldVersion:     "6.0",
			expectedResult: true,
		},
		{
			description: "When transition for upgrade version is restore, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: restore, ToVersion: "7.1"},
					{Method: restore, ToVersion: "7.2"},
				},
			},
			upgradeVersion: "7.1",
			oldVersion:     "7.0",
			expectedResult: false,
		},
		{
			description: "When transition does not exist for upgrade version, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: restore, ToVersion: "7.1"},
					{Method: inPlace, ToVersion: "7.2"},
				},
			},
			oldVersion:     "7.0",
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When no transitions exist, Expect false",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
			oldVersion:     "14",
			upgradeVersion: "7.2",
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.version.isVersionUpgradeAllowed(tc.upgradeVersion)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestIsSkipBackupUpgradeAllowed(t *testing.T) {
	testcases := []struct {
		description    string
		version        Version
		oldVersion     string
		upgradeVersion string
		expectedResult bool
	}{
		{
			description: "When an in-place transition for upgrade version allows skip backup, expect true.",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: inPlace, ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: inPlace, ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			oldVersion:     "6.0",
			upgradeVersion: "6.1",
			expectedResult: true,
		},
		{
			description: "When an in-place transition for upgrade version DOES NOT allow skip backup, expect false.",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: inPlace, ToVersion: "8", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			oldVersion:     "7.0",
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When an in-place transition for upgrade version exists but skip backup is nil, expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: inPlace, ToVersion: "8", SkipBackupSupported: nil},
				},
			},
			oldVersion:     "7.0",
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When no transitions exist for upgrade version, expect false.",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
			oldVersion:     "14",
			upgradeVersion: "15",
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.version.isSkipBackupUpgradeAllowed(tc.upgradeVersion)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestHasUpgradeVersions(t *testing.T) {
	testcases := []struct {
		description    string
		version        Version
		expectedResult bool
	}{
		{
			description: "When in-place transitions exist, expect true.",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: inPlace, ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: inPlace, ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: true,
		},
		{
			description: "When only restore transitions exist, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: restore, ToVersion: "8", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: false,
		},
		{
			description: "When no transitions exist, expect false.",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.version.hasUpgradeVersions()
			assert.Equal(t, tc.expectedResult, result)
		})

	}
}

func TestGetAllowedVersionsList(t *testing.T) {
	testcases := []struct {
		description    string
		version        Version
		expectedResult []string
	}{
		{
			description: "When in-place transitions exist, Expect list of in-place versions",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: restore, ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: restore, ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
					{Method: inPlace, ToVersion: "8.0", SkipBackupSupported: core.BoolPtr(false)},
					{Method: inPlace, ToVersion: "9.5", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: []string{"8.0", "9.5"},
		},
		{
			description: "When only restore transitions exist, Expect nil",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: restore, ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: restore, ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: nil,
		},
		{
			description: "When no transitions exist, Expect nil",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
			expectedResult: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.version.getAllowedVersionsList()
			assert.DeepEqual(t, tc.expectedResult, result)
		})
	}
}

type MockMeta struct{}

func MockGetDeploymentCapability(capability string, instanceID string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error) {
	return &clouddatabasesv5.Capability{
		Versions: []clouddatabasesv5.VersionsCapabilityItem{
			{
				Version:     core.StringPtr("6"),
				Status:      core.StringPtr("active"),
				IsPreferred: core.BoolPtr(false),
				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{
					{
						Application:         core.StringPtr("mongodb"),
						FromVersion:         core.StringPtr("6"),
						ToVersion:           core.StringPtr("7"),
						Method:              core.StringPtr(inPlace),
						SkipBackupSupported: core.BoolPtr(false),
					},
					{
						Application:         core.StringPtr("mongodb"),
						FromVersion:         core.StringPtr("6"),
						ToVersion:           core.StringPtr("8"),
						Method:              core.StringPtr(inPlace),
						SkipBackupSupported: core.BoolPtr(true),
					},
					{
						Application:         core.StringPtr("mongodb"),
						FromVersion:         core.StringPtr("6"),
						ToVersion:           core.StringPtr("10"),
						Method:              core.StringPtr(restore),
						SkipBackupSupported: core.BoolPtr(true),
					},
				},
			},
		},
	}, nil
}

func MockGetDeploymentCapabilityNoTransitions(capability string, instanceID string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error) {
	return &clouddatabasesv5.Capability{
		Versions: []clouddatabasesv5.VersionsCapabilityItem{
			{
				Type:        core.StringPtr("mongodb"),
				Version:     core.StringPtr("9"),
				Status:      core.StringPtr("active"),
				IsPreferred: core.BoolPtr(false),
				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{},
			},
		},
	}, nil
}

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		description        string
		instanceID         string
		location           string
		oldVersion         string
		upgradeVersion     string
		skipBackup         bool
		mockCapabilityFunc func(capability string, instanceID string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error)
		expectedError      string
	}{
		{
			description:        "When there are no upgrade paths for the upgrade version, Expect no upgrade versions error",
			instanceID:         "test-instance",
			location:           "us-south",
			oldVersion:         "8",
			upgradeVersion:     "9",
			skipBackup:         false,
			mockCapabilityFunc: MockGetDeploymentCapabilityNoTransitions,
			expectedError:      "No available upgrade versions for version 8",
		},
		{
			description:        "When the upgrade version is no a valid version, Expect allowed versions error",
			instanceID:         "test-instance",
			location:           "us-south",
			oldVersion:         "9",
			upgradeVersion:     "10",
			skipBackup:         false,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "Version 10 is not a valid upgrade version. Allowed versions: [7 8]",
		},
		{
			description:        "When skip backup is not allowed for upgrade version, Expect skip backup error",
			instanceID:         "test-instance",
			location:           "eu-gb",
			oldVersion:         "6",
			upgradeVersion:     "7",
			skipBackup:         true,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "Skipping backup is not allowed when upgrading to version 7",
		},
		{
			description:        "When the upgrade version is valid, Expect no error",
			instanceID:         "test-instance",
			location:           "eu-gb",
			oldVersion:         "6",
			upgradeVersion:     "7",
			skipBackup:         false,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "",
		},
	}

	originalFetchFunc := fetchDeploymentVersionFn
	defer func() { fetchDeploymentVersionFn = originalFetchFunc }()

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			fetchDeploymentVersionFn = func(instanceID string, location string, meta interface{}) *Version {
				capability, err := tc.mockCapabilityFunc(versions, instanceID, classicPlatform, location, meta)
				require.NoError(t, err)
				return expandVersion(capability.Versions[0])
			}

			err := validateUpgradeVersion(tc.instanceID, tc.location, tc.oldVersion, tc.upgradeVersion, tc.skipBackup, &MockMeta{})

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateGroupScaling(t *testing.T) {
	tests := []struct {
		description   string
		groupId       string
		resourceName  string
		value         int
		resource      *GroupResource
		nodeCount     int
		expectedError string
	}{
		{
			description:  "When group is NOT adjustable, Expect error",
			groupId:      "member",
			resourceName: "members",
			value:        6,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   3,
				Minimum:                      3,
				Maximum:                      9,
				StepSize:                     1,
				IsAdjustable:                 false,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "member can not change members value after create",
		},
		{
			description:  "When new value is invalid step size, Expect error specifying increments",
			groupId:      "member",
			resourceName: "members",
			value:        5,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   3,
				Minimum:                      3,
				Maximum:                      9,
				StepSize:                     3,
				IsAdjustable:                 true,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "member group members must be >= 3 and <= 9 in increments of 3",
		},
		{
			description:  "When new value is less than the min, Expect error specifying increments",
			groupId:      "member",
			resourceName: "members",
			value:        2,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   3,
				Minimum:                      3,
				Maximum:                      9,
				StepSize:                     1,
				IsAdjustable:                 true,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "member group members must be >= 3 and <= 9 in increments of 1",
		},
		{
			description:  "When new value is more than the max, Expect error specifying increments",
			groupId:      "member",
			resourceName: "members",
			value:        10,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   4,
				Minimum:                      2,
				Maximum:                      8,
				StepSize:                     2,
				IsAdjustable:                 true,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "member group members must be >= 2 and <= 8 in increments of 2",
		},
		{
			description:  "When new value is less than current value, Expect error cannot scale down increments",
			groupId:      "member",
			resourceName: "members",
			value:        4,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   5,
				Minimum:                      3,
				Maximum:                      9,
				StepSize:                     1,
				IsAdjustable:                 true,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "can not scale member group members below 5 to 4",
		},
		{
			description:  "When new value is a valid increment, Expect no error",
			groupId:      "member",
			resourceName: "members",
			value:        6,
			resource: &GroupResource{
				Units:                        "count",
				Allocation:                   3,
				Minimum:                      3,
				Maximum:                      9,
				StepSize:                     3,
				IsAdjustable:                 true,
				IsOptional:                   false,
				CanScaleDown:                 false,
				CPUEnforcementRatioCeilingMb: 0,
				CPUEnforcementRatioMb:        0,
			},
			nodeCount:     1,
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			err := validateGroupScaling(tc.groupId, tc.resourceName, tc.value, tc.resource, tc.nodeCount)

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
