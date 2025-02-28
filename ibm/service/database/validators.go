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
type AllowedUpgrade struct {
	ToVersion           string `json:"to_version"`
	SkipBackupSupported bool   `json:"skip_backup_supported"`
}

func getAllowedUpgradeVersions(capability clouddatabasesv5.Capability) []AllowedUpgrade {
	var allowedVersions []AllowedUpgrade

	for _, version := range capability.Versions {
		for _, transition := range version.Transitions {
			// TODO change to in-place
			if transition.Method != nil && *transition.Method == "restore" {
				allowedVersions = append(allowedVersions, AllowedUpgrade{
					ToVersion:           *transition.ToVersion,
					SkipBackupSupported: *transition.SkipBackupSupported,
				})
			}
		}
	}

	return allowedVersions

}

func isVersionAllowed(newVersion string, allowedVersions []AllowedUpgrade) bool {
	for _, upgrade := range allowedVersions {
		if upgrade.ToVersion == newVersion {
			return true
		}
	}
	return false
}

func isSkipBackupAllowed(newVersion string, allowedVersions []AllowedUpgrade) bool {
	for _, upgrade := range allowedVersions {
		if upgrade.ToVersion == newVersion {
			return upgrade.SkipBackupSupported
		}
	}
	return false
}

var (
	// TODO: FInd a better way to do this, this allows mocking for unit tests
	getDeploymentCapabilityFunc = getDeploymentCapability
)

func validateVersion(instanceID string, version string, skipBackup bool, meta interface{}) (err error) {
	// TODO make more ipu specific
	// Get available versions for deployment
	capability, err := getDeploymentCapabilityFunc("versions", instanceID, "classic", "us-south", meta)
	if err != nil {
		log.Fatalf("Error fetching capability: %v", err)
	}
	allowedVersions := getAllowedUpgradeVersions(*capability)

	if len(allowedVersions) == 0 {
		return fmt.Errorf("You are not allowed to upgrade version, there are no approved upgrade paths for your current version, please look at our docs here")
	}

	isAllowed := isVersionAllowed(version, allowedVersions)

	if isAllowed == false {
		allowedVersionList := []string{}

		for _, upgrade := range allowedVersions {
			allowedVersionList = append(allowedVersionList, upgrade.ToVersion)
		}
		return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions %v", version, allowedVersionList)
	}

	if skipBackup == true {
		isAllowedSkipBackup := isSkipBackupAllowed(version, allowedVersions)

		if isAllowedSkipBackup != true {
			return fmt.Errorf("You are not allowed to skip taking a backup when upgrading to version %s. Please remove version_upgrade_skip_backup or update field to false", version)
		}
	}

	return nil
}

/* VERSION VALIDATOR END */
