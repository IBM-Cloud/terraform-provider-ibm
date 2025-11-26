// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isFloatingIPAddress       = "address"
	isFloatingIPCRN           = "crn"
	isFloatingIPName          = "name"
	isFloatingIPStatus        = "status"
	isFloatingIPZone          = "zone"
	isFloatingIPTarget        = "target"
	isFloatingIPResourceGroup = "resource_group"
	isFloatingIPTags          = "tags"

	isFloatingIPPending   = "pending"
	isFloatingIPAvailable = "available"
	isFloatingIPDeleting  = "deleting"
	isFloatingIPDeleted   = "done"

	isFloatingIPAccessTags = "access_tags"
)

func ResourceIBMISFloatingIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISFloatingIPCreate,
		ReadContext:   resourceIBMISFloatingIPRead,
		UpdateContext: resourceIBMISFloatingIPUpdate,
		DeleteContext: resourceIBMISFloatingIPDelete,
		Exists:        resourceIBMISFloatingIPExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
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
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {

					if diff.HasChange(isFloatingIPTarget) {
						if !diff.NewValueKnown(isFloatingIPTarget) {
							diff.ForceNew(isFloatingIPTarget)
							return nil
						}
						old, new := diff.GetChange(isFloatingIPTarget)
						if old != "" || new != "" {
							sess, err := vpcClient(v)
							if err != nil {
								return err
							}
							if checkIfZoneChanged(old.(string), new.(string), diff.Get(isFloatingIPZone).(string), sess) {
								diff.ForceNew(isFloatingIPTarget)
							}
						}
					}
					return nil
				},
			),
		),

		Schema: map[string]*schema.Schema{
			isFloatingIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP address",
			},

			isFloatingIPName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_floating_ip", isFloatingIPName),
				Description:  "Name of the floating IP",
			},

			isFloatingIPStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP status",
			},

			isFloatingIPZone: {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPTarget},
				ExactlyOneOf:  []string{isFloatingIPTarget, isFloatingIPZone},
				Description:   "Zone name",
			},

			isFloatingIPTarget: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isFloatingIPZone},
				ExactlyOneOf:  []string{isFloatingIPTarget, isFloatingIPZone},
				Description:   "Target info",
			},
			floatingIPTargets: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target of this floating IP.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						floatingIPTargetsDeleted: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									floatingIPTargetsMoreInfo: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						floatingIPTargetsHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network interface.",
						},
						floatingIPTargetsId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this network interface.",
						},
						floatingIPTargetsName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this network interface.",
						},
						floatingIpPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									floatingIpPrimaryIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									floatingIpPrimaryIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									floatingIpPrimaryIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									floatingIpPrimaryIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
									floatingIpPrimaryIpResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type",
									},
								},
							},
						},
						floatingIPTargetsResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						floatingIPTargetsCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this public gateway.",
						},
					},
				},
			},

			isFloatingIPResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group info",
			},

			isFloatingIPTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_floating_ip", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Floating IP tags",
			},

			isFloatingIPAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_floating_ip", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
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
				Description: "The crn of the resource",
			},

			isFloatingIPCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func vpcClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	return sess, err
}

func ResourceIBMISFloatingIPValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isFloatingIPName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
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

	ibmISFloatingIPResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_floating_ip", Schema: validateSchema}
	return &ibmISFloatingIPResourceValidator
}

func resourceIBMISFloatingIPCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get(isFloatingIPName).(string)
	err := fipCreate(context, d, meta, name)
	if err != nil {
		return err
	}

	return resourceIBMISFloatingIPRead(context, d, meta)
}

