package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func resourceIBMISInstanceGroupManagerPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupManagerPolicyCreate,
		Read:     resourceIBMISInstanceGroupManagerPolicyRead,
		Update:   resourceIBMISInstanceGroupManagerPolicyUpdate,
		Delete:   resourceIBMISInstanceGroupManagerPolicyDelete,
		Exists:   resourceIBMISInstanceGroupManagerPolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance group manager policy name",
			},

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

			"metric_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of metric to be evaluated",
			},

			"metric_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The metric value to be evaluated",
			},

			"policy_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Policy for the Instance Group",
			},
		},
	}
}

func resourceIBMISInstanceGroupManagerPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	instanceGroupID := d.Get("instance_group_id").(string)
	instanceGroupManagerID := d.Get("instance_group_manager_id").(string)

	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerPolicyPrototype := vpcv1.InstanceGroupManagerPolicyPrototype{}

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		instanceGroupManagerPolicyPrototype.Name = &name
	}

	if v, ok := d.GetOk("metric_type"); ok {
		metricType := v.(string)
		instanceGroupManagerPolicyPrototype.MetricType = &metricType
	}

	if v, ok := d.GetOk("metric_value"); ok {
		metricValue := int64(v.(int))
		instanceGroupManagerPolicyPrototype.MetricValue = &metricValue
	}

	if v, ok := d.GetOk("policy_type"); ok {
		policyType := v.(string)
		instanceGroupManagerPolicyPrototype.PolicyType = &policyType
	}

	createInstanceGroupManagerPolicyOptions := vpcv1.CreateInstanceGroupManagerPolicyOptions{
		InstanceGroupID:                     &instanceGroupID,
		InstanceGroupManagerID:              &instanceGroupManagerID,
		InstanceGroupManagerPolicyPrototype: &instanceGroupManagerPolicyPrototype,
	}
	data, response, err := sess.CreateInstanceGroupManagerPolicy(&createInstanceGroupManagerPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Creating InstanceGroup Manager Policy: %s\n%s", err, response)
	}
	instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
	d.SetId(*instanceGroupManagerPolicy.ID)

	return resourceIBMISInstanceGroupManagerPolicyUpdate(d, meta)

}

func resourceIBMISInstanceGroupManagerPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	var changed bool
	updateInstanceGroupManagerPolicyOptions := vpcv1.UpdateInstanceGroupManagerPolicyOptions{}

	if d.HasChange("name") && !d.IsNewResource() {
		name := d.Get("name").(string)
		updateInstanceGroupManagerPolicyOptions.Name = &name
		changed = true
	}

	if d.HasChange("metric_type") && !d.IsNewResource() {
		metricType := d.Get("metric_type").(string)
		updateInstanceGroupManagerPolicyOptions.MetricType = &metricType
		changed = true
	}

	if d.HasChange("metric_value") && !d.IsNewResource() {
		metricValue := int64(d.Get("metric_value").(int))
		updateInstanceGroupManagerPolicyOptions.MetricValue = &metricValue
		changed = true
	}

	if changed {
		instanceGroupManagerPolicyID := d.Id()
		instanceGroupID := d.Get("instance_group_id").(string)
		instanceGroupManagerID := d.Get("instance_group_manager_id").(string)
		updateInstanceGroupManagerPolicyOptions.ID = &instanceGroupManagerPolicyID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupID = &instanceGroupID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupManagerID = &instanceGroupManagerID

		_, response, err := sess.UpdateInstanceGroupManagerPolicy(&updateInstanceGroupManagerPolicyOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error Updating InstanceGroup Manager Policy: %s\n%s", err, response)
		}
	}
	return resourceIBMISInstanceGroupManagerPolicyRead(d, meta)
}

func resourceIBMISInstanceGroupManagerPolicyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerPolicyID := d.Id()
	instanceGroupManagerID := d.Get("instance_group_manager_id").(string)
	instanceGroupID := d.Get("instance_group_id").(string)

	getInstanceGroupManagerPolicyOptions := vpcv1.GetInstanceGroupManagerPolicyOptions{
		ID:                     &instanceGroupManagerPolicyID,
		InstanceGroupID:        &instanceGroupID,
		InstanceGroupManagerID: &instanceGroupManagerID,
	}
	data, response, err := sess.GetInstanceGroupManagerPolicy(&getInstanceGroupManagerPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Manager Policy: %s\n%s", err, response)
	}
	instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
	d.Set("name", *instanceGroupManagerPolicy.Name)
	d.Set("metric_value", instanceGroupManagerPolicy.MetricValue)
	d.Set("metric_type", instanceGroupManagerPolicy.MetricType)
	d.Set("policy_type", instanceGroupManagerPolicy.PolicyType)

	return nil
}

func resourceIBMISInstanceGroupManagerPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}
	instanceGroupManagerPolicyID := d.Id()
	instanceGroupID := d.Get("instance_group_id").(string)
	instanceGroupManagerID := d.Get("instance_group_manager_id").(string)
	deleteInstanceGroupManagerPolicyOptions := vpcv1.DeleteInstanceGroupManagerPolicyOptions{
		ID:                     &instanceGroupManagerPolicyID,
		InstanceGroupManagerID: &instanceGroupManagerID,
		InstanceGroupID:        &instanceGroupID,
	}
	response, err := sess.DeleteInstanceGroupManagerPolicy(&deleteInstanceGroupManagerPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Deleting the InstanceGroup Manager Policy: %s\n%s", err, response)
	}
	return nil
}

func resourceIBMISInstanceGroupManagerPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := myvpcClient(meta)
	if err != nil {
		return false, err
	}

	instanceGroupManagerPolicyID := d.Id()
	instanceGroupManagerID := d.Get("instance_group_manager_id").(string)
	instanceGroupID := d.Get("instance_group_id").(string)

	getInstanceGroupManagerPolicyOptions := vpcv1.GetInstanceGroupManagerPolicyOptions{
		ID:                     &instanceGroupManagerPolicyID,
		InstanceGroupManagerID: &instanceGroupManagerID,
		InstanceGroupID:        &instanceGroupID,
	}

	_, response, err := sess.GetInstanceGroupManagerPolicy(&getInstanceGroupManagerPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error Getting InstanceGroup Manager Policy: %s\n%s", err, response)
	}
	return true, nil
}
