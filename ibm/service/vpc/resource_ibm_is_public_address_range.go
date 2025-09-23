// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
)

const (
	isPublicAddressRangeDeleting   = "deleting"
	isPublicAddressRangeDeleted    = "deleted"
	isPublicAddressRangeAvailable  = "stable"
	isPublicAddressRangeFailed     = "failed"
	isPublicAddressRangePending    = "pending"
	isPublicAddressRangeSuspended  = "suspended"
	isPublicAddressRangeUpdating   = "updating"
	isPublicAddressRangeWaiting    = "waiting"
	isPublicAddressRangeUserTags   = "tags"
	isPublicAddressRangeAccessTags = "access_tags"
)

func ResourceIBMPublicAddressRange() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPublicAddressRangeCreate,
		ReadContext:   resourceIBMPublicAddressRangeRead,
		UpdateContext: resourceIBMPublicAddressRangeUpdate,
		DeleteContext: resourceIBMPublicAddressRangeDelete,
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
				}),
		),

		Schema: map[string]*schema.Schema{
			"ipv4_address_count": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of IPv4 addresses in this public address range.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_public_address_range", "name"),
				Description:  "The name for this public address range. The name is unique across all public address ranges in the region.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The resource group for this public address range.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource group.",
						},
					},
				},
			},
			"target": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The target this public address range is bound to.If absent, this pubic address range is not bound to a target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The VPC this public address range is bound to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The CRN for this VPC.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The URL for this VPC.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this VPC.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this VPC. The name is unique across all VPCs in the region.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"zone": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The zone this public address range resides in.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The URL for this zone.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The globally unique name for this zone.",
									},
								},
							},
						},
					},
				},
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The public IPv4 range, expressed in CIDR format.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the public address range was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this public address range.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this public address range.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the public address range.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			isPublicAddressRangeUserTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_public_address_range", isPublicAddressRangeUserTags)},
				Set:         flex.ResourceIBMVPCHash,
				Description: "User Tags for the PublicAddressRange",
			},
			isPublicAddressRangeAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_public_address_range", isPublicAddressRangeAccessTags)},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func ResourceIBMPublicAddressRangeValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 isPublicAddressRangeUserTags,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 isPublicAddressRangeAccessTags,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_public_address_range", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPublicAddressRangeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_public_address_range", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createPublicAddressRangeOptions := &vpcv1.CreatePublicAddressRangeOptions{}

	createPublicAddressRangeOptions.SetIpv4AddressCount(int64(d.Get("ipv4_address_count").(int)))
	if _, ok := d.GetOk("name"); ok {
		createPublicAddressRangeOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		resourceGroupModel, err := ResourceIBMPublicAddressRangeMapToResourceGroupIdentity(d.Get("resource_group.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "create", "parse-resource_group").GetDiag()
		}
		createPublicAddressRangeOptions.SetResourceGroup(resourceGroupModel)
	}
	if _, ok := d.GetOk("target"); ok {
		targetModel, err := ResourceIBMPublicAddressRangeMapToPublicAddressRangeTargetPrototype(d.Get("target.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "create", "parse-target").GetDiag()
		}
		createPublicAddressRangeOptions.SetTarget(targetModel)
	}

	publicAddressRange, _, err := vpcClient.CreatePublicAddressRangeWithContext(context, createPublicAddressRangeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*publicAddressRange.ID)
	log.Printf("[INFO] PublicAddressRange : %s", *publicAddressRange.ID)

	_, err = isWaitForPublicAddressRangeAvailable(vpcClient, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPublicAddressRangeAvailable failed: %s", err.Error()), "ibm_public_address_range", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isPublicAddressRangeUserTags); ok || v != "" {
		oldList, newList := d.GetChange(isPublicAddressRangeUserTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicAddressRange.CRN, "", "user")
		if err != nil {
			log.Printf(
				"Error on create of resource public address range (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isPublicAddressRangeAccessTags); ok {
		oldList, newList := d.GetChange(isPublicAddressRangeAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *publicAddressRange.CRN, "", "access")
		if err != nil {
			log.Printf(
				"Error on create of resource public address range (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMPublicAddressRangeRead(context, d, meta)
}

func isPublicAddressRangeRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{
			ID: &id,
		}
		PublicAddressRange, response, err := sess.GetPublicAddressRange(getPublicAddressRangeOptions)
		if err != nil {
			return nil, isPublicAddressRangeFailed, fmt.Errorf("[ERROR] Error getting PublicAddressRange : %s\n%s", err, response)
		}

		if *PublicAddressRange.LifecycleState == isPublicAddressRangeAvailable {
			return PublicAddressRange, *PublicAddressRange.LifecycleState, nil
		} else if *PublicAddressRange.LifecycleState == isPublicAddressRangeFailed {
			return PublicAddressRange, *PublicAddressRange.LifecycleState, fmt.Errorf("PublicAddressRange (%s) went into failed state during the operation \n [WARNING] Running terraform apply again will remove the tainted PublicAddressRange and attempt to create the PublicAddressRange again replacing the previous configuration", *PublicAddressRange.ID)
		}

		return PublicAddressRange, isPublicAddressRangePending, nil
	}
}

func isWaitForPublicAddressRangeAvailable(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for public address range (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isPublicAddressRangePending},
		Target:     []string{isPublicAddressRangeAvailable, isPublicAddressRangeFailed},
		Refresh:    isPublicAddressRangeRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMPublicAddressRangeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{}

	getPublicAddressRangeOptions.SetID(d.Id())

	publicAddressRange, response, err := vpcClient.GetPublicAddressRangeWithContext(context, getPublicAddressRangeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("ipv4_address_count", flex.IntValue(publicAddressRange.Ipv4AddressCount)); err != nil {
		err = fmt.Errorf("Error setting ipv4_address_count: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-ipv4_address_count").GetDiag()
	}
	if !core.IsNil(publicAddressRange.Name) {
		if err = d.Set("name", publicAddressRange.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(publicAddressRange.ResourceGroup) {
		resourceGroupMap, err := ResourceIBMPublicAddressRangeResourceGroupReferenceToMap(publicAddressRange.ResourceGroup)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "resource_group-to-map").GetDiag()
		}
		if err = d.Set("resource_group", []map[string]interface{}{resourceGroupMap}); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-resource_group").GetDiag()
		}
	}
	if !core.IsNil(publicAddressRange.Target) {
		targetMap, err := ResourceIBMPublicAddressRangePublicAddressRangeTargetToMap(publicAddressRange.Target)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "target-to-map").GetDiag()
		}
		if err = d.Set("target", []map[string]interface{}{targetMap}); err != nil {
			err = fmt.Errorf("Error setting target: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-target").GetDiag()
		}
	}
	if err = d.Set("cidr", publicAddressRange.CIDR); err != nil {
		err = fmt.Errorf("Error setting cidr: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-cidr").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(publicAddressRange.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("crn", publicAddressRange.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-crn").GetDiag()
	}
	if err = d.Set("href", publicAddressRange.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-href").GetDiag()
	}
	if err = d.Set("lifecycle_state", publicAddressRange.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", publicAddressRange.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "read", "set-resource_type").GetDiag()
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *publicAddressRange.CRN, "", "user")
	if err != nil {
		log.Printf(
			"Error on get of resource public address range (%s) tags: %s", d.Id(), err)
	}
	d.Set(isPublicAddressRangeUserTags, tags)

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *publicAddressRange.CRN, "", "access")
	if err != nil {
		log.Printf(
			"Error on get of resource public address range (%s) access tags: %s", d.Id(), err)
	}
	d.Set(isPublicAddressRangeAccessTags, accesstags)

	return nil
}

func resourceIBMPublicAddressRangeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if d.HasChange(isPublicAddressRangeUserTags) {
		oldList, newList := d.GetChange(isPublicAddressRangeUserTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get("crn").(string), "", "user")
		if err != nil {
			log.Printf(
				"Error on update of resource public address range (%s) tags: %s", d.Id(), err)
		}
	}

	if d.HasChange(isPublicAddressRangeAccessTags) {
		oldList, newList := d.GetChange(isPublicAddressRangeAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get("crn").(string), "", "access")
		if err != nil {
			log.Printf(
				"Error on update of resource public address range (%s) access tags: %s", d.Id(), err)
		}
	}

	updatePublicAddressRangeOptions := &vpcv1.UpdatePublicAddressRangeOptions{}

	updatePublicAddressRangeOptions.SetID(d.Id())

	patchVals := &vpcv1.PublicAddressRangePatch{}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		updatePublicAddressRangeOptions.PublicAddressRangePatch, _ = patchVals.AsPatch()
		_, _, err = vpcClient.UpdatePublicAddressRangeWithContext(context, updatePublicAddressRangeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	if d.HasChange("target") {
		targetRemoved := false
		if _, ok := d.GetOk("target"); !ok {
			targetRemoved = true
		}
		target, err := ResourceIBMPublicAddressRangeMapToPublicAddressRangeTargetPatch(d.Get("target.0").(map[string]interface{}), d)
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Target = target
		updatePublicAddressRangeOptions.PublicAddressRangePatch, _ = patchVals.AsPatch()
		if targetRemoved {
			updatePublicAddressRangeOptions.PublicAddressRangePatch["target"] = nil
		}
		_, _, err = vpcClient.UpdatePublicAddressRangeWithContext(context, updatePublicAddressRangeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForPublicAddressRangeUpdate(vpcClient, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMPublicAddressRangeRead(context, d, meta)
}

func isWaitForPublicAddressRangeUpdate(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for PublicAddressRange (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isPublicAddressRangeUpdating},
		Target:     []string{isPublicAddressRangeAvailable, isPublicAddressRangeFailed},
		Refresh:    isPublicAddressRangeUpdateRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isPublicAddressRangeUpdateRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{
			ID: &id,
		}
		publicAddressRange, response, err := sess.GetPublicAddressRange(getPublicAddressRangeOptions)
		if err != nil {
			return nil, isPublicAddressRangeFailed, fmt.Errorf("[ERROR] Error getting PublicAddressRange : %s\n%s", err, response)
		}

		if *publicAddressRange.LifecycleState == isPublicAddressRangeAvailable || *publicAddressRange.LifecycleState == isPublicAddressRangeFailed {
			return publicAddressRange, *publicAddressRange.LifecycleState, nil
		} else if *publicAddressRange.LifecycleState == isPublicAddressRangeFailed {
			return publicAddressRange, *publicAddressRange.LifecycleState, fmt.Errorf("PublicAddressRange (%s) went into failed state during the operation \n [WARNING] Running terraform apply again will remove the tainted PublicAddressRange and attempt to create the PublicAddressRange again replacing the previous configuration", *publicAddressRange.ID)
		}

		return publicAddressRange, isPublicAddressRangeUpdating, nil
	}
}

func resourceIBMPublicAddressRangeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_public_address_range", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePublicAddressRangeOptions := &vpcv1.DeletePublicAddressRangeOptions{}

	deletePublicAddressRangeOptions.SetID(d.Id())

	_, _, err = vpcClient.DeletePublicAddressRangeWithContext(context, deletePublicAddressRangeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForPublicAddressRangeDeleted(vpcClient, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePublicAddressRangeWithContext failed: %s", err.Error()), "ibm_public_address_range", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isWaitForPublicAddressRangeDeleted(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for PublicAddressRange (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isPublicAddressRangeDeleting},
		Target:     []string{isPublicAddressRangeDeleted, isPublicAddressRangeFailed},
		Refresh:    isPublicAddressRangeDeleteRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPublicAddressRangeDeleteRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Refresh function for PublicAddressRange delete.")
		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{
			ID: &id,
		}
		publicAddressRange, response, err := sess.GetPublicAddressRange(getPublicAddressRangeOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return publicAddressRange, isPublicAddressRangeDeleted, nil
			}
			return nil, isPublicAddressRangeFailed, fmt.Errorf("[ERROR] The PublicAddressRange %s failed to delete: %s\n%s", id, err, response)
		}
		return publicAddressRange, *publicAddressRange.LifecycleState, nil
	}
}

func ResourceIBMPublicAddressRangeMapToResourceGroupIdentity(modelMap map[string]interface{}) (vpcv1.ResourceGroupIdentityIntf, error) {
	model := &vpcv1.ResourceGroupIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToPublicAddressRangeTargetPrototype(modelMap map[string]interface{}) (*vpcv1.PublicAddressRangeTargetPrototype, error) {
	model := &vpcv1.PublicAddressRangeTargetPrototype{}
	VPCModel, err := ResourceIBMPublicAddressRangeMapToVPCIdentity(modelMap["vpc"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.VPC = VPCModel
	ZoneModel, err := ResourceIBMPublicAddressRangeMapToZoneIdentity(modelMap["zone"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Zone = ZoneModel
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToVPCIdentity(modelMap map[string]interface{}) (vpcv1.VPCIdentityIntf, error) {
	model := &vpcv1.VPCIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["crn"] != nil && modelMap["crn"].(string) != "" {
		model.CRN = core.StringPtr(modelMap["crn"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToVPCIdentityPatch(modelMap map[string]interface{}, d *schema.ResourceData) (vpcv1.VPCIdentityIntf, error) {
	model := &vpcv1.VPCIdentity{}

	if d.HasChange("target.0.vpc.0.id") && modelMap["id"] != nil && modelMap["id"].(string) != "" {
		log.Println("d.HasChange('target.0.vpc.0.id')")
		log.Println(d.HasChange("target.0.vpc.0.id"))
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if d.HasChange("target.0.vpc.0.crn") && modelMap["crn"] != nil && modelMap["crn"].(string) != "" {
		log.Println("d.HasChange('target.0.vpc.0.crn')")
		log.Println(d.HasChange("target.0.vpc.0.crn"))
		model.CRN = core.StringPtr(modelMap["crn"].(string))
	}
	if d.HasChange("target.0.vpc.0.href") && modelMap["href"] != nil && modelMap["href"].(string) != "" {
		log.Println("d.HasChange('target.0.vpc.0.href')")
		log.Println(d.HasChange("target.0.vpc.0.href"))
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToZoneIdentity(modelMap map[string]interface{}) (vpcv1.ZoneIdentityIntf, error) {
	model := &vpcv1.ZoneIdentity{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToZoneIdentityPatch(modelMap map[string]interface{}, d *schema.ResourceData) (vpcv1.ZoneIdentityIntf, error) {
	model := &vpcv1.ZoneIdentity{}
	if d.HasChange("target.0.zone.0.name") && modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if d.HasChange("target.0.zone.0.href") && modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeMapToPublicAddressRangeTargetPatch(modelMap map[string]interface{}, d *schema.ResourceData) (*vpcv1.PublicAddressRangeTargetPatch, error) {
	model := &vpcv1.PublicAddressRangeTargetPatch{}
	if d.HasChange("target.0.vpc") && modelMap["vpc"] != nil && len(modelMap["vpc"].([]interface{})) > 0 {
		VPCModel, err := ResourceIBMPublicAddressRangeMapToVPCIdentityPatch(modelMap["vpc"].([]interface{})[0].(map[string]interface{}), d)
		if err != nil {
			return model, err
		}
		model.VPC = VPCModel
	}
	if d.HasChange("target.0.zone") && modelMap["zone"] != nil && len(modelMap["zone"].([]interface{})) > 0 {
		ZoneModel, err := ResourceIBMPublicAddressRangeMapToZoneIdentityPatch(modelMap["zone"].([]interface{})[0].(map[string]interface{}), d)
		if err != nil {
			return model, err
		}
		model.Zone = ZoneModel
	}
	return model, nil
}

func ResourceIBMPublicAddressRangeResourceGroupReferenceToMap(model *vpcv1.ResourceGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMPublicAddressRangePublicAddressRangeTargetToMap(model *vpcv1.PublicAddressRangeTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	vpcMap, err := ResourceIBMPublicAddressRangeVPCReferenceToMap(model.VPC)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc"] = []map[string]interface{}{vpcMap}
	zoneMap, err := ResourceIBMPublicAddressRangeZoneReferenceToMap(model.Zone)
	if err != nil {
		return modelMap, err
	}
	modelMap["zone"] = []map[string]interface{}{zoneMap}
	return modelMap, nil
}

func ResourceIBMPublicAddressRangeVPCReferenceToMap(model *vpcv1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMPublicAddressRangeDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func ResourceIBMPublicAddressRangeDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMPublicAddressRangeZoneReferenceToMap(model *vpcv1.ZoneReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	return modelMap, nil
}
