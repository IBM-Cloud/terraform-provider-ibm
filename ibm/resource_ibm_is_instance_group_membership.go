// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMISInstanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceGroupMembershipUpdate,
		Read:     resourceIBMISInstanceGroupMembershipRead,
		Update:   resourceIBMISInstanceGroupMembershipUpdate,
		Delete:   resourceIBMISInstanceGroupMembershipDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"instance_group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_membership", "instance_group"),
				Description:  "The instance group identifier.",
			},
			"instance_group_membership": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_membership", "instance_group_membership"),
				Description:  "The unique identifier for this instance group membership.",
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_is_instance_group_membership", "name"),
				Description:  "The user-defined name for this instance group membership. Names must be unique within the instance group.",
			},
			"action_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "The delete flag for this instance group membership. Must be set to true to delete instance group membership.",
			},
			"delete_instance_on_membership_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, when deleting the membership the instance will also be deleted.",
			},
			"instance": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
						},
						"virtual_server_instance": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
					},
				},
			},
			"instance_template": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this instance template.",
						},
						"instance_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance template.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this instance template.",
						},
					},
				},
			},
			"load_balancer_pool_member": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this load balancer pool member.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the instance group membership- `deleting`: Membership is deleting dependent resources- `failed`: Membership was unable to maintain dependent resources- `healthy`: Membership is active and serving in the group- `pending`: Membership is waiting for dependent resources- `unhealthy`: Membership has unhealthy dependent resources.",
			},
		},
	}
}

func resourceIBMISInstanceGroupMembershipValidator() *ResourceValidator {

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
			Identifier:                 "instance_group",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "instance_group_membership",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})
	ibmISInstanceGroupMembershipResourceValidator := ResourceValidator{ResourceName: "ibm_is_instance_group_membership", Schema: validateSchema}
	return &ibmISInstanceGroupMembershipResourceValidator
}

func resourceIBMISInstanceGroupMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	instanceGroupID := d.Get("instance_group").(string)
	instanceGroupMembershipID := d.Get("instance_group_membership").(string)

	getInstanceGroupMembershipOptions := vpcv1.GetInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}

	instanceGroupMembership, response, err := sess.GetInstanceGroupMembership(&getInstanceGroupMembershipOptions)
	if err != nil || instanceGroupMembership == nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Membership: %s\n%s", err, response)
	}

	if v, ok := d.GetOk("action_delete"); ok {
		actionDelete := v.(bool)
		if actionDelete {
			return resourceIBMISInstanceGroupMembershipDelete(d, meta)
		}
	}

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		if name != *instanceGroupMembership.Name {
			d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, instanceGroupMembershipID))

			updateInstanceGroupMembershipOptions := vpcv1.UpdateInstanceGroupMembershipOptions{}
			instanceGroupMembershipPatchModel := &vpcv1.InstanceGroupMembershipPatch{}
			instanceGroupMembershipPatchModel.Name = &name

			updateInstanceGroupMembershipOptions.ID = &instanceGroupMembershipID
			updateInstanceGroupMembershipOptions.InstanceGroupID = &instanceGroupID
			instanceGroupMembershipPatch, err := instanceGroupMembershipPatchModel.AsPatch()
			if err != nil {
				return fmt.Errorf("Error calling asPatch for InstanceGroupMembershipPatch: %s", err)
			}
			updateInstanceGroupMembershipOptions.InstanceGroupMembershipPatch = instanceGroupMembershipPatch
			_, response, err := sess.UpdateInstanceGroupMembership(&updateInstanceGroupMembershipOptions)
			if err != nil {
				return fmt.Errorf("Error updating InstanceGroup Membership: %s\n%s", err, response)
			}
		}
	}
	return resourceIBMISInstanceGroupMembershipRead(d, meta)
}

func resourceIBMISInstanceGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	instanceGroupID := parts[0]
	instanceGroupMembershipID := parts[1]

	getInstanceGroupMembershipOptions := vpcv1.GetInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupMembership, response, err := sess.GetInstanceGroupMembership(&getInstanceGroupMembershipOptions)
	if err != nil || instanceGroupMembership == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting InstanceGroup Membership: %s\n%s", err, response)
	}
	d.Set("delete_instance_on_membership_delete", *instanceGroupMembership.DeleteInstanceOnMembershipDelete)
	d.Set("instance_group_membership", *instanceGroupMembership.ID)
	d.Set("status", *instanceGroupMembership.Status)

	instances := make([]map[string]interface{}, 0)
	if instanceGroupMembership.Instance != nil {
		instance := map[string]interface{}{
			"crn":                     *instanceGroupMembership.Instance.CRN,
			"virtual_server_instance": *instanceGroupMembership.Instance.ID,
			"name":                    *instanceGroupMembership.Instance.Name,
		}
		instances = append(instances, instance)
	}
	d.Set("instance", instances)

	instance_templates := make([]map[string]interface{}, 0)
	if instanceGroupMembership.InstanceTemplate != nil {
		instance_template := map[string]interface{}{
			"crn":               *instanceGroupMembership.InstanceTemplate.CRN,
			"instance_template": *instanceGroupMembership.InstanceTemplate.ID,
			"name":              *instanceGroupMembership.InstanceTemplate.Name,
		}
		instance_templates = append(instance_templates, instance_template)
	}
	d.Set("instance_template", instance_templates)

	if instanceGroupMembership.PoolMember != nil && instanceGroupMembership.PoolMember.ID != nil {
		d.Set("load_balancer_pool_member", *instanceGroupMembership.PoolMember.ID)
	}
	return nil
}

func resourceIBMISInstanceGroupMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	instanceGroupID := parts[0]
	instanceGroupMembershipID := parts[1]

	deleteInstanceGroupMembershipOptions := vpcv1.DeleteInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}
	response, err := sess.DeleteInstanceGroupMembership(&deleteInstanceGroupMembershipOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Deleting the InstanceGroup Membership: %s\n%s", err, response)
	}
	return nil
}
