// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsBackupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsBackupPolicyCreate,
		ReadContext:   resourceIBMIsBackupPolicyRead,
		UpdateContext: resourceIBMIsBackupPolicyUpdate,
		DeleteContext: resourceIBMIsBackupPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"match_resource_types": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				// ValidateFunc: InvokeValidator("ibm_is_backup_policy", "match_resource_types"),
				// Default:     ["volume"],
				Description: "A resource type this backup policy applies to. Resources that have both a matching type and a matching user tag will be subject to the backup policy.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"match_user_tags": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				// ValidateFunc: InvokeValidator("ibm_is_backup_policy", "match_user_tags"),
				Description: "The user tags this backup policy applies to. Resources that have both a matching user tag and a matching type will be subject to the backup policy.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: InvokeValidator("ibm_is_backup_policy", "name"),
				Description: "The user-defined name for this backup policy. Names must be unique within the region this backup policy resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"plans": &schema.Schema{
				Type: schema.TypeList,
				// MaxItems:    1,
				MinItems:    1,
				Optional:    true,
				Description: "The prototype objects for backup plans to be created for this backup policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the plan is active.",
						},
						"attach_user_tags": &schema.Schema{
							Type:        schema.TypeSet,
							Optional:    true,
							Set:         schema.HashString,
							Description: "User tags to attach to each backup (snapshot) created by this plan. If unspecified, no user tags will be attached.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"clone_policy": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_snapshots": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum number of recent snapshots (per source) that will keep clones.",
									},
									"zones": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "The zone this backup policy plan will create snapshot clones in.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The globally unique name for this zone.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The URL for this zone.",
												},
											},
										},
									},
								},
							},
						},
						"copy_user_tags": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Indicates whether to copy the source's user tags to the created backups (snapshots).",
						},
						"cron_spec": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							// ValidateFunc: InvokeValidator("ibm_is_backup_policy", "cron_spec"),
							Description: "The cron specification for the backup schedule.",
						},
						"deletion_trigger": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delete_after": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     30,
										Description: "The maximum number of days to keep each backup after creation.",
									},
									"delete_over_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum number of recent backups to keep. If unspecified, there will be no maximum.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							// ValidateFunc: InvokeValidator("ibm_is_backup_policy", "name"),
							Description: "The user-defined name for this backup policy plan. Names must be unique within the backup policy this plan resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
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
					},
				},
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique identifier for this resource group.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the backup policy was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this backup policy.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this backup policy.",
			},
			"last_job_completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the most recent job for this backup policy completed.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the backup policy.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMIsBackupPolicyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
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
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cron_spec",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^((((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5,7})$`,
			MinValueLength:             9,
			MaxValueLength:             63,
		},
	)
	// validateSchema = append(validateSchema,
	// 	ValidateSchema{
	// 		Identifier:                 "id",
	// 		ValidateFunctionIdentifier: ValidateRegexpLen,
	// 		Type:                       TypeString,
	// 		Optional:                   true,
	// 		Regexp:                     `^[0-9a-f]{32}$`,
	// 		MinValueLength:             9,
	// 		MaxValueLength:             63,
	// 	},
	// )
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "attach_user_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             0,
			MaxValueLength:             128,
		},
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "match_user_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "match_resource_types",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)
	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_backup_policy", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsBackupPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	createBackupPolicyOptions := &vpcv1.CreateBackupPolicyOptions{}

	if _, ok := d.GetOk("match_resource_types"); ok {
		createBackupPolicyOptions.SetMatchResourceTypes(flex.ExpandStringList((d.Get("match_resource_types").(*schema.Set)).List()))
	}
	if _, ok := d.GetOk("match_user_tags"); ok {
		createBackupPolicyOptions.SetMatchUserTags((flex.ExpandStringList((d.Get("match_user_tags").(*schema.Set)).List())))
	}
	if _, ok := d.GetOk("name"); ok {
		createBackupPolicyOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("plans"); ok {
		var plans []vpcv1.BackupPolicyPlanPrototype
		for _, e := range d.Get("plans").([]interface{}) {
			value := e.(map[string]interface{})
			plansItem := resourceIBMIsBackupPolicyMapToBackupPolicyPlanPrototype(value)
			plans = append(plans, plansItem)
		}
		createBackupPolicyOptions.SetPlans(plans)
	}
	if _, ok := d.GetOk("resource_group"); ok {
		resourceGroup := resourceIBMIsBackupPolicyMapToResourceGroupIdentity(d.Get("resource_group.0").(map[string]interface{}))
		// createBackupPolicyOptions.SetResourceGroup(resourceGroup)
		createBackupPolicyOptions.ResourceGroup = resourceGroup
	}

	backupPolicy, response, err := vpcClient.CreateBackupPolicyWithContext(context, createBackupPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateBackupPolicyWithContext failed %s\n%s", err, response))
	}

	d.SetId(*backupPolicy.ID)

	return resourceIBMIsBackupPolicyRead(context, d, meta)
}

