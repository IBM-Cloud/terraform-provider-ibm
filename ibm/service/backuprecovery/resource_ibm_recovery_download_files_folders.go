// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name of the recovery task. This field must be set and must be a unique name.",
			},
			"object": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
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
									"os_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the operating system type of the object.",
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
												"cloud_platform": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the cloud platform to enable tiering.",
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
																			Computed:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
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
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the id of a Recovery.",
			},
			"request_initiator_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery", "request_initiator_type"),
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"snapshot_environment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery", "snapshot_environment"),
				Description: "Specifies the type of snapshot environment for which the Recovery was performed.",
			},
			"physical_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Specifies the recovery options specific to Physical environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Specifies the list of Recover Object parameters. For recovering files, specifies the object contains the file to recover.",
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
												"os_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the operating system type of the object.",
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
															"cloud_platform": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the cloud platform to enable tiering.",
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
																						Computed:    true,
																						Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																					},
																					"move_after": &schema.Schema{
																						Type:        schema.TypeInt,
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
						"recovery_action": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the type of recover action to be performed.",
						},
						"recover_volume_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to recover Physical Volumes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_environment": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
									},
									"physical_target_params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the params for recovering to a physical target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mount_target": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the target entity where the volumes are being mounted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the id of the object.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the object.",
															},
														},
													},
												},
												"volume_mapping": &schema.Schema{
													Type:        schema.TypeList,
													Required:    true,
													Description: "Specifies the mapping from source volumes to destination volumes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"source_volume_guid": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the guid of the source volume.",
															},
															"destination_volume_guid": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the guid of the destination volume.",
															},
														},
													},
												},
												"force_unmount_volume": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false.",
												},
												"vlan_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
															},
															"disable_vlan": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.",
															},
															"interface_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Interface group to use for Recovery.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"mount_volume_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to mount Physical Volumes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_environment": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
									},
									"physical_target_params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the params for recovering to a physical target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mount_to_original_target": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.",
												},
												"original_target_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the configuration for mounting to the original target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server_credentials": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies credentials to access the target server. This is required if the server is of Linux OS.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"username": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the username to access target entity.",
																		},
																		"password": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the password to access target entity.",
																		},
																	},
																},
															},
														},
													},
												},
												"new_target_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the configuration for mounting to a new target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mount_target": &schema.Schema{
																Type:        schema.TypeList,
																MinItems:    1,
																MaxItems:    1,
																Required:    true,
																Description: "Specifies the target entity to recover to.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the id of the object.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the object.",
																		},
																		"parent_source_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the id of the parent source of the target.",
																		},
																		"parent_source_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the parent source of the target.",
																		},
																	},
																},
															},
															"server_credentials": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies credentials to access the target server. This is required if the server is of Linux OS.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"username": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the username to access target entity.",
																		},
																		"password": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the password to access target entity.",
																		},
																	},
																},
															},
														},
													},
												},
												"read_only_mount": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to perform a read-only mount. Default is false.",
												},
												"volume_names": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"mounted_volume_mapping": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the mapping of original volumes and mounted volumes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"original_volume": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the original volume.",
															},
															"mounted_volume": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the point where the volume is mounted.",
															},
															"file_system_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the file system of the volume.",
															},
														},
													},
												},
												"vlan_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
															},
															"disable_vlan": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.",
															},
															"interface_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Interface group to use for Recovery.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"recover_file_and_folder_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to perform a file and folder recovery.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"files_and_folders": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "Specifies the information about the files and folders to be recovered.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"absolute_path": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the absolute path to the file or folder.",
												},
												"destination_dir": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the destination directory where the file/directory was copied.",
												},
												"is_directory": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether this is a directory or not.",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the recovery status for this file or folder.",
												},
												"messages": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specify error messages about the file during recovery.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"is_view_file_recovery": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specify if the recovery is of type view file/folder.",
												},
											},
										},
									},
									"target_environment": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
									},
									"physical_target_params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the parameters to recover to a Physical target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"recover_target": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the target entity where the volumes are being mounted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the id of the object.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the object.",
															},
															"parent_source_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the id of the parent source of the target.",
															},
															"parent_source_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the parent source of the target.",
															},
														},
													},
												},
												"restore_to_original_paths": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If this is true, then files will be restored to original paths.",
												},
												"overwrite_existing": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to overwrite existing file/folder during recovery.",
												},
												"alternate_restore_directory": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the directory path where restore should happen if restore_to_original_paths is set to false.",
												},
												"preserve_attributes": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to preserve file/folder attributes during recovery.",
												},
												"preserve_timestamps": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to preserve the original time stamps.",
												},
												"preserve_acls": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to preserve the ACLs of the original file.",
												},
												"continue_on_error": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.",
												},
												"save_success_files": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to save success files or not. Default value is false.",
												},
												"vlan_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
															},
															"disable_vlan": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.",
															},
															"interface_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Interface group to use for Recovery.",
															},
														},
													},
												},
												"restore_entity_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the restore type (restore everything or ACLs only) when restoring or downloading files or folders from a Physical file based or block based backup snapshot.",
												},
											},
										},
									},
								},
							},
						},
						"download_file_and_folder_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to download files and folders.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"files_and_folders": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the info about the files and folders to be recovered.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"absolute_path": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the absolute path to the file or folder.",
												},
												"destination_dir": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the destination directory where the file/directory was copied.",
												},
												"is_directory": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether this is a directory or not.",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the recovery status for this file or folder.",
												},
												"messages": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specify error messages about the file during recovery.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"is_view_file_recovery": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specify if the recovery is of type view file/folder.",
												},
											},
										},
									},
									"download_file_path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the path location to download the files and folders.",
									},
								},
							},
						},
						"system_recovery_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to perform a system recovery.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"full_nas_path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the path to the recovery view.",
									},
								},
							},
						},
					},
				},
			},
			"oracle_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Specifies the recovery options specific to oracle environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Specifies the list of parameters for list of objects to be recovered.",
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
												"os_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the operating system type of the object.",
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
															"cloud_platform": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the cloud platform to enable tiering.",
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
																						Computed:    true,
																						Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																					},
																					"move_after": &schema.Schema{
																						Type:        schema.TypeInt,
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
									"instant_recovery_info": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the info about instant recovery. This is only applicable for RecoverOracle.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"progress_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the progress monitor id.",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of the recovery.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time in Unix timestamp epoch in microseconds.",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time in Unix timestamp epoch in microseconds.",
												},
											},
										},
									},
								},
							},
						},
						"recovery_action": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the type of recover action to be performed.",
						},
						"recover_app_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters to recover Oracle databases.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_environment": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
									},
									"oracle_target_params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the params for recovering to a oracle host. Provided oracle backup should be recovered to same type of target host. For Example: If you have oracle backup taken from a physical host then that should be recovered to physical host only.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"recover_to_new_source": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Specifies the parameter whether the recovery should be performed to a new source or an original Source Target.",
												},
												"new_source_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"host": &schema.Schema{
																Type:        schema.TypeList,
																MinItems:    1,
																MaxItems:    1,
																Required:    true,
																Description: "Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the id of the object.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the object.",
																		},
																	},
																},
															},
															"recovery_target": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies if recovery target is a database or a view.",
															},
															"recover_database_params": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies recovery parameters when recovering to a database.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"restore_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
																		},
																		"db_channels": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"archive_log_retention_days": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																					},
																					"archive_log_retention_hours": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																					},
																					"credentials": &schema.Schema{
																						Type:        schema.TypeList,
																						MaxItems:    1,
																						Optional:    true,
																						Description: "Specifies the object to hold username and password.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"username": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies the username to access target entity.",
																								},
																								"password": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies the password to access target entity.",
																								},
																							},
																						},
																					},
																					"database_unique_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the unique Name of the database.",
																					},
																					"database_uuid": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																					},
																					"default_channel_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																					},
																					"database_node_list": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"host_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies the id of the database host from which backup is allowed.",
																								},
																								"channel_count": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																								},
																								"port": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies the port where the Database is listening.",
																								},
																								"sbt_host_params": &schema.Schema{
																									Type:        schema.TypeList,
																									MaxItems:    1,
																									Optional:    true,
																									Description: "Specifies details about capturing Oracle SBT host info.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"sbt_library_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the path of sbt library.",
																											},
																											"view_fs_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the Cohesity view path.",
																											},
																											"vip_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Optional:    true,
																												Description: "Specifies the list of Cohesity primary VIPs.",
																												Elem:        &schema.Schema{Type: schema.TypeString},
																											},
																											"vlan_info_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Optional:    true,
																												Description: "Specifies the Vlan information for Cohesity cluster.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"ip_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Optional:    true,
																															Description: "Specifies the list of Ips in this VLAN.",
																															Elem:        &schema.Schema{Type: schema.TypeString},
																														},
																														"gateway": &schema.Schema{
																															Type:        schema.TypeString,
																															Optional:    true,
																															Description: "Specifies the gateway of this VLAN.",
																														},
																														"id": &schema.Schema{
																															Type:        schema.TypeInt,
																															Optional:    true,
																															Description: "Specifies the Id of this VLAN.",
																														},
																														"subnet_ip": &schema.Schema{
																															Type:        schema.TypeString,
																															Optional:    true,
																															Description: "Specifies the subnet Ip for this VLAN.",
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																					"max_host_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																					},
																					"enable_dg_primary_backup": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																					},
																					"rman_backup_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the type of Oracle RMAN backup requested.",
																					},
																				},
																			},
																		},
																		"recovery_mode": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies if database should be left in recovery mode.",
																		},
																		"shell_evironment_vars": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "key.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "value.",
																					},
																				},
																			},
																		},
																		"granular_restore_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies information about list of objects (PDBs) to restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"granularity_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies type of granular restore.",
																					},
																					"pdb_restore_params": &schema.Schema{
																						Type:        schema.TypeList,
																						MaxItems:    1,
																						Optional:    true,
																						Description: "Specifies information about the list of pdbs to be restored.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"drop_duplicate_pdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																								},
																								"pdb_objects": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies list of PDB objects to restore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"db_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "Specifies pluggable database id.",
																											},
																											"db_name": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "Specifies name of the DB.",
																											},
																										},
																									},
																								},
																								"restore_to_existing_cdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies if pdbs should be restored to an existing CDB.",
																								},
																								"rename_pdb_map": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies the new PDB name mapping to existing PDBs.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"key": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "key.",
																											},
																											"value": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "value.",
																											},
																										},
																									},
																								},
																								"include_in_restore": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies whether to restore or skip the provided PDBs list.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"oracle_archive_log_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"range_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the type of range.",
																					},
																					"range_info_vec": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies an array of oracle restore ranges.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"start_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																								},
																								"end_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																								},
																								"protection_group_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																								},
																								"reset_log_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"incarnation_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"thread_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																								},
																							},
																						},
																					},
																					"archive_log_restore_dest": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies destination where archive logs are to be restored.",
																					},
																				},
																			},
																		},
																		"oracle_recovery_validation_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies parameters related to Oracle Recovery Validation.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"create_dummy_instance": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																					},
																				},
																			},
																		},
																		"restore_spfile_or_pfile_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies parameters related to spfile/pfile restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"should_restore_spfile_or_pfile": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether to restore spfile/pfile or skip it.",
																					},
																					"file_location": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																					},
																				},
																			},
																		},
																		"use_scn_for_restore": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether database recovery performed should use scn value or not.",
																		},
																		"database_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.",
																		},
																		"oracle_base_folder": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the oracle base folder at selected host.",
																		},
																		"oracle_home_folder": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the oracle home folder at selected host.",
																		},
																		"db_files_destination": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the location to restore database files.",
																		},
																		"db_config_file_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the config file path on selected host which configures the restored database.",
																		},
																		"enable_archive_log_mode": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies archive log mode for oracle restore.",
																		},
																		"pfile_parameter_map": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies a key value pair for pfile parameters.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "key.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "value.",
																					},
																				},
																			},
																		},
																		"bct_file_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies BCT file path.",
																		},
																		"num_tempfiles": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies no. of tempfiles to be used for the recovered database.",
																		},
																		"redo_log_config": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies redo log config.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"num_groups": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies no. of redo log groups.",
																					},
																					"member_prefix": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies Log member name prefix.",
																					},
																					"size_m_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies Size of the member in MB.",
																					},
																					"group_members": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies list of members of this redo log group.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																				},
																			},
																		},
																		"is_multi_stage_restore": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether this task is a multistage restore task. If set, we migrate the DB after clone completes.",
																		},
																		"oracle_update_restore_options": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies the parameters that are needed for updating oracle restore options.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"delay_secs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies when the migration of the oracle instance should be started after successful recovery.",
																					},
																					"target_path_vec": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies the target paths to be used for DB migration.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																				},
																			},
																		},
																		"skip_clone_nid": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Whether or not to skip the nid step in Oracle Clone workflow. Applicable to both smart and old clone workflow.",
																		},
																		"no_filename_check": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to validate filenames or not in Oracle alternate restore workflow.",
																		},
																		"new_name_clause": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies newname clause for db files which allows user to have full control on how their database files can be renamed during the oracle alternate restore workflow.",
																		},
																	},
																},
															},
															"recover_view_params": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies recovery parameters when recovering to a view.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"restore_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
																		},
																		"db_channels": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"archive_log_retention_days": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																					},
																					"archive_log_retention_hours": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																					},
																					"credentials": &schema.Schema{
																						Type:        schema.TypeList,
																						MaxItems:    1,
																						Optional:    true,
																						Description: "Specifies the object to hold username and password.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"username": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies the username to access target entity.",
																								},
																								"password": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies the password to access target entity.",
																								},
																							},
																						},
																					},
																					"database_unique_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the unique Name of the database.",
																					},
																					"database_uuid": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																					},
																					"default_channel_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																					},
																					"database_node_list": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"host_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies the id of the database host from which backup is allowed.",
																								},
																								"channel_count": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																								},
																								"port": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies the port where the Database is listening.",
																								},
																								"sbt_host_params": &schema.Schema{
																									Type:        schema.TypeList,
																									MaxItems:    1,
																									Optional:    true,
																									Description: "Specifies details about capturing Oracle SBT host info.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"sbt_library_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the path of sbt library.",
																											},
																											"view_fs_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the Cohesity view path.",
																											},
																											"vip_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Optional:    true,
																												Description: "Specifies the list of Cohesity primary VIPs.",
																												Elem:        &schema.Schema{Type: schema.TypeString},
																											},
																											"vlan_info_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Optional:    true,
																												Description: "Specifies the Vlan information for Cohesity cluster.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"ip_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Optional:    true,
																															Description: "Specifies the list of Ips in this VLAN.",
																															Elem:        &schema.Schema{Type: schema.TypeString},
																														},
																														"gateway": &schema.Schema{
																															Type:        schema.TypeString,
																															Optional:    true,
																															Description: "Specifies the gateway of this VLAN.",
																														},
																														"id": &schema.Schema{
																															Type:        schema.TypeInt,
																															Optional:    true,
																															Description: "Specifies the Id of this VLAN.",
																														},
																														"subnet_ip": &schema.Schema{
																															Type:        schema.TypeString,
																															Optional:    true,
																															Description: "Specifies the subnet Ip for this VLAN.",
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																					"max_host_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																					},
																					"enable_dg_primary_backup": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																					},
																					"rman_backup_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the type of Oracle RMAN backup requested.",
																					},
																				},
																			},
																		},
																		"recovery_mode": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies if database should be left in recovery mode.",
																		},
																		"shell_evironment_vars": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "key.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "value.",
																					},
																				},
																			},
																		},
																		"granular_restore_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies information about list of objects (PDBs) to restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"granularity_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies type of granular restore.",
																					},
																					"pdb_restore_params": &schema.Schema{
																						Type:        schema.TypeList,
																						MaxItems:    1,
																						Optional:    true,
																						Description: "Specifies information about the list of pdbs to be restored.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"drop_duplicate_pdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																								},
																								"pdb_objects": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies list of PDB objects to restore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"db_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "Specifies pluggable database id.",
																											},
																											"db_name": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "Specifies name of the DB.",
																											},
																										},
																									},
																								},
																								"restore_to_existing_cdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies if pdbs should be restored to an existing CDB.",
																								},
																								"rename_pdb_map": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies the new PDB name mapping to existing PDBs.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"key": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "key.",
																											},
																											"value": &schema.Schema{
																												Type:        schema.TypeString,
																												Required:    true,
																												Description: "value.",
																											},
																										},
																									},
																								},
																								"include_in_restore": &schema.Schema{
																									Type:        schema.TypeBool,
																									Optional:    true,
																									Description: "Specifies whether to restore or skip the provided PDBs list.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"oracle_archive_log_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"range_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the type of range.",
																					},
																					"range_info_vec": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies an array of oracle restore ranges.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"start_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																								},
																								"end_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																								},
																								"protection_group_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																								},
																								"reset_log_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"incarnation_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"thread_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																								},
																							},
																						},
																					},
																					"archive_log_restore_dest": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies destination where archive logs are to be restored.",
																					},
																				},
																			},
																		},
																		"oracle_recovery_validation_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies parameters related to Oracle Recovery Validation.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"create_dummy_instance": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																					},
																				},
																			},
																		},
																		"restore_spfile_or_pfile_info": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies parameters related to spfile/pfile restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"should_restore_spfile_or_pfile": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether to restore spfile/pfile or skip it.",
																					},
																					"file_location": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																					},
																				},
																			},
																		},
																		"use_scn_for_restore": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether database recovery performed should use scn value or not.",
																		},
																		"view_mount_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the directory where cohesity view for app recovery will be mounted.",
																		},
																	},
																},
															},
														},
													},
												},
												"original_source_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"restore_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
															},
															"db_channels": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"archive_log_retention_days": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																		},
																		"archive_log_retention_hours": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																		},
																		"credentials": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies the object to hold username and password.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"username": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Specifies the username to access target entity.",
																					},
																					"password": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Specifies the password to access target entity.",
																					},
																				},
																			},
																		},
																		"database_unique_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the unique Name of the database.",
																		},
																		"database_uuid": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																		},
																		"default_channel_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																		},
																		"database_node_list": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"host_id": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies the id of the database host from which backup is allowed.",
																					},
																					"channel_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																					},
																					"port": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the port where the Database is listening.",
																					},
																					"sbt_host_params": &schema.Schema{
																						Type:        schema.TypeList,
																						MaxItems:    1,
																						Optional:    true,
																						Description: "Specifies details about capturing Oracle SBT host info.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"sbt_library_path": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies the path of sbt library.",
																								},
																								"view_fs_path": &schema.Schema{
																									Type:        schema.TypeString,
																									Optional:    true,
																									Description: "Specifies the Cohesity view path.",
																								},
																								"vip_list": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies the list of Cohesity primary VIPs.",
																									Elem:        &schema.Schema{Type: schema.TypeString},
																								},
																								"vlan_info_list": &schema.Schema{
																									Type:        schema.TypeList,
																									Optional:    true,
																									Description: "Specifies the Vlan information for Cohesity cluster.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"ip_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Optional:    true,
																												Description: "Specifies the list of Ips in this VLAN.",
																												Elem:        &schema.Schema{Type: schema.TypeString},
																											},
																											"gateway": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the gateway of this VLAN.",
																											},
																											"id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Optional:    true,
																												Description: "Specifies the Id of this VLAN.",
																											},
																											"subnet_ip": &schema.Schema{
																												Type:        schema.TypeString,
																												Optional:    true,
																												Description: "Specifies the subnet Ip for this VLAN.",
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"max_host_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																		},
																		"enable_dg_primary_backup": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																		},
																		"rman_backup_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the type of Oracle RMAN backup requested.",
																		},
																	},
																},
															},
															"recovery_mode": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies if database should be left in recovery mode.",
															},
															"shell_evironment_vars": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "key.",
																		},
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "value.",
																		},
																	},
																},
															},
															"granular_restore_info": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies information about list of objects (PDBs) to restore.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"granularity_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies type of granular restore.",
																		},
																		"pdb_restore_params": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies information about the list of pdbs to be restored.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"drop_duplicate_pdb": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																					},
																					"pdb_objects": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies list of PDB objects to restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"db_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies pluggable database id.",
																								},
																								"db_name": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies name of the DB.",
																								},
																							},
																						},
																					},
																					"restore_to_existing_cdb": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies if pdbs should be restored to an existing CDB.",
																					},
																					"rename_pdb_map": &schema.Schema{
																						Type:        schema.TypeList,
																						Optional:    true,
																						Description: "Specifies the new PDB name mapping to existing PDBs.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"key": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "key.",
																								},
																								"value": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "value.",
																								},
																							},
																						},
																					},
																					"include_in_restore": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether to restore or skip the provided PDBs list.",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"oracle_archive_log_info": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"range_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the type of range.",
																		},
																		"range_info_vec": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies an array of oracle restore ranges.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"start_of_range": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																					},
																					"end_of_range": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																					},
																					"protection_group_id": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																					},
																					"reset_log_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																					},
																					"incarnation_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																					},
																					"thread_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																					},
																				},
																			},
																		},
																		"archive_log_restore_dest": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies destination where archive logs are to be restored.",
																		},
																	},
																},
															},
															"oracle_recovery_validation_info": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies parameters related to Oracle Recovery Validation.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"create_dummy_instance": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																		},
																	},
																},
															},
															"restore_spfile_or_pfile_info": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies parameters related to spfile/pfile restore.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"should_restore_spfile_or_pfile": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to restore spfile/pfile or skip it.",
																		},
																		"file_location": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																		},
																	},
																},
															},
															"use_scn_for_restore": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether database recovery performed should use scn value or not.",
															},
															"roll_forward_log_path_vec": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "List of archive logs to apply on Database after overwrite restore.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"attempt_complete_recovery": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether or not this is a complete recovery attempt.",
															},
															"roll_forward_time_msecs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "UTC time in msecs till which we have to roll-forward the database.",
															},
															"stop_active_passive": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether allowed to automatically stop active passive resource.",
															},
														},
													},
												},
											},
										},
									},
									"vlan_config": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
												},
												"disable_vlan": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.",
												},
												"interface_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Interface group to use for Recovery.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
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
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.",
			},
			"progress_task_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Progress monitor task id for Recovery.",
			},
			"recovery_action": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the type of recover action.",
			},
			"permissions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of tenants that have permissions for this recovery.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tenant id.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Tenant.",
						},
					},
				},
			},
			"creation_info": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the information about the creation of the protection group or recovery.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the user who created the protection group or recovery.",
						},
					},
				},
			},
			"can_tear_down": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether it's possible to tear down the objects created by the recovery.",
			},
			"tear_down_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed.",
			},
			"tear_down_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the error message about the tear down operation if it fails.",
			},
			"messages": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies messages about the recovery.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"is_parent_recovery": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery.",
			},
			"retrieve_archive_tasks": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of persistent state of a retrieve of an archive task.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_uid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the globally unique id for this retrieval of an archive task.",
						},
						"uptier_expiry_times": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies how much time the retrieved entity is present in the hot-tiers.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
			"is_multi_stage_restore": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case.",
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
		return diag.FromErr(err)
	}

	createDownloadFilesAndFoldersRecoveryOptions := &backuprecoveryv1.CreateDownloadFilesAndFoldersRecoveryOptions{}

	createDownloadFilesAndFoldersRecoveryOptions.SetName(d.Get("name").(string))
	objectModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(d.Get("object.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createDownloadFilesAndFoldersRecoveryOptions.SetObject(objectModel)
	var filesAndFolders []backuprecoveryv1.FilesAndFoldersObject
	for _, v := range d.Get("files_and_folders").([]interface{}) {
		value := v.(map[string]interface{})
		filesAndFoldersItem, err := resourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(value)
		if err != nil {
			return diag.FromErr(err)
		}
		filesAndFolders = append(filesAndFolders, *filesAndFoldersItem)
	}
	createDownloadFilesAndFoldersRecoveryOptions.SetFilesAndFolders(filesAndFolders)
	if _, ok := d.GetOk("parent_recovery_id"); ok {
		createDownloadFilesAndFoldersRecoveryOptions.SetParentRecoveryID(d.Get("parent_recovery_id").(string))
	}
	if _, ok := d.GetOk("glacier_retrieval_type"); ok {
		createDownloadFilesAndFoldersRecoveryOptions.SetGlacierRetrievalType(d.Get("glacier_retrieval_type").(string))
	}

	recovery, response, err := backupRecoveryClient.CreateDownloadFilesAndFoldersRecoveryWithContext(context, createDownloadFilesAndFoldersRecoveryOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateDownloadFilesAndFoldersRecoveryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateDownloadFilesAndFoldersRecoveryWithContext failed %s\n%s", err, response))
	}

	d.SetId(*recovery.ID)

	return resourceIbmRecoveryDownloadFilesFoldersRead(context, d, meta)
}

func resourceIbmRecoveryDownloadFilesFoldersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

	getRecoveryByIdOptions.SetID(d.Id())

	recovery, response, err := backupRecoveryClient.GetRecoveryByIDWithContext(context, getRecoveryByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetRecoveryByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRecoveryByIDWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", recovery.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("snapshot_environment", recovery.SnapshotEnvironment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting snapshot_environment: %s", err))
	}
	if !core.IsNil(recovery.PhysicalParams) {
		physicalParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsToMap(recovery.PhysicalParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("physical_params", []map[string]interface{}{physicalParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting physical_params: %s", err))
		}
	}
	if !core.IsNil(recovery.OracleParams) {
		oracleParamsMap, err := resourceIbmRecoveryRecoverOracleParamsToMap(recovery.OracleParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("oracle_params", []map[string]interface{}{oracleParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting oracle_params: %s", err))
		}
	}
	if !core.IsNil(recovery.StartTimeUsecs) {
		if err = d.Set("start_time_usecs", flex.IntValue(recovery.StartTimeUsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting start_time_usecs: %s", err))
		}
	}
	if !core.IsNil(recovery.EndTimeUsecs) {
		if err = d.Set("end_time_usecs", flex.IntValue(recovery.EndTimeUsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting end_time_usecs: %s", err))
		}
	}
	if !core.IsNil(recovery.Status) {
		if err = d.Set("status", recovery.Status); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
		}
	}
	if !core.IsNil(recovery.ProgressTaskID) {
		if err = d.Set("progress_task_id", recovery.ProgressTaskID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting progress_task_id: %s", err))
		}
	}
	if !core.IsNil(recovery.RecoveryAction) {
		if err = d.Set("recovery_action", recovery.RecoveryAction); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting recovery_action: %s", err))
		}
	}
	if !core.IsNil(recovery.Permissions) {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range recovery.Permissions {
			permissionsItemMap, err := resourceIbmRecoveryTenantToMap(&permissionsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			permissions = append(permissions, permissionsItemMap)
		}
		if err = d.Set("permissions", permissions); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting permissions: %s", err))
		}
	}
	if !core.IsNil(recovery.CreationInfo) {
		creationInfoMap, err := resourceIbmRecoveryCreationInfoToMap(recovery.CreationInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("creation_info", []map[string]interface{}{creationInfoMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting creation_info: %s", err))
		}
	}
	if !core.IsNil(recovery.CanTearDown) {
		if err = d.Set("can_tear_down", recovery.CanTearDown); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting can_tear_down: %s", err))
		}
	}
	if !core.IsNil(recovery.TearDownStatus) {
		if err = d.Set("tear_down_status", recovery.TearDownStatus); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tear_down_status: %s", err))
		}
	}
	if !core.IsNil(recovery.TearDownMessage) {
		if err = d.Set("tear_down_message", recovery.TearDownMessage); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tear_down_message: %s", err))
		}
	}
	if !core.IsNil(recovery.Messages) {
		if err = d.Set("messages", recovery.Messages); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting messages: %s", err))
		}
	}
	if !core.IsNil(recovery.IsParentRecovery) {
		if err = d.Set("is_parent_recovery", recovery.IsParentRecovery); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_parent_recovery: %s", err))
		}
	}
	if !core.IsNil(recovery.ParentRecoveryID) {
		if err = d.Set("parent_recovery_id", recovery.ParentRecoveryID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting parent_recovery_id: %s", err))
		}
	}
	if !core.IsNil(recovery.RetrieveArchiveTasks) {
		retrieveArchiveTasks := []map[string]interface{}{}
		for _, retrieveArchiveTasksItem := range recovery.RetrieveArchiveTasks {
			retrieveArchiveTasksItemMap, err := resourceIbmRecoveryRetrieveArchiveTaskToMap(&retrieveArchiveTasksItem)
			if err != nil {
				return diag.FromErr(err)
			}
			retrieveArchiveTasks = append(retrieveArchiveTasks, retrieveArchiveTasksItemMap)
		}
		if err = d.Set("retrieve_archive_tasks", retrieveArchiveTasks); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting retrieve_archive_tasks: %s", err))
		}
	}
	if !core.IsNil(recovery.IsMultiStageRestore) {
		if err = d.Set("is_multi_stage_restore", recovery.IsMultiStageRestore); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_multi_stage_restore: %s", err))
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

func resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParams, error) {
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
		ObjectInfoModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap["object_info"].([]interface{})[0].(map[string]interface{}))
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
		ArchivalTargetInfoModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap["archival_target_info"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo, error) {
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
	if modelMap["os_type"] != nil && modelMap["os_type"].(string) != "" {
		model.OsType = core.StringPtr(modelMap["os_type"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo, error) {
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
		TierSettingsModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(modelMap["tier_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TierSettings = TierSettingsModel
	}
	return model, nil
}

func resourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(modelMap map[string]interface{}) (*backuprecoveryv1.ArchivalTargetTierInfo, error) {
	model := &backuprecoveryv1.ArchivalTargetTierInfo{}
	model.CloudPlatform = core.StringPtr(modelMap["cloud_platform"].(string))
	if modelMap["oracle_tiering"] != nil && len(modelMap["oracle_tiering"].([]interface{})) > 0 {
		OracleTieringModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(modelMap["oracle_tiering"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTiers, error) {
	model := &backuprecoveryv1.OracleTiers{}
	tiers := []backuprecoveryv1.OracleTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := resourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func resourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTier, error) {
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

func resourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(modelMap map[string]interface{}) (*backuprecoveryv1.FilesAndFoldersObject, error) {
	model := &backuprecoveryv1.FilesAndFoldersObject{}
	model.AbsolutePath = core.StringPtr(modelMap["absolute_path"].(string))
	if modelMap["is_directory"] != nil {
		model.IsDirectory = core.BoolPtr(modelMap["is_directory"].(bool))
	}
	return model, nil
}

func resourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["snapshot_id"] = model.SnapshotID
	if model.PointInTimeUsecs != nil {
		modelMap["point_in_time_usecs"] = flex.IntValue(model.PointInTimeUsecs)
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.SnapshotCreationTimeUsecs != nil {
		modelMap["snapshot_creation_time_usecs"] = flex.IntValue(model.SnapshotCreationTimeUsecs)
	}
	if model.ObjectInfo != nil {
		objectInfoMap, err := resourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["object_info"] = []map[string]interface{}{objectInfoMap}
	}
	if model.SnapshotTargetType != nil {
		modelMap["snapshot_target_type"] = model.SnapshotTargetType
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.ArchivalTargetInfo != nil {
		archivalTargetInfoMap, err := resourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_target_info"] = []map[string]interface{}{archivalTargetInfoMap}
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.RecoverFromStandby != nil {
		modelMap["recover_from_standby"] = model.RecoverFromStandby
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
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

func resourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo) (map[string]interface{}, error) {
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

func resourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetID != nil {
		modelMap["target_id"] = flex.IntValue(model.TargetID)
	}
	if model.ArchivalTaskID != nil {
		modelMap["archival_task_id"] = model.ArchivalTaskID
	}
	if model.TargetName != nil {
		modelMap["target_name"] = model.TargetName
	}
	if model.TargetType != nil {
		modelMap["target_type"] = model.TargetType
	}
	if model.UsageType != nil {
		modelMap["usage_type"] = model.UsageType
	}
	if model.OwnershipContext != nil {
		modelMap["ownership_context"] = model.OwnershipContext
	}
	if model.TierSettings != nil {
		tierSettingsMap, err := resourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := resourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(model.OracleTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_tiering"] = []map[string]interface{}{oracleTieringMap}
	}
	if model.CurrentTierType != nil {
		modelMap["current_tier_type"] = model.CurrentTierType
	}
	return modelMap, nil
}

func resourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := resourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func resourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MoveAfterUnit != nil {
		modelMap["move_after_unit"] = model.MoveAfterUnit
	}
	if model.MoveAfter != nil {
		modelMap["move_after"] = flex.IntValue(model.MoveAfter)
	}
	modelMap["tier_type"] = model.TierType
	return modelMap, nil
}

func resourceIbmRecoveryDownloadFilesFoldersFilesAndFoldersObjectToMap(model *backuprecoveryv1.FilesAndFoldersObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["absolute_path"] = model.AbsolutePath
	if model.IsDirectory != nil {
		modelMap["is_directory"] = model.IsDirectory
	}
	return modelMap, nil
}
