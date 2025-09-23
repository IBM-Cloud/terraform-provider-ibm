// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerNetworkInterfaceFloatingIPID = "floating_ip"
)

func DataSourceIBMIsBareMetalServerNetworkInterfaceFloatingIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerNetworkInterfaceFloatingIPRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},
			isBareMetalServerNetworkInterface: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier of bare metal server",
			},
			isBareMetalServerNetworkInterfaceFloatingIPID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The floating ip identifier of the network interface associated with the bare metal server",
			},
			floatingIPName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the floating IP",
			},

			floatingIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP address",
			},

			floatingIPStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP status",
			},

			floatingIPZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			floatingIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target info",
			},

			floatingIPCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP crn",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerNetworkInterfaceFloatingIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	nicID := d.Get(isBareMetalServerNetworkInterface).(string)
	fipID := d.Get(isBareMetalServerNetworkInterfaceFloatingIPID).(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerNetworkInterfaceFloatingIPOptions{
		BareMetalServerID:  &bareMetalServerID,
		NetworkInterfaceID: &nicID,
		ID:                 &fipID,
	}

	ip, _, err := sess.GetBareMetalServerNetworkInterfaceFloatingIPWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceFloatingIPWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(floatingIPName, *ip.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-name").GetDiag()
	}

	if err = d.Set(floatingIPAddress, *ip.Address); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-address").GetDiag()
	}

	if err = d.Set(floatingIPStatus, *ip.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-status").GetDiag()
	}
	if err = d.Set(floatingIPZone, *ip.Zone.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-zone").GetDiag()
	}

	if err = d.Set(floatingIPCRN, *ip.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-crn").GetDiag()
	}
	target, ok := ip.Target.(*vpcv1.FloatingIPTarget)
	if ok {

		if err = d.Set(floatingIPTarget, target.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_floating_ip", "read", "set-target").GetDiag()
		}
	}

	d.SetId(*ip.ID)

	return nil
}
