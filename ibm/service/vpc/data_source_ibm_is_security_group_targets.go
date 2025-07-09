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

func DataSourceIBMISSecurityGroupTargets() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceIBMISSecurityGroupTargetsRead,

		Schema: map[string]*schema.Schema{

			"security_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group id",
			},

			"targets": {
				Type:        schema.TypeList,
				Description: "List of targets",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"target": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "security group target identifier",
						},

						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this target",
						},

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group target name",
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
				},
			},
		},
	}
}

func dataSourceIBMISSecurityGroupTargetsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_security_group_targets", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	securityGroupID := d.Get("security_group").(string)

	// Support for pagination
	start := ""
	allrecs := []vpcv1.SecurityGroupTargetReferenceIntf{}

	for {
		listSecurityGroupTargetsOptions := sess.NewListSecurityGroupTargetsOptions(securityGroupID)
		if start != "" {
			listSecurityGroupTargetsOptions.Start = &start
		}
		groups, _, err := sess.ListSecurityGroupTargetsWithContext(context, listSecurityGroupTargetsOptions)
		if err != nil || groups == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SecurityGroupTargetsPager.GetAll() failed %s", err), "(Data) ibm_is_security_group_targets", "read")
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

	targets := make([]map[string]interface{}, 0)
	for _, securityGroupTargetReferenceIntf := range allrecs {
		securityGroupTargetReference := securityGroupTargetReferenceIntf.(*vpcv1.SecurityGroupTargetReference)
		tr := map[string]interface{}{
			"name":   *securityGroupTargetReference.Name,
			"target": *securityGroupTargetReference.ID,
			"crn":    securityGroupTargetReference.CRN,
			// "resource_type": *securityGroupTargetReference.ResourceType,
		}
		if securityGroupTargetReference.Deleted != nil {
			tr["more_info"] = *securityGroupTargetReference.Deleted.MoreInfo
		}
		if securityGroupTargetReference != nil && securityGroupTargetReference.ResourceType != nil {
			tr["resource_type"] = *securityGroupTargetReference.ResourceType
		}
		targets = append(targets, tr)
	}
	if err = d.Set("targets", targets); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting targets %s", err), "(Data) ibm_is_security_group_targets", "read", "targets-set").GetDiag()
	}
	d.SetId(securityGroupID)
	return nil
}
