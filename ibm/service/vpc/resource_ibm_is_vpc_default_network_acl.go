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

func ResourceIBMISVPCDefaultNetworkACL() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPCDefaultNetworkACLCreate,
		ReadContext:   resourceIBMISVPCDefaultNetworkACLRead,
		UpdateContext: resourceIBMISVPCDefaultNetworkACLUpdate,
		DeleteContext: resourceIBMISVPCDefaultNetworkACLDelete,
		Exists:        resourceIBMISVPCDefaultNetworkACLExists,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// For import, the ID should be in format vpc_id/network_acl_id
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
			isVPCDefaultNetworkACL: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC ID",
			},
			isNetworkACLName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_network_acl", isNetworkACLName),
				Description:  "Network ACL name",
			},
			isNetworkACLTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_network_acl", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},
			isNetworkACLAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_network_acl", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
			isNetworkACLCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the resource",
			},
			isNetworkACLResourceGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this network ACL",
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
			"subnets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The subnets to which this network ACL is attached",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this subnet",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this subnet",
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

func ResourceIBMISVPCDefaultNetworkACLValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isNetworkACLName,
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

	ibmISVPCDefaultNetworkACLResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpc_default_network_acl", Schema: validateSchema}
	return &ibmISVPCDefaultNetworkACLResourceValidator
}

func resourceIBMISVPCDefaultNetworkACLCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := d.Get(isVPCDefaultNetworkACL).(string)

	// Get the VPC to obtain the default network ACL
	getVpcOptions := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}

	vpc, _, err := sess.GetVPCWithContext(context, getVpcOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCWithContext failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if vpc.DefaultNetworkACL == nil {
		tfErr := flex.TerraformErrorf(fmt.Errorf("VPC does not have a default network ACL"), "VPC does not have a default network ACL", "ibm_is_vpc_default_network_acl", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	defaultNetworkACLID := *vpc.DefaultNetworkACL.ID
	d.SetId(fmt.Sprintf("%s/%s", vpcID, defaultNetworkACLID))

	// Handle tags on creation
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isNetworkACLTags); ok || v != "" {
		oldList, newList := d.GetChange(isNetworkACLTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpc.DefaultNetworkACL.CRN, "", isUserTagType)
		if err != nil {
			log.Printf("Error on create of resource default network acl (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isNetworkACLAccessTags); ok {
		oldList, newList := d.GetChange(isNetworkACLAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vpc.DefaultNetworkACL.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf("Error on create of resource default network acl (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISVPCDefaultNetworkACLRead(context, d, meta)
}

func resourceIBMISVPCDefaultNetworkACLRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Parse ID to get VPC ID and Network ACL ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("invalid ID format"), "ID should be in format vpc_id/network_acl_id", "ibm_is_vpc_default_network_acl", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := parts[0]
	networkACLID := parts[1]

	getNetworkAclOptions := &vpcv1.GetNetworkACLOptions{
		ID: &networkACLID,
	}
	nwacl, response, err := sess.GetNetworkACLWithContext(context, getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(isVPCDefaultNetworkACL, vpcID)
	d.Set("default_network_acl", networkACLID)

	if !core.IsNil(nwacl.Name) {
		if err = d.Set(isNetworkACLName, *nwacl.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-name").GetDiag()
		}
	}

	if err = d.Set(isNetworkACLCRN, *nwacl.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-crn").GetDiag()
	}

	// Handle resource group
	resourceGroupList := []map[string]interface{}{}
	if nwacl.ResourceGroup != nil {
		resourceGroupMap := map[string]interface{}{
			"href": *nwacl.ResourceGroup.Href,
			"id":   *nwacl.ResourceGroup.ID,
			"name": *nwacl.ResourceGroup.Name,
		}
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		if err = d.Set(flex.ResourceGroupName, *nwacl.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-resource_group_name").GetDiag()
		}
	}
	if err = d.Set(isNetworkACLResourceGroup, resourceGroupList); err != nil {
		err = fmt.Errorf("Error setting resource_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-resource_group").GetDiag()
	}

	// Handle subnets
	subnets := make([]map[string]interface{}, 0)
	for _, subnet := range nwacl.Subnets {
		subnetMap := map[string]interface{}{
			"id":   *subnet.ID,
			"name": *subnet.Name,
		}
		subnets = append(subnets, subnetMap)
	}
	if err = d.Set("subnets", subnets); err != nil {
		err = fmt.Errorf("Error setting subnets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-subnets").GetDiag()
	}

	// Handle tags
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *nwacl.CRN, "", isUserTagType)
	if err != nil {
		log.Printf("Error on get of resource default network acl (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isNetworkACLTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-tags").GetDiag()
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *nwacl.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf("Error on get of resource default network acl (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isNetworkACLAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/acl"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *nwacl.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *nwacl.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "read", "set-resource_crn").GetDiag()
	}

	return nil
}

func resourceIBMISVPCDefaultNetworkACLUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Parse ID to get Network ACL ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("invalid ID format"), "ID should be in format vpc_id/network_acl_id", "ibm_is_vpc_default_network_acl", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	networkACLID := parts[1]
	hasChanged := false

	// Handle name change
	if d.HasChange(isNetworkACLName) {
		name := d.Get(isNetworkACLName).(string)
		updateNetworkACLOptions := &vpcv1.UpdateNetworkACLOptions{
			ID: &networkACLID,
		}
		networkACLPatchModel := &vpcv1.NetworkACLPatch{
			Name: &name,
		}
		networkACLPatch, err := networkACLPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("networkACLPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateNetworkACLOptions.NetworkACLPatch = networkACLPatch
		_, _, err = sess.UpdateNetworkACLWithContext(context, updateNetworkACLOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateNetworkACLWithContext failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		hasChanged = true
	}

	// Handle tags update
	if d.HasChange(isNetworkACLTags) {
		oldList, newList := d.GetChange(isNetworkACLTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isNetworkACLCRN).(string), "", isUserTagType)
		if err != nil {
			log.Printf("Error on update of resource default network acl (%s) tags: %s", d.Id(), err)
		}
		hasChanged = true
	}

	if d.HasChange(isNetworkACLAccessTags) {
		oldList, newList := d.GetChange(isNetworkACLAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isNetworkACLCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf("Error on update of resource default network acl (%s) access tags: %s", d.Id(), err)
		}
		hasChanged = true
	}

	if hasChanged {
		return resourceIBMISVPCDefaultNetworkACLRead(context, d, meta)
	}

	return nil
}

func resourceIBMISVPCDefaultNetworkACLDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Default network ACL cannot be deleted, just remove from Terraform state
	d.SetId("")
	return nil
}

func resourceIBMISVPCDefaultNetworkACLExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_default_network_acl", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	// Parse ID to get Network ACL ID
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of vpcID/networkACLID", d.Id())
	}

	networkACLID := parts[1]
	getNetworkAclOptions := &vpcv1.GetNetworkACLOptions{
		ID: &networkACLID,
	}
	_, response, err := sess.GetNetworkACL(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkACL failed: %s", err.Error()), "ibm_is_vpc_default_network_acl", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
