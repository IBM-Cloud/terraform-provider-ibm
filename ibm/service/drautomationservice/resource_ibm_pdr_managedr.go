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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func ResourceIbmPdrManagedr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmPdrManagedrCreate,
		ReadContext:   resourceIbmPdrManagedrRead,
		DeleteContext: resourceIbmPdrManagedrDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "instance id of instance to provision.",
			},
			"stand_by_redeploy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					allowed := []string{"true", "false"}
					for _, a := range allowed {
						if v == a {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be one of %v, got: %s", key, allowed, v))
					return
				},
				Description: "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\").",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The language requested for the return document.",
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
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Indicates whether to proceed with asynchronous operation after all configuration details are updated in the database.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Sensitive:   true,
				ForceNew:    true,
				Optional:    true,
				Description: "The API key associated with the IBM Cloud service instance.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Client ID for MFA Authentication.",
			},
			"proxy_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Proxy IP for the Communication between Orchestrator and Service.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Client Secret for MFA Authentication.",
			},
			"tenant_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Tenant Name for MFA Authentication.",
			},
			"orchestrator_ha": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Flag to enable or disable High Availability (HA) for the service instance.",
			},
			"guid": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "The globally unique identifier of the service instance.",
			},
			"location_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The Location or data center identifier where the service instance is deployed.",
			},
			"machine_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Machine type or flavor used for virtual machines in the service instance.",
			},
			"orchestrator_location_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Type of orchestrator cluster used in the service instance.",
			},
			"orchestrator_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username for the orchestrator management interface.",
			},
			"orchestrator_password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				ForceNew:    true,
				Required:    true,
				Description: "Password for the orchestrator management interface.",
			},
			"orchestrator_workspace_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the orchestrator workspace.",
			},
			"region_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Cloud region where the service instance is deployed.",
			},
			"resource_instance": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "ID of the associated IBM Cloud resource instance.",
			},
			"secret": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Secret name or identifier used for retrieving credentials from Secrets Manager.",
			},
			"secret_group": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Secret group name in IBM Cloud Secrets Manager containing sensitive data for * the service instance.",
			},
			"ssh_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Name of the SSH key stored in the cloud provider.",
			},
			"standby_machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Machine type or flavor used for standby virtual machines.",
			},
			"standby_orchestrator_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Username for the standby orchestrator management interface.",
			},
			"standby_orchestrator_workspace_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the standby orchestrator workspace.",
			},
			"standby_tier": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Tier of the standby service instance.",
			},
			"tier": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Tier of the service instance.",
			},
			"standby_ssh_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "standy ssh key name of the service instance.",
			},
			"orchestrator_network_ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.All(
						validation.StringLenBetween(36, 64),
						validation.StringMatch(
							regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`),
							"must contain only alphanumeric characters, hyphen, or underscore",
						),
					),
				},
				Description: "List of network IDs for primary orchestrator VM.",
			},
			"standby_orchestrator_network_ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.All(
						validation.StringLenBetween(36, 64),
						validation.StringMatch(
							regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`),
							"must contain only alphanumeric characters, hyphen, or underscore",
						),
					),
				},
				Description: "List of network IDs for standby orchestrator VM.",
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

func resourceIbmPdrManagedrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ctx, cancel := context.WithTimeout(ctx, 90*time.Minute)
	defer cancel()

	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createManageDrOptions := &drautomationservicev1.CreateManageDrOptions{}

	createManageDrOptions.SetInstanceID(d.Get("instance_id").(string))
	createManageDrOptions.SetStandByRedeploy(d.Get("stand_by_redeploy").(string))
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
	if _, ok := d.GetOk("region_id"); ok {
		createManageDrOptions.SetRegionID(d.Get("region_id").(string))
	}
	if _, ok := d.GetOk("client_id"); ok {
		createManageDrOptions.SetClientID(d.Get("client_id").(string))
	}
	if _, ok := d.GetOk("proxy_ip"); ok {
		createManageDrOptions.SetProxyIP(d.Get("proxy_ip").(string))
	}
	if _, ok := d.GetOk("client_secret"); ok {
		createManageDrOptions.SetClientSecret(d.Get("client_secret").(string))
	}
	if _, ok := d.GetOk("tenant_name"); ok {
		createManageDrOptions.SetTenantName(d.Get("tenant_name").(string))
	}
	if _, ok := d.GetOk("resource_instance"); ok {
		createManageDrOptions.SetResourceInstance(d.Get("resource_instance").(string))
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
	if _, ok := d.GetOk("standby_tier"); ok {
		createManageDrOptions.SetStandbyTier(d.Get("standby_tier").(string))
	}
	if _, ok := d.GetOk("tier"); ok {
		createManageDrOptions.SetTier(d.Get("tier").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("accepts_incomplete"); ok {
		createManageDrOptions.SetAcceptsIncomplete(d.Get("accepts_incomplete").(bool))
	}
	if _, ok := d.GetOk("standby_ssh_key_name"); ok {
		createManageDrOptions.SetStandbySSHKeyName(d.Get("standby_ssh_key_name").(string))
	}
	if _, ok := d.GetOk("orchestrator_network_ids"); ok {
		var orchestratorNetworkIds []string
		for _, v := range d.Get("orchestrator_network_ids").([]interface{}) {
			orchestratorNetworkIdsItem := v.(string)
			orchestratorNetworkIds = append(orchestratorNetworkIds, orchestratorNetworkIdsItem)
		}
		createManageDrOptions.SetOrchestratorNetworkIds(orchestratorNetworkIds)
	}
	if _, ok := d.GetOk("standby_orchestrator_network_ids"); ok {
		var standbyOrchestratorNetworkIds []string
		for _, v := range d.Get("standby_orchestrator_network_ids").([]interface{}) {
			standbyOrchestratorNetworkIdsItem := v.(string)
			standbyOrchestratorNetworkIds = append(standbyOrchestratorNetworkIds, standbyOrchestratorNetworkIdsItem)
		}
		createManageDrOptions.SetStandbyOrchestratorNetworkIds(standbyOrchestratorNetworkIds)
	}

	_, response, err := drAutomationServiceClient.CreateManageDrWithContext(ctx, createManageDrOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateManageDrWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateManageDrWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_managedr", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(d.Get("instance_id").(string))
	// Step 2: Poll Last Operation status every 5 minutes until Active or Fail
	instanceID := *createManageDrOptions.InstanceID
	const (
		pollInterval = 1 * time.Minute
		maxWaitTime  = 90 * time.Minute // optional, can extend as needed
	)
	timeout := time.After(maxWaitTime)
	ticker := time.NewTicker(pollInterval)
	// defer ticker.Stop()

	log.Printf("[INFO] Started polling last operation status for instance %s every %s", instanceID, pollInterval)
	enableha, _ := d.GetOk("orchestrator_ha")

	if !enableha.(bool) {
		for {
			select {
			case <-timeout:
				errMsg := fmt.Sprintf("Timeout exceeded while waiting for Manage DR to become Active (instance_id: %s)", instanceID)
				tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")
				log.Printf("[ERROR] %s", errMsg)
				ticker.Stop()
				return tfErr.GetDiag()

			case <-ticker.C:
				status, _, _, statusErr, _ := checkLastOperationStatus(ctx, drAutomationServiceClient, instanceID)
				if statusErr != nil {
					tfErr := flex.TerraformErrorf(statusErr, fmt.Sprintf("GetLastOperation failed: %s", statusErr.Error()), "ibm_pdr_managedr", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					ticker.Stop()
					return tfErr.GetDiag()
				}

				log.Printf("[INFO] Current Last Operation status for instance %s: %s", instanceID, *status.Status)

				switch strings.ToLower(*status.Status) {
				case "active":
					log.Printf("[INFO] Manage DR operation completed successfully for instance %s", instanceID)
					ticker.Stop()
					return resourceIbmPdrManagedrRead(ctx, d, meta)

				case "fail", "failed", "error":
					errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.PrimaryDescription)
					tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

					log.Printf("[ERROR] %s", errMsg)
					ticker.Stop()
					return tfErr.GetDiag()

				default:
					log.Printf("[DEBUG] Manage DR still in progress... retrying in %v", pollInterval)
				}
			}
		}
	} else {
		for {
			select {
			case <-timeout:
				errMsg := fmt.Sprintf("Timeout exceeded while waiting for Manage DR to become Active (instance_id: %s)", instanceID)
				tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")
				log.Printf("[ERROR] %s", errMsg)
				ticker.Stop()
				return tfErr.GetDiag()

			case <-ticker.C:
				status, _, standbyStatus, statusErr, _ := checkLastOperationStatus(ctx, drAutomationServiceClient, instanceID)
				if statusErr != nil {
					tfErr := flex.TerraformErrorf(statusErr, fmt.Sprintf("GetLastOperation failed: %s", statusErr.Error()), "ibm_pdr_managedr", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					ticker.Stop()
					return tfErr.GetDiag()
				}

				switch strings.ToLower(*status.Status) {
				case "active":
					if strings.ToLower(standbyStatus) == "active" {
						log.Printf("[INFO] Manage DR operation completed successfully for instance %s (both primary and standby active)", instanceID)
						ticker.Stop()
						return resourceIbmPdrManagedrRead(ctx, d, meta)
					}
					if strings.ToLower(standbyStatus) == "failed" {
						errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.StandbyDescription)
						tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

						log.Printf("[ERROR] %s", errMsg)
						ticker.Stop()
						return tfErr.GetDiag()
					}

					// If standby still initializing
					log.Printf("[INFO] Manage DR overall status is Active, but standby orchestrator still in progress (status: %s). Retrying in %v...",
						standbyStatus, pollInterval)
					continue

				case "fail", "failed", "error":
					errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.PrimaryDescription)
					tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

					log.Printf("[ERROR] %s", errMsg)
					ticker.Stop()
					return tfErr.GetDiag()

				default:
					log.Printf("[DEBUG] Manage DR still in progress ... retrying in %v", pollInterval)
				}
			}
		}
	}

	// return resourceIbmPdrManagedrRead(context, d, meta)

}

func resourceIbmPdrManagedrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getManageDrOptions := &drautomationservicev1.GetManageDrOptions{}

	instanceID := d.Id()

	log.Printf("[DEBUG] Read operation using instance ID from resource: %s", instanceID)

	getManageDrOptions.SetInstanceID(instanceID)

	// if _, ok := d.GetOk("accept_language"); ok {
	// 	getManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	// }

	serviceInstanceManageDr, response, err := drAutomationServiceClient.GetManageDrWithContext(context, getManageDrOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetManageDrWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetManageDrWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_managedr", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	if !core.IsNil(serviceInstanceManageDr.DashboardURL) {
		if err = d.Set("dashboard_url", serviceInstanceManageDr.DashboardURL); err != nil {
			err = fmt.Errorf("Error setting dashboard_url: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-dashboard_url").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ID) {
		if err = d.Set("instance_id", d.Id()); err != nil {
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

func checkLastOperationStatus(ctx context.Context, client *drautomationservicev1.DrAutomationServiceV1, instanceID string) (*drautomationservicev1.ServiceInstanceStatus, string, string, error, error) {
	opts := &drautomationservicev1.GetLastOperationOptions{}
	opts.SetInstanceID(instanceID)

	statusResponse, _, err := client.GetLastOperationWithContext(ctx, opts)
	if err != nil {
		return nil, "", "", err, nil
	}

	if statusResponse.Status == nil {
		return statusResponse, "", "", fmt.Errorf("received nil status for instance %s", instanceID), nil
	}

	status := strings.ToLower(*statusResponse.Status)
	primaryStatus := strings.ToLower(*statusResponse.PrimaryOrchestratorStatus)
	standbyStatus := strings.ToLower(*statusResponse.StandbyStatus)

	// --- Custom error logic based on your conditions ---
	if status == "failed" {
		switch {
		case primaryStatus == "failed" && standbyStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s \n %s", *statusResponse.PrimaryDescription, *statusResponse.StandbyDescription)
		case primaryStatus == "failed" && (standbyStatus == "" || standbyStatus == "na"):
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s", *statusResponse.PrimaryDescription)
		case primaryStatus == "active" && (standbyStatus != "" || standbyStatus == "failed"):
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s \n %s", *statusResponse.PrimaryDescription, *statusResponse.StandbyDescription)
		case primaryStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("primary orchestrator failed for instance %s", instanceID)
		case standbyStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("standby orchestrator failed for instance %s", instanceID)
		default:
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("operation failed for instance %s with unknown cause", instanceID)
		}
	}

	return statusResponse, primaryStatus, standbyStatus, nil, nil
}
