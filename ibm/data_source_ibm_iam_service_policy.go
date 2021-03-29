// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Data source to find the policY for a serviceID, policy ID
func dataSourceIBMIAMServicePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMServicePolicyRead,

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IAM ID of ServiceID",
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Role names of the policy definition",
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service name of the policy definition",
						},
						"resource_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "ID of resource instance of the policy definition",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region of the policy definition",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type of the policy definition",
						},
						"resource": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource of the policy definition",
						},
						"resource_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the resource group.",
						},
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMIAMServicePolicyRead(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	servicepolicyID := d.Get("policy_id").(string)
	parts, err := idParts(servicepolicyID)
	if err != nil {
		return err
	}
	policyID := parts[1]

	servicePolicy, err := iampapClient.V1Policy().Get(policyID)
	if err != nil {
		return fmt.Errorf("Error retrieving servicePolicy: %s", err)
	}

	d.SetId(policyID)
	roles := make([]string, len(servicePolicy.Roles))
	for i, role := range servicePolicy.Roles {
		roles[i] = role.Name
	}
	d.Set("roles", roles)
	d.Set("version", servicePolicy.Version)
	d.Set("resources", flattenPolicyResource(servicePolicy.Resources))
	return nil
}
