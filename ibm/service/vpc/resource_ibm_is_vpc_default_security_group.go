// Copyright IBM Corp. 2017, 2025 All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISVPCDefaultSecurityGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPCDefaultSecurityGroupCreate,
		ReadContext:   resourceIBMISVPCDefaultSecurityGroupRead,
		UpdateContext: resourceIBMISVPCDefaultSecurityGroupUpdate,
		DeleteContext: resourceIBMISVPCDefaultSecurityGroupDelete,
		Exists:        resourceIBMISVPCDefaultSecurityGroupExists,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// For import, the ID should be in format vpc_id/security_group_id
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isVPCDefaultSecurityGroup: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC ID",
			},
			isSecurityGroupName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_security_group", isSecurityGroupName),
				Description:  "Security group name",
			},
			isSecurityGroupTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_security_group", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},
			isSecurityGroupAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_security_group", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
			isSecurityGroupCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the resource",
			},
			isSecurityGroupResourceGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this security group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group",
						},
					},
				},
			},
			"targets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The targets attached to this security group",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this target",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this target",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of the target",
						},
					},
				},
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},
			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the resource",
			},
			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func ResourceIBMISVPCDefaultSecurityGroupValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
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

	ibmISVPCDefaultSecurityGroupResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpc_default_security_group", Schema: validateSchema}
	return &ibmISVPCDefaultSecurityGroupResourceValidator
}

func resourceIBMISVPCDefaultSecurityGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := d.Get(isVPCDefaultSecurityGroup).(string)

	// Get the VPC to obtain the default security group
	getVpcOptions := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}

	vpc, _, err := sess.GetVPCWithContext(context, getVpcOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCWithContext failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if vpc.DefaultSecurityGroup == nil {
		tfErr := flex.TerraformErrorf(fmt.Errorf("VPC does not have a default security group"), "VPC does not have a default security group", "ibm_is_vpc_default_security_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	defaultSecurityGroupID := *vpc.DefaultSecurityGroup.ID
	d.SetId(fmt.Sprintf("%s/%s", vpcID, defaultSecurityGroupID))

	// Handle tags on creation
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isSecurityGroupTags); ok || v != "" {
		oldList, newList := d.GetChange(isSecurityGroupTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpc.DefaultSecurityGroup.CRN, "", isUserTagType)
		if err != nil {
			log.Printf("Error on create of resource default security group (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isSecurityGroupAccessTags); ok {
		oldList, newList := d.GetChange(isSecurityGroupAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpc.DefaultSecurityGroup.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf("Error on create of resource default security group (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISVPCDefaultSecurityGroupRead(context, d, meta)
}

func resourceIBMISVPCDefaultSecurityGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Parse ID to get VPC ID and Security Group ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("invalid ID format"), "ID should be in format vpc_id/security_group_id", "ibm_is_vpc_default_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := parts[0]
	securityGroupID := parts[1]

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &securityGroupID,
	}
	securityGroup, response, err := sess.GetSecurityGroupWithContext(context, getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(isVPCDefaultSecurityGroup, vpcID)
	d.Set("default_security_group", securityGroupID)

	if !core.IsNil(securityGroup.Name) {
		if err = d.Set(isSecurityGroupName, *securityGroup.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-name").GetDiag()
		}
	}

	if err = d.Set(isSecurityGroupCRN, *securityGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-crn").GetDiag()
	}

	// Handle resource group
	resourceGroupList := []map[string]interface{}{}
	if securityGroup.ResourceGroup != nil {
		resourceGroupMap := map[string]interface{}{
			"href": *securityGroup.ResourceGroup.Href,
			"id":   *securityGroup.ResourceGroup.ID,
			"name": *securityGroup.ResourceGroup.Name,
		}
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		if err = d.Set(flex.ResourceGroupName, *securityGroup.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-resource_group_name").GetDiag()
		}
	}
	if err = d.Set(isSecurityGroupResourceGroup, resourceGroupList); err != nil {
		err = fmt.Errorf("Error setting resource_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-resource_group").GetDiag()
	}

	// Handle targets
	listSecurityGroupTargetsOptions := &vpcv1.ListSecurityGroupTargetsOptions{
		SecurityGroupID: &securityGroupID,
	}
	sgTargets, response, err := sess.ListSecurityGroupTargetsWithContext(context, listSecurityGroupTargetsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupTargetsWithContext failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	targets := make([]map[string]interface{}, 0)
	if sgTargets != nil && len(sgTargets.Targets) > 0 {
		for _, target := range sgTargets.Targets {
			if target != nil {
				targetMap := map[string]interface{}{}
				switch targetType := target.(type) {
				case *vpcv1.SecurityGroupTargetReferenceNetworkInterfaceReferenceTargetContext:
					targetMap["id"] = *targetType.ID
					if targetType.Name != nil {
						targetMap["name"] = *targetType.Name
					}
					if targetType.ResourceType != nil {
						targetMap["resource_type"] = *targetType.ResourceType
					}
				case *vpcv1.SecurityGroupTargetReferenceLoadBalancerReference:
					targetMap["id"] = *targetType.ID
					if targetType.Name != nil {
						targetMap["name"] = *targetType.Name
					}
					if targetType.ResourceType != nil {
						targetMap["resource_type"] = *targetType.ResourceType
					}
				case *vpcv1.SecurityGroupTargetReferenceVPNServerReference:
					targetMap["id"] = *targetType.ID
					if targetType.Name != nil {
						targetMap["name"] = *targetType.Name
					}
					if targetType.ResourceType != nil {
						targetMap["resource_type"] = *targetType.ResourceType
					}
				}
				if len(targetMap) > 0 {
					targets = append(targets, targetMap)
				}
			}
		}
	}
	if err = d.Set("targets", targets); err != nil {
		err = fmt.Errorf("Error setting targets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-targets").GetDiag()
	}

	// Handle tags
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isUserTagType)
	if err != nil {
		log.Printf("Error on get of resource default security group (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isSecurityGroupTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-tags").GetDiag()
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf("Error on get of resource default security group (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isSecurityGroupAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/securityGroups"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *securityGroup.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *securityGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "read", "set-resource_crn").GetDiag()
	}

	return nil
}

func resourceIBMISVPCDefaultSecurityGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Parse ID to get Security Group ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("invalid ID format"), "ID should be in format vpc_id/security_group_id", "ibm_is_vpc_default_security_group", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	securityGroupID := parts[1]
	hasChanged := false

	// Handle name change
	if d.HasChange(isSecurityGroupName) {
		name := d.Get(isSecurityGroupName).(string)
		updateSecurityGroupOptions := &vpcv1.UpdateSecurityGroupOptions{
			ID: &securityGroupID,
		}
		securityGroupPatchModel := &vpcv1.SecurityGroupPatch{
			Name: &name,
		}
		securityGroupPatch, err := securityGroupPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("securityGroupPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateSecurityGroupOptions.SecurityGroupPatch = securityGroupPatch
		_, _, err = sess.UpdateSecurityGroupWithContext(context, updateSecurityGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		hasChanged = true
	}

	// Handle tags update
	if d.HasChange(isSecurityGroupTags) {
		oldList, newList := d.GetChange(isSecurityGroupTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSecurityGroupCRN).(string), "", isUserTagType)
		if err != nil {
			log.Printf("Error on update of resource default security group (%s) tags: %s", d.Id(), err)
		}
		hasChanged = true
	}

	if d.HasChange(isSecurityGroupAccessTags) {
		oldList, newList := d.GetChange(isSecurityGroupAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSecurityGroupCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf("Error on update of resource default security group (%s) access tags: %s", d.Id(), err)
		}
		hasChanged = true
	}

	if hasChanged {
		return resourceIBMISVPCDefaultSecurityGroupRead(context, d, meta)
	}

	return nil
}

func resourceIBMISVPCDefaultSecurityGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Default security group cannot be deleted, just remove from Terraform state
	d.SetId("")
	return nil
}

func resourceIBMISVPCDefaultSecurityGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_security_group", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	// Parse ID to get Security Group ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of vpcID/securityGroupID", d.Id())
	}

	securityGroupID := parts[1]
	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &securityGroupID,
	}
	_, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroup failed: %s", err.Error()), "ibm_is_vpc_default_security_group", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