func resourceIBMIsBackupPolicyMapToBackupPolicyPlanPrototype(backupPolicyPlanPrototypeMap map[string]interface{}) vpcv1.BackupPolicyPlanPrototype {

	backupPolicyPlanPrototype := vpcv1.BackupPolicyPlanPrototype{}

	if backupPolicyPlanPrototypeMap["active"] != nil {
		backupPolicyPlanPrototype.Active = core.BoolPtr(backupPolicyPlanPrototypeMap["active"].(bool))
	}

	if len(backupPolicyPlanPrototypeMap["attach_user_tags"].(*schema.Set).List()) != 0 {
		attachUserTags := []string{}
		for _, attachUserTagsItem := range backupPolicyPlanPrototypeMap["attach_user_tags"].(*schema.Set).List() {
			attachUserTags = append(attachUserTags, attachUserTagsItem.(string))
		}
		backupPolicyPlanPrototype.AttachUserTags = attachUserTags
	}

	// if backupPolicyPlanPrototypeMap["clone_policy"] != nil {
	// 	backupPolicyPlanClonePolicyPrototypeDeletionTrigger := backupPolicyPlanPrototypeMap["clone_policy"].([]interface{})[0]
	// 	backupPolicyPlanClonePolicyPrototype := resourceIBMIsBackupPolicyMapToBackupPolicyPlanClonePolicyPrototype(backupPolicyPlanClonePolicyPrototypeDeletionTrigger.(map[string]interface{}))
	// 	backupPolicyPlanPrototype.ClonePolicy = &backupPolicyPlanClonePolicyPrototype
	// }

	if backupPolicyPlanPrototypeMap["copy_user_tags"] != nil {
		backupPolicyPlanPrototype.CopyUserTags = core.BoolPtr(backupPolicyPlanPrototypeMap["copy_user_tags"].(bool))
	}

	backupPolicyPlanPrototype.CronSpec = core.StringPtr(backupPolicyPlanPrototypeMap["cron_spec"].(string))

	if backupPolicyPlanPrototypeMap["deletion_trigger"] != nil {
		backupPolicyPlanPrototypeMapDeletionTrigger := backupPolicyPlanPrototypeMap["deletion_trigger"].([]interface{})[0]
		clonePolicy := resourceIBMIsBackupPolicyMapToBackupPolicyPlanDeletionTriggerPrototype(backupPolicyPlanPrototypeMapDeletionTrigger.(map[string]interface{}))
		backupPolicyPlanPrototype.DeletionTrigger = &clonePolicy
	}
	if backupPolicyPlanPrototypeMap["name"] != nil {
		backupPolicyPlanPrototype.Name = core.StringPtr(backupPolicyPlanPrototypeMap["name"].(string))
	}

	return backupPolicyPlanPrototype
}

// func resourceIBMIsBackupPolicyMapToBackupPolicyPlanClonePolicyPrototype(backupPolicyPlanClonePolicyPrototypeMap map[string]interface{}) vpcv1.BackupPolicyPlanClonePolicyPrototype {
// 	backupPolicyPlanClonePolicyPrototype := vpcv1.BackupPolicyPlanClonePolicyPrototype{}

