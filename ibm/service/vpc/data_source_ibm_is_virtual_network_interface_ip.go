// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVirtualNetworkInterfaceIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIsVirtualNetworkInterfaceIPRead,

		Schema: map[string]*schema.Schema{
			"virtual_network_interface": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual network interface identifier",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reserved IP identifier.",
			},
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this reserved IP.",
			},
			"reserved_ip": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for this reserved IP.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}

func dataSourceIsVirtualNetworkInterfaceIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVirtualNetworkInterfaceIPOptions := &vpcv1.GetVirtualNetworkInterfaceIPOptions{}

	getVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(d.Get("virtual_network_interface").(string))
	getVirtualNetworkInterfaceIPOptions.SetID(d.Get("reserved_ip").(string))

	reservedIP, _, err := vpcClient.GetVirtualNetworkInterfaceIPWithContext(context, getVirtualNetworkInterfaceIPOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVirtualNetworkInterfaceIPWithContext failed: %s", err.Error()), "(Data) ibm_is_virtual_network_interface_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reservedIP.ID)

	if err = d.Set("address", reservedIP.Address); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_is_virtual_network_interface_ip", "read", "set-address").GetDiag()
	}
	if err = d.Set("href", reservedIP.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_virtual_network_interface_ip", "read", "set-href").GetDiag()
	}

	if err = d.Set("reserved_ip", reservedIP.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reserved_ip: %s", err), "(Data) ibm_is_virtual_network_interface_ip", "read", "set-reserved_ip").GetDiag()
	}

	if err = d.Set("name", reservedIP.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_virtual_network_interface_ip", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", reservedIP.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_virtual_network_interface_ip", "read", "set-resource_type").GetDiag()
	}

	return nil
}
