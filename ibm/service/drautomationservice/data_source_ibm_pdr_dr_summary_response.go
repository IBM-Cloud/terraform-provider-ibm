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
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIBMPdrDrSummaryResponse() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrDrSummaryResponseRead,

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

			"managed_vm_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vm_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"core": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dr_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dr_average_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"workspace_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"workgroup_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"service_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the DR automation service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status of the service.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Service description.",
						},
						"orchestrator_ha": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The flag indicating whether orchestartor HA is enabled.",
						},
						"deployment_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the deployment.",
						},
						"recovery_location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The disaster recovery location.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Resource group name.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deployment crn.",
						},
						"primary_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service Orchestator primary IP address.",
						},
						"standby_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service Orchestator standby IP address.",
						},
						"standby_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The standby orchestrator current status details.",
						},
						"standby_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The standby orchestrator current status.",
						},
						"primary_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Primary Orchestrator Dashboard URL.",
						},
						"standby_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Standby Orchestrator Dashboard URL.",
						},
						"plan_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The plan name.",
						},
					},
				},
			},
			"orchestrator_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the orchestrator configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the primary orchestrator.",
						},
						"orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the primary orchestrator.",
						},
						"ssh_key_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH key name used for the orchestrator.",
						},
						"standby_ssh_key_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH key name used for the standby orchestrator.",
						},
						"schematic_workspace_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the schematic workspace.",
						},
						"schematic_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the schematic workspace.",
						},
						"standby_orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the standby orchestrator.",
						},
						"standby_orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the standby orchestrator.",
						},
						"orchestrator_config_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration status of the orchestrator.",
						},
						"orchestrator_group_leader": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The leader node of the orchestrator group.",
						},
						"orchestrator_cluster_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message regarding orchestrator cluster status.",
						},
						"orch_standby_node_addition_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of standby node addition.",
						},
						"orchestrator_location_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of orchestrator Location.",
						},
						"location_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of location.",
						},
						"vpc_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the VPC.",
						},
						"transit_gateway_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit gateway.",
						},
						"proxy_ip": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the proxy.",
						},
						"orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the orchestrator workspace.",
						},
						"standby_orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the standby orchestrator workspace.",
						},
						"orch_ext_connectivity_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The external connectivity status of the orchestrator.",
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
						"latest_orchestrator_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Latest Orchestrator Time in COS.",
						},
						"mfa_enabled": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "indicates if Multi Factor Authentication is enabled or not.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPdrDrSummaryResponseRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_dr_summary_response", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrSummaryOptions := &drautomationservicev1.GetDrSummaryOptions{}

	getDrSummaryOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrSummaryOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	drAutomationGetSummaryResponse, response, err := drAutomationServiceClient.GetDrSummaryWithContext(context, getDrSummaryOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetDrSummaryWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetDrSummaryWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_dr_summary_response", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrDrSummaryResponseID(d))

	if !core.IsNil(drAutomationGetSummaryResponse.ServiceDetails) {
		serviceDetails := []map[string]interface{}{}
		serviceDetailsMap, err := DataSourceIBMPdrDrSummaryResponseServiceDetailsToMap(drAutomationGetSummaryResponse.ServiceDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_dr_summary_response", "read", "service_details-to-map").GetDiag()
		}
		serviceDetails = append(serviceDetails, serviceDetailsMap)
		if err = d.Set("service_details", serviceDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_details: %s", err), "(Data) ibm_pdr_dr_summary_response", "read", "set-service_details").GetDiag()
		}
	}

	if !core.IsNil(drAutomationGetSummaryResponse.OrchestratorDetails) {
		orchestratorDetails := []map[string]interface{}{}
		orchestratorDetailsMap, err := DataSourceIBMPdrDrSummaryResponseOrchestratorDetailsToMap(drAutomationGetSummaryResponse.OrchestratorDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_dr_summary_response", "read", "orchestrator_details-to-map").GetDiag()
		}
		orchestratorDetails = append(orchestratorDetails, orchestratorDetailsMap)
		if err = d.Set("orchestrator_details", orchestratorDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_details: %s", err), "(Data) ibm_pdr_dr_summary_response", "read", "set-orchestrator_details").GetDiag()
		}
	}
	fmt.Println("==============================================")
	fmt.Println("==============================================")
	fmt.Println("==============================================")
	fmt.Println("form terraform")
	fmt.Println(drAutomationGetSummaryResponse.ManagedVMList)
	fmt.Println("==============================================")
	fmt.Println("==============================================")
	fmt.Println("==============================================")
	managedVMList := make([]map[string]interface{}, 0, len(drAutomationGetSummaryResponse.ManagedVMList))

	for _, vmDetails := range drAutomationGetSummaryResponse.ManagedVMList {
		obj := map[string]interface{}{
			"vm_id": *vmDetails.VMID,
		}

		if vmDetails.Core != nil {
			obj["core"] = *vmDetails.Core
		}
		if vmDetails.DrAverageTime != nil {
			obj["dr_average_time"] = *vmDetails.DrAverageTime
		}
		if vmDetails.DrRegion != nil {
			obj["dr_region"] = *vmDetails.DrRegion
		}
		if vmDetails.Memory != nil {
			obj["memory"] = *vmDetails.Memory
		}
		if vmDetails.Region != nil {
			obj["region"] = *vmDetails.Region
		}
		if vmDetails.VMName != nil {
			obj["vm_name"] = *vmDetails.VMName
		}
		if vmDetails.WorkgroupName != nil {
			obj["workgroup_name"] = *vmDetails.WorkgroupName
		}
		if vmDetails.WorkspaceName != nil {
			obj["workspace_name"] = *vmDetails.WorkspaceName
		}

		managedVMList = append(managedVMList, obj)
	}
	fmt.Println("===========================================")
	fmt.Println("===========================================")
	fmt.Println("===========================================")
	fmt.Println(managedVMList)
	fmt.Println("===========================================")
	fmt.Println("===========================================")
	fmt.Println("===========================================")

	if err := d.Set("managed_vm_list", managedVMList); err != nil {
		return diag.Errorf("failed to set managed_vm_list: %s", err)
	}
	return nil
}

