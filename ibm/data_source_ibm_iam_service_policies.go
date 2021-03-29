// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
)

// Data source to find all the policies for a serviceID
func dataSourceIBMIAMServicePolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMServicePoliciesRead,

		Schema: map[string]*schema.Schema{
			"iam_service_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"iam_service_id", "iam_id"},
				Description:  "UUID of ServiceID",
			},
			"iam_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"iam_service_id", "iam_id"},
				Description:  "IAM ID of ServiceID",
			},
			"sort": {
				Description: "Sort query for policies",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
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
					},
				},
			},
		},
	}
}

func dataSourceIBMIAMServicePoliciesRead(d *schema.ResourceData, meta interface{}) error {

	var iamID string
	if v, ok := d.GetOk("iam_service_id"); ok && v != nil {

		serviceIDUUID := v.(string)
		iamClient, err := meta.(ClientSession).IAMAPI()
		if err != nil {
			return err
		}
		serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
		if err != nil {
			return err
		}
		iamID = serviceID.IAMID
	}
	if v, ok := d.GetOk("iam_id"); ok && v != nil {
		iamID = v.(string)
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	query := iampapv1.SearchParams{
		AccountID: userDetails.userAccount,
		Type:      iampapv1.AccessPolicyType,
		IAMID:     iamID,
	}

	if v, ok := d.GetOk("sort"); ok {
		query.Sort = v.(string)
	}

	policies, err := iampapClient.V1Policy().List(query)
	if err != nil {
		return err
	}

	servicePolicies := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {
		roles := make([]string, len(policy.Roles))
		for i, role := range policy.Roles {
			roles[i] = role.Name
		}
		resources := flattenPolicyResource(policy.Resources)
		p := map[string]interface{}{
			"roles":     roles,
			"resources": resources,
		}
		if v, ok := d.GetOk("iam_service_id"); ok && v != nil {
			serviceIDUUID := v.(string)
			p["id"] = fmt.Sprintf("%s/%s", serviceIDUUID, policy.ID)
		} else if v, ok := d.GetOk("iam_id"); ok && v != nil {
			iamID := v.(string)
			p["id"] = fmt.Sprintf("%s/%s", iamID, policy.ID)
		}
		servicePolicies = append(servicePolicies, p)
	}

	if v, ok := d.GetOk("iam_service_id"); ok && v != nil {
		serviceIDUUID := v.(string)
		d.SetId(serviceIDUUID)
	} else if v, ok := d.GetOk("iam_id"); ok && v != nil {
		iamID := v.(string)
		d.SetId(iamID)
	}
	d.Set("policies", servicePolicies)
	return nil
}