func fipCreate(context context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	floatingIPPrototype := &vpcv1.FloatingIPPrototype{
		Name: &name,
	}
	zone, target := "", ""
	if zn, ok := d.GetOk(isFloatingIPZone); ok {
		zone = zn.(string)
		floatingIPPrototype.Zone = &vpcv1.ZoneIdentity{
			Name: &zone,
		}
	}

	if tgt, ok := d.GetOk(isFloatingIPTarget); ok {
		target = tgt.(string)
		floatingIPPrototype.Target = &vpcv1.FloatingIPTargetPrototypeNetworkInterfaceIdentity{
			ID: &target,
		}
	}

	if zone == "" && target == "" {
		err = fmt.Errorf("%s or %s need to be provided", isFloatingIPZone, isFloatingIPTarget)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "create", "validation").GetDiag()
	}

	if rgrp, ok := d.GetOk(isFloatingIPResourceGroup); ok {
		rg := rgrp.(string)
		floatingIPPrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	createFloatingIPOptions := &vpcv1.CreateFloatingIPOptions{
		FloatingIPPrototype: floatingIPPrototype,
	}

	floatingip, response, err := vpcClient.CreateFloatingIPWithContext(context, createFloatingIPOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*floatingip.ID)
	log.Printf("[INFO] Floating IP : %s[%s]", *floatingip.ID, *floatingip.Address)
	_, err = isWaitForInstanceFloatingIP(vpcClient, d.Id(), d)
	if err != nil {
		return diag.FromErr(err)
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isFloatingIPTags); ok || v != "" {
		oldList, newList := d.GetChange(isFloatingIPTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *floatingip.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc Floating IP (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isFloatingIPAccessTags); ok {
		oldList, newList := d.GetChange(isFloatingIPAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *floatingip.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource Floating IP (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISFloatingIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := fipGet(context, d, meta, id)
	if err != nil {
		return err
	}

	return nil
}

func fipGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getFloatingIPOptions := &vpcv1.GetFloatingIPOptions{
		ID: &id,
	}
	floatingip, response, err := vpcClient.GetFloatingIPWithContext(context, getFloatingIPOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()

	}
	if err = d.Set(isFloatingIPName, *floatingip.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-name").GetDiag()
	}
	if err = d.Set(isFloatingIPAddress, *floatingip.Address); err != nil {
		err = fmt.Errorf("Error setting address: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-address").GetDiag()
	}
	if err = d.Set(isFloatingIPStatus, *floatingip.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-status").GetDiag()
	}
	if err = d.Set(isFloatingIPZone, *floatingip.Zone.Name); err != nil {
		err = fmt.Errorf("Error setting zone: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-zone").GetDiag()
	}
	if err = d.Set(isFloatingIPCRN, *floatingip.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-crn").GetDiag()
	}

	targetId, targetMap := floatingIPCollectionFloatingIpTargetToMap(floatingip.Target)
	if targetId != "" {
		if err = d.Set(isFloatingIPTarget, targetId); err != nil {
			err = fmt.Errorf("Error setting target: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-target").GetDiag()
		}
	} else {
		if err = d.Set(isFloatingIPTarget, ""); err != nil {
			err = fmt.Errorf("Error setting target: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-target").GetDiag()
		}
	}
	targetList := make([]map[string]interface{}, 0)
	targetList = append(targetList, targetMap)
	if err = d.Set(floatingIPTargets, targetList); err != nil {
		err = fmt.Errorf("Error setting targets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-targets").GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *floatingip.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of vpc Floating IP (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isFloatingIPTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-tags").GetDiag()
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *floatingip.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource Floating IP (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isFloatingIPAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/floatingIPs"); err != nil {
		err = fmt.Errorf("Error setting controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *floatingip.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-flex_name").GetDiag()

	}
	if err = d.Set(flex.ResourceCRN, *floatingip.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-flex_crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *floatingip.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-flex_status").GetDiag()
	}
	if floatingip.ResourceGroup != nil {
		if err = d.Set(flex.ResourceGroupName, floatingip.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-flex_resource_group_name").GetDiag()
		}
		if err = d.Set(isFloatingIPResourceGroup, floatingip.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "set-resource_group").GetDiag()
		}
	}
	return nil
}

func resourceIBMISFloatingIPUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := fipUpdate(context, d, meta, id)
	if err != nil {
		return err
	}
	return resourceIBMISFloatingIPRead(context, d, meta)
}

func fipUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if d.HasChange(isFloatingIPTags) {
		options := &vpcv1.GetFloatingIPOptions{
			ID: &id,
		}
		fip, response, err := vpcClient.GetFloatingIPWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		oldList, newList := d.GetChange(isFloatingIPTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *fip.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of vpc Floating IP (%s) tags: %s", id, err)
		}
	}
	if d.HasChange(isFloatingIPAccessTags) {
		options := &vpcv1.GetFloatingIPOptions{
			ID: &id,
		}
		fip, response, err := vpcClient.GetFloatingIPWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		oldList, newList := d.GetChange(isFloatingIPAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *fip.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource Floating IP (%s) access tags: %s", d.Id(), err)
		}
	}

	hasChanged := false
	options := &vpcv1.UpdateFloatingIPOptions{
		ID: &id,
	}
	floatingIPPatchModel := &vpcv1.FloatingIPPatch{}
	if d.HasChange(isFloatingIPName) {
		name := d.Get(isFloatingIPName).(string)
		floatingIPPatchModel.Name = &name
		hasChanged = true
		floatingIPPatch, err := floatingIPPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for FloatingIPPatch: %s", err), "ibm_is_floating_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.FloatingIPPatch = floatingIPPatch
	}

	if d.HasChange(isFloatingIPTarget) {
		target := d.Get(isFloatingIPTarget).(string)
		floatingIPPatchModel.Target = &vpcv1.FloatingIPTargetPatch{
			ID: &target,
		}
		hasChanged = true
		floatingIPPatch, err := floatingIPPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error calling asPatch for FloatingIPPatch: %s", err), "ibm_is_floating_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.FloatingIPPatch = floatingIPPatch
	}
	if hasChanged {
		_, response, err := vpcClient.UpdateFloatingIPWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISFloatingIPDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := fipDelete(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func fipDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getFloatingIpOptions := &vpcv1.GetFloatingIPOptions{
		ID: &id,
	}
	_, response, err := vpcClient.GetFloatingIPWithContext(context, getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.DeleteFloatingIPOptions{
		ID: &id,
	}
	response, err = vpcClient.DeleteFloatingIPWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteFloatingIPWithContext failed: %s\n%s", err, response), "ibm_is_floating_ip", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForFloatingIPDeleted(vpcClient, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func resourceIBMISFloatingIPExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	exists, err := fipExists(d, meta, id)
	return exists, err
}

func fipExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		return false, err
	}
	getFloatingIpOptions := &vpcv1.GetFloatingIPOptions{
		ID: &id,
	}
	_, response, err := vpcClient.GetFloatingIP(getFloatingIpOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting floating IP: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForFloatingIPDeleted(fip *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for FloatingIP (%s) to be deleted.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{isFloatingIPPending, isFloatingIPDeleting},
		Target:     []string{"", isFloatingIPDeleted},
		Refresh:    isFloatingIPDeleteRefreshFunc(fip, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isFloatingIPDeleteRefreshFunc(fip *vpcv1.VpcV1, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] floating ip delete function here")
		getfipoptions := &vpcv1.GetFloatingIPOptions{
			ID: &id,
		}
		FloatingIP, response, err := fip.GetFloatingIP(getfipoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return FloatingIP, isFloatingIPDeleted, nil
			}
			return FloatingIP, "", fmt.Errorf("[ERROR] Error Getting Floating IP: %s\n%s", err, response)
		}
		return FloatingIP, isFloatingIPDeleting, err
	}
}

func isWaitForInstanceFloatingIP(floatingipC *vpcv1.VpcV1, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for floating IP (%s) to be available.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{isFloatingIPPending},
		Target:     []string{isFloatingIPAvailable, ""},
		Refresh:    isInstanceFloatingIPRefreshFunc(floatingipC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceFloatingIPRefreshFunc(floatingipC *vpcv1.VpcV1, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getfipoptions := &vpcv1.GetFloatingIPOptions{
			ID: &id,
		}
		instance, response, err := floatingipC.GetFloatingIP(getfipoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Floating IP for the instance: %s\n%s", err, response)
		}

		if *instance.Status == "available" {
			return instance, isFloatingIPAvailable, nil
		}

		return instance, isFloatingIPPending, nil
	}
}

func checkIfZoneChanged(oldNic, newNic, currentZone string, floatingipC *vpcv1.VpcV1) bool {
	var oldZone, newZone string
	listInstancesOptions := &vpcv1.ListInstancesOptions{}
	start := ""
	allrecs := []vpcv1.Instance{}
	for {

		if start != "" {
			listInstancesOptions.Start = &start
		}

		instances, _, err := floatingipC.ListInstances(listInstancesOptions)
		if err != nil {
			return false
		}
		start = flex.GetNext(instances.Next)
		allrecs = append(allrecs, instances.Instances...)
		if start == "" {
			break
		}
	}
	for _, instance := range allrecs {
		for _, nic := range instance.NetworkInterfaces {
			if oldNic == *nic.ID {
				oldZone = *instance.Zone.Name
			}
			if newNic == *nic.ID {
				newZone = *instance.Zone.Name
			}
		}
	}
	if newZone != oldZone {
		if oldZone == "" && newZone == currentZone {
			return false
		}
		return true
	}
	return false
}

func floatingIPCollectionFloatingIpTargetToMap(targetItemIntf vpcv1.FloatingIPTargetIntf) (targetId string, targetMap map[string]interface{}) {
	targetMap = map[string]interface{}{}
	targetId = ""
	if targetItemIntf != nil {
		switch reflect.TypeOf(targetItemIntf).String() {
		case "*vpcv1.FloatingIPTargetNetworkInterfaceReference":
			{
				targetItem := targetItemIntf.(*vpcv1.FloatingIPTargetNetworkInterfaceReference)
				targetId = *targetItem.ID
				if targetItem.Deleted != nil {
					deletedList := []map[string]interface{}{}
					deletedMap := floatingIPTargetNicDeletedToMap(*targetItem.Deleted)
					deletedList = append(deletedList, deletedMap)
					targetMap[floatingIPTargetsDeleted] = deletedList
				}
				if targetItem.Href != nil {
					targetMap[floatingIPTargetsHref] = targetItem.Href
				}
				if targetItem.ID != nil {
					targetMap[floatingIPTargetsId] = targetItem.ID
				}
				if targetItem.Name != nil {
					targetMap[floatingIPTargetsName] = targetItem.Name
				}
				if targetItem.PrimaryIP != nil {
					primaryIpList := make([]map[string]interface{}, 0)
					currentIP := map[string]interface{}{}
					if targetItem.PrimaryIP.Address != nil {
						currentIP[floatingIpPrimaryIpAddress] = *targetItem.PrimaryIP.Address
					}
					if targetItem.PrimaryIP.Href != nil {
						currentIP[floatingIpPrimaryIpHref] = *targetItem.PrimaryIP.Href
					}
					if targetItem.PrimaryIP.Name != nil {
						currentIP[floatingIpPrimaryIpName] = *targetItem.PrimaryIP.Name
					}
					if targetItem.PrimaryIP.ID != nil {
						currentIP[floatingIpPrimaryIpId] = *targetItem.PrimaryIP.ID
					}
					if targetItem.PrimaryIP.ResourceType != nil {
						currentIP[floatingIpPrimaryIpResourceType] = *targetItem.PrimaryIP.ResourceType
					}
					primaryIpList = append(primaryIpList, currentIP)
					targetMap[floatingIpPrimaryIP] = primaryIpList
				}
				if targetItem.ResourceType != nil {
					targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
				}
			}
		case "*vpcv1.FloatingIPTargetPublicGatewayReference":
			{
				targetItem := targetItemIntf.(*vpcv1.FloatingIPTargetPublicGatewayReference)
				targetId = *targetItem.ID
				if targetItem.Deleted != nil {
					deletedList := []map[string]interface{}{}
					deletedMap := floatingIPTargetPgDeletedToMap(*targetItem.Deleted)
					deletedList = append(deletedList, deletedMap)
					targetMap[floatingIPTargetsDeleted] = deletedList
				}
				if targetItem.Href != nil {
					targetMap[floatingIPTargetsHref] = targetItem.Href
				}
				if targetItem.ID != nil {
					targetMap[floatingIPTargetsId] = targetItem.ID
				}
				if targetItem.Name != nil {
					targetMap[floatingIPTargetsName] = targetItem.Name
				}
				if targetItem.ResourceType != nil {
					targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
				}
				if targetItem.CRN != nil {
					targetMap[floatingIPTargetsCrn] = targetItem.CRN
				}
			}
		case "*vpcv1.FloatingIPTarget":
			{
				targetItem := targetItemIntf.(*vpcv1.FloatingIPTarget)
				targetId = *targetItem.ID
				if targetItem.Deleted != nil {
					deletedList := []map[string]interface{}{}
					deletedMap := floatingIPTargetNicDeletedToMap(*targetItem.Deleted)
					deletedList = append(deletedList, deletedMap)
					targetMap[floatingIPTargetsDeleted] = deletedList
				}
				if targetItem.Href != nil {
					targetMap[floatingIPTargetsHref] = targetItem.Href
				}
				if targetItem.ID != nil {
					targetMap[floatingIPTargetsId] = targetItem.ID
				}
				if targetItem.Name != nil {
					targetMap[floatingIPTargetsName] = targetItem.Name
				}
				if targetItem.PrimaryIP != nil && targetItem.PrimaryIP.Address != nil {
					primaryIpList := make([]map[string]interface{}, 0)
					currentIP := map[string]interface{}{}
					if targetItem.PrimaryIP.Address != nil {
						currentIP[floatingIpPrimaryIpAddress] = *targetItem.PrimaryIP.Address
					}
					if targetItem.PrimaryIP.Href != nil {
						currentIP[floatingIpPrimaryIpHref] = *targetItem.PrimaryIP.Href
					}
					if targetItem.PrimaryIP.Name != nil {
						currentIP[floatingIpPrimaryIpName] = *targetItem.PrimaryIP.Name
					}
					if targetItem.PrimaryIP.ID != nil {
						currentIP[floatingIpPrimaryIpId] = *targetItem.PrimaryIP.ID
					}
					if targetItem.PrimaryIP.ResourceType != nil {
						currentIP[floatingIpPrimaryIpResourceType] = *targetItem.PrimaryIP.ResourceType
					}
					primaryIpList = append(primaryIpList, currentIP)
					targetMap[floatingIpPrimaryIP] = primaryIpList
				}
				if targetItem.ResourceType != nil {
					targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
				}
				if targetItem.CRN != nil {
					targetMap[floatingIPTargetsCrn] = targetItem.CRN
				}
			}
		}
	}

	return targetId, targetMap
}

func floatingIPTargetNicDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap[floatingIPTargetsMoreInfo] = deletedItem.MoreInfo
	}

	return deletedMap
}
func floatingIPTargetPgDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap[floatingIPTargetsMoreInfo] = deletedItem.MoreInfo
	}

	return deletedMap
}
