package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
)

/*
	TODO LORNA In Place upgrades
	User changes version value
	- check if version is valid CustomDiff
	- Get capability versions
		oldVersion, newVersion := diff.GetChange("version")
       check new version in response transitions
	- Validation:
	 If insatnce exists check transition, otherwise just check valid version
		-  version not allowed. If not message with allowed versions
		- check if remote_leader_id error
		- If skip_backup field (do we display a warning?)
		if timeout is 1 hour
		- If expiration date???? check that it isnt >24hours from now (default: 5mins) use the timeouts update instead. Whatever this is set to thats the expiry. Just wait for it to start
	- Call version endpoint with version, backup and expiration datetime
	- Check for new provisions
*/

type AllowedUpgrade struct {
	ToVersion           string `json:"to_version"`
	SkipBackupSupported bool   `json:"skip_backup_supported"`
}

func getDeploymentCapability(capabilityId string, deploymentId string, platform string, location string, meta interface{}) (*clouddatabasesv5.Capability, error) {

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return nil, err
	}

	getDeploymentCapabilityOptions := &clouddatabasesv5.GetDeploymentCapabilityOptions{
		ID:             core.StringPtr(deploymentId),
		CapabilityID:   core.StringPtr(capabilityId),
		TargetPlatform: core.StringPtr(fmt.Sprintf("target_platform=%s", platform)),
		TargetLocation: core.StringPtr(fmt.Sprintf("target_location=%s", location)),
	}
	getDeploymentCapabilityResponse, response, err := cloudDatabasesClient.GetDeploymentCapability(getDeploymentCapabilityOptions)

	// Check if response is nil before proceeding
	if getDeploymentCapabilityResponse == nil || getDeploymentCapabilityResponse.Capability == nil {
		return nil, fmt.Errorf("capability '%s' field is nil in response %s", capabilityId, response)
	}

	// Convert response to JSON and unmarshal into Capability struct
	jsonData, err := json.Marshal(getDeploymentCapabilityResponse.Capability)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal capability response: %w", err)
	}

	var capability clouddatabasesv5.Capability
	err = json.Unmarshal(jsonData, &capability)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal capability response: %w", err)
	}

	return &capability, nil
}

func getAllowedUpgradeVersions(capability clouddatabasesv5.Capability) []AllowedUpgrade {
	var allowedVersions []AllowedUpgrade

	for _, version := range capability.Versions {
		for _, transition := range version.Transitions {
			// TODO change to in-place, fix skip backup once API is complete
			if transition.Method != nil && *transition.Method == "restore" {
				allowedVersions = append(allowedVersions, AllowedUpgrade{
					ToVersion:           *transition.ToVersion,
					SkipBackupSupported: false, // transition.SkipBackupSupported,
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

func validateVersion(instanceID string, newVersion string, skipBackup bool, meta interface{}) (err error) {
	capability, err := getDeploymentCapability("versions", instanceID, "classic", "us-south", meta)
	if err != nil {
		log.Fatalf("Error fetching capability: %v", err)
	}

	allowedVersions := getAllowedUpgradeVersions(*capability)

	if len(allowedVersions) == 0 {
		return fmt.Errorf("You are not allowed to upgrade version, there are no approved upgrade paths for your current version, please look at our docs here")
	}

	isAllowed := isVersionAllowed(newVersion, allowedVersions)

	if isAllowed == false {
		allowedVersionList := []string{}

		for _, upgrade := range allowedVersions {
			allowedVersionList = append(allowedVersionList, upgrade.ToVersion)
		}
		return fmt.Errorf("Version %s is not a valid upgrade version. Allowed versions %v", newVersion, allowedVersionList)
	}

	if skipBackup == true {
		isAllowedSkipBackup := isSkipBackupAllowed(newVersion, allowedVersions)

		if isAllowedSkipBackup != true {
			return fmt.Errorf("You are not allowed to skip taking a backup when upgrading to version %s. Please remove skip_backup or update field to false", newVersion)
		}
	}

	return nil
}
