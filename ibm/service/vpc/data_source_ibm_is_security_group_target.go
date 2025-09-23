// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISSecurityGroupTarget() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceIBMISSecurityGroupTargetRead,

		Schema: map[string]*schema.Schema{

			"security_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group id",
			},

			"target": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "security group target identifier",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group target name",
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this security group target",
			},

			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Type",
			},

			"more_info": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link to documentation about deleted resources",
			},
		},
	}
}

func dataSourceIBMISSecurityGroupTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_security_group_target", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	securityGroupID := d.Get("security_group").(string)
	name := d.Get("name").(string)

	// Support for pagination
	start := ""
	allrecs := []vpcv1.SecurityGroupTargetReferenceIntf{}

	for {
		listSecurityGroupTargetsOptions := sess.NewListSecurityGroupTargetsOptions(securityGroupID)
		if start != "" {
			listSecurityGroupTargetsOptions.Start = &start
		}
		groups, _, err := sess.ListSecurityGroupTargetsWithContext(context, listSecurityGroupTargetsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupTargetsWithContext failed %s", err), "(Data) ibm_is_security_group_target", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if *groups.TotalCount == int64(0) {
			break
		}

		start = flex.GetNext(groups.Next)
		allrecs = append(allrecs, groups.Targets...)

		if start == "" {
			break
		}

	}

	for _, securityGroupTargetReferenceIntf := range allrecs {
		securityGroupTargetReference := securityGroupTargetReferenceIntf.(*vpcv1.SecurityGroupTargetReference)
		if *securityGroupTargetReference.Name == name {
			if err = d.Set("target", *securityGroupTargetReference.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target %s", err), "(Data) ibm_is_security_group_target", "read", "set-target").GetDiag()
			}
			if err = d.Set("crn", securityGroupTargetReference.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn %s", err), "(Data) ibm_is_security_group_target", "read", "set-crn").GetDiag()
			}
			// d.Set("resource_type", *securityGroupTargetReference.ResourceType)
			if securityGroupTargetReference.Deleted != nil {
				if err = d.Set("more_info", *securityGroupTargetReference.Deleted.MoreInfo); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting targets %s", err), "(Data) ibm_is_security_group_target", "read", "set-more_info").GetDiag()
				}
			}
			if securityGroupTargetReference != nil && securityGroupTargetReference.ResourceType != nil {
				if err = d.Set("resource_type", *securityGroupTargetReference.ResourceType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type %s", err), "(Data) ibm_is_security_group_target", "read", "set-resource_type").GetDiag()
				}
			}
			d.SetId(fmt.Sprintf("%s/%s", securityGroupID, *securityGroupTargetReference.ID))

			return nil
		}
	}
	err = fmt.Errorf("Security Group Target %s not found", name)
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupTargetsWithContext failed %s", err), "(Data) ibm_is_security_group_target", "read")
	log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
