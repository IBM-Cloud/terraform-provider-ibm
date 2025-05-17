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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceId                             = "instance"
	isInstanceVolAttVol                      = "volume"
	isInstanceVolAttTags                     = "tags"
	isInstanceVolAttId                       = "volume_attachment_id"
	isInstanceVolAttIops                     = "volume_iops"
	isInstanceExistingVolume                 = "existing"
	isInstanceVolAttName                     = "name"
	isInstanceVolAttVolume                   = "volume"
	isInstanceVolumeDeleteOnInstanceDelete   = "delete_volume_on_instance_delete"
	isInstanceVolumeDeleteOnAttachmentDelete = "delete_volume_on_attachment_delete"
	isInstanceVolCapacity                    = "capacity"
	isInstanceVolIops                        = "iops"
	isInstanceVolEncryptionKey               = "encryption_key"
	isInstanceVolProfile                     = "profile"
)

func ResourceIBMISInstanceVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMisInstanceVolumeAttachmentCreate,
		ReadContext:   resourceIBMisInstanceVolumeAttachmentRead,
		UpdateContext: resourceIBMisInstanceVolumeAttachmentUpdate,
		DeleteContext: resourceIBMisInstanceVolumeAttachmentDelete,
		Exists:        resourceIBMisInstanceVolumeAttachmentExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},
		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceVolumeValidate(diff)
				}),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				}),
		),
		Schema: map[string]*schema.Schema{
			isInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceId),
				Description:  "Instance id",
			},
			isInstanceVolAttId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this volume attachment",
			},

			isInstanceVolAttName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceVolAttName),
				Description:  "The user-defined name for this volume attachment.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.",
			},

			isInstanceVolumeDeleteOnInstanceDelete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "If set to true, when deleting the instance the volume will also be deleted.",
			},
			isInstanceVolumeDeleteOnAttachmentDelete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If set to true, when deleting the attachment, the volume will also be deleted. Default value for this true.",
			},
			isInstanceVolAttVol: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{isInstanceVolIops, isInstanceVolumeAttVolumeReferenceName, isInstanceVolProfile, isInstanceVolCapacity, isInstanceVolumeSnapshot, isInstanceVolAttTags},
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceName),
				Description:   "Instance id",
			},

			isInstanceVolIops: {
				Type:          schema.TypeInt,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{isInstanceVolAttVol},
				Description:   "The maximum I/O operations per second (IOPS) for the volume.",
			},

			isInstanceVolumeAttVolumeReferenceName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceVolumeAttVolumeReferenceName),
				Description:  "The unique user-defined name for this volume",
			},

			isInstanceVolAttTags: {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isInstanceVolAttVol},
				Elem:          &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_volume_attachment", "tags")},
				Set:           flex.ResourceIBMVPCHash,
				Description:   "UserTags for the volume instance",
			},

			isInstanceVolProfile: {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{isInstanceVolAttVol},
				Computed:      true,
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceVolProfile),
				Description:   "The  globally unique name for the volume profile to use for this volume.",
			},

			isInstanceVolCapacity: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				AtLeastOneOf:  []string{isInstanceVolAttVol, isInstanceVolCapacity, isInstanceVolumeSnapshot, isInstanceVolumeSnapshotCrn},
				ConflictsWith: []string{isInstanceVolAttVol},
				ValidateFunc:  validate.InvokeValidator("ibm_is_instance_volume_attachment", isInstanceVolCapacity),
				Description:   "The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.",
			},
			isInstanceVolEncryptionKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.",
			},
			isInstanceVolumeSnapshot: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				AtLeastOneOf:  []string{isInstanceVolAttVol, isInstanceVolCapacity, isInstanceVolumeSnapshot, isInstanceVolumeSnapshotCrn},
				ConflictsWith: []string{isInstanceVolAttVol, isInstanceVolumeSnapshotCrn},
				Description:   "The snapshot ID of the volume to be attached",
			},
			isInstanceVolumeSnapshotCrn: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				AtLeastOneOf:  []string{isInstanceVolAttVol, isInstanceVolCapacity, isInstanceVolumeSnapshot, isInstanceVolumeSnapshotCrn},
				ConflictsWith: []string{isInstanceVolAttVol, isInstanceVolumeSnapshot},
				Description:   "The snapshot crn of the volume to be attached",
			},
			isInstanceVolumeAttVolumeReferenceCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this volume",
			},
			isInstanceVolumeAttVolumeReferenceDeleted: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link to documentation about deleted resources",
			},
			isInstanceVolumeAttVolumeReferenceHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this volume",
			},

			isInstanceVolumeAttDevice: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique identifier for the device which is exposed to the instance operating system",
			},

			isInstanceVolumeAttHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this volume attachment",
			},

			isInstanceVolumeAttStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this volume attachment, one of [ attached, attaching, deleting, detaching ]",
			},

			isInstanceVolumeAttType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of volume attachment one of [ boot, data ]",
			},

			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMISInstanceVolumeAttachmentValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceId,
			ValidateFunctionIdentifier: validate.ValidateNoZeroValues,
			Type:                       validate.TypeString})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceVolAttName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceVolCapacity,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "10",
			MaxValue:                   "64000"})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceVolumeAttVolumeReferenceName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceVolProfile,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "general-purpose, 5iops-tier, 10iops-tier, custom, sdp",
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISInstanceVolumeAttachmentValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_volume_attachment", Schema: validateSchema}
	return &ibmISInstanceVolumeAttachmentValidator
}

func instanceVolAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}, instanceId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	instanceVolAttproto := &vpcv1.CreateInstanceVolumeAttachmentOptions{
		InstanceID: &instanceId,
	}
	volumeIdStr := ""
	if volumeId, ok := d.GetOk(isInstanceVolAttVol); ok {
		volumeIdStr = volumeId.(string)
	}
	if volumeIdStr != "" {
		var volProtoVol = &vpcv1.VolumeAttachmentPrototypeVolumeVolumeIdentity{}
		volProtoVol.ID = &volumeIdStr
		instanceVolAttproto.Volume = volProtoVol
	} else {
		var volProtoVol = &vpcv1.VolumeAttachmentPrototypeVolumeVolumePrototypeInstanceContext{}
		if volname, ok := d.GetOk(isInstanceVolumeAttVolumeReferenceName); ok {
			volnamestr := volname.(string)
			volProtoVol.Name = &volnamestr
		}
		// var userTags *schema.Set
		// if v, ok := d.GetOk(isInstanceVolAttTags); ok {
		// 	userTags = v.(*schema.Set)
		// 	if userTags != nil && userTags.Len() != 0 {
		// 		userTagsArray := make([]string, userTags.Len())
		// 		for i, userTag := range userTags.List() {
		// 			userTagStr := userTag.(string)
		// 			userTagsArray[i] = userTagStr
		// 		}
		// 		schematicTags := os.Getenv("IC_ENV_TAGS")
		// 		var envTags []string
		// 		if schematicTags != "" {
		// 			envTags = strings.Split(schematicTags, ",")
		// 			userTagsArray = append(userTagsArray, envTags...)
		// 		}
		// 		volProtoVol.UserTags = userTagsArray
		// 	}
		// }
		volSnapshotStr := ""
		volSnapshotCrnStr := ""
		if volSnapshot, ok := d.GetOk(isInstanceVolumeSnapshot); ok {
			volSnapshotStr = volSnapshot.(string)
			volProtoVol.SourceSnapshot = &vpcv1.SnapshotIdentity{
				ID: &volSnapshotStr,
			}
		}
		if volSnapshotCrn, ok := d.GetOk(isInstanceVolumeSnapshotCrn); ok {
			volSnapshotCrnStr = volSnapshotCrn.(string)
			volProtoVol.SourceSnapshot = &vpcv1.SnapshotIdentity{
				CRN: &volSnapshotCrnStr,
			}
		}
		encryptionCRNStr := ""
		if encryptionCRN, ok := d.GetOk(isInstanceVolEncryptionKey); ok {
			encryptionCRNStr = encryptionCRN.(string)
			volProtoVol.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
				CRN: &encryptionCRNStr,
			}
		}
		var snapCapacity int64
		if volSnapshotStr != "" || volSnapshotCrnStr != "" {
			if volSnapshotStr == "" {
				volSnapshotStr = volSnapshotCrnStr[strings.LastIndex(volSnapshotCrnStr, ":")+1:]
			}
			snapshotGet, _, err := sess.GetSnapshotWithContext(context, &vpcv1.GetSnapshotOptions{
				ID: &volSnapshotStr,
			})
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			snapCapacity = int64(int(*snapshotGet.MinimumCapacity))
		}
		var volCapacityInt int64
		if volCapacity, ok := d.GetOk(isInstanceVolCapacity); ok {
			volCapacityInt = int64(volCapacity.(int))
			if volCapacityInt != 0 && volCapacityInt > snapCapacity {
				volProtoVol.Capacity = &volCapacityInt
			}
		}
		// bandwidth changes
		var volBandwidthInt int64
		if volBandwidth, ok := d.GetOk("bandwidth"); ok {
			volBandwidthInt = int64(volBandwidth.(int))
			if volBandwidthInt != 0 {
				volProtoVol.Bandwidth = &volBandwidthInt
			}
		}
		var iops int64
		if volIops, ok := d.GetOk(isInstanceVolIops); ok {
			iops = int64(volIops.(int))
			if iops != 0 {
				volProtoVol.Iops = &iops
			}
			volProfileStr := d.Get(isInstanceVolProfile).(string)
			if volProfileStr == "" {
				volProfileStr = "custom"
			}
			volProtoVol.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &volProfileStr,
			}
		} else {
			volProfileStr := "general-purpose"
			if volProfile, ok := d.GetOk(isInstanceVolProfile); ok {
				volProfileStr = volProfile.(string)
				volProtoVol.Profile = &vpcv1.VolumeProfileIdentity{
					Name: &volProfileStr,
				}
			} else {
				volProtoVol.Profile = &vpcv1.VolumeProfileIdentity{
					Name: &volProfileStr,
				}
			}
		}

		instanceVolAttproto.Volume = volProtoVol
	}

	if autoDelete, ok := d.GetOk(isInstanceVolumeDeleteOnInstanceDelete); ok {
		autoDeleteBool := autoDelete.(bool)
		instanceVolAttproto.DeleteVolumeOnInstanceDelete = &autoDeleteBool
	}
	if name, ok := d.GetOk(isInstanceVolAttName); ok {
		namestr := name.(string)
		instanceVolAttproto.Name = &namestr
	}

	isInstanceKey := "instance_key_" + instanceId
	conns.IbmMutexKV.Lock(isInstanceKey)
	defer conns.IbmMutexKV.Unlock(isInstanceKey)

	instanceVolAtt, _, err := sess.CreateInstanceVolumeAttachmentWithContext(context, instanceVolAttproto)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceVolumeAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(makeTerraformVolAttID(instanceId, *instanceVolAtt.ID))
	volAtt, err := isWaitForInstanceVolumeAttached(sess, d, instanceId, *instanceVolAtt.ID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceVolumeAttached failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isInstanceVolAttTags); ok || v != "" {
		volAttRef := volAtt.(*vpcv1.VolumeAttachment)
		oldList, newList := d.GetChange(isInstanceVolAttTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *volAttRef.Volume.CRN, "", isInstanceUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource instance volume attachment (%s) tags: %s", d.Id(), err)
		}
	}
	log.Printf("[INFO] Instance (%s) volume attachment : %s", instanceId, *instanceVolAtt.ID)
	return nil
}

func resourceIBMisInstanceVolumeAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceId := d.Get(isInstanceId).(string)
	err := instanceVolAttachmentCreate(context, d, meta, instanceId)
	if err != nil {
		return err
	}
	return resourceIBMisInstanceVolumeAttachmentRead(context, d, meta)
}

func resourceIBMisInstanceVolumeAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceID, id, err := parseVolAttTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "sep-id-parts").GetDiag()
	}
	diagErr := instanceVolumeAttachmentGet(context, d, meta, instanceID, id)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func instanceVolumeAttachmentGet(context context.Context, d *schema.ResourceData, meta interface{}, instanceId, id string) diag.Diagnostics {
	instanceC, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getinsVolAttOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
		InstanceID: &instanceId,
		ID:         &id,
	}
	volumeAtt, response, err := instanceC.GetInstanceVolumeAttachmentWithContext(context, getinsVolAttOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceVolumeAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(isInstanceId, instanceId)

	if volumeAtt.Volume != nil {
		if err = d.Set(isInstanceVolumeAttVolumeReferenceName, *volumeAtt.Volume.Name); err != nil {
			err = fmt.Errorf("Error setting volume_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume_name").GetDiag()
		}
		if err = d.Set(isInstanceVolumeAttVolumeReferenceCrn, *volumeAtt.Volume.CRN); err != nil {
			err = fmt.Errorf("Error setting volume_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume_crn").GetDiag()
		}
		if volumeAtt.Volume.Deleted != nil {
			if err = d.Set(isInstanceVolumeAttVolumeReferenceDeleted, *volumeAtt.Volume.Deleted.MoreInfo); err != nil {
				err = fmt.Errorf("Error setting volume_deleted: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume_deleted").GetDiag()
			}
		}
		if err = d.Set(isInstanceVolumeAttVolumeReferenceHref, *volumeAtt.Volume.Href); err != nil {
			err = fmt.Errorf("Error setting volume_href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume_href").GetDiag()
		}
	}
	if err = d.Set(isInstanceVolumeDeleteOnInstanceDelete, *volumeAtt.DeleteVolumeOnInstanceDelete); err != nil {
		err = fmt.Errorf("Error setting delete_volume_on_instance_delete: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-delete_volume_on_instance_delete").GetDiag()
	}
	if err = d.Set(isInstanceVolAttName, *volumeAtt.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-name").GetDiag()
	}
	if volumeAtt.Device != nil {
		if err = d.Set(isInstanceVolumeAttDevice, *volumeAtt.Device.ID); err != nil {
			err = fmt.Errorf("Error setting device: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-device").GetDiag()
		}
	}
	if err = d.Set(isInstanceVolumeAttHref, *volumeAtt.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-href").GetDiag()
	}
	if err = d.Set(isInstanceVolAttId, *volumeAtt.ID); err != nil {
		err = fmt.Errorf("Error setting volume_attachment_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume_attachment_id").GetDiag()
	}
	if err = d.Set(isInstanceVolumeAttStatus, *volumeAtt.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-status").GetDiag()
	}
	if err = d.Set(isInstanceVolumeAttType, *volumeAtt.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-type").GetDiag()
	}
	if err = d.Set("version", response.Headers.Get("Etag")); err != nil {
		err = fmt.Errorf("Error setting version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-version").GetDiag()
	}
	volId := *volumeAtt.Volume.ID
	getVolOptions := &vpcv1.GetVolumeOptions{
		ID: &volId,
	}
	volumeDetail, _, err := instanceC.GetVolumeWithContext(context, getVolOptions)
	if err != nil || volumeDetail == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(isInstanceVolAttVol, *volumeDetail.ID); err != nil {
		err = fmt.Errorf("Error setting volume: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-volume").GetDiag()
	}
	if err = d.Set(isInstanceVolIops, *volumeDetail.Iops); err != nil {
		err = fmt.Errorf("Error setting iops: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-iops").GetDiag()
	}
	if err = d.Set(isInstanceVolProfile, *volumeDetail.Profile.Name); err != nil {
		err = fmt.Errorf("Error setting profile: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-profile").GetDiag()
	}
	if err = d.Set(isInstanceVolCapacity, *volumeDetail.Capacity); err != nil {
		err = fmt.Errorf("Error setting capacity: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-capacity").GetDiag()
	}
	// bandwidth changes
	if err = d.Set("bandwidth", volumeDetail.Bandwidth); err != nil {
		err = fmt.Errorf("Error setting bandwidth: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-bandwidth").GetDiag()
	}
	if volumeDetail.EncryptionKey != nil {
		if err = d.Set(isInstanceVolEncryptionKey, *volumeDetail.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-encryption_key").GetDiag()
		}
	}
	if volumeDetail.SourceSnapshot != nil {
		if err = d.Set(isInstanceVolumeSnapshot, *volumeDetail.SourceSnapshot.ID); err != nil {
			err = fmt.Errorf("Error setting snapshot: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-snapshot").GetDiag()
		}
		if err = d.Set(isInstanceVolumeSnapshotCrn, *volumeDetail.SourceSnapshot.CRN); err != nil {
			err = fmt.Errorf("Error setting snapshot_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-snapshot_crn").GetDiag()
		}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *volumeDetail.CRN, "", isInstanceUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource Instance  volume attachment (%s) tags: %s", d.Id(), err)
	}

	if err = d.Set(isInstanceTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "read", "set-tags").GetDiag()
	}
	return nil
}

func instanceVolAttUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceC, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceId, id, err := parseVolAttTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "update", "sep-id-parts").GetDiag()
	}
	if volumecrnok, ok := d.GetOk("volume_crn"); ok {
		volumecrn := volumecrnok.(string)
		if d.HasChange(isInstanceTags) {
			oldList, newList := d.GetChange(isInstanceTags)
			err = flex.UpdateTagsUsingCRN(oldList, newList, meta, volumecrn)
			if err != nil {
				log.Printf(
					"Error on update of resource Instance volume attachment (%s) tags: %s", d.Id(), err)
			}
		}
	}
	updateInstanceVolAttOptions := &vpcv1.UpdateInstanceVolumeAttachmentOptions{
		InstanceID: &instanceId,
		ID:         &id,
	}
	flag := false

	// name && auto delete change
	volAttNamePatchModel := &vpcv1.VolumeAttachmentPatch{}
	if d.HasChange(isInstanceVolumeDeleteOnInstanceDelete) {
		autoDelete := d.Get(isInstanceVolumeDeleteOnInstanceDelete).(bool)
		volAttNamePatchModel.DeleteVolumeOnInstanceDelete = &autoDelete
		flag = true
	}

	if d.HasChange(isInstanceVolAttName) {
		name := d.Get(isInstanceVolAttName).(string)
		volAttNamePatchModel.Name = &name
		flag = true
	}
	if flag {
		volAttNamePatchModelAsPatch, err := volAttNamePatchModel.AsPatch()
		if err != nil || volAttNamePatchModelAsPatch == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volAttNamePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateInstanceVolAttOptions.VolumeAttachmentPatch = volAttNamePatchModelAsPatch

		instanceVolAttUpdate, _, err := instanceC.UpdateInstanceVolumeAttachmentWithContext(context, updateInstanceVolAttOptions)
		if err != nil || instanceVolAttUpdate == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceVolumeAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	hasNameChanged := d.HasChange(isInstanceVolumeAttVolumeReferenceName)
	hasBandwidthChanged := d.HasChange("bandwidth")
	if hasNameChanged || hasBandwidthChanged {
		volid := d.Get(isInstanceVolAttVol).(string)
		voloptions := &vpcv1.UpdateVolumeOptions{
			ID: &volid,
		}
		if hasNameChanged {
			volumeNamePatchModel := &vpcv1.VolumePatch{}
			newname := d.Get(isInstanceVolumeAttVolumeReferenceName).(string)
			volumeNamePatchModel.Name = &newname
			volumePatch, err := volumeNamePatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeNamePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			voloptions.VolumePatch = volumePatch
			_, _, err = instanceC.UpdateVolumeWithContext(context, voloptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
		if hasBandwidthChanged {
			volumeBandwidthPatchModel := &vpcv1.VolumePatch{}
			newBandwidth := int64(d.Get("bandwidth").(int))
			volumeBandwidthPatchModel.Bandwidth = &newBandwidth
			volumePatch, err := volumeBandwidthPatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeBandwidthPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			voloptions.VolumePatch = volumePatch
			_, _, err = instanceC.UpdateVolumeWithContext(context, voloptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}

	}

	// profile/iops update

	volId := ""
	if volIdOk, ok := d.GetOk(isInstanceVolAttVol); ok {
		volId = volIdOk.(string)
	}
	volProfile := ""
	if volProfileOk, ok := d.GetOk(isInstanceVolProfile); ok {
		volProfile = volProfileOk.(string)
	}
	if volId != "" && d.HasChange(isInstanceVolIops) && !d.HasChange(isInstanceVolProfile) && volProfile == "sdp" { // || d.HasChange(isInstanceVolAttTags)
		updateVolumeProfileOptions := &vpcv1.UpdateVolumeOptions{
			ID: &volId,
		}
		volumeProfilePatchModel := &vpcv1.VolumePatch{}
		if d.HasChange(isVolumeIops) {
			iops := int64(d.Get(isVolumeIops).(int))
			volumeProfilePatchModel.Iops = &iops
		}
		volumeProfilePatch, err := volumeProfilePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeProfilePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		optionsget := &vpcv1.GetVolumeOptions{
			ID: &volId,
		}
		_, response, err := instanceC.GetVolumeWithContext(context, optionsget)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")
		updateVolumeProfileOptions.IfMatch = &eTag
		updateVolumeProfileOptions.VolumePatch = volumeProfilePatch
		_, response, err = instanceC.UpdateVolumeWithContext(context, updateVolumeProfileOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(instanceC, volId, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else if volId != "" && (d.HasChange(isInstanceVolIops) || d.HasChange(isInstanceVolProfile)) { // || d.HasChange(isInstanceVolAttTags)
		insId := d.Get(isInstanceId).(string)
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &insId,
		}
		instance, response, err := instanceC.GetInstanceWithContext(context, getinsOptions)
		if err != nil || instance == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		if instance != nil && *instance.Status != "running" {
			actiontype := "start"
			createinsactoptions := &vpcv1.CreateInstanceActionOptions{
				InstanceID: &insId,
				Type:       &actiontype,
			}
			_, response, err = instanceC.CreateInstanceActionWithContext(context, createinsactoptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForInstanceAvailable(instanceC, insId, d.Timeout(schema.TimeoutCreate), d)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
		updateVolumeProfileOptions := &vpcv1.UpdateVolumeOptions{
			ID: &volId,
		}
		volumeProfilePatchModel := &vpcv1.VolumePatch{}
		if d.HasChange(isInstanceVolProfile) {
			profile := d.Get(isInstanceVolProfile).(string)
			volumeProfilePatchModel.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &profile,
			}
		} else if d.HasChange(isVolumeIops) {
			profile := d.Get(isInstanceVolProfile).(string)
			volumeProfilePatchModel.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &profile,
			}
			iops := int64(d.Get(isVolumeIops).(int))
			volumeProfilePatchModel.Iops = &iops
		}
		// if d.HasChange(isInstanceVolAttTags) && !d.IsNewResource() {
		// 	if v, ok := d.GetOk(isInstanceVolAttTags); ok {
		// 		userTags := v.(*schema.Set)
		// 		if userTags != nil && userTags.Len() != 0 {
		// 			userTagsArray := make([]string, userTags.Len())
		// 			for i, userTag := range userTags.List() {
		// 				userTagStr := userTag.(string)
		// 				userTagsArray[i] = userTagStr
		// 			}
		// 			schematicTags := os.Getenv("IC_ENV_TAGS")
		// 			var envTags []string
		// 			if schematicTags != "" {
		// 				envTags = strings.Split(schematicTags, ",")
		// 				userTagsArray = append(userTagsArray, envTags...)
		// 			}
		// 			volumeProfilePatchModel.UserTags = userTagsArray
		// 		}
		// 	}

		// }

		volumeProfilePatch, err := volumeProfilePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeProfilePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		optionsget := &vpcv1.GetVolumeOptions{
			ID: &volId,
		}
		_, response, err = instanceC.GetVolumeWithContext(context, optionsget)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")
		updateVolumeProfileOptions.IfMatch = &eTag
		updateVolumeProfileOptions.VolumePatch = volumeProfilePatch
		_, response, err = instanceC.UpdateVolumeWithContext(context, updateVolumeProfileOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(instanceC, volId, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	// capacity update

	if volId != "" && d.HasChange(isInstanceVolCapacity) {

		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &volId,
		}
		vol, _, err := instanceC.GetVolumeWithContext(context, getvolumeoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if *vol.Profile.Name != "sdp" {
			if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) == 0 || *vol.VolumeAttachments[0].Name == "" {
				err = fmt.Errorf("Error volume capacity can't be updated since volume %s is not attached to any instance for VolumePatch", id)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			getinsOptions := &vpcv1.GetInstanceOptions{
				ID: &instanceId,
			}
			instance, _, err := instanceC.GetInstanceWithContext(context, getinsOptions)
			if err != nil || instance == nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if instance != nil && *instance.Status != "running" {
				actiontype := "start"
				createinsactoptions := &vpcv1.CreateInstanceActionOptions{
					InstanceID: &instanceId,
					Type:       &actiontype,
				}
				_, _, err = instanceC.CreateInstanceActionWithContext(context, createinsactoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForInstanceAvailable(instanceC, instanceId, d.Timeout(schema.TimeoutCreate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
		capacity := int64(d.Get(isVolumeCapacity).(int))
		updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
			ID: &volId,
		}
		volumeCapacityPatchModel := &vpcv1.VolumePatch{}
		volumeCapacityPatchModel.Capacity = &capacity
		volumeCapacityPatch, err := volumeCapacityPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeCapacityPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateVolumeOptions.VolumePatch = volumeCapacityPatch
		_, _, err = instanceC.UpdateVolumeWithContext(context, updateVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(instanceC, volId, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMisInstanceVolumeAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	err := instanceVolAttUpdate(context, d, meta)
	if err != nil {
		return err
	}
	return resourceIBMisInstanceVolumeAttachmentRead(context, d, meta)
}

func instanceVolAttDelete(context context.Context, d *schema.ResourceData, meta interface{}, instanceId, id, volId string, volDelete bool) diag.Diagnostics {
	instanceC, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceVolAttOptions := &vpcv1.DeleteInstanceVolumeAttachmentOptions{
		InstanceID: &instanceId,
		ID:         &id,
	}

	isInstanceKey := "instance_key_" + instanceId
	conns.IbmMutexKV.Lock(isInstanceKey)
	defer conns.IbmMutexKV.Unlock(isInstanceKey)

	_, err = instanceC.DeleteInstanceVolumeAttachmentWithContext(context, deleteInstanceVolAttOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceVolumeAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForInstanceVolumeDetached(instanceC, d, instanceId, id)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceVolumeDetached failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if volDelete {
		deleteVolumeOptions := &vpcv1.DeleteVolumeOptions{
			ID: &volId,
		}
		_, err := instanceC.DeleteVolumeWithContext(context, deleteVolumeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVolumeWithContext failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeDeleted(instanceC, volId, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeDeleted failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMisInstanceVolumeAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instanceId, id, err := parseVolAttTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "delete", "sep-id-parts").GetDiag()
	}

	volDelete := false
	if volDeleteOk, ok := d.GetOk(isInstanceVolumeDeleteOnAttachmentDelete); ok {
		volDelete = volDeleteOk.(bool)
	}
	volId := ""
	if volIdOk, ok := d.GetOk(isInstanceVolAttVol); ok {
		volId = volIdOk.(string)
	}

	diagErr := instanceVolAttDelete(context, d, meta, instanceId, id, volId, volDelete)
	if diagErr != nil {
		return diagErr
	}
	d.SetId("")
	return nil
}

func resourceIBMisInstanceVolumeAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	instanceId, id, err := parseVolAttTerraformID(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "exists", "sep-id-parts")
	}
	exists, err := instanceVolAttExists(d, meta, instanceId, id)
	return exists, err
}

func instanceVolAttExists(d *schema.ResourceData, meta interface{}, instanceId, id string) (bool, error) {
	instanceC, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_volume_attachment", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getinsvolattOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
		InstanceID: &instanceId,
		ID:         &id,
	}
	_, response, err := instanceC.GetInstanceVolumeAttachment(getinsvolattOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceVolumeAttachment failed: %s", err.Error()), "ibm_is_instance_volume_attachment", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func makeTerraformVolAttID(id1, id2 string) string {
	// Include both instance id and volume attachment to create a unique Terraform id.  As a bonus,
	// we can extract the instance id as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}

func parseVolAttTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}
