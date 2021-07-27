// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMISInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupManagerCreate,
		Read:     resourceIBMISInstanceGroupManagerRead,
		Update:   resourceIBMISInstanceGroupManagerUpdate,
		Delete:   resourceIBMISInstanceGroupManagerDelete,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "name"),
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
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "manager_type"),
				Description:  "The type of instance group manager.",
			},

			"aggregation_window": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      90,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "aggregation_window"),
				Description:  "The time window in seconds to aggregate metrics prior to evaluation",
			},

			"cooldown": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      300,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "cooldown"),
				Description:  "The duration of time in seconds to pause further scale actions after scaling has taken place",
			},

			"max_membership_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "max_membership_count"),
				Description:  "The maximum number of members in a managed instance group",
			},

			"min_membership_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_manager", "min_membership_count"),
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

func resourceIBMISInstanceGroupManagerValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	managerType := "autoscale, scheduled"
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "manager_type",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              managerType})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "aggregation_window",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "90",
			MaxValue:                   "600"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "cooldown",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "120",
			MaxValue:                   "3600"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "max_membership_count",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "1",
			MaxValue:                   "1000"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "min_membership_count",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "1",
			MaxValue:                   "1000"})

	ibmISInstanceGroupManagerResourceValidator := ResourceValidator{ResourceName: "ibm_is_instance_group_manager", Schema: validateSchema}
	return &ibmISInstanceGroupManagerResourceValidator
}

func resourceIBMISInstanceGroupManagerCreate(d *schema.ResourceData, meta interface{}) error {

	instanceGroupID := d.Get("instance_group").(string)
	managerType := d.Get("manager_type").(string)

	sess, err := vpcClient(meta)
	if err != nil {
		return err
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
		instanceGroupManagerIntf, response, err := sess.CreateInstanceGroupManager(&createInstanceGroupManagerOptions)
		if err != nil || instanceGroupManagerIntf == nil {
			return fmt.Errorf("Error creating InstanceGroup manager: %s\n%s", err, response)
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
			return healthError
		}

		instanceGroupManagerIntf, response, err := sess.CreateInstanceGroupManager(&createInstanceGroupManagerOptions)
		if err != nil || instanceGroupManagerIntf == nil {
			return fmt.Errorf("Error creating InstanceGroup manager: %s\n%s", err, response)
		}
		instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)

		d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, *instanceGroupManager.ID))

	}

	return resourceIBMISInstanceGroupManagerRead(d, meta)

}

func resourceIBMISInstanceGroupManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
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
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		instanceGroupID := parts[0]
		instanceGroupManagerID := parts[1]
		updateInstanceGroupManagerOptions.ID = &instanceGroupManagerID
		updateInstanceGroupManagerOptions.InstanceGroupID = &instanceGroupID
		instanceGroupManagerPatch, err := instanceGroupManagerPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("Error calling asPatch for InstanceGroupManagerPatch: %s", err)
		}
		updateInstanceGroupManagerOptions.InstanceGroupManagerPatch = instanceGroupManagerPatch

		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			return healthError
		}

		_, response, err := sess.UpdateInstanceGroupManager(&updateInstanceGroupManagerOptions)
		if err != nil {
			return fmt.Errorf("Error updating InstanceGroup manager: %s\n%s", err, response)
		}
	}
	return resourceIBMISInstanceGroupManagerRead(d, meta)
}

func resourceIBMISInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]

	getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupManagerIntf, response, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
	if err != nil || instanceGroupManagerIntf == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Manager: %s\n%s", err, response)
	}
	instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)

	managerType := *instanceGroupManager.ManagerType

	if managerType == "scheduled" {
		d.Set("name", *instanceGroupManager.Name)
		d.Set("enable_manager", *instanceGroupManager.ManagementEnabled)
		d.Set("manager_id", instanceGroupManagerID)
		d.Set("instance_group", instanceGroupID)
		d.Set("manager_type", *instanceGroupManager.ManagerType)

	} else {

		d.Set("name", *instanceGroupManager.Name)
		d.Set("aggregation_window", *instanceGroupManager.AggregationWindow)
		d.Set("cooldown", *instanceGroupManager.Cooldown)
		d.Set("max_membership_count", *instanceGroupManager.MaxMembershipCount)
		d.Set("min_membership_count", *instanceGroupManager.MinMembershipCount)
		d.Set("enable_manager", *instanceGroupManager.ManagementEnabled)
		d.Set("manager_id", instanceGroupManagerID)
		d.Set("instance_group", instanceGroupID)
		d.Set("manager_type", *instanceGroupManager.ManagerType)

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
	d.Set("actions", actions)

	policies := make([]string, 0)

	for i := 0; i < len(instanceGroupManager.Policies); i++ {
		policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
	}
	d.Set("policies", policies)
	return nil
}

func resourceIBMISInstanceGroupManagerDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]

	deleteInstanceGroupManagerOptions := vpcv1.DeleteInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutDelete))
	if healthError != nil {
		return healthError
	}

	response, err := sess.DeleteInstanceGroupManager(&deleteInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Deleting the InstanceGroup Manager: %s\n%s", err, response)
	}
	return nil
}
