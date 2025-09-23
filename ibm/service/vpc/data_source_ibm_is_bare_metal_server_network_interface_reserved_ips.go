// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Define all the constants that matches with the given terrafrom attribute
const (
	// Request Param Constants
	isBareMetalServerNICReservedIPLimit  = "limit"
	isBareMetalServerNICReservedIPSort   = "sort"
	isBareMetalServerNICReservedIPs      = "reserved_ips"
	isBareMetalServerNICReservedIPsCount = "total_count"
)

func DataSourceIBMISBareMetalServerNICReservedIPs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerNICReservedIPsRead,
		Schema: map[string]*schema.Schema{
			/*
				Request Parameters
				==================
				These are mandatory req parameters
			*/
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The BareMetalServer identifier.",
			},
			isBareMetalServerNicID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The BareMetalServer network interface identifier.",
			},
			/*
				Response Parameters
				===================
				All of these are computed and an user doesn't need to provide
				these from outside.
			*/

			isBareMetalServerNICReservedIPs: {
				Type:        schema.TypeList,
				Description: "Collection of all reserved IPs bound to a network interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address",
						},
						isBareMetalServerNicIpAutoDelete: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If reserved ip shall be deleted automatically",
						},
						isBareMetalServerNICReservedIPCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the reserved IP was created.",
						},
						isBareMetalServerNICReservedIPhref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP.",
						},
						isBareMetalServerNicIpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this reserved IP",
						},
						isBareMetalServerNicIpName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined or system-provided name for this reserved IP.",
						},
						isBareMetalServerNICReservedIPOwner: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner of a reserved IP, defining whether it is managed by the user or the provider.",
						},
						isBareMetalServerNICReservedIPType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						isBareMetalServerNICReservedIPTarget: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reserved IP target id",
						},
					},
				},
			},
			isBareMetalServerNICReservedIPsCount: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources across all pages",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerNICReservedIPsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	nicID := d.Get(isBareMetalServerNicID).(string)

	// Flatten all the reserved IPs
	allrecs := []vpcv1.ReservedIP{}
	options := &vpcv1.ListBareMetalServerNetworkInterfaceIpsOptions{
		BareMetalServerID:  &bareMetalServerID,
		NetworkInterfaceID: &nicID,
	}

	result, response, err := sess.ListBareMetalServerNetworkInterfaceIpsWithContext(context, options)
	if err != nil || response == nil || result == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBareMetalServerNetworkInterfaceIpsWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allrecs = append(allrecs, result.Ips...)

	// Now store all the reserved IP info with their response tags
	reservedIPs := []map[string]interface{}{}
	for _, data := range allrecs {
		ipsOutput := map[string]interface{}{}
		ipsOutput[isBareMetalServerNicIpAddress] = *data.Address
		ipsOutput[isBareMetalServerNicIpAutoDelete] = *data.AutoDelete
		ipsOutput[isBareMetalServerNICReservedIPCreatedAt] = (*data.CreatedAt).String()
		ipsOutput[isBareMetalServerNICReservedIPhref] = *data.Href
		ipsOutput[isBareMetalServerNicIpID] = *data.ID
		ipsOutput[isBareMetalServerNicIpName] = *data.Name
		ipsOutput[isBareMetalServerNICReservedIPOwner] = *data.Owner
		ipsOutput[isBareMetalServerNICReservedIPType] = *data.ResourceType
		target, ok := data.Target.(*vpcv1.ReservedIPTarget)
		if ok {
			ipsOutput[isReservedIPTarget] = target.ID
		}
		reservedIPs = append(reservedIPs, ipsOutput)
	}

	d.SetId(time.Now().UTC().String()) // This is not any reserved ip or BareMetalServer id but state id
	if err = d.Set(isBareMetalServerNICReservedIPs, reservedIPs); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reserved_ips: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read", "set-reserved_ips").GetDiag()
	}
	if err = d.Set(isBareMetalServerNICReservedIPsCount, len(reservedIPs)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_count: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read", "set-total_count").GetDiag()
	}
	if err = d.Set(isBareMetalServerID, bareMetalServerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bare_metal_server: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read", "set-bare_metal_server").GetDiag()
	}
	if err = d.Set(isBareMetalServerNicID, nicID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interface: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ips", "read", "set-auto_delete").GetDiag()
	}
	return nil
}
