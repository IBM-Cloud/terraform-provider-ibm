// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmRecoveryDownloadFilesFolders() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmRecoveryDownloadFilesFoldersCreate,
		ReadContext:   resourceIbmRecoveryDownloadFilesFoldersRead,
		DeleteContext: resourceIbmRecoveryDownloadFilesFoldersDelete,
		UpdateContext: resourceIbmRecoveryDownloadFilesFoldersUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"documents": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_directory": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether the document is a directory. Since currently only files are supported this should always be false.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the item id of the document.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "Specifies the name of the recovery task. This field must be set and must be a unique name.",
			},
			"object": &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				// ForceNew:    true,
				Description: "Specifies the common snapshot parameters for a protected object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"snapshot_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the snapshot id.",
						},
						"point_in_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.",
						},
						"protection_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the protection group id of the object snapshot.",
						},
						"protection_group_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the protection group name of the object snapshot.",
						},
						"snapshot_creation_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.",
						},
						"object_info": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the information about the object for which the snapshot is taken.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies object id.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the name of the object.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies registered source id to which object belongs.",
									},
									"source_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies registered source name to which object belongs.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the environment of the object.",
									},
									"object_hash": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the hash identifier of the object.",
									},
									"object_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the type of the object.",
									},
									"logical_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the logical size of object in bytes.",
									},
									"uuid": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the uuid which is a unique identifier of the object.",
									},
									"global_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the global id which is a unique identifier of the object.",
									},
									"protection_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the protection type of the object if any.",
									},
									"sharepoint_site_summary": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the common parameters for Sharepoint site objects.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"site_web_url": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the web url for the Sharepoint site.",
												},
											},
										},
									},
									"os_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the operating system type of the object.",
									},
									"child_objects": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies child object details.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies object id.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the name of the object.",
												},
												"source_id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies registered source id to which object belongs.",
												},
												"source_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies registered source name to which object belongs.",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the environment of the object.",
												},
												"object_hash": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the hash identifier of the object.",
												},
												"object_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the type of the object.",
												},
												"logical_size_bytes": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the logical size of object in bytes.",
												},
												"uuid": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the uuid which is a unique identifier of the object.",
												},
												"global_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the global id which is a unique identifier of the object.",
												},
												"protection_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the protection type of the object if any.",
												},
												"sharepoint_site_summary": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the common parameters for Sharepoint site objects.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"site_web_url": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the web url for the Sharepoint site.",
															},
														},
													},
												},
												"os_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the operating system type of the object.",
												},
												"child_objects": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies child object details.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},
												"v_center_summary": &schema.Schema{
													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_cloud_env": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.",
															},
														},
													},
												},
												"windows_cluster_summary": &schema.Schema{
													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_source_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the type of cluster resource this source represents.",
															},
														},
													},
												},
											},
										},
									},
									"v_center_summary": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_cloud_env": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.",
												},
											},
										},
									},
									"windows_cluster_summary": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_source_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the type of cluster resource this source represents.",
												},
											},
										},
									},
								},
							},
						},
						"snapshot_target_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the snapshot target type.",
						},
						"storage_domain_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the ID of the Storage Domain where this snapshot is stored.",
						},
						"archival_target_info": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the archival target information if the snapshot is an archival snapshot.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies the archival target ID.",
									},
									"archival_task_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.",
									},
									"target_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the archival target name.",
									},
									"target_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the archival target type.",
									},
									"usage_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the usage type for the target.",
									},
									"ownership_context": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the ownership context for the target.",
									},
									"tier_settings": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the tier info for archival.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aws_tiering": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies aws tiers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tiers": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"move_after_unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																		},
																		"tier_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the AWS tier types.",
																		},
																	},
																},
															},
														},
													},
												},
												"azure_tiering": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies Azure tiers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tiers": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"move_after_unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																		},
																		"tier_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the Azure tier types.",
																		},
																	},
																},
															},
														},
													},
												},
												"cloud_platform": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the cloud platform to enable tiering.",
												},
												"google_tiering": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies Google tiers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tiers": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"move_after_unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																		},
																		"tier_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the Google tier types.",
																		},
																	},
																},
															},
														},
													},
												},
												"oracle_tiering": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies Oracle tiers.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tiers": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"move_after_unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																			Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																		},
																		"tier_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the Oracle tier types.",
																		},
																	},
																},
															},
														},
													},
												},
												"current_tier_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.",
												},
											},
										},
									},
								},
							},
						},
						"progress_task_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Progress monitor task id for Recovery of VM.",
						},
						"recover_from_standby": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies that user wants to perform standby restore if it is enabled for this object.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.",
						},
						"start_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.",
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.",
						},
						"messages": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specify error messages about the object.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"bytes_restored": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specify the total bytes restored.",
						},
					},
				},
			},
			"parent_recovery_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery_download_files_folders", "parent_recovery_id"),
				Description: "If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery.",
			},
			"files_and_folders": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the list of files and folders to download.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"absolute_path": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the absolute path of the file or folder.",
						},
						"is_directory": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether the file or folder object is a directory.",
						},
					},
				},
			},
			"glacier_retrieval_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery_download_files_folders", "glacier_retrieval_type"),
				Description: "Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot.",
			},
		},
	}
}

