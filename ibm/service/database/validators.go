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

func transformVersions(versions []clouddatabasesv5.VersionsCapabilityItem) []*Version {
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

		// Process transitions as an array of VersionTransition
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

func validateVersion(instanceID string, targetVersion string, skipBackup bool, meta interface{}) (err error) {
	// Get available versions for deployment
	capability, err := getDeploymentCapability("versions", instanceID, "classic", "us-south", meta)
	if err != nil {
		log.Fatalf("Error fetching capability: %v", err)
	}

	if capability != nil && capability.Versions != nil {
		// Convert capability.Versions to []interface{}
		var versions []Version
		for _, v := range transformVersions(capability.Versions) {
			versions = append(versions, *v)
		}

		// Now pass the versionsData to expandVersions
		allowedVersions := getAllowedUpgradeVersions(versions)

		if len(allowedVersions) == 0 {
			log.Fatalf("No approved upgrade paths for your current version. Check the docs.")
		}

		// Now you can work with expandedVersions, which is []*Version
		fmt.Println("Trans Versions:", versions)

		for _, version := range versions {
			if version.isVersionAllowed(targetVersion) {
				// validate here
				if skipBackup && !version.isSkipBackupAllowed(targetVersion) {
					return fmt.Errorf("Skipping backup is not allowed when upgrading to %s", targetVersion)
				}
				// Upgrade allowed, proceed
				return nil
			}

			// If we reach here, the version was not allowed
			var allowedVersionList []string
			for _, upgrade := range allowedVersions {
				allowedVersionList = append(allowedVersionList, upgrade.ToVersion)
			}
			return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions: %v", targetVersion, allowedVersionList)

		}

	}

	// if skipBackup == true {
	// 	isAllowedSkipBackup := isSkipBackupAllowed(version, allowedVersions)

	// 	if isAllowedSkipBackup != true {
	// 		return fmt.Errorf("You are not allowed to skip taking a backup when upgrading to version %s. Please remove version_upgrade_skip_backup or update field to false", version)
	// 	}
	// }

	return nil
}

/* VERSION VALIDATOR END */
