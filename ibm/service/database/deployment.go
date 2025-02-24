// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"encoding/json"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
)

/* Deployment api calls
   TODO Move other deployment api endpoints in here */

/*
	TODO LORNA In Place upgrades
	User changes version value
	- check if version is valid CustomDiff DONE
	- Get capability versions DONE
		oldVersion, newVersion := diff.GetChange("version") DONE
       check new version in response transitionsDONE
	- Validation:
	 If insatnce exists check transition, otherwise just check valid version
		-  version not allowed. If not message with allowed versions DONE
		- check if remote_leader_id error API validation
		- If skip_backup field (do we display a warning?)
		if timeout is 1 hour
		- If expiration date???? check that it isnt >24hours from now (default: 5mins) use the timeouts update instead. Whatever this is set to thats the expiry. Just wait for it to start
	- Call version endpoint with version, backup and expiration datetime
	- Check for new provisions
*/

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

	if getDeploymentCapabilityResponse == nil || getDeploymentCapabilityResponse.Capability == nil {
		return nil, fmt.Errorf("capability '%s' field is nil in response %s", capabilityId, response)
	}

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

func listDeploymentTasks(deploymentId string, meta interface{}) (*clouddatabasesv5.Tasks, error) {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return nil, err
	}

	listDeploymentTasksOptions := &clouddatabasesv5.ListDeploymentTasksOptions{
		ID: core.StringPtr(deploymentId),
	}

	listDeploymentTasksResponse, response, err := cloudDatabasesClient.ListDeploymentTasks(listDeploymentTasksOptions)

	if listDeploymentTasksResponse == nil {
		return nil, fmt.Errorf("field is nil in response tasks response %s", response)
	}

	jsonData, err := json.Marshal(listDeploymentTasksResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tasks response: %w", err)
	}

	var tasks clouddatabasesv5.Tasks
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks response: %w", err)
	}

	return &tasks, nil
}
