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

func DataSourceIBMISInstanceGroupManagerAction() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISInstanceGroupManagerActionRead,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
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

			"instance_group_manager": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance group manager ID of type scheduled",
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
	}
}

func dataSourceIBMISInstanceGroupManagerActionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_group_manager_action", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)
	actionName := d.Get("name").(string)

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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceGroupManagerActionsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_group_manager_action", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
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

	for _, data := range allrecs {
		instanceGroupManagerAction := data.(*vpcv1.InstanceGroupManagerAction)
		if actionName == *instanceGroupManagerAction.Name {
			d.SetId(fmt.Sprintf("%s/%s/%s", instanceGroupID, instanceGroupManagerID, *instanceGroupManagerAction.ID))

			if err = d.Set("auto_delete", instanceGroupManagerAction.AutoDelete); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-auto_delete").GetDiag()
			}

			if err = d.Set("auto_delete_timeout", flex.IntValue(instanceGroupManagerAction.AutoDeleteTimeout)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete_timeout: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-auto_delete_timeout").GetDiag()
			}
			if err = d.Set("created_at", flex.DateTimeToString(instanceGroupManagerAction.CreatedAt)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-created_at").GetDiag()
			}
			if err = d.Set("action_id", *instanceGroupManagerAction.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action_id: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-action_id").GetDiag()
			}
			if err = d.Set("resource_type", instanceGroupManagerAction.ResourceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-resource_type").GetDiag()
			}
			if err = d.Set("status", instanceGroupManagerAction.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-status").GetDiag()
			}
			if err = d.Set("updated_at", flex.DateTimeToString(instanceGroupManagerAction.UpdatedAt)); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting updated_at: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-updated_at").GetDiag()
			}
			if !core.IsNil(instanceGroupManagerAction.ActionType) {
				if err = d.Set("action_type", instanceGroupManagerAction.ActionType); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action_type: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-action_type").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManagerAction.CronSpec) {
				if err = d.Set("cron_spec", instanceGroupManagerAction.CronSpec); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cron_spec: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-cron_spec").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManagerAction.LastAppliedAt) {
				if err = d.Set("last_applied_at", flex.DateTimeToString(instanceGroupManagerAction.LastAppliedAt)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_applied_at: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-last_applied_at").GetDiag()
				}
			}
			if !core.IsNil(instanceGroupManagerAction.NextRunAt) {
				if err = d.Set("next_run_at", flex.DateTimeToString(instanceGroupManagerAction.NextRunAt)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting next_run_at: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-next_run_at").GetDiag()
				}
			}

			instanceGroupManagerScheduledActionGroupGroup := instanceGroupManagerAction.Group
			if instanceGroupManagerScheduledActionGroupGroup != nil && instanceGroupManagerScheduledActionGroupGroup.MembershipCount != nil {
				if err = d.Set("membership_count", flex.IntValue(instanceGroupManagerScheduledActionGroupGroup.MembershipCount)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting membership_count: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-membership_count").GetDiag()
				}
			}
			instanceGroupManagerScheduledActionManagerManagerInt := instanceGroupManagerAction.Manager
			if instanceGroupManagerScheduledActionManagerManagerInt != nil {
				instanceGroupManagerScheduledActionManagerManager := instanceGroupManagerScheduledActionManagerManagerInt.(*vpcv1.InstanceGroupManagerScheduledActionManager)
				if instanceGroupManagerScheduledActionManagerManager != nil && instanceGroupManagerScheduledActionManagerManager.ID != nil {

					if instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount != nil {
						if err = d.Set("max_membership_count", flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount)); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting max_membership_count: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-max_membership_count").GetDiag()
						}

					}
					if err = d.Set("min_membership_count", flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MinMembershipCount)); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting min_membership_count: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-min_membership_count").GetDiag()
					}
					if err = d.Set("target_manager_name", *instanceGroupManagerScheduledActionManagerManager.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target_manager_name: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-target_manager_name").GetDiag()
					}
					if err = d.Set("target_manager", *instanceGroupManagerScheduledActionManagerManager.ID); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target_manager: %s", err), "(Data) ibm_is_instance_group_manager_action", "read", "set-target_manager").GetDiag()
					}

				}
			}
			return nil
		}
	}
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance group manager action %s not found", actionName), "(Data) ibm_is_instance_group_manager_action", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
