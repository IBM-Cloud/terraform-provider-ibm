// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isImageHref                   = "href"
	isImageName                   = "name"
	isImageTags                   = "tags"
	isImageOperatingSystem        = "operating_system"
	isImageStatus                 = "status"
	isImageVisibility             = "visibility"
	isImageFile                   = "file"
	isImageVolume                 = "source_volume"
	isImageMinimumProvisionedSize = "size"

	isImageResourceGroup    = "resource_group"
	isImageEncryptedDataKey = "encrypted_data_key"
	isImageEncryptionKey    = "encryption_key"
	isImageEncryption       = "encryption"
	isImageCheckSum         = "checksum"
	IsImageCRN              = "crn"

	isImageProvisioning     = "provisioning"
	isImageProvisioningDone = "done"
	isImageDeleting         = "deleting"
	isImageDeleted          = "done"

	isImageAccessTags    = "access_tags"
	isImageUserTagType   = "user"
	isImageAccessTagType = "access"

	isImageDeprecate      = "deprecate"
	isImageObsolete       = "obsolete"
	isImageUserDataFormat = "user_data_format"
)

func ResourceIBMISImage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISImageCreate,
		ReadContext:   resourceIBMISImageRead,
		UpdateContext: resourceIBMISImageUpdate,
		DeleteContext: resourceIBMISImageDelete,
		Exists:        resourceIBMISImageExists,
		Importer:      &schema.ResourceImporter{},

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
			isImageHref: {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				RequiredWith:     []string{isImageOperatingSystem},
				ExactlyOneOf:     []string{isImageHref, isImageVolume},
				Description:      "Image Href value",
			},

			isImageName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_image", isImageName),
				Description:  "Image name",
			},

			isImageEncryptedDataKey: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A base64-encoded, encrypted representation of the key that was used to encrypt the data for this image",
			},

			isImageCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the image was created",
			},
			isImageDeprecate: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set to deprecate. You can set an image to `deprecated` as a warning to transition away from soon-to-be obsolete images. Deprecated images can be used to provision resources.",
			},
			isImageObsolete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set to obsolete. You can set an image to `obsolete` as a warning to transition away from soon-to-be deleted images. You can't use obsolete images to provision resources.",
			},
			isImageDeprecationAt: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The deprecation date and time (UTC) for this image. If absent, no deprecation date and time has been set.",
			},
			isImageObsolescenceAt: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The obsolescence date and time (UTC) for this image. If absent, no obsolescence date and time has been set.",
			},

			isImageEncryptionKey: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource",
			},
			isImageTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_image", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Tags for the image",
			},

			isImageOperatingSystem: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{isImageHref},
				Computed:     true,
				Description:  "Image Operating system",
			},

			isImageEncryption: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of encryption used on the image",
			},
			isImageStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this image",
			},

			isImageMinimumProvisionedSize: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum size (in gigabytes) of a volume onto which this image may be provisioned",
			},

			isImageVisibility: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the image is publicly visible or private to the account",
			},

			isImageFile: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Details for the stored image file",
			},

			isImageVolume: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{isImageHref, isImageVolume},
				Description:  "Image volume id",
			},

			isImageResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group for this image",
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

			IsImageCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			isImageCheckSum: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA256 checksum of this image",
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
			isImageAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_image", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
			isImageUserDataFormat: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user data format for this image",
			},
		},
	}
}

func ResourceIBMISImageValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isImageName,
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
	ibmISImageResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_image", Schema: validateSchema}
	return &ibmISImageResourceValidator
}

func resourceIBMISImageCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] Image create")
	href := d.Get(isImageHref).(string)
	name := d.Get(isImageName).(string)
	operatingSystem := d.Get(isImageOperatingSystem).(string)
	volume := d.Get(isImageVolume).(string)

	if volume != "" {
		err := imgCreateByVolume(context, d, meta, name, volume)
		if err != nil {
			return err
		}
	} else {
		err := imgCreateByFile(context, d, meta, href, name, operatingSystem)
		if err != nil {
			return err
		}
	}

	return resourceIBMISImageRead(context, d, meta)
}

