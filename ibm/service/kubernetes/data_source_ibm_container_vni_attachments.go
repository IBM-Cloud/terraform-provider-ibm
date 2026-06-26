// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMContainerVNIAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerVNIAttachmentsRead,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"cluster", "worker"},
				Description:  "The cluster ID or name to list all attachments",
			},
			"worker": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"cluster", "worker"},
				Description:  "The worker ID to list attachments for specific worker",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group",
			},
			"attachments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of VNI attachments",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cluster ID",
						},
						"worker_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The worker ID",
						},
						"vni_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VNI ID",
						},
						"vlan_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The VLAN ID",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the attachment",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the attachment was created",
						},
						"auto_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the VNI will be deleted when detached",
						},
						"primary_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The primary IP address of the VNI",
						},
						"mac_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MAC address of the VNI",
						},
						"vni_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the VNI",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMContainerVNIAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	cluster, hasCluster := d.GetOk("cluster")
	worker, _ := d.GetOk("worker")

	// Get VNI client
	vniClient, err := getVNIClient(meta)
	if err != nil {
		return fmt.Errorf("error creating VNI client: %s", err)
	}

	// Get target header
	targetEnv, err := getVNITargetHeader(d, meta)
	if err != nil {
		return fmt.Errorf("error creating target header: %s", err)
	}

	// Determine node ID for query
	var nodeID string
	if hasCluster {
		nodeID = cluster.(string)
	} else {
		nodeID = worker.(string)
	}

	// List all attachments with automatic pagination
	allAttachments, err := listAllVNIAttachments(vniClient, nodeID, targetEnv)
	if err != nil {
		return fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Set ID
	d.SetId(nodeID)

	// Convert to list format
	attachmentsList := make([]map[string]interface{}, 0, len(allAttachments))
	for _, attachment := range allAttachments {
		attachmentMap := map[string]interface{}{
			"worker_id": attachment.AttachedTo.ID,
			"vni_id":    attachment.VirtualNetworkInterface.ExternalID,
		}

		// Set cluster ID if available
		if hasCluster {
			attachmentMap["cluster_id"] = cluster.(string)
		}

		if attachment.VlanID != nil && *attachment.VlanID > 0 {
			attachmentMap["vlan_id"] = *attachment.VlanID
		}
		if attachment.Status != "" {
			attachmentMap["status"] = attachment.Status
		}
		if attachment.CreatedAt != "" {
			attachmentMap["created_at"] = attachment.CreatedAt
		}
		if attachment.VirtualNetworkInterface.AutoDelete != nil {
			attachmentMap["auto_delete"] = *attachment.VirtualNetworkInterface.AutoDelete
		}
		if attachment.VirtualNetworkInterface.PrimaryIPAddress != nil {
			attachmentMap["primary_ip_address"] = *attachment.VirtualNetworkInterface.PrimaryIPAddress
		}
		if attachment.VirtualNetworkInterface.MACAddress != nil {
			attachmentMap["mac_address"] = *attachment.VirtualNetworkInterface.MACAddress
		}
		if attachment.VirtualNetworkInterface.Name != nil {
			attachmentMap["vni_name"] = *attachment.VirtualNetworkInterface.Name
		}

		attachmentsList = append(attachmentsList, attachmentMap)
	}

	d.Set("attachments", attachmentsList)

	return nil
}

// listAllVNIAttachments fetches all VNI attachments with automatic pagination
func listAllVNIAttachments(client graphql.VNIs, nodeID string, targetEnv containerv1.ClusterTargetHeader) ([]graphql.VNIAttachment, error) {
	var allAttachments []graphql.VNIAttachment
	var cursor *string

	for {
		input := graphql.ListVNIAttachmentsInput{
			NodeID: nodeID,
			After:  cursor,
		}

		resp, err := client.ListAttachments(input, targetEnv)
		if err != nil {
			return nil, err
		}

		// Collect attachments from this page
		for _, edge := range resp.Connection.Edges {
			allAttachments = append(allAttachments, edge.Node)
		}

		// Check if there are more pages
		if !resp.Connection.PageInfo.HasNextPage {
			break
		}

		cursor = resp.Connection.PageInfo.EndCursor
	}

	return allAttachments, nil
}
