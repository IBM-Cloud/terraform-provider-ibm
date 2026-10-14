// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"

	// "github.com/IBM/dra-go-sdk/drautomationservicev1"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrLastOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrLastOperationRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service Instance ID.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"primary_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the progress details of primary orchestrator creation.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the primary orchestrator.",
			},
			"standby_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the progress details of primary orchestrator creation.",
			},
			"standby_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the standby orchestrator.",
			},
			"orchestrator_ha": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether high availability (HA) is enabled for the orchestrator.",
			},
			"deployment_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the service instance deployment.",
			},
			"recovery_location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The disaster recovery location associated with the instance.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group to which the service instance belongs.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service instance crn.",
			},
			"primary_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the primary orchestrator VM.",
			},
			"standby_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the standby orchestrator VM.",
			},
			"orchestrator_config_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration status of the orchestrator cluster.",
			},
			"primary_orchestrator_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration status of the orchestrator cluster.",
			},
			"orchestrator_cluster_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the primary orchestrator VM.",
			},
			"orch_standby_node_addition_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of standby node in the Orchestrator cluster.",
			},
			"orch_ext_connectivity_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of standby node addition to the orchestrator cluster.",
			},
			"last_updated_orchestrator_deployment_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deployment time of primary orchestrator VM.",
			},
			"last_updated_standby_orchestrator_deployment_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deployment time of StandBy orchestrator VM.",
			},
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the DR Automation plan.",
			},
			"mfa_enabled": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicated whether multi factor authentication is ennabled or not.",
			},
			"primary_error_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Capture the error while creating primary orchestrator.",
			},
			"standby_error_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Capture the error while creating standby orchestrator.",
			},
			"is_api_key_expired": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the API key used for the deployment is expired.",
			},
		},
	}
}

func dataSourceIBMPdrLastOperationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_last_operation", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getLastOperationOptions := &drautomationservicev1.GetLastOperationOptions{}

	getLastOperationOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getLastOperationOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	serviceInstanceStatus, response, err := drAutomationServiceClient.GetLastOperationWithContext(context, getLastOperationOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetLastOperationWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetLastOperationWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_last_operation", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrLastOperationID(d))

	if !core.IsNil(serviceInstanceStatus.PrimaryDescription) {
		if err = d.Set("primary_description", serviceInstanceStatus.PrimaryDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_description").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.Status) {
		if err = d.Set("status", serviceInstanceStatus.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.StandbyDescription) {
		if err = d.Set("standby_description", serviceInstanceStatus.StandbyDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_description").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.StandbyStatus) {
		if err = d.Set("standby_status", serviceInstanceStatus.StandbyStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.OrchestratorHa) {
		if err = d.Set("orchestrator_ha", serviceInstanceStatus.OrchestratorHa); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_ha: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_ha").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.DeploymentName) {
		if err = d.Set("deployment_name", serviceInstanceStatus.DeploymentName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deployment_name: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-deployment_name").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.RecoveryLocation) {
		if err = d.Set("recovery_location", serviceInstanceStatus.RecoveryLocation); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting recovery_location: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-recovery_location").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.ResourceGroup) {
		if err = d.Set("resource_group", serviceInstanceStatus.ResourceGroup); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-resource_group").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.CRN) {
		if err = d.Set("crn", serviceInstanceStatus.CRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-crn").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.PrimaryIPAddress) {
		if err = d.Set("primary_ip_address", serviceInstanceStatus.PrimaryIPAddress); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_ip_address: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_ip_address").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.StandbyIPAddress) {
		if err = d.Set("standby_ip_address", serviceInstanceStatus.StandbyIPAddress); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_ip_address: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_ip_address").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.OrchestratorConfigStatus) {
		if err = d.Set("orchestrator_config_status", serviceInstanceStatus.OrchestratorConfigStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_config_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_config_status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.PrimaryOrchestratorStatus) {
		if err = d.Set("primary_orchestrator_status", serviceInstanceStatus.PrimaryOrchestratorStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_orchestrator_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_orchestrator_status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.OrchestratorClusterMessage) {
		if err = d.Set("orchestrator_cluster_message", serviceInstanceStatus.OrchestratorClusterMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_cluster_message: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_cluster_message").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.OrchStandbyNodeAdditionStatus) {
		if err = d.Set("orch_standby_node_addition_status", serviceInstanceStatus.OrchStandbyNodeAdditionStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_standby_node_addition_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orch_standby_node_addition_status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.OrchExtConnectivityStatus) {
		if err = d.Set("orch_ext_connectivity_status", serviceInstanceStatus.OrchExtConnectivityStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_ext_connectivity_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orch_ext_connectivity_status").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.LastUpdatedOrchestratorDeploymentTime) {
		if err = d.Set("last_updated_orchestrator_deployment_time", flex.DateTimeToString(serviceInstanceStatus.LastUpdatedOrchestratorDeploymentTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_updated_orchestrator_deployment_time: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-last_updated_orchestrator_deployment_time").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.LastUpdatedStandbyOrchestratorDeploymentTime) {
		if err = d.Set("last_updated_standby_orchestrator_deployment_time", flex.DateTimeToString(serviceInstanceStatus.LastUpdatedStandbyOrchestratorDeploymentTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_updated_standby_orchestrator_deployment_time: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-last_updated_standby_orchestrator_deployment_time").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.PlanName) {
		if err = d.Set("plan_name", serviceInstanceStatus.PlanName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan_name: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-plan_name").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.MfaEnabled) {
		if err = d.Set("mfa_enabled", serviceInstanceStatus.MfaEnabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mfa_enabled: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-mfa_enabled").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.PrimaryErrorDescription) {
		if err = d.Set("primary_error_description", serviceInstanceStatus.PrimaryErrorDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_error_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_error_description").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.StandbyErrorDescription) {
		if err = d.Set("standby_error_description", serviceInstanceStatus.StandbyErrorDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_error_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_error_description").GetDiag()
		}
	}

	if !core.IsNil(serviceInstanceStatus.IsAPIKeyExpired) {
		if err = d.Set("is_api_key_expired", serviceInstanceStatus.IsAPIKeyExpired); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_api_key_expired: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-is_api_key_expired").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrLastOperationID returns a reasonable ID for the list.
func dataSourceIBMPdrLastOperationID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}