// 	if backupPolicyPlanClonePolicyPrototypeMap["max_snapshots"] != nil {
// 		backupPolicyPlanClonePolicyPrototype.MaxSnapshots = core.Int64Ptr(int64(backupPolicyPlanClonePolicyPrototypeMap["max_snapshots"].(int)))
// 	}
// 	zones := []vpcv1.ZoneIdentityIntf{}
// 	for _, zonesItem := range backupPolicyPlanClonePolicyPrototypeMap["zones"].([]interface{}) {
// 		zonesItemModel := resourceIBMIsBackupPolicyMapToZoneIdentity(zonesItem.(map[string]interface{}))
// 		zones = append(zones, zonesItemModel)
// 	}
// 	backupPolicyPlanClonePolicyPrototype.Zones = zones

// 	return backupPolicyPlanClonePolicyPrototype
// }

func resourceIBMIsBackupPolicyMapToZoneIdentity(zoneIdentityMap map[string]interface{}) vpcv1.ZoneIdentityIntf {
	zoneIdentity := vpcv1.ZoneIdentity{}

	if zoneIdentityMap["name"] != nil {
		zoneIdentity.Name = core.StringPtr(zoneIdentityMap["name"].(string))
	}
	if zoneIdentityMap["href"] != nil {
		zoneIdentity.Href = core.StringPtr(zoneIdentityMap["href"].(string))
	}

	return &zoneIdentity
}

func resourceIBMIsBackupPolicyMapToZoneIdentityByName(zoneIdentityByNameMap map[string]interface{}) vpcv1.ZoneIdentityByName {
	zoneIdentityByName := vpcv1.ZoneIdentityByName{}

	zoneIdentityByName.Name = core.StringPtr(zoneIdentityByNameMap["name"].(string))

	return zoneIdentityByName
}

func resourceIBMIsBackupPolicyMapToZoneIdentityByHref(zoneIdentityByHrefMap map[string]interface{}) vpcv1.ZoneIdentityByHref {
	zoneIdentityByHref := vpcv1.ZoneIdentityByHref{}

	zoneIdentityByHref.Href = core.StringPtr(zoneIdentityByHrefMap["href"].(string))

	return zoneIdentityByHref
}

func resourceIBMIsBackupPolicyMapToBackupPolicyPlanDeletionTriggerPrototype(backupPolicyPlanDeletionTriggerPrototypeMap map[string]interface{}) vpcv1.BackupPolicyPlanDeletionTriggerPrototype {
	backupPolicyPlanDeletionTriggerPrototype := vpcv1.BackupPolicyPlanDeletionTriggerPrototype{}

	if backupPolicyPlanDeletionTriggerPrototypeMap["delete_after"] != nil {
		backupPolicyPlanDeletionTriggerPrototype.DeleteAfter = core.Int64Ptr(int64(backupPolicyPlanDeletionTriggerPrototypeMap["delete_after"].(int)))
	}
	if backupPolicyPlanDeletionTriggerPrototypeMap["delete_over_count"] != nil {
		backupPolicyPlanDeletionTriggerPrototype.DeleteOverCount = core.Int64Ptr(int64(backupPolicyPlanDeletionTriggerPrototypeMap["delete_over_count"].(int)))
	}

	return backupPolicyPlanDeletionTriggerPrototype
}

func resourceIBMIsBackupPolicyMapToResourceGroupIdentity(resourceGroupIdentityMap map[string]interface{}) vpcv1.ResourceGroupIdentityIntf {
	resourceGroupIdentity := vpcv1.ResourceGroupIdentity{}

	if resourceGroupIdentityMap["id"] != nil {
		resourceGroupIdentity.ID = core.StringPtr(resourceGroupIdentityMap["id"].(string))
	}

	return &resourceGroupIdentity
}

