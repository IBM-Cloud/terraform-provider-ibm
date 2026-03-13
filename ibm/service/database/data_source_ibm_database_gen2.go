package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dataSourceIBMDatabaseGen2Backend struct{}

func newDataSourceIBMDatabaseGen2Backend() dataSourceIBMDatabaseBackend {
	return &dataSourceIBMDatabaseGen2Backend{}
}

func (g *dataSourceIBMDatabaseGen2Backend) Read(d *schema.ResourceData, meta interface{}) error {
	// NOTE - Edge case: potential stale values for unsupported Gen2 attributes.
	// If this data source was previously resolved to a Classic instance, all
	// attributes (including ones not supported by Gen2) would have been set.
	// If the same filters later resolve to a Gen2 instance (e.g., name/location/service),
	// Terraform will not automatically clear attributes that are no longer set,
	// unlike a resource which would ForceNew on such changes.
	// As a result, the Gen2 read path may only set supported attributes while
	// previously populated Classic-only attributes remain stale in state.
	// There is no clean mechanism to fully reset datasource state, and doing so
	// is generally considered an anti-pattern.
	// If this becomes an issue, unsupported attributes could be explicitly set
	// to null via d.Set() to ensure stale values are cleared.

	instance, err := findInstance(d, meta)
	if err != nil {
		return err
	}
	if instance == nil || instance.ID == nil {
		return fmt.Errorf("database instance not found")
	}
	d.SetId(*instance.ID)

	tags, err := flex.GetTagsUsingCRN(meta, d.Id())
	if err != nil {
		log.Printf(
			"Error on get of ibm Database tags (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)

	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	d.Set("guid", instance.GUID)
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return err
	}
	options := globalcatalogv1.GetCatalogEntryOptions{

		ID: instance.ResourceID,
	}
	service, _, err := globalClient.GetCatalogEntry(&options)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving service offering: %s", err)
	}

	d.Set("service", service.Name)

	planOptions := globalcatalogv1.GetCatalogEntryOptions{

		ID: instance.ResourcePlanID,
	}
	plan, _, err := globalClient.GetCatalogEntry(&planOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving plan: %s", err)
	}
	d.Set("plan", plan.Name)

	d.Set(flex.ResourceName, instance.Name)
	d.Set(flex.ResourceCRN, instance.CRN)
	d.Set(flex.ResourceStatus, instance.State)

	rMgtClient, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return err
	}
	GetResourceGroup := rg.GetResourceGroupOptions{
		ID: instance.ResourceGroupID,
	}
	resourceGroup, resp, err := rMgtClient.GetResourceGroup(&GetResourceGroup)
	if err != nil || resourceGroup == nil {
		log.Printf("[ERROR] Error retrieving resource group: %s %s", err, resp)
	}
	if resourceGroup != nil && resourceGroup.Name != nil {
		d.Set(flex.ResourceGroupName, resourceGroup.Name)
	}

	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))

	// Admin user is not available in Gen2. Users should manage credentials using ibm_resource_key.
	// Clear it from state if it was previously set (e.g., if the state was carried forward from a Classic instance).
	d.Set("adminuser", nil)

	// Extract version from instance.Extensions based on database type
	var version string
	if instance.Extensions != nil {
		dbType := getDatabaseTypeFromResourceID(*instance.ResourceID)
		if dbType != "" {
			if dataservices, ok := instance.Extensions["dataservices"].(map[string]interface{}); ok {
				if dbTypeData, ok := dataservices[dbType].(map[string]interface{}); ok {
					if v, ok := dbTypeData["version"].(string); ok {
						version = v
					}
				}
			}
		}
	}
	d.Set("version", version)

	// Extract platform_options from instance.Extensions for Gen2
	if instance.Extensions != nil {
		d.Set("platform_options", expandPlatformOptionsFromInstance(instance.Extensions))
	}

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting database client settings: %s", err)
	}

	listDeploymentScalingGroupsOptions := &clouddatabasesv5.ListDeploymentScalingGroupsOptions{
		ID: instance.ID,
	}

	groupList, _, err := cloudDatabasesClient.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting database groups: %s", err)
	}
	d.Set("groups", flex.FlattenIcdGroups(groupList))

	// Auto scaling is currently not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("auto_scaling", nil)

	// Allowlist is not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("allowlist", nil)

	return nil
}
