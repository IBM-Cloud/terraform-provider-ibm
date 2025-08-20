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

func DataSourceIbmPdrGetDeploymentStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetDeploymentStatusRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"orch_ext_connectivity_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orch_standby_node_addition_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_cluster_message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_cluster_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_config_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_group_leader": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"orchestrator_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schematic_workspace_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schematic_workspace_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_key_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"standby_orchestrator_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"standby_orchestrator_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIbmPdrGetDeploymentStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_deployment_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceDrdeploymentOptions := &drautomationservicev1.ServiceInstanceDrdeploymentOptions{}

	serviceInstanceDrdeploymentOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceDrdeploymentOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	deploymentStatus, _, err := drAutomationServiceClient.ServiceInstanceDrdeploymentWithContext(context, serviceInstanceDrdeploymentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceDrdeploymentWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_deployment_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetDeploymentStatusID(d))

	if !core.IsNil(deploymentStatus.OrchExtConnectivityStatus) {
		if err = d.Set("orch_ext_connectivity_status", deploymentStatus.OrchExtConnectivityStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_ext_connectivity_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orch_ext_connectivity_status").GetDiag()
		}
	}

	if err = d.Set("orch_standby_node_addition_status", deploymentStatus.OrchStandbyNodeAdditionStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orch_standby_node_addition_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orch_standby_node_addition_status").GetDiag()
	}

	if err = d.Set("orchestrator_cluster_message", deploymentStatus.OrchestratorClusterMessage); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_cluster_message: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_cluster_message").GetDiag()
	}

	if err = d.Set("orchestrator_cluster_type", deploymentStatus.OrchestratorClusterType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_cluster_type: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_cluster_type").GetDiag()
	}

	if err = d.Set("orchestrator_config_status", deploymentStatus.OrchestratorConfigStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_config_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_config_status").GetDiag()
	}

	if err = d.Set("orchestrator_group_leader", deploymentStatus.OrchestratorGroupLeader); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_group_leader: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_group_leader").GetDiag()
	}

	if err = d.Set("orchestrator_name", deploymentStatus.OrchestratorName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_name: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_name").GetDiag()
	}

	if err = d.Set("orchestrator_status", deploymentStatus.OrchestratorStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-orchestrator_status").GetDiag()
	}

	if err = d.Set("schematic_workspace_name", deploymentStatus.SchematicWorkspaceName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting schematic_workspace_name: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-schematic_workspace_name").GetDiag()
	}

	if err = d.Set("schematic_workspace_status", deploymentStatus.SchematicWorkspaceStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting schematic_workspace_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-schematic_workspace_status").GetDiag()
	}

	if err = d.Set("ssh_key_name", deploymentStatus.SshKeyName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ssh_key_name: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-ssh_key_name").GetDiag()
	}

	if err = d.Set("standby_orchestrator_name", deploymentStatus.StandbyOrchestratorName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_orchestrator_name: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-standby_orchestrator_name").GetDiag()
	}

	if err = d.Set("standby_orchestrator_status", deploymentStatus.StandbyOrchestratorStatus); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_orchestrator_status: %s", err), "(Data) ibm_pdr_get_deployment_status", "read", "set-standby_orchestrator_status").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrGetDeploymentStatusID returns a reasonable ID for the list.
func dataSourceIbmPdrGetDeploymentStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
