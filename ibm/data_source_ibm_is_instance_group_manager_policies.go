package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func dataSourceIBMISInstanceGroupManagerPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceGroupManagerPolicyRead,

		Schema: map[string]*schema.Schema{

			"instance_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"instance_group_manager_id": {
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
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupManagerPolicyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerID := d.Get("instance_group_manager_id").(string)
	instanceGroupID := d.Get("instance_group_id").(string)

	listInstanceGroupManagerPoliciesOptions := vpcv1.ListInstanceGroupManagerPoliciesOptions{
		InstanceGroupID:        &instanceGroupID,
		InstanceGroupManagerID: &instanceGroupManagerID,
	}
	instanceGroupManagerPolicyCollection, response, err := sess.ListInstanceGroupManagerPolicies(&listInstanceGroupManagerPoliciesOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Manager Policies: %s\n%s", err, response)
	}
	policies := make([]map[string]interface{}, 0)
	for _, data := range instanceGroupManagerPolicyCollection.Policies {
		instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
		policy := map[string]interface{}{
			"id":           *instanceGroupManagerPolicy.ID,
			"name":         *instanceGroupManagerPolicy.Name,
			"metric_value": *instanceGroupManagerPolicy.MetricValue,
			"metric_type":  *instanceGroupManagerPolicy.MetricType,
			"policy_type":  *instanceGroupManagerPolicy.PolicyType,
		}
		policies = append(policies, policy)
	}
	d.Set("instance_group_manager_policies", policies)
	d.SetId(dataSourceIBMISInstanceGroupManagerPoliciesID(d))
	return nil
}

// dataSourceIBMISInstanceGroupManagerPoliciesID returns a reasonable ID for a instance group manager policies list.
func dataSourceIBMISInstanceGroupManagerPoliciesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
