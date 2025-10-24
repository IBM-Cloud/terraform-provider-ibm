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
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIbmPdrGetDrSummaryResponse() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetDrSummaryResponseRead,

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
			"managed_vm_list": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"orchestrator_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Contains details about the orchestrator configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location identifier.",
						},
						"orch_ext_connectivity_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External connectivity status of the orchestrator.",
						},
						"orch_standby_node_addition_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of standby node addition.",
						},
						"orchestrator_cluster_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message regarding orchestrator cluster status.",
						},
						"orchestrator_cluster_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of orchestrator cluster.",
						},
						"orchestrator_config_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configuration status of the orchestrator.",
						},
						"orchestrator_group_leader": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Leader node of the orchestrator group.",
						},
						"orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the primary orchestrator.",
						},
						"orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the primary orchestrator.",
						},
						"orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the orchestrator workspace.",
						},
						"proxy_ip": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address of the proxy.",
						},
						"schematic_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the schematic workspace.",
						},
						"schematic_workspace_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the schematic workspace.",
						},
						"ssh_key_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH key name used for the orchestrator.",
						},
						"standby_orchestrator_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the standby orchestrator.",
						},
						"standby_orchestrator_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the standby orchestrator.",
						},
						"standby_orchestrator_workspace_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the standby orchestrator workspace.",
						},
						"transit_gateway_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the transit gateway.",
						},
						"vpc_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VPC.",
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
							Description: "Cloud Resource Name identifier.",
						},
						"deployment_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the deployment.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the primary service.",
						},
						"is_ksys_ha": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag indicating if KSYS HA is enabled.",
						},
						"primary_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address of the primary service.",
						},
						"primary_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Primary Orchestrator Dashboard URL.",
						},
						"recovery_location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location for disaster recovery.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource group name.",
						},
						"standby_description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the standby service.",
						},
						"standby_ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address of the standby service.",
						},
						"standby_orchestrator_dashboard_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Standby Orchestrator Dashboard URL.",
						},
						"standby_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the standby service.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the primary service.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmPdrGetDrSummaryResponseRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if _, ok := d.GetOk("if_none_match"); ok {
		getDrSummaryOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	drAutomationGetSummaryResponse, _, err := drAutomationServiceClient.GetDrSummaryWithContext(context, getDrSummaryOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDrSummaryWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_dr_summary_response", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetDrSummaryResponseID(d))

	convertedMap := make(map[string]interface{}, len(drAutomationGetSummaryResponse.ManagedVMList))
	for k, v := range drAutomationGetSummaryResponse.ManagedVMList {
		convertedMap[k] = v
	}
	if err = d.Set("managed_vm_list", flex.Flatten(convertedMap)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vm_list: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-managed_vm_list").GetDiag()
	}

	orchestratorDetails := []map[string]interface{}{}
	orchestratorDetailsMap, err := DataSourceIbmPdrGetDrSummaryResponseOrchestratorDetailsToMap(drAutomationGetSummaryResponse.OrchestratorDetails)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_summary_response", "read", "orchestrator_details-to-map").GetDiag()
	}
	orchestratorDetails = append(orchestratorDetails, orchestratorDetailsMap)
	if err = d.Set("orchestrator_details", orchestratorDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting orchestrator_details: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-orchestrator_details").GetDiag()
	}

	serviceDetails := []map[string]interface{}{}
	serviceDetailsMap, err := DataSourceIbmPdrGetDrSummaryResponseServiceDetailsToMap(drAutomationGetSummaryResponse.ServiceDetails)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_summary_response", "read", "service_details-to-map").GetDiag()
	}
	serviceDetails = append(serviceDetails, serviceDetailsMap)
	if err = d.Set("service_details", serviceDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_details: %s", err), "(Data) ibm_pdr_get_dr_summary_response", "read", "set-service_details").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrGetDrSummaryResponseID returns a reasonable ID for the list.
func dataSourceIbmPdrGetDrSummaryResponseID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrGetDrSummaryResponseOrchestratorDetailsToMap(model *drautomationservicev1.OrchestratorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["location_id"] = *model.LocationID
	modelMap["orch_ext_connectivity_status"] = *model.OrchExtConnectivityStatus
	modelMap["orch_standby_node_addition_status"] = *model.OrchStandbyNodeAdditionStatus
	modelMap["orchestrator_cluster_message"] = *model.OrchestratorClusterMessage
	modelMap["orchestrator_cluster_type"] = *model.OrchestratorClusterType
	modelMap["orchestrator_config_status"] = *model.OrchestratorConfigStatus
	modelMap["orchestrator_group_leader"] = *model.OrchestratorGroupLeader
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

func DataSourceIbmPdrGetDrSummaryResponseServiceDetailsToMap(model *drautomationservicev1.ServiceDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	modelMap["deployment_name"] = *model.DeploymentName
	modelMap["description"] = *model.Description
	modelMap["is_ksys_ha"] = *model.IsKsysHa
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