func resourceIBMIsBackupPolicyMapToResourceGroupIdentityByID(resourceGroupIdentityByIDMap map[string]interface{}) vpcv1.ResourceGroupIdentityByID {
	resourceGroupIdentityByID := vpcv1.ResourceGroupIdentityByID{}

	resourceGroupIdentityByID.ID = core.StringPtr(resourceGroupIdentityByIDMap["id"].(string))

	return resourceGroupIdentityByID
}

func resourceIBMIsBackupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}

	getBackupPolicyOptions.SetID(d.Id())

	backupPolicy, response, err := vpcClient.GetBackupPolicyWithContext(context, getBackupPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetBackupPolicyWithContext failed %s\n%s", err, response))
	}

	if backupPolicy.MatchResourceTypes != nil {
		if err = d.Set("match_resource_types", backupPolicy.MatchResourceTypes); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting match_resource_types: %s", err))
		}
	}
	if backupPolicy.MatchUserTags != nil {
		if err = d.Set("match_user_tags", backupPolicy.MatchUserTags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting match_user_tags: %s", err))
		}
	}
	if backupPolicy.Name != nil {
		if err = d.Set("name", backupPolicy.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if backupPolicy.Plans != nil {
		plans := []map[string]interface{}{}
		for _, plansItem := range backupPolicy.Plans {

			backupPolicyPlanPrototypeMap := map[string]interface{}{}
			if plansItem.Deleted != nil {
				DeletedMap := resourceIbmIsBackupPolicyBackupPolicyPlanReferenceDeletedToMap(*plansItem.Deleted)
				backupPolicyPlanPrototypeMap["deleted"] = []map[string]interface{}{DeletedMap}
			}
			backupPolicyPlanPrototypeMap["href"] = plansItem.Href
			backupPolicyPlanPrototypeMap["id"] = *plansItem.ID
			backupPolicyPlanPrototypeMap["name"] = plansItem.Name
			backupPolicyPlanPrototypeMap["resource_type"] = plansItem.ResourceType

			getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}
			getBackupPolicyPlanOptions.SetBackupPolicyID(d.Id())
			getBackupPolicyPlanOptions.SetID(*plansItem.ID)
			backupPolicyPlan, response, err := vpcClient.GetBackupPolicyPlanWithContext(context, getBackupPolicyPlanOptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					d.SetId("")
					return nil
				}
				log.Printf("[DEBUG] GetBackupPolicyPlanWithContext failed %s\n%s", err, response)
				return diag.FromErr(fmt.Errorf("GetBackupPolicyPlanWithContext failed %s\n%s", err, response))
			}
			if backupPolicyPlan.CronSpec != nil {
				backupPolicyPlanPrototypeMap["cron_spec"] = backupPolicyPlan.CronSpec
			}

			if backupPolicyPlan.Active != nil {
				backupPolicyPlanPrototypeMap["active"] = backupPolicyPlan.Active
			}

			if backupPolicyPlan.AttachUserTags != nil {
				backupPolicyPlanPrototypeMap["attach_user_tags"] = backupPolicyPlan.AttachUserTags
			}

			if backupPolicyPlan.CopyUserTags != nil {
				backupPolicyPlanPrototypeMap["copy_user_tags"] = backupPolicyPlan.CopyUserTags
			}

			if backupPolicyPlan.DeletionTrigger != nil {
				deletionTriggerMap := resourceIBMIsBackupPolicyPlanBackupPolicyPlanDeletionTriggerPrototypeToMap(*backupPolicyPlan.DeletionTrigger)
				backupPolicyPlanPrototypeMap["deletion_trigger"] = []map[string]interface{}{deletionTriggerMap}
			}

			if backupPolicyPlan.Name != nil {
				backupPolicyPlanPrototypeMap["name"] = backupPolicyPlan.Name
			}

			if backupPolicyPlan.CreatedAt != nil {
				backupPolicyPlanPrototypeMap["created_at"] = flex.DateTimeToString(backupPolicyPlan.CreatedAt)
			}

			if backupPolicyPlan.Href != nil {
				backupPolicyPlanPrototypeMap["href"] = backupPolicyPlan.Href
			}

			if backupPolicyPlan.LifecycleState != nil {
				backupPolicyPlanPrototypeMap["lifecycle_state"] = backupPolicyPlan.LifecycleState
			}

			if backupPolicyPlan.ResourceType != nil {
				backupPolicyPlanPrototypeMap["resource_type"] = backupPolicyPlan.ResourceType
			}
			// plansItemMap := resourceIBMIsBackupPolicyBackupPolicyPlanReferenceToMap(context, d.Id(), meta, plansItem)
			plans = append(plans, backupPolicyPlanPrototypeMap)
		}
		log.Println("plans")
		log.Println(plans)
		log.Println("&plans")
		log.Println(&plans)
		if err = d.Set("plans", plans); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting plans: %s", err))
		}
	}
	if backupPolicy.ResourceGroup != nil {
		resourceGroupMap := resourceIBMIsBackupPolicyResourceGroupIdentityToMap(*backupPolicy.ResourceGroup)
		if err = d.Set("resource_group", []map[string]interface{}{resourceGroupMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
		}
	}
	if backupPolicy.CreatedAt != nil {
		if err = d.Set("created_at", flex.DateTimeToString(backupPolicy.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}

	if backupPolicy.CRN != nil {
		if err = d.Set("crn", backupPolicy.CRN); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
		}
	}

	if backupPolicy.Href != nil {
		if err = d.Set("href", backupPolicy.Href); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
		}
	}

	if backupPolicy.LastJobCompletedAt != nil {
		if err = d.Set("last_job_completed_at", flex.DateTimeToString(backupPolicy.LastJobCompletedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_job_completed_at: %s", err))
		}
	}

	if backupPolicy.LifecycleState != nil {
		if err = d.Set("lifecycle_state", backupPolicy.LifecycleState); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting lifecycle_state: %s", err))
		}
	}

	if backupPolicy.ResourceType != nil {
		if err = d.Set("resource_type", backupPolicy.ResourceType); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
		}
	}

	if err = d.Set("version", response.Headers.Get("Etag")); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	return nil
}

