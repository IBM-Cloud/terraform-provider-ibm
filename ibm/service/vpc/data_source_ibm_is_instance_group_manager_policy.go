// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISInstanceGroupManagerPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupManagerPolicyRead,

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

			"name": {
				Type:        schema.TypeString,
				Required:    true,
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
	}
}

func dataSourceIBMISInstanceGroupManagerPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_manager_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)
	policyName := d.Get("name").(string)

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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupManagerPoliciesWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_group_manager_policy", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(instanceGroupManagerPolicyCollection.Next)
		allrecs = append(allrecs, instanceGroupManagerPolicyCollection.Policies...)
		if start == "" {
			break
		}
	}

	for _, data := range allrecs {
		instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
		if policyName == *instanceGroupManagerPolicy.Name {
			d.SetId(fmt.Sprintf("%s/%s/%s", instanceGroupID, instanceGroupManagerID, *instanceGroupManagerPolicy.ID))
			if err = d.Set("policy_id", *instanceGroupManagerPolicy.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policy_id: %s", err), "(Data) ibm_is_instance_group_manager_policy", "read", "set-policy_id").GetDiag()
			}
			if !core.IsNil(instanceGroupManagerPolicy.MetricValue) {
				if err = d.Set("metric_value", flex.IntValue(instanceGroupManagerPolicy.MetricValue)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metric_value: %s", err), "(Data) ibm_is_instance_group_manager_policy", "read", "set-metric_value").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManagerPolicy.MetricType) {
				if err = d.Set("metric_type", instanceGroupManagerPolicy.MetricType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metric_type: %s", err), "(Data) ibm_is_instance_group_manager_policy", "read", "set-metric_type").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManagerPolicy.PolicyType) {
				if err = d.Set("policy_type", instanceGroupManagerPolicy.PolicyType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policy_type: %s", err), "(Data) ibm_is_instance_group_manager_policy", "read", "set-policy_type").GetDiag()
				}
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance group manager policy %s not found", policyName), "(Data) ibm_is_instance_group_manager_policy", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
