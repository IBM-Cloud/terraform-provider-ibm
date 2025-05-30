// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISInstanceGroupManagerPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupManagerPoliciesRead,

		Schema: map[string]*schema.Schema{

			"instance_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"instance_group_manager": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance group manager ID",
			},

			"instance_group_manager_policies": {
				Type:        schema.TypeList,
				Description: "List of instance group manager policies",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the instance group manager policy.",
						},

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance group manager policy",
						},

						"metric_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of metric to be evaluated",
						},

						"metric_value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The metric value to be evaluated",
						},

						"policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of Policy for the Instance Group",
						},
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy ID",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupManagerPoliciesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_manager_policies", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)

	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroupManagerPolicyIntf{}

	for {
		listInstanceGroupManagerPoliciesOptions := vpcv1.ListInstanceGroupManagerPoliciesOptions{
			InstanceGroupID:        &instanceGroupID,
			InstanceGroupManagerID: &instanceGroupManagerID,
		}
		if start != "" {
			listInstanceGroupManagerPoliciesOptions.Start = &start
		}
		instanceGroupManagerPolicyCollection, _, err := sess.ListInstanceGroupManagerPoliciesWithContext(context, &listInstanceGroupManagerPoliciesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupManagerPoliciesWithContext failed %s", err), "(Data) ibm_is_instance_group_manager_policies", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(instanceGroupManagerPolicyCollection.Next)
		allrecs = append(allrecs, instanceGroupManagerPolicyCollection.Policies...)
		if start == "" {
			break
		}
	}

	policies := make([]map[string]interface{}, 0)
	for _, data := range allrecs {
		instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
		policy := map[string]interface{}{
			"id":           fmt.Sprintf("%s/%s/%s", instanceGroupID, instanceGroupManagerID, *instanceGroupManagerPolicy.ID),
			"name":         *instanceGroupManagerPolicy.Name,
			"metric_value": *instanceGroupManagerPolicy.MetricValue,
			"metric_type":  *instanceGroupManagerPolicy.MetricType,
			"policy_type":  *instanceGroupManagerPolicy.PolicyType,
			"policy_id":    *instanceGroupManagerPolicy.ID,
		}
		policies = append(policies, policy)
	}
	if err = d.Set("instance_group_manager_policies", policies); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_group_manager_policies %s", err), "(Data) ibm_is_instance_group_manager_policies", "read", "instance_group_manager_policies-set").GetDiag()
	}
	d.SetId(dataSourceIBMISInstanceGroupManagerPoliciesID(d))
	return nil
}

// dataSourceIBMISInstanceGroupManagerPoliciesID returns a reasonable ID for a instance group manager policies list.
func dataSourceIBMISInstanceGroupManagerPoliciesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
