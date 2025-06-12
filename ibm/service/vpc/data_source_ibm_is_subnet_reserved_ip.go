// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Define all the constants that matches with the given terrafrom attribute
const (
	// Request Param Constants
	isSubNetID     = "subnet"
	isReservedIPID = "reserved_ip"

	// Response Param Constants
	isReservedIPAddress    = "address"
	isReservedIPAutoDelete = "auto_delete"
	isReservedIPCreatedAt  = "created_at"
	isReservedIPhref       = "href"
	isReservedIPName       = "name"
	isReservedIPOwner      = "owner"
	isReservedIPType       = "resource_type"
)

func DataSourceIBMISReservedIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSdataSourceIBMISReservedIPRead,
		Schema: map[string]*schema.Schema{
			/*
				Request Parameters
				==================
				These are mandatory req parameters
			*/
			isSubNetID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet identifier.",
			},
			isReservedIPID: {
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

			isReservedIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address",
			},
			isReservedIPLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the reserved IP",
			},
			isReservedIPAutoDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, this reserved IP will be automatically deleted",
			},
			isReservedIPCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the reserved IP was created.",
			},
			isReservedIPhref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this reserved IP.",
			},
			isReservedIPName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user-defined or system-provided name for this reserved IP.",
			},
			isReservedIPOwner: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner of a reserved IP, defining whether it is managed by the user or the provider.",
			},
			isReservedIPType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			isReservedIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reserved IP target id.",
			},
			isReservedIPTargetCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn for target.",
			},
			"target_reference": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target this reserved IP is bound to.If absent, this reserved IP is provider-owned or unbound.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this endpoint gateway.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this endpoint gateway.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this endpoint gateway.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
		},
	}
}

// dataSdataSourceIBMISReservedIPRead is used when the reserved IPs are read from the vpc
func dataSdataSourceIBMISReservedIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_subnet_reserved_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	subnetID := d.Get(isSubNetID).(string)
	reservedIPID := d.Get(isReservedIPID).(string)

	options := sess.NewGetSubnetReservedIPOptions(subnetID, reservedIPID)
	reservedIP, response, err := sess.GetSubnetReservedIPWithContext(context, options)

	if err != nil || response == nil || reservedIP == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "(Data) ibm_is_subnet_reserved_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reservedIP.ID)
	if err = d.Set("auto_delete", reservedIP.AutoDelete); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-auto_delete").GetDiag()
	}
	if err = d.Set("address", reservedIP.Address); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-address").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(reservedIP.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", reservedIP.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-href").GetDiag()
	}
	if err = d.Set("name", reservedIP.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-name").GetDiag()
	}
	if err = d.Set("owner", reservedIP.Owner); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-owner").GetDiag()
	}
	if reservedIP.LifecycleState != nil {
		if err = d.Set("lifecycle_state", reservedIP.LifecycleState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-lifecycle_state").GetDiag()
		}
	}
	if err = d.Set("resource_type", reservedIP.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-resource_type").GetDiag()
	}
	target := []map[string]interface{}{}
	if reservedIP.Target != nil {
		modelMap, err := dataSourceIBMIsReservedIPReservedIPTargetToMap(reservedIP.Target)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_subnet_reserved_ip", "read", "target_reference-to-map").GetDiag()
		}
		target = append(target, modelMap)
	}
	if err = d.Set("target_reference", target); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target_reference: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-target_reference").GetDiag()
	}
	if len(target) > 0 {

		if err = d.Set(isReservedIPTarget, target[0]["id"]); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
		}

		if err = d.Set(isReservedIPTargetCrn, target[0]["crn"]); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target_crn: %s", err), "(Data) ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
		}
	}
	return nil // By default there should be no error
}
