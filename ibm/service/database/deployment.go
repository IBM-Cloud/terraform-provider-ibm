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
   TODO Move other deployment api endpoints in here
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
