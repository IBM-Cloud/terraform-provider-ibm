// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMISInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupManagerRead,

		Schema: map[string]*schema.Schema{

			"instance_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the instance group manager.",
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

			"manager_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of instance group manager.",
			},

			"policies": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of Policies associated with instancegroup manager",
			},

			"actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_group_manager_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_group_manager_action_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupManagerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_manager", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupID := d.Get("instance_group").(string)

	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroupManagerIntf{}

	for {
		listInstanceGroupManagerOptions := vpcv1.ListInstanceGroupManagersOptions{
			InstanceGroupID: &instanceGroupID,
		}
		if start != "" {
			listInstanceGroupManagerOptions.Start = &start
		}
		instanceGroupManagerCollections, _, err := sess.ListInstanceGroupManagersWithContext(context, &listInstanceGroupManagerOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_group_manager", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(instanceGroupManagerCollections.Next)
		allrecs = append(allrecs, instanceGroupManagerCollections.Managers...)

		if start == "" {
			break
		}
	}

	instanceGroupManagerName := d.Get("name").(string)
	for _, instanceGroupManagerIntf := range allrecs {
		instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)
		if instanceGroupManagerName == *instanceGroupManager.Name {
			d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupManager.ID))
			if !core.IsNil(instanceGroupManager.ManagerType) {
				if err = d.Set("manager_type", instanceGroupManager.ManagerType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting manager_type: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-manager_type").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManager.ID) {
				if err = d.Set("manager_id", *instanceGroupManager.ID); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting manager_id: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-manager_id").GetDiag()
				}
			}

			if *instanceGroupManager.ManagerType == "scheduled" {

				actions := make([]map[string]interface{}, 0)
				if instanceGroupManager.Actions != nil {
					for _, action := range instanceGroupManager.Actions {
						actn := map[string]interface{}{
							"instance_group_manager_action":      action.ID,
							"instance_group_manager_action_name": action.Name,
							"resource_type":                      action.ResourceType,
						}
						actions = append(actions, actn)
					}

					if err = d.Set("actions", actions); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting actions: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-actions").GetDiag()
					}
				}

			} else {
				if !core.IsNil(instanceGroupManager.AggregationWindow) {
					if err = d.Set("aggregation_window", flex.IntValue(instanceGroupManager.AggregationWindow)); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting aggregation_window: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-aggregation_window").GetDiag()
					}
				}
				if !core.IsNil(instanceGroupManager.Cooldown) {
					if err = d.Set("cooldown", flex.IntValue(instanceGroupManager.Cooldown)); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cooldown: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-cooldown").GetDiag()
					}
				}
				if !core.IsNil(instanceGroupManager.MaxMembershipCount) {
					if err = d.Set("max_membership_count", flex.IntValue(instanceGroupManager.MaxMembershipCount)); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting max_membership_count: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-max_membership_count").GetDiag()
					}
				}
				if !core.IsNil(instanceGroupManager.MinMembershipCount) {
					if err = d.Set("min_membership_count", flex.IntValue(instanceGroupManager.MinMembershipCount)); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting min_membership_count: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-min_membership_count").GetDiag()
					}
				}
				policies := make([]string, 0)
				if instanceGroupManager.Policies != nil {
					for i := 0; i < len(instanceGroupManager.Policies); i++ {
						policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
					}
				}
				if err = d.Set("policies", policies); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policies: %s", err), "(Data) ibm_is_instance_group_manager", "read", "set-policies").GetDiag()
				}
			}

			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance group manager %s not found", instanceGroupManagerName), "(Data) ibm_is_instance_group_manager", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
