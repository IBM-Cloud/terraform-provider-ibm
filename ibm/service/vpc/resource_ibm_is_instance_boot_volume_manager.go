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
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceBootVolumeManagerDelete = "delete"
)

func ResourceIBMISInstanceBootVolumeManager() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISInstanceBootVolumeManagerCreate,
		Read:     resourceIBMISInstanceBootVolumeManagerRead,
		Update:   resourceIBMISInstanceBootVolumeManagerUpdate,
		Delete:   resourceIBMISInstanceBootVolumeManagerDelete,
		Exists:   resourceIBMISInstanceBootVolumeManagerExists,
		Importer: &schema.ResourceImporter{},

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
					return flex.ResourceVolumeValidate(diff)
				}),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{

			isInstanceBootVolumeId: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeName),
				Description:  "Volume name",
			},
			isVolumeName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeName),
				Description:  "Volume name",
			},

			isVolumeProfileName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeProfileName),
				Description:  "Volume profile name",
			},

			isVolumeZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			isVolumeEncryptionKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume encryption key info",
			},

			isInstanceBootVolumeManagerDelete: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Volume encryption key info",
			},

			isVolumeEncryptionType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume encryption type info",
			},

			isVolumeCapacity: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeCapacity),
				Description:  "Volume capacity value",
			},
			isVolumeSourceSnapshot: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this snapshot",
			},
			isVolumeResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group name",
			},
			isVolumeIops: {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", isVolumeIops),
				Description:  "IOPS value for the Volume",
			},
			isVolumeCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN value for the volume instance",
			},
			isVolumeStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume status",
			},

			isVolumeStatusReasons: {
				Type:     schema.TypeList,
				Computed: true,
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
			isVolumeHealthReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeHealthReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},

						isVolumeHealthReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},

						isVolumeHealthReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},

			isVolumeHealthState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.",
			},
			isVolumeDeleteAllSnapshots: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Deletes all snapshots created from this volume",
			},
			isVolumeTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "UserTags for the volume instance",
			},
			isVolumeAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_instance_boot_volume_manager", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Access management tags for the volume instance",
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

			isVolumeBandwidth: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume",
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
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceBootVolumeId,
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
			MaxValue:                   "250"})
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

	ibmISVolumeResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_boot_volume_manager", Schema: validateSchema}
	return &ibmISVolumeResourceValidator
}

func resourceIBMISInstanceBootVolumeManagerCreate(d *schema.ResourceData, meta interface{}) error {

	volId := d.Get(isInstanceBootVolumeId).(string)
	d.SetId(volId)
	err := resourceIBMISInstanceBootVolumeManagerRead(d, meta)
	if err != nil {
		return err
	}

	return resourceIBMISInstanceBootVolumeManagerUpdate(d, meta)
}

