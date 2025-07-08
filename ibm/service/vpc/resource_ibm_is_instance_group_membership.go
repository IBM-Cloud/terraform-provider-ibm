// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceGroupMembership                                = "instance_group_membership"
	isInstanceGroup                                          = "instance_group"
	isInstanceGroupMembershipName                            = "name"
	isInstanceGroupMemershipActionDelete                     = "action_delete"
	isInstanceGroupMemershipDeleteInstanceOnMembershipDelete = "delete_instance_on_membership_delete"
	isInstanceGroupMemershipInstance                         = "instance"
	isInstanceGroupMemershipInstanceName                     = "name"
	isInstanceGroupMemershipInstanceTemplate                 = "instance_template"
	isInstanceGroupMemershipInstanceTemplateName             = "name"
	isInstanceGroupMembershipCrn                             = "crn"
	isInstanceGroupMembershipVirtualServerInstance           = "virtual_server_instance"
	isInstanceGroupMembershipLoadBalancerPoolMember          = "load_balancer_pool_member"
	isInstanceGroupMembershipStatus                          = "status"
)

func ResourceIBMISInstanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceGroupMembershipUpdate,
		ReadContext:   resourceIBMISInstanceGroupMembershipRead,
		UpdateContext: resourceIBMISInstanceGroupMembershipUpdate,
		DeleteContext: resourceIBMISInstanceGroupMembershipDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isInstanceGroup: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_membership", isInstanceGroup),
				Description:  "The instance group identifier.",
			},
			isInstanceGroupMembership: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_membership", isInstanceGroupMembership),
				Description:  "The unique identifier for this instance group membership.",
			},
			isInstanceGroupMembershipName: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group_membership", isInstanceGroupMembershipName),
				Description:  "The user-defined name for this instance group membership. Names must be unique within the instance group.",
			},
			isInstanceGroupMemershipActionDelete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "The delete flag for this instance group membership. Must be set to true to delete instance group membership.",
			},
			isInstanceGroupMemershipDeleteInstanceOnMembershipDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If set to true, when deleting the membership the instance will also be deleted.",
			},
			isInstanceGroupMemershipInstance: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGroupMembershipCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this virtual server instance.",
						},
						isInstanceGroupMembershipVirtualServerInstance: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						isInstanceGroupMemershipInstanceName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
					},
				},
			},
			isInstanceGroupMemershipInstanceTemplate: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceGroupMembershipCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this instance template.",
						},
						isInstanceGroupMemershipInstanceTemplate: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance template.",
						},
						isInstanceGroupMemershipInstanceTemplateName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this instance template.",
						},
					},
				},
			},
			isInstanceGroupMembershipLoadBalancerPoolMember: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this load balancer pool member.",
			},
			isInstanceGroupMembershipStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the instance group membership- `deleting`: Membership is deleting dependent resources- `failed`: Membership was unable to maintain dependent resources- `healthy`: Membership is active and serving in the group- `pending`: Membership is waiting for dependent resources- `unhealthy`: Membership has unhealthy dependent resources.",
			},
		},
	}
}

func ResourceIBMISInstanceGroupMembershipValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceGroupMembershipName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceGroup,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceGroupMembership,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})
	ibmISInstanceGroupMembershipResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_group_membership", Schema: validateSchema}
	return &ibmISInstanceGroupMembershipResourceValidator
}

func resourceIBMISInstanceGroupMembershipUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupID := d.Get(isInstanceGroup).(string)
	instanceGroupMembershipID := d.Get(isInstanceGroupMembership).(string)

	getInstanceGroupMembershipOptions := vpcv1.GetInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}

	instanceGroupMembership, _, err := sess.GetInstanceGroupMembershipWithContext(context, &getInstanceGroupMembershipOptions)
	if err != nil || instanceGroupMembership == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupMembershipWithContext failed: %s", err.Error()), "ibm_is_instance_group_membership", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(fmt.Sprintf("%s/%s", instanceGroupID, instanceGroupMembershipID))

	if v, ok := d.GetOk(isInstanceGroupMemershipActionDelete); ok {
		actionDelete := v.(bool)
		if actionDelete {
			return resourceIBMISInstanceGroupMembershipDelete(context, d, meta)
		}
	}

	if v, ok := d.GetOk(isInstanceGroupMembershipName); ok {
		name := v.(string)
		if name != *instanceGroupMembership.Name {

			updateInstanceGroupMembershipOptions := vpcv1.UpdateInstanceGroupMembershipOptions{}
			instanceGroupMembershipPatchModel := &vpcv1.InstanceGroupMembershipPatch{}
			instanceGroupMembershipPatchModel.Name = &name

			updateInstanceGroupMembershipOptions.ID = &instanceGroupMembershipID
			updateInstanceGroupMembershipOptions.InstanceGroupID = &instanceGroupID
			instanceGroupMembershipPatch, err := instanceGroupMembershipPatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupMembershipPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_group_membership", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			updateInstanceGroupMembershipOptions.InstanceGroupMembershipPatch = instanceGroupMembershipPatch
			_, _, err = sess.UpdateInstanceGroupMembershipWithContext(context, &updateInstanceGroupMembershipOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupMembershipWithContext failed: %s", err.Error()), "ibm_is_instance_group_membership", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}
	return resourceIBMISInstanceGroupMembershipRead(context, d, meta)
}

func resourceIBMISInstanceGroupMembershipRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupMembershipID := parts[1]

	getInstanceGroupMembershipOptions := vpcv1.GetInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}
	instanceGroupMembership, response, err := sess.GetInstanceGroupMembershipWithContext(context, &getInstanceGroupMembershipOptions)
	if err != nil || instanceGroupMembership == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupMembershipWithContext failed: %s", err.Error()), "ibm_is_instance_group_membership", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isInstanceGroupMemershipDeleteInstanceOnMembershipDelete, *instanceGroupMembership.DeleteInstanceOnMembershipDelete); err != nil {
		err = fmt.Errorf("Error setting delete_instance_on_membership_delete: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-delete_instance_on_membership_delete").GetDiag()
	}
	if d.Set(isInstanceGroupMembership, *instanceGroupMembership.ID); err != nil {
		err = fmt.Errorf("Error setting instance_group_membership: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-instance_group_membership").GetDiag()
	}
	if err = d.Set(isInstanceGroupMembershipStatus, *instanceGroupMembership.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-status").GetDiag()
	}

	instances := make([]map[string]interface{}, 0)
	if instanceGroupMembership.Instance != nil {
		instance := map[string]interface{}{
			isInstanceGroupMembershipCrn:                   *instanceGroupMembership.Instance.CRN,
			isInstanceGroupMembershipVirtualServerInstance: *instanceGroupMembership.Instance.ID,
			isInstanceGroupMemershipInstanceName:           *instanceGroupMembership.Instance.Name,
		}
		instances = append(instances, instance)
	}

	if err = d.Set(isInstanceGroupMemershipInstance, instances); err != nil {
		err = fmt.Errorf("Error setting instance: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-instance").GetDiag()
	}
	instance_templates := make([]map[string]interface{}, 0)
	if instanceGroupMembership.InstanceTemplate != nil {
		instance_template := map[string]interface{}{
			isInstanceGroupMembershipCrn:                 *instanceGroupMembership.InstanceTemplate.CRN,
			isInstanceGroupMemershipInstanceTemplate:     *instanceGroupMembership.InstanceTemplate.ID,
			isInstanceGroupMemershipInstanceTemplateName: *instanceGroupMembership.InstanceTemplate.Name,
		}
		instance_templates = append(instance_templates, instance_template)
	}

	if err = d.Set(isInstanceGroupMemershipInstanceTemplate, instance_templates); err != nil {
		err = fmt.Errorf("Error setting instance_template: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-instance_template").GetDiag()
	}
	if !core.IsNil(instanceGroupMembership.PoolMember) {
		if err = d.Set(isInstanceGroupMembershipLoadBalancerPoolMember, *instanceGroupMembership.PoolMember.ID); err != nil {
			err = fmt.Errorf("Error setting load_balancer_pool_member: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "read", "set-load_balancer_pool_member").GetDiag()
		}
	}
	return nil
}

func resourceIBMISInstanceGroupMembershipDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group_membership", "delete", "sep-id-parts").GetDiag()
	}
	instanceGroupID := parts[0]
	instanceGroupMembershipID := parts[1]

	deleteInstanceGroupMembershipOptions := vpcv1.DeleteInstanceGroupMembershipOptions{
		ID:              &instanceGroupMembershipID,
		InstanceGroupID: &instanceGroupID,
	}
	response, err := sess.DeleteInstanceGroupMembershipWithContext(context, &deleteInstanceGroupMembershipOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceGroupMembershipWithContext failed: %s", err.Error()), "ibm_is_instance_group_membership", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}
