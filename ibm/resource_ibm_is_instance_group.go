// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// SCALING ...
	SCALING = "scaling"
	// HEALTHY ...
	HEALTHY = "healthy"
	// DELETING ...
	DELETING = "deleting"
	// UNHEALTHY
	UNHEALTHY = "unhealthy"
)

func resourceIBMISInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupCreate,
		Read:     resourceIBMISInstanceGroupRead,
		Update:   resourceIBMISInstanceGroupUpdate,
		Delete:   resourceIBMISInstanceGroupDelete,
		Exists:   resourceIBMISInstanceGroupExists,
		Importer: &schema.ResourceImporter{},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group", "name"),
				Description:  "The user-defined name for this instance group",
			},

			"instance_template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance template ID",
			},

			"instance_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: InvokeValidator("ibm_is_instance_group", "instance_count"),
				Description:  "The number of instances in the instance group",
			},

			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID",
			},

			"subnets": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: "list of subnet IDs",
			},

			"application_port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group", "application_port"),
				Description:  "Used by the instance group when scaling up instances to supply the port for the load balancer pool member.",
			},

			"load_balancer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "load balancer ID",
			},

			"load_balancer_pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "load balancer pool ID",
			},

			"managers": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "list of Managers associated with instancegroup",
			},

			"instances": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "number of instances in the intances group",
			},

			"vpc": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "vpc instance",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance group status - deleting, healthy, scaling, unhealthy",
			},

			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: InvokeValidator("ibm_is_instance_group", "tag")},
				Set:         resourceIBMVPCHash,
				Description: "List of tags for instance group",
			},
		},
	}
}

func resourceIBMISInstanceGroupValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "instance_count",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "0",
			MaxValue:                   "1000"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "application_port",
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "tag",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISInstanceGroupResourceValidator := ResourceValidator{ResourceName: "ibm_is_instance_group", Schema: validateSchema}
	return &ibmISInstanceGroupResourceValidator
}

