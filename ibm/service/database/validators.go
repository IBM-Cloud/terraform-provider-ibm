// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* TODO move other validators in here */

/* VERSION VALIDATOR */

const (
	versions        = "versions"
	classicPlatform = "classic"
	inPlace         = "in-place"
	restore         = "restore"
)

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
		if transition.ToVersion == version && transition.Method == inPlace {
			return true
		}
	}
	return false
}

func (v *Version) isSkipBackupUpgradeAllowed(version string) bool {
	for _, transition := range v.Transitions {
		if transition.ToVersion == version && transition.SkipBackupSupported != nil && transition.Method == inPlace {
			return *transition.SkipBackupSupported
		}
	}
	return false
}

func (v *Version) hasUpgradeVersions() bool {
	for _, transition := range v.Transitions {
		if transition.Method == inPlace {
			return true
		}
	}
	return false
}

func (v *Version) getAllowedVersionsList() []string {
	var allowedList []string
	for _, transition := range v.Transitions {
		if transition.Method == inPlace {
			allowedList = append(allowedList, transition.ToVersion)
		}
	}
	return allowedList
}

var fetchDeploymentVersionFn = fetchDeploymentVersion

func fetchDeploymentVersion(instanceId string, location string, meta interface{}) *Version {
	options := DeploymentCapabilityOptions{
		Platform:      classicPlatform,
		Location:      location,
		IncludeHidden: core.BoolPtr(true),
		IncludeBeta:   core.BoolPtr(true),
	}

	capability, err := getDeploymentCapability(versions, instanceId, options, meta)
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

func validateUpgradeVersion(instanceId string, location string, oldVersion string, upgradeVersion string, skipBackup bool, meta interface{}) (err error) {
	deploymentVersion := fetchDeploymentVersionFn(instanceId, location, meta)

	if deploymentVersion == nil || deploymentVersion.hasUpgradeVersions() == false {
		return fmt.Errorf("No available upgrade versions for version %s", oldVersion)
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

func validateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return pickResourceBackendFromDiff(d).ValidateUnsupportedAttrsDiff(ctx, d, meta)
}
