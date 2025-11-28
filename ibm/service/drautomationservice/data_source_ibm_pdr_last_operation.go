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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrLastOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrLastOperationRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service instance crn.",
			},
			"deployment_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the service instance deployment.",
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
			"mfa_enabled": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicated whether multi factor authentication is ennabled or not.",
			},
			"orch_ext_connectivity_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of standby node addition to the orchestrator cluster.",
			},
			"orch_standby_node_addtion_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of standby node in the Orchestrator cluster.",
			},
			"orchestrator_cluster_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the primary orchestrator VM.",
			},
			"orchestrator_config_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration status of the orchestrator cluster.",
			},
			"orchestrator_ha": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether high availability (HA) is enabled for the orchestrator.",
			},
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the DR Automation plan.",
			},
			"primary_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the progress details of primary orchestrator creation.",
			},
			"primary_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the primary orchestrator VM.",
			},
			"primary_orchestrator_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration status of the orchestrator cluster.",
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
			"standby_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the progress details of primary orchestrator creation.",
			},
			"standby_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the standby orchestrator VM.",
			},
			"standby_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the standby orchestrator.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the primary orchestrator.",
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
	// if _, ok := d.GetOk("if_none_match"); ok {
	// 	getLastOperationOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	// }

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

	if err = d.Set("crn", serviceInstanceStatus.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-crn").GetDiag()
	}

	if err = d.Set("deployment_name", serviceInstanceStatus.DeploymentName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deployment_name: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-deployment_name").GetDiag()
	}

	if err = d.Set("last_updated_orchestrator_deployment_time", flex.DateTimeToString(serviceInstanceStatus.LastUpdatedOrchestratorDeploymentTime)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_updated_orchestrator_deployment_time: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-last_updated_orchestrator_deployment_time").GetDiag()
	}

	if err = d.Set("last_updated_standby_orchestrator_deployment_time", flex.DateTimeToString(serviceInstanceStatus.LastUpdatedStandbyOrchestratorDeploymentTime)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_updated_standby_orchestrator_deployment_time: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-last_updated_standby_orchestrator_deployment_time").GetDiag()
	}

	if err = d.Set("mfa_enabled", serviceInstanceStatus.MfaEnabled); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mfa_enabled: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-mfa_enabled").GetDiag()
	}

	if !core.IsNil(serviceInstanceStatus.OrchExtConnectivityStatus) {
		if err = d.Set("orch_ext_connectivity_status", serviceInstanceStatus.OrchExtConnectivityStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_ext_connectivity_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orch_ext_connectivity_status").GetDiag()
		}
	}

	if err = d.Set("orch_standby_node_addtion_status", serviceInstanceStatus.OrchStandbyNodeAddtionStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_standby_node_addtion_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orch_standby_node_addtion_status").GetDiag()
	}

	if err = d.Set("orchestrator_cluster_message", serviceInstanceStatus.OrchestratorClusterMessage); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_cluster_message: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_cluster_message").GetDiag()
	}

	if err = d.Set("orchestrator_config_status", serviceInstanceStatus.OrchestratorConfigStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_config_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_config_status").GetDiag()
	}

	if err = d.Set("orchestrator_ha", serviceInstanceStatus.OrchestratorHa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_ha: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-orchestrator_ha").GetDiag()
	}

	if err = d.Set("plan_name", serviceInstanceStatus.PlanName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan_name: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-plan_name").GetDiag()
	}

	if err = d.Set("primary_description", serviceInstanceStatus.PrimaryDescription); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_description").GetDiag()
	}

	if err = d.Set("primary_ip_address", serviceInstanceStatus.PrimaryIPAddress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_ip_address: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_ip_address").GetDiag()
	}

	if err = d.Set("primary_orchestrator_status", serviceInstanceStatus.PrimaryOrchestratorStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_orchestrator_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-primary_orchestrator_status").GetDiag()
	}

	if err = d.Set("recovery_location", serviceInstanceStatus.RecoveryLocation); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting recovery_location: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-recovery_location").GetDiag()
	}

	if err = d.Set("resource_group", serviceInstanceStatus.ResourceGroup); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-resource_group").GetDiag()
	}

	if err = d.Set("standby_description", serviceInstanceStatus.StandbyDescription); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_description: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_description").GetDiag()
	}

	if err = d.Set("standby_ip_address", serviceInstanceStatus.StandbyIPAddress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_ip_address: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_ip_address").GetDiag()
	}

	if err = d.Set("standby_status", serviceInstanceStatus.StandbyStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-standby_status").GetDiag()
	}

	if err = d.Set("status", serviceInstanceStatus.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-status").GetDiag()
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