func ResourceIbmRecoveryDownloadFilesFoldersValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "parent_recovery_id",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^\d+:\d+:\d+$`,
		},
		validate.ValidateSchema{
			Identifier:                 "glacier_retrieval_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "kExpeditedNoPCU, kExpeditedWithPCU, kStandard",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_recovery_download_files_folders", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmRecoveryDownloadFilesFoldersCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_recovery_download_files_folders", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDownloadFilesAndFoldersRecoveryOptions := &backuprecoveryv1.CreateDownloadFilesAndFoldersRecoveryOptions{}

	createDownloadFilesAndFoldersRecoveryOptions.SetName(d.Get("name").(string))
	objectModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(d.Get("object.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_recovery_download_files_folders", "create", "parse-object").GetDiag()
	}
	createDownloadFilesAndFoldersRecoveryOptions.SetObject(objectModel)
	var filesAndFolders []backuprecoveryv1.FilesAndFoldersObject
	for _, v := range d.Get("files_and_folders").([]interface{}) {
		value := v.(map[string]interface{})
		filesAndFoldersItem, err := ResourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_recovery_download_files_folders", "create", "parse-files_and_folders").GetDiag()
		}
		filesAndFolders = append(filesAndFolders, *filesAndFoldersItem)
	}
	createDownloadFilesAndFoldersRecoveryOptions.SetFilesAndFolders(filesAndFolders)
	if _, ok := d.GetOk("documents"); ok {
		var documents []backuprecoveryv1.DocumentObject
		for _, v := range d.Get("documents").([]interface{}) {
			value := v.(map[string]interface{})
			documentsItem, err := ResourceIbmRecoveryDownloadFilesFoldersMapToDocumentObject(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_recovery_download_files_folders", "create", "parse-documents").GetDiag()
			}
			documents = append(documents, *documentsItem)
		}
		createDownloadFilesAndFoldersRecoveryOptions.SetDocuments(documents)
	}
	if _, ok := d.GetOk("parent_recovery_id"); ok {
		createDownloadFilesAndFoldersRecoveryOptions.SetParentRecoveryID(d.Get("parent_recovery_id").(string))
	}
	if _, ok := d.GetOk("glacier_retrieval_type"); ok {
		createDownloadFilesAndFoldersRecoveryOptions.SetGlacierRetrievalType(d.Get("glacier_retrieval_type").(string))
	}

	recovery, _, err := backupRecoveryClient.CreateDownloadFilesAndFoldersRecoveryWithContext(context, createDownloadFilesAndFoldersRecoveryOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDownloadFilesAndFoldersRecoveryWithContext failed: %s", err.Error()), "ibm_recovery_download_files_folders", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*recovery.ID)

	return resourceIbmRecoveryDownloadFilesFoldersRead(context, d, meta)
}

func resourceIbmRecoveryDownloadFilesFoldersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_recovery_download_files_folders", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

	getRecoveryByIdOptions.SetID(d.Id())

	recovery, response, err := backupRecoveryClient.GetRecoveryByIDWithContext(context, getRecoveryByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetRecoveryByIDWithContext failed: %s", err.Error()), "ibm_recovery_download_files_folders", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getRecoveryByIdOptions.ID)

	if !core.IsNil(recovery.Name) {
		if err = d.Set("name", recovery.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_recovery", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(recovery.StartTimeUsecs) {
		if err = d.Set("start_time_usecs", flex.IntValue(recovery.StartTimeUsecs)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting start_time_usecs: %s", err), "(Data) ibm_recovery", "read", "set-start_time_usecs").GetDiag()
		}
	}

	if !core.IsNil(recovery.EndTimeUsecs) {
		if err = d.Set("end_time_usecs", flex.IntValue(recovery.EndTimeUsecs)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting end_time_usecs: %s", err), "(Data) ibm_recovery", "read", "set-end_time_usecs").GetDiag()
		}
	}

	if !core.IsNil(recovery.Status) {
		if err = d.Set("status", recovery.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_recovery", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(recovery.ProgressTaskID) {
		if err = d.Set("progress_task_id", recovery.ProgressTaskID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting progress_task_id: %s", err), "(Data) ibm_recovery", "read", "set-progress_task_id").GetDiag()
		}
	}

	if !core.IsNil(recovery.SnapshotEnvironment) {
		if err = d.Set("snapshot_environment", recovery.SnapshotEnvironment); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting snapshot_environment: %s", err), "(Data) ibm_recovery", "read", "set-snapshot_environment").GetDiag()
		}
	}

	if !core.IsNil(recovery.RecoveryAction) {
		if err = d.Set("recovery_action", recovery.RecoveryAction); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting recovery_action: %s", err), "(Data) ibm_recovery", "read", "set-recovery_action").GetDiag()
		}
	}

	if !core.IsNil(recovery.Permissions) {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range recovery.Permissions {
			permissionsItemMap, err := DataSourceIbmRecoveryTenantToMap(&permissionsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_recovery", "read", "permissions-to-map").GetDiag()
			}
			permissions = append(permissions, permissionsItemMap)
		}
		if err = d.Set("permissions", permissions); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting permissions: %s", err), "(Data) ibm_recovery", "read", "set-permissions").GetDiag()
		}
	}

	if !core.IsNil(recovery.CreationInfo) {
		creationInfo := []map[string]interface{}{}
		creationInfoMap, err := DataSourceIbmRecoveryCreationInfoToMap(recovery.CreationInfo)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_recovery", "read", "creation_info-to-map").GetDiag()
		}
		creationInfo = append(creationInfo, creationInfoMap)
		if err = d.Set("creation_info", creationInfo); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting creation_info: %s", err), "(Data) ibm_recovery", "read", "set-creation_info").GetDiag()
		}
	}

	if !core.IsNil(recovery.CanTearDown) {
		if err = d.Set("can_tear_down", recovery.CanTearDown); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting can_tear_down: %s", err), "(Data) ibm_recovery", "read", "set-can_tear_down").GetDiag()
		}
	}

	if !core.IsNil(recovery.TearDownStatus) {
		if err = d.Set("tear_down_status", recovery.TearDownStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tear_down_status: %s", err), "(Data) ibm_recovery", "read", "set-tear_down_status").GetDiag()
		}
	}

	if !core.IsNil(recovery.TearDownMessage) {
		if err = d.Set("tear_down_message", recovery.TearDownMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tear_down_message: %s", err), "(Data) ibm_recovery", "read", "set-tear_down_message").GetDiag()
		}
	}

	if !core.IsNil(recovery.Messages) {
		messages := []interface{}{}
		for _, messagesItem := range recovery.Messages {
			messages = append(messages, messagesItem)
		}
		if err = d.Set("messages", messages); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting messages: %s", err), "(Data) ibm_recovery", "read", "set-messages").GetDiag()
		}
	}

	if !core.IsNil(recovery.IsParentRecovery) {
		if err = d.Set("is_parent_recovery", recovery.IsParentRecovery); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_parent_recovery: %s", err), "(Data) ibm_recovery", "read", "set-is_parent_recovery").GetDiag()
		}
	}

	if !core.IsNil(recovery.ParentRecoveryID) {
		if err = d.Set("parent_recovery_id", recovery.ParentRecoveryID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting parent_recovery_id: %s", err), "(Data) ibm_recovery", "read", "set-parent_recovery_id").GetDiag()
		}
	}

	if !core.IsNil(recovery.RetrieveArchiveTasks) {
		retrieveArchiveTasks := []map[string]interface{}{}
		for _, retrieveArchiveTasksItem := range recovery.RetrieveArchiveTasks {
			retrieveArchiveTasksItemMap, err := DataSourceIbmRecoveryRetrieveArchiveTaskToMap(&retrieveArchiveTasksItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_recovery", "read", "retrieve_archive_tasks-to-map").GetDiag()
			}
			retrieveArchiveTasks = append(retrieveArchiveTasks, retrieveArchiveTasksItemMap)
		}
		if err = d.Set("retrieve_archive_tasks", retrieveArchiveTasks); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting retrieve_archive_tasks: %s", err), "(Data) ibm_recovery", "read", "set-retrieve_archive_tasks").GetDiag()
		}
	}

	if !core.IsNil(recovery.IsMultiStageRestore) {
		if err = d.Set("is_multi_stage_restore", recovery.IsMultiStageRestore); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_multi_stage_restore: %s", err), "(Data) ibm_recovery", "read", "set-is_multi_stage_restore").GetDiag()
		}
	}

	if !core.IsNil(recovery.PhysicalParams) {
		physicalParams := []map[string]interface{}{}
		physicalParamsMap, err := DataSourceIbmRecoveryRecoverPhysicalParamsToMap(recovery.PhysicalParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_recovery", "read", "physical_params-to-map").GetDiag()
		}
		physicalParams = append(physicalParams, physicalParamsMap)
		if err = d.Set("physical_params", physicalParams); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting physical_params: %s", err), "(Data) ibm_recovery", "read", "set-physical_params").GetDiag()
		}
	}

	if !core.IsNil(recovery.MssqlParams) {
		mssqlParams := []map[string]interface{}{}
		mssqlParamsMap, err := DataSourceIbmRecoveryRecoverSqlParamsToMap(recovery.MssqlParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_recovery", "read", "mssql_params-to-map").GetDiag()
		}
		mssqlParams = append(mssqlParams, mssqlParamsMap)
		if err = d.Set("mssql_params", mssqlParams); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mssql_params: %s", err), "(Data) ibm_recovery", "read", "set-mssql_params").GetDiag()
		}
	}

	return nil
}

func resourceIbmRecoveryDownloadFilesFoldersDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryDownloadFilesFoldersUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParams, error) {
	model := &backuprecoveryv1.CommonRecoverObjectSnapshotParams{}
	model.SnapshotID = core.StringPtr(modelMap["snapshot_id"].(string))
	if modelMap["point_in_time_usecs"] != nil {
		model.PointInTimeUsecs = core.Int64Ptr(int64(modelMap["point_in_time_usecs"].(int)))
	}
	if modelMap["protection_group_id"] != nil && modelMap["protection_group_id"].(string) != "" {
		model.ProtectionGroupID = core.StringPtr(modelMap["protection_group_id"].(string))
	}
	if modelMap["protection_group_name"] != nil && modelMap["protection_group_name"].(string) != "" {
		model.ProtectionGroupName = core.StringPtr(modelMap["protection_group_name"].(string))
	}
	if modelMap["snapshot_creation_time_usecs"] != nil {
		model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(modelMap["snapshot_creation_time_usecs"].(int)))
	}
	if modelMap["object_info"] != nil && len(modelMap["object_info"].([]interface{})) > 0 {
		ObjectInfoModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap["object_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ObjectInfo = ObjectInfoModel
	}
	if modelMap["snapshot_target_type"] != nil && modelMap["snapshot_target_type"].(string) != "" {
		model.SnapshotTargetType = core.StringPtr(modelMap["snapshot_target_type"].(string))
	}
	if modelMap["storage_domain_id"] != nil {
		model.StorageDomainID = core.Int64Ptr(int64(modelMap["storage_domain_id"].(int)))
	}
	if modelMap["archival_target_info"] != nil && len(modelMap["archival_target_info"].([]interface{})) > 0 {
		ArchivalTargetInfoModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap["archival_target_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ArchivalTargetInfo = ArchivalTargetInfoModel
	}
	if modelMap["progress_task_id"] != nil && modelMap["progress_task_id"].(string) != "" {
		model.ProgressTaskID = core.StringPtr(modelMap["progress_task_id"].(string))
	}
	if modelMap["recover_from_standby"] != nil {
		model.RecoverFromStandby = core.BoolPtr(modelMap["recover_from_standby"].(bool))
	}
	if modelMap["status"] != nil && modelMap["status"].(string) != "" {
		model.Status = core.StringPtr(modelMap["status"].(string))
	}
	if modelMap["start_time_usecs"] != nil {
		model.StartTimeUsecs = core.Int64Ptr(int64(modelMap["start_time_usecs"].(int)))
	}
	if modelMap["end_time_usecs"] != nil {
		model.EndTimeUsecs = core.Int64Ptr(int64(modelMap["end_time_usecs"].(int)))
	}
	if modelMap["messages"] != nil {
		messages := []string{}
		for _, messagesItem := range modelMap["messages"].([]interface{}) {
			messages = append(messages, messagesItem.(string))
		}
		model.Messages = messages
	}
	if modelMap["bytes_restored"] != nil {
		model.BytesRestored = core.Int64Ptr(int64(modelMap["bytes_restored"].(int)))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo, error) {
	model := &backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["source_id"] != nil {
		model.SourceID = core.Int64Ptr(int64(modelMap["source_id"].(int)))
	}
	if modelMap["source_name"] != nil && modelMap["source_name"].(string) != "" {
		model.SourceName = core.StringPtr(modelMap["source_name"].(string))
	}
	if modelMap["environment"] != nil && modelMap["environment"].(string) != "" {
		model.Environment = core.StringPtr(modelMap["environment"].(string))
	}
	if modelMap["object_hash"] != nil && modelMap["object_hash"].(string) != "" {
		model.ObjectHash = core.StringPtr(modelMap["object_hash"].(string))
	}
	if modelMap["object_type"] != nil && modelMap["object_type"].(string) != "" {
		model.ObjectType = core.StringPtr(modelMap["object_type"].(string))
	}
	if modelMap["logical_size_bytes"] != nil {
		model.LogicalSizeBytes = core.Int64Ptr(int64(modelMap["logical_size_bytes"].(int)))
	}
	if modelMap["uuid"] != nil && modelMap["uuid"].(string) != "" {
		model.UUID = core.StringPtr(modelMap["uuid"].(string))
	}
	if modelMap["global_id"] != nil && modelMap["global_id"].(string) != "" {
		model.GlobalID = core.StringPtr(modelMap["global_id"].(string))
	}
	if modelMap["protection_type"] != nil && modelMap["protection_type"].(string) != "" {
		model.ProtectionType = core.StringPtr(modelMap["protection_type"].(string))
	}
	if modelMap["sharepoint_site_summary"] != nil && len(modelMap["sharepoint_site_summary"].([]interface{})) > 0 {
		SharepointSiteSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToSharepointObjectParams(modelMap["sharepoint_site_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SharepointSiteSummary = SharepointSiteSummaryModel
	}
	if modelMap["os_type"] != nil && modelMap["os_type"].(string) != "" {
		model.OsType = core.StringPtr(modelMap["os_type"].(string))
	}
	if modelMap["child_objects"] != nil {
		childObjects := []backuprecoveryv1.ObjectSummary{}
		for _, childObjectsItem := range modelMap["child_objects"].([]interface{}) {
			childObjectsItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectSummary(childObjectsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			childObjects = append(childObjects, *childObjectsItemModel)
		}
		model.ChildObjects = childObjects
	}
	if modelMap["v_center_summary"] != nil && len(modelMap["v_center_summary"].([]interface{})) > 0 {
		VCenterSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeVCenterParams(modelMap["v_center_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VCenterSummary = VCenterSummaryModel
	}
	if modelMap["windows_cluster_summary"] != nil && len(modelMap["windows_cluster_summary"].([]interface{})) > 0 {
		WindowsClusterSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeWindowsClusterParams(modelMap["windows_cluster_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WindowsClusterSummary = WindowsClusterSummaryModel
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToSharepointObjectParams(modelMap map[string]interface{}) (*backuprecoveryv1.SharepointObjectParams, error) {
	model := &backuprecoveryv1.SharepointObjectParams{}
	if modelMap["site_web_url"] != nil && modelMap["site_web_url"].(string) != "" {
		model.SiteWebURL = core.StringPtr(modelMap["site_web_url"].(string))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToObjectSummary(modelMap map[string]interface{}) (*backuprecoveryv1.ObjectSummary, error) {
	model := &backuprecoveryv1.ObjectSummary{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["source_id"] != nil {
		model.SourceID = core.Int64Ptr(int64(modelMap["source_id"].(int)))
	}
	if modelMap["source_name"] != nil && modelMap["source_name"].(string) != "" {
		model.SourceName = core.StringPtr(modelMap["source_name"].(string))
	}
	if modelMap["environment"] != nil && modelMap["environment"].(string) != "" {
		model.Environment = core.StringPtr(modelMap["environment"].(string))
	}
	if modelMap["object_hash"] != nil && modelMap["object_hash"].(string) != "" {
		model.ObjectHash = core.StringPtr(modelMap["object_hash"].(string))
	}
	if modelMap["object_type"] != nil && modelMap["object_type"].(string) != "" {
		model.ObjectType = core.StringPtr(modelMap["object_type"].(string))
	}
	if modelMap["logical_size_bytes"] != nil {
		model.LogicalSizeBytes = core.Int64Ptr(int64(modelMap["logical_size_bytes"].(int)))
	}
	if modelMap["uuid"] != nil && modelMap["uuid"].(string) != "" {
		model.UUID = core.StringPtr(modelMap["uuid"].(string))
	}
	if modelMap["global_id"] != nil && modelMap["global_id"].(string) != "" {
		model.GlobalID = core.StringPtr(modelMap["global_id"].(string))
	}
	if modelMap["protection_type"] != nil && modelMap["protection_type"].(string) != "" {
		model.ProtectionType = core.StringPtr(modelMap["protection_type"].(string))
	}
	if modelMap["sharepoint_site_summary"] != nil && len(modelMap["sharepoint_site_summary"].([]interface{})) > 0 {
		SharepointSiteSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToSharepointObjectParams(modelMap["sharepoint_site_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SharepointSiteSummary = SharepointSiteSummaryModel
	}
	if modelMap["os_type"] != nil && modelMap["os_type"].(string) != "" {
		model.OsType = core.StringPtr(modelMap["os_type"].(string))
	}
	if modelMap["child_objects"] != nil {
		childObjects := []backuprecoveryv1.ObjectSummary{}
		for _, childObjectsItem := range modelMap["child_objects"].([]interface{}) {
			childObjectsItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectSummary(childObjectsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			childObjects = append(childObjects, *childObjectsItemModel)
		}
		model.ChildObjects = childObjects
	}
	if modelMap["v_center_summary"] != nil && len(modelMap["v_center_summary"].([]interface{})) > 0 {
		VCenterSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeVCenterParams(modelMap["v_center_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VCenterSummary = VCenterSummaryModel
	}
	if modelMap["windows_cluster_summary"] != nil && len(modelMap["windows_cluster_summary"].([]interface{})) > 0 {
		WindowsClusterSummaryModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeWindowsClusterParams(modelMap["windows_cluster_summary"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WindowsClusterSummary = WindowsClusterSummaryModel
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeVCenterParams(modelMap map[string]interface{}) (*backuprecoveryv1.ObjectTypeVCenterParams, error) {
	model := &backuprecoveryv1.ObjectTypeVCenterParams{}
	if modelMap["is_cloud_env"] != nil {
		model.IsCloudEnv = core.BoolPtr(modelMap["is_cloud_env"].(bool))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeWindowsClusterParams(modelMap map[string]interface{}) (*backuprecoveryv1.ObjectTypeWindowsClusterParams, error) {
	model := &backuprecoveryv1.ObjectTypeWindowsClusterParams{}
	if modelMap["cluster_source_type"] != nil && modelMap["cluster_source_type"].(string) != "" {
		model.ClusterSourceType = core.StringPtr(modelMap["cluster_source_type"].(string))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo, error) {
	model := &backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo{}
	if modelMap["target_id"] != nil {
		model.TargetID = core.Int64Ptr(int64(modelMap["target_id"].(int)))
	}
	if modelMap["archival_task_id"] != nil && modelMap["archival_task_id"].(string) != "" {
		model.ArchivalTaskID = core.StringPtr(modelMap["archival_task_id"].(string))
	}
	if modelMap["target_name"] != nil && modelMap["target_name"].(string) != "" {
		model.TargetName = core.StringPtr(modelMap["target_name"].(string))
	}
	if modelMap["target_type"] != nil && modelMap["target_type"].(string) != "" {
		model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	}
	if modelMap["usage_type"] != nil && modelMap["usage_type"].(string) != "" {
		model.UsageType = core.StringPtr(modelMap["usage_type"].(string))
	}
	if modelMap["ownership_context"] != nil && modelMap["ownership_context"].(string) != "" {
		model.OwnershipContext = core.StringPtr(modelMap["ownership_context"].(string))
	}
	if modelMap["tier_settings"] != nil && len(modelMap["tier_settings"].([]interface{})) > 0 {
		TierSettingsModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(modelMap["tier_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TierSettings = TierSettingsModel
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(modelMap map[string]interface{}) (*backuprecoveryv1.ArchivalTargetTierInfo, error) {
	model := &backuprecoveryv1.ArchivalTargetTierInfo{}
	if modelMap["aws_tiering"] != nil && len(modelMap["aws_tiering"].([]interface{})) > 0 {
		AwsTieringModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTiers(modelMap["aws_tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.AwsTiering = AwsTieringModel
	}
	if modelMap["azure_tiering"] != nil && len(modelMap["azure_tiering"].([]interface{})) > 0 {
		AzureTieringModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTiers(modelMap["azure_tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.AzureTiering = AzureTieringModel
	}
	model.CloudPlatform = core.StringPtr(modelMap["cloud_platform"].(string))
	if modelMap["google_tiering"] != nil && len(modelMap["google_tiering"].([]interface{})) > 0 {
		GoogleTieringModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTiers(modelMap["google_tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.GoogleTiering = GoogleTieringModel
	}
	if modelMap["oracle_tiering"] != nil && len(modelMap["oracle_tiering"].([]interface{})) > 0 {
		OracleTieringModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(modelMap["oracle_tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleTiering = OracleTieringModel
	}
	if modelMap["current_tier_type"] != nil && modelMap["current_tier_type"].(string) != "" {
		model.CurrentTierType = core.StringPtr(modelMap["current_tier_type"].(string))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTiers(modelMap map[string]interface{}) (*backuprecoveryv1.AWSTiers, error) {
	model := &backuprecoveryv1.AWSTiers{}
	tiers := []backuprecoveryv1.AWSTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTier(modelMap map[string]interface{}) (*backuprecoveryv1.AWSTier, error) {
	model := &backuprecoveryv1.AWSTier{}
	if modelMap["move_after_unit"] != nil && modelMap["move_after_unit"].(string) != "" {
		model.MoveAfterUnit = core.StringPtr(modelMap["move_after_unit"].(string))
	}
	if modelMap["move_after"] != nil {
		model.MoveAfter = core.Int64Ptr(int64(modelMap["move_after"].(int)))
	}
	model.TierType = core.StringPtr(modelMap["tier_type"].(string))
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTiers(modelMap map[string]interface{}) (*backuprecoveryv1.AzureTiers, error) {
	model := &backuprecoveryv1.AzureTiers{}
	if modelMap["tiers"] != nil {
		tiers := []backuprecoveryv1.AzureTier{}
		for _, tiersItem := range modelMap["tiers"].([]interface{}) {
			tiersItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTier(tiersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			tiers = append(tiers, *tiersItemModel)
		}
		model.Tiers = tiers
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTier(modelMap map[string]interface{}) (*backuprecoveryv1.AzureTier, error) {
	model := &backuprecoveryv1.AzureTier{}
	if modelMap["move_after_unit"] != nil && modelMap["move_after_unit"].(string) != "" {
		model.MoveAfterUnit = core.StringPtr(modelMap["move_after_unit"].(string))
	}
	if modelMap["move_after"] != nil {
		model.MoveAfter = core.Int64Ptr(int64(modelMap["move_after"].(int)))
	}
	model.TierType = core.StringPtr(modelMap["tier_type"].(string))
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTiers(modelMap map[string]interface{}) (*backuprecoveryv1.GoogleTiers, error) {
	model := &backuprecoveryv1.GoogleTiers{}
	tiers := []backuprecoveryv1.GoogleTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTier(modelMap map[string]interface{}) (*backuprecoveryv1.GoogleTier, error) {
	model := &backuprecoveryv1.GoogleTier{}
	if modelMap["move_after_unit"] != nil && modelMap["move_after_unit"].(string) != "" {
		model.MoveAfterUnit = core.StringPtr(modelMap["move_after_unit"].(string))
	}
	if modelMap["move_after"] != nil {
		model.MoveAfter = core.Int64Ptr(int64(modelMap["move_after"].(int)))
	}
	model.TierType = core.StringPtr(modelMap["tier_type"].(string))
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTiers, error) {
	model := &backuprecoveryv1.OracleTiers{}
	tiers := []backuprecoveryv1.OracleTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTier, error) {
	model := &backuprecoveryv1.OracleTier{}
	if modelMap["move_after_unit"] != nil && modelMap["move_after_unit"].(string) != "" {
		model.MoveAfterUnit = core.StringPtr(modelMap["move_after_unit"].(string))
	}
	if modelMap["move_after"] != nil {
		model.MoveAfter = core.Int64Ptr(int64(modelMap["move_after"].(int)))
	}
	model.TierType = core.StringPtr(modelMap["tier_type"].(string))
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(modelMap map[string]interface{}) (*backuprecoveryv1.FilesAndFoldersObject, error) {
	model := &backuprecoveryv1.FilesAndFoldersObject{}
	model.AbsolutePath = core.StringPtr(modelMap["absolute_path"].(string))
	if modelMap["is_directory"] != nil {
		model.IsDirectory = core.BoolPtr(modelMap["is_directory"].(bool))
	}
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersMapToDocumentObject(modelMap map[string]interface{}) (*backuprecoveryv1.DocumentObject, error) {
	model := &backuprecoveryv1.DocumentObject{}
	if modelMap["is_directory"] != nil {
		model.IsDirectory = core.BoolPtr(modelMap["is_directory"].(bool))
	}
	model.ItemID = core.StringPtr(modelMap["item_id"].(string))
	return model, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersDocumentObjectToMap(model *backuprecoveryv1.DocumentObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsDirectory != nil {
		modelMap["is_directory"] = *model.IsDirectory
	}
	modelMap["item_id"] = *model.ItemID
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["snapshot_id"] = *model.SnapshotID
	if model.PointInTimeUsecs != nil {
		modelMap["point_in_time_usecs"] = flex.IntValue(model.PointInTimeUsecs)
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = *model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = *model.ProtectionGroupName
	}
	if model.SnapshotCreationTimeUsecs != nil {
		modelMap["snapshot_creation_time_usecs"] = flex.IntValue(model.SnapshotCreationTimeUsecs)
	}
	if model.ObjectInfo != nil {
		objectInfoMap, err := ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["object_info"] = []map[string]interface{}{objectInfoMap}
	}
	if model.SnapshotTargetType != nil {
		modelMap["snapshot_target_type"] = *model.SnapshotTargetType
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.ArchivalTargetInfo != nil {
		archivalTargetInfoMap, err := ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_target_info"] = []map[string]interface{}{archivalTargetInfoMap}
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = *model.ProgressTaskID
	}
	if model.RecoverFromStandby != nil {
		modelMap["recover_from_standby"] = *model.RecoverFromStandby
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Messages != nil {
		modelMap["messages"] = model.Messages
	}
	if model.BytesRestored != nil {
		modelMap["bytes_restored"] = flex.IntValue(model.BytesRestored)
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = *model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = *model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = *model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = *model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = *model.ProtectionType
	}
	if model.SharepointSiteSummary != nil {
		sharepointSiteSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersSharepointObjectParamsToMap(model.SharepointSiteSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["sharepoint_site_summary"] = []map[string]interface{}{sharepointSiteSummaryMap}
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.ChildObjects != nil {
		childObjects := []map[string]interface{}{}
		for _, childObjectsItem := range model.ChildObjects {
			childObjectsItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectSummaryToMap(&childObjectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			childObjects = append(childObjects, childObjectsItemMap)
		}
		modelMap["child_objects"] = childObjects
	}
	if model.VCenterSummary != nil {
		vCenterSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectTypeVCenterParamsToMap(model.VCenterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["v_center_summary"] = []map[string]interface{}{vCenterSummaryMap}
	}
	if model.WindowsClusterSummary != nil {
		windowsClusterSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectTypeWindowsClusterParamsToMap(model.WindowsClusterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["windows_cluster_summary"] = []map[string]interface{}{windowsClusterSummaryMap}
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersSharepointObjectParamsToMap(model *backuprecoveryv1.SharepointObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SiteWebURL != nil {
		modelMap["site_web_url"] = *model.SiteWebURL
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersObjectSummaryToMap(model *backuprecoveryv1.ObjectSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.SourceName != nil {
		modelMap["source_name"] = *model.SourceName
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ObjectHash != nil {
		modelMap["object_hash"] = *model.ObjectHash
	}
	if model.ObjectType != nil {
		modelMap["object_type"] = *model.ObjectType
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	if model.GlobalID != nil {
		modelMap["global_id"] = *model.GlobalID
	}
	if model.ProtectionType != nil {
		modelMap["protection_type"] = *model.ProtectionType
	}
	if model.SharepointSiteSummary != nil {
		sharepointSiteSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersSharepointObjectParamsToMap(model.SharepointSiteSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["sharepoint_site_summary"] = []map[string]interface{}{sharepointSiteSummaryMap}
	}
	if model.OsType != nil {
		modelMap["os_type"] = *model.OsType
	}
	if model.ChildObjects != nil {
		childObjects := []map[string]interface{}{}
		for _, childObjectsItem := range model.ChildObjects {
			childObjectsItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectSummaryToMap(&childObjectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			childObjects = append(childObjects, childObjectsItemMap)
		}
		modelMap["child_objects"] = childObjects
	}
	if model.VCenterSummary != nil {
		vCenterSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectTypeVCenterParamsToMap(model.VCenterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["v_center_summary"] = []map[string]interface{}{vCenterSummaryMap}
	}
	if model.WindowsClusterSummary != nil {
		windowsClusterSummaryMap, err := ResourceIbmRecoveryDownloadFilesFoldersObjectTypeWindowsClusterParamsToMap(model.WindowsClusterSummary)
		if err != nil {
			return modelMap, err
		}
		modelMap["windows_cluster_summary"] = []map[string]interface{}{windowsClusterSummaryMap}
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersObjectTypeVCenterParamsToMap(model *backuprecoveryv1.ObjectTypeVCenterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsCloudEnv != nil {
		modelMap["is_cloud_env"] = *model.IsCloudEnv
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersObjectTypeWindowsClusterParamsToMap(model *backuprecoveryv1.ObjectTypeWindowsClusterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterSourceType != nil {
		modelMap["cluster_source_type"] = *model.ClusterSourceType
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetID != nil {
		modelMap["target_id"] = flex.IntValue(model.TargetID)
	}
	if model.ArchivalTaskID != nil {
		modelMap["archival_task_id"] = *model.ArchivalTaskID
	}
	if model.TargetName != nil {
		modelMap["target_name"] = *model.TargetName
	}
	if model.TargetType != nil {
		modelMap["target_type"] = *model.TargetType
	}
	if model.UsageType != nil {
		modelMap["usage_type"] = *model.UsageType
	}
	if model.OwnershipContext != nil {
		modelMap["ownership_context"] = *model.OwnershipContext
	}
	if model.TierSettings != nil {
		tierSettingsMap, err := ResourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AwsTiering != nil {
		awsTieringMap, err := ResourceIbmRecoveryDownloadFilesFoldersAWSTiersToMap(model.AwsTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["aws_tiering"] = []map[string]interface{}{awsTieringMap}
	}
	if model.AzureTiering != nil {
		azureTieringMap, err := ResourceIbmRecoveryDownloadFilesFoldersAzureTiersToMap(model.AzureTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["azure_tiering"] = []map[string]interface{}{azureTieringMap}
	}
	modelMap["cloud_platform"] = *model.CloudPlatform
	if model.GoogleTiering != nil {
		googleTieringMap, err := ResourceIbmRecoveryDownloadFilesFoldersGoogleTiersToMap(model.GoogleTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["google_tiering"] = []map[string]interface{}{googleTieringMap}
	}
	if model.OracleTiering != nil {
		oracleTieringMap, err := ResourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(model.OracleTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_tiering"] = []map[string]interface{}{oracleTieringMap}
	}
	if model.CurrentTierType != nil {
		modelMap["current_tier_type"] = *model.CurrentTierType
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersAWSTiersToMap(model *backuprecoveryv1.AWSTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersAWSTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersAWSTierToMap(model *backuprecoveryv1.AWSTier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoveAfterUnit != nil {
		modelMap["move_after_unit"] = *model.MoveAfterUnit
	}
	if model.MoveAfter != nil {
		modelMap["move_after"] = flex.IntValue(model.MoveAfter)
	}
	modelMap["tier_type"] = *model.TierType
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersAzureTiersToMap(model *backuprecoveryv1.AzureTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Tiers != nil {
		tiers := []map[string]interface{}{}
		for _, tiersItem := range model.Tiers {
			tiersItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersAzureTierToMap(&tiersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			tiers = append(tiers, tiersItemMap)
		}
		modelMap["tiers"] = tiers
	}
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersAzureTierToMap(model *backuprecoveryv1.AzureTier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoveAfterUnit != nil {
		modelMap["move_after_unit"] = *model.MoveAfterUnit
	}
	if model.MoveAfter != nil {
		modelMap["move_after"] = flex.IntValue(model.MoveAfter)
	}
	modelMap["tier_type"] = *model.TierType
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersGoogleTiersToMap(model *backuprecoveryv1.GoogleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersGoogleTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersGoogleTierToMap(model *backuprecoveryv1.GoogleTier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoveAfterUnit != nil {
		modelMap["move_after_unit"] = *model.MoveAfterUnit
	}
	if model.MoveAfter != nil {
		modelMap["move_after"] = flex.IntValue(model.MoveAfter)
	}
	modelMap["tier_type"] = *model.TierType
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := ResourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoveAfterUnit != nil {
		modelMap["move_after_unit"] = *model.MoveAfterUnit
	}
	if model.MoveAfter != nil {
		modelMap["move_after"] = flex.IntValue(model.MoveAfter)
	}
	modelMap["tier_type"] = *model.TierType
	return modelMap, nil
}

func ResourceIbmRecoveryDownloadFilesFoldersFilesAndFoldersObjectToMap(model *backuprecoveryv1.FilesAndFoldersObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["absolute_path"] = *model.AbsolutePath
	if model.IsDirectory != nil {
		modelMap["is_directory"] = *model.IsDirectory
	}
	return modelMap, nil
}