// func stringValue(v *string) string {
// 	if v == nil {
// 		return ""
// 	}
// 	return *v
// }

// dataSourceIBMPdrDrSummaryResponseID returns a reasonable ID for the list.
func dataSourceIBMPdrDrSummaryResponseID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrDrSummaryResponseServiceDetailsToMap(model *drautomationservicev1.ServiceDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.OrchestratorHa != nil {
		modelMap["orchestrator_ha"] = *model.OrchestratorHa
	}
	if model.DeploymentName != nil {
		modelMap["deployment_name"] = *model.DeploymentName
	}
	if model.RecoveryLocation != nil {
		modelMap["recovery_location"] = *model.RecoveryLocation
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = *model.ResourceGroup
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.PrimaryIPAddress != nil {
		modelMap["primary_ip_address"] = *model.PrimaryIPAddress
	}
	if model.StandbyIPAddress != nil {
		modelMap["standby_ip_address"] = *model.StandbyIPAddress
	}
	if model.StandbyDescription != nil {
		modelMap["standby_description"] = *model.StandbyDescription
	}
	if model.StandbyStatus != nil {
		modelMap["standby_status"] = *model.StandbyStatus
	}
	if model.PrimaryOrchestratorDashboardURL != nil {
		modelMap["primary_orchestrator_dashboard_url"] = *model.PrimaryOrchestratorDashboardURL
	}
	if model.StandbyOrchestratorDashboardURL != nil {
		modelMap["standby_orchestrator_dashboard_url"] = *model.StandbyOrchestratorDashboardURL
	}
	if model.PlanName != nil {
		modelMap["plan_name"] = *model.PlanName
	}
	return modelMap, nil
}

func DataSourceIBMPdrDrSummaryResponseOrchestratorDetailsToMap(model *drautomationservicev1.OrchestratorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.OrchestratorName != nil {
		modelMap["orchestrator_name"] = *model.OrchestratorName
	}
	if model.OrchestratorStatus != nil {
		modelMap["orchestrator_status"] = *model.OrchestratorStatus
	}
	if model.SSHKeyName != nil {
		modelMap["ssh_key_name"] = *model.SSHKeyName
	}
	if model.StandbySSHKeyName != nil {
		modelMap["standby_ssh_key_name"] = *model.StandbySSHKeyName
	}
	if model.SchematicWorkspaceStatus != nil {
		modelMap["schematic_workspace_status"] = *model.SchematicWorkspaceStatus
	}
	if model.SchematicWorkspaceName != nil {
		modelMap["schematic_workspace_name"] = *model.SchematicWorkspaceName
	}
	if model.StandbyOrchestratorName != nil {
		modelMap["standby_orchestrator_name"] = *model.StandbyOrchestratorName
	}
	if model.StandbyOrchestratorStatus != nil {
		modelMap["standby_orchestrator_status"] = *model.StandbyOrchestratorStatus
	}
	if model.OrchestratorConfigStatus != nil {
		modelMap["orchestrator_config_status"] = *model.OrchestratorConfigStatus
	}
	if model.OrchestratorGroupLeader != nil {
		modelMap["orchestrator_group_leader"] = *model.OrchestratorGroupLeader
	}
	if model.OrchestratorClusterMessage != nil {
		modelMap["orchestrator_cluster_message"] = *model.OrchestratorClusterMessage
	}
	if model.OrchStandbyNodeAdditionStatus != nil {
		modelMap["orch_standby_node_addition_status"] = *model.OrchStandbyNodeAdditionStatus
	}
	if model.OrchestratorLocationType != nil {
		modelMap["orchestrator_location_type"] = *model.OrchestratorLocationType
	}
	if model.LocationID != nil {
		modelMap["location_id"] = *model.LocationID
	}
	if model.VPCName != nil {
		modelMap["vpc_name"] = *model.VPCName
	}
	if model.TransitGatewayName != nil {
		modelMap["transit_gateway_name"] = *model.TransitGatewayName
	}
	if model.ProxyIP != nil {
		modelMap["proxy_ip"] = *model.ProxyIP
	}
	if model.OrchestratorWorkspaceName != nil {
		modelMap["orchestrator_workspace_name"] = *model.OrchestratorWorkspaceName
	}
	if model.StandbyOrchestratorWorkspaceName != nil {
		modelMap["standby_orchestrator_workspace_name"] = *model.StandbyOrchestratorWorkspaceName
	}
	if model.OrchExtConnectivityStatus != nil {
		modelMap["orch_ext_connectivity_status"] = *model.OrchExtConnectivityStatus
	}
	if model.LastUpdatedOrchestratorDeploymentTime != nil {
		modelMap["last_updated_orchestrator_deployment_time"] = model.LastUpdatedOrchestratorDeploymentTime.String()
	}
	if model.LastUpdatedStandbyOrchestratorDeploymentTime != nil {
		modelMap["last_updated_standby_orchestrator_deployment_time"] = model.LastUpdatedStandbyOrchestratorDeploymentTime.String()
	}
	if model.LatestOrchestratorTime != nil {
		modelMap["latest_orchestrator_time"] = model.LatestOrchestratorTime.String()
	}
	if model.MfaEnabled != nil {
		modelMap["mfa_enabled"] = *model.MfaEnabled
	}
	return modelMap, nil
}
