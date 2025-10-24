// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice

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
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func ResourceIbmPdrManagedr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmPdrManagedrCreate,
		ReadContext:   resourceIbmPdrManagedrRead,
		DeleteContext: resourceIbmPdrManagedrDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "instance id of instance to provision.",
			},
			"stand_by_redeploy": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "stand_by_redeploy"),
				Description:  "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\").",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"accepts_incomplete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.",
			},
			"dashboard_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the dashboard for managing the DR service instance in IBM Cloud.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN (Cloud Resource Name) of the DR service instance.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"api_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				ForceNew:  true,
				Optional:  true,
			},
			"orchestrator_ha": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"guid": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"location_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"machine_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"orchestrator_location_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"orchestrator_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"orchestrator_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				ForceNew:  true,
				Optional:  true,
			},
			"orchestrator_workspace_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"orchestrator_workspace_location": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"proxy_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"resource_instance": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"schematic_workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secondary_workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secret": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secret_group": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ssh_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ssh_public_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_machine_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_orchestrator_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_orchestrator_workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_orchestrator_workspace_location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_schematic_workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standby_tier": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"transit_gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func ResourceIbmPdrManagedrValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "stand_by_redeploy",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "false, true",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pdr_managedr", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmPdrManagedrCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createManageDrOptions := &drautomationservicev1.CreateManageDrOptions{}

	createManageDrOptions.SetInstanceID(d.Get("instance_id").(string))
	createManageDrOptions.SetStandByRedeploy(d.Get("stand_by_redeploy").(string))
	if _, ok := d.GetOk("action"); ok {
		createManageDrOptions.SetAction(d.Get("action").(string))
	}
	if _, ok := d.GetOk("api_key"); ok {
		createManageDrOptions.SetAPIKey(d.Get("api_key").(string))
	}
	if v, ok := d.GetOk("orchestrator_ha"); ok {
		enableHA := v.(bool)
		createManageDrOptions.OrchestratorHa = core.BoolPtr(enableHA)
	}
	if _, ok := d.GetOk("guid"); ok {
		createManageDrOptions.SetGUID(d.Get("guid").(string))
	}
	if _, ok := d.GetOk("location_id"); ok {
		createManageDrOptions.SetLocationID(d.Get("location_id").(string))
	}
	if _, ok := d.GetOk("machine_type"); ok {
		createManageDrOptions.SetMachineType(d.Get("machine_type").(string))
	}
	if _, ok := d.GetOk("orchestrator_location_type"); ok {
		createManageDrOptions.SetOrchestratorLocationType(d.Get("orchestrator_location_type").(string))
	}
	if _, ok := d.GetOk("orchestrator_name"); ok {
		createManageDrOptions.SetOrchestratorName(d.Get("orchestrator_name").(string))
	}
	if _, ok := d.GetOk("orchestrator_password"); ok {
		createManageDrOptions.SetOrchestratorPassword(d.Get("orchestrator_password").(string))
	}
	if _, ok := d.GetOk("orchestrator_workspace_id"); ok {
		createManageDrOptions.SetOrchestratorWorkspaceID(d.Get("orchestrator_workspace_id").(string))
	}
	if _, ok := d.GetOk("orchestrator_workspace_location"); ok {
		createManageDrOptions.SetOrchestratorWorkspaceLocation(d.Get("orchestrator_workspace_location").(string))
	}
	if _, ok := d.GetOk("proxy_ip"); ok {
		createManageDrOptions.SetProxyIP(d.Get("proxy_ip").(string))
	}
	if _, ok := d.GetOk("region_id"); ok {
		createManageDrOptions.SetRegionID(d.Get("region_id").(string))
	}
	if _, ok := d.GetOk("resource_instance"); ok {
		createManageDrOptions.SetResourceInstance(d.Get("resource_instance").(string))
	}
	if _, ok := d.GetOk("schematic_workspace_id"); ok {
		createManageDrOptions.SetSchematicWorkspaceID(d.Get("schematic_workspace_id").(string))
	}
	if _, ok := d.GetOk("secondary_workspace_id"); ok {
		createManageDrOptions.SetSecondaryWorkspaceID(d.Get("secondary_workspace_id").(string))
	}
	if _, ok := d.GetOk("secret"); ok {
		createManageDrOptions.SetSecret(d.Get("secret").(string))
	}
	if _, ok := d.GetOk("secret_group"); ok {
		createManageDrOptions.SetSecretGroup(d.Get("secret_group").(string))
	}
	if _, ok := d.GetOk("ssh_key_name"); ok {
		createManageDrOptions.SetSSHKeyName(d.Get("ssh_key_name").(string))
	}
	if _, ok := d.GetOk("standby_machine_type"); ok {
		createManageDrOptions.SetStandbyMachineType(d.Get("standby_machine_type").(string))
	}
	if _, ok := d.GetOk("standby_orchestrator_name"); ok {
		createManageDrOptions.SetStandbyOrchestratorName(d.Get("standby_orchestrator_name").(string))
	}
	if _, ok := d.GetOk("standby_orchestrator_workspace_id"); ok {
		createManageDrOptions.SetStandbyOrchestratorWorkspaceID(d.Get("standby_orchestrator_workspace_id").(string))
	}
	if _, ok := d.GetOk("standby_orchestrator_workspace_location"); ok {
		createManageDrOptions.SetStandbyOrchestratorWorkspaceLocation(d.Get("standby_orchestrator_workspace_location").(string))
	}
	if _, ok := d.GetOk("standby_schematic_workspace_id"); ok {
		createManageDrOptions.SetStandbySchematicWorkspaceID(d.Get("standby_schematic_workspace_id").(string))
	}
	if _, ok := d.GetOk("standby_tier"); ok {
		createManageDrOptions.SetStandbyTier(d.Get("standby_tier").(string))
	}
	if _, ok := d.GetOk("tier"); ok {
		createManageDrOptions.SetTier(d.Get("tier").(string))
	}
	if _, ok := d.GetOk("transit_gateway_id"); ok {
		createManageDrOptions.SetTransitGatewayID(d.Get("transit_gateway_id").(string))
	}
	if _, ok := d.GetOk("vpc_id"); ok {
		createManageDrOptions.SetVPCID(d.Get("vpc_id").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		createManageDrOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}
	if _, ok := d.GetOk("accepts_incomplete"); ok {
		createManageDrOptions.SetAcceptsIncomplete(d.Get("accepts_incomplete").(bool))
	}

	serviceInstanceManageDr, _, err := drAutomationServiceClient.CreateManageDrWithContext(context, createManageDrOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateManageDrWithContext failed: %s", err.Error()), "ibm_pdr_managedr", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createManageDrOptions.InstanceID, *serviceInstanceManageDr.ID))

	return resourceIbmPdrManagedrRead(context, d, meta)
}

func resourceIbmPdrManagedrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getManageDrOptions := &drautomationservicev1.GetManageDrOptions{}

	// parts, err := flex.SepIdParts(d.Id(), "/")
	// if err != nil {
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "sep-id-parts").GetDiag()
	// }
	instanceID := d.Get("instance_id").(string)

	log.Printf("[DEBUG] Read operation using instance ID from resource: %s", instanceID)

	getManageDrOptions.SetInstanceID(instanceID)

	// getManageDrOptions.SetInstanceID(parts[0])
	// getManageDrOptions.SetInstanceID(parts[1])
	if _, ok := d.GetOk("accept_language"); ok {
		getManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getManageDrOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	serviceInstanceManageDr, response, err := drAutomationServiceClient.GetManageDrWithContext(context, getManageDrOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetManageDrWithContext failed: %s", err.Error()), "ibm_pdr_managedr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(serviceInstanceManageDr.DashboardURL) {
		if err = d.Set("dashboard_url", serviceInstanceManageDr.DashboardURL); err != nil {
			err = fmt.Errorf("Error setting dashboard_url: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-dashboard_url").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ID) {
		if err = d.Set("instance_id", serviceInstanceManageDr.ID); err != nil {
			err = fmt.Errorf("Error setting instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pdr_managedr", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmPdrManagedrDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