func resourceIBMISInstanceGroupCreate(d *schema.ResourceData, meta interface{}) error {

	name := d.Get("name").(string)
	instanceTemplate := d.Get("instance_template").(string)

	subnets := d.Get("subnets")

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	var subnetIDs []vpcv1.SubnetIdentityIntf
	for _, s := range subnets.([]interface{}) {
		subnet := s.(string)
		subnetIDs = append(subnetIDs, &vpcv1.SubnetIdentity{ID: &subnet})
	}

	instanceGroupOptions := vpcv1.CreateInstanceGroupOptions{
		InstanceTemplate: &vpcv1.InstanceTemplateIdentity{
			ID: &instanceTemplate,
		},
		Subnets: subnetIDs,
		Name:    &name,
	}

	var membershipCount int
	if v, ok := d.GetOk("instance_count"); ok {
		membershipCount = v.(int)
		mc := int64(membershipCount)
		instanceGroupOptions.MembershipCount = &mc
	}

	if v, ok := d.GetOk("load_balancer"); ok {
		lbID := v.(string)
		instanceGroupOptions.LoadBalancer = &vpcv1.LoadBalancerIdentity{ID: &lbID}
	}

	if v, ok := d.GetOk("load_balancer_pool"); ok {
		lbPoolID := v.(string)
		instanceGroupOptions.LoadBalancerPool = &vpcv1.LoadBalancerPoolIdentity{ID: &lbPoolID}
	}

	if v, ok := d.GetOk("resource_group"); ok {
		resourceGroup := v.(string)
		instanceGroupOptions.ResourceGroup = &vpcv1.ResourceGroupIdentity{ID: &resourceGroup}
	}

	if v, ok := d.GetOk("application_port"); ok {
		applicatioPort := int64(v.(int))
		instanceGroupOptions.ApplicationPort = &applicatioPort
	}

	instanceGroup, response, err := sess.CreateInstanceGroup(&instanceGroupOptions)
	if err != nil || instanceGroup == nil {
		return fmt.Errorf("Error Creating InstanceGroup: %s\n%s", err, response)
	}
	d.SetId(*instanceGroup.ID)

	_, healthError := waitForHealthyInstanceGroup(d, meta, d.Timeout(schema.TimeoutCreate))
	if healthError != nil {
		return healthError
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = UpdateTagsUsingCRN(oldList, newList, meta, *instanceGroup.CRN)
		if err != nil {
			log.Printf(
				"Error on create of instance group (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISInstanceGroupRead(d, meta)

}

func resourceIBMISInstanceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	var changed bool
	instanceGroupUpdateOptions := vpcv1.UpdateInstanceGroupOptions{}
	instanceGroupPatchModel := vpcv1.InstanceGroupPatch{}

	if d.HasChange("tags") {
		instanceGroupID := d.Id()
		getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
		instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
		if err != nil || instanceGroup == nil {
			return fmt.Errorf("Error getting instance group: %s\n%s", err, response)
		}
		oldList, newList := d.GetChange("tags")
		err = UpdateTagsUsingCRN(oldList, newList, meta, *instanceGroup.CRN)
		if err != nil {
			log.Printf(
				"Error on update of instance group (%s) tags: %s", d.Id(), err)
		}
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		instanceGroupPatchModel.Name = &name
		changed = true
	}

	if d.HasChange("instance_template") {
		instanceTemplate := d.Get("instance_template").(string)
		instanceGroupPatchModel.InstanceTemplate = &vpcv1.InstanceTemplateIdentity{
			ID: &instanceTemplate,
		}
		changed = true
	}

	if d.HasChange("instance_count") {
		membershipCount := d.Get("instance_count").(int)
		mc := int64(membershipCount)
		instanceGroupPatchModel.MembershipCount = &mc
		changed = true
	}

	if d.HasChange("subnets") {
		subnets := d.Get("subnets")
		var subnetIDs []vpcv1.SubnetIdentityIntf
		for _, s := range subnets.([]interface{}) {
			subnet := s.(string)
			subnetIDs = append(subnetIDs, &vpcv1.SubnetIdentity{ID: &subnet})
		}
		instanceGroupPatchModel.Subnets = subnetIDs
		changed = true
	}

	if d.HasChange("application_port") || d.HasChange("load_balancer") || d.HasChange("load_balancer_pool") {
		applicationPort := int64(d.Get("application_port").(int))
		lbID := d.Get("load_balancer").(string)
		lbPoolID := d.Get("load_balancer_pool").(string)
		instanceGroupPatchModel.ApplicationPort = &applicationPort
		instanceGroupPatchModel.LoadBalancer = &vpcv1.LoadBalancerIdentity{ID: &lbID}
		instanceGroupPatchModel.LoadBalancerPool = &vpcv1.LoadBalancerPoolIdentity{ID: &lbPoolID}
		changed = true
	}

	if changed {
		instanceGroupID := d.Id()
		instanceGroupUpdateOptions.ID = &instanceGroupID
		instanceGroupPatch, err := instanceGroupPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("Error calling asPatch for InstanceGroupPatch: %s", err)
		}
		instanceGroupUpdateOptions.InstanceGroupPatch = instanceGroupPatch
		_, response, err := sess.UpdateInstanceGroup(&instanceGroupUpdateOptions)
		if err != nil {
			return fmt.Errorf("Error Updating InstanceGroup: %s\n%s", err, response)
		}

		// wait for instance group health update with update timeout configured.
		_, healthError := waitForHealthyInstanceGroup(d, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			return healthError
		}
	}
	return resourceIBMISInstanceGroupRead(d, meta)
}

func resourceIBMISInstanceGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
	if err != nil || instanceGroup == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
	}
	d.Set("name", *instanceGroup.Name)
	d.Set("instance_template", *instanceGroup.InstanceTemplate.ID)
	d.Set("instances", *instanceGroup.MembershipCount)
	d.Set("resource_group", *instanceGroup.ResourceGroup.ID)
	if instanceGroup.ApplicationPort != nil {
		d.Set("application_port", *instanceGroup.ApplicationPort)
	}

	subnets := make([]string, 0)

	for i := 0; i < len(instanceGroup.Subnets); i++ {
		subnets = append(subnets, string(*(instanceGroup.Subnets[i].ID)))
	}
	if instanceGroup.LoadBalancerPool != nil {
		d.Set("load_balancer_pool", *instanceGroup.LoadBalancerPool.ID)
	}
	d.Set("subnets", subnets)
	managers := make([]string, 0)

	for i := 0; i < len(instanceGroup.Managers); i++ {
		managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
	}
	d.Set("managers", managers)

	d.Set("status", *instanceGroup.Status)
	d.Set("vpc", *instanceGroup.VPC.ID)
	tags, err := GetTagsUsingCRN(meta, *instanceGroup.CRN)
	if err != nil {
		log.Printf(
			"Error on get of instance group (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)
	return nil
}

func resourceIBMISInstanceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	instanceGroup, _, _ := sess.GetInstanceGroup(&getInstanceGroupOptions)

	if *instanceGroup.Status == UNHEALTHY {
		intf, healthError := waitForHealthyInstanceGroupAfterDetachingInstances(d, meta, d.Timeout(schema.TimeoutDelete))
		ig := intf.(*vpcv1.InstanceGroup)
		if *ig.Status == UNHEALTHY || healthError != nil {
			return fmt.Errorf("Error deleting instance group, group %s", *ig.Status)
		}
	}
	// Inorder to delete instance group, need to update membership count to 0

	if *instanceGroup.MembershipCount > 0 {
		zeroMembers := int64(0)
		instanceGroupUpdateOptions := vpcv1.UpdateInstanceGroupOptions{}
		instanceGroupPatchModel := vpcv1.InstanceGroupPatch{}

		instanceGroupPatchModel.MembershipCount = &zeroMembers
		instanceGroupPatch, err := instanceGroupPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("Error calling asPatch for IGPatch: %s", err)
		}

		instanceGroupUpdateOptions.ID = &instanceGroupID
		instanceGroupUpdateOptions.InstanceGroupPatch = instanceGroupPatch
		_, response, err := sess.UpdateInstanceGroup(&instanceGroupUpdateOptions)
		if err != nil {
			return fmt.Errorf("Error updating instanceGroup's instance count to 0 : %s\n%s", err, response)
		}
		intf, healthError := waitForHealthyInstanceGroupAfterDetachingInstances(d, meta, d.Timeout(schema.TimeoutDelete))
		ig := intf.(*vpcv1.InstanceGroup)

		if *ig.Status == UNHEALTHY || healthError != nil {
			return fmt.Errorf("Error deleting instance group, group %s", *ig.Status)
		}

	}
	if *instanceGroup.Status == HEALTHY {
		deleteInstanceGroupOptions := vpcv1.DeleteInstanceGroupOptions{ID: &instanceGroupID}
		response, Err := sess.DeleteInstanceGroup(&deleteInstanceGroupOptions)
		if Err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error Deleting the InstanceGroup: %s\n%s", Err, response)
		}

		_, deleteError := waitForInstanceGroupDelete(d, meta)
		if deleteError != nil {
			return deleteError
		}
	} else {
		return fmt.Errorf("Error Deleting the InstanceGroup: group in %s state", *instanceGroup.Status)
	}
	return nil
}

func resourceIBMISInstanceGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	_, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
	}
	return true, nil
}

func waitForHealthyInstanceGroup(d *schema.ResourceData, meta interface{}, timeout time.Duration) (interface{}, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return nil, err
	}

	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}

	healthStateConf := &resource.StateChangeConf{
		Pending: []string{SCALING},
		Target:  []string{HEALTHY},
		Refresh: func() (interface{}, string, error) {
			instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
			if err != nil || instanceGroup == nil {
				return nil, SCALING, fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
			}
			log.Println("Status : ", *instanceGroup.Status)

			if *instanceGroup.Status == "" {
				return instanceGroup, SCALING, nil
			}
			return instanceGroup, *instanceGroup.Status, nil
		},
		Timeout:      timeout,
		Delay:        20 * time.Second,
		MinTimeout:   5 * time.Second,
		PollInterval: 10 * time.Second,
	}

	return healthStateConf.WaitForState()

}

func waitForHealthyInstanceGroupAfterDetachingInstances(d *schema.ResourceData, meta interface{}, timeout time.Duration) (interface{}, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return nil, err
	}

	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}

	healthStateConf := &resource.StateChangeConf{
		Pending: []string{SCALING, UNHEALTHY},
		Target:  []string{HEALTHY},
		Refresh: func() (interface{}, string, error) {
			instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
			if err != nil || instanceGroup == nil {
				return nil, SCALING, fmt.Errorf("Error Getting InstanceGroup: %s\n%s", err, response)
			}
			log.Println("Status : ", *instanceGroup.Status)

			if *instanceGroup.Status == "" {
				return instanceGroup, SCALING, nil
			}
			return instanceGroup, *instanceGroup.Status, nil
		},
		Timeout:      timeout,
		Delay:        20 * time.Second,
		MinTimeout:   5 * time.Second,
		PollInterval: 10 * time.Second,
	}

	return healthStateConf.WaitForState()

}

func waitForInstanceGroupDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	healthStateConf := &resource.StateChangeConf{
		Pending: []string{HEALTHY},
		Target:  []string{DELETING},
		Refresh: func() (interface{}, string, error) {
			resp, err := resourceIBMISInstanceGroupExists(d, meta)
			if resp {
				return resp, HEALTHY, nil
			}
			return resp, DELETING, err
		},
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        20 * time.Second,
		MinTimeout:   5 * time.Second,
		PollInterval: 10 * time.Second,
	}

	return healthStateConf.WaitForState()

}