func resourceIBMISInstanceBootVolumeManagerRead(d *schema.ResourceData, meta interface{}) error {

	id := d.Id()
	err := instancebootvolGet(d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func instancebootvolGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	vol, response, err := sess.GetVolume(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error getting Instance boot volume (%s): %s\n%s", id, err, response)
	}
	d.SetId(*vol.ID)
	d.Set(isVolumeName, *vol.Name)
	d.Set(isVolumeProfileName, *vol.Profile.Name)
	d.Set(isVolumeZone, *vol.Zone.Name)
	if vol.EncryptionKey != nil {
		d.Set(isVolumeEncryptionKey, vol.EncryptionKey.CRN)
	}
	if vol.Encryption != nil {
		d.Set(isVolumeEncryptionType, vol.Encryption)
	}
	d.Set(isVolumeIops, *vol.Iops)
	d.Set(isVolumeCapacity, *vol.Capacity)
	d.Set(isVolumeCrn, *vol.CRN)
	if vol.SourceSnapshot != nil {
		d.Set(isVolumeSourceSnapshot, *vol.SourceSnapshot.ID)
	}
	d.Set(isVolumeStatus, *vol.Status)
	if vol.HealthState != nil {
		d.Set(isVolumeHealthState, *vol.HealthState)
	}
	d.Set(isVolumeBandwidth, int(*vol.Bandwidth))
	//set the status reasons
	if vol.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range vol.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isVolumeStatusReasonsCode] = *sr.Code
				currentSR[isVolumeStatusReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isVolumeStatusReasonsMoreInfo] = *sr.Message
				}
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
		d.Set(isVolumeStatusReasons, statusReasonsList)
	}
	if vol.UserTags != nil {
		if err = d.Set(isVolumeTags, vol.UserTags); err != nil {
			return fmt.Errorf("Error setting user tags: %s", err)
		}
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *vol.CRN, "", isVolumeAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource volume (%s) access tags: %s", d.Id(), err)
	}
	d.Set(isVolumeAccessTags, accesstags)
	if vol.HealthReasons != nil {
		healthReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range vol.HealthReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isVolumeHealthReasonsCode] = *sr.Code
				currentSR[isVolumeHealthReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isVolumeHealthReasonsMoreInfo] = *sr.Message
				}
				healthReasonsList = append(healthReasonsList, currentSR)
			}
		}
		d.Set(isVolumeHealthReasons, healthReasonsList)
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/storage/storageVolumes")
	d.Set(flex.ResourceName, *vol.Name)
	d.Set(flex.ResourceCRN, *vol.CRN)
	d.Set(flex.ResourceStatus, *vol.Status)
	if vol.ResourceGroup != nil {
		d.Set(flex.ResourceGroupName, vol.ResourceGroup.Name)
		d.Set(isVolumeResourceGroup, *vol.ResourceGroup.ID)
	}
	return nil
}

func resourceIBMISInstanceBootVolumeManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	err := instancebootvolUpdate(d, meta)
	if err != nil {
		return err
	}
	return resourceIBMISInstanceBootVolumeManagerRead(d, meta)
}

func instancebootvolUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id := d.Id()
	name := ""
	hasNameChanged := false
	delete := false

	if delete_all_snapshots, ok := d.GetOk(isVolumeDeleteAllSnapshots); ok && delete_all_snapshots.(bool) {
		delete = true
	}

	if d.HasChange(isVolumeName) {
		name = d.Get(isVolumeName).(string)
		hasNameChanged = true
	}
	var capacity int64
	if delete {
		deleteAllInstanceBootSnapshots(sess, id)
	}

	if d.HasChange(isVolumeAccessTags) {
		options := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := sess.GetVolume(options)
		if err != nil {
			return fmt.Errorf("[ERROR]Error getting Instance boot volume : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isVolumeAccessTags)

		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vol.CRN, "", isVolumeAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource Instance boot volume (%s) access tags: %s", id, err)
		}
	}

	optionsget := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	_, response, err := sess.GetVolume(optionsget)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error getting Instance boot volume (%s): %s\n%s", id, err, response)
	}
	eTag := response.Headers.Get("ETag")
	options := &vpcv1.UpdateVolumeOptions{
		ID: &id,
	}
	options.IfMatch = &eTag

	//name update
	volumeNamePatchModel := &vpcv1.VolumePatch{}
	if hasNameChanged {
		volumeNamePatchModel.Name = &name
		volumeNamePatch, err := volumeNamePatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("[ERROR] Error calling asPatch for volumeNamePatch in Instance boot volume : %s", err)
		}
		options.VolumePatch = volumeNamePatch
		_, response, err = sess.UpdateVolume(options)
		eTag = response.Headers.Get("ETag")
		if err != nil {
			return err
		}
		_, err = isWaitForInstanceBootVolumeManagerAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), &eTag)
		if err != nil {
			return err
		}
	}

	// profile/ iops update
	if d.HasChange(isVolumeProfileName) || d.HasChange(isVolumeIops) {
		volumeProfilePatchModel := &vpcv1.VolumePatch{}
		volId := d.Id()
		getvoloptions := &vpcv1.GetVolumeOptions{
			ID: &volId,
		}
		vol, response, err := sess.GetVolume(getvoloptions)
		if err != nil || vol == nil {
			return fmt.Errorf("[ERROR] Error retrieving Instance boot volume (%s) details: %s\n%s", volId, err, response)
		}
		if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) < 1 {
			return fmt.Errorf("[ERROR] Error updating Instance boot volume profile/iops because the specified volume %s is not attached to a virtual server instance ", volId)
		}
		volAtt := &vol.VolumeAttachments[0]
		insId := *volAtt.Instance.ID
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &insId,
		}
		instance, response, err := sess.GetInstance(getinsOptions)
		if err != nil || instance == nil {
			return fmt.Errorf("[ERROR] Error retrieving Instance (%s) to which the boot volume (%s) is attached : %s\n%s", insId, volId, err, response)
		}
		if instance != nil && *instance.Status != "running" {
			actiontype := "start"
			createinsactoptions := &vpcv1.CreateInstanceActionOptions{
				InstanceID: &insId,
				Type:       &actiontype,
			}
			_, response, err = sess.CreateInstanceAction(createinsactoptions)
			if err != nil {
				return fmt.Errorf("[ERROR] Error starting Instance (%s) to which the boot volume (%s) is attached  : %s\n%s", insId, volId, err, response)
			}
			_, err = isWaitForInstanceAvailable(sess, insId, d.Timeout(schema.TimeoutCreate), d)
			if err != nil {
				return err
			}
		}
		if d.HasChange(isVolumeProfileName) {
			profile := d.Get(isVolumeProfileName).(string)
			volumeProfilePatchModel.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &profile,
			}
		} else if d.HasChange(isVolumeIops) {
			profile := d.Get(isVolumeProfileName).(string)
			volumeProfilePatchModel.Profile = &vpcv1.VolumeProfileIdentity{
				Name: &profile,
			}
			iops := int64(d.Get(isVolumeIops).(int))
			volumeProfilePatchModel.Iops = &iops
		}

		volumeProfilePatch, err := volumeProfilePatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("[ERROR] Error calling asPatch for VolumeProfilePatch in Instance boot volume : %s", err)
		}
		options.VolumePatch = volumeProfilePatch
		_, response, err = sess.UpdateVolume(options)
		_, err = isWaitForInstanceBootVolumeManagerAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), &eTag)
		if err != nil {
			return err
		}
	}

	// capacity update
	if d.HasChange(isVolumeCapacity) {
		id := d.Id()
		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := sess.GetVolume(getvolumeoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Error Getting Instance boot volume (%s): %s\n%s", id, err, response)
		}
		if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) == 0 || *vol.VolumeAttachments[0].ID == "" {
			return fmt.Errorf("[ERROR] Error volume capacity can't be updated since Instance boot volume %s is not attached to any instance for VolumePatch", id)
		}
		insId := vol.VolumeAttachments[0].Instance.ID
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: insId,
		}
		instance, response, err := sess.GetInstance(getinsOptions)
		if err != nil || instance == nil {
			return fmt.Errorf("[ERROR] Error retrieving Instance (%s) : %s\n%s", *insId, err, response)
		}
		if instance != nil && *instance.Status != "running" {
			actiontype := "start"
			createinsactoptions := &vpcv1.CreateInstanceActionOptions{
				InstanceID: insId,
				Type:       &actiontype,
			}
			_, response, err = sess.CreateInstanceAction(createinsactoptions)
			if err != nil {
				return fmt.Errorf("[ERROR] Error starting Instance (%s) : %s\n%s", *insId, err, response)
			}
			_, err = isWaitForInstanceAvailable(sess, *insId, d.Timeout(schema.TimeoutCreate), d)
			if err != nil {
				return err
			}
		}
		capacity = int64(d.Get(isVolumeCapacity).(int))
		volumeCapacityPatchModel := &vpcv1.VolumePatch{}
		volumeCapacityPatchModel.Capacity = &capacity

		volumeCapacityPatch, err := volumeCapacityPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("[ERROR] Error calling asPatch for volumeCapacityPatch in Instance boot volume : %s", err)
		}
		options.VolumePatch = volumeCapacityPatch
		_, response, err = sess.UpdateVolume(options)
		if err != nil {
			return fmt.Errorf("[ERROR] Error updating Instance boot volume: %s\n%s", err, response)
		}
		_, err = isWaitForInstanceBootVolumeManagerAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), &eTag)
		if err != nil {
			return err
		}
	}

	// user tags update
	if d.HasChange(isVolumeTags) {
		var userTags *schema.Set
		if v, ok := d.GetOk(isVolumeTags); ok {
			userTags = v.(*schema.Set)
			if userTags != nil && userTags.Len() != 0 {
				userTagsArray := make([]string, userTags.Len())
				for i, userTag := range userTags.List() {
					userTagStr := userTag.(string)
					userTagsArray[i] = userTagStr
				}
				schematicTags := os.Getenv("IC_ENV_TAGS")
				var envTags []string
				if schematicTags != "" {
					envTags = strings.Split(schematicTags, ",")
					userTagsArray = append(userTagsArray, envTags...)
				}
				volumeNamePatchModel := &vpcv1.VolumePatch{}
				volumeNamePatchModel.UserTags = userTagsArray
				volumeNamePatch, err := volumeNamePatchModel.AsPatch()
				if err != nil {
					return fmt.Errorf("[ERROR] Error calling asPatch for volumeNamePatch in Instance boot volume: %s", err)
				}
				options.IfMatch = &eTag
				options.VolumePatch = volumeNamePatch
				_, response, err := sess.UpdateVolume(options)
				if err != nil {
					return fmt.Errorf("[ERROR] Error updating Instance boot volume : %s\n%s", err, response)
				}
				_, err = isWaitForInstanceBootVolumeManagerAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate), &eTag)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func resourceIBMISInstanceBootVolumeManagerDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	// check if force delete is true
	if d.Get(isInstanceBootVolumeManagerDelete).(bool) {
		sess, err := vpcClient(meta)
		if err != nil {
			return err
		}

		options := &vpcv1.DeleteVolumeOptions{
			ID: &id,
		}
		response, err := sess.DeleteVolume(options)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting Volume : %s\n%s", err, response)
		}
		_, err = isWaitForInstanceBootVolumeManagerDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return err
		}
	}
	d.SetId("")
	return nil
}

