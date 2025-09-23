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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isVolumeName                  = "name"
	isVolumeProfileName           = "profile"
	isVolumeZone                  = "zone"
	isVolumeEncryptionKey         = "encryption_key"
	isVolumeEncryptionType        = "encryption_type"
	isVolumeCapacity              = "capacity"
	isVolumeIops                  = "iops"
	isVolumeCrn                   = "crn"
	isVolumeTags                  = "tags"
	isVolumeStatus                = "status"
	isVolumeStatusReasons         = "status_reasons"
	isVolumeStatusReasonsCode     = "code"
	isVolumeStatusReasonsMessage  = "message"
	isVolumeStatusReasonsMoreInfo = "more_info"
	isVolumeDeleting              = "deleting"
	isVolumeDeleted               = "done"
	isVolumeProvisioning          = "provisioning"
	isVolumeProvisioningDone      = "done"
	isVolumeResourceGroup         = "resource_group"
	isVolumeSourceSnapshot        = "source_snapshot"
	isVolumeSourceSnapshotCrn     = "source_snapshot_crn"
	isVolumeDeleteAllSnapshots    = "delete_all_snapshots"
	isVolumeBandwidth             = "bandwidth"
	isVolumeAccessTags            = "access_tags"
	isVolumeUserTagType           = "user"
	isVolumeAccessTagType         = "access"
	isVolumeHealthReasons         = "health_reasons"
	isVolumeHealthReasonsCode     = "code"
	isVolumeHealthReasonsMessage  = "message"
	isVolumeHealthReasonsMoreInfo = "more_info"
	isVolumeHealthState           = "health_state"

	isVolumeCatalogOffering           = "catalog_offering"
	isVolumeCatalogOfferingPlanCrn    = "plan_crn"
	isVolumeCatalogOfferingVersionCrn = "version_crn"
)

func ResourceIBMISVolume() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVolumeCreate,
		ReadContext:   resourceIBMISVolumeRead,
		UpdateContext: resourceIBMISVolumeUpdate,
		DeleteContext: resourceIBMISVolumeDelete,
		Exists:        resourceIBMISVolumeExists,
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
					return flex.ResourceVolumeValidate(diff)
				}),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{

			isVolumeName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume", isVolumeName),
				Description:  "Volume name",
			},

			isVolumeProfileName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume", isVolumeProfileName),
				Description:  "Volume profile name",
			},

			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The maximum bandwidth (in megabits per second) for the volume. For this property to be specified, the volume storage_generation must be 2.",
			},
			isVolumeZone: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone name",
			},

			isVolumeEncryptionKey: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Volume encryption key info",
			},

			isVolumeEncryptionType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Volume encryption type info",
			},
			"storage_generation": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "storage_generation indicates which generation the profile family belongs to. For the custom and tiered profiles, this value is 1.",
			},

			isVolumeCapacity: {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Computed: true,
				// ValidateFunc: validate.InvokeValidator("ibm_is_volume", isVolumeCapacity),
				Description: "Volume capacity value",
			},
			isVolumeSourceSnapshot: {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{isVolumeSourceSnapshotCrn},
				ValidateFunc:  validate.InvokeValidator("ibm_is_volume", isVolumeSourceSnapshot),
				Description:   "The unique identifier for this snapshot",
			},
			isVolumeSourceSnapshotCrn: {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{isVolumeSourceSnapshot},
				Description:   "The crn for this snapshot",
			},
			isVolumeResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource group name",
			},
			isVolumeIops: {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// ValidateFunc: validate.InvokeValidator("ibm_is_volume", isVolumeIops),
				Description: "IOPS value for the Volume",
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
			isVolumeCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering this volume was created from. If a virtual server instance is provisioned with a boot_volume_attachment specifying this volume, the virtual server instance will use this volume's catalog offering, including its pricing plan.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeCatalogOfferingPlanCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this catalog offering version's billing plan",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						isVolumeCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this version of a catalog offering",
						},
					},
				},
			},
			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The usage constraints to match against the requested instance or bare metal server properties to determine compatibility.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_version": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_volume", "allowed_use.api_version"),
							Description:  "The API version with which to evaluate the expressions.",
						},
						"bare_metal_server": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_volume", "allowed_use.bare_metal_server"),
							Description:  "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume.",
						},
						"instance": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_volume", "allowed_use.instance"),
							Description:  "The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume.",
						},
					},
				},
			},
			// defined_performance changes
			"adjustable_capacity_states": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attachment states that support adjustable capacity for this volume.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"adjustable_iops_states": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attachment states that support adjustable IOPS for this volume.",
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_volume", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "UserTags for the volume instance",
			},
			isVolumeAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_volume", "accesstag")},
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

			isVolumesOperatingSystem: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The operating system associated with this volume. If absent, this volume was notcreated from an image, or the image did not include an operating system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVolumeArchitecture: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operating system architecture.",
						},
						isVolumeDHOnly: &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Images with this operating system can only be used on dedicated hosts or dedicated host groups.",
						},
						isVolumeDisplayName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique, display-friendly name for the operating system.",
						},
						isVolumeOSFamily: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The software family for this operating system.",
						},

						isVolumesOperatingSystemHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this operating system.",
						},
						isVolumesOperatingSystemName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this operating system.",
						},
						isVolumeOSVendor: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vendor of the operating system.",
						},
						isVolumeOSVersion: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The major release version of this operating system.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISVolumeValidator() *validate.ResourceValidator {

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
			Identifier:                 isVolumeSourceSnapshot,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})
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
			AllowedValues:              "general-purpose, 5iops-tier, 10iops-tier, custom, sdp",
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeCapacity,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "10"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isVolumeIops,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "100"})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "allowed_use.api_version",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "allowed_use.bare_metal_server",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$`})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "allowed_use.instance",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$`})

	ibmISVolumeResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_volume", Schema: validateSchema}
	return &ibmISVolumeResourceValidator
}

func resourceIBMISVolumeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	volName := d.Get(isVolumeName).(string)
	profile := d.Get(isVolumeProfileName).(string)
	zone := d.Get(isVolumeZone).(string)

	err := volCreate(context, d, meta, volName, profile, zone)
	if err != nil {
		return err
	}

	return resourceIBMISVolumeRead(context, d, meta)
}

func volCreate(context context.Context, d *schema.ResourceData, meta interface{}, volName, profile, zone string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateVolumeOptions{}
	volTemplate := &vpcv1.VolumePrototype{
		Name: &volName,
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
		Profile: &vpcv1.VolumeProfileIdentity{
			Name: &profile,
		},
	}

	if sourceSnapsht, ok := d.GetOk(isVolumeSourceSnapshot); ok {
		sourceSnapshot := sourceSnapsht.(string)
		snapshotIdentity := &vpcv1.SnapshotIdentity{
			ID: &sourceSnapshot,
		}
		volTemplate.SourceSnapshot = snapshotIdentity
		if capacity, ok := d.GetOk(isVolumeCapacity); ok {
			if int64(capacity.(int)) > 0 {
				volCapacity := int64(capacity.(int))
				volTemplate.Capacity = &volCapacity
			}
		}
	} else if sourceSnapshtCrn, ok := d.GetOk(isVolumeSourceSnapshotCrn); ok {
		sourceSnapshot := sourceSnapshtCrn.(string)

		snapshotIdentity := &vpcv1.SnapshotIdentity{
			CRN: &sourceSnapshot,
		}
		volTemplate.SourceSnapshot = snapshotIdentity
		if capacity, ok := d.GetOk(isVolumeCapacity); ok {
			if int64(capacity.(int)) > 0 {
				volCapacity := int64(capacity.(int))
				volTemplate.Capacity = &volCapacity
			}
		}
	} else if capacity, ok := d.GetOk(isVolumeCapacity); ok {
		if int64(capacity.(int)) > 0 {
			volCapacity := int64(capacity.(int))
			volTemplate.Capacity = &volCapacity
		}
	} else {
		volCapacity := int64(100)
		volTemplate.Capacity = &volCapacity
	}

	if key, ok := d.GetOk(isVolumeEncryptionKey); ok {
		encryptionKey := key.(string)
		volTemplate.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
			CRN: &encryptionKey,
		}
	}

	if rgrp, ok := d.GetOk(isVolumeResourceGroup); ok {
		rg := rgrp.(string)
		volTemplate.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	if i, ok := d.GetOk(isVolumeIops); ok {
		iops := int64(i.(int))
		volTemplate.Iops = &iops
	}
	// bandwidth changes
	if b, ok := d.GetOk("bandwidth"); ok {
		bandwidth := int64(b.(int))
		volTemplate.Bandwidth = &bandwidth
	}

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
			volTemplate.UserTags = userTagsArray
		}
	}
	if allowedUse, ok := d.GetOk("allowed_use"); ok && len(allowedUse.([]interface{})) > 0 {
		allowedUseModel, _ := ResourceIBMIsVolumeAllowedUseMapToVolumeAllowedUsePrototype(allowedUse.([]interface{})[0].(map[string]interface{}))
		volTemplate.AllowedUse = allowedUseModel
	}

	options.VolumePrototype = volTemplate
	vol, _, err := sess.CreateVolumeWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*vol.ID)
	log.Printf("[INFO] Volume : %s", *vol.ID)
	_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := d.GetOk(isVolumeAccessTags); ok {
		oldList, newList := d.GetChange(isVolumeAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vol.CRN, "", isVolumeAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vpc volume (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISVolumeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	err := volGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func volGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	volume, response, err := sess.GetVolumeWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*volume.ID)
	if !core.IsNil(volume.Name) {
		if err = d.Set("name", volume.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("storage_generation", flex.IntValue(volume.StorageGeneration)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_generation: %s", err), "(Data) ibm_is_volume", "read", "set-storage_generation").GetDiag()
	}
	if !core.IsNil(volume.Profile) {
		if err = d.Set("profile", *volume.Profile.Name); err != nil {
			err = fmt.Errorf("Error setting profile: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-profile").GetDiag()
		}
	}
	if !core.IsNil(volume.Zone) {
		if err = d.Set(isVolumeZone, *volume.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-zone").GetDiag()
		}
	}
	if volume.EncryptionKey != nil {
		if err = d.Set(isVolumeEncryptionKey, volume.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-encryption_key").GetDiag()
		}
	}
	if err = d.Set("encryption_type", volume.Encryption); err != nil {
		err = fmt.Errorf("Error setting encryption_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-encryption_type").GetDiag()
	}
	if !core.IsNil(volume.Iops) {
		if err = d.Set("iops", flex.IntValue(volume.Iops)); err != nil {
			err = fmt.Errorf("Error setting iops: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-iops").GetDiag()
		}
	}
	if !core.IsNil(volume.Capacity) {
		if err = d.Set("capacity", flex.IntValue(volume.Capacity)); err != nil {
			err = fmt.Errorf("Error setting capacity: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-capacity").GetDiag()
		}
	}
	if err = d.Set("crn", volume.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-crn").GetDiag()
	}
	if volume.SourceSnapshot != nil {
		if err = d.Set(isVolumeSourceSnapshot, *volume.SourceSnapshot.ID); err != nil {
			err = fmt.Errorf("Error setting source_snapshot: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-source_snapshot").GetDiag()
		}
		if err = d.Set(isVolumeSourceSnapshotCrn, *volume.SourceSnapshot.CRN); err != nil {
			err = fmt.Errorf("Error setting source_snapshot_crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-source_snapshot_crn").GetDiag()
		}
	}
	if err = d.Set("status", volume.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-status").GetDiag()
	}
	if err = d.Set("health_state", volume.HealthState); err != nil {
		err = fmt.Errorf("Error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-health_state").GetDiag()
	}
	if !core.IsNil(volume.Bandwidth) {
		if err = d.Set("bandwidth", flex.IntValue(volume.Bandwidth)); err != nil {
			err = fmt.Errorf("Error setting bandwidth: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-bandwidth").GetDiag()
		}
	}
	//set the status reasons
	if volume.StatusReasons != nil {
		statusReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range volume.StatusReasons {
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
		if err = d.Set(isVolumeStatusReasons, statusReasonsList); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-status_reasons").GetDiag()
		}
	}
	if volume.UserTags != nil {
		if err = d.Set(isVolumeTags, volume.UserTags); err != nil {
			err = fmt.Errorf("Error setting tags: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-tags").GetDiag()
		}
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *volume.CRN, "", isVolumeAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource volume (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isVolumeAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-access_tags").GetDiag()
	}
	if volume.HealthReasons != nil {
		healthReasonsList := make([]map[string]interface{}, 0)
		for _, sr := range volume.HealthReasons {
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
		if err = d.Set(isVolumeHealthReasons, healthReasonsList); err != nil {
			err = fmt.Errorf("Error setting health_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-health_reasons").GetDiag()
		}
	}
	allowedUses := []map[string]interface{}{}
	if volume.AllowedUse != nil {
		modelMap, err := ResourceceIBMIsVolumeAllowedUseToMap(volume.AllowedUse)
		if err != nil {
			err = fmt.Errorf("Error setting allowed_use: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-allowed_use").GetDiag()
		}
		allowedUses = append(allowedUses, modelMap)
	}
	if err = d.Set("allowed_use", allowedUses); err != nil {
		err = fmt.Errorf("Error setting allowed_use: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-allowed_use").GetDiag()
	}
	// catalog
	catalogList := make([]map[string]interface{}, 0)
	if volume.CatalogOffering != nil {
		versionCrn := ""
		if volume.CatalogOffering.Version != nil && volume.CatalogOffering.Version.CRN != nil {
			versionCrn = *volume.CatalogOffering.Version.CRN
		}
		catalogMap := map[string]interface{}{}
		if versionCrn != "" {
			catalogMap[isVolumeCatalogOfferingVersionCrn] = versionCrn
		}
		if volume.CatalogOffering.Plan != nil {
			planCrn := ""
			if volume.CatalogOffering.Plan.CRN != nil {
				planCrn = *volume.CatalogOffering.Plan.CRN
			}
			if planCrn != "" {
				catalogMap[isVolumeCatalogOfferingPlanCrn] = *volume.CatalogOffering.Plan.CRN
			}
			if volume.CatalogOffering.Plan.Deleted != nil {
				deletedMap := resourceIbmIsVolumeCatalogOfferingVersionPlanReferenceDeletedToMap(*volume.CatalogOffering.Plan.Deleted)
				catalogMap["deleted"] = []map[string]interface{}{deletedMap}
			}
		}
		catalogList = append(catalogList, catalogMap)
	}

	if err = d.Set(isVolumeCatalogOffering, catalogList); err != nil {
		err = fmt.Errorf("Error setting catalog_offering: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-catalog_offering").GetDiag()
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_volume", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// defined_performance changes

	if err = d.Set("adjustable_capacity_states", volume.AdjustableCapacityStates); err != nil {
		err = fmt.Errorf("Error setting adjustable_capacity_states: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-adjustable_capacity_states").GetDiag()
	}
	if err = d.Set("adjustable_iops_states", volume.AdjustableIopsStates); err != nil {
		err = fmt.Errorf("Error setting adjustable_iops_states: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-adjustable_iops_states").GetDiag()
	}
	if err = d.Set("resource_controller_url", controller+"/vpc-ext/storage/storageVolumes"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set("resource_name", volume.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set("resource_crn", volume.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("resource_status", volume.Status); err != nil {
		err = fmt.Errorf("Error setting resource_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_status").GetDiag()
	}
	if volume.ResourceGroup != nil {
		if err = d.Set(flex.ResourceGroupName, volume.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_group_name").GetDiag()
		}
		if err = d.Set(isVolumeResourceGroup, *volume.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-resource_group").GetDiag()
		}
	}
	operatingSystemList := []map[string]interface{}{}
	if volume.OperatingSystem != nil {
		operatingSystemMap := dataSourceVolumeCollectionVolumesOperatingSystemToMap(*volume.OperatingSystem)
		operatingSystemList = append(operatingSystemList, operatingSystemMap)
	}
	if err = d.Set(isVolumesOperatingSystem, operatingSystemList); err != nil {
		err = fmt.Errorf("Error setting operating_system: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "read", "set-operating_system").GetDiag()
	}

	return nil
}

func resourceIBMISVolumeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

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

	err := volUpdate(context, d, meta, id, name, hasNameChanged, delete)
	if err != nil {
		return err
	}
	return resourceIBMISVolumeRead(context, d, meta)
}

func volUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, name string, hasNameChanged, delete bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var capacity int64
	if delete {
		deleteAllSnapshots(sess, id)
	}

	if d.HasChange(isVolumeAccessTags) {
		options := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, _, err := sess.GetVolumeWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isVolumeAccessTags)

		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *vol.CRN, "", isVolumeAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc volume (%s) access tags: %s", id, err)
		}
	}

	optionsget := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	oldVol, response, err := sess.GetVolumeWithContext(context, optionsget)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	eTag := response.Headers.Get("ETag")
	options := &vpcv1.UpdateVolumeOptions{
		ID: &id,
	}
	options.IfMatch = &eTag

	// bandwidth update
	hasBandwidthChanged := false
	if d.HasChange("bandwidth") {
		hasBandwidthChanged = true
	}

	//name || bandwidth update
	volumeNamePatchModel := &vpcv1.VolumePatch{}
	if hasNameChanged || hasBandwidthChanged {
		if hasNameChanged {
			volumeNamePatchModel.Name = &name
			volumeNamePatch, err := volumeNamePatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeNamePatchModel.AsPatch() for name failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			options.VolumePatch = volumeNamePatch
			_, response, err = sess.UpdateVolumeWithContext(context, options)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			eTag = response.Headers.Get("ETag")
			options.IfMatch = &eTag
		}
		if hasBandwidthChanged {
			volumeNamePatchModel = &vpcv1.VolumePatch{}
			bandwidth := int64(d.Get("bandwidth").(int))
			volumeNamePatchModel.Bandwidth = &bandwidth
			volumeNamePatch, err := volumeNamePatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeNamePatchModel.AsPatch() for bandwidth failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			options.VolumePatch = volumeNamePatch
			_, response, err = sess.UpdateVolumeWithContext(context, options)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			eTag = response.Headers.Get("ETag")
			options.IfMatch = &eTag
		}
	}

	if d.HasChange("allowed_use") {
		allowedUseModel, _ := ResourceIBMIsInstanceMapToVolumeAllowedUsePatchPrototype(d.Get("allowed_use").([]interface{})[0].(map[string]interface{}))
		optionsget := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		_, response, err := sess.GetVolume(optionsget)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")
		options := &vpcv1.UpdateVolumeOptions{
			ID: &id,
		}
		options.IfMatch = &eTag
		volumePatchModel := &vpcv1.VolumePatch{}
		volumePatchModel.AllowedUse = allowedUseModel
		volumePatch, err := volumePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeProfilePatchModel.AsPatch() for iops failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.VolumePatch = volumePatch
		_, _, err = sess.UpdateVolumeWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	// profile/ iops update
	if !d.HasChange(isVolumeProfileName) && *oldVol.Profile.Name == "sdp" && d.HasChange(isVolumeIops) {
		volumeProfilePatchModel := &vpcv1.VolumePatch{}
		iops := int64(d.Get(isVolumeIops).(int))
		volumeProfilePatchModel.Iops = &iops
		volumeProfilePatch, err := volumeProfilePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeProfilePatchModel.AsPatch() for iops failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.VolumePatch = volumeProfilePatch
		_, response, err = sess.UpdateVolumeWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag = response.Headers.Get("ETag")
		options.IfMatch = &eTag
	} else if d.HasChange(isVolumeProfileName) || d.HasChange(isVolumeIops) {
		volumeProfilePatchModel := &vpcv1.VolumePatch{}
		volId := d.Id()
		getvoloptions := &vpcv1.GetVolumeOptions{
			ID: &volId,
		}
		vol, response, err := sess.GetVolumeWithContext(context, getvoloptions)
		if err != nil || vol == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) < 1 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error updating Volume profile/iops because the specified volume %s is not attached to a virtual server instance", volId), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		volAtt := &vol.VolumeAttachments[0]
		insId := *volAtt.Instance.ID
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &insId,
		}
		instance, response, err := sess.GetInstanceWithContext(context, getinsOptions)
		if err != nil || instance == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error retrieving Instance (%s) to which the volume (%s) is attached : %s\n%s", insId, volId, err, response), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if instance != nil && *instance.Status != "running" {
			actiontype := "start"
			createinsactoptions := &vpcv1.CreateInstanceActionOptions{
				InstanceID: &insId,
				Type:       &actiontype,
			}
			_, response, err = sess.CreateInstanceActionWithContext(context, createinsactoptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForInstanceAvailable(sess, insId, d.Timeout(schema.TimeoutCreate), d)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeProfilePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.VolumePatch = volumeProfilePatch
		_, response, err = sess.UpdateVolumeWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag = response.Headers.Get("ETag")
		options.IfMatch = &eTag
		_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	// capacity update
	if d.HasChange(isVolumeCapacity) {
		id := d.Id()
		getvolumeoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := sess.GetVolumeWithContext(context, getvolumeoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag = response.Headers.Get("ETag")
		options.IfMatch = &eTag
		if *vol.Profile.Name != "sdp" {
			if vol.VolumeAttachments == nil || len(vol.VolumeAttachments) == 0 || *vol.VolumeAttachments[0].ID == "" {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error volume capacity can't be updated since volume %s is not attached to any instance for VolumePatch", id), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			insId := vol.VolumeAttachments[0].Instance.ID
			getinsOptions := &vpcv1.GetInstanceOptions{
				ID: insId,
			}
			instance, _, err := sess.GetInstanceWithContext(context, getinsOptions)
			if err != nil || instance == nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if instance != nil && *instance.Status != "running" {
				actiontype := "start"
				createinsactoptions := &vpcv1.CreateInstanceActionOptions{
					InstanceID: insId,
					Type:       &actiontype,
				}
				_, response, err = sess.CreateInstanceActionWithContext(context, createinsactoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForInstanceAvailable(sess, *insId, d.Timeout(schema.TimeoutCreate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		capacity = int64(d.Get(isVolumeCapacity).(int))
		volumeCapacityPatchModel := &vpcv1.VolumePatch{}
		volumeCapacityPatchModel.Capacity = &capacity

		volumeCapacityPatch, err := volumeCapacityPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeCapacityPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.VolumePatch = volumeCapacityPatch
		_, response, err = sess.UpdateVolumeWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
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
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("volumeNamePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_volume", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				options.IfMatch = &eTag
				options.VolumePatch = volumeNamePatch
				_, response, err = sess.UpdateVolumeWithContext(context, options)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForVolumeAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeAvailable failed: %s", err.Error()), "ibm_is_volume", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	return nil
}

func resourceIBMISVolumeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	err := volDelete(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func volDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getvoloptions := &vpcv1.GetVolumeOptions{
		ID: &id,
	}
	volDetails, response, err := sess.GetVolumeWithContext(context, getvoloptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if volDetails.VolumeAttachments != nil {
		for _, volAtt := range volDetails.VolumeAttachments {
			deleteVolumeAttachment := &vpcv1.DeleteInstanceVolumeAttachmentOptions{
				InstanceID: volAtt.Instance.ID,
				ID:         volAtt.ID,
			}
			_, err := sess.DeleteInstanceVolumeAttachmentWithContext(context, deleteVolumeAttachment)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceVolumeAttachmentWithContext failed: %s", err.Error()), "ibm_is_volume", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForInstanceVolumeDetached(sess, d, d.Id(), *volAtt.ID)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceVolumeDetached failed: %s", err.Error()), "ibm_is_volume", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

		}
	}

	options := &vpcv1.DeleteVolumeOptions{
		ID: &id,
	}
	response, err = sess.DeleteVolumeWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVolumeWithContext failed: %s", err.Error()), "ibm_is_volume", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForVolumeDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVolumeDeleted failed: %s", err.Error()), "ibm_is_volume", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForVolumeDeleted(vol *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeDeleting},
		Target:     []string{"done", ""},
		Refresh:    isVolumeDeleteRefreshFunc(vol, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVolumeDeleteRefreshFunc(vol *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		volgetoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := vol.GetVolume(volgetoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return vol, isVolumeDeleted, nil
			}
			return vol, "", fmt.Errorf("[ERROR] Error getting Volume: %s\n%s", err, response)
		}
		return vol, isVolumeDeleting, err
	}
}

func resourceIBMISVolumeExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	id := d.Id()

	exists, err := volExists(d, meta, id)
	return exists, err
}

func volExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
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
		return false, fmt.Errorf("[ERROR] Error getting Volume: %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForVolumeAvailable(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVolumeProvisioning},
		Target:     []string{isVolumeProvisioningDone, ""},
		Refresh:    isVolumeRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVolumeRefreshFunc(client *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		volgetoptions := &vpcv1.GetVolumeOptions{
			ID: &id,
		}
		vol, response, err := client.GetVolume(volgetoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting volume: %s\n%s", err, response)
		}

		if *vol.Status == "available" {
			return vol, isVolumeProvisioningDone, nil
		}

		return vol, isVolumeProvisioning, nil
	}
}

func deleteAllSnapshots(sess *vpcv1.VpcV1, id string) error {
	delete_all_snapshots := new(vpcv1.DeleteSnapshotsOptions)
	delete_all_snapshots.SourceVolumeID = &id
	response, err := sess.DeleteSnapshots(delete_all_snapshots)
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting snapshots from volume %s\n%s", err, response)
	}
	return nil
}

func ResourceIBMIsVolumeAllowedUseMapToVolumeAllowedUsePrototype(modelMap map[string]interface{}) (*vpcv1.VolumeAllowedUsePrototype, error) {
	model := &vpcv1.VolumeAllowedUsePrototype{}
	if modelMap["api_version"] != nil && modelMap["api_version"].(string) != "" {
		model.ApiVersion = core.StringPtr(modelMap["api_version"].(string))
	}
	if modelMap["bare_metal_server"] != nil && modelMap["bare_metal_server"].(string) != "" {
		model.BareMetalServer = core.StringPtr(modelMap["bare_metal_server"].(string))
	}
	if modelMap["instance"] != nil && modelMap["instance"].(string) != "" {
		model.Instance = core.StringPtr(modelMap["instance"].(string))
	}
	return model, nil
}

func ResourceIBMIsInstanceMapToVolumeAllowedUsePatchPrototype(modelMap map[string]interface{}) (*vpcv1.VolumeAllowedUsePatch, error) {
	model := &vpcv1.VolumeAllowedUsePatch{}
	if modelMap["api_version"] != nil && modelMap["api_version"].(string) != "" {
		model.ApiVersion = core.StringPtr(modelMap["api_version"].(string))
	}
	if modelMap["bare_metal_server"] != nil && modelMap["bare_metal_server"].(string) != "" {
		model.BareMetalServer = core.StringPtr(modelMap["bare_metal_server"].(string))
	}
	if modelMap["instance"] != nil && modelMap["instance"].(string) != "" {
		model.Instance = core.StringPtr(modelMap["instance"].(string))
	}
	return model, nil
}

func ResourceIBMIsVolumeAllowedUseToMap(model *vpcv1.VolumeAllowedUsePrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.Instance != nil {
		modelMap["instance"] = *model.Instance
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
}

func ResourceceIBMIsVolumeAllowedUseToMap(model *vpcv1.VolumeAllowedUse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.Instance != nil {
		modelMap["instance"] = *model.Instance
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
}
