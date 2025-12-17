// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
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
)

func DataSourceIBMPdrGetDrSummaryResponse() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrGetDrSummaryResponseRead,

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
			"managed_vm_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id":           {Type: schema.TypeString, Computed: true},
						"core":            {Type: schema.TypeString, Computed: true},
						"dr_average_time": {Type: schema.TypeString, Computed: true},
						"dr_region":       {Type: schema.TypeString, Computed: true},
						"memory":          {Type: schema.TypeString, Computed: true},
						"region":          {Type: schema.TypeString, Computed: true},
						"vm_name":         {Type: schema.TypeString, Computed: true},
						"workgroup_name":  {Type: schema.TypeString, Computed: true},
						"workspace_name":  {Type: schema.TypeString, Computed: true},
					},
				},
			},
			"orchestrator_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the orchestrator configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"location_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier of location.",
						},
						"mfa_enabled": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "indicates if Multi Factor Authentication is enabled or not.",
						},
						"orch_ext_connectivity_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The external connectivity status of the orchestrator.",
						},
						"orch_standby_node_addition_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of standby node addition.",
						},
						"orchestrator_cluster_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message regarding orchestrator cluster status.",
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
						"orchestrator_location_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of orchestrator Location.",
						},
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
						"orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the orchestrator workspace.",
						},
						"proxy_ip": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address of the proxy.",
						},
						"schematic_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the schematic workspace.",
						},
						"schematic_workspace_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the schematic workspace.",
						},
						"ssh_key_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH key name used for the orchestrator.",
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
						"standby_orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the standby orchestrator workspace.",
						},
						"transit_gateway_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit gateway.",
						},
						"vpc_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the VPC.",
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
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deployment crn.",
						},
						"deployment_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the deployment.",
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
						"plan_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The plan name.",
						},
						"primary_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service Orchestator primary IP address.",
						},
						"primary_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Primary Orchestrator Dashboard URL.",
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
						"standby_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The standby orchestrator current status details.",
						},
						"standby_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service Orchestator standby IP address.",
						},
						"standby_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Standby Orchestrator Dashboard URL.",
						},
						"standby_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The standby orchestrator current status.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status of the service.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPdrGetDrSummaryResponseRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_summary_response", "read", "initialize-client")
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
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_get_dr_summary_response", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrGetDrSummaryResponseID(d))

	vmList := make([]map[string]interface{}, 0)

	for vmID, rawVM := range drAutomationGetSummaryResponse.ManagedVMList {

		vmMap, ok := rawVM.(map[string]interface{})
		if !ok {
			return flex.DiscriminatedTerraformErrorf(
				fmt.Errorf("managed_vm_list[%s] is not an object", vmID),
				"(Data) ibm_pdr_get_dr_summary_response",
				"read",
				"invalid-managed_vm_list",
				"",
			).GetDiag()
		}

		item := map[string]interface{}{
			"vm_id": vmID,
		}

		if v, ok := vmMap["core"]; ok {
			item["core"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["dr_average_time"]; ok {
			item["dr_average_time"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["dr_region"]; ok {
			item["dr_region"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["memory"]; ok {
			item["memory"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["region"]; ok {
			item["region"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["vm_name"]; ok {
			item["vm_name"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["workgroup_name"]; ok {
			item["workgroup_name"] = fmt.Sprintf("%v", v)
		}
		if v, ok := vmMap["workspace_name"]; ok {
			item["workspace_name"] = fmt.Sprintf("%v", v)
		}

		vmList = append(vmList, item)
	}

	if err := d.Set("managed_vm_list", vmList); err != nil {
		return flex.DiscriminatedTerraformErrorf(
			err,
			fmt.Sprintf("Error setting managed_vm_list: %s", err),
			"(Data) ibm_pdr_get_dr_summary_response",
			"read",
			"set-managed_vm_list",
		).GetDiag()
	}

	// if err = d.Set("managed_vm_lists", converted); err != nil {
	// 	return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vm_list: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-managed_vm_list").GetDiag()
	// }

	orchestratorDetails := []map[string]interface{}{}
	orchestratorDetailsMap, err := DataSourceIBMPdrGetDrSummaryResponseOrchestratorDetailsToMap(drAutomationGetSummaryResponse.OrchestratorDetails)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_summary_response", "read", "orchestrator_details-to-map").GetDiag()
	}
	orchestratorDetails = append(orchestratorDetails, orchestratorDetailsMap)
	if err = d.Set("orchestrator_details", orchestratorDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_details: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-orchestrator_details").GetDiag()
	}

	serviceDetails := []map[string]interface{}{}
	serviceDetailsMap, err := DataSourceIBMPdrGetDrSummaryResponseServiceDetailsToMap(drAutomationGetSummaryResponse.ServiceDetails)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_summary_response", "read", "service_details-to-map").GetDiag()
	}
	serviceDetails = append(serviceDetails, serviceDetailsMap)
	if err = d.Set("service_details", serviceDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_details: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-service_details").GetDiag()
	}

	return nil
}

// dataSourceIBMPdrGetDrSummaryResponseID returns a reasonable ID for the list.
func dataSourceIBMPdrGetDrSummaryResponseID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrGetDrSummaryResponseOrchestratorDetailsToMap(model *drautomationservicev1.OrchestratorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["last_updated_orchestrator_deployment_time"] = model.LastUpdatedOrchestratorDeploymentTime.String()
	modelMap["last_updated_standby_orchestrator_deployment_time"] = model.LastUpdatedStandbyOrchestratorDeploymentTime.String()
	if model.LatestOrchestratorTime != nil {
		modelMap["latest_orchestrator_time"] = model.LatestOrchestratorTime.String()
	}
	modelMap["location_id"] = *model.LocationID
	modelMap["mfa_enabled"] = *model.MfaEnabled
	modelMap["orch_ext_connectivity_status"] = *model.OrchExtConnectivityStatus
	modelMap["orch_standby_node_addition_status"] = *model.OrchStandbyNodeAdditionStatus
	modelMap["orchestrator_cluster_message"] = *model.OrchestratorClusterMessage
	modelMap["orchestrator_config_status"] = *model.OrchestratorConfigStatus
	modelMap["orchestrator_group_leader"] = *model.OrchestratorGroupLeader
	modelMap["orchestrator_location_type"] = *model.OrchestratorLocationType
	modelMap["orchestrator_name"] = *model.OrchestratorName
	modelMap["orchestrator_status"] = *model.OrchestratorStatus
	modelMap["orchestrator_workspace_name"] = *model.OrchestratorWorkspaceName
	modelMap["proxy_ip"] = *model.ProxyIP
	modelMap["schematic_workspace_name"] = *model.SchematicWorkspaceName
	modelMap["schematic_workspace_status"] = *model.SchematicWorkspaceStatus
	modelMap["ssh_key_name"] = *model.SSHKeyName
	modelMap["standby_orchestrator_name"] = *model.StandbyOrchestratorName
	modelMap["standby_orchestrator_status"] = *model.StandbyOrchestratorStatus
	modelMap["standby_orchestrator_workspace_name"] = *model.StandbyOrchestratorWorkspaceName
	modelMap["transit_gateway_name"] = *model.TransitGatewayName
	modelMap["vpc_name"] = *model.VPCName
	return modelMap, nil
}

func DataSourceIBMPdrGetDrSummaryResponseServiceDetailsToMap(model *drautomationservicev1.ServiceDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	modelMap["deployment_name"] = *model.DeploymentName
	modelMap["description"] = *model.Description
	modelMap["orchestrator_ha"] = *model.OrchestratorHa
	modelMap["plan_name"] = *model.PlanName
	modelMap["primary_ip_address"] = *model.PrimaryIPAddress
	modelMap["primary_orchestrator_dashboard_url"] = *model.PrimaryOrchestratorDashboardURL
	modelMap["recovery_location"] = *model.RecoveryLocation
	modelMap["resource_group"] = *model.ResourceGroup
	modelMap["standby_description"] = *model.StandbyDescription
	modelMap["standby_ip_address"] = *model.StandbyIPAddress
	modelMap["standby_orchestrator_dashboard_url"] = *model.StandbyOrchestratorDashboardURL
	modelMap["standby_status"] = *model.StandbyStatus
	modelMap["status"] = *model.Status
	return modelMap, nil
}
