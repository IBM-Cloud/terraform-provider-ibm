// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
)

/* Deployment api calls
   TODO Move other deployment api endpoints in here
*/

type CapabilityOptions struct {
	Platform      string
	Location      string
	IncludeHidden *bool
	IncludeBeta   *bool
}

func getDeploymentCapability(capabilityId string, deploymentId string, options CapabilityOptions, meta interface{}) (*clouddatabasesv5.Capability, error) {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return nil, err
	}

	getDeploymentCapabilityOptions := &clouddatabasesv5.GetDeploymentCapabilityOptions{
		ID:             core.StringPtr(deploymentId),
		CapabilityID:   core.StringPtr(capabilityId),
		TargetPlatform: core.StringPtr(fmt.Sprintf("target_platform=%s", options.Platform)),
		TargetLocation: core.StringPtr(fmt.Sprintf("target_location=%s", options.Location)),
	}

	if options.IncludeHidden != nil {
		getDeploymentCapabilityOptions.IncludeHidden = core.BoolPtr(*options.IncludeHidden)
	}

	if options.IncludeBeta != nil {
		getDeploymentCapabilityOptions.IncludeBeta = core.BoolPtr(*options.IncludeBeta)
	}

	getDeploymentCapabilityResponse, response, err := cloudDatabasesClient.GetDeploymentCapability(getDeploymentCapabilityOptions)

	if getDeploymentCapabilityResponse == nil || getDeploymentCapabilityResponse.Capability == nil {
		return nil, fmt.Errorf("capability '%s' field is nil in response %s", capabilityId, response)
	}

	return getDeploymentCapabilityResponse.Capability, nil
}
