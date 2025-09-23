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
	isSnapshotName              = "name"
	isSnapshotResourceGroup     = "resource_group"
	isSnapshotSourceVolume      = "source_volume"
	isSnapshotSourceImage       = "source_image"
	isSnapshotSourceSnapshot    = "source_snapshot"
	isSnapshotSourceSnapshotCRN = "source_snapshot_crn"
	isSnapshotCopies            = "copies"
	isSnapshotUserTags          = "tags"
	isSnapshotAccessTags        = "access_tags"
	isSnapshotCRN               = "crn"
	isSnapshotHref              = "href"
	isSnapshotEncryption        = "encryption"
	isSnapshotEncryptionKey     = "encryption_key"
	isSnapshotOperatingSystem   = "operating_system"
	isSnapshotLCState           = "lifecycle_state"
	isSnapshotMinCapacity       = "minimum_capacity"
	isSnapshotResourceType      = "resource_type"
	isSnapshotSize              = "size"
	isSnapshotBootable          = "bootable"
	isSnapshotDeleting          = "deleting"
	isSnapshotDeleted           = "deleted"
	isSnapshotAvailable         = "stable"
	isSnapshotFailed            = "failed"
	isSnapshotPending           = "pending"
	isSnapshotSuspended         = "suspended"
	isSnapshotUpdating          = "updating"
	isSnapshotWaiting           = "waiting"
	isSnapshotCapturedAt        = "captured_at"
	isSnapshotBackupPolicyPlan  = "backup_policy_plan"

	isSnapshotCatalogOffering           = "catalog_offering"
	isSnapshotCatalogOfferingPlanCrn    = "plan_crn"
	isSnapshotCatalogOfferingVersionCrn = "version_crn"
)

func ResourceIBMSnapshot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSnapshotCreate,
		ReadContext:   resourceIBMISSnapshotRead,
		UpdateContext: resourceIBMISSnapshotUpdate,
		DeleteContext: resourceIBMISSnapshotDelete,
		Exists:        resourceIBMISSnapshotExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				}),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{

			isSnapshotName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", isSnapshotName),
				Description:  "Snapshot name",
			},

			"service_tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The [service tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags) prefixed with `is.snapshot:` associated with this snapshot.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			isSnapshotSourceSnapshotCRN: {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Description:  "Source Snapshot CRN",
				ExactlyOneOf: []string{isSnapshotSourceSnapshotCRN, isSnapshotSourceVolume},
			},

			isSnapshotCopies: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The copies of this snapshot in other regions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for the copied snapshot.",
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
							Computed:    true,
							Description: "The URL for the copied snapshot.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the copied snapshot.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the copied snapshot. The name is unique across all snapshots in the copied snapshot's native region.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource is remote to this region,and identifies the native region.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this region.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this region.",
									},
								},
							},
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},

			isSnapshotResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource group info",
			},

			isSnapshotSourceSnapshot: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, the source snapshot this snapshot was created from.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the source snapshot.",
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
							Computed:    true,
							Description: "The URL for the source snapshot.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the source snapshot.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the source snapshot. The name is unique across all snapshots in the source snapshot's native region.",
						},
						"remote": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource is remote to this region,and identifies the native region.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this region.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this region.",
									},
								},
							},
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},

			isSnapshotConsistencyGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The snapshot consistency group which created this snapshot.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of this snapshot consistency group.",
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
							Computed:    true,
							Description: "The URL for the snapshot consistency group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the snapshot consistency group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for the snapshot consistency group. The name is unique across all snapshot consistency groups in the region.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},

			isSnapshotSourceVolume: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "Snapshot source volume",
				ExactlyOneOf: []string{isSnapshotSourceSnapshotCRN, isSnapshotSourceVolume},
			},

			isSnapshotSourceImage: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If present, the image id from which the data on this volume was most directly provisioned.",
			},

			isSnapshotOperatingSystem: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The globally unique name for the operating system included in this image",
			},

			isSnapshotBootable: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if a boot volume attachment can be created with a volume created from this snapshot",
			},

			isSnapshotLCState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Snapshot lifecycle state",
			},
			isSnapshotCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},
			isSnapshotEncryption: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encryption type of the snapshot",
			},
			isSnapshotEncryptionKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A reference to the root key used to wrap the data encryption key for the source volume.",
			},

			isSnapshotHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for the snapshot",
			},

			isSnapshotMinCapacity: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum capacity of the snapshot",
			},
			isSnapshotResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type of the snapshot",
			},

			isSnapshotAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			isSnapshotSize: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of the snapshot",
			},

			isSnapshotClones: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Zones for creating the snapshot clone",
			},

			isSnapshotUserTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", isSnapshotUserTags)},
				Set:         flex.ResourceIBMVPCHash,
				Description: "User Tags for the snapshot",
			},
			isSnapshotCatalogOffering: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The catalog offering inherited from the snapshot's source. If a virtual server instance is provisioned with a source_snapshot specifying this snapshot, the virtual server instance will use this snapshot's catalog offering, including its pricing plan.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSnapshotCatalogOfferingPlanCrn: {
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
						isSnapshotCatalogOfferingVersionCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this version of a catalog offering",
						},
					},
				},
			},

			isSnapshotBackupPolicyPlan: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, the backup policy plan which created this snapshot.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
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
							Computed:    true,
							Description: "The URL for this backup policy plan.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this backup policy plan.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this backup policy plan.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced",
						},
					},
				},
			},

			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The usage constraints to match against the requested instance or bare metal server properties to determine compatibility. Can only be specified for bootable snapshots.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_version": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", "allowed_use.api_version"),
							Description:  "The expression that must be satisfied by the properties of a virtual server instance provisioned using this snapshot.",
						},
						"bare_metal_server": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", "allowed_use.bare_metal_server"),
							Description:  "The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this snapshot.",
						},
						"instance": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_snapshot", "allowed_use.instance"),
							Description:  "The expression that must be satisfied by a virtual server instance provisioned using this image.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISSnapshotValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSnapshotName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSnapshotUserTags,
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
	ibmISSnapshotResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_snapshot", Schema: validateSchema}
	return &ibmISSnapshotResourceValidator
}

func resourceIBMISSnapshotCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	snapbyVolFlag := false
	options := &vpcv1.CreateSnapshotOptions{}

	snapshotprototypeoptions := &vpcv1.SnapshotPrototypeSnapshotBySourceVolume{}
	snapshotprototypeoptionsbysourcesnapshot := &vpcv1.SnapshotPrototypeSnapshotBySourceSnapshot{}

	// snapshot by source volume
	if sourceVolume, oksv := d.GetOk(isSnapshotSourceVolume); oksv {
		sv := sourceVolume.(string)
		snapshotprototypeoptions.SourceVolume = &vpcv1.VolumeIdentity{
			ID: &sv,
		}
		snapbyVolFlag = true

		if snapshotName, ok := d.GetOk(isSnapshotName); ok {
			name := snapshotName.(string)
			snapshotprototypeoptions.Name = &name
		}

		if allowedUse, ok := d.GetOk("allowed_use"); ok && len(allowedUse.([]interface{})) > 0 {
			allowedUseModel, _ := ResourceIBMIsSnapshotMapToSnapshotAllowedUse(allowedUse.([]interface{})[0].(map[string]interface{}))
			snapshotprototypeoptions.AllowedUse = allowedUseModel
		}

		if grp, ok := d.GetOk(isVPCResourceGroup); ok {
			rg := grp.(string)
			snapshotprototypeoptions.ResourceGroup = &vpcv1.ResourceGroupIdentity{
				ID: &rg,
			}
		}
	} else if sourceSnapshot, okss := d.GetOk(isSnapshotSourceSnapshotCRN); okss {
		ss := sourceSnapshot.(string)
		snapshotprototypeoptionsbysourcesnapshot.SourceSnapshot = &vpcv1.SnapshotIdentityByCRN{
			CRN: &ss,
		}
		snapbyVolFlag = false
		if snapshotName, ok := d.GetOk(isSnapshotName); ok {
			name := snapshotName.(string)
			snapshotprototypeoptionsbysourcesnapshot.Name = &name
		}

		if encryptionKey, ok := d.GetOk(isSnapshotEncryptionKey); ok {
			encryptionKeyString := encryptionKey.(string)
			snapshotprototypeoptionsbysourcesnapshot.EncryptionKey = &vpcv1.EncryptionKeyIdentity{
				CRN: &encryptionKeyString,
			}
		}

		if allowedUse, ok := d.GetOk("allowed_use"); ok && len(allowedUse.([]interface{})) > 0 {
			allowedUseModel, _ := ResourceIBMIsSnapshotMapToSnapshotAllowedUse(allowedUse.([]interface{})[0].(map[string]interface{}))
			snapshotprototypeoptionsbysourcesnapshot.AllowedUse = allowedUseModel
		}

		if grp, ok := d.GetOk(isVPCResourceGroup); ok {
			rg := grp.(string)
			snapshotprototypeoptionsbysourcesnapshot.ResourceGroup = &vpcv1.ResourceGroupIdentity{
				ID: &rg,
			}
		}

	}
	if clones, ok := d.GetOk(isSnapshotClones); ok {
		cloneSet := clones.(*schema.Set)
		if cloneSet.Len() != 0 {
			cloneobjs := make([]vpcv1.SnapshotClonePrototype, cloneSet.Len())
			for i, clone := range cloneSet.List() {
				clonestr := clone.(string)
				cloneobjs[i] = vpcv1.SnapshotClonePrototype{
					Zone: &vpcv1.ZoneIdentity{
						Name: &clonestr,
					},
				}
			}
			if snapbyVolFlag {
				snapshotprototypeoptions.Clones = cloneobjs
			} else {
				snapshotprototypeoptionsbysourcesnapshot.Clones = cloneobjs
			}
		}
	}

	var userTags *schema.Set
	if v, ok := d.GetOk(isSnapshotUserTags); ok {
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
			if snapbyVolFlag {
				snapshotprototypeoptions.UserTags = userTagsArray
			} else {
				snapshotprototypeoptionsbysourcesnapshot.UserTags = userTagsArray
			}
		}
	}
	if snapbyVolFlag {
		options.SnapshotPrototype = snapshotprototypeoptions
	} else {
		options.SnapshotPrototype = snapshotprototypeoptionsbysourcesnapshot
	}
	log.Printf("[DEBUG] Snapshot create")

	snapshot, _, err := sess.CreateSnapshotWithContext(context, options)
	if err != nil || snapshot == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*snapshot.ID)
	log.Printf("[INFO] Snapshot : %s", *snapshot.ID)

	_, err = isWaitForSnapshotAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))

	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot available failed: %s", err.Error()), "ibm_is_snapshot", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := d.GetOk(isSnapshotAccessTags); ok {
		oldList, newList := d.GetChange(isSubnetAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *snapshot.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"[ERROR] Error on create of resource snapshot (%s) access tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISSnapshotRead(context, d, meta)
}

