// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmSearchIndexedObject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSearchIndexedObjectCreate,
		ReadContext:   resourceIbmSearchIndexedObjectRead,
		DeleteContext: resourceIbmSearchIndexedObjectDelete,
		UpdateContext: resourceIbmSearchIndexedObjectUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"protection_group_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"storage_domain_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "TenantId contains id of the tenant for which objects are to be returned.",
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "\"This field is deprecated. Please use mightHaveTagIds.\".",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "\"This field is deprecated. Please use mightHaveSnapshotTagIds.\".",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"must_have_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies tags which must be all present in the document.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"might_have_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"must_have_snapshot_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies snapshot tags which must be all present in the document.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"might_have_snapshot_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"pagination_cookie": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the pagination cookie with which subsequent parts of the response can be fetched.",
			},
			"count": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the number of indexed objects to be fetched for the specified pagination cookie.",
			},
			"object_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_search_indexed_object", "object_type"),
				Description: "Specifies the object type to be searched for.",
			},
			"use_cached_data": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.",
			},
			"files": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the request parameters to search for files and file folders.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_string": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the search string to filter the files. User can specify a wildcard character '*' as a suffix to a string where all files name are matched with the prefix string.",
						},
						"types": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of file types. Only files within the given types will be returned.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"source_environments": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of the source environments. Only files from these types of source will be returned.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"source_ids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of source ids. Only files found in these sources will be returned.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"object_ids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of object ids. Only files found in these objects will be returned.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
			"public_folders": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the request parameters to search for Public Folder items.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_string": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the search string to filter the items. User can specify a wildcard character '*' as a suffix to a string where all item names are matched with the prefix string.",
						},
						"types": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of public folder item types. Only items within the given types will be returned.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"has_attachment": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Filters the public folder items which have attachment.",
						},
						"sender_address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Filters the public folder items which are received from specified user's email address.",
						},
						"recipient_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Filters the public folder items which are sent to specified email addresses.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"cc_recipient_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Filters the public folder items which are sent to specified email addresses in CC.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"bcc_recipient_addresses": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Filters the public folder items which are sent to specified email addresses in BCC.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"received_start_time_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the start time in Unix timestamp epoch in seconds where the received time of the public folder item is more than specified value.",
						},
						"received_end_time_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the end time in Unix timestamp epoch in seconds where the received time of the public folder items is less than specified value.",
						},
					},
				},
			},
			"public_folder_items": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the indexed Public folder items.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies tag applied to the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Id of tag applied to the object.",
									},
								},
							},
						},
						"snapshot_tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies snapshot tags applied to the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Id of tag applied to the object.",
									},
									"run_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies runs the tags are applied to.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the object.",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the path of the object.",
						},
						"protection_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "\"Specifies the protection group id which contains this object.\".",
						},
						"protection_group_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "\"Specifies the protection group name which contains this object.\".",
						},
						"policy_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the protection policy id for this file.",
						},
						"policy_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the protection policy name for this file.",
						},
						"storage_domain_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "\"Specifies the Storage Domain id where the backup data of Object is present.\".",
						},
						"source_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Source Object information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies object id.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the object.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies registered source id to which object belongs.",
									},
									"source_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies registered source name to which object belongs.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the environment of the object.",
									},
									"object_hash": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the hash identifier of the object.",
									},
									"object_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of the object.",
									},
									"logical_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the logical size of object in bytes.",
									},
									"uuid": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the uuid which is a unique identifier of the object.",
									},
									"global_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the global id which is a unique identifier of the object.",
									},
									"protection_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the protection type of the object if any.",
									},
									"os_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the operating system type of the object.",
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the Public folder item type.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the indexed item.",
						},
						"subject": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the subject of the indexed item.",
						},
						"has_attachments": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether the item has any attachments.",
						},
						"item_class": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the item class of the indexed item.",
						},
						"received_time_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the Unix timestamp epoch in seconds at which this item is received.",
						},
						"item_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the size in bytes for the indexed item.",
						},
						"parent_folder_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of parent folder the indexed item.",
						},
					},
				},
			},
		},
	}
}

func ResourceIbmSearchIndexedObjectValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "object_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "Files, PublicFolders",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_search_indexed_object", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSearchIndexedObjectCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	searchIndexedObjectsOptions := &backuprecoveryv1.SearchIndexedObjectsOptions{}

	searchIndexedObjectsOptions.SetObjectType(d.Get("object_type").(string))
	if _, ok := d.GetOk("protection_group_ids"); ok {
		var protectionGroupIds []string
		for _, v := range d.Get("protection_group_ids").([]interface{}) {
			protectionGroupIdsItem := v.(string)
			protectionGroupIds = append(protectionGroupIds, protectionGroupIdsItem)
		}
		searchIndexedObjectsOptions.SetProtectionGroupIds(protectionGroupIds)
	}
	if _, ok := d.GetOk("storage_domain_ids"); ok {
		var storageDomainIds []int64
		for _, v := range d.Get("storage_domain_ids").([]interface{}) {
			storageDomainIdsItem := int64(v.(int))
			storageDomainIds = append(storageDomainIds, storageDomainIdsItem)
		}
		searchIndexedObjectsOptions.SetStorageDomainIds(storageDomainIds)
	}
	if _, ok := d.GetOk("tenant_id"); ok {
		searchIndexedObjectsOptions.SetTenantID(d.Get("tenant_id").(string))
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		searchIndexedObjectsOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("tags"); ok {
		var tags []string
		for _, v := range d.Get("tags").([]interface{}) {
			tagsItem := v.(string)
			tags = append(tags, tagsItem)
		}
		searchIndexedObjectsOptions.SetTags(tags)
	}
	if _, ok := d.GetOk("snapshot_tags"); ok {
		var snapshotTags []string
		for _, v := range d.Get("snapshot_tags").([]interface{}) {
			snapshotTagsItem := v.(string)
			snapshotTags = append(snapshotTags, snapshotTagsItem)
		}
		searchIndexedObjectsOptions.SetSnapshotTags(snapshotTags)
	}
	if _, ok := d.GetOk("must_have_tag_ids"); ok {
		var mustHaveTagIds []string
		for _, v := range d.Get("must_have_tag_ids").([]interface{}) {
			mustHaveTagIdsItem := v.(string)
			mustHaveTagIds = append(mustHaveTagIds, mustHaveTagIdsItem)
		}
		searchIndexedObjectsOptions.SetMustHaveTagIds(mustHaveTagIds)
	}
	if _, ok := d.GetOk("might_have_tag_ids"); ok {
		var mightHaveTagIds []string
		for _, v := range d.Get("might_have_tag_ids").([]interface{}) {
			mightHaveTagIdsItem := v.(string)
			mightHaveTagIds = append(mightHaveTagIds, mightHaveTagIdsItem)
		}
		searchIndexedObjectsOptions.SetMightHaveTagIds(mightHaveTagIds)
	}
	if _, ok := d.GetOk("must_have_snapshot_tag_ids"); ok {
		var mustHaveSnapshotTagIds []string
		for _, v := range d.Get("must_have_snapshot_tag_ids").([]interface{}) {
			mustHaveSnapshotTagIdsItem := v.(string)
			mustHaveSnapshotTagIds = append(mustHaveSnapshotTagIds, mustHaveSnapshotTagIdsItem)
		}
		searchIndexedObjectsOptions.SetMustHaveSnapshotTagIds(mustHaveSnapshotTagIds)
	}
	if _, ok := d.GetOk("might_have_snapshot_tag_ids"); ok {
		var mightHaveSnapshotTagIds []string
		for _, v := range d.Get("might_have_snapshot_tag_ids").([]interface{}) {
			mightHaveSnapshotTagIdsItem := v.(string)
			mightHaveSnapshotTagIds = append(mightHaveSnapshotTagIds, mightHaveSnapshotTagIdsItem)
		}
		searchIndexedObjectsOptions.SetMightHaveSnapshotTagIds(mightHaveSnapshotTagIds)
	}
	if _, ok := d.GetOk("pagination_cookie"); ok {
		searchIndexedObjectsOptions.SetPaginationCookie(d.Get("pagination_cookie").(string))
	}
	if _, ok := d.GetOk("count"); ok {
		searchIndexedObjectsOptions.SetCount(int64(d.Get("count").(int)))
	}
	if _, ok := d.GetOk("use_cached_data"); ok {
		searchIndexedObjectsOptions.SetUseCachedData(d.Get("use_cached_data").(bool))
	}
	if _, ok := d.GetOk("files"); ok {
		filesModel, err := resourceIbmSearchIndexedObjectMapToSearchFileRequestParams(d.Get("files.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		searchIndexedObjectsOptions.SetFiles(filesModel)
	}
	if _, ok := d.GetOk("public_folders"); ok {
		publicFoldersModel, err := resourceIbmSearchIndexedObjectMapToSearchPublicFolderRequestParams(d.Get("public_folders.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		searchIndexedObjectsOptions.SetPublicFolders(publicFoldersModel)
	}

	searchIndexedObjectsResponse, response, err := backupRecoveryClient.SearchIndexedObjectsWithContext(context, searchIndexedObjectsOptions)
	if err != nil {
		log.Printf("[DEBUG] SearchIndexedObjectsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("SearchIndexedObjectsWithContext failed %s\n%s", err, response))
	}

	d.SetId(resourceIbmSearchIndexedObjectID(d))

	if err = d.Set("object_type", searchIndexedObjectsResponse.ObjectType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting object_type: %s", err))
	}
	if !core.IsNil(searchIndexedObjectsResponse.Count) {
		if err = d.Set("count", flex.IntValue(searchIndexedObjectsResponse.Count)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting count: %s", err))
		}
	}
	if !core.IsNil(searchIndexedObjectsResponse.PaginationCookie) {
		if err = d.Set("pagination_cookie", searchIndexedObjectsResponse.PaginationCookie); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting pagination_cookie: %s", err))
		}
	}
	if !core.IsNil(searchIndexedObjectsResponse.Files) {
		files := []map[string]interface{}{}
		for _, filesItem := range searchIndexedObjectsResponse.Files {
			filesItemMap, err := resourceIbmSearchIndexedObjectFileToMap(&filesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			files = append(files, filesItemMap)
		}
		if err = d.Set("files", files); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting files: %s", err))
		}
	}
	if !core.IsNil(searchIndexedObjectsResponse.PublicFolderItems) {
		publicFolderItems := []map[string]interface{}{}
		for _, publicFolderItemsItem := range searchIndexedObjectsResponse.PublicFolderItems {
			publicFolderItemsItemMap, err := resourceIbmSearchIndexedObjectPublicFolderItemToMap(&publicFolderItemsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			publicFolderItems = append(publicFolderItems, publicFolderItemsItemMap)
		}
		if err = d.Set("public_folder_items", publicFolderItems); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting public_folder_items: %s", err))
		}
	}

	return nil

	// return resourceIbmSearchIndexedObjectRead(context, d, meta)
}

func resourceIbmSearchIndexedObjectID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmSearchIndexedObjectRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}
	getRecoveryByIdOptions.SetID("")
	return nil
}

func resourceIbmSearchIndexedObjectDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmSearchIndexedObjectUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "Update operation is not supported for this resource. No changes will be applied.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmSearchIndexedObjectPublicFolderItemToMap(model *backuprecoveryv1.PublicFolderItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Tags != nil {
		tags := []map[string]interface{}{}
		for _, tagsItem := range model.Tags {
			tagsItemMap, err := resourceIbmSearchIndexedObjectTagInfoToMap(&tagsItem)
			if err != nil {
				return modelMap, err
			}
			tags = append(tags, tagsItemMap)
		}
		modelMap["tags"] = tags
	}
	if model.SnapshotTags != nil {
		snapshotTags := []map[string]interface{}{}
		for _, snapshotTagsItem := range model.SnapshotTags {
			snapshotTagsItemMap, err := resourceIbmSearchIndexedObjectSnapshotTagInfoToMap(&snapshotTagsItem)
			if err != nil {
				return modelMap, err
			}
			snapshotTags = append(snapshotTags, snapshotTagsItemMap)
		}
		modelMap["snapshot_tags"] = snapshotTags
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = model.PolicyID
	}
	if model.PolicyName != nil {
		modelMap["policy_name"] = model.PolicyName
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.SourceInfo != nil {
		sourceInfoMap, err := resourceIbmSearchIndexedObjectPublicFolderItemSourceInfoToMap(model.SourceInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["source_info"] = []map[string]interface{}{sourceInfoMap}
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Subject != nil {
		modelMap["subject"] = model.Subject
	}
	if model.HasAttachments != nil {
		modelMap["has_attachments"] = model.HasAttachments
	}
	if model.ItemClass != nil {
		modelMap["item_class"] = model.ItemClass
	}
	if model.ReceivedTimeSecs != nil {
		modelMap["received_time_secs"] = flex.IntValue(model.ReceivedTimeSecs)
	}
	if model.ItemSize != nil {
		modelMap["item_size"] = flex.IntValue(model.ItemSize)
	}
	if model.ParentFolderID != nil {
		modelMap["parent_folder_id"] = model.ParentFolderID
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectPublicFolderItemSourceInfoToMap(model *backuprecoveryv1.PublicFolderItemSourceInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = model.ProtectionType
	}
	if model.OsType != nil {
		modelMap["os_type"] = model.OsType
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectSnapshotTagInfoToMap(model *backuprecoveryv1.SnapshotTagInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["tag_id"] = model.TagID
	if model.RunIds != nil {
		modelMap["run_ids"] = model.RunIds
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectFileToMap(model *backuprecoveryv1.File) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Tags != nil {
		tags := []map[string]interface{}{}
		for _, tagsItem := range model.Tags {
			tagsItemMap, err := resourceIbmSearchIndexedObjectTagInfoToMap(&tagsItem)
			if err != nil {
				return modelMap, err
			}
			tags = append(tags, tagsItemMap)
		}
		modelMap["tags"] = tags
	}
	if model.SnapshotTags != nil {
		snapshotTags := []map[string]interface{}{}
		for _, snapshotTagsItem := range model.SnapshotTags {
			snapshotTagsItemMap, err := resourceIbmSearchIndexedObjectSnapshotTagInfoToMap(&snapshotTagsItem)
			if err != nil {
				return modelMap, err
			}
			snapshotTags = append(snapshotTags, snapshotTagsItemMap)
		}
		modelMap["snapshot_tags"] = snapshotTags
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = model.PolicyID
	}
	if model.PolicyName != nil {
		modelMap["policy_name"] = model.PolicyName
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.SourceInfo != nil {
		sourceInfoMap, err := resourceIbmSearchIndexedObjectFileSourceInfoToMap(model.SourceInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["source_info"] = []map[string]interface{}{sourceInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectTagInfoToMap(model *backuprecoveryv1.TagInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["tag_id"] = model.TagID
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectFileSourceInfoToMap(model *backuprecoveryv1.FileSourceInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = model.ProtectionType
	}
	if model.OsType != nil {
		modelMap["os_type"] = model.OsType
	}
	if model.ProtectionStats != nil {
		protectionStats := []map[string]interface{}{}
		for _, protectionStatsItem := range model.ProtectionStats {
			protectionStatsItemMap, err := resourceIbmSearchIndexedObjectObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := resourceIbmSearchIndexedObjectPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := resourceIbmSearchIndexedObjectObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := resourceIbmSearchIndexedObjectObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectObjectPhysicalParamsToMap(model *backuprecoveryv1.ObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectObjectOracleParamsToMap(model *backuprecoveryv1.ObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := resourceIbmSearchIndexedObjectDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := resourceIbmSearchIndexedObjectHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := resourceIbmSearchIndexedObjectPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := resourceIbmSearchIndexedObjectDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.DeletedProtectedCount != nil {
		modelMap["deleted_protected_count"] = flex.IntValue(model.DeletedProtectedCount)
	}
	if model.ProtectedSizeBytes != nil {
		modelMap["protected_size_bytes"] = flex.IntValue(model.ProtectedSizeBytes)
	}
	if model.UnprotectedSizeBytes != nil {
		modelMap["unprotected_size_bytes"] = flex.IntValue(model.UnprotectedSizeBytes)
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := resourceIbmSearchIndexedObjectUserToMap(&usersItem)
			if err != nil {
				return modelMap, err
			}
			users = append(users, usersItemMap)
		}
		modelMap["users"] = users
	}
	if model.Groups != nil {
		groups := []map[string]interface{}{}
		for _, groupsItem := range model.Groups {
			groupsItemMap, err := resourceIbmSearchIndexedObjectGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := resourceIbmSearchIndexedObjectTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = model.Domain
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = model.Domain
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectMapToSearchFileRequestParams(modelMap map[string]interface{}) (*backuprecoveryv1.SearchFileRequestParams, error) {
	model := &backuprecoveryv1.SearchFileRequestParams{}
	if modelMap["search_string"] != nil && modelMap["search_string"].(string) != "" {
		model.SearchString = core.StringPtr(modelMap["search_string"].(string))
	}
	if modelMap["types"] != nil {
		types := []string{}
		for _, typesItem := range modelMap["types"].([]interface{}) {
			types = append(types, typesItem.(string))
		}
		model.Types = types
	}
	if modelMap["source_environments"] != nil {
		sourceEnvironments := []string{}
		for _, sourceEnvironmentsItem := range modelMap["source_environments"].([]interface{}) {
			sourceEnvironments = append(sourceEnvironments, sourceEnvironmentsItem.(string))
		}
		model.SourceEnvironments = sourceEnvironments
	}
	if modelMap["source_ids"] != nil {
		sourceIds := []int64{}
		for _, sourceIdsItem := range modelMap["source_ids"].([]interface{}) {
			sourceIds = append(sourceIds, int64(sourceIdsItem.(int)))
		}
		model.SourceIds = sourceIds
	}
	if modelMap["object_ids"] != nil {
		objectIds := []int64{}
		for _, objectIdsItem := range modelMap["object_ids"].([]interface{}) {
			objectIds = append(objectIds, int64(objectIdsItem.(int)))
		}
		model.ObjectIds = objectIds
	}
	return model, nil
}

func resourceIbmSearchIndexedObjectMapToSearchPublicFolderRequestParams(modelMap map[string]interface{}) (*backuprecoveryv1.SearchPublicFolderRequestParams, error) {
	model := &backuprecoveryv1.SearchPublicFolderRequestParams{}
	if modelMap["search_string"] != nil && modelMap["search_string"].(string) != "" {
		model.SearchString = core.StringPtr(modelMap["search_string"].(string))
	}
	if modelMap["types"] != nil {
		types := []string{}
		for _, typesItem := range modelMap["types"].([]interface{}) {
			types = append(types, typesItem.(string))
		}
		model.Types = types
	}
	if modelMap["has_attachment"] != nil {
		model.HasAttachment = core.BoolPtr(modelMap["has_attachment"].(bool))
	}
	if modelMap["sender_address"] != nil && modelMap["sender_address"].(string) != "" {
		model.SenderAddress = core.StringPtr(modelMap["sender_address"].(string))
	}
	if modelMap["recipient_addresses"] != nil {
		recipientAddresses := []string{}
		for _, recipientAddressesItem := range modelMap["recipient_addresses"].([]interface{}) {
			recipientAddresses = append(recipientAddresses, recipientAddressesItem.(string))
		}
		model.RecipientAddresses = recipientAddresses
	}
	if modelMap["cc_recipient_addresses"] != nil {
		ccRecipientAddresses := []string{}
		for _, ccRecipientAddressesItem := range modelMap["cc_recipient_addresses"].([]interface{}) {
			ccRecipientAddresses = append(ccRecipientAddresses, ccRecipientAddressesItem.(string))
		}
		model.CcRecipientAddresses = ccRecipientAddresses
	}
	if modelMap["bcc_recipient_addresses"] != nil {
		bccRecipientAddresses := []string{}
		for _, bccRecipientAddressesItem := range modelMap["bcc_recipient_addresses"].([]interface{}) {
			bccRecipientAddresses = append(bccRecipientAddresses, bccRecipientAddressesItem.(string))
		}
		model.BccRecipientAddresses = bccRecipientAddresses
	}
	if modelMap["received_start_time_secs"] != nil {
		model.ReceivedStartTimeSecs = core.Int64Ptr(int64(modelMap["received_start_time_secs"].(int)))
	}
	if modelMap["received_end_time_secs"] != nil {
		model.ReceivedEndTimeSecs = core.Int64Ptr(int64(modelMap["received_end_time_secs"].(int)))
	}
	return model, nil
}

func resourceIbmSearchIndexedObjectSearchFileRequestParamsToMap(model *backuprecoveryv1.SearchFileRequestParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SearchString != nil {
		modelMap["search_string"] = model.SearchString
	}
	if model.Types != nil {
		modelMap["types"] = model.Types
	}
	if model.SourceEnvironments != nil {
		modelMap["source_environments"] = model.SourceEnvironments
	}
	if model.SourceIds != nil {
		modelMap["source_ids"] = model.SourceIds
	}
	if model.ObjectIds != nil {
		modelMap["object_ids"] = model.ObjectIds
	}
	return modelMap, nil
}

func resourceIbmSearchIndexedObjectSearchPublicFolderRequestParamsToMap(model *backuprecoveryv1.SearchPublicFolderRequestParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SearchString != nil {
		modelMap["search_string"] = model.SearchString
	}
	if model.Types != nil {
		modelMap["types"] = model.Types
	}
	if model.HasAttachment != nil {
		modelMap["has_attachment"] = model.HasAttachment
	}
	if model.SenderAddress != nil {
		modelMap["sender_address"] = model.SenderAddress
	}
	if model.RecipientAddresses != nil {
		modelMap["recipient_addresses"] = model.RecipientAddresses
	}
	if model.CcRecipientAddresses != nil {
		modelMap["cc_recipient_addresses"] = model.CcRecipientAddresses
	}
	if model.BccRecipientAddresses != nil {
		modelMap["bcc_recipient_addresses"] = model.BccRecipientAddresses
	}
	if model.ReceivedStartTimeSecs != nil {
		modelMap["received_start_time_secs"] = flex.IntValue(model.ReceivedStartTimeSecs)
	}
	if model.ReceivedEndTimeSecs != nil {
		modelMap["received_end_time_secs"] = flex.IntValue(model.ReceivedEndTimeSecs)
	}
	return modelMap, nil
}