func isWaitForInstanceBootVolumeManagerDeleted(vol *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeDeleting},
		Target:     []string{"done", ""},
		Refresh:    isInstanceBootVolumeManagerDeleteRefreshFunc(vol, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceBootVolumeManagerDeleteRefreshFunc(vol *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		volgetoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := vol.GetVolume(volgetoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vol, isVolumeDeleted, nil
			}
			return vol, "", fmt.Errorf("[ERROR] Error getting Instance boot volume: %s\n%s", err, response)
		}
		return vol, isVolumeDeleting, err
	}
}

func resourceIBMISInstanceBootVolumeManagerExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()

	exists, err := instancebootvolExists(d, meta, id)
	return exists, err
}

func instancebootvolExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	options := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	_, response, err := sess.GetVolume(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Instance boot volume: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForInstanceBootVolumeManagerAvailable(client *vpcv1.VpcV1, id string, timeout time.Duration, eTag *string) (interface{}, error) {
	log.Printf("Waiting for Instance boot volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeProvisioning},
		Target:     []string{isVolumeProvisioningDone, ""},
		Refresh:    isInstanceBootVolumeManagerRefreshFunc(client, id, eTag),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isInstanceBootVolumeManagerRefreshFunc(client *vpcv1.VpcV1, id string, eTag *string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		volgetoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := client.GetVolume(volgetoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting Instance boot volume: %s\n%s", err, response)
		}

		if *vol.Status == "available" {
			return vol, isVolumeProvisioningDone, nil
		}
		*eTag = response.Headers.Get("ETag")

		return vol, isVolumeProvisioning, nil
	}
}

func deleteAllInstanceBootSnapshots(sess *vpcv1.VpcV1, id string) error {
	delete_all_snapshots := new(vpcv1.DeleteSnapshotsOptions)
	delete_all_snapshots.SourceVolumeID = &id
	response, err := sess.DeleteSnapshots(delete_all_snapshots)
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting snapshots from Instance boot volume %s\n%s", err, response)
	}
	return nil
}