// func resourceIBMIsBackupPolicyBackupPolicyPlanReferenceToMap(context context.Context, backupPolicyId string, meta interface{}, backupPolicyPlanReference vpcv1.BackupPolicyPlanReference) map[string]interface{} {

// 	backupPolicyPlanPrototypeMap := map[string]interface{}{}
// 	if backupPolicyPlanReference.Deleted != nil {
// 		DeletedMap := resourceIbmIsBackupPolicyBackupPolicyPlanReferenceDeletedToMap(*backupPolicyPlanReference.Deleted)
// 		backupPolicyPlanPrototypeMap["deleted"] = []map[string]interface{}{DeletedMap}
// 	}
// 	backupPolicyPlanPrototypeMap["href"] = backupPolicyPlanReference.Href
// 	backupPolicyPlanPrototypeMap["id"] = *backupPolicyPlanReference.ID
// 	backupPolicyPlanPrototypeMap["name"] = backupPolicyPlanReference.Name
// 	backupPolicyPlanPrototypeMap["resource_type"] = backupPolicyPlanReference.ResourceType

// 	getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}
// 	getBackupPolicyPlanOptions.SetBackupPolicyID(backupPolicyId)
// 	getBackupPolicyPlanOptions.SetID(*backupPolicyPlanReference.ID)
// 	backupPolicyPlan, response, err := vpcClient.GetBackupPolicyPlanWithContext(context, getBackupPolicyPlanOptions)
// 	if err != nil {
// 		log.Printf("[DEBUG] GetBackupPolicyPlanWithContext failed %s\n%s", err, response)
// 		// return diag.FromErr(fmt.Errorf("GetBackupPolicyPlanWithContext failed %s\n%s", err, response))
// 	}
// 	if backupPolicyPlan.CronSpec != nil {
// 		backupPolicyPlanPrototypeMap["cron_spec"] = backupPolicyPlan.CronSpec
// 	}

