// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/core"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestIsVersionUpgradeAllowed(t *testing.T) {
	testcases := []struct {
		description    string
		version        Version
		upgradeVersion string
		expectedResult bool
	}{
		{
			description: "When transition for upgrade version is in-place, Expect true",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: "in-place", ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(false)},
					{Method: "in-place", ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			upgradeVersion: "7.0",
			expectedResult: true,
		},
		{
			description: "When transition for upgrade version is restore, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: "restore", ToVersion: "7.1"},
					{Method: "restore", ToVersion: "7.2"},
				},
			},
			upgradeVersion: "7.1",
			expectedResult: false,
		},
		{
			description: "When transition does not exist for upgrade version, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: "restore", ToVersion: "7.1"},
					{Method: "in-place", ToVersion: "7.2"},
				},
			},
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When no transitions exist, Expect false",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
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
		upgradeVersion string
		expectedResult bool
	}{
		{
			description: "When an in-place transition for upgrade version allows skip backup, expect true.",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: "in-place", ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: "in-place", ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			upgradeVersion: "6.1",
			expectedResult: true,
		},
		{
			description: "When an in-place transition for upgrade version DOES NOT allow skip backup, expect false.",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: "in-place", ToVersion: "8", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When an in-place transition for upgrade version exists but skip backup is nil, expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: "in-place", ToVersion: "8", SkipBackupSupported: nil},
				},
			},
			upgradeVersion: "8",
			expectedResult: false,
		},
		{
			description: "When no transitions exist for upgrade version, expect false.",
			version: Version{
				Version:     "14",
				Transitions: []VersionTransition{},
			},
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
					{Method: "in-place", ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: "in-place", ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: true,
		},
		{
			description: "When only restore transitions exist, Expect false",
			version: Version{
				Version: "7.0",
				Transitions: []VersionTransition{
					{Method: "restore", ToVersion: "8", SkipBackupSupported: core.BoolPtr(false)},
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
					{Method: "restore", ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: "restore", ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
					{Method: "in-place", ToVersion: "8.0", SkipBackupSupported: core.BoolPtr(false)},
					{Method: "in-place", ToVersion: "9.5", SkipBackupSupported: core.BoolPtr(false)},
				},
			},
			expectedResult: []string{"8.0", "9.5"},
		},
		{
			description: "When only restore transitions exist, Expect nil",
			version: Version{
				Version: "6.0",
				Transitions: []VersionTransition{
					{Method: "restore", ToVersion: "6.1", SkipBackupSupported: core.BoolPtr(true)},
					{Method: "restore", ToVersion: "7.0", SkipBackupSupported: core.BoolPtr(false)},
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
						Method:              core.StringPtr("in-place"),
						SkipBackupSupported: core.BoolPtr(false),
					},
					{
						Application:         core.StringPtr("mongodb"),
						FromVersion:         core.StringPtr("6"),
						ToVersion:           core.StringPtr("8"),
						Method:              core.StringPtr("in-place"),
						SkipBackupSupported: core.BoolPtr(true),
					},
					{
						Application:         core.StringPtr("mongodb"),
						FromVersion:         core.StringPtr("6"),
						ToVersion:           core.StringPtr("10"),
						Method:              core.StringPtr("restore"),
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
		upgradeVersion     string
		skipBackup         bool
		mockCapabilityFunc func(capability string, instanceID string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error)
		expectedError      string
	}{
		{
			description:        "When there are no upgrade paths for the upgrade version, Expect no upgrade versions error",
			instanceID:         "test-instance",
			location:           "us-south",
			upgradeVersion:     "9",
			skipBackup:         false,
			mockCapabilityFunc: MockGetDeploymentCapabilityNoTransitions,
			expectedError:      "No available upgrade versions for your current version",
		},
		{
			description:        "When the upgrade version is no a valid version, Expect allowed versions error",
			instanceID:         "test-instance",
			location:           "us-south",
			upgradeVersion:     "10",
			skipBackup:         false,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "Version 10 is not a valid upgrade version. Allowed versions: [7 8]",
		},
		{
			description:        "When skip backup is not allowed for upgrade version, Expect skip backup error",
			instanceID:         "test-instance",
			location:           "eu-gb",
			upgradeVersion:     "7",
			skipBackup:         true,
			mockCapabilityFunc: MockGetDeploymentCapability,
			expectedError:      "Skipping backup is not allowed when upgrading to version 7",
		},
		{
			description:        "When the upgrade version is valid, Expect no error",
			instanceID:         "test-instance",
			location:           "eu-gb",
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
				capability, err := tc.mockCapabilityFunc("versions", instanceID, "classic", location, meta)
				require.NoError(t, err)
				return expandVersion(capability.Versions[0])
			}

			err := validateUpgradeVersion(tc.instanceID, tc.location, tc.upgradeVersion, tc.skipBackup, &MockMeta{})

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
