// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISInstanceGroupManagerAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceGroupManagerActionCreate,
		ReadContext:   resourceIBMISInstanceGroupManagerActionRead,
		UpdateContext: resourceIBMISInstanceGroupManagerActionUpdate,
		DeleteContext: resourceIBMISInstanceGroupManagerActionDelete,
		Exists:        resourceIBMISInstanceGroupManagerActionExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager_action", "name"),
				Description:  "instance group manager action name",
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
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The date and time the scheduled action will run.",
				ConflictsWith: []string{"cron_spec"},
			},

			"cron_spec": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_group_manager_action", "cron_spec"),
				Description:   "The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.",
				ConflictsWith: []string{"run_at"},
			},

			"membership_count": {
				Type:          schema.TypeInt,
				Optional:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_group_manager_action", "membership_count"),
				Description:   "The number of members the instance group should have at the scheduled time.",
				ConflictsWith: []string{"target_manager", "max_membership_count", "min_membership_count"},
				AtLeastOneOf:  []string{"target_manager", "membership_count"},
			},

			"max_membership_count": {
				Type:          schema.TypeInt,
				Optional:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_group_manager_action", "max_membership_count"),
				Description:   "The maximum number of members in a managed instance group",
				ConflictsWith: []string{"membership_count"},
				RequiredWith:  []string{"target_manager", "min_membership_count"},
			},

			"min_membership_count": {
				Type:          schema.TypeInt,
				Optional:      true,
				Default:       1,
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_group_manager_action", "min_membership_count"),
				Description:   "The minimum number of members in a managed instance group",
				ConflictsWith: []string{"membership_count"},
			},

			"target_manager": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The unique identifier for this instance group manager of type autoscale.",
				ConflictsWith: []string{"membership_count"},
				RequiredWith:  []string{"min_membership_count", "max_membership_count"},
				AtLeastOneOf:  []string{"target_manager", "membership_count"},
			},

			"target_manager_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance group manager name of type autoscale.",
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

func ResourceIBMISInstanceGroupManagerActionValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "max_membership_count",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "1000"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "min_membership_count",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "1000"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cron_spec",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Regexp:                     `^((((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5,7})$`,
			MinValueLength:             9,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "membership_count",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "100"})

	ibmISInstanceGroupManagerResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_group_manager_action", Schema: validateSchema}
	return &ibmISInstanceGroupManagerResourceValidator
}

func resourceIBMISInstanceGroupManagerActionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceGroupID := d.Get("instance_group").(string)
	instancegroupmanagerscheduledID := d.Get("instance_group_manager").(string)

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerActionOptions := vpcv1.CreateInstanceGroupManagerActionOptions{}
	instanceGroupManagerActionOptions.InstanceGroupID = &instanceGroupID
	instanceGroupManagerActionOptions.InstanceGroupManagerID = &instancegroupmanagerscheduledID

	instanceGroupManagerActionPrototype := vpcv1.InstanceGroupManagerActionPrototype{}

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		instanceGroupManagerActionPrototype.Name = &name
	}

	if v, ok := d.GetOk("run_at"); ok {
		runat := v.(string)
		datetime, err := strfmt.ParseDateTime(runat)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "create", "parse-run_at").GetDiag()
		}
		instanceGroupManagerActionPrototype.RunAt = &datetime
	}

	if v, ok := d.GetOk("cron_spec"); ok {
		cron_spec := v.(string)
		instanceGroupManagerActionPrototype.CronSpec = &cron_spec
	}

	if v, ok := d.GetOk("membership_count"); ok {
		membershipCount := int64(v.(int))
		instanceGroupManagerScheduledActionGroupPrototype := vpcv1.InstanceGroupManagerScheduledActionGroupPrototype{}
		instanceGroupManagerScheduledActionGroupPrototype.MembershipCount = &membershipCount
		instanceGroupManagerActionPrototype.Group = &instanceGroupManagerScheduledActionGroupPrototype
	}

	instanceGroupManagerScheduledActionByManagerManager := vpcv1.InstanceGroupManagerScheduledActionManagerPrototype{}
	if v, ok := d.GetOk("min_membership_count"); ok {
		minmembershipCount := int64(v.(int))
		instanceGroupManagerScheduledActionByManagerManager.MinMembershipCount = &minmembershipCount
	}

	if v, ok := d.GetOk("max_membership_count"); ok {
		maxmembershipCount := int64(v.(int))
		instanceGroupManagerScheduledActionByManagerManager.MaxMembershipCount = &maxmembershipCount
	}

	if v, ok := d.GetOk("target_manager"); ok {
		instanceGroupManagerAutoScale := v.(string)
		instanceGroupManagerScheduledActionByManagerManager.ID = &instanceGroupManagerAutoScale
		instanceGroupManagerActionPrototype.Manager = &instanceGroupManagerScheduledActionByManagerManager
	}

	instanceGroupManagerActionOptions.InstanceGroupManagerActionPrototype = &instanceGroupManagerActionPrototype

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutCreate))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupManagerActionIntf, _, err := sess.CreateInstanceGroupManagerActionWithContext(context, &instanceGroupManagerActionOptions)
	if err != nil || instanceGroupManagerActionIntf == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceGroupManagerActionWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupManagerAction := instanceGroupManagerActionIntf.(*vpcv1.InstanceGroupManagerAction)
	d.SetId(fmt.Sprintf("%s/%s/%s", instanceGroupID, instancegroupmanagerscheduledID, *instanceGroupManagerAction.ID))

	return resourceIBMISInstanceGroupManagerActionRead(context, d, meta)

}

func resourceIBMISInstanceGroupManagerActionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var changed bool
	instanceGroupManagerActionPatchModel := &vpcv1.InstanceGroupManagerActionPatch{}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		instanceGroupManagerActionPatchModel.Name = &name
		changed = true
	}

	if d.HasChange("cron_spec") {
		cronspec := d.Get("cron_spec").(string)
		instanceGroupManagerActionPatchModel.CronSpec = &cronspec
		changed = true
	}

	if d.HasChange("run_at") {
		runat := d.Get("run_at").(string)
		datetime, err := strfmt.ParseDateTime(runat)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "update", "parse-run_at").GetDiag()
		}
		instanceGroupManagerActionPatchModel.RunAt = &datetime
		changed = true
	}

	if d.HasChange("membership_count") {
		membershipCount := int64(d.Get("membership_count").(int))
		instanceGroupManagerScheduledActionGroupPatch := vpcv1.InstanceGroupManagerActionGroupPatch{}
		instanceGroupManagerScheduledActionGroupPatch.MembershipCount = &membershipCount
		instanceGroupManagerActionPatchModel.Group = &instanceGroupManagerScheduledActionGroupPatch
		changed = true
	}

	instanceGroupManagerScheduledActionByManagerPatchManager := vpcv1.InstanceGroupManagerActionManagerPatch{}

	if d.HasChange("min_membership_count") {
		minmembershipCount := int64(d.Get("min_membership_count").(int))
		instanceGroupManagerScheduledActionByManagerPatchManager.MinMembershipCount = &minmembershipCount
		changed = true
	}

	if d.HasChange("max_membership_count") {
		minmembershipCount := int64(d.Get("max_membership_count").(int))
		instanceGroupManagerScheduledActionByManagerPatchManager.MinMembershipCount = &minmembershipCount
		changed = true
	}
	instanceGroupManagerActionPatchModel.Manager = &instanceGroupManagerScheduledActionByManagerPatchManager

	if changed {

		parts, err := flex.IdParts(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "update", "sep-id-parts").GetDiag()
		}

		instanceGroupID := parts[0]
		instancegroupmanagerscheduledID := parts[1]
		instanceGroupManagerActionID := parts[2]

		updateInstanceGroupManagerActionOptions := &vpcv1.UpdateInstanceGroupManagerActionOptions{}
		updateInstanceGroupManagerActionOptions.InstanceGroupID = &instanceGroupID
		updateInstanceGroupManagerActionOptions.InstanceGroupManagerID = &instancegroupmanagerscheduledID
		updateInstanceGroupManagerActionOptions.ID = &instanceGroupManagerActionID

		instanceGroupManagerActionPatch, err := instanceGroupManagerActionPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupManagerActionPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateInstanceGroupManagerActionOptions.InstanceGroupManagerActionPatch = instanceGroupManagerActionPatch

		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_action", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, _, err = sess.UpdateInstanceGroupManagerActionWithContext(context, updateInstanceGroupManagerActionOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupManagerActionWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISInstanceGroupManagerActionRead(context, d, meta)
}

func resourceIBMISInstanceGroupManagerActionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instancegroupmanagerscheduledID := parts[1]
	instanceGroupManagerActionID := parts[2]

	getInstanceGroupManagerActionOptions := &vpcv1.GetInstanceGroupManagerActionOptions{
		InstanceGroupID:        &instanceGroupID,
		InstanceGroupManagerID: &instancegroupmanagerscheduledID,
		ID:                     &instanceGroupManagerActionID,
	}

	instanceGroupManagerActionIntf, response, err := sess.GetInstanceGroupManagerActionWithContext(context, getInstanceGroupManagerActionOptions)
	if err != nil || instanceGroupManagerActionIntf == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerActionWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupManagerAction := instanceGroupManagerActionIntf.(*vpcv1.InstanceGroupManagerAction)
	if err = d.Set("auto_delete", instanceGroupManagerAction.AutoDelete); err != nil {
		err = fmt.Errorf("Error setting auto_delete: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-auto_delete").GetDiag()
	}
	if err = d.Set("auto_delete_timeout", flex.IntValue(instanceGroupManagerAction.AutoDeleteTimeout)); err != nil {
		err = fmt.Errorf("Error setting auto_delete_timeout: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-auto_delete_timeout").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(instanceGroupManagerAction.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("action_id", *instanceGroupManagerAction.ID); err != nil {
		err = fmt.Errorf("Error setting action_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-action_id").GetDiag()
	}
	if !core.IsNil(instanceGroupManagerAction.Name) {
		if err = d.Set("name", instanceGroupManagerAction.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("resource_type", instanceGroupManagerAction.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("status", instanceGroupManagerAction.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-status").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(instanceGroupManagerAction.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-updated_at").GetDiag()
	}
	if !core.IsNil(instanceGroupManagerAction.ActionType) {
		if err = d.Set("action_type", instanceGroupManagerAction.ActionType); err != nil {
			err = fmt.Errorf("Error setting action_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-action_type").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerAction.CronSpec) {
		if err = d.Set("cron_spec", instanceGroupManagerAction.CronSpec); err != nil {
			err = fmt.Errorf("Error setting cron_spec: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-cron_spec").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerAction.LastAppliedAt) {
		if err = d.Set("last_applied_at", flex.DateTimeToString(instanceGroupManagerAction.LastAppliedAt)); err != nil {
			err = fmt.Errorf("Error setting last_applied_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-last_applied_at").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerAction.NextRunAt) {
		if err = d.Set("next_run_at", flex.DateTimeToString(instanceGroupManagerAction.NextRunAt)); err != nil {
			err = fmt.Errorf("Error setting next_run_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-next_run_at").GetDiag()
		}
	}
	instanceGroupManagerScheduledActionGroupGroup := instanceGroupManagerAction.Group
	if instanceGroupManagerScheduledActionGroupGroup != nil && instanceGroupManagerScheduledActionGroupGroup.MembershipCount != nil {
		if err = d.Set("membership_count", flex.IntValue(instanceGroupManagerScheduledActionGroupGroup.MembershipCount)); err != nil {
			err = fmt.Errorf("Error setting membership_count: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-membership_count").GetDiag()
		}
	}
	instanceGroupManagerScheduledActionManagerManagerInt := instanceGroupManagerAction.Manager
	if instanceGroupManagerScheduledActionManagerManagerInt != nil {
		instanceGroupManagerScheduledActionManagerManager := instanceGroupManagerScheduledActionManagerManagerInt.(*vpcv1.InstanceGroupManagerScheduledActionManager)
		if instanceGroupManagerScheduledActionManagerManager != nil && instanceGroupManagerScheduledActionManagerManager.ID != nil {

			if instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount != nil {
				if err = d.Set("max_membership_count", flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MaxMembershipCount)); err != nil {
					err = fmt.Errorf("Error setting max_membership_count: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-max_membership_count").GetDiag()
				}

			}
			if err = d.Set("min_membership_count", flex.IntValue(instanceGroupManagerScheduledActionManagerManager.MinMembershipCount)); err != nil {
				err = fmt.Errorf("Error setting min_membership_count: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-min_membership_count").GetDiag()
			}
			if err = d.Set("target_manager_name", *instanceGroupManagerScheduledActionManagerManager.Name); err != nil {
				err = fmt.Errorf("Error setting target_manager_name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-target_manager_name").GetDiag()
			}
			if err = d.Set("target_manager", *instanceGroupManagerScheduledActionManagerManager.ID); err != nil {
				err = fmt.Errorf("Error setting target_manager: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "read", "set-target_manager").GetDiag()
			}

		}
	}

	return nil
}

func resourceIBMISInstanceGroupManagerActionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "delete", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instancegroupmanagerscheduledID := parts[1]
	instanceGroupManagerActionID := parts[2]

	deleteInstanceGroupManagerActionOptions := &vpcv1.DeleteInstanceGroupManagerActionOptions{}
	deleteInstanceGroupManagerActionOptions.InstanceGroupID = &instanceGroupID
	deleteInstanceGroupManagerActionOptions.InstanceGroupManagerID = &instancegroupmanagerscheduledID
	deleteInstanceGroupManagerActionOptions.ID = &instanceGroupManagerActionID

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutDelete))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_action", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	response, err := sess.DeleteInstanceGroupManagerActionWithContext(context, deleteInstanceGroupManagerActionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceGroupManagerActionWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}

func resourceIBMISInstanceGroupManagerActionExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_action", "delete", "sep-id-parts")
	}
	instanceGroupID := parts[0]
	instancegroupmanagerscheduledID := parts[1]
	instanceGroupManagerActionID := parts[2]

	getInstanceGroupManagerActionOptions := &vpcv1.GetInstanceGroupManagerActionOptions{
		InstanceGroupID:        &instanceGroupID,
		InstanceGroupManagerID: &instancegroupmanagerscheduledID,
		ID:                     &instanceGroupManagerActionID,
	}

	_, response, err := sess.GetInstanceGroupManagerAction(getInstanceGroupManagerActionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerAction failed: %s", err.Error()), "ibm_is_instance_group_manager_action", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	return true, nil
}
