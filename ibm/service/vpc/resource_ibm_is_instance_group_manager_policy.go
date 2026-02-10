// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISInstanceGroupManagerPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceGroupManagerPolicyCreate,
		ReadContext:   resourceIBMISInstanceGroupManagerPolicyRead,
		UpdateContext: resourceIBMISInstanceGroupManagerPolicyUpdate,
		DeleteContext: resourceIBMISInstanceGroupManagerPolicyDelete,
		Exists:        resourceIBMISInstanceGroupManagerPolicyExists,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager_policy", "name"),
				Description:  "instance group manager policy name",
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager_policy", "metric_type"),
				Description:  "The type of metric to be evaluated",
			},

			"metric_value": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The metric value to be evaluated",
			},

			"policy_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_manager_policy", "policy_type"),
				Description:  "The type of Policy for the Instance Group",
			},

			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Policy ID",
			},
		},
	}
}

func ResourceIBMISInstanceGroupManagerPolicyValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	metricTypes := "cpu,memory,network_in,network_out"
	policyType := "target"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "metric_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              metricTypes})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "policy_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              policyType})

	ibmISInstanceGroupManagerPolicyResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_group_manager_policy", Schema: validateSchema}
	return &ibmISInstanceGroupManagerPolicyResourceValidator
}

func resourceIBMISInstanceGroupManagerPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceGroupID := d.Get("instance_group").(string)
	instanceGroupManagerID := d.Get("instance_group_manager").(string)

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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

	isInsGrpKey := "Instance_Group_Key_" + instanceGroupID
	conns.IbmMutexKV.Lock(isInsGrpKey)
	defer conns.IbmMutexKV.Unlock(isInsGrpKey)

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutCreate))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	data, _, err := sess.CreateInstanceGroupManagerPolicyWithContext(context, &createInstanceGroupManagerPolicyOptions)
	if err != nil || data == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceGroupManagerPolicyWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceGroupID, instanceGroupManagerID, *instanceGroupManagerPolicy.ID))

	return resourceIBMISInstanceGroupManagerPolicyRead(context, d, meta)

}

func resourceIBMISInstanceGroupManagerPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var changed bool
	updateInstanceGroupManagerPolicyOptions := vpcv1.UpdateInstanceGroupManagerPolicyOptions{}
	instanceGroupManagerPolicyPatchModel := &vpcv1.InstanceGroupManagerPolicyPatch{}
	if d.HasChange("name") {
		name := d.Get("name").(string)
		instanceGroupManagerPolicyPatchModel.Name = &name
		changed = true
	}

	if d.HasChange("metric_type") {
		metricType := d.Get("metric_type").(string)
		instanceGroupManagerPolicyPatchModel.MetricType = &metricType
		changed = true
	}

	if d.HasChange("metric_value") {
		metricValue := int64(d.Get("metric_value").(int))
		instanceGroupManagerPolicyPatchModel.MetricValue = &metricValue
		changed = true
	}

	if changed {
		parts, err := flex.IdParts(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "update", "sep-id-parts").GetDiag()
		}
		instanceGroupID := parts[0]
		instanceGroupManagerID := parts[1]
		instanceGroupManagerPolicyID := parts[2]

		updateInstanceGroupManagerPolicyOptions.ID = &instanceGroupManagerPolicyID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupID = &instanceGroupID
		updateInstanceGroupManagerPolicyOptions.InstanceGroupManagerID = &instanceGroupManagerID

		instanceGroupManagerPolicyAsPatch, asPatchErr := instanceGroupManagerPolicyPatchModel.AsPatch()
		if asPatchErr != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupManagerPolicyPatchModel.AsPatch() failed: %s", asPatchErr.Error()), "ibm_is_instance_group_manager_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateInstanceGroupManagerPolicyOptions.InstanceGroupManagerPolicyPatch = instanceGroupManagerPolicyAsPatch

		isInsGrpKey := "Instance_Group_Key_" + instanceGroupID
		conns.IbmMutexKV.Lock(isInsGrpKey)
		defer conns.IbmMutexKV.Unlock(isInsGrpKey)

		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, _, err = sess.UpdateInstanceGroupManagerPolicyWithContext(context, &updateInstanceGroupManagerPolicyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupManagerPolicyWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISInstanceGroupManagerPolicyRead(context, d, meta)
}

func resourceIBMISInstanceGroupManagerPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]
	instanceGroupManagerPolicyID := parts[2]

	getInstanceGroupManagerPolicyOptions := vpcv1.GetInstanceGroupManagerPolicyOptions{
		ID:                     &instanceGroupManagerPolicyID,
		InstanceGroupID:        &instanceGroupID,
		InstanceGroupManagerID: &instanceGroupManagerID,
	}
	data, response, err := sess.GetInstanceGroupManagerPolicyWithContext(context, &getInstanceGroupManagerPolicyOptions)
	if err != nil || data == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerPolicyWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupManagerPolicy := data.(*vpcv1.InstanceGroupManagerPolicy)
	if !core.IsNil(instanceGroupManagerPolicy.Name) {
		if err = d.Set("name", instanceGroupManagerPolicy.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerPolicy.MetricValue) {
		if err = d.Set("metric_value", flex.IntValue(instanceGroupManagerPolicy.MetricValue)); err != nil {
			err = fmt.Errorf("Error setting metric_value: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-metric_value").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerPolicy.MetricType) {
		if err = d.Set("metric_type", instanceGroupManagerPolicy.MetricType); err != nil {
			err = fmt.Errorf("Error setting metric_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-metric_type").GetDiag()
		}
	}
	if !core.IsNil(instanceGroupManagerPolicy.PolicyType) {
		if err = d.Set("policy_type", instanceGroupManagerPolicy.PolicyType); err != nil {
			err = fmt.Errorf("Error setting policy_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-policy_type").GetDiag()
		}
	}
	if err = d.Set("policy_id", instanceGroupManagerPolicyID); err != nil {
		err = fmt.Errorf("Error setting policy_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-policy_id").GetDiag()
	}
	if err = d.Set("instance_group", instanceGroupID); err != nil {
		err = fmt.Errorf("Error setting instance_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-instance_group").GetDiag()
	}
	if d.Set("instance_group_manager", instanceGroupManagerID); err != nil {
		err = fmt.Errorf("Error setting instance_group_manager: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "read", "set-instance_group_manager").GetDiag()
	}

	return nil
}

func resourceIBMISInstanceGroupManagerPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "delete", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]
	instanceGroupManagerPolicyID := parts[2]

	deleteInstanceGroupManagerPolicyOptions := vpcv1.DeleteInstanceGroupManagerPolicyOptions{
		ID:                     &instanceGroupManagerPolicyID,
		InstanceGroupManagerID: &instanceGroupManagerID,
		InstanceGroupID:        &instanceGroupID,
	}

	isInsGrpKey := "Instance_Group_Key_" + instanceGroupID
	conns.IbmMutexKV.Lock(isInsGrpKey)
	defer conns.IbmMutexKV.Unlock(isInsGrpKey)

	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutDelete))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group_manager_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	response, err := sess.DeleteInstanceGroupManagerPolicyWithContext(context, &deleteInstanceGroupManagerPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceGroupManagerPolicyWithContext failed: %s", err.Error()), "ibm_is_instance_group_manager_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}

func resourceIBMISInstanceGroupManagerPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "exists", "sep-id-parts")
	}

	if len(parts) != 3 {
		err = fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of instanceGroupID/instanceGroupManagerID/instanceGroupManagerPolicyID", d.Id())
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_manager_policy", "exists", "sep-id-parts")
	}
	instanceGroupID := parts[0]
	instanceGroupManagerID := parts[1]
	instanceGroupManagerPolicyID := parts[2]

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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupManagerPolicy failed: %s", err.Error()), "ibm_is_instance_group_manager_policy", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
