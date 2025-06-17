// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerDisk = "disk"
)

func DataSourceIBMIsBareMetalServerDisk() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerDiskRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},

			isBareMetalServerDisk: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server disk identifier",
			},
			//disks

			isBareMetalServerDiskHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this bare metal server disk",
			},
			isBareMetalServerDiskID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this bare metal server disk",
			},
			isBareMetalServerDiskInterfaceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The disk interface used for attaching the disk. Supported values are [ nvme, sata ]",
			},
			isBareMetalServerDiskName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user-defined name for this disk",
			},
			isBareMetalServerDiskResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type",
			},
			isBareMetalServerDiskSize: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of the disk in GB (gigabytes)",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerDiskRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	bareMetalServerDiskID := d.Get(isBareMetalServerDisk).(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_disk", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerDiskOptions{
		BareMetalServerID: &bareMetalServerID,
		ID:                &bareMetalServerDiskID,
	}

	bareMetalServerDisk, _, err := sess.GetBareMetalServerDiskWithContext(context, options)
	if err != nil || bareMetalServerDisk == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerDiskWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_disk", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*bareMetalServerDisk.ID)
	if err = d.Set("href", bareMetalServerDisk.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_disk", "read", "set-href").GetDiag()
	}
	if err = d.Set("interface_type", bareMetalServerDisk.InterfaceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting interface_type: %s", err), "(Data) ibm_is_bare_metal_server_disk", "read", "set-interface_type").GetDiag()
	}
	if err = d.Set("name", bareMetalServerDisk.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_disk", "read", "set-name").GetDiag()
	}
	if err = d.Set("resource_type", bareMetalServerDisk.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_disk", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("size", flex.IntValue(bareMetalServerDisk.Size)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting size: %s", err), "(Data) ibm_is_bare_metal_server_disk", "read", "set-size").GetDiag()
	}

	return nil
}
