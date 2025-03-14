// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
)

/* TODO move other validators in here */

/* VERSION VALIDATOR */

type Version struct {
	Version     string
	Type        string
	Status      string
	IsPreferred bool
	Transitions []VersionTransition
}

type VersionTransition struct {
	Application         string
	Method              string
	SkipBackupSupported bool
	FromVersion         string
	ToVersion           string
}
type AllowedUpgrade struct {
	ToVersion           string `json:"to_version"`
	SkipBackupSupported bool   `json:"skip_backup_supported"`
}

func getAllowedUpgradeVersions(versions []Version) []AllowedUpgrade {
	var allowedVersions []AllowedUpgrade

	for _, version := range versions {
		for _, transition := range version.Transitions {
			if transition.Method == "restore" {
				allowedVersions = append(allowedVersions, AllowedUpgrade{
					ToVersion:           transition.ToVersion,
					SkipBackupSupported: transition.SkipBackupSupported,
				})
			}
		}
	}

	return allowedVersions
}

func (v *Version) isVersionAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version {
			return true
		}
	}
	return false
}

func (v *Version) isSkipBackupAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version {
			return transition.SkipBackupSupported
		}
	}
	return false
}

func expandVersions(versions []clouddatabasesv5.VersionsCapabilityItem) []*Version {
	if len(versions) == 0 {
		return nil
	}

	expandedVersions := make([]*Version, 0, len(versions))

	for _, capabilityVersion := range versions {
		version := Version{
			Version:     *capabilityVersion.Version,
			Status:      *capabilityVersion.Status,
			IsPreferred: *capabilityVersion.IsPreferred,
		}

		if capabilityVersion.Transitions != nil {
			var transitions []VersionTransition

			for _, transition := range capabilityVersion.Transitions {
				transitions = append(transitions, VersionTransition{
					Application:         *transition.Application,
					Method:              *transition.Method,
					SkipBackupSupported: *transition.SkipBackupSupported,
					FromVersion:         *transition.FromVersion,
					ToVersion:           *transition.ToVersion,
				})
			}
			version.Transitions = transitions
		}
		expandedVersions = append(expandedVersions, &version)
	}

	return expandedVersions
}

// This allows us to mock getDeploymentCapability in tests
var getDeploymentCapabilityFunc = getDeploymentCapability

func validateVersion(instanceID string, targetVersion string, skipBackup bool, meta interface{}) (err error) {
	// Get available versions for deployment
	capability, err := getDeploymentCapabilityFunc("versions", instanceID, "classic", "us-south", meta)
	if err != nil {
		log.Fatalf("Error fetching capability: %v", err)
	}

	if capability != nil && capability.Versions != nil {
		var versions []Version
		for _, v := range expandVersions(capability.Versions) {
			versions = append(versions, *v)
		}

		allowedVersions := getAllowedUpgradeVersions(versions)

		if len(allowedVersions) == 0 {
			log.Fatalf("No approved upgrade paths for your current version.")
		}

		for _, version := range versions {
			if version.isVersionAllowed(targetVersion) {
				if skipBackup && !version.isSkipBackupAllowed(targetVersion) {
					return fmt.Errorf("Skipping backup is not allowed when upgrading to version %s", targetVersion)
				}
				return nil
			}

			// Version is not allowed
			var allowedVersionList []string
			for _, upgrade := range allowedVersions {
				allowedVersionList = append(allowedVersionList, upgrade.ToVersion)
			}
			return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions: %v", targetVersion, allowedVersionList)

		}
	}

	return nil
}

/* VERSION VALIDATOR END */
