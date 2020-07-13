package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func dataSourceIBMISInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISInstanceGroupManagerRead,

		Schema: map[string]*schema.Schema{

			"instance_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"instance_group_managers": {
				Type:        schema.TypeList,
				Description: "List of instance group managers",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the instance group manager.",
						},

						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the instance group manager.",
						},

						"manager_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of instance group manager.",
						},

						"aggregation_window": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The time window in seconds to aggregate metrics prior to evaluation",
						},

						"cooldown": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The duration of time in seconds to pause further scale actions after scaling has taken place",
						},

						"max_membership_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of members in a managed instance group",
						},
						"min_membership_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The minimum number of members in a managed instance group",
						},

						"policies": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "list of Policies associated with instancegroup manager",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupID := d.Get("instance_group_id").(string)

	listInstanceGroupManagerOptions := vpcv1.ListInstanceGroupManagersOptions{
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupManagerCollections, response, err := sess.ListInstanceGroupManagers(&listInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Manager: %s\n%s", err, response)
	}
	instanceGroupMnagers := make([]map[string]interface{}, 0)
	for _, instanceGroupManager := range instanceGroupManagerCollections.Managers {
		manager := map[string]interface{}{
			"id":                   *instanceGroupManager.ID,
			"name":                 *instanceGroupManager.Name,
			"aggregation_window":   *instanceGroupManager.AggregationWindow,
			"cooldown":             *instanceGroupManager.Cooldown,
			"max_membership_count": *instanceGroupManager.MaxMembershipCount,
			"min_membership_count": *instanceGroupManager.MinMembershipCount,
			"manager_type":         *instanceGroupManager.ManagerType,
		}

		policies := make([]string, 0)
		for i := 0; i < len(instanceGroupManager.Policies); i++ {
			policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
		}
		manager["policies"] = policies
		instanceGroupMnagers = append(instanceGroupMnagers, manager)
	}
	d.Set("instance_group_managers", instanceGroupMnagers)
	d.SetId(dataSourceIBMISInstanceGroupManagersID(d))
	return nil
}

// dataSourceIBMISInstanceGroupManagersID returns a reasonable ID for a instance group manager list.
func dataSourceIBMISInstanceGroupManagersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
