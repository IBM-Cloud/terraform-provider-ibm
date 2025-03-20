// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

// func TestIsVersionAllowed(t *testing.T) {
// 	testcases := []struct {
// 		version        Version
// 		targetVersion  string
// 		expectedResult bool
// 	}{
// 		{
// 			version: Version{
// 				Version: "6.0",
// 				Transitions: []VersionTransition{
// 					{ToVersion: "6.1"},
// 					{ToVersion: "7.0"},
// 				},
// 			},
// 			targetVersion:  "6.1",
// 			expectedResult: true,
// 		},
// 		{
// 			version: Version{
// 				Version: "7.0",
// 				Transitions: []VersionTransition{
// 					{ToVersion: "7.1"},
// 					{ToVersion: "7.2"},
// 				},
// 			},
// 			targetVersion:  "13.0",
// 			expectedResult: false,
// 		},
// 		{
// 			version: Version{
// 				Version:     "14",
// 				Transitions: []VersionTransition{},
// 			},
// 			targetVersion:  "15",
// 			expectedResult: false,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		result := tc.version.isVersionAllowed(tc.targetVersion)
// 		assert.Equal(t, tc.expectedResult, result)
// 	}
// }

// func TestIsSkipBackupAllowed(t *testing.T) {
// 	testcases := []struct {
// 		version        Version
// 		targetVersion  string
// 		expectedResult bool
// 	}{
// 		{
// 			version: Version{
// 				Version: "6.0",
// 				Transitions: []VersionTransition{
// 					{ToVersion: "6.1", SkipBackupSupported: true},
// 					{ToVersion: "7.0", SkipBackupSupported: false},
// 				},
// 			},
// 			targetVersion:  "6.1",
// 			expectedResult: true,
// 		},
// 		{
// 			version: Version{
// 				Version: "7.0",
// 				Transitions: []VersionTransition{
// 					{ToVersion: "8", SkipBackupSupported: false},
// 				},
// 			},
// 			targetVersion:  "8",
// 			expectedResult: false,
// 		},
// 		{
// 			version: Version{
// 				Version:     "14",
// 				Transitions: []VersionTransition{},
// 			},
// 			targetVersion:  "15",
// 			expectedResult: false,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		result := tc.version.isSkipBackupAllowed(tc.targetVersion)
// 		assert.Equal(t, tc.expectedResult, result)
// 	}
// }

// func TestValidateUpgradeVersion(t *testing.T) {
// 	testcases := []struct {
// 		versions       []Version
// 		targetVersion  string
// 		skipBackup     bool
// 		expectedError  string
// 	}{
// 		{
// 			versions: []Version{
// 				{
// 					Version: "1.0",
// 					Transitions: []VersionTransition{
// 						{ToVersion: "1.1", SkipBackupSupported: true},
// 					},
// 				},
// 			},
// 			targetVersion: "1.1",
// 			skipBackup:    false,
// 			expectedError: "",
// 		},
// 		{
// 			versions: []Version{
// 				{
// 					Version: "1.0",
// 					Transitions: []VersionTransition{
// 						{ToVersion: "1.1", SkipBackupSupported: true},
// 					},
// 				},
// 			},
// 			targetVersion: "2.0",
// 			skipBackup:    false,
// 			expectedError: "Version 2.0 is not a valid upgrade version. Allowed versions: [1.1]",
// 		},
// 		{
// 			versions: []Version{
// 				{
// 					Version: "1.0",
// 					Transitions: []VersionTransition{
// 						{ToVersion: "1.1", SkipBackupSupported: false},
// 					},
// 				},
// 			},
// 			targetVersion: "1.1",
// 			skipBackup:    true,
// 			expectedError: "Skipping backup is not allowed when upgrading to 1.1.",
// 		},
// 	}

// 	for _, tc := range testcases {
// 		err := validateVersion(tc.instanceID, tc.targetVersion, tc.skipBackup, mockMeta)

// 		if tc.expectedError == "" {
// 			assert.NoError(t, err)
// 		} else {
// 			assert.EqualError(t, err, tc.expectedError)
// 		}
// 	}
// }

// type MockMeta struct{}

// func MockGetDeploymentCapability(capability, instanceID, platform, location string, meta interface{}) (*clouddatabasesv5.Capability, error) {
// 	return &clouddatabasesv5.Capability{
// 		Versions: []clouddatabasesv5.VersionsCapabilityItem{
// 			{
// 				Version: core.StringPtr("5"),
// 				Status:  core.StringPtr("deprecated"),
// 				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{
// 					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("6")},
// 					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("7")},
// 					{Method: core.StringPtr("restore"), FromVersion: core.StringPtr("5"), ToVersion: core.StringPtr("8")},
// 				},
// 			},
// 		},
// 	}, nil
// }

// func MockGetDeploymentCapabilityNoTransitions(capability, instanceID, platform, location string, meta interface{}) (*clouddatabasesv5.Capability, error) {
// 	return &clouddatabasesv5.Capability{
// 		Versions: []clouddatabasesv5.VersionsCapabilityItem{
// 			{
// 				Version:     core.StringPtr("8"),
// 				Status:      core.StringPtr("preview"),
// 				Transitions: []clouddatabasesv5.VersionsCapabilityItemTransitionsItem{},
// 			},
// 		},
// 	}, nil
// }

// func TestValidateVersion(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		instanceID         string
// 		targetVersion      string
// 		skipBackup         bool
// 		mockCapabilityFunc func(capability string, instanceID string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error)
// 		expectedError      string
// 	}{
// 		{
// 			name:               "No upgrade paths",
// 			instanceID:         "test-instance",
// 			targetVersion:      "9",
// 			skipBackup:         false,
// 			mockCapabilityFunc: MockGetDeploymentCapabilityNoTransitions,
// 			expectedError:      "You are not allowed to upgrade version, there are no approved upgrade paths for your current version, please look at our docs here",
// 		},
// 		{
// 			name:               "Upgrade version not allowed",
// 			instanceID:         "test-instance",
// 			targetVersion:      "10",
// 			skipBackup:         false,
// 			mockCapabilityFunc: MockGetDeploymentCapability,
// 			expectedError:      "Version 10 is not a valid upgrade version. Allowed versions [6 7 8]",
// 		},
// 		{
// 			name:               "Valid version, no error",
// 			instanceID:         "test-instance",
// 			targetVersion:      "7",
// 			skipBackup:         false,
// 			mockCapabilityFunc: MockGetDeploymentCapability,
// 			expectedError:      "",
// 		},
// 	}

// 	for _, tc := range tests {
// 		getDeploymentCapabilityFunc = tc.mockCapabilityFunc

// 		err := validateVersion(tc.instanceID, tc.targetVersion, tc.skipBackup, &MockMeta{})

// 		// TODO fix actual versus expected
// 		// Log the actual result for debugging
// 		if err != nil {
// 			t.Logf("Test %s: received error: %v", tc.name, err)
// 		} else {
// 			t.Logf("Test %s: received no error", tc.name)
// 		}
// 		// TODO check nil error message
// 		if tc.expectedError != "" {
// 			assert.Error(t, err, tc.expectedError)
// 		} else {
// 			assert.NoError(t, err)
// 		}
// 	}
// }