func isWaitForSnapshotAvailable(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Snapshot (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isSnapshotPending},
		Target:     []string{isSnapshotAvailable, isSnapshotFailed},
		Refresh:    isSnapshotRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSnapshotRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSnapshotOptions := &vpcv1.GetSnapshotOptions{
			ID: &id,
		}
		snapshot, response, err := sess.GetSnapshot(getSnapshotOptions)
		if err != nil {
			return nil, isSnapshotFailed, fmt.Errorf("[ERROR] Error getting Snapshot : %s\n%s", err, response)
		}

		if *snapshot.LifecycleState == isSnapshotAvailable {
			return snapshot, *snapshot.LifecycleState, nil
		} else if *snapshot.LifecycleState == isSnapshotFailed {
			return snapshot, *snapshot.LifecycleState, fmt.Errorf("Snapshot (%s) went into failed state during the operation \n [WARNING] Running terraform apply again will remove the tainted snapshot and attempt to create the snapshot again replacing the previous configuration", *snapshot.ID)
		}

		return snapshot, isSnapshotPending, nil
	}
}

func resourceIBMISSnapshotRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := snapshotGet(context, d, meta, id)
	if err != nil {
		return err
	}
	return nil
}

func snapshotGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getSnapshotOptions := &vpcv1.GetSnapshotOptions{
		ID: &id,
	}
	snapshot, response, err := sess.GetSnapshotWithContext(context, getSnapshotOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*snapshot.ID)
	if !core.IsNil(snapshot.Name) {
		if err = d.Set("name", snapshot.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("href", snapshot.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-href").GetDiag()
	}
	if err = d.Set("crn", snapshot.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-crn").GetDiag()
	}
	if err = d.Set("minimum_capacity", flex.IntValue(snapshot.MinimumCapacity)); err != nil {
		err = fmt.Errorf("Error setting minimum_capacity: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-minimum_capacity").GetDiag()
	}
	if err = d.Set("size", flex.IntValue(snapshot.Size)); err != nil {
		err = fmt.Errorf("Error setting size: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-size").GetDiag()
	}
	if err = d.Set("encryption", snapshot.Encryption); err != nil {
		err = fmt.Errorf("Error setting encryption: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-encryption").GetDiag()
	}
	if snapshot.EncryptionKey != nil && snapshot.EncryptionKey.CRN != nil {
		if err = d.Set(isSnapshotEncryptionKey, *snapshot.EncryptionKey.CRN); err != nil {
			err = fmt.Errorf("Error setting encryption_key: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-encryption_key").GetDiag()
		}
	}
	if err = d.Set("service_tags", snapshot.ServiceTags); err != nil {
		err = fmt.Errorf("Error setting service_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-service_tags").GetDiag()
	}
	if err = d.Set("lifecycle_state", snapshot.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", snapshot.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("bootable", snapshot.Bootable); err != nil {
		err = fmt.Errorf("Error setting bootable: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-bootable").GetDiag()
	}
	if !core.IsNil(snapshot.UserTags) {
		if err = d.Set("tags", snapshot.UserTags); err != nil {
			err = fmt.Errorf("Error setting tags: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-tags").GetDiag()
		}
	}
	sourceSnapshotList := []map[string]interface{}{}
	if snapshot.SourceSnapshot != nil {
		sourceSnapshot := map[string]interface{}{}
		sourceSnapshot["crn"] = snapshot.SourceSnapshot.CRN
		sourceSnapshot["href"] = *snapshot.SourceSnapshot.Href
		if snapshot.SourceSnapshot.Deleted != nil {
			snapshotSourceSnapshotDeletedMap := map[string]interface{}{}
			snapshotSourceSnapshotDeletedMap["more_info"] = *snapshot.SourceSnapshot.Deleted.MoreInfo
			sourceSnapshot["deleted"] = []map[string]interface{}{snapshotSourceSnapshotDeletedMap}
		}
		sourceSnapshot["id"] = *snapshot.SourceSnapshot.ID
		sourceSnapshot["name"] = *snapshot.SourceSnapshot.Name
		sourceSnapshot["resource_type"] = *snapshot.SourceSnapshot.ResourceType
		sourceSnapshotList = append(sourceSnapshotList, sourceSnapshot)
	}
	if err = d.Set("source_snapshot", sourceSnapshotList); err != nil {
		err = fmt.Errorf("Error setting source_snapshot: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-source_snapshot").GetDiag()
	}

	// snapshot consistency group
	snapshotConsistencyGroupList := []map[string]interface{}{}
	if snapshot.SnapshotConsistencyGroup != nil {
		snapshotConsistencyGroup := map[string]interface{}{}
		snapshotConsistencyGroup["href"] = snapshot.SnapshotConsistencyGroup.Href
		snapshotConsistencyGroup["crn"] = snapshot.SnapshotConsistencyGroup.CRN
		if snapshot.SnapshotConsistencyGroup.Deleted != nil {
			snapshotConsistencyGroupDeletedMap := map[string]interface{}{}
			snapshotConsistencyGroupDeletedMap["more_info"] = snapshot.SnapshotConsistencyGroup.Deleted.MoreInfo
			snapshotConsistencyGroup["deleted"] = []map[string]interface{}{snapshotConsistencyGroupDeletedMap}
		}
		snapshotConsistencyGroup["id"] = snapshot.SnapshotConsistencyGroup.ID
		snapshotConsistencyGroup["name"] = snapshot.SnapshotConsistencyGroup.Name
		snapshotConsistencyGroup["resource_type"] = snapshot.SnapshotConsistencyGroup.ResourceType
		snapshotConsistencyGroupList = append(snapshotConsistencyGroupList, snapshotConsistencyGroup)
	}
	if err = d.Set("snapshot_consistency_group", snapshotConsistencyGroupList); err != nil {
		err = fmt.Errorf("Error setting snapshot_consistency_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-snapshot_consistency_group").GetDiag()
	}
	snapshotCopies := []map[string]interface{}{}
	if snapshot.Copies != nil {
		for _, copiesItem := range snapshot.Copies {
			copiesMap, err := dataSourceIBMIsSnapshotsSnapshotCopiesItemToMap(&copiesItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "copies-to-map").GetDiag()
			}
			snapshotCopies = append(snapshotCopies, copiesMap)
		}
	}
	if err = d.Set("copies", snapshotCopies); err != nil {
		err = fmt.Errorf("Error setting copies: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-copies").GetDiag()
	}

	if snapshot.ResourceGroup != nil && snapshot.ResourceGroup.ID != nil {
		if err = d.Set("resource_group", *snapshot.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-resource_group").GetDiag()
		}
	}
	if snapshot.SourceVolume != nil && snapshot.SourceVolume.ID != nil {
		if err = d.Set("source_volume", *snapshot.SourceVolume.ID); err != nil {
			err = fmt.Errorf("Error setting source_volume: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-source_volume").GetDiag()
		}
	}

	if snapshot.SourceImage != nil && snapshot.SourceImage.ID != nil {
		if err = d.Set(isSnapshotSourceImage, *snapshot.SourceImage.ID); err != nil {
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-source_volume").GetDiag()
			}
		}
	}

	if snapshot.OperatingSystem != nil && snapshot.OperatingSystem.Name != nil {
		if err = d.Set("operating_system", *snapshot.OperatingSystem.Name); err != nil {
			err = fmt.Errorf("Error setting operating_system: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-operating_system").GetDiag()
		}
	}
	var clones []string
	clones = make([]string, 0)
	if snapshot.Clones != nil {
		for _, clone := range snapshot.Clones {
			if clone.Zone != nil {
				clones = append(clones, *clone.Zone.Name)
			}
		}
	}
	if err = d.Set("clones", flex.NewStringSet(schema.HashString, clones)); err != nil {
		err = fmt.Errorf("Error setting clones: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-clones").GetDiag()
	}

	// catalog
	catalogList := make([]map[string]interface{}, 0)
	if snapshot.CatalogOffering != nil {
		versionCrn := ""
		if snapshot.CatalogOffering.Version != nil && snapshot.CatalogOffering.Version.CRN != nil {
			versionCrn = *snapshot.CatalogOffering.Version.CRN
		}
		catalogMap := map[string]interface{}{}
		if versionCrn != "" {
			catalogMap[isSnapshotCatalogOfferingVersionCrn] = versionCrn
		}
		if snapshot.CatalogOffering.Plan != nil {
			planCrn := ""
			if snapshot.CatalogOffering.Plan != nil && snapshot.CatalogOffering.Plan.CRN != nil {
				planCrn = *snapshot.CatalogOffering.Plan.CRN
			}
			if planCrn != "" {
				catalogMap[isSnapshotCatalogOfferingPlanCrn] = planCrn
			}
			if snapshot.CatalogOffering.Plan.Deleted != nil {
				deletedMap := resourceIbmIsSnapshotCatalogOfferingVersionPlanReferenceDeletedToMap(*snapshot.CatalogOffering.Plan.Deleted)
				catalogMap["deleted"] = []map[string]interface{}{deletedMap}
			}
		}
		catalogList = append(catalogList, catalogMap)
	}
	if err = d.Set("catalog_offering", catalogList); err != nil {
		err = fmt.Errorf("Error setting catalog_offering: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-catalog_offering").GetDiag()
	}

	backupPolicyPlanList := []map[string]interface{}{}
	if snapshot.BackupPolicyPlan != nil {
		backupPolicyPlan := map[string]interface{}{}
		if snapshot.BackupPolicyPlan.Deleted != nil {
			snapshotBackupPolicyPlanDeletedMap := map[string]interface{}{}
			snapshotBackupPolicyPlanDeletedMap["more_info"] = snapshot.BackupPolicyPlan.Deleted.MoreInfo
			backupPolicyPlan["deleted"] = []map[string]interface{}{snapshotBackupPolicyPlanDeletedMap}
		}
		backupPolicyPlan["href"] = snapshot.BackupPolicyPlan.Href
		backupPolicyPlan["id"] = snapshot.BackupPolicyPlan.ID
		backupPolicyPlan["name"] = snapshot.BackupPolicyPlan.Name
		backupPolicyPlan["resource_type"] = snapshot.BackupPolicyPlan.ResourceType
		backupPolicyPlanList = append(backupPolicyPlanList, backupPolicyPlan)
	}
	if err = d.Set("backup_policy_plan", backupPolicyPlanList); err != nil {
		err = fmt.Errorf("Error setting backup_policy_plan: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-backup_policy_plan").GetDiag()
	}
	allowedUses := []map[string]interface{}{}
	if snapshot.AllowedUse != nil {
		modelMap, err := DataSourceIBMIsSnapshotAllowedUseToMap(snapshot.AllowedUse)
		if err != nil {
			err = fmt.Errorf("Error setting allowed_use: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-allowed_use").GetDiag()
		}
		allowedUses = append(allowedUses, modelMap)
	}
	if err = d.Set("allowed_use", allowedUses); err != nil {
		err = fmt.Errorf("Error setting allowed_use: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-allowed_use").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *snapshot.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"[ERROR] Error on get of resource snapshot (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set("access_tags", accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "read", "set-access_tags").GetDiag()
	}
	return nil
}

func resourceIBMISSnapshotUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	name := ""
	hasChanged := false

	if d.HasChange(isSnapshotName) {
		name = d.Get(isSnapshotName).(string)
		hasChanged = true
	}
	err := snapshotUpdate(context, d, meta, id, name, hasChanged)
	if err != nil {
		return err
	}
	return resourceIBMISSnapshotRead(context, d, meta)
}

func snapshotUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSnapshotOptions := &vpcv1.GetSnapshotOptions{
		ID: &id,
	}
	_, response, err := sess.GetSnapshotWithContext(context, getSnapshotOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	eTag := response.Headers.Get("ETag")

	updateSnapshotOptions := &vpcv1.UpdateSnapshotOptions{
		ID: &id,
	}
	updateSnapshotOptions.IfMatch = &eTag

	// user tags update
	if d.HasChange(isSnapshotUserTags) {
		var userTags *schema.Set
		if v, ok := d.GetOk(isSnapshotUserTags); ok {

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
				snapshotPatchModel := &vpcv1.SnapshotPatch{}
				snapshotPatchModel.UserTags = userTagsArray
				snapshotPatch, err := snapshotPatchModel.AsPatch()
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("snapshotPatchModel.AsPatch failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				updateSnapshotOptions.SnapshotPatch = snapshotPatch
				_, _, err = sess.UpdateSnapshot(updateSnapshotOptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForSnapshotUpdate(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot update failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	if d.HasChange(isSnapshotName) {
		updateSnapshotOptions := &vpcv1.UpdateSnapshotOptions{
			ID: &id,
		}
		snapshotPatchModel := &vpcv1.SnapshotPatch{
			Name: &name,
		}
		snapshotPatch, err := snapshotPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("snapshotPatchModel.AsPatch failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateSnapshotOptions.SnapshotPatch = snapshotPatch
		_, _, err = sess.UpdateSnapshot(updateSnapshotOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForSnapshotUpdate(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot update failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}

	if d.HasChange("allowed_use") && len(d.Get("allowed_use").([]interface{})) > 0 {
		getSnapshotOptions := &vpcv1.GetSnapshotOptions{
			ID: &id,
		}
		_, response, err := sess.GetSnapshot(getSnapshotOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		eTag := response.Headers.Get("ETag")

		updateSnapshotOptions := &vpcv1.UpdateSnapshotOptions{
			ID: &id,
		}
		updateSnapshotOptions.IfMatch = &eTag
		allowedUseModel, err := ResourceIBMIsSnapshotMapToSnapshotAllowedUsePatch(d.Get("allowed_use").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error in settting allowed_use: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		snapshotPatchModel := &vpcv1.SnapshotPatch{
			AllowedUse: allowedUseModel,
		}
		snapshotPatch, err := snapshotPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("snapshotPatchModel.AsPatch failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateSnapshotOptions.SnapshotPatch = snapshotPatch
		_, response, err = sess.UpdateSnapshot(updateSnapshotOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForSnapshotUpdate(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot update failed: %s", err.Error()), "ibm_is_snapshot", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if d.HasChange(isSnapshotClones) {
		ovs, nvs := d.GetChange(isSnapshotClones)
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)

		remove := flex.ExpandStringList(ov.Difference(nv).List())
		add := flex.ExpandStringList(nv.Difference(ov).List())

		if len(add) > 0 {
			for i := range add {
				createCloneOptions := &vpcv1.CreateSnapshotCloneOptions{
					ID:       &id,
					ZoneName: &add[i],
				}
				_, _, err := sess.CreateSnapshotCloneWithContext(context, createCloneOptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSnapshotCloneWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForCloneAvailable(sess, d, id, add[i])
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot clone available failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				delCloneOptions := &vpcv1.DeleteSnapshotCloneOptions{
					ID:       &id,
					ZoneName: &remove[i],
				}
				_, err := sess.DeleteSnapshotCloneWithContext(context, delCloneOptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSnapshotCloneWithContext failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForCloneDeleted(sess, d, d.Id(), remove[i])
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot clone deleted failed: %s", err.Error()), "ibm_is_snapshot", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	if d.HasChange(isSnapshotAccessTags) {
		oldList, newList := d.GetChange(isSnapshotAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSnapshotCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf(
				"[ERROR] Error on update of resource snapshot (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func isWaitForSnapshotUpdate(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Snapshot (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isSnapshotUpdating},
		Target:     []string{isSnapshotAvailable, isSnapshotFailed},
		Refresh:    isSnapshotUpdateRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isSnapshotUpdateRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSnapshotOptions := &vpcv1.GetSnapshotOptions{
			ID: &id,
		}
		snapshot, response, err := sess.GetSnapshot(getSnapshotOptions)
		if err != nil {
			return nil, isSnapshotFailed, fmt.Errorf("[ERROR] Error getting Snapshot : %s\n%s", err, response)
		}

		if *snapshot.LifecycleState == isSnapshotAvailable || *snapshot.LifecycleState == isSnapshotFailed {
			return snapshot, *snapshot.LifecycleState, nil
		} else if *snapshot.LifecycleState == isSnapshotFailed {
			return snapshot, *snapshot.LifecycleState, fmt.Errorf("Snapshot (%s) went into failed state during the operation \n [WARNING] Running terraform apply again will remove the tainted snapshot and attempt to create the snapshot again replacing the previous configuration", *snapshot.ID)
		}

		return snapshot, isSnapshotUpdating, nil
	}
}
func isWaitForCloneAvailable(sess *vpcv1.VpcV1, d *schema.ResourceData, id, zoneName string) (interface{}, error) {
	log.Printf("Waiting for Snapshot (%s) clone (%s) to be available.", id, zoneName)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"false"},
		Target:     []string{"true", "deleted"},
		Refresh:    isSnapshotCloneRefreshFunc(sess, id, zoneName),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isSnapshotCloneRefreshFunc(sess *vpcv1.VpcV1, id, zoneName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSnapshotCloneOptions := &vpcv1.GetSnapshotCloneOptions{
			ID:       &id,
			ZoneName: &zoneName,
		}
		clone, response, err := sess.GetSnapshotClone(getSnapshotCloneOptions)
		if err != nil {
			if response.StatusCode == 404 {
				return nil, "deleted", nil
			}
			return nil, "deleted", fmt.Errorf("[ERROR] Error getting Snapshot clone : %s\n%s", err, response)
		}

		if *clone.Available == true {
			return clone, "true", nil
		}

		return clone, "false", nil
	}
}

func isSnapshotCloneDeleteRefreshFunc(sess *vpcv1.VpcV1, id, zoneName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getSnapshotCloneOptions := &vpcv1.GetSnapshotCloneOptions{
			ID:       &id,
			ZoneName: &zoneName,
		}
		clone, response, err := sess.GetSnapshotClone(getSnapshotCloneOptions)
		if err != nil {
			if response.StatusCode == 404 {
				return clone, "deleted", nil
			}
			return clone, "false", fmt.Errorf("[ERROR] Error getting Snapshot clone : %s\n%s", err, response)
		}

		return clone, "true", nil
	}
}

func isWaitForCloneDeleted(sess *vpcv1.VpcV1, d *schema.ResourceData, id, zoneName string) (interface{}, error) {
	log.Printf("Waiting for Snapshot (%s) clone (%s) to be deleted.", id, zoneName)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"true"},
		Target:     []string{"false", "deleted"},
		Refresh:    isSnapshotCloneDeleteRefreshFunc(sess, id, zoneName),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func resourceIBMISSnapshotDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	err := snapshotDelete(context, d, meta, id)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func snapshotDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSnapshotOptions := &vpcv1.GetSnapshotOptions{
		ID: &id,
	}
	_, response, err := sess.GetSnapshotWithContext(context, getSnapshotOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteSnapshotOptions := &vpcv1.DeleteSnapshotOptions{
		ID: &id,
	}
	response, err = sess.DeleteSnapshotWithContext(context, deleteSnapshotOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSnapshotWithContext failed: %s", err.Error()), "ibm_is_snapshot", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForSnapshotDeleted(sess, id, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Wait for Snapshot deleted failed: %s", err.Error()), "ibm_is_snapshot", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func isWaitForSnapshotDeleted(sess *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Snapshot (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isSnapshotDeleting},
		Target:     []string{isSnapshotDeleted, isSnapshotFailed},
		Refresh:    isSnapshotDeleteRefreshFunc(sess, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isSnapshotDeleteRefreshFunc(sess *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Refresh function for Snapshot delete.")
		getSnapshotOptions := &vpcv1.GetSnapshotOptions{
			ID: &id,
		}
		snapshot, response, err := sess.GetSnapshot(getSnapshotOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return snapshot, isSnapshotDeleted, nil
			}
			return nil, isSnapshotFailed, fmt.Errorf("[ERROR] The Snapshot %s failed to delete: %s\n%s", id, err, response)
		}
		return snapshot, *snapshot.LifecycleState, nil
	}
}

func resourceIBMISSnapshotExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	exists, err := snapshotExists(d, meta, id)
	return exists, err
}

func snapshotExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_snapshot", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	getSnapshotOptions := &vpcv1.GetSnapshotOptions{
		ID: &id,
	}
	_, response, err := sess.GetSnapshot(getSnapshotOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] GetSnapshot failed: %s\n%s", err, response)
	}
	return true, nil
}
func ResourceIBMIsSnapshotMapToSnapshotAllowedUsePatch(modelMap map[string]interface{}) (*vpcv1.SnapshotAllowedUsePatch, error) {
	model := &vpcv1.SnapshotAllowedUsePatch{}
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

func ResourceIBMIsSnapshotMapToSnapshotAllowedUse(modelMap map[string]interface{}) (*vpcv1.SnapshotAllowedUsePrototype, error) {
	model := &vpcv1.SnapshotAllowedUsePrototype{}
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
