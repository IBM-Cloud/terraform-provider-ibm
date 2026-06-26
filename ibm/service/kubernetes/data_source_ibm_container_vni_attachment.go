// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/container/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMContainerVNIAttachment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerVNIAttachmentRead,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The cluster ID or name",
			},
			"worker": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The worker ID",
			},
			"vni_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VNI ID",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the resource group",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The VLAN ID for the bare metal worker",
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
				Description: "Whether the VNI will be deleted when the attachment is destroyed",
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
	}
}

func dataSourceIBMContainerVNIAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	workerID := d.Get("worker").(string)
	vniID := d.Get("vni_id").(string)
	clusterID := d.Get("cluster").(string)

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

	// List attachments for the worker to find this specific VNI
	input := graphql.ListVNIAttachmentsInput{
		NodeID: workerID,
	}

	resp, err := vniClient.ListAttachments(input, targetEnv)
	if err != nil {
		return fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Find the specific VNI attachment
	var attachment *graphql.VNIAttachment
	for _, edge := range resp.Connection.Edges {
		if edge.Node.VirtualNetworkInterface.ExternalID == vniID {
			attachment = &edge.Node
			break
		}
	}

	if attachment == nil {
		return fmt.Errorf("VNI attachment not found for worker %s and VNI %s", workerID, vniID)
	}

	// Set ID (vni_id is unique and avoids issues with worker ID format)
	d.SetId(vniID)

	// Set attributes
	d.Set("vni_id", attachment.VirtualNetworkInterface.ExternalID)
	d.Set("worker", attachment.AttachedTo.ID)
	d.Set("cluster", clusterID)

	if attachment.VlanID != nil && *attachment.VlanID > 0 {
		d.Set("vlan_id", *attachment.VlanID)
	}
	if attachment.Status != "" {
		d.Set("status", attachment.Status)
	}
	if attachment.CreatedAt != "" {
		d.Set("created_at", attachment.CreatedAt)
	}
	if attachment.VirtualNetworkInterface.AutoDelete != nil {
		d.Set("auto_delete", *attachment.VirtualNetworkInterface.AutoDelete)
	}
	if attachment.VirtualNetworkInterface.PrimaryIPAddress != nil {
		d.Set("primary_ip_address", *attachment.VirtualNetworkInterface.PrimaryIPAddress)
	}
	if attachment.VirtualNetworkInterface.MACAddress != nil {
		d.Set("mac_address", *attachment.VirtualNetworkInterface.MACAddress)
	}
	if attachment.VirtualNetworkInterface.Name != nil {
		d.Set("vni_name", *attachment.VirtualNetworkInterface.Name)
	}

	return nil
}
