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

const (
	isInstanceBootVolumeIdentifier    = "boot_volume"
	isInstanceBootVolumeManagerDelete = "delete_volume"
)

func ResourceIBMISInstanceBootVolumeManager() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceBootVolumeManagerCreate,
		ReadContext:   resourceIBMISInstanceBootVolumeManagerRead,
		UpdateContext: resourceIBMISInstanceBootVolumeManagerUpdate,
		DeleteContext: resourceIBMISInstanceBootVolumeManagerDelete,
		Exists:        resourceIBMISInstanceBootVolumeManagerExists,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
			isInstanceBootVolumeIdentifier: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier for the boot volume",
			},
			isVolumeName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeName),
				Description:  "The user-defined name for this boot volume",
			},
			isVolumeProfileName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeProfileName),
				Description:  "The globally unique name of the volume profile to use for this volume",
			},
			isVolumeCapacity: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeCapacity),
				Description:  "The capacity of the volume in gigabytes",
			},
			isVolumeIops: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeIops),
				Description:  "The maximum I/O operations per second (IOPS) for the volume",
			},
			isInstanceBootVolumeManagerDelete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to true, the boot volume will be deleted when this resource is destroyed",
			},
			isVolumeTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "User tags for the boot volume",
			},
			isVolumeAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Access management tags for the boot volume",
			},
			isVolumeDeleteAllSnapshots: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to true, all snapshots created from this volume will be deleted when the volume is deleted",
			},

			// Computed attributes
			isVolumeZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone where this volume resides",
			},
			isVolumeEncryptionKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the encryption key used to encrypt this volume",
			},
			isVolumeEncryptionType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of encryption used on the volume",
			},
			isVolumeSourceSnapshot: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for the snapshot from which this volume was created",
			},
			isVolumeResourceGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this volume",
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
			isVolumeCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the volume",
			},
			isVolumeStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the volume",
			},
			isVolumeHealthState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health state of this volume",
			},
			isVolumeHealthReasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current health_state (if any)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeHealthReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state",
						},
						isVolumeHealthReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state",
						},
						isVolumeHealthReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state",
						},
					},
				},
			},
			isVolumeStatusReasons: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeStatusReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason",
						},
						isVolumeStatusReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason",
						},
						isVolumeStatusReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason",
						},
					},
				},
			},
			isVolumeBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume",
			},
			"volume_attachments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The volume attachments for this volume",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_volume_on_instance_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, this volume will be deleted when the instance is deleted",
						},
						"device": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A unique identifier for the device which is exposed to the instance operating system",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A unique identifier for the device which is exposed to the instance operating system",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this volume attachment",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this volume attachment",
						},
						"instance": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The attached instance",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this virtual server instance",
									},
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this virtual server instance",
									},
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this virtual server instance",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this virtual server instance",
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this volume attachment",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of volume attachment",
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

func ResourceIBMISInstanceBootVolumeManagerValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceBootVolumeIdentifier,
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})
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
			Identifier:                 isVolumeProfileName,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "general-purpose, 5iops-tier, 10iops-tier, custom",
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeCapacity,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "10",
			MaxValue:                   "16000"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeIops,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "100",
			MaxValue:                   "48000"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISInstanceBootVolumeManagerResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_boot_volume_manager", Schema: validateSchema}
	return &ibmISInstanceBootVolumeManagerResourceValidator
}

func resourceIBMISInstanceBootVolumeManagerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bootVolumeID := d.Get(isInstanceBootVolumeIdentifier).(string)

	// Verify the boot volume exists
	getVolumeOptions := &vpcv1.GetVolumeOptions{
		ID: &bootVolumeID,
	}
	volume, _, err := sess.GetVolumeWithContext(context, getVolumeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(bootVolumeID)

	// Handle tags on creation
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isVolumeTags); ok || v != "" {
		_, newList := d.GetChange(isVolumeTags)
		userTagsArray := make([]string, 0)
		if newList != nil {
			for _, tag := range newList.(*schema.Set).List() {
				userTagsArray = append(userTagsArray, tag.(string))
			}
		}
		if v != "" {
			envTags := strings.Split(v, ",")
			userTagsArray = append(userTagsArray, envTags...)
		}
		if len(userTagsArray) > 0 {
			volumePatchModel := &vpcv1.VolumePatch{
				UserTags: userTagsArray,
			}
			volumePatch, err := volumePatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
				ID:          &bootVolumeID,
				VolumePatch: volumePatch,
			}
			_, _, err = sess.UpdateVolumeWithContext(context, updateVolumeOptions)
			if err != nil {
				log.Printf("Error on create of boot volume manager (%s) tags: %s", d.Id(), err)
			}
		}
	}

	if _, ok := d.GetOk(isVolumeAccessTags); ok {
		oldList, newList := d.GetChange(isVolumeAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *volume.CRN, "", isVolumeAccessTagType)
		if err != nil {
			log.Printf("Error on create of boot volume manager (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISInstanceBootVolumeManagerUpdate(context, d, meta)
}

func resourceIBMISInstanceBootVolumeManagerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	id := d.Id()
	getVolumeOptions := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	volume, response, err := sess.GetVolumeWithContext(context, getVolumeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(isInstanceBootVolumeId, id); err != nil {
		err = fmt.Errorf("Error setting boot_volume: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-boot_volume").GetDiag()
	}

	if !core.IsNil(volume.Name) {
		if err = d.Set(isVolumeName, *volume.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(volume.Profile) {
		if err = d.Set(isVolumeProfileName, *volume.Profile.Name); err != nil {
			err = fmt.Errorf("Error setting profile: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-profile").GetDiag()
		}
	}

	if !core.IsNil(volume.Zone) {
		if err = d.Set(isVolumeZone, *volume.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-zone").GetDiag()
		}
	}

	if volume.EncryptionKey != nil {
		if err = d.Set(isVolumeEncryptionKey, volume.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-encryption_key").GetDiag()
		}
	}

	if volume.Encryption != nil {
		if err = d.Set(isVolumeEncryptionType, *volume.Encryption); err != nil {
			err = fmt.Errorf("Error setting encryption_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-encryption_type").GetDiag()
		}
	}

	if !core.IsNil(volume.Capacity) {
		if err = d.Set(isVolumeCapacity, flex.IntValue(volume.Capacity)); err != nil {
			err = fmt.Errorf("Error setting capacity: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-capacity").GetDiag()
		}
	}

	if !core.IsNil(volume.Iops) {
		if err = d.Set(isVolumeIops, flex.IntValue(volume.Iops)); err != nil {
			err = fmt.Errorf("Error setting iops: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-iops").GetDiag()
		}
	}

	if !core.IsNil(volume.CRN) {
		if err = d.Set(isVolumeCrn, *volume.CRN); err != nil {
			err = fmt.Errorf("Error setting crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-crn").GetDiag()
		}
	}

	if volume.SourceSnapshot != nil {
		if err = d.Set(isVolumeSourceSnapshot, *volume.SourceSnapshot.ID); err != nil {
			err = fmt.Errorf("Error setting source_snapshot: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-source_snapshot").GetDiag()
		}
	}

	if !core.IsNil(volume.Status) {
		if err = d.Set(isVolumeStatus, *volume.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-status").GetDiag()
		}
	}

	if volume.HealthState != nil {
		if err = d.Set(isVolumeHealthState, *volume.HealthState); err != nil {
			err = fmt.Errorf("Error setting health_state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-health_state").GetDiag()
		}
	}

	if !core.IsNil(volume.Bandwidth) {
		if err = d.Set(isVolumeBandwidth, flex.IntValue(volume.Bandwidth)); err != nil {
			err = fmt.Errorf("Error setting bandwidth: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-bandwidth").GetDiag()
		}
	}

	// Handle resource group
	resourceGroupList := []map[string]interface{}{}
	if volume.ResourceGroup != nil {
		resourceGroupMap := map[string]interface{}{
			"href": *volume.ResourceGroup.Href,
			"id":   *volume.ResourceGroup.ID,
			"name": *volume.ResourceGroup.Name,
		}
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		if err = d.Set(flex.ResourceGroupName, *volume.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_group_name").GetDiag()
		}
	}
	if err = d.Set(isVolumeResourceGroup, resourceGroupList); err != nil {
		err = fmt.Errorf("Error setting resource_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_group").GetDiag()
	}

	// Handle status reasons
	if volume.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range volume.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isVolumeStatusReasonsCode] = *sr.Code
				currentSR[isVolumeStatusReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isVolumeStatusReasonsMoreInfo] = *sr.MoreInfo
				}
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
		if err = d.Set(isVolumeStatusReasons, statusReasonsList); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-status_reasons").GetDiag()
		}
	}

	// Handle health reasons
	if volume.HealthReasons != nil {
		healthReasonsList := make([]map[string]interface{}, 0)
		for _, hr := range volume.HealthReasons {
			currentHR := map[string]interface{}{}
			if hr.Code != nil && hr.Message != nil {
				currentHR[isVolumeHealthReasonsCode] = *hr.Code
				currentHR[isVolumeHealthReasonsMessage] = *hr.Message
				if hr.MoreInfo != nil {
					currentHR[isVolumeHealthReasonsMoreInfo] = *hr.MoreInfo
				}
				healthReasonsList = append(healthReasonsList, currentHR)
			}
		}
		if err = d.Set(isVolumeHealthReasons, healthReasonsList); err != nil {
			err = fmt.Errorf("Error setting health_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-health_reasons").GetDiag()
		}
	}

	// Handle volume attachments
	if volume.VolumeAttachments != nil {
		volumeAttachmentsList := make([]map[string]interface{}, 0)
		for _, va := range volume.VolumeAttachments {
			volumeAttachmentMap := map[string]interface{}{}
			if va.DeleteVolumeOnInstanceDelete != nil {
				volumeAttachmentMap["delete_volume_on_instance_delete"] = *va.DeleteVolumeOnInstanceDelete
			}
			if va.Device != nil {
				deviceList := []map[string]interface{}{}
				deviceMap := map[string]interface{}{
					"id": *va.Device.ID,
				}
				deviceList = append(deviceList, deviceMap)
				volumeAttachmentMap["device"] = deviceList
			}
			if va.Href != nil {
				volumeAttachmentMap["href"] = *va.Href
			}
			if va.ID != nil {
				volumeAttachmentMap["id"] = *va.ID
			}
			if va.Instance != nil {
				instanceList := []map[string]interface{}{}
				instanceMap := map[string]interface{}{}
				if va.Instance.CRN != nil {
					instanceMap["crn"] = *va.Instance.CRN
				}
				if va.Instance.Href != nil {
					instanceMap["href"] = *va.Instance.Href
				}
				if va.Instance.ID != nil {
					instanceMap["id"] = *va.Instance.ID
				}
				if va.Instance.Name != nil {
					instanceMap["name"] = *va.Instance.Name
				}
				instanceList = append(instanceList, instanceMap)
				volumeAttachmentMap["instance"] = instanceList
			}
			if va.Name != nil {
				volumeAttachmentMap["name"] = *va.Name
			}
			if va.Type != nil {
				volumeAttachmentMap["type"] = *va.Type
			}
			volumeAttachmentsList = append(volumeAttachmentsList, volumeAttachmentMap)
		}
		if err = d.Set("volume_attachments", volumeAttachmentsList); err != nil {
			err = fmt.Errorf("Error setting volume_attachments: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-volume_attachments").GetDiag()
		}
	}

	// Handle tags
	if volume.UserTags != nil {
		if err = d.Set(isVolumeTags, volume.UserTags); err != nil {
			err = fmt.Errorf("Error setting tags: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-tags").GetDiag()
		}
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *volume.CRN, "", isVolumeAccessTagType)
	if err != nil {
		log.Printf("Error on get of boot volume manager (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isVolumeAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-access_tags").GetDiag()
	}

	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/storage/storageVolumes"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *volume.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *volume.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *volume.Status); err != nil {
		err = fmt.Errorf("Error setting resource_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "read", "set-resource_status").GetDiag()
	}

	return nil
}

func resourceIBMISInstanceBootVolumeManagerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	id := d.Id()
	hasChanged := false

	// Handle volume property updates with instance state management
	if d.HasChange(isVolumeProfileName) || d.HasChange(isVolumeIops) || d.HasChange(isVolumeCapacity) {
		// Get volume details to check attachments
		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		volume, response, err := sess.GetVolumeWithContext(context, getVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		// Check if volume is attached to an instance
		if volume.VolumeAttachments != nil && len(volume.VolumeAttachments) > 0 {
			instanceID := *volume.VolumeAttachments[0].Instance.ID

			// Get instance status
			getInstanceOptions := &vpcv1.GetInstanceOptions{
				ID: &instanceID,
			}
			instance, _, err := sess.GetInstanceWithContext(context, getInstanceOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			// Ensure instance is running for volume updates
			if instance != nil && *instance.Status != "running" {
				actionType := "start"
				createInstanceActionOptions := &vpcv1.CreateInstanceActionOptions{
					InstanceID: &instanceID,
					Type:       &actionType,
				}
				_, _, err = sess.CreateInstanceActionWithContext(context, createInstanceActionOptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForInstanceAvailable(sess, instanceID, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		// Get fresh ETag for the update
		getVolumeOptions = &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		_, response, err = sess.GetVolumeWithContext(context, getVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")

		volumePatchModel := &vpcv1.VolumePatch{}

		if d.HasChange(isVolumeProfileName) {
			profile := d.Get(isVolumeProfileName).(string)
			volumePatchModel.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &profile,
			}
		}

		if d.HasChange(isVolumeIops) {
			iops := int64(d.Get(isVolumeIops).(int))
			volumePatchModel.Iops = &iops
		}

		if d.HasChange(isVolumeCapacity) {
			capacity := int64(d.Get(isVolumeCapacity).(int))
			volumePatchModel.Capacity = &capacity
		}

		volumePatch, err := volumePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
			ID:          &id,
			VolumePatch: volumePatch,
			IfMatch:     &eTag,
		}

		_, _, err = sess.UpdateVolumeWithContext(context, updateVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForVolumeAvailable(sess, id, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		hasChanged = true
	}

	// Handle name update
	if d.HasChange(isVolumeName) {
		name := d.Get(isVolumeName).(string)
		volumePatchModel := &vpcv1.VolumePatch{
			Name: &name,
		}
		volumePatch, err := volumePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
			ID:          &id,
			VolumePatch: volumePatch,
		}

		_, _, err = sess.UpdateVolumeWithContext(context, updateVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		hasChanged = true
	}

	// Handle tags update
	if d.HasChange(isVolumeTags) {
		userTagsSet := d.Get(isVolumeTags).(*schema.Set)
		userTagsArray := make([]string, 0)
		if userTagsSet != nil && userTagsSet.Len() > 0 {
			for _, userTag := range userTagsSet.List() {
				userTagsArray = append(userTagsArray, userTag.(string))
			}
		}

		// Add environment tags
		envTags := os.Getenv("IC_ENV_TAGS")
		if envTags != "" {
			envTagsArray := strings.Split(envTags, ",")
			userTagsArray = append(userTagsArray, envTagsArray...)
		}

		volumePatchModel := &vpcv1.VolumePatch{
			UserTags: userTagsArray,
		}
		volumePatch, err := volumePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
			ID:          &id,
			VolumePatch: volumePatch,
		}

		_, _, err = sess.UpdateVolumeWithContext(context, updateVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		hasChanged = true
	}

	// Handle access tags update
	if d.HasChange(isVolumeAccessTags) {
		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		volume, _, err := sess.GetVolumeWithContext(context, getVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		oldList, newList := d.GetChange(isVolumeAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *volume.CRN, "", isVolumeAccessTagType)
		if err != nil {
			log.Printf("Error on update of boot volume manager (%s) access tags: %s", id, err)
		}
		hasChanged = true
	}

	// Handle snapshot deletion if requested
	if deleteAllSnapshots, ok := d.GetOk(isVolumeDeleteAllSnapshots); ok && deleteAllSnapshots.(bool) {
		deleteSnapshotsOptions := &vpcv1.DeleteSnapshotsOptions{
			SourceVolumeID: &id,
		}
		_, err := sess.DeleteSnapshotsWithContext(context, deleteSnapshotsOptions)
		if err != nil {
			log.Printf("Error deleting snapshots from boot volume (%s): %s", id, err)
		}
	}

	if hasChanged {
		return resourceIBMISInstanceBootVolumeManagerRead(context, d, meta)
	}

	return nil
}

func resourceIBMISInstanceBootVolumeManagerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	id := d.Id()

	// Check if force delete is enabled
	if d.Get(isInstanceBootVolumeManagerDelete).(bool) {
		// Handle snapshot deletion if requested
		if deleteAllSnapshots, ok := d.GetOk(isVolumeDeleteAllSnapshots); ok && deleteAllSnapshots.(bool) {
			deleteSnapshotsOptions := &vpcv1.DeleteSnapshotsOptions{
				SourceVolumeID: &id,
			}
			_, err := sess.DeleteSnapshotsWithContext(context, deleteSnapshotsOptions)
			if err != nil {
				log.Printf("Error deleting snapshots from boot volume (%s): %s", id, err)
			}
		}

		deleteVolumeOptions := &vpcv1.DeleteVolumeOptions{
			ID: &id,
		}
		_, err := sess.DeleteVolumeWithContext(context, deleteVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForVolumeDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeDeleted failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId("")
	return nil
}

func resourceIBMISInstanceBootVolumeManagerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_boot_volume_manager", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	id := d.Id()
	getVolumeOptions := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	_, response, err := sess.GetVolume(getVolumeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolume failed: %s", err.Error()), "ibm_is_instance_boot_volume_manager", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
