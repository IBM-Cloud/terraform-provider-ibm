// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMISSecurityGroupTarget() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMISSecurityGroupTargetRead,

		Schema: map[string]*schema.Schema{

			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Security group id",
			},

			"target_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "security group target identifier",
				ValidateFunc: InvokeValidator("ibm_is_security_group_target", "target_id"),
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
	}
}

func dataSourceIBMISSecurityGroupTargetRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	securityGroupID := d.Get("security_group_id").(string)
	securityGroupTargetID := d.Get("target_id").(string)

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}

	data, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Security Group Target : %s\n%s", err, response)
	}

	target := data.(*vpcv1.SecurityGroupTargetReference)

	if target.Deleted != nil {
		d.Set("more_info", target.Deleted.MoreInfo)
	}

	d.Set("name", *target.Name)
	d.Set(isSecurityGroupResourceType, *target.ResourceType)
	d.SetId(fmt.Sprintf("%s/%s", securityGroupID, *target.ID))

	return nil
}
