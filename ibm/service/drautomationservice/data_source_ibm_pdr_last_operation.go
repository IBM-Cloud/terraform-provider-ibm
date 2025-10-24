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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIbmPdrLastOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrLastOperationRead,

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
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name (CRN) of the service instance.",
			},
			"deployment_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the service instance deployment.",
			},
			"is_ksys_ha": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether high availability (HA) is enabled for the orchestrator.",
			},
			"orch_ext_connectivity_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of standby node addition to the orchestrator cluster.",
			},
			"orch_standby_node_addtion_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Health or informational message about the orchestrator cluster.",
			},
			"orchestrator_cluster_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the primary orchestrator VM.",
			},
			"orchestrator_config_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration status of the orchestrator cluster.",
			},
			"primary_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Detailed status message for the primary orchestrator VM.",
			},
			"primary_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address of the primary orchestrator VM.",
			},
			"primary_orchestrator_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration status of the orchestrator cluster.",
			},
			"recovery_location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Disaster recovery location associated with the instance.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group to which the service instance belongs.",
			},
			"standby_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Detailed status message for the standby orchestrator VM.",
			},
			"standby_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address of the standby orchestrator VM.",
			},
			"standby_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current state of the standby orchestrator VM.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Overall status of the service instance.",
			},
		},
	}
}

func dataSourceIbmPdrLastOperationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if _, ok := d.GetOk("if_none_match"); ok {
		getLastOperationOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	serviceInstanceStatus, _, err := drAutomationServiceClient.GetLastOperationWithContext(context, getLastOperationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLastOperationWithContext failed: %s", err.Error()), "(Data) ibm_pdr_last_operation", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrLastOperationID(d))

	if err = d.Set("crn", serviceInstanceStatus.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-crn").GetDiag()
	}

	if err = d.Set("deployment_name", serviceInstanceStatus.DeploymentName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deployment_name: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-deployment_name").GetDiag()
	}

	if err = d.Set("is_ksys_ha", serviceInstanceStatus.IsKsysHa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_ksys_ha: %s", err), "(Data) ibm_pdr_last_operation", "read", "set-is_ksys_ha").GetDiag()
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

// dataSourceIbmPdrLastOperationID returns a reasonable ID for the list.
func dataSourceIbmPdrLastOperationID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
