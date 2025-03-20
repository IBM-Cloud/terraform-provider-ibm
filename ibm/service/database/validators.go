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

/* TODO
How to mock fetch function
How to test against a branch for go dsk etc
*/

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

// This allows us to mock getDeploymentCapability in tests
var getDeploymentCapabilityFunc = getDeploymentCapability

func fetchDeploymentVersion(instanceId string, meta interface{}) *Version {
	capability, err := getDeploymentCapabilityFunc("versions", instanceId, "classic", "us-south", meta)
	if err != nil {
		log.Fatalf("Error fetching versions: %v", err)
	}

	if capability == nil || capability.Versions == nil || len(capability.Versions) == 0 {
		return nil
	}

	// Expand the first version as there's only one expected
	version := expandVersion(capability.Versions[0])

	return version
}

func (v *Version) isVersionAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version && transition.Method == "restore" { //TODO replace restore with in-place
			return true
		}
	}
	return false
}

func (v *Version) isSkipBackupAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version && transition.SkipBackupSupported != nil && transition.Method == "restore" { //TODO replace restore with in-place
			return *transition.SkipBackupSupported
		}
	}
	return false
}

func (v *Version) hasUpgradeVersions() bool {
	for _, transition := range v.Transitions {
		if transition.Method == "restore" { //TODO replace restore with in-place
			return true
		}
	}
	return false
}

func (v *Version) getAllowedVersionsList() []string {
	var allowedList []string
	for _, transition := range v.Transitions {
		if transition.Method == "restore" { //TODO replace restore with in-place
			allowedList = append(allowedList, transition.ToVersion)
		}
	}
	return allowedList
}

func validateVersion(instanceId string, targetVersion string, skipBackup bool, meta interface{}) (err error) {
	deploymentVersion := fetchDeploymentVersion(instanceId, meta)

	if deploymentVersion == nil || deploymentVersion.hasUpgradeVersions() == false {
		return fmt.Errorf("No available upgrade versions for your current version.")
	}

	if !deploymentVersion.isVersionAllowed(targetVersion) {
		allowedList := deploymentVersion.getAllowedVersionsList()
		return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions: %v", targetVersion, allowedList)
	}

	if skipBackup && !deploymentVersion.isSkipBackupAllowed(targetVersion) {
		return fmt.Errorf("Skipping backup is not allowed when upgrading to version %s", targetVersion)
	}

	return nil
}

/* VERSION VALIDATOR END */
