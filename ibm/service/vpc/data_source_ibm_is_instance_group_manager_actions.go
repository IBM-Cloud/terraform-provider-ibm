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

func DataSourceIBMISInstanceGroupManagerActions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupManagerActionsRead,

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

			"instance_group_manager_actions": {
				Type:        schema.TypeList,
				Description: "List of instance group manager actions",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "instance group manager action name",
						},

						"action_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance group manager action ID",
						},

						"instance_group": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "instance group ID",
						},

						"run_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the scheduled action will run.",
						},

						"membership_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of members the instance group should have at the scheduled time.",
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

						"target_manager": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance group manager of type autoscale.",
						},

						"target_manager_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance group manager name of type autoscale.",
						},

						"cron_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.",
						},

						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the instance group action- `active`: Action is ready to be run- `completed`: Action was completed successfully- `failed`: Action could not be completed successfully- `incompatible`: Action parameters are not compatible with the group or manager- `omitted`: Action was not applied because this action's manager was disabled.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the instance group manager action was modified.",
						},
						"action_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of action for the instance group.",
						},

						"last_applied_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the scheduled action was last applied. If empty the action has never been applied.",
						},
						"next_run_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.",
						},
						"auto_delete": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"auto_delete_timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the instance group manager action was modified.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISInstanceGroupManagerActionsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_manager_actions", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)

	// Support for pagination
	start := ""
	allrecs := []vpcv1.InstanceGroupManagerActionIntf{}

	for {
		listInstanceGroupManagerActionsOptions := vpcv1.ListInstanceGroupManagerActionsOptions{
			InstanceGroupID:        &instanceGroupID,
			InstanceGroupManagerID: &instanceGroupManagerID,
		}
		if start != "" {
			listInstanceGroupManagerActionsOptions.Start = &start
		}
		instanceGroupManagerActionsCollection, _, err := sess.ListInstanceGroupManagerActionsWithContext(context, &listInstanceGroupManagerActionsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupManagerActionsWithContext failed %s", err), "(Data) ibm_is_instance_group_manager_actions", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if instanceGroupManagerActionsCollection != nil && *instanceGroupManagerActionsCollection.TotalCount == int64(0) {
			break
		}
		start = flex.GetNext(instanceGroupManagerActionsCollection.Next)
		allrecs = append(allrecs, instanceGroupManagerActionsCollection.Actions...)
		if start == "" {
			break
		}
	}

	actions := make([]map[string]interface{}, 0)
	for _, data := range allrecs {
		instanceGroupManagerAction := data.(*vpcv1.InstanceGroupManagerAction)

		action := map[string]interface{}{
			"name":                *instanceGroupManagerAction.Name,
			"auto_delete":         *instanceGroupManagerAction.AutoDelete,
			"auto_delete_timeout": flex.IntValue(instanceGroupManagerAction.AutoDeleteTimeout),
			"created_at":          instanceGroupManagerAction.CreatedAt.String(),
			"action_id":           *instanceGroupManagerAction.ID,
			"resource_type":       *instanceGroupManagerAction.ResourceType,
			"status":              *instanceGroupManagerAction.Status,
			"updated_at":          instanceGroupManagerAction.UpdatedAt.String(),
			"action_type":         *instanceGroupManagerAction.ActionType,
		}
		if instanceGroupManagerAction.CronSpec != nil {
			action["cron_spec"] = *instanceGroupManagerAction.CronSpec
		}
		if instanceGroupManagerAction.LastAppliedAt != nil {
			action["last_applied_at"] = instanceGroupManagerAction.LastAppliedAt.String()
		}
		if instanceGroupManagerAction.NextRunAt != nil {
			action["last_applied_at"] = instanceGroupManagerAction.NextRunAt.String()
		}
		instanceGroupManagerScheduledActionGroupGroup := instanceGroupManagerAction.Group
		if instanceGroupManagerScheduledActionGroupGroup != nil && instanceGroupManagerScheduledActionGroupGroup.MembershipCount != nil {
			action["membership_count"] = flex.IntValue(instanceGroupManagerScheduledActionGroupGroup.MembershipCount)
		}
		instanceGroupManagerScheduledActionManagerManagerInt := instanceGroupManagerAction.Manager
		if instanceGroupManagerScheduledActionManagerManagerInt != nil {
			instanceGroupManagerScheduledActionManagerManager := instanceGroupManagerScheduledActionManagerManagerInt.(*vpcv1.InstanceGroupManagerScheduledActionManager)
			if instanceGroupManagerScheduledActionManagerManager != nil && instanceGroupManagerScheduledActionManagerManager.ID != nil {

				if instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount != nil {
					action["max_membership_count"] = flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount)
				}
				action["min_membership_count"] = flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MinMembershipCount)
				action["target_manager_name"] = *instanceGroupManagerScheduledActionManagerManager.Name
				action["target_manager"] = *instanceGroupManagerScheduledActionManagerManager.ID
			}
		}
		actions = append(actions, action)
	}
	if err = d.Set("instance_group_manager_actions", actions); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting actions %s", err), "(Data) ibm_is_instance_group_manager_actions", "read", "actions-set").GetDiag()
	}

	d.SetId(dataSourceIBMISInstanceGroupManagerActionsID(d))
	return nil
}

func dataSourceIBMISInstanceGroupManagerActionsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
