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

// Define all the constants that matches with the given terrafrom attribute
const (
	// Request Param Constants
	isInstanceNICID           = "network_interface"
	isInstanceNICReservedIPID = "reserved_ip"

	// Response Param Constants
	isInstanceNICReservedIPAddress    = "address"
	isInstanceNICReservedIPAutoDelete = "auto_delete"
	isInstanceNICReservedIPCreatedAt  = "created_at"
	isInstanceNICReservedIPhref       = "href"
	isInstanceNICReservedIPName       = "name"
	isInstanceNICReservedIPOwner      = "owner"
	isInstanceNICReservedIPType       = "resource_type"
	isInstanceNICReservedIPTarget     = "target"
)

func DataSourceIBMISInstanceNICReservedIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceNICReservedIPRead,
		Schema: map[string]*schema.Schema{
			/*
				Request Parameters
				==================
				These are mandatory req parameters
			*/
			isInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance identifier.",
			},
			isInstanceNICID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance network interface identifier.",
			},
			isInstanceNICReservedIPID: {
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

			isInstanceNICReservedIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address",
			},
			isInstanceNICReservedIPAutoDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, this reserved IP will be automatically deleted",
			},
			isInstanceNICReservedIPCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the reserved IP was created.",
			},
			isInstanceNICReservedIPhref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this reserved IP.",
			},
			isInstanceNICReservedIPName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user-defined or system-provided name for this reserved IP.",
			},
			isInstanceNICReservedIPOwner: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner of a reserved IP, defining whether it is managed by the user or the provider.",
			},
			isInstanceNICReservedIPType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			isInstanceNICReservedIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reserved IP target id.",
			},
		},
	}
}

// dataSourceIBMISInstanceNICReservedIPRead is used when the reserved IPs are read from the vpc
func dataSourceIBMISInstanceNICReservedIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instance := d.Get(isInstanceID).(string)
	instanceNICId := d.Get(isInstanceNICID).(string)
	reservedIPID := d.Get(isInstanceNICReservedIPID).(string)

	options := sess.NewGetInstanceNetworkInterfaceIPOptions(instance, instanceNICId, reservedIPID)
	reserveIP, response, err := sess.GetInstanceNetworkInterfaceIPWithContext(context, options)

	if err != nil || response == nil || reserveIP == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceProfileWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_network_interface_reserved_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reserveIP.ID)

	if err = d.Set(isInstanceNICReservedIPAutoDelete, *reserveIP.AutoDelete); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-auto_delete").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPCreatedAt, (*reserveIP.CreatedAt).String()); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-created_at").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPhref, *reserveIP.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-href").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPName, *reserveIP.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-name").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPOwner, *reserveIP.Owner); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-owner").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPType, *reserveIP.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set(isInstanceNICReservedIPAddress, *reserveIP.Address); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-address").GetDiag()
	}
	if reserveIP.Target != nil {
		target, ok := reserveIP.Target.(*vpcv1.ReservedIPTarget)
		if ok {
			if err = d.Set(isInstanceNICReservedIPTarget, target.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_instance_network_interface_reserved_ip", "read", "set-target").GetDiag()
			}
		}
	}
	return nil // By default there should be no error
}
