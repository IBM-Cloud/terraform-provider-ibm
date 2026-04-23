// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaDeploymentRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the provisioned instance.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"cloud_account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud account identifier.",
			},
			"connectivity_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of network connectivity.",
			},
			"creation_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp expressing creation time.",
			},
			"custom_network": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of custom network CIDRs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deprovision_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp expressing deprovision time.",
			},
			"guid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Global unique identifier.",
			},
			"is_duplicate": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether deployment is duplicate.",
			},
			"plan_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier for the service plan.",
			},
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of service plan.",
			},
			"powerha_cluster_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the PowerHA cluster.",
			},
			"powerha_cluster_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of PowerHA cluster.",
			},
			"powerha_level": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PowerHA version level.",
			},
			"primary_cluster_nodes_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of primary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Number of CPU cores allocated to the node.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address assigned to the virtual machine.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory allocated to the virtual machine in MB or GB.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the virtual machine is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the virtual machine.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the virtual machine.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current operational status of the virtual machine.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workspace identifier associated with the node.",
						},
					},
				},
			},
			"primary_location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary cluster location.",
			},
			"primary_region_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the primary workspace region.",
			},
			"primary_workspace": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary workspace identifier.",
			},
			"primary_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the primary workspace.",
			},
			"provision_end_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp provisioning completed.",
			},
			"provision_start_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp provisioning started.",
			},
			"provision_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current provision status.",
			},
			"region_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Deployment region identifier.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the resource group.",
			},
			"resource_group_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of associated resource group.",
			},
			"resource_instance": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource instance identifier.",
			},
			"secondary_cluster_nodes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of secondary cluster nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the PHA agent running on the node.",
						},
						"cores": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Number of CPU cores allocated to the node.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address assigned to the virtual machine.",
						},
						"memory": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory allocated to the virtual machine in MB or GB.",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the virtual machine is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the virtual machine.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the virtual machine.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current operational status of the virtual machine.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workspace identifier associated with the node.",
						},
					},
				},
			},
			"secondary_location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secondary cluster location.",
			},
			"secondary_workspace": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secondary workspace identifier.",
			},
			"service_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of provisioned service.",
			},
			"service_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier for the service.",
			},
			"service_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of service.",
			},
			"standby_region_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the standby workspace region.",
			},
			"standby_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the standby workspace.",
			},
			"user_tags": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User defined tags.",
			},
		},
	}
}

func dataSourceIBMPhaDeploymentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_deployment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPhaDeploymentOptions := &powerhaautomationservicev1.GetPhaDeploymentOptions{}

	getPhaDeploymentOptions.SetPhaInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getPhaDeploymentOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaDeploymentResponse, response, err := powerhaAutomationServiceClient.GetPhaDeploymentWithContext(context, getPhaDeploymentOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetPhaDeploymentWithContext failed: %s", err.Error())

		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetPhaDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}

		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_deployment", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}
	d.SetId(*getPhaDeploymentOptions.PhaInstanceID)

	if !core.IsNil(phaDeploymentResponse.CloudAccountID) {
		if err = d.Set("cloud_account_id", phaDeploymentResponse.CloudAccountID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cloud_account_id: %s", err), "(Data) ibm_pha_deployment", "read", "set-cloud_account_id").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ConnectivityType) {
		if err = d.Set("connectivity_type", phaDeploymentResponse.ConnectivityType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connectivity_type: %s", err), "(Data) ibm_pha_deployment", "read", "set-connectivity_type").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.CreationTime) {
		if err = d.Set("creation_time", phaDeploymentResponse.CreationTime); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting creation_time: %s", err), "(Data) ibm_pha_deployment", "read", "set-creation_time").GetDiag()
		}
	}

	customNetwork := []interface{}{}
	for _, customNetworkItem := range phaDeploymentResponse.CustomNetwork {
		customNetwork = append(customNetwork, customNetworkItem)
	}
	if err = d.Set("custom_network", customNetwork); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting custom_network: %s", err), "(Data) ibm_pha_deployment", "read", "set-custom_network").GetDiag()
	}

	if !core.IsNil(phaDeploymentResponse.DeprovisionTime) {
		if err = d.Set("deprovision_time", phaDeploymentResponse.DeprovisionTime); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deprovision_time: %s", err), "(Data) ibm_pha_deployment", "read", "set-deprovision_time").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.GUID) {
		if err = d.Set("guid", phaDeploymentResponse.GUID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting guid: %s", err), "(Data) ibm_pha_deployment", "read", "set-guid").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.IsDuplicate) {
		if err = d.Set("is_duplicate", phaDeploymentResponse.IsDuplicate); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_duplicate: %s", err), "(Data) ibm_pha_deployment", "read", "set-is_duplicate").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PlanID) {
		if err = d.Set("plan_id", phaDeploymentResponse.PlanID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan_id: %s", err), "(Data) ibm_pha_deployment", "read", "set-plan_id").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PlanName) {
		if err = d.Set("plan_name", phaDeploymentResponse.PlanName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-plan_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PowerhaClusterName) {
		if err = d.Set("powerha_cluster_name", phaDeploymentResponse.PowerhaClusterName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting powerha_cluster_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-powerha_cluster_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PowerhaClusterType) {
		if err = d.Set("powerha_cluster_type", phaDeploymentResponse.PowerhaClusterType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting powerha_cluster_type: %s", err), "(Data) ibm_pha_deployment", "read", "set-powerha_cluster_type").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PowerhaLevel) {
		if err = d.Set("powerha_level", phaDeploymentResponse.PowerhaLevel); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting powerha_level: %s", err), "(Data) ibm_pha_deployment", "read", "set-powerha_level").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryClusterNodesDetails) {
		primaryClusterNodesDetails := []map[string]interface{}{}
		for _, primaryClusterNodesDetailsItem := range phaDeploymentResponse.PrimaryClusterNodesDetails {
			primaryClusterNodesDetailsItemMap, err := DataSourceIBMPhaDeploymentClusterNodeInfoToMap(&primaryClusterNodesDetailsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_deployment", "read", "primary_cluster_nodes_details-to-map").GetDiag()
			}
			primaryClusterNodesDetails = append(primaryClusterNodesDetails, primaryClusterNodesDetailsItemMap)
		}
		if err = d.Set("primary_cluster_nodes_details", primaryClusterNodesDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_cluster_nodes_details: %s", err), "(Data) ibm_pha_deployment", "read", "set-primary_cluster_nodes_details").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryLocation) {
		if err = d.Set("primary_location", phaDeploymentResponse.PrimaryLocation); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_location: %s", err), "(Data) ibm_pha_deployment", "read", "set-primary_location").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryRegionName) {
		if err = d.Set("primary_region_name", phaDeploymentResponse.PrimaryRegionName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_region_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-primary_region_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryWorkspace) {
		if err = d.Set("primary_workspace", phaDeploymentResponse.PrimaryWorkspace); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_workspace: %s", err), "(Data) ibm_pha_deployment", "read", "set-primary_workspace").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.PrimaryWorkspaceName) {
		if err = d.Set("primary_workspace_name", phaDeploymentResponse.PrimaryWorkspaceName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_workspace_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-primary_workspace_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ProvisionEndTime) {
		if err = d.Set("provision_end_time", phaDeploymentResponse.ProvisionEndTime); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provision_end_time: %s", err), "(Data) ibm_pha_deployment", "read", "set-provision_end_time").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ProvisionStartTime) {
		if err = d.Set("provision_start_time", phaDeploymentResponse.ProvisionStartTime); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provision_start_time: %s", err), "(Data) ibm_pha_deployment", "read", "set-provision_start_time").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ProvisionStatus) {
		if err = d.Set("provision_status", phaDeploymentResponse.ProvisionStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provision_status: %s", err), "(Data) ibm_pha_deployment", "read", "set-provision_status").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.RegionID) {
		if err = d.Set("region_id", phaDeploymentResponse.RegionID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting region_id: %s", err), "(Data) ibm_pha_deployment", "read", "set-region_id").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ResourceGroup) {
		if err = d.Set("resource_group", phaDeploymentResponse.ResourceGroup); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_pha_deployment", "read", "set-resource_group").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ResourceGroupCRN) {
		if err = d.Set("resource_group_crn", phaDeploymentResponse.ResourceGroupCRN); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_crn: %s", err), "(Data) ibm_pha_deployment", "read", "set-resource_group_crn").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ResourceInstance) {
		if err = d.Set("resource_instance", phaDeploymentResponse.ResourceInstance); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_instance: %s", err), "(Data) ibm_pha_deployment", "read", "set-resource_instance").GetDiag()
		}
	}

	secondaryClusterNodes := []map[string]interface{}{}
	for _, secondaryClusterNodesItem := range phaDeploymentResponse.SecondaryClusterNodes {
		secondaryClusterNodesItemMap, err := DataSourceIBMPhaDeploymentClusterNodeInfoToMap(&secondaryClusterNodesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_deployment", "read", "secondary_cluster_nodes-to-map").GetDiag()
		}
		secondaryClusterNodes = append(secondaryClusterNodes, secondaryClusterNodesItemMap)
	}
	if err = d.Set("secondary_cluster_nodes", secondaryClusterNodes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting secondary_cluster_nodes: %s", err), "(Data) ibm_pha_deployment", "read", "set-secondary_cluster_nodes").GetDiag()
	}

	if !core.IsNil(phaDeploymentResponse.SecondaryLocation) {
		if err = d.Set("secondary_location", phaDeploymentResponse.SecondaryLocation); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting secondary_location: %s", err), "(Data) ibm_pha_deployment", "read", "set-secondary_location").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.SecondaryWorkspace) {
		if err = d.Set("secondary_workspace", phaDeploymentResponse.SecondaryWorkspace); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting secondary_workspace: %s", err), "(Data) ibm_pha_deployment", "read", "set-secondary_workspace").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ServiceDescription) {
		if err = d.Set("service_description", phaDeploymentResponse.ServiceDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_description: %s", err), "(Data) ibm_pha_deployment", "read", "set-service_description").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ServiceID) {
		if err = d.Set("service_id", phaDeploymentResponse.ServiceID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_id: %s", err), "(Data) ibm_pha_deployment", "read", "set-service_id").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.ServiceName) {
		if err = d.Set("service_name", phaDeploymentResponse.ServiceName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-service_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.StandbyRegionName) {
		if err = d.Set("standby_region_name", phaDeploymentResponse.StandbyRegionName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_region_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-standby_region_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.StandbyWorkspaceName) {
		if err = d.Set("standby_workspace_name", phaDeploymentResponse.StandbyWorkspaceName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting standby_workspace_name: %s", err), "(Data) ibm_pha_deployment", "read", "set-standby_workspace_name").GetDiag()
		}
	}

	if !core.IsNil(phaDeploymentResponse.UserTags) {
		if err = d.Set("user_tags", phaDeploymentResponse.UserTags); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_tags: %s", err), "(Data) ibm_pha_deployment", "read", "set-user_tags").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMPhaDeploymentClusterNodeInfoToMap(model *powerhaautomationservicev1.ClusterNodeInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AgentStatus != nil {
		modelMap["agent_status"] = *model.AgentStatus
	}
	if model.Cores != nil {
		modelMap["cores"] = flex.Float64Value(model.Cores)
	}
	if model.IPAddress != nil {
		modelMap["ip_address"] = *model.IPAddress
	}
	if model.Memory != nil {
		modelMap["memory"] = flex.IntValue(model.Memory)
	}
	if model.PhaLevel != nil {
		modelMap["pha_level"] = *model.PhaLevel
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.VMID != nil {
		modelMap["vm_id"] = *model.VMID
	}
	if model.VMName != nil {
		modelMap["vm_name"] = *model.VMName
	}
	if model.VMStatus != nil {
		modelMap["vm_status"] = *model.VMStatus
	}
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = *model.WorkspaceID
	}
	return modelMap, nil
}
