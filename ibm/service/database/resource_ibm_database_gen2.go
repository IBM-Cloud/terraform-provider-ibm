package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var gen2UnsupportedAttrs = []string{
	// TODO: update the list
	"backup_policy",
	"users",
}

type resourceIBMDatabaseGen2Backend struct{}

func newResourceIBMDatabaseGen2Backend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseGen2Backend{}
}

func (g *resourceIBMDatabaseGen2Backend) Create(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet (plan=%q)", d.Get("plan").(string))
}

func (g *resourceIBMDatabaseGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceID := d.Id()
	rsInst := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}
	instance, response, err := rsConClient.GetResourceInstance(&rsInst)
	if err != nil {
		if strings.Contains(err.Error(), "Object not found") ||
			strings.Contains(err.Error(), "status code: 404") {
			log.Printf("[WARN] Removing record from state because it's not found via the API")
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response))
	}
	if strings.Contains(*instance.State, "removed") {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return nil
	}

	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf(
			"Error on get of ibm Database tags (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)
	d.Set("name", *instance.Name)
	d.Set("status", *instance.State)
	d.Set("resource_group_id", *instance.ResourceGroupID)
	var instanceLocation string
	if instance.CRN != nil {
		location := strings.Split(*instance.CRN, ":")
		if len(location) > 5 {
			instanceLocation = location[5]
			d.Set("location", instanceLocation)
		}
	}
	d.Set("guid", *instance.GUID)

	if instance.Parameters != nil {
		if endpoint, ok := instance.Parameters["service-endpoints"]; ok {
			d.Set("service_endpoints", endpoint)
		}
	}

	d.Set(flex.ResourceName, *instance.Name)
	d.Set(flex.ResourceCRN, *instance.CRN)
	d.Set(flex.ResourceStatus, *instance.State)
	d.Set(flex.ResourceGroupName, *instance.ResourceGroupCRN)

	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))

	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return diag.FromErr(err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.GetServiceName(*instance.ResourceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving service offering: %s", err))
	}

	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(*instance.ResourcePlanID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving plan: %s", err))
	}
	d.Set("plan", servicePlan)

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

	// Get groups data from GlobalCatalog for Gen2
	// Find the deployment by getting plan's children and matching by location
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	var catalogDeployment *globalcatalogv1.CatalogEntry
	kind := "deployment"
	childOptions := globalcatalogv1.GetChildObjectsOptions{
		ID:   instance.ResourcePlanID,
		Kind: &kind,
	}
	children, _, err := globalClient.GetChildObjects(&childOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving plan children: %s", err))
	}

	if children != nil && children.Resources != nil {
		for _, child := range children.Resources {
			// Check if this deployment's location matches the instance region
			if child.Metadata != nil &&
				child.Metadata.Deployment != nil &&
				child.Metadata.Deployment.Location != nil &&
				*child.Metadata.Deployment.Location == instanceLocation {
				catalogDeployment = &child
				break
			}
		}
	}

	if catalogDeployment == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Could not find deployment catalog entry for region %s", instanceLocation))
	}

	// Extract resources from deployment metadata
	var catalogResources []interface{}
	if catalogDeployment.Metadata != nil && catalogDeployment.Metadata.Other != nil {
		if resources, ok := catalogDeployment.Metadata.Other["resources"].([]interface{}); ok {
			catalogResources = resources
		}
	}

	// Flatten groups using instance extensions and catalog metadata
	if instance.Extensions != nil && len(catalogResources) > 0 {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	// Auto scaling is currently not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("auto_scaling", nil)

	// Allowlist is not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("allowlist", nil)

	// Users are not managed in Gen2 via this resource. Users should manage credentials using ibm_resource_key.
	// Clear it from state if it was previously set (e.g., if the state was carried forward from a Classic instance).
	d.Set("users", nil)

	// Configuration schema is currently not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("configuration_schema", nil)

	return nil

}

func (g *resourceIBMDatabaseGen2Backend) Update(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) Delete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, fmt.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (g *resourceIBMDatabaseGen2Backend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	var bad []string
	for _, k := range gen2UnsupportedAttrs {
		if isAttrConfiguredInDiff(d, k) {
			bad = append(bad, k)
		}
	}
	if len(bad) == 0 {
		return nil
	}

	planRaw, _ := d.GetOk("plan")
	plan, _ := planRaw.(string)

	return fmt.Errorf(
		"plan %q indicates Gen2. The following attributes are not supported for Gen2 and must be removed: %s",
		strings.TrimSpace(plan),
		strings.Join(bad, ", "),
	)
}
