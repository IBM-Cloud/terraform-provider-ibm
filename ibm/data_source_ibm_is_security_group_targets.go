// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMISSecurityGroupTargets() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMISSecurityGroupTargetsRead,

		Schema: map[string]*schema.Schema{

			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group id",
				ForceNew:    true,
			},

			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     50,
				Description: "The maximum number of resources that can be returned by the request",
			},

			"targets": {
				Type:        schema.TypeList,
				Description: "List of subnets",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"target_id": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "security group target identifier",
							ForceNew:     true,
							ValidateFunc: InvokeValidator("ibm_is_security_group_target", isSecurityGroupTargetID),
						},

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group name",
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

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources across all pages",
			},
		},
	}
}

func dataSourceIBMISSecurityGroupTargetsRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	var limit int64
	if lim, ok := d.GetOk("limit"); ok {
		limit = int64(lim.(int))
	}

	securityGroupID := d.Get("security_group_id").(string)
	d.SetId(fmt.Sprintf("%s", securityGroupID))
	listSecurityGroupTargetsOptions := sess.NewListSecurityGroupTargetsOptions(securityGroupID)

	listSecurityGroupTargetsOptions.Limit = &limit

	groups, response, err := sess.ListSecurityGroupTargets(listSecurityGroupTargetsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Security Group Targets : %s\n%s", err, response)
	}

	d.Set("limit", *groups.Limit)
	d.Set("total_count", *groups.TotalCount)

	targets := make([]map[string]interface{}, 0)
	if groups.Targets != nil {
		for _, target := range groups.Targets {

			tr := target.(*vpcv1.SecurityGroupTargetReference)
			var moreinfo string
			if tr.Deleted != nil {
				moreinfo = *tr.Deleted.MoreInfo
			}
			if tr != nil {
				l := map[string]interface{}{
					"name":          *tr.Name,
					"target_id":     *tr.ID,
					"resource_type": *tr.ResourceType,
					"more_info":     moreinfo,
				}
				if l != nil {
					targets = append(targets, l)
				}
			}
		}
	}
	if targets != nil {
		d.Set("targets", targets)
	}

	return nil
}
