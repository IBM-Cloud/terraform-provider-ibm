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
	FromVersion         string
	ToVersion           string
	SkipBackupSupported *bool
}

func expandVersion(version clouddatabasesv5.VersionsCapabilityItem) *Version {
	expandedVersion := &Version{
		Version:     *version.Version,
		Status:      *version.Status,
		IsPreferred: *version.IsPreferred,
	}

	if version.Transitions != nil {
		var transitions []VersionTransition

		for _, transition := range version.Transitions {
			versionTransition := VersionTransition{
				Application: *transition.Application,
				Method:      *transition.Method,
				FromVersion: *transition.FromVersion,
				ToVersion:   *transition.ToVersion,
			}

			if skipBackup := transition.SkipBackupSupported; skipBackup != nil {
				versionTransition.SkipBackupSupported = skipBackup
			}

			transitions = append(transitions, versionTransition)
		}
		expandedVersion.Transitions = transitions
	}

	return expandedVersion
}

func (v *Version) isVersionUpgradeAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version && transition.Method == "in-place" {
			return true
		}
	}
	return false
}

func (v *Version) isSkipBackupUpgradeAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version && transition.SkipBackupSupported != nil && transition.Method == "in-place" {
			return *transition.SkipBackupSupported
		}
	}
	return false
}

func (v *Version) hasUpgradeVersions() bool {
	for _, transition := range v.Transitions {
		if transition.Method == "in-place" {
			return true
		}
	}
	return false
}

func (v *Version) getAllowedVersionsList() []string {
	var allowedList []string
	for _, transition := range v.Transitions {
		if transition.Method == "in-place" {
			allowedList = append(allowedList, transition.ToVersion)
		}
	}
	return allowedList
}

var fetchDeploymentVersionFn = fetchDeploymentVersion

func fetchDeploymentVersion(instanceId string, location string, meta interface{}) *Version {
	capability, err := getDeploymentCapability("versions", instanceId, "classic", location, meta)
	if err != nil {
		log.Fatalf("Error fetching deployment versions: %v", err)
	}

	if capability == nil || capability.Versions == nil || len(capability.Versions) == 0 {
		return nil
	}

	// Expand the first version as there's only one expected
	version := expandVersion(capability.Versions[0])

	return version
}

func validateUpgradeVersion(instanceId string, location string, upgradeVersion string, skipBackup bool, meta interface{}) (err error) {
	deploymentVersion := fetchDeploymentVersionFn(instanceId, location, meta)

	if deploymentVersion == nil || deploymentVersion.hasUpgradeVersions() == false {
		return fmt.Errorf("No available upgrade versions for your current version.")
	}

	if !deploymentVersion.isVersionUpgradeAllowed(upgradeVersion) {
		allowedList := deploymentVersion.getAllowedVersionsList()
		return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions: %v", upgradeVersion, allowedList)
	}

	if skipBackup && !deploymentVersion.isSkipBackupUpgradeAllowed(upgradeVersion) {
		return fmt.Errorf("Skipping backup is not allowed when upgrading to version %s", upgradeVersion)
	}

	return nil
}

/* VERSION VALIDATOR END */
