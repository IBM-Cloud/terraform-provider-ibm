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
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaClusterNodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaClusterNodesRead,

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
			"primary_node_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details of the primary cluster nodes.",
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
							Description: "Number of CPU cores allocated to the VM.",
						},
						"ip_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of IP addresses assigned to the VM.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"memory": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Amount of memory allocated to the VM (in GB).",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the VM is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the VM.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VM.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the VM.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the workspace associated with the VM.",
						},
					},
				},
			},
			"secondary_node_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details of the secondary cluster nodes.",
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
							Description: "Number of CPU cores allocated to the VM.",
						},
						"ip_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of IP addresses assigned to the VM.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"memory": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Amount of memory allocated to the VM (in GB).",
						},
						"pha_level": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PowerHA version level installed on the node.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the VM is deployed.",
						},
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the VM.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VM.",
						},
						"vm_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the VM.",
						},
						"workspace_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the workspace associated with the VM.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPhaClusterNodesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_cluster_nodes", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getClusterNodeOptions := &powerhaautomationservicev1.GetClusterNodeOptions{}

	getClusterNodeOptions.SetPhaInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getClusterNodeOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	clusterNodeResponse, _, err := powerhaAutomationServiceClient.GetClusterNodeWithContext(context, getClusterNodeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetClusterNodeWithContext failed: %s", err.Error()), "(Data) ibm_pha_cluster_nodes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getClusterNodeOptions.PhaInstanceID)

	primaryNodeDetails := []map[string]interface{}{}
	for _, primaryNodeDetailsItem := range clusterNodeResponse.PrimaryNodeDetails {
		primaryNodeDetailsItemMap, err := DataSourceIBMPhaClusterNodesNodeDetailToMap(&primaryNodeDetailsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_cluster_nodes", "read", "primary_node_details-to-map").GetDiag()
		}
		primaryNodeDetails = append(primaryNodeDetails, primaryNodeDetailsItemMap)
	}
	if err = d.Set("primary_node_details", primaryNodeDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_node_details: %s", err), "(Data) ibm_pha_cluster_nodes", "read", "set-primary_node_details").GetDiag()
	}

	secondaryNodeDetails := []map[string]interface{}{}
	for _, secondaryNodeDetailsItem := range clusterNodeResponse.SecondaryNodeDetails {
		secondaryNodeDetailsItemMap, err := DataSourceIBMPhaClusterNodesNodeDetailToMap(&secondaryNodeDetailsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_cluster_nodes", "read", "secondary_node_details-to-map").GetDiag()
		}
		secondaryNodeDetails = append(secondaryNodeDetails, secondaryNodeDetailsItemMap)
	}
	if err = d.Set("secondary_node_details", secondaryNodeDetails); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting secondary_node_details: %s", err), "(Data) ibm_pha_cluster_nodes", "read", "set-secondary_node_details").GetDiag()
	}

	return nil
}

func DataSourceIBMPhaClusterNodesNodeDetailToMap(model *powerhaautomationservicev1.NodeDetail) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AgentStatus != nil {
		modelMap["agent_status"] = *model.AgentStatus
	}
	if model.Cores != nil {
		modelMap["cores"] = flex.Float64Value(model.Cores)
	}
	modelMap["ip_addresses"] = model.IPAddresses
	if model.Memory != nil {
		modelMap["memory"] = flex.Float64Value(model.Memory)
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
