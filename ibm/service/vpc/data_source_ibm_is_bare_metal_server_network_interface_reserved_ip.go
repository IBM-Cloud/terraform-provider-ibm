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

// Define all the constants that matches with the given terrafrom attribute
const (
	// Response Param Constants

	isBareMetalServerNICReservedIPCreatedAt      = "created_at"
	isBareMetalServerNICReservedIPhref           = "href"
	isBareMetalServerNICReservedIPLifecycleState = "lifecycle_state"
	isBareMetalServerNICReservedIPOwner          = "owner"
	isBareMetalServerNICReservedIPType           = "resource_type"
	isBareMetalServerNICReservedIPTarget         = "target"
)

func DataSourceIBMISBareMetalServerNICReservedIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerNICReservedIPRead,
		Schema: map[string]*schema.Schema{
			/*
				Request Parameters
				==================
				These are mandatory req parameters
			*/
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Bare Metal Server identifier.",
			},
			isBareMetalServerNicID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Bare Metal Server network interface identifier.",
			},
			isBareMetalServerNicIpID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The reserved IP identifier.",
			},

			/*
				Response Parameters
				===================
				All of these are computed and an user doesn't need to provide
				these from outside.
			*/

			isBareMetalServerNicIpAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address",
			},
			isBareMetalServerNicIpAutoDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, this reserved IP will be automatically deleted",
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
				Description: "Reserved IP target id.",
			},
		},
	}
}

// dataSourceIBMISBareMetalServerNICReservedIPRead is used when the reserved IPs are read from the vpc
func dataSourceIBMISBareMetalServerNICReservedIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bareMetalServer := d.Get(isBareMetalServerID).(string)
	bareMetalServerNICId := d.Get(isBareMetalServerNicID).(string)
	reservedIPID := d.Get(isBareMetalServerNicIpID).(string)

	options := sess.NewGetBareMetalServerNetworkInterfaceIPOptions(bareMetalServer, bareMetalServerNICId, reservedIPID)
	reserveIP, response, err := sess.GetBareMetalServerNetworkInterfaceIPWithContext(context, options)

	if err != nil || response == nil || reserveIP == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceIPWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reserveIP.ID)
	if err = d.Set(isBareMetalServerNicIpAutoDelete, *reserveIP.AutoDelete); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-auto_delete").GetDiag()
	}
	if err = d.Set(isBareMetalServerNICReservedIPCreatedAt, (*reserveIP.CreatedAt).String()); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-created_at").GetDiag()
	}
	if err = d.Set(isBareMetalServerNICReservedIPhref, *reserveIP.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-href").GetDiag()
	}
	if err = d.Set(isBareMetalServerNicIpName, *reserveIP.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-name").GetDiag()
	}
	if err = d.Set(isBareMetalServerNICReservedIPOwner, *reserveIP.Owner); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-owner").GetDiag()
	}
	if err = d.Set(isBareMetalServerNICReservedIPType, *reserveIP.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set(isBareMetalServerNicIpAddress, *reserveIP.Address); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-address").GetDiag()
	}
	if reserveIP.Target != nil {
		target, ok := reserveIP.Target.(*vpcv1.ReservedIPTarget)
		if ok {
			if err = d.Set(isBareMetalServerNICReservedIPTarget, target.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_bare_metal_server_network_interface_reserved_ip", "read", "set-target").GetDiag()
			}
		}
	}
	return nil // By default there should be no error
}
