package database

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dataSourceIBMDatabasePointInTimeRecoveryGen2Backend struct{}

func newDataSourceIBMDatabasePointInTimeRecoveryGen2Backend() *dataSourceIBMDatabasePointInTimeRecoveryGen2Backend {
	return &dataSourceIBMDatabasePointInTimeRecoveryGen2Backend{}
}

func (g *dataSourceIBMDatabasePointInTimeRecoveryGen2Backend) Read(d *schema.ResourceData, meta interface{}) error {
	instance, err := g.findInstanceByDeploymentID(d, meta)
	if err != nil {
		return err
	}

	location := ""
	if instance.RegionID != nil {
		location = *instance.RegionID
	}
	if location == "" {
		return fmt.Errorf("failed to determine location for deployment %s", d.Get("deployment_id").(string))
	}

	d.SetId(d.Get("deployment_id").(string))

	return fmt.Errorf("point-in-time recovery metadata for Gen2 deployment %s is not exposed through Resource Controller APIs. Resolved instance id=%s, location=%s", d.Get("deployment_id").(string), *instance.ID, location)
}

func (g *dataSourceIBMDatabasePointInTimeRecoveryGen2Backend) findInstanceByDeploymentID(d *schema.ResourceData, meta interface{}) (*rc.ResourceInstance, error) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, err
	}

	deploymentID := d.Get("deployment_id").(string)
	listOptions := &rc.ListResourceInstancesOptions{
		GUID: &deploymentID,
	}

	var instances []rc.ResourceInstance
	nextURL := ""

	for {
		if nextURL != "" {
			listOptions.SetStart(nextURL)
		}

		listResponse, resp, err := rsConClient.ListResourceInstances(listOptions)
		if err != nil {
			return nil, fmt.Errorf("[ERROR] Error retrieving resource instance for deployment [%s]: %s with resp code: %s", deploymentID, err, resp)
		}

		instances = append(instances, listResponse.Resources...)

		nextURL, err = getInstancesNext(listResponse.NextURL)
		if err != nil {
			return nil, fmt.Errorf("[DEBUG] ListResourceInstances failed. Error occurred while parsing NextURL: %s", err)
		}

		if nextURL == "" {
			break
		}
	}

	if len(instances) == 0 {
		return nil, fmt.Errorf("[ERROR] No resource instance found with deployment_id [%s]", deploymentID)
	}

	if len(instances) > 1 {
		return nil, fmt.Errorf("More than one resource instance found with deployment_id [%s]", deploymentID)
	}

	return &instances[0], nil
}
