package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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

			"metric_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of metric to be evaluated",
			},

			"metric_value": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The metric value to be evaluated",
			},

			"policy_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of Policy for the Instance Group",
			},
		},
	}
}

func resourceIBMISInstanceGroupManagerPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	instanceGroupID := d.Get("instance_group").(string)
	instanceGroupManagerID := d.Get("instance_group_manager").(string)

	sess, err := myvpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupManagerPolicyPrototype := vpcv1.InstanceGroupManagerPolicyPrototype{}

	name := d.Get("name").(string)
	metricType := d.Get("metric_type").(string)
	metricValue := int64(d.Get("metric_value").(int))
	policyType := d.Get("policy_type").(string)

	instanceGroupManagerPolicyPrototype.Name = &name
	instanceGroupManagerPolicyPrototype.MetricType = &metricType
	instanceGroupManagerPolicyPrototype.MetricValue = &metricValue
	instanceGroupManagerPolicyPrototype.PolicyType = &policyType

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

	// Construct an instance of the InstanceGroupManagerPolicyPatch model
	instanceGroupManagerPolicyPatchModel := new(vpcv1.InstanceGroupManagerPolicyPatch)

	if d.HasChange("name") && !d.IsNewResource() {
		name := d.Get("name").(string)
		instanceGroupManagerPolicyPatchModel.Name = &name
		changed = true
	}

	if d.HasChange("metric_type") && !d.IsNewResource() {
		metricType := d.Get("metric_type").(string)
		instanceGroupManagerPolicyPatchModel.MetricType = &metricType
		changed = true
	}

	if d.HasChange("metric_value") && !d.IsNewResource() {
		metricValue := int64(d.Get("metric_value").(int))
		instanceGroupManagerPolicyPatchModel.MetricValue = &metricValue
		changed = true
	}

	if changed {
		instanceGroupManagerPolicyPatchModelAsPatch, _ := instanceGroupManagerPolicyPatchModel.AsPatch()

		instanceGroupManagerPolicyID := d.Id()
		instanceGroupID := d.Get("instance_group").(string)
		instanceGroupManagerID := d.Get("instance_group_manager").(string)
		updateInstanceGroupManagerPolicyOptions.ID = &instanceGroupManagerPolicyID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupID = &instanceGroupID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupManagerID = &instanceGroupManagerID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupManagerPolicyPatch = instanceGroupManagerPolicyPatchModelAsPatch

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
	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)

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
	instanceGroupID := d.Get("instance_group").(string)
	instanceGroupManagerID := d.Get("instance_group_manager").(string)
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
	instanceGroupManagerID := d.Get("instance_group_manager").(string)
	instanceGroupID := d.Get("instance_group").(string)

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