// 	if backupPolicyPlan.Active != nil {
// 		backupPolicyPlanPrototypeMap["active"] = backupPolicyPlan.Active
// 	}

// 	if backupPolicyPlan.AttachUserTags != nil {
// 		backupPolicyPlanPrototypeMap["attach_user_tags"] = backupPolicyPlan.AttachUserTags
// 	}

// 	if backupPolicyPlan.CopyUserTags != nil {
// 		backupPolicyPlanPrototypeMap["copy_user_tags"] = backupPolicyPlan.CopyUserTag
// 	}

// 	if backupPolicyPlan.DeletionTrigger != nil {
// 		deletionTriggerMap := resourceIBMIsBackupPolicyPlanBackupPolicyPlanDeletionTriggerPrototypeToMap(*backupPolicyPlan.DeletionTrigger)
// 		backupPolicyPlanPrototypeMap["deletion_trigger"] = []map[string]interface{}{deletionTriggerMap}
// 	}

// 	if backupPolicyPlan.Name != nil {
// 		backupPolicyPlanPrototypeMap["name"] = backupPolicyPlan.Name
// 	}

// 	if backupPolicyPlan.CreatedAt != nil {
// 		backupPolicyPlanPrototypeMap["created_at"] = backupPolicyPlan.CreatedAt
// 	}

// 	if backupPolicyPlan.Href != nil {
// 		backupPolicyPlanPrototypeMap["href"] = backupPolicyPlan.Href
// 	}

// 	if backupPolicyPlan.LifecycleState != nil {
// 		backupPolicyPlanPrototypeMap["lifecycle_state"] = backupPolicyPlan.LifecycleState
// 	}

// 	if backupPolicyPlan.ResourceType != nil {
// 		backupPolicyPlanPrototypeMap["resource_type"] = backupPolicyPlan.ResourceType
// 	}
// 	return backupPolicyPlanPrototypeMap
// }

func resourceIbmIsBackupPolicyBackupPolicyPlanReferenceDeletedToMap(backupPolicyPlanReferenceDeleted vpcv1.BackupPolicyPlanReferenceDeleted) map[string]interface{} {
	backupPolicyPlanReferenceDeletedMap := map[string]interface{}{}

	backupPolicyPlanReferenceDeletedMap["more_info"] = backupPolicyPlanReferenceDeleted.MoreInfo

	return backupPolicyPlanReferenceDeletedMap
}

// func resourceIBMIsBackupPolicyBackupPolicyPlanClonePolicyPrototypeToMap(backupPolicyPlanClonePolicyPrototype vpcv1.BackupPolicyPlanClonePolicyPrototype) map[string]interface{} {
// 	backupPolicyPlanClonePolicyPrototypeMap := map[string]interface{}{}

// 	if backupPolicyPlanClonePolicyPrototype.MaxSnapshots != nil {
// 		backupPolicyPlanClonePolicyPrototypeMap["max_snapshots"] = intValue(backupPolicyPlanClonePolicyPrototype.MaxSnapshots)
// 	}
// 	zones := []map[string]interface{}{}
// 	for _, zonesItem := range backupPolicyPlanClonePolicyPrototype.Zones {
// 		zonesItemMap := resourceIBMIsBackupPolicyZoneIdentityToMap(zonesItem)
// 		zones = append(zones, zonesItemMap)
// 	}
// 	backupPolicyPlanClonePolicyPrototypeMap["zones"] = zones

// 	return backupPolicyPlanClonePolicyPrototypeMap
// }

func resourceIBMIsBackupPolicyZoneIdentityToMap(zoneIdentityIntf vpcv1.ZoneIdentityIntf) map[string]interface{} {
	zoneIdentityMap := map[string]interface{}{}
	zoneIdentity := zoneIdentityIntf.(*vpcv1.ZoneIdentity)
	zoneIdentityMap["name"] = zoneIdentity.Name
	zoneIdentityMap["href"] = zoneIdentity.Href
	return zoneIdentityMap
}

