// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISPublicGateway() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISPublicGatewayRead,

		Schema: map[string]*schema.Schema{
			isPublicGatewayName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Public gateway Name",
			},

			isPublicGatewayFloatingIP: {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Public gateway floating IP",
			},

			isPublicGatewayStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway instance status",
			},

			isPublicGatewayResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Public gateway resource group info",
			},

			isPublicGatewayVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway VPC info",
			},

			isPublicGatewayZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public gateway zone info",
			},

			isPublicGatewayTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Service tags for the public gateway instance",
			},

			isPublicGatewayAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			isPublicGatewayCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func dataSourceIBMISPublicGatewayRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_public_gateway", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	name := d.Get(isPublicGatewayName).(string)
	rgroup := ""
	if rg, ok := d.GetOk(isPublicGatewayResourceGroup); ok {
		rgroup = rg.(string)
	}
	start := ""
	allrecs := []vpcv1.PublicGateway{}
	for {
		listPublicGatewaysOptions := &vpcv1.ListPublicGatewaysOptions{}
		if start != "" {
			listPublicGatewaysOptions.Start = &start
		}
		if rgroup != "" {
			listPublicGatewaysOptions.ResourceGroupID = &rgroup
		}
		publicgws, _, err := sess.ListPublicGatewaysWithContext(context, listPublicGatewaysOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListPublicGatewaysWithContext failed: %s", err.Error()), "(Data) ibm_is_public_gateway", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(publicgws.Next)
		allrecs = append(allrecs, publicgws.PublicGateways...)
		if start == "" {
			break
		}
	}
	for _, publicgw := range allrecs {
		if *publicgw.Name == name {
			d.SetId(*publicgw.ID)
			if err = d.Set("name", publicgw.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_public_gateway", "read", "set-name").GetDiag()
			}
			if publicgw.FloatingIP != nil {
				floatIP := map[string]interface{}{
					"id":                             *publicgw.FloatingIP.ID,
					isPublicGatewayFloatingIPAddress: *publicgw.FloatingIP.Address,
				}
				if err = d.Set("floating_ip", floatIP); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting floating_ip: %s", err), "(Data) ibm_is_public_gateway", "read", "set-floating_ip").GetDiag()
				}
			}
			if err = d.Set("status", publicgw.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_public_gateway", "read", "set-status").GetDiag()
			}
			if err = d.Set("zone", publicgw.Zone.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_is_public_gateway", "read", "set-zone").GetDiag()
			}
			if err = d.Set("vpc", *publicgw.VPC.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_public_gateway", "read", "set-vpc").GetDiag()
			}
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *publicgw.CRN, "", isUserTagType)
			if err != nil {
				log.Printf(
					"Error on get of vpc public gateway (%s) tags: %s", *publicgw.ID, err)
			}
			if err = d.Set("tags", tags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_public_gateway", "read", "set-tags").GetDiag()
			}
			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *publicgw.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error on get of vpc public gateway (%s) access tags: %s", d.Id(), err)
			}
			if err = d.Set("access_tags", accesstags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_public_gateway", "read", "set-access_tags").GetDiag()
			}

			controller, err := flex.GetBaseController(meta)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_public_gateway", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/publicGateways")
			if err = d.Set("resource_name", publicgw.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_public_gateway", "read", "set-resource_name").GetDiag()
			}

			if err = d.Set("resource_crn", publicgw.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_public_gateway", "read", "set-resource_crn").GetDiag()
			}
			if err = d.Set("crn", publicgw.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_public_gateway", "read", "set-crn").GetDiag()
			}
			if err = d.Set("resource_status", publicgw.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_status: %s", err), "(Data) ibm_is_public_gateway", "read", "set-resource_status").GetDiag()
			}
			if publicgw.ResourceGroup != nil {
				if err = d.Set("resource_group", publicgw.ResourceGroup.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_public_gateway", "read", "set-resource_group").GetDiag()
				}
				if err = d.Set("resource_group_name", publicgw.ResourceGroup.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_public_gateway", "read", "set-resource_group_name").GetDiag()
				}
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No Public gateway found with name: %s", name), "(Data) ibm_is_public_gateway", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
