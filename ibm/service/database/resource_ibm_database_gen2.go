package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
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
	if instance.CRN != nil {
		location := strings.Split(*instance.CRN, ":")
		if len(location) > 5 {
			d.Set("location", location[5])
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

	icdClient, err := meta.(conns.ClientSession).ICDAPI()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database client settings: %s", err))
	}

	icdId := flex.EscapeUrlParm(instanceID)

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database client settings: %s", err))
	}

	getDeploymentInfoOptions := &clouddatabasesv5.GetDeploymentInfoOptions{
		ID: core.StringPtr(instanceID),
	}
	getDeploymentInfoResponse, response, err := cloudDatabasesClient.GetDeploymentInfo(getDeploymentInfoOptions)

	if err != nil {
		if response.StatusCode == 404 {
			return diag.FromErr(fmt.Errorf("[ERROR] The database instance was not found in the region set for the Provider, or the default of us-south. Specify the correct region in the provider definition, or create a provider alias for the correct region. %v", err))
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database config while updating adminpassword for: %s with error %s", instanceID, err))
	}

	deployment := getDeploymentInfoResponse.Deployment

	d.Set("adminuser", deployment.AdminUsernames["database"])
	d.Set("version", deployment.Version)

	listDeploymentScalingGroupsOptions := &clouddatabasesv5.ListDeploymentScalingGroupsOptions{
		ID: core.StringPtr(instanceID),
	}
	groupList, _, err := cloudDatabasesClient.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database groups: %s", err))
	}
	if len(groupList.Groups) == 0 || groupList.Groups[0].Members == nil || groupList.Groups[0].Members.AllocationCount == nil || *groupList.Groups[0].Members.AllocationCount == 0 {
		return diag.FromErr(fmt.Errorf("[ERROR] This database appears to have have 0 members. Unable to proceed"))
	}

	d.Set("groups", flex.FlattenIcdGroups(groupList))

	getAutoscalingConditionsOptions := &clouddatabasesv5.GetAutoscalingConditionsOptions{
		ID:      instance.ID,
		GroupID: core.StringPtr("member"),
	}

	autoscalingGroup, _, err := cloudDatabasesClient.GetAutoscalingConditions(getAutoscalingConditionsOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database autoscaling groups: %s\n Hint: Check if there is a mismatch between your database location and IBMCLOUD_REGION", err))
	}
	d.Set("auto_scaling", flattenAutoScalingGroup(*autoscalingGroup))

	alEntry := &clouddatabasesv5.GetAllowlistOptions{
		ID: &instanceID,
	}

	allowlist, _, err := cloudDatabasesClient.GetAllowlist(alEntry)

	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database allowlist: %s", err))
	}

	d.Set("allowlist", flex.FlattenAllowlist(allowlist.IPAddresses))

	//ICD does not implement a GetUsers API. Users populated from tf configuration.
	tfusers := d.Get("users").(*schema.Set)
	users := flex.ExpandUsers(tfusers)
	user := icdv4.User{
		UserName: deployment.AdminUsernames["database"],
	}
	users = append(users, user)

	if serviceOff == "databases-for-postgresql" || serviceOff == "databases-for-redis" || serviceOff == "databases-for-enterprisedb" {
		configSchema, err := icdClient.Configurations().GetConfiguration(icdId)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error getting database (%s) configuration schema : %s", icdId, err))
		}
		s, err := json.Marshal(configSchema)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error marshalling the database configuration schema: %s", err))
		}

		// depends on GET configuration/schema - checking if we have something similar in CDP
		if err = d.Set("configuration_schema", string(s)); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting the database configuration schema: %s", err))
		}
	}

	endpoint, _ := instance.Parameters["service-endpoints"]
	if endpoint == "public" || endpoint == "public-and-private" {
		return publicServiceEndpointsWarning()
	}

	tm := &TaskManager{
		Client:     cloudDatabasesClient,
		InstanceID: instanceID,
	}

	upgradeInProgress, upgradeTask, err := tm.matchingTaskInProgress(taskUpgrade)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting tasks for instance: %w", err))
	}

	if upgradeInProgress {
		return upgradeInProgressWarning(upgradeTask)
	}

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
