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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceGroupManagerCreate,
		ReadContext:   resourceIBMISInstanceGroupManagerRead,
		UpdateContext: resourceIBMISInstanceGroupManagerUpdate,
		DeleteContext: resourceIBMISInstanceGroupManagerDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "name"),
				Description:  "instance group manager name",
			},

			"enable_manager": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "enable instance group manager",
			},

			"instance_group": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"manager_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "autoscale",
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "manager_type"),
				Description:  "The type of instance group manager.",
			},

			"aggregation_window": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      90,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "aggregation_window"),
				Description:  "The time window in seconds to aggregate metrics prior to evaluation",
			},

			"cooldown": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      300,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "cooldown"),
				Description:  "The duration of time in seconds to pause further scale actions after scaling has taken place",
			},

			"max_membership_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "max_membership_count"),
				Description:  "The maximum number of members in a managed instance group",
			},

			"min_membership_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager", "min_membership_count"),
				Description:  "The minimum number of members in a managed instance group",
			},

			"manager_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "instance group manager ID",
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

func ResourceIBMISInstanceGroupManagerValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	managerType := "autoscale, scheduled"
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
			Identifier:                 "manager_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              managerType})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "aggregation_window",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "90",
			MaxValue:                   "600"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cooldown",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "120",
			MaxValue:                   "3600"})
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

	ibmISInstanceGroupManagerResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_group_manager", Schema: validateSchema}
	return &ibmISInstanceGroupManagerResourceValidator
}

func resourceIBMISInstanceGroupManagerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	instanceGroupID := d.Get("instance_group").(string)
	managerType := d.Get("manager_type").(string)

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if managerType == "scheduled" {
		instanceGroupManagerPrototype := vpcv1.InstanceGroupManagerPrototypeInstanceGroupManagerScheduledPrototype{}
		instanceGroupManagerPrototype.ManagerType = &managerType

		if v, ok := d.GetOk("name"); ok {
			name := v.(string)
			instanceGroupManagerPrototype.Name = &name
		}

		if v, ok := d.GetOk("enable_manager"); ok {
			enableManager := v.(bool)
			instanceGroupManagerPrototype.ManagementEnabled = &enableManager
		}

		createInstanceGroupManagerOptions := vpcv1.CreateInstanceGroupManagerOptions{
			InstanceGroupID:               &instanceGroupID,
			InstanceGroupManagerPrototype: &instanceGroupManagerPrototype,
		}
		instanceGroupManagerIntf, _, err := sess.CreateInstanceGroupManagerWithContext(context, &createInstanceGroupManagerOptions)
		if err != nil || instanceGroupManagerIntf == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceGroupManagerWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)
		d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupManager.ID))

	} else {

		instanceGroupManagerPrototype := vpcv1.InstanceGroupManagerPrototypeInstanceGroupManagerAutoScalePrototype{}
		instanceGroupManagerPrototype.ManagerType = &managerType

		if v, ok := d.GetOk("name"); ok {
			name := v.(string)
			instanceGroupManagerPrototype.Name = &name
		}

		if v, ok := d.GetOk("enable_manager"); ok {
			enableManager := v.(bool)
			instanceGroupManagerPrototype.ManagementEnabled = &enableManager
		}

		if v, ok := d.GetOk("aggregation_window"); ok {
			aggregationWindow := int64(v.(int))
			instanceGroupManagerPrototype.AggregationWindow = &aggregationWindow
		}

		if v, ok := d.GetOk("cooldown"); ok {
			cooldown := int64(v.(int))
			instanceGroupManagerPrototype.Cooldown = &cooldown
		}

		if v, ok := d.GetOk("min_membership_count"); ok {
			minMembershipCount := int64(v.(int))
			instanceGroupManagerPrototype.MinMembershipCount = &minMembershipCount
		}

		if v, ok := d.GetOk("max_membership_count"); ok {
			maxMembershipCount := int64(v.(int))
			instanceGroupManagerPrototype.MaxMembershipCount = &maxMembershipCount
		}

		createInstanceGroupManagerOptions := vpcv1.CreateInstanceGroupManagerOptions{
			InstanceGroupID:               &instanceGroupID,
			InstanceGroupManagerPrototype: &instanceGroupManagerPrototype,
		}

		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutCreate))
		if healthError != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", err.Error()), "ibm_is_instance_group_manager", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		instanceGroupManagerIntf, _, err := sess.CreateInstanceGroupManagerWithContext(context, &createInstanceGroupManagerOptions)
		if err != nil || instanceGroupManagerIntf == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceGroupManagerWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)

		d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupManager.ID))

	}

	return resourceIBMISInstanceGroupManagerRead(context, d, meta)

}

func resourceIBMISInstanceGroupManagerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	managerType := d.Get("manager_type").(string)

	var changed bool
	updateInstanceGroupManagerOptions := vpcv1.UpdateInstanceGroupManagerOptions{}
	instanceGroupManagerPatchModel := &vpcv1.InstanceGroupManagerPatch{}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		instanceGroupManagerPatchModel.Name = &name
		changed = true
	}
	if managerType == "autoscale" {
		if d.HasChange("aggregation_window") {
			aggregationWindow := int64(d.Get("aggregation_window").(int))
			instanceGroupManagerPatchModel.AggregationWindow = &aggregationWindow
			changed = true
		}

		if d.HasChange("cooldown") {
			cooldown := int64(d.Get("cooldown").(int))
			instanceGroupManagerPatchModel.Cooldown = &cooldown
			changed = true
		}

		if d.HasChange("max_membership_count") {
			maxMembershipCount := int64(d.Get("max_membership_count").(int))
			instanceGroupManagerPatchModel.MaxMembershipCount = &maxMembershipCount
			changed = true
		}

		if d.HasChange("min_membership_count") {
			minMembershipCount := int64(d.Get("min_membership_count").(int))
			instanceGroupManagerPatchModel.MinMembershipCount = &minMembershipCount
			changed = true
		}
	}

	if d.HasChange("enable_manager") {
		enableManager := d.Get("enable_manager").(bool)
		instanceGroupManagerPatchModel.ManagementEnabled = &enableManager
		changed = true
	}

	if changed {
		parts, err := flex.IdParts(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "update", "sep-id-parts").GetDiag()
		}
		instanceGroupID := parts[0]
		instanceGroupManagerID := parts[1]
		updateInstanceGroupManagerOptions.ID = &instanceGroupManagerID
		updateInstanceGroupManagerOptions.InstanceGroupID = &instanceGroupID
		instanceGroupManagerPatch, err := instanceGroupManagerPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupManagerPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_group_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateInstanceGroupManagerOptions.InstanceGroupManagerPatch = instanceGroupManagerPatch

		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", err.Error()), "ibm_is_instance_group_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, _, err = sess.UpdateInstanceGroupManagerWithContext(context, &updateInstanceGroupManagerOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupManagerWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISInstanceGroupManagerRead(context, d, meta)
}

func resourceIBMISInstanceGroupManagerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]

	getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupManagerIntf, response, err := sess.GetInstanceGroupManagerWithContext(context, &getInstanceGroupManagerOptions)
	if err != nil || instanceGroupManagerIntf == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)

	managerType := *instanceGroupManager.ManagerType

	if managerType == "scheduled" {
		if !core.IsNil(instanceGroupManager.Name) {
			if err = d.Set("name", instanceGroupManager.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-name").GetDiag()
			}
		}
		if !core.IsNil(instanceGroupManager.ManagementEnabled) {
			if err = d.Set("enable_manager", instanceGroupManager.ManagementEnabled); err != nil {
				err = fmt.Errorf("Error setting enable_manager: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-enable_manager").GetDiag()
			}
		}
		if err = d.Set("manager_id", instanceGroupManagerID); err != nil {
			err = fmt.Errorf("Error setting manager_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-manager_id").GetDiag()
		}
		if err = d.Set("instance_group", instanceGroupID); err != nil {
			err = fmt.Errorf("Error setting instance_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-instance_group").GetDiag()
		}
		if err = d.Set("manager_type", *instanceGroupManager.ManagerType); err != nil {
			err = fmt.Errorf("Error setting manager_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-manager_type").GetDiag()
		}
	} else {
		if !core.IsNil(instanceGroupManager.Name) {
			if err = d.Set("name", instanceGroupManager.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-name").GetDiag()
			}
		}
		d.Set("aggregation_window", *instanceGroupManager.AggregationWindow)
		if !core.IsNil(instanceGroupManager.AggregationWindow) {
			if err = d.Set("aggregation_window", flex.IntValue(instanceGroupManager.AggregationWindow)); err != nil {
				err = fmt.Errorf("Error setting aggregation_window: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-aggregation_window").GetDiag()
			}
		}
		if !core.IsNil(instanceGroupManager.Cooldown) {
			if err = d.Set("cooldown", flex.IntValue(instanceGroupManager.Cooldown)); err != nil {
				err = fmt.Errorf("Error setting cooldown: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-cooldown").GetDiag()
			}
		}
		if !core.IsNil(instanceGroupManager.MaxMembershipCount) {
			if err = d.Set("max_membership_count", flex.IntValue(instanceGroupManager.MaxMembershipCount)); err != nil {
				err = fmt.Errorf("Error setting max_membership_count: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-max_membership_count").GetDiag()
			}
		}
		if !core.IsNil(instanceGroupManager.MinMembershipCount) {
			if err = d.Set("min_membership_count", flex.IntValue(instanceGroupManager.MinMembershipCount)); err != nil {
				err = fmt.Errorf("Error setting min_membership_count: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-min_membership_count").GetDiag()
			}
		}
		if !core.IsNil(instanceGroupManager.ManagementEnabled) {
			if err = d.Set("enable_manager", instanceGroupManager.ManagementEnabled); err != nil {
				err = fmt.Errorf("Error setting enable_manager: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-enable_manager").GetDiag()
			}
		}
		if err = d.Set("manager_id", instanceGroupManagerID); err != nil {
			err = fmt.Errorf("Error setting manager_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-manager_id").GetDiag()
		}
		if err = d.Set("instance_group", instanceGroupID); err != nil {
			err = fmt.Errorf("Error setting instance_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-instance_group").GetDiag()
		}
		if err = d.Set("manager_type", *instanceGroupManager.ManagerType); err != nil {
			err = fmt.Errorf("Error setting manager_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-manager_type").GetDiag()
		}

	}

	actions := make([]map[string]interface{}, 0)
	if instanceGroupManager.Actions != nil {
		for _, action := range instanceGroupManager.Actions {
			actn := map[string]interface{}{
				"instance_group_manager_action":      *action.ID,
				"instance_group_manager_action_name": *action.Name,
				"resource_type":                      *action.ResourceType,
			}
			actions = append(actions, actn)
		}
	}
	if err = d.Set("actions", actions); err != nil {
		err = fmt.Errorf("Error setting actions: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-actions").GetDiag()
	}

	policies := make([]string, 0)

	for i := 0; i < len(instanceGroupManager.Policies); i++ {
		policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
	}
	if err = d.Set("policies", policies); err != nil {
		err = fmt.Errorf("Error setting policies: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "read", "set-policies").GetDiag()
	}
	return nil
}

func resourceIBMISInstanceGroupManagerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager", "delete", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]

	deleteInstanceGroupManagerOptions := vpcv1.DeleteInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutDelete))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", err.Error()), "ibm_is_instance_group_manager", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	response, err := sess.DeleteInstanceGroupManagerWithContext(context, &deleteInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceGroupManagerWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}