func imgCreateByFile(context context.Context, d *schema.ResourceData, meta interface{}, href, name, operatingSystem string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	imagePrototype := &vpcv1.ImagePrototypeImageByFile{
		Name: &name,
		File: &vpcv1.ImageFilePrototype{
			Href: &href,
		},
		OperatingSystem: &vpcv1.OperatingSystemIdentity{
			Name: &operatingSystem,
		},
	}
	if obsoleteAtOk, ok := d.GetOk(isImageObsolescenceAt); ok {
		obsoleteAt, err := strfmt.ParseDateTime(obsoleteAtOk.(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "parse-obsolescence_at").GetDiag()
		}
		imagePrototype.ObsolescenceAt = &obsoleteAt
	}
	if deprecateAtOk, ok := d.GetOk(isImageDeprecationAt); ok {
		deprecateAt, err := strfmt.ParseDateTime(deprecateAtOk.(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "parse-deprecation_at").GetDiag()
		}
		imagePrototype.DeprecationAt = &deprecateAt
	}
	if encryptionKey, ok := d.GetOk(isImageEncryptionKey); ok {
		encryptionKeyStr := encryptionKey.(string)
		// Construct an instance of the EncryptionKeyReference model
		encryptionKeyReferenceModel := new(vpcv1.EncryptionKeyIdentity)
		encryptionKeyReferenceModel.CRN = &encryptionKeyStr
		imagePrototype.EncryptionKey = encryptionKeyReferenceModel
	}
	if encDataKey, ok := d.GetOk(isImageEncryptedDataKey); ok {
		encDataKeyStr := encDataKey.(string)
		imagePrototype.EncryptedDataKey = &encDataKeyStr
	}
	if rgrp, ok := d.GetOk(isImageResourceGroup); ok {
		rg := rgrp.(string)
		imagePrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	options := &vpcv1.CreateImageOptions{
		ImagePrototype: imagePrototype,
	}
	image, _, err := sess.CreateImageWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateImageWithContext failed: %s", err.Error()), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*image.ID)
	log.Printf("[INFO] Image ID : %s", *image.ID)
	_, err = isWaitForImageAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForImageAvailable failed: %s", err.Error()), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isImageTags); ok || v != "" {
		oldList, newList := d.GetChange(isImageTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc Image (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isImageAccessTags); ok {
		oldList, newList := d.GetChange(isImageAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc Image (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}
func imgCreateByVolume(context context.Context, d *schema.ResourceData, meta interface{}, name, volume string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	imagePrototype := &vpcv1.ImagePrototypeImageBySourceVolume{
		Name: &name,
	}
	var insId string
	imagePrototype.SourceVolume = &vpcv1.VolumeIdentity{
		ID: &volume,
	}
	options := &vpcv1.GetVolumeOptions{
		ID: &volume,
	}
	vol, response, err := sess.GetVolumeWithContext(context, options)
	if err != nil || vol == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) == 0 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error creating Image because the specified source_volume %s is not attached to a virtual server instance", volume), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	volAtt := &vol.VolumeAttachments[0]
	if *volAtt.Type != "boot" {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error creating Image because the specified source_volume %s is not boot volume", volume), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	insId = *volAtt.Instance.ID
	getinsOptions := &vpcv1.GetInstanceOptions{
		ID: &insId,
	}
	instance, response, err := sess.GetInstanceWithContext(context, getinsOptions)
	if err != nil || instance == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error retrieving Instance (%s) to which the source_volume (%s) is attached : %s\n%s", insId, volume, err, response), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instance != nil && *instance.Status == "running" {
		actiontype := "stop"
		createinsactoptions := &vpcv1.CreateInstanceActionOptions{
			InstanceID: &insId,
			Type:       &actiontype,
		}
		_, response, err = sess.CreateInstanceActionWithContext(context, createinsactoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error stopping Instance (%s) to which the source_volume (%s) is attached  : %s\n%s", insId, volume, err, response), "ibm_is_image", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForInstanceActionStop(sess, d.Timeout(schema.TimeoutCreate), insId, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStop failed: %s", err.Error()), "ibm_is_image", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else if *instance.Status != "stopped" {
		_, err = isWaitForInstanceActionStop(sess, d.Timeout(schema.TimeoutCreate), insId, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStop failed: %s", err.Error()), "ibm_is_image", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if obsoleteAtOk, ok := d.GetOk(isImageObsolescenceAt); ok {
		obsoleteAt, err := strfmt.ParseDateTime(obsoleteAtOk.(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "parse-obsolescence_at").GetDiag()
		}
		imagePrototype.ObsolescenceAt = &obsoleteAt
	}
	if deprecateAtOk, ok := d.GetOk(isImageDeprecationAt); ok {
		deprecateAt, err := strfmt.ParseDateTime(deprecateAtOk.(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "create", "parse-deprecation_at").GetDiag()
		}
		imagePrototype.DeprecationAt = &deprecateAt
	}

	if encryptionKey, ok := d.GetOk(isImageEncryptionKey); ok {
		encryptionKeyStr := encryptionKey.(string)
		// Construct an instance of the EncryptionKeyReference model
		encryptionKeyReferenceModel := new(vpcv1.EncryptionKeyIdentity)
		encryptionKeyReferenceModel.CRN = &encryptionKeyStr
		imagePrototype.EncryptionKey = encryptionKeyReferenceModel
	}
	if rgrp, ok := d.GetOk(isImageResourceGroup); ok {
		rg := rgrp.(string)
		imagePrototype.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	imagOptions := &vpcv1.CreateImageOptions{
		ImagePrototype: imagePrototype,
	}
	image, response, err := sess.CreateImageWithContext(context, imagOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateImageWithContext failed: %s", err.Error()), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*image.ID)
	log.Printf("[INFO] Image ID : %s", *image.ID)
	_, err = isWaitForImageAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForImageAvailable failed: %s", err.Error()), "ibm_is_image", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isImageTags); ok || v != "" {
		oldList, newList := d.GetChange(isImageTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc Image (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk(isImageAccessTags); ok {
		oldList, newList := d.GetChange(isImageAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc Image (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func isWaitForImageAvailable(imageC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for image (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isImageProvisioning},
		Target:     []string{isImageProvisioningDone, ""},
		Refresh:    isImageRefreshFunc(imageC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
func isImageRefreshFunc(imageC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getimgoptions := &vpcv1.GetImageOptions{
			ID: &id,
		}
		image, response, err := imageC.GetImage(getimgoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Image: %s\n%s", err, response)
		}

		if *image.Status == "available" || *image.Status == "failed" {
			return image, isImageProvisioningDone, nil
		}

		return image, isImageProvisioning, nil
	}
}

func resourceIBMISImageUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	name := ""
	hasChanged := false

	if d.HasChange(isImageName) {
		name = d.Get(isImageName).(string)
		hasChanged = true
	}
	err := imgUpdate(context, d, meta, id, name, hasChanged)
	if err != nil {
		return err
	}

	return resourceIBMISImageRead(context, d, meta)
}

func imgUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, name string, hasNameChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isImageDeprecate) && !d.IsNewResource() {
		deprecateTrue := d.Get(isImageDeprecate).(bool)
		if deprecateTrue {
			deprecateImageOptions := &vpcv1.DeprecateImageOptions{
				ID: &id,
			}
			_, err := sess.DeprecateImageWithContext(context, deprecateImageOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeprecateImageWithContext failed: %s", err.Error()), "ibm_is_image", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForImageDeprecate(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForImageDeprecate failed: %s", err.Error()), "ibm_is_image", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}
	if d.HasChange(isImageObsolete) && !d.IsNewResource() {
		obsoleteTrue := d.Get(isImageObsolete).(bool)
		if obsoleteTrue {
			obsoleteImageOptions := &vpcv1.ObsoleteImageOptions{
				ID: &id,
			}
			_, err := sess.ObsoleteImageWithContext(context, obsoleteImageOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ObsoleteImageWithContext failed: %s", err.Error()), "ibm_is_image", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForImageObsolete(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForImageObsolete failed: %s", err.Error()), "ibm_is_image", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}
	if d.HasChange(isImageTags) {
		options := &vpcv1.GetImageOptions{
			ID: &id,
		}
		image, _, err := sess.GetImageWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetImageWithContext failed: %s", err.Error()), "ibm_is_image", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isImageTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Image (%s) tags: %s", id, err)
		}
	}
	if d.HasChange(isImageAccessTags) {
		options := &vpcv1.GetImageOptions{
			ID: &id,
		}
		image, _, err := sess.GetImageWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetImageWithContext failed: %s", err.Error()), "ibm_is_image", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isImageAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *image.CRN, "", isImageAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc Image (%s) access tags: %s", id, err)
		}
	}

	options := &vpcv1.UpdateImageOptions{
		ID: &id,
	}
	imagePatchModel := &vpcv1.ImagePatch{}
	if hasNameChanged {
		imagePatchModel.Name = &name
	}
	nullObsolescence := false
	nullDeprecate := false
	if d.HasChange(isImageObsolescenceAt) {
		obsolescenceAt := d.Get(isImageObsolescenceAt).(string)
		if obsolescenceAt == "null" {
			nullObsolescence = true
		} else {
			obsoleteAt, err := strfmt.ParseDateTime(obsolescenceAt)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "update", "parse-obsolescence_at").GetDiag()
			}
			imagePatchModel.ObsolescenceAt = &obsoleteAt
		}
	}

	if d.HasChange(isImageDeprecationAt) {
		deprecationAt := d.Get(isImageDeprecationAt).(string)
		if deprecationAt == "null" {
			nullDeprecate = true
		} else {
			deprecateAt, err := strfmt.ParseDateTime(deprecationAt)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "update", "parse-deprecation_at").GetDiag()
			}
			imagePatchModel.DeprecationAt = &deprecateAt
		}
	}
	imagePatch, err := imagePatchModel.AsPatch()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("imagePatchModel.AsPatch failed: %s", err.Error()), "ibm_is_image", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if nullDeprecate {
		imagePatch["deprecation_at"] = nil
	}
	if nullObsolescence {
		imagePatch["obsolescence_at"] = nil
	}
	options.ImagePatch = imagePatch
	_, _, err = sess.UpdateImage(options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateImageWithContext failed: %s", err.Error()), "ibm_is_image", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func resourceIBMISImageRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	err := imgGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func imgGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetImageOptions{
		ID: &id,
	}
	image, response, err := sess.GetImageWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetImageWithContext failed: %s", err.Error()), "ibm_is_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// d.Set(isImageArchitecure, image.Architecture)
	if !core.IsNil(image.MinimumProvisionedSize) {
		if err = d.Set("size", flex.IntValue(image.MinimumProvisionedSize)); err != nil {
			err = fmt.Errorf("Error setting size: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-size").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(image.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(image.DeprecationAt) {
		if err = d.Set("deprecation_at", flex.DateTimeToString(image.DeprecationAt)); err != nil {
			err = fmt.Errorf("Error setting deprecation_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-deprecation_at").GetDiag()
		}
	}
	if !core.IsNil(image.ObsolescenceAt) {
		if err = d.Set("obsolescence_at", flex.DateTimeToString(image.ObsolescenceAt)); err != nil {
			err = fmt.Errorf("Error setting obsolescence_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-obsolescence_at").GetDiag()
		}
	}
	if !core.IsNil(image.Name) {
		if err = d.Set("name", image.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(image.OperatingSystem) {
		if err = d.Set("operating_system", image.OperatingSystem.Name); err != nil {
			err = fmt.Errorf("Error setting operating_system: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-operating_system").GetDiag()
		}
	}
	// d.Set(isImageFormat, image.Format)
	if err = d.Set("encryption", image.Encryption); err != nil {
		err = fmt.Errorf("Error setting encryption: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-encryption").GetDiag()
	}
	if image.EncryptionKey != nil {
		if err = d.Set("encryption_key", *image.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-encryption_key").GetDiag()
		}
	}
	if image.File != nil && image.File.Size != nil {
		if err = d.Set("file", *image.File.Size); err != nil {
			err = fmt.Errorf("Error setting file: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-file").GetDiag()
		}
	}
	if image.SourceVolume != nil {
		if err = d.Set("source_volume", *image.SourceVolume.ID); err != nil {
			err = fmt.Errorf("Error setting source_volume: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-source_volume").GetDiag()
		}
	}

	if err = d.Set("href", image.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-href").GetDiag()
	}
	if err = d.Set("status", image.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-status").GetDiag()
	}
	if err = d.Set("visibility", image.Visibility); err != nil {
		err = fmt.Errorf("Error setting visibility: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-visibility").GetDiag()
	}
	if image.UserDataFormat != nil {
		if err = d.Set("user_data_format", image.UserDataFormat); err != nil {
			err = fmt.Errorf("Error setting user_data_format: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-user_data_format").GetDiag()
		}
	}
	if err = d.Set("encryption", image.Encryption); err != nil {
		err = fmt.Errorf("Error setting encryption: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-encryption").GetDiag()
	}
	if image.EncryptionKey != nil {
		if err = d.Set("encryption_key", *image.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-encryption_key").GetDiag()
		}
	}
	if image.File != nil && image.File.Checksums != nil {
		if err = d.Set(isImageCheckSum, *image.File.Checksums.Sha256); err != nil {
			err = fmt.Errorf("Error setting checksum: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-checksum").GetDiag()
		}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *image.CRN, "", isImageUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc Image (%s) tags: %s", d.Id(), err)
	}
	d.Set(isImageTags, tags)
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *image.CRN, "", isImageAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc Image (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set("access_tags", accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-access_tags").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_image", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/compute/image")
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/compute/image"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-resource_controller_url").GetDiag()
	}
	d.Set(flex.ResourceName, *image.Name)
	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/compute/image"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceStatus, *image.Status); err != nil {
		err = fmt.Errorf("Error setting resource_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-resource_status").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *image.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set(IsImageCRN, *image.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-crn").GetDiag()
	}
	if image.ResourceGroup != nil {
		if err = d.Set(isImageResourceGroup, *image.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "read", "set-resource_group").GetDiag()
		}
	}
	return nil
}

func resourceIBMISImageDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	err := imgDelete(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func imgDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getImageOptions := &vpcv1.GetImageOptions{
		ID: &id,
	}
	_, response, err := sess.GetImageWithContext(context, getImageOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetImageWithContext failed: %s", err.Error()), "ibm_is_image", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.DeleteImageOptions{
		ID: &id,
	}
	response, err = sess.DeleteImageWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteImageWithContext failed: %s", err.Error()), "ibm_is_image", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForImageDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForImageDeleted failed: %s", err.Error()), "ibm_is_image", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForImageDeleted(imageC *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for image (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isImageDeleting},
		Target:     []string{"", isImageDeleted},
		Refresh:    isImageDeleteRefreshFunc(imageC, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isImageDeleteRefreshFunc(imageC *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] is image delete function here")
		getimgoptions := &vpcv1.GetImageOptions{
			ID: &id,
		}
		image, response, err := imageC.GetImage(getimgoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return image, isImageDeleted, nil
			}
			return image, "", fmt.Errorf("[ERROR] Error Getting Image: %s\n%s", err, response)
		}
		return image, isImageDeleting, err
	}
}
func resourceIBMISImageExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	exists, err := imgExists(d, meta, id)
	return exists, err
}

func imgExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_image", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	options := &vpcv1.GetImageOptions{
		ID: &id,
	}
	_, response, err := sess.GetImage(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Image: %s\n%s", err, response)
	}
	return true, nil
}