func resourceIBMIsBackupPolicyZoneIdentityByNameToMap(zoneIdentityByName vpcv1.ZoneIdentityByName) map[string]interface{} {
	zoneIdentityByNameMap := map[string]interface{}{}
	zoneIdentityByNameMap["name"] = zoneIdentityByName.Name
	return zoneIdentityByNameMap
}

func resourceIBMIsBackupPolicyZoneIdentityByHrefToMap(zoneIdentityByHref vpcv1.ZoneIdentityByHref) map[string]interface{} {
	zoneIdentityByHrefMap := map[string]interface{}{}
	zoneIdentityByHrefMap["href"] = zoneIdentityByHref.Href
	return zoneIdentityByHrefMap
}

func resourceIBMIsBackupPolicyBackupPolicyPlanDeletionTriggerPrototypeToMap(backupPolicyPlanDeletionTriggerPrototype vpcv1.BackupPolicyPlanDeletionTriggerPrototype) map[string]interface{} {
	backupPolicyPlanDeletionTriggerPrototypeMap := map[string]interface{}{}
	if backupPolicyPlanDeletionTriggerPrototype.DeleteAfter != nil {
		backupPolicyPlanDeletionTriggerPrototypeMap["delete_after"] = flex.IntValue(backupPolicyPlanDeletionTriggerPrototype.DeleteAfter)
	}
	if backupPolicyPlanDeletionTriggerPrototype.DeleteOverCount != nil {
		backupPolicyPlanDeletionTriggerPrototypeMap["delete_over_count"] = flex.IntValue(backupPolicyPlanDeletionTriggerPrototype.DeleteOverCount)
	}
	return backupPolicyPlanDeletionTriggerPrototypeMap
}

func resourceIBMIsBackupPolicyResourceGroupIdentityToMap(resourceGroupIdentity vpcv1.ResourceGroupReference) map[string]interface{} {
	resourceGroupIdentityMap := map[string]interface{}{}
	resourceGroupIdentityMap["id"] = resourceGroupIdentity.ID
	return resourceGroupIdentityMap
}

func resourceIBMIsBackupPolicyResourceGroupIdentityByIDToMap(resourceGroupIdentityByID vpcv1.ResourceGroupIdentityByID) map[string]interface{} {
	resourceGroupIdentityByIDMap := map[string]interface{}{}
	resourceGroupIdentityByIDMap["id"] = resourceGroupIdentityByID.ID
	return resourceGroupIdentityByIDMap
}

func resourceIBMIsBackupPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	updateBackupPolicyOptions := &vpcv1.UpdateBackupPolicyOptions{}
	updateBackupPolicyOptions.SetID(d.Id())
	hasChange := false
	patchVals := &vpcv1.BackupPolicyPatch{}
	if d.HasChange("match_user_tags") {
		patchVals.MatchUserTags = (flex.ExpandStringList((d.Get("match_user_tags").(*schema.Set)).List()))
		hasChange = true
	}
	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	updateBackupPolicyOptions.SetIfMatch(d.Get("version").(string))
	if hasChange {
		updateBackupPolicyOptions.BackupPolicyPatch, _ = patchVals.AsPatch()
		_, response, err := vpcClient.UpdateBackupPolicyWithContext(context, updateBackupPolicyOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateBackupPolicyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateBackupPolicyWithContext failed %s\n%s", err, response))
		}
	}

	if d.HasChange("plans") && !d.IsNewResource() {
		getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}
		getBackupPolicyOptions.SetID(d.Id())
		backupPolicy, response, err := vpcClient.GetBackupPolicyWithContext(context, getBackupPolicyOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] GetBackupPolicyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetBackupPolicyWithContext failed %s\n%s", err, response))
		}
		backupPolicyPlanID := backupPolicy.Plans[0].ID
		getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}
		getBackupPolicyPlanOptions.SetBackupPolicyID(d.Id())
		getBackupPolicyPlanOptions.SetID(*backupPolicyPlanID)
		_, response, err = vpcClient.GetBackupPolicyPlanWithContext(context, getBackupPolicyPlanOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] GetBackupPolicyPlanWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetBackupPolicyPlanWithContext failed %s\n%s", err, response))
		}
		etag := response.Headers.Get("E-tag")
		plans := d.Get("plans").([]interface{})
		for i := range plans {
			updateBackupPolicyPlanOptions := &vpcv1.UpdateBackupPolicyPlanOptions{}
			updateBackupPolicyPlanOptions.SetBackupPolicyID(d.Id())
			updateBackupPolicyPlanOptions.SetID(*backupPolicyPlanID)
			hasChange := false
			patchVals := &vpcv1.BackupPolicyPlanPatch{}

			cronSpec := fmt.Sprintf("plans.%d.cron_spec", i)
			if d.HasChange(cronSpec) {
				patchVals.CronSpec = core.StringPtr(d.Get(cronSpec).(string))
				hasChange = true
			}

			active := fmt.Sprintf("plans.%d.active", i)
			if d.HasChange(active) {
				patchVals.Active = core.BoolPtr(d.Get(active).(bool))
				hasChange = true
			}

			attachUserTags := fmt.Sprintf("plans.%d.attach_user_tags", i)
			if d.HasChange(attachUserTags) {
				patchVals.AttachUserTags = (flex.ExpandStringList((d.Get(attachUserTags).(*schema.Set)).List()))
				hasChange = true
			}

			// clonePolicy := fmt.Sprintf("plans.%d.clone_policy", i)
			// if d.HasChange(clonePolicy) {
			// 	clonePolicyItem := fmt.Sprintf("plans.%d.clone_policy.0", i)
			// 	clonePolicy := resourceIBMIsBackupPolicyPlanMapToBackupPolicyPlanClonePolicyPatch(d.Get(clonePolicyItem).(map[string]interface{}))
			// 	patchVals.ClonePolicy = &clonePolicy
			// 	hasChange = true
			// }

			copyUserTags := fmt.Sprintf("plans.%d.copy_user_tags", i)
			if d.HasChange(copyUserTags) {
				patchVals.CopyUserTags = core.BoolPtr(d.Get(copyUserTags).(bool))
				hasChange = true
			}

			deletionTrigger := fmt.Sprintf("plans.%d.deletion_trigger", i)
			if d.HasChange(deletionTrigger) {
				deletionTriggerItem := fmt.Sprintf("plans.%d.deletion_trigger.0", i)
				deletionTrigger := resourceIBMIsBackupPolicyPlanMapToBackupPolicyPlanDeletionTriggerPatch(d.Get(deletionTriggerItem).(map[string]interface{}))
				patchVals.DeletionTrigger = &deletionTrigger
				hasChange = true
			}

			name := fmt.Sprintf("plans.%d.name", i)
			if d.HasChange(name) {
				patchVals.Name = core.StringPtr(d.Get(name).(string))
				hasChange = true
			}
			updateBackupPolicyPlanOptions.SetIfMatch(etag)

			if hasChange {
				updateBackupPolicyPlanOptions.BackupPolicyPlanPatch, _ = patchVals.AsPatch()
				_, response, err := vpcClient.UpdateBackupPolicyPlanWithContext(context, updateBackupPolicyPlanOptions)
				if err != nil {
					log.Printf("[DEBUG] UpdateBackupPolicyPlanWithContext failed %s\n%s", err, response)
					return diag.FromErr(fmt.Errorf("UpdateBackupPolicyPlanWithContext failed %s\n%s", err, response))
				}
			}

		}

	}

	return resourceIBMIsBackupPolicyRead(context, d, meta)
}

func resourceIBMIsBackupPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	deleteBackupPolicyOptions := &vpcv1.DeleteBackupPolicyOptions{}
	deleteBackupPolicyOptions.SetID(d.Id())
	deleteBackupPolicyOptions.SetIfMatch(d.Get("version").(string))
	_, response, err := vpcClient.DeleteBackupPolicyWithContext(context, deleteBackupPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteBackupPolicyWithContext failed %s\n%s", err, response))
	}
	d.SetId("")
	return nil
}
