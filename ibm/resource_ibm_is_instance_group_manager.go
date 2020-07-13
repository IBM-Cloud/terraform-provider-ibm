package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func resourceIBMISInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupManagerCreate,
		Read:     resourceIBMISInstanceGroupManagerRead,
		Update:   resourceIBMISInstanceGroupManagerUpdate,
		Delete:   resourceIBMISInstanceGroupManagerDelete,
		Exists:   resourceIBMISInstanceGroupManagerExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance group manager name",
			},

			"instance_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance group ID",
			},

			"manager_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "autoscale",
				Description: "The type of instance group manager.",
			},

			"aggregation_window": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     90,
				Description: "The time window in seconds to aggregate metrics prior to evaluation",
			},

			"cooldown": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     300,
				Description: "The duration of time in seconds to pause further scale actions after scaling has taken place",
			},

			"max_membership_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The maximum number of members in a managed instance group",
			},

			"min_membership_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The minimum number of members in a managed instance group",
			},

			"policies": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of Policies associated with instancegroup manager",
			},
		},
	}
}

func resourceIBMISInstanceGroupManagerCreate(d *schema.ResourceData, meta interface{}) error {

	instanceGroupID := d.Get("instance_group_id").(string)
	maxMembershipCount := int64(d.Get("max_membership_count").(int))

	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerPrototype := vpcv1.InstanceGroupManagerPrototype{}
	//instanceGroupManagerPrototype.ManagerType = &managerType
	instanceGroupManagerPrototype.MaxMembershipCount = &maxMembershipCount

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		instanceGroupManagerPrototype.Name = &name
	}

	if v, ok := d.GetOk("manager_type"); ok {
		managerType := v.(string)
		instanceGroupManagerPrototype.ManagerType = &managerType
	}

	if v, ok := d.GetOk("min_membership_count"); ok {
		minMembershipCount := int64(v.(int))
		instanceGroupManagerPrototype.MinMembershipCount = &minMembershipCount
	}

	if v, ok := d.GetOk("cooldown"); ok {
		cooldown := int64(v.(int))
		instanceGroupManagerPrototype.Cooldown = &cooldown
	}

	if v, ok := d.GetOk("aggregation_window"); ok {
		aggregationWindow := int64(v.(int))
		instanceGroupManagerPrototype.AggregationWindow = &aggregationWindow
	}

	//managerPrototype := &vpcv1.InstanceGroupManagerPrototypeIntf{}
	createInstanceGroupManagerOptions := vpcv1.CreateInstanceGroupManagerOptions{
		InstanceGroupID:               &instanceGroupID,
		InstanceGroupManagerPrototype: &instanceGroupManagerPrototype,
	}
	instanceGroupManager, response, err := sess.CreateInstanceGroupManager(&createInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Creating InstanceGroup Manager: %s\n%s", err, response)
	}
	d.SetId(*instanceGroupManager.ID)

	return resourceIBMISInstanceGroupManagerUpdate(d, meta)

}

func resourceIBMISInstanceGroupManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	var changed bool
	updateInstanceGroupManagerOptions := vpcv1.UpdateInstanceGroupManagerOptions{}

	if d.HasChange("name") && !d.IsNewResource() {
		name := d.Get("name").(string)
		updateInstanceGroupManagerOptions.Name = &name
		changed = true
	}

	if d.HasChange("aggregation_window") && !d.IsNewResource() {
		aggregationWindow := int64(d.Get("aggregation_window").(int))
		updateInstanceGroupManagerOptions.AggregationWindow = &aggregationWindow
		changed = true
	}

	if d.HasChange("cooldown") && !d.IsNewResource() {
		cooldown := int64(d.Get("cooldown").(int))
		updateInstanceGroupManagerOptions.Cooldown = &cooldown
		changed = true
	}

	if d.HasChange("max_membership_count") && !d.IsNewResource() {
		maxMembershipCount := int64(d.Get("max_membership_count").(int))
		updateInstanceGroupManagerOptions.MaxMembershipCount = &maxMembershipCount
		changed = true
	}

	if d.HasChange("min_membership_count") && !d.IsNewResource() {
		minMembershipCount := int64(d.Get("min_membership_count").(int))
		updateInstanceGroupManagerOptions.MinMembershipCount = &minMembershipCount
		changed = true
	}

	if changed {
		instanceGroupManagerID := d.Id()
		instanceGroupID := d.Get("instance_group_id").(string)
		updateInstanceGroupManagerOptions.ID = &instanceGroupManagerID
		updateInstanceGroupManagerOptions.InstanceGroupID = &instanceGroupID
		_, response, err := sess.UpdateInstanceGroupManager(&updateInstanceGroupManagerOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error Updating InstanceGroup Manager: %s\n%s", err, response)
		}
	}
	return resourceIBMISInstanceGroupManagerRead(d, meta)
}

func resourceIBMISInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerID := d.Id()
	instanceGroupID := d.Get("instance_group_id").(string)

	getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupManager, response, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Manager: %s\n%s", err, response)
	}
	d.Set("name", *instanceGroupManager.Name)
	d.Set("aggregation_window", *instanceGroupManager.AggregationWindow)
	d.Set("cooldown", *instanceGroupManager.Cooldown)
	d.Set("max_membership_count", *instanceGroupManager.MaxMembershipCount)
	d.Set("min_membership_count", *instanceGroupManager.MinMembershipCount)

	policies := make([]string, 0)

	for i := 0; i < len(instanceGroupManager.Policies); i++ {
		policies = append(policies, string(*(instanceGroupManager.Policies[i].ID)))
	}
	d.Set("policies", policies)
	return nil
}

func resourceIBMISInstanceGroupManagerDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	instanceGroupManagerID := d.Id()
	instanceGroupID := d.Get("instance_group_id").(string)
	deleteInstanceGroupManagerOptions := vpcv1.DeleteInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
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

func resourceIBMISInstanceGroupManagerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := myvpcClient(meta)
	if err != nil {
		return false, err
	}

	instanceGroupManagerID := d.Id()
	instanceGroupID := d.Get("instance_group_id").(string)

	getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
		ID:              &instanceGroupManagerID,
		InstanceGroupID: &instanceGroupID,
	}

	_, response, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error Getting InstanceGroup Manager: %s\n%s", err, response)
	}
	return true, nil
}
