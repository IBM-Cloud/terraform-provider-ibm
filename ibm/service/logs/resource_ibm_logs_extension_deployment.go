// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func ResourceIbmLogsExtensionDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsExtensionDeploymentCreate,
		ReadContext:   resourceIbmLogsExtensionDeploymentRead,
		UpdateContext: resourceIbmLogsExtensionDeploymentUpdate,
		DeleteContext: resourceIbmLogsExtensionDeploymentDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the IBM Cloud Logs instance.",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region of the IBM Cloud Logs instance.",
			},
			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "public or private.",
			},
			"extension_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the extension to deploy.",
			},
			"extension_deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ExtensionDeployment Id.",
			},
			"version": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_extension_deployment", "version"),
				Description:  "The version of the Extension revision to deploy.",
			},
			"item_ids": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of Extension item IDs to deploy.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"applications": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Applications that the Extension is deployed for. When this is empty, it is applied to all applications.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"subsystems": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIbmLogsExtensionDeploymentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "version",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$`,
			MinValueLength:             1,
			MaxValueLength:             4096,
		},
	)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "code",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "bad_request_or_unspecified, unauthorized, forbidden, not_found, method_internal_error, conflict, unauthenticated, resource_exhausted, deadline_exceeded",
		},
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "target_domain",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "alert_definition, alert, enrichment, rule_group, view, dashboard, events_to_metrics",
		},
	)
	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_extension_deployment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsExtensionDeploymentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_extension_deployment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("getClientWithLogsInstanceEndpoint failed: %s", err.Error()), "ibm_logs_extension_deployment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateExtensionDeploymentOptions := &logsv0.UpdateExtensionDeploymentOptions{}

	// Set the extension ID - deployment is tied to a specific extension
	extensionId := d.Get("extension_id").(string)
	updateExtensionDeploymentOptions.SetID(extensionId)

	updateExtensionDeploymentOptions.SetVersion(d.Get("version").(string))
	var newItemIds []string
	for _, v := range d.Get("item_ids").([]interface{}) {
		newItemIdsItem := v.(string)
		newItemIds = append(newItemIds, newItemIdsItem)
	}
	updateExtensionDeploymentOptions.SetItemIds(newItemIds)
	if _, ok := d.GetOk("applications"); ok {
		var newApplications []string
		for _, v := range d.Get("applications").([]interface{}) {
			newApplicationsItem := v.(string)
			newApplications = append(newApplications, newApplicationsItem)
		}
		updateExtensionDeploymentOptions.SetApplications(newApplications)
	}
	if _, ok := d.GetOk("subsystems"); ok {
		var newSubsystems []string
		for _, v := range d.Get("subsystems").([]interface{}) {
			newSubsystemsItem := v.(string)
			newSubsystems = append(newSubsystems, newSubsystemsItem)
		}
		updateExtensionDeploymentOptions.SetSubsystems(newSubsystems)
	}

	_, _, err = logsClient.UpdateExtensionDeploymentWithContext(context, updateExtensionDeploymentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateExtensionDeploymentWithContext failed: %s", err.Error()), "ibm_logs_extension_deployment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use extension ID as the deployment identifier
	d.SetId(fmt.Sprintf("%s/%s/%s", region, instanceId, extensionId))

	return resourceIbmLogsExtensionDeploymentRead(context, d, meta)
}

func resourceIbmLogsExtensionDeploymentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_extension_deployment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, region, instanceId, extensionId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("updateClientURLWithInstanceEndpoint failed: %s", err.Error()), "ibm_logs_extension_deployment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getExtensionDeploymentOptions := &logsv0.GetExtensionDeploymentOptions{}
	getExtensionDeploymentOptions.SetID(extensionId)

	extensionDeployment, response, err := logsClient.GetExtensionDeploymentWithContext(context, getExtensionDeploymentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetExtensionDeploymentWithContext failed: %s", err.Error()), "ibm_logs_extension_deployment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Set additional attributes required for import
	if err = d.Set("extension_id", extensionId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting extension_id: %s", err))
	}
	// extension_deployment_id is the same as extension_id since deployments are identified by extension ID
	if err = d.Set("extension_deployment_id", extensionId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting extension_deployment_id: %s", err))
	}
	if err = d.Set("instance_id", instanceId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if err = d.Set("region", region); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
	}

	if err = d.Set("version", extensionDeployment.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("item_ids", extensionDeployment.ItemIds); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting item_ids: %s", err))
	}
	if !core.IsNil(extensionDeployment.Applications) {
		if err = d.Set("applications", extensionDeployment.Applications); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting applications: %s", err))
		}
	}
	if !core.IsNil(extensionDeployment.Subsystems) {
		if err = d.Set("subsystems", extensionDeployment.Subsystems); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting subsystems: %s", err))
		}
	}

	return nil
}

func resourceIbmLogsExtensionDeploymentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_extension_deployment", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, _, _, _, err = updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("updateClientURLWithInstanceEndpoint failed: %s", err.Error()), "ibm_logs_extension_deployment", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateExtensionDeploymentOptions := &logsv0.UpdateExtensionDeploymentOptions{}

	// Use extension_id from the resource
	extensionId := d.Get("extension_id").(string)
	updateExtensionDeploymentOptions.SetID(extensionId)

	hasChange := false

	if d.HasChange("version") || d.HasChange("item_ids") {
		updateExtensionDeploymentOptions.SetVersion(d.Get("version").(string))
		var newItemIds []string
		for _, v := range d.Get("item_ids").([]interface{}) {
			newItemIdsItem := v.(string)
			newItemIds = append(newItemIds, newItemIdsItem)
		}
		updateExtensionDeploymentOptions.SetItemIds(newItemIds)
		hasChange = true
	}
	if d.HasChange("applications") {
		var newApplications []string
		for _, v := range d.Get("applications").([]interface{}) {
			newApplicationsItem := v.(string)
			newApplications = append(newApplications, newApplicationsItem)
		}
		updateExtensionDeploymentOptions.SetApplications(newApplications)
		hasChange = true
	}
	if d.HasChange("subsystems") {
		var newSubsystems []string
		for _, v := range d.Get("subsystems").([]interface{}) {
			newSubsystemsItem := v.(string)
			newSubsystems = append(newSubsystems, newSubsystemsItem)
		}
		updateExtensionDeploymentOptions.SetSubsystems(newSubsystems)
		hasChange = true
	}
	// Sub-resource: id field not updatable

	if hasChange {
		_, _, err := logsClient.UpdateExtensionDeploymentWithContext(context, updateExtensionDeploymentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateExtensionDeploymentWithContext failed: %s", err.Error()), "ibm_logs_extension_deployment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmLogsExtensionDeploymentRead(context, d, meta)
}

func resourceIbmLogsExtensionDeploymentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_extension_deployment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, _, _, extensionId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("updateClientURLWithInstanceEndpoint failed: %s", err.Error()), "ibm_logs_extension_deployment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteExtensionDeploymentOptions := &logsv0.DeleteExtensionDeploymentOptions{}
	deleteExtensionDeploymentOptions.SetID(extensionId)

	_, _, err = logsClient.DeleteExtensionDeploymentWithContext(context, deleteExtensionDeploymentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteExtensionDeploymentWithContext failed: %s", err.Error()), "ibm_logs_extension_deployment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
