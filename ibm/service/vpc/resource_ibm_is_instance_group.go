// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
	DELETING                     = "deleting"
	isInstanceGroupAccessTags    = "access_tags"
	isInstanceGroupUserTagType   = "user"
	isInstanceGroupAccessTagType = "access"
)

func ResourceIBMISInstanceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceGroupCreate,
		ReadContext:   resourceIBMISInstanceGroupRead,
		UpdateContext: resourceIBMISInstanceGroupUpdate,
		DeleteContext: resourceIBMISInstanceGroupDelete,
		Exists:        resourceIBMISInstanceGroupExists,
		Importer:      &schema.ResourceImporter{},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				},
			),
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
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group", "name"),
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
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group", "instance_count"),
				Description:  "The number of instances in the instance group",
			},

			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of this instance group",
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
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_group", "application_port"),
				RequiredWith: []string{"load_balancer", "load_balancer_pool"},
				Description:  "Used by the instance group when scaling up instances to supply the port for the load balancer pool member.",
			},

			"load_balancer": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"application_port", "load_balancer_pool"},
				Description:  "load balancer ID",
			},

			"load_balancer_pool": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"application_port", "load_balancer"},
				Description:  "load balancer pool ID",
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
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_group", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags for instance group",
			},

			isInstanceGroupAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_group", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func ResourceIBMISInstanceGroupValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
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
			Identifier:                 "instance_count",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "1000"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "application_port",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "1",
			MaxValue:                   "65535"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISInstanceGroupResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_group", Schema: validateSchema}
	return &ibmISInstanceGroupResourceValidator
}

func resourceIBMISInstanceGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get("name").(string)
	instanceTemplate := d.Get("instance_template").(string)

	subnets := d.Get("subnets")

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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

	instanceGroup, _, err := sess.CreateInstanceGroupWithContext(context, &instanceGroupOptions)
	if err != nil || instanceGroup == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*instanceGroup.ID)

	_, healthError := waitForHealthyInstanceGroup(d.Id(), meta, d.Timeout(schema.TimeoutCreate))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", healthError.Error()), "ibm_is_instance_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *instanceGroup.CRN, "", isInstanceGroupUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of instance group (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isInstanceGroupAccessTags); ok {
		oldList, newList := d.GetChange(isInstanceGroupAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *instanceGroup.CRN, "", isInstanceGroupAccessTagType)
		if err != nil {
			log.Printf(
				"[ERROR] Error on create of instance group (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISInstanceGroupRead(context, d, meta)

}

func resourceIBMISInstanceGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var changed bool
	instanceGroupUpdateOptions := vpcv1.UpdateInstanceGroupOptions{}
	instanceGroupPatchModel := vpcv1.InstanceGroupPatch{}

	if d.HasChange("tags") {
		instanceGroupID := d.Id()
		getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
		instanceGroup, _, err := sess.GetInstanceGroupWithContext(context, &getInstanceGroupOptions)
		if err != nil || instanceGroup == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateTagsUsingCRN(oldList, newList, meta, *instanceGroup.CRN)
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		instanceGroupUpdateOptions.InstanceGroupPatch = instanceGroupPatch
		_, _, err = sess.UpdateInstanceGroupWithContext(context, &instanceGroupUpdateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		// wait for instance group health update with update timeout configured.
		_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
		if healthError != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", err.Error()), "ibm_is_instance_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISInstanceGroupRead(context, d, meta)
}

func resourceIBMISInstanceGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupID := d.Id()
	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	instanceGroup, response, err := sess.GetInstanceGroupWithContext(context, &getInstanceGroupOptions)
	if err != nil || instanceGroup == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(instanceGroup.Name) {
		if err = d.Set("name", instanceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(instanceGroup.InstanceTemplate) {
		if err = d.Set("instance_template", *instanceGroup.InstanceTemplate.ID); err != nil {
			err = fmt.Errorf("Error setting instance_template: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-instance_template").GetDiag()
		}
	}
	if err = d.Set("instances", *instanceGroup.MembershipCount); err != nil {
		err = fmt.Errorf("Error setting instances: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-instances").GetDiag()
	}
	if err = d.Set("instance_count", *instanceGroup.MembershipCount); err != nil {
		err = fmt.Errorf("Error setting instance_count: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-instance_count").GetDiag()
	}
	if !core.IsNil(instanceGroup.ResourceGroup) {
		if err = d.Set("resource_group", *instanceGroup.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-resource_group").GetDiag()
		}
	}
	if !core.IsNil(instanceGroup.ApplicationPort) {
		if err = d.Set("application_port", flex.IntValue(instanceGroup.ApplicationPort)); err != nil {
			err = fmt.Errorf("Error setting application_port: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-application_port").GetDiag()
		}
	}
	subnets := make([]string, 0)

	for i := 0; i < len(instanceGroup.Subnets); i++ {
		subnets = append(subnets, string(*(instanceGroup.Subnets[i].ID)))
	}
	if instanceGroup.LoadBalancerPool != nil {
		if err = d.Set("load_balancer_pool", *instanceGroup.LoadBalancerPool.ID); err != nil {
			err = fmt.Errorf("Error setting load_balancer_pool: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-load_balancer_pool").GetDiag()
		}
	}
	if err = d.Set("subnets", subnets); err != nil {
		err = fmt.Errorf("Error setting subnets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-subnets").GetDiag()
	}

	managers := make([]string, 0)

	for i := 0; i < len(instanceGroup.Managers); i++ {
		managers = append(managers, string(*(instanceGroup.Managers[i].ID)))
	}
	if err = d.Set("managers", managers); err != nil {
		err = fmt.Errorf("Error setting managers: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-managers").GetDiag()
	}
	if err = d.Set("status", *instanceGroup.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-status").GetDiag()
	}
	if err = d.Set("vpc", *instanceGroup.VPC.ID); err != nil {
		err = fmt.Errorf("Error setting vpc: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-vpc").GetDiag()
	}
	if err = d.Set("crn", *instanceGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-crn").GetDiag()
	}

	tags, err := flex.GetTagsUsingCRN(meta, *instanceGroup.CRN)
	if err != nil {
		log.Printf(
			"Error on get of instance group (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set("tags", tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "read", "set-tags").GetDiag()
	}
	return nil
}

func getLBStatus(context context.Context, sess *vpcv1.VpcV1, lbId string) (string, error) {
	getlboptions := &vpcv1.GetLoadBalancerOptions{
		ID: &lbId,
	}
	lb, response, err := sess.GetLoadBalancerWithContext(context, getlboptions)
	if err != nil || lb == nil {
		return "", fmt.Errorf("[ERROR] Error Getting Load Balancer : %s\n%s", err, response)
	}
	return *lb.ProvisioningStatus, nil
}

func resourceIBMISInstanceGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceGroupID := d.Id()

	// Before we delete the instance group, we need to
	// know if the load balancer attached is in active state

	// First, get the instance
	igOpts := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}
	instanceGroup, response, err := sess.GetInstanceGroupWithContext(context, &igOpts)
	if err != nil || instanceGroup == nil {
		if response != nil && response.StatusCode == 404 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Instance Group with id:[%s] not found!!", instanceGroupID), "ibm_is_instance_group", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Internal Error fetching info for instance group [%s]", instanceGroupID), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Inorder to delete instance group, need to update membership count to 0
	zeroMembers := int64(0)
	instanceGroupUpdateOptions := vpcv1.UpdateInstanceGroupOptions{}
	instanceGroupPatchModel := vpcv1.InstanceGroupPatch{}

	instanceGroupPatchModel.MembershipCount = &zeroMembers
	instanceGroupPatch, err := instanceGroupPatchModel.AsPatch()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("instanceGroupPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceGroupUpdateOptions.ID = &instanceGroupID
	instanceGroupUpdateOptions.InstanceGroupPatch = instanceGroupPatch
	_, response, err = sess.UpdateInstanceGroupWithContext(context, &instanceGroupUpdateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, healthError := waitForHealthyInstanceGroup(instanceGroupID, meta, d.Timeout(schema.TimeoutUpdate))
	if healthError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHealthyInstanceGroup failed: %s", err.Error()), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// If there is any load balancer, please check if it is active
	if instanceGroup.LoadBalancerPool != nil {
		loadBalancerPool := *instanceGroup.LoadBalancerPool.Href
		// The sixth component is the Load Balancer ID
		loadBalancerID := strings.Split(loadBalancerPool, "/")[5]

		// Now check if the load balancer is in active state or not
		lbStatus, err := getLBStatus(context, sess, loadBalancerID)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("getLBStatus failed: %s", err.Error()), "ibm_is_instance_group", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if lbStatus != "active" {
			log.Printf("Load Balancer [%s] is not active....Waiting it to be active!\n", loadBalancerID)
			_, err := isWaitForLBAvailable(sess, loadBalancerID, d.Timeout(schema.TimeoutDelete))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBAvailable failed: %s", err.Error()), "ibm_is_instance_group", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			lbStatus, err = getLBStatus(context, sess, loadBalancerID)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("getLBStatus failed: %s", err.Error()), "ibm_is_instance_group", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if lbStatus != "active" {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("LoadBalancer [%s] is not active yet! Current Load Balancer status is [%s]", loadBalancerID, lbStatus), "ibm_is_instance_group", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}

	deleteInstanceGroupOptions := vpcv1.DeleteInstanceGroupOptions{ID: &instanceGroupID}
	response, Err := sess.DeleteInstanceGroupWithContext(context, &deleteInstanceGroupOptions)
	if Err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceGroupWithContext failed: %s", err.Error()), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, deleteError := waitForInstanceGroupDelete(d, meta)
	if deleteError != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForInstanceGroupDelete failed: %s", err.Error()), "ibm_is_instance_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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
		return false, fmt.Errorf("[ERROR] Error Getting InstanceGroup: %s\n%s", err, response)
	}
	return true, nil
}

func waitForHealthyInstanceGroup(instanceGroupID string, meta interface{}, timeout time.Duration) (interface{}, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return nil, err
	}

	getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &instanceGroupID}

	healthStateConf := &resource.StateChangeConf{
		Pending: []string{SCALING},
		Target:  []string{HEALTHY},
		Refresh: func() (interface{}, string, error) {
			instanceGroup, response, err := sess.GetInstanceGroup(&getInstanceGroupOptions)
			if err != nil || instanceGroup == nil {
				return nil, SCALING, fmt.Errorf("[ERROR] Error Getting InstanceGroup: %s\n%s", err, response)
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
