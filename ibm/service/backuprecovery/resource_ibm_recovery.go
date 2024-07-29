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

func ResourceIbmRecovery() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmRecoveryCreate,
		ReadContext:   resourceIbmRecoveryRead,
		DeleteContext: resourceIbmRecoveryDelete,
		UpdateContext: resourceIbmRecoveryUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the id of a Recovery.",
			},
			"request_initiator_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery", "request_initiator_type"),
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name of the Recovery.",
			},
			"snapshot_environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_recovery", "snapshot_environment"),
				Description: "Specifies the type of snapshot environment for which the Recovery was performed.",
			},
			"physical_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
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
				Optional:    true,
				ForceNew:    true,
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
			"parent_recovery_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery.",
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

func ResourceIbmRecoveryValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "request_initiator_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "Helios, UIAuto, UIUser",
		},
		validate.ValidateSchema{
			Identifier:                 "snapshot_environment",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kOracle, kPhysical, kSQL",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_recovery", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmRecoveryCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createRecoveryOptions := &backuprecoveryv1.CreateRecoveryOptions{}

	createRecoveryOptions.SetName(d.Get("name").(string))
	createRecoveryOptions.SetSnapshotEnvironment(d.Get("snapshot_environment").(string))
	if _, ok := d.GetOk("physical_params"); ok {
		physicalParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createRecoveryOptions.SetPhysicalParams(physicalParamsModel)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		oracleParamsModel, err := resourceIbmRecoveryMapToRecoverOracleParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createRecoveryOptions.SetOracleParams(oracleParamsModel)
	}
	if _, ok := d.GetOk("request_initiator_type"); ok {
		createRecoveryOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}

	recovery, response, err := backupRecoveryClient.CreateRecoveryWithContext(context, createRecoveryOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateRecoveryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateRecoveryWithContext failed %s\n%s", err, response))
	}

	d.SetId(*recovery.ID)

	return resourceIbmRecoveryRead(context, d, meta)
}

func resourceIbmRecoveryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryMapToRecoverPhysicalParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParams{}
	objects := []backuprecoveryv1.CommonRecoverObjectSnapshotParams{}
	for _, objectsItem := range modelMap["objects"].([]interface{}) {
		objectsItemModel, err := resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParams(objectsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		objects = append(objects, *objectsItemModel)
	}
	model.Objects = objects
	model.RecoveryAction = core.StringPtr(modelMap["recovery_action"].(string))
	if modelMap["recover_volume_params"] != nil && len(modelMap["recover_volume_params"].([]interface{})) > 0 {
		RecoverVolumeParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParamsRecoverVolumeParams(modelMap["recover_volume_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RecoverVolumeParams = RecoverVolumeParamsModel
	}
	if modelMap["mount_volume_params"] != nil && len(modelMap["mount_volume_params"].([]interface{})) > 0 {
		MountVolumeParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParamsMountVolumeParams(modelMap["mount_volume_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MountVolumeParams = MountVolumeParamsModel
	}
	if modelMap["recover_file_and_folder_params"] != nil && len(modelMap["recover_file_and_folder_params"].([]interface{})) > 0 {
		RecoverFileAndFolderParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParamsRecoverFileAndFolderParams(modelMap["recover_file_and_folder_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RecoverFileAndFolderParams = RecoverFileAndFolderParamsModel
	}
	if modelMap["download_file_and_folder_params"] != nil && len(modelMap["download_file_and_folder_params"].([]interface{})) > 0 {
		DownloadFileAndFolderParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParamsDownloadFileAndFolderParams(modelMap["download_file_and_folder_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DownloadFileAndFolderParams = DownloadFileAndFolderParamsModel
	}
	if modelMap["system_recovery_params"] != nil && len(modelMap["system_recovery_params"].([]interface{})) > 0 {
		SystemRecoveryParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalParamsSystemRecoveryParams(modelMap["system_recovery_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemRecoveryParams = SystemRecoveryParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParams(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParams, error) {
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
		ObjectInfoModel, err := resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap["object_info"].([]interface{})[0].(map[string]interface{}))
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
		ArchivalTargetInfoModel, err := resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap["archival_target_info"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo, error) {
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

func resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo, error) {
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
		TierSettingsModel, err := resourceIbmRecoveryMapToArchivalTargetTierInfo(modelMap["tier_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TierSettings = TierSettingsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToArchivalTargetTierInfo(modelMap map[string]interface{}) (*backuprecoveryv1.ArchivalTargetTierInfo, error) {
	model := &backuprecoveryv1.ArchivalTargetTierInfo{}
	model.CloudPlatform = core.StringPtr(modelMap["cloud_platform"].(string))
	if modelMap["oracle_tiering"] != nil && len(modelMap["oracle_tiering"].([]interface{})) > 0 {
		OracleTieringModel, err := resourceIbmRecoveryMapToOracleTiers(modelMap["oracle_tiering"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmRecoveryMapToOracleTiers(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTiers, error) {
	model := &backuprecoveryv1.OracleTiers{}
	tiers := []backuprecoveryv1.OracleTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := resourceIbmRecoveryMapToOracleTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func resourceIbmRecoveryMapToOracleTier(modelMap map[string]interface{}) (*backuprecoveryv1.OracleTier, error) {
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

func resourceIbmRecoveryMapToRecoverPhysicalParamsRecoverVolumeParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams{}
	model.TargetEnvironment = core.StringPtr(modelMap["target_environment"].(string))
	if modelMap["physical_target_params"] != nil && len(modelMap["physical_target_params"].([]interface{})) > 0 {
		PhysicalTargetParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalVolumeParamsPhysicalTargetParams(modelMap["physical_target_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PhysicalTargetParams = PhysicalTargetParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalVolumeParamsPhysicalTargetParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams{}
	MountTargetModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverVolumeMountTarget(modelMap["mount_target"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.MountTarget = MountTargetModel
	volumeMapping := []backuprecoveryv1.RecoverVolumeMapping{}
	for _, volumeMappingItem := range modelMap["volume_mapping"].([]interface{}) {
		volumeMappingItemModel, err := resourceIbmRecoveryMapToRecoverVolumeMapping(volumeMappingItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		volumeMapping = append(volumeMapping, *volumeMappingItemModel)
	}
	model.VolumeMapping = volumeMapping
	if modelMap["force_unmount_volume"] != nil {
		model.ForceUnmountVolume = core.BoolPtr(modelMap["force_unmount_volume"].(bool))
	}
	if modelMap["vlan_config"] != nil && len(modelMap["vlan_config"].([]interface{})) > 0 {
		VlanConfigModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverVolumeVlanConfig(modelMap["vlan_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanConfig = VlanConfigModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverVolumeMountTarget(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverVolumeMapping(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverVolumeMapping, error) {
	model := &backuprecoveryv1.RecoverVolumeMapping{}
	model.SourceVolumeGuid = core.StringPtr(modelMap["source_volume_guid"].(string))
	model.DestinationVolumeGuid = core.StringPtr(modelMap["destination_volume_guid"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverVolumeVlanConfig(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalParamsMountVolumeParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams{}
	model.TargetEnvironment = core.StringPtr(modelMap["target_environment"].(string))
	if modelMap["physical_target_params"] != nil && len(modelMap["physical_target_params"].([]interface{})) > 0 {
		PhysicalTargetParamsModel, err := resourceIbmRecoveryMapToMountPhysicalVolumeParamsPhysicalTargetParams(modelMap["physical_target_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PhysicalTargetParams = PhysicalTargetParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToMountPhysicalVolumeParamsPhysicalTargetParams(modelMap map[string]interface{}) (*backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams, error) {
	model := &backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams{}
	model.MountToOriginalTarget = core.BoolPtr(modelMap["mount_to_original_target"].(bool))
	if modelMap["original_target_config"] != nil && len(modelMap["original_target_config"].([]interface{})) > 0 {
		OriginalTargetConfigModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeOriginalTargetConfig(modelMap["original_target_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OriginalTargetConfig = OriginalTargetConfigModel
	}
	if modelMap["new_target_config"] != nil && len(modelMap["new_target_config"].([]interface{})) > 0 {
		NewTargetConfigModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeNewTargetConfig(modelMap["new_target_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NewTargetConfig = NewTargetConfigModel
	}
	if modelMap["read_only_mount"] != nil {
		model.ReadOnlyMount = core.BoolPtr(modelMap["read_only_mount"].(bool))
	}
	if modelMap["volume_names"] != nil {
		volumeNames := []string{}
		for _, volumeNamesItem := range modelMap["volume_names"].([]interface{}) {
			volumeNames = append(volumeNames, volumeNamesItem.(string))
		}
		model.VolumeNames = volumeNames
	}
	if modelMap["mounted_volume_mapping"] != nil {
		mountedVolumeMapping := []backuprecoveryv1.MountedVolumeMapping{}
		for _, mountedVolumeMappingItem := range modelMap["mounted_volume_mapping"].([]interface{}) {
			mountedVolumeMappingItemModel, err := resourceIbmRecoveryMapToMountedVolumeMapping(mountedVolumeMappingItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			mountedVolumeMapping = append(mountedVolumeMapping, *mountedVolumeMappingItemModel)
		}
		model.MountedVolumeMapping = mountedVolumeMapping
	}
	if modelMap["vlan_config"] != nil && len(modelMap["vlan_config"].([]interface{})) > 0 {
		VlanConfigModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeVlanConfig(modelMap["vlan_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanConfig = VlanConfigModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeOriginalTargetConfig(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig{}
	if modelMap["server_credentials"] != nil && len(modelMap["server_credentials"].([]interface{})) > 0 {
		ServerCredentialsModel, err := resourceIbmRecoveryMapToPhysicalMountVolumesOriginalTargetConfigServerCredentials(modelMap["server_credentials"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ServerCredentials = ServerCredentialsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalMountVolumesOriginalTargetConfigServerCredentials(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials, error) {
	model := &backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials{}
	model.Username = core.StringPtr(modelMap["username"].(string))
	model.Password = core.StringPtr(modelMap["password"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeNewTargetConfig(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig{}
	MountTargetModel, err := resourceIbmRecoveryMapToRecoverTarget(modelMap["mount_target"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.MountTarget = MountTargetModel
	if modelMap["server_credentials"] != nil && len(modelMap["server_credentials"].([]interface{})) > 0 {
		ServerCredentialsModel, err := resourceIbmRecoveryMapToPhysicalMountVolumesNewTargetConfigServerCredentials(modelMap["server_credentials"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ServerCredentials = ServerCredentialsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverTarget(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverTarget, error) {
	model := &backuprecoveryv1.RecoverTarget{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["parent_source_id"] != nil {
		model.ParentSourceID = core.Int64Ptr(int64(modelMap["parent_source_id"].(int)))
	}
	if modelMap["parent_source_name"] != nil && modelMap["parent_source_name"].(string) != "" {
		model.ParentSourceName = core.StringPtr(modelMap["parent_source_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalMountVolumesNewTargetConfigServerCredentials(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials, error) {
	model := &backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials{}
	model.Username = core.StringPtr(modelMap["username"].(string))
	model.Password = core.StringPtr(modelMap["password"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToMountedVolumeMapping(modelMap map[string]interface{}) (*backuprecoveryv1.MountedVolumeMapping, error) {
	model := &backuprecoveryv1.MountedVolumeMapping{}
	if modelMap["original_volume"] != nil && modelMap["original_volume"].(string) != "" {
		model.OriginalVolume = core.StringPtr(modelMap["original_volume"].(string))
	}
	if modelMap["mounted_volume"] != nil && modelMap["mounted_volume"].(string) != "" {
		model.MountedVolume = core.StringPtr(modelMap["mounted_volume"].(string))
	}
	if modelMap["file_system_type"] != nil && modelMap["file_system_type"].(string) != "" {
		model.FileSystemType = core.StringPtr(modelMap["file_system_type"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForMountVolumeVlanConfig(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalParamsRecoverFileAndFolderParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams{}
	filesAndFolders := []backuprecoveryv1.CommonRecoverFileAndFolderInfo{}
	for _, filesAndFoldersItem := range modelMap["files_and_folders"].([]interface{}) {
		filesAndFoldersItemModel, err := resourceIbmRecoveryMapToCommonRecoverFileAndFolderInfo(filesAndFoldersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		filesAndFolders = append(filesAndFolders, *filesAndFoldersItemModel)
	}
	model.FilesAndFolders = filesAndFolders
	model.TargetEnvironment = core.StringPtr(modelMap["target_environment"].(string))
	if modelMap["physical_target_params"] != nil && len(modelMap["physical_target_params"].([]interface{})) > 0 {
		PhysicalTargetParamsModel, err := resourceIbmRecoveryMapToRecoverPhysicalFileAndFolderParamsPhysicalTargetParams(modelMap["physical_target_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PhysicalTargetParams = PhysicalTargetParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonRecoverFileAndFolderInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverFileAndFolderInfo, error) {
	model := &backuprecoveryv1.CommonRecoverFileAndFolderInfo{}
	model.AbsolutePath = core.StringPtr(modelMap["absolute_path"].(string))
	if modelMap["destination_dir"] != nil && modelMap["destination_dir"].(string) != "" {
		model.DestinationDir = core.StringPtr(modelMap["destination_dir"].(string))
	}
	if modelMap["is_directory"] != nil {
		model.IsDirectory = core.BoolPtr(modelMap["is_directory"].(bool))
	}
	if modelMap["status"] != nil && modelMap["status"].(string) != "" {
		model.Status = core.StringPtr(modelMap["status"].(string))
	}
	if modelMap["messages"] != nil {
		messages := []string{}
		for _, messagesItem := range modelMap["messages"].([]interface{}) {
			messages = append(messages, messagesItem.(string))
		}
		model.Messages = messages
	}
	if modelMap["is_view_file_recovery"] != nil {
		model.IsViewFileRecovery = core.BoolPtr(modelMap["is_view_file_recovery"].(bool))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalFileAndFolderParamsPhysicalTargetParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams{}
	RecoverTargetModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverFileAndFolderRecoverTarget(modelMap["recover_target"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.RecoverTarget = RecoverTargetModel
	if modelMap["restore_to_original_paths"] != nil {
		model.RestoreToOriginalPaths = core.BoolPtr(modelMap["restore_to_original_paths"].(bool))
	}
	if modelMap["overwrite_existing"] != nil {
		model.OverwriteExisting = core.BoolPtr(modelMap["overwrite_existing"].(bool))
	}
	if modelMap["alternate_restore_directory"] != nil && modelMap["alternate_restore_directory"].(string) != "" {
		model.AlternateRestoreDirectory = core.StringPtr(modelMap["alternate_restore_directory"].(string))
	}
	if modelMap["preserve_attributes"] != nil {
		model.PreserveAttributes = core.BoolPtr(modelMap["preserve_attributes"].(bool))
	}
	if modelMap["preserve_timestamps"] != nil {
		model.PreserveTimestamps = core.BoolPtr(modelMap["preserve_timestamps"].(bool))
	}
	if modelMap["preserve_acls"] != nil {
		model.PreserveAcls = core.BoolPtr(modelMap["preserve_acls"].(bool))
	}
	if modelMap["continue_on_error"] != nil {
		model.ContinueOnError = core.BoolPtr(modelMap["continue_on_error"].(bool))
	}
	if modelMap["save_success_files"] != nil {
		model.SaveSuccessFiles = core.BoolPtr(modelMap["save_success_files"].(bool))
	}
	if modelMap["vlan_config"] != nil && len(modelMap["vlan_config"].([]interface{})) > 0 {
		VlanConfigModel, err := resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverFileAndFolderVlanConfig(modelMap["vlan_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanConfig = VlanConfigModel
	}
	if modelMap["restore_entity_type"] != nil && modelMap["restore_entity_type"].(string) != "" {
		model.RestoreEntityType = core.StringPtr(modelMap["restore_entity_type"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverFileAndFolderRecoverTarget(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["parent_source_id"] != nil {
		model.ParentSourceID = core.Int64Ptr(int64(modelMap["parent_source_id"].(int)))
	}
	if modelMap["parent_source_name"] != nil && modelMap["parent_source_name"].(string) != "" {
		model.ParentSourceName = core.StringPtr(modelMap["parent_source_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToPhysicalTargetParamsForRecoverFileAndFolderVlanConfig(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig, error) {
	model := &backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalParamsDownloadFileAndFolderParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams{}
	if modelMap["files_and_folders"] != nil {
		filesAndFolders := []backuprecoveryv1.CommonRecoverFileAndFolderInfo{}
		for _, filesAndFoldersItem := range modelMap["files_and_folders"].([]interface{}) {
			filesAndFoldersItemModel, err := resourceIbmRecoveryMapToCommonRecoverFileAndFolderInfo(filesAndFoldersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			filesAndFolders = append(filesAndFolders, *filesAndFoldersItemModel)
		}
		model.FilesAndFolders = filesAndFolders
	}
	if modelMap["download_file_path"] != nil && modelMap["download_file_path"].(string) != "" {
		model.DownloadFilePath = core.StringPtr(modelMap["download_file_path"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverPhysicalParamsSystemRecoveryParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams, error) {
	model := &backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams{}
	if modelMap["full_nas_path"] != nil && modelMap["full_nas_path"].(string) != "" {
		model.FullNasPath = core.StringPtr(modelMap["full_nas_path"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleParams, error) {
	model := &backuprecoveryv1.RecoverOracleParams{}
	objects := []backuprecoveryv1.RecoverOracleDbSnapshotParams{}
	for _, objectsItem := range modelMap["objects"].([]interface{}) {
		objectsItemModel, err := resourceIbmRecoveryMapToRecoverOracleDbSnapshotParams(objectsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		objects = append(objects, *objectsItemModel)
	}
	model.Objects = objects
	model.RecoveryAction = core.StringPtr(modelMap["recovery_action"].(string))
	if modelMap["recover_app_params"] != nil && len(modelMap["recover_app_params"].([]interface{})) > 0 {
		RecoverAppParamsModel, err := resourceIbmRecoveryMapToRecoverOracleParamsRecoverAppParams(modelMap["recover_app_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RecoverAppParams = RecoverAppParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleDbSnapshotParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleDbSnapshotParams, error) {
	model := &backuprecoveryv1.RecoverOracleDbSnapshotParams{}
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
		ObjectInfoModel, err := resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsObjectInfo(modelMap["object_info"].([]interface{})[0].(map[string]interface{}))
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
		ArchivalTargetInfoModel, err := resourceIbmRecoveryMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(modelMap["archival_target_info"].([]interface{})[0].(map[string]interface{}))
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
	if modelMap["instant_recovery_info"] != nil && len(modelMap["instant_recovery_info"].([]interface{})) > 0 {
		InstantRecoveryInfoModel, err := resourceIbmRecoveryMapToRecoverOracleDbSnapshotParamsInstantRecoveryInfo(modelMap["instant_recovery_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.InstantRecoveryInfo = InstantRecoveryInfoModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleDbSnapshotParamsInstantRecoveryInfo(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleDbSnapshotParamsInstantRecoveryInfo, error) {
	model := &backuprecoveryv1.RecoverOracleDbSnapshotParamsInstantRecoveryInfo{}
	if modelMap["progress_task_id"] != nil && modelMap["progress_task_id"].(string) != "" {
		model.ProgressTaskID = core.StringPtr(modelMap["progress_task_id"].(string))
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
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleParamsRecoverAppParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleParamsRecoverAppParams, error) {
	model := &backuprecoveryv1.RecoverOracleParamsRecoverAppParams{}
	model.TargetEnvironment = core.StringPtr(modelMap["target_environment"].(string))
	if modelMap["oracle_target_params"] != nil && len(modelMap["oracle_target_params"].([]interface{})) > 0 {
		OracleTargetParamsModel, err := resourceIbmRecoveryMapToRecoverOracleAppParamsOracleTargetParams(modelMap["oracle_target_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleTargetParams = OracleTargetParamsModel
	}
	if modelMap["vlan_config"] != nil && len(modelMap["vlan_config"].([]interface{})) > 0 {
		VlanConfigModel, err := resourceIbmRecoveryMapToRecoverOracleAppParamsVlanConfig(modelMap["vlan_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanConfig = VlanConfigModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleAppParamsOracleTargetParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleAppParamsOracleTargetParams, error) {
	model := &backuprecoveryv1.RecoverOracleAppParamsOracleTargetParams{}
	model.RecoverToNewSource = core.BoolPtr(modelMap["recover_to_new_source"].(bool))
	if modelMap["new_source_config"] != nil && len(modelMap["new_source_config"].([]interface{})) > 0 {
		NewSourceConfigModel, err := resourceIbmRecoveryMapToCommonRecoverOracleAppTargetParamsNewSourceConfig(modelMap["new_source_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NewSourceConfig = NewSourceConfigModel
	}
	if modelMap["original_source_config"] != nil && len(modelMap["original_source_config"].([]interface{})) > 0 {
		OriginalSourceConfigModel, err := resourceIbmRecoveryMapToCommonRecoverOracleAppTargetParamsOriginalSourceConfig(modelMap["original_source_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OriginalSourceConfig = OriginalSourceConfigModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonRecoverOracleAppTargetParamsNewSourceConfig(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverOracleAppTargetParamsNewSourceConfig, error) {
	model := &backuprecoveryv1.CommonRecoverOracleAppTargetParamsNewSourceConfig{}
	HostModel, err := resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigHost(modelMap["host"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Host = HostModel
	if modelMap["recovery_target"] != nil && modelMap["recovery_target"].(string) != "" {
		model.RecoveryTarget = core.StringPtr(modelMap["recovery_target"].(string))
	}
	if modelMap["recover_database_params"] != nil && len(modelMap["recover_database_params"].([]interface{})) > 0 {
		RecoverDatabaseParamsModel, err := resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigRecoverDatabaseParams(modelMap["recover_database_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RecoverDatabaseParams = RecoverDatabaseParamsModel
	}
	if modelMap["recover_view_params"] != nil && len(modelMap["recover_view_params"].([]interface{})) > 0 {
		RecoverViewParamsModel, err := resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigRecoverViewParams(modelMap["recover_view_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RecoverViewParams = RecoverViewParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigHost(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleAppNewSourceConfigHost, error) {
	model := &backuprecoveryv1.RecoverOracleAppNewSourceConfigHost{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigRecoverDatabaseParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverDatabaseParams, error) {
	model := &backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverDatabaseParams{}
	if modelMap["restore_time_usecs"] != nil {
		model.RestoreTimeUsecs = core.Int64Ptr(int64(modelMap["restore_time_usecs"].(int)))
	}
	if modelMap["db_channels"] != nil {
		dbChannels := []backuprecoveryv1.OracleDbChannel{}
		for _, dbChannelsItem := range modelMap["db_channels"].([]interface{}) {
			dbChannelsItemModel, err := resourceIbmRecoveryMapToOracleDbChannel(dbChannelsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			dbChannels = append(dbChannels, *dbChannelsItemModel)
		}
		model.DbChannels = dbChannels
	}
	if modelMap["recovery_mode"] != nil {
		model.RecoveryMode = core.BoolPtr(modelMap["recovery_mode"].(bool))
	}
	if modelMap["shell_evironment_vars"] != nil {
		shellEvironmentVars := []backuprecoveryv1.KeyValuePair{}
		for _, shellEvironmentVarsItem := range modelMap["shell_evironment_vars"].([]interface{}) {
			shellEvironmentVarsItemModel, err := resourceIbmRecoveryMapToKeyValuePair(shellEvironmentVarsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			shellEvironmentVars = append(shellEvironmentVars, *shellEvironmentVarsItemModel)
		}
		model.ShellEvironmentVars = shellEvironmentVars
	}
	if modelMap["granular_restore_info"] != nil && len(modelMap["granular_restore_info"].([]interface{})) > 0 {
		GranularRestoreInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigGranularRestoreInfo(modelMap["granular_restore_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.GranularRestoreInfo = GranularRestoreInfoModel
	}
	if modelMap["oracle_archive_log_info"] != nil && len(modelMap["oracle_archive_log_info"].([]interface{})) > 0 {
		OracleArchiveLogInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleArchiveLogInfo(modelMap["oracle_archive_log_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleArchiveLogInfo = OracleArchiveLogInfoModel
	}
	if modelMap["oracle_recovery_validation_info"] != nil && len(modelMap["oracle_recovery_validation_info"].([]interface{})) > 0 {
		OracleRecoveryValidationInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleRecoveryValidationInfo(modelMap["oracle_recovery_validation_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleRecoveryValidationInfo = OracleRecoveryValidationInfoModel
	}
	if modelMap["restore_spfile_or_pfile_info"] != nil && len(modelMap["restore_spfile_or_pfile_info"].([]interface{})) > 0 {
		RestoreSpfileOrPfileInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigRestoreSpfileOrPfileInfo(modelMap["restore_spfile_or_pfile_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RestoreSpfileOrPfileInfo = RestoreSpfileOrPfileInfoModel
	}
	if modelMap["use_scn_for_restore"] != nil {
		model.UseScnForRestore = core.BoolPtr(modelMap["use_scn_for_restore"].(bool))
	}
	if modelMap["database_name"] != nil && modelMap["database_name"].(string) != "" {
		model.DatabaseName = core.StringPtr(modelMap["database_name"].(string))
	}
	if modelMap["oracle_base_folder"] != nil && modelMap["oracle_base_folder"].(string) != "" {
		model.OracleBaseFolder = core.StringPtr(modelMap["oracle_base_folder"].(string))
	}
	if modelMap["oracle_home_folder"] != nil && modelMap["oracle_home_folder"].(string) != "" {
		model.OracleHomeFolder = core.StringPtr(modelMap["oracle_home_folder"].(string))
	}
	if modelMap["db_files_destination"] != nil && modelMap["db_files_destination"].(string) != "" {
		model.DbFilesDestination = core.StringPtr(modelMap["db_files_destination"].(string))
	}
	if modelMap["db_config_file_path"] != nil && modelMap["db_config_file_path"].(string) != "" {
		model.DbConfigFilePath = core.StringPtr(modelMap["db_config_file_path"].(string))
	}
	if modelMap["enable_archive_log_mode"] != nil {
		model.EnableArchiveLogMode = core.BoolPtr(modelMap["enable_archive_log_mode"].(bool))
	}
	if modelMap["pfile_parameter_map"] != nil {
		pfileParameterMap := []backuprecoveryv1.KeyValuePair{}
		for _, pfileParameterMapItem := range modelMap["pfile_parameter_map"].([]interface{}) {
			pfileParameterMapItemModel, err := resourceIbmRecoveryMapToKeyValuePair(pfileParameterMapItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			pfileParameterMap = append(pfileParameterMap, *pfileParameterMapItemModel)
		}
		model.PfileParameterMap = pfileParameterMap
	}
	if modelMap["bct_file_path"] != nil && modelMap["bct_file_path"].(string) != "" {
		model.BctFilePath = core.StringPtr(modelMap["bct_file_path"].(string))
	}
	if modelMap["num_tempfiles"] != nil {
		model.NumTempfiles = core.Int64Ptr(int64(modelMap["num_tempfiles"].(int)))
	}
	if modelMap["redo_log_config"] != nil && len(modelMap["redo_log_config"].([]interface{})) > 0 {
		RedoLogConfigModel, err := resourceIbmRecoveryMapToRecoverOracleNewTargetDatabaseConfigRedoLogConfig(modelMap["redo_log_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RedoLogConfig = RedoLogConfigModel
	}
	if modelMap["is_multi_stage_restore"] != nil {
		model.IsMultiStageRestore = core.BoolPtr(modelMap["is_multi_stage_restore"].(bool))
	}
	if modelMap["oracle_update_restore_options"] != nil && len(modelMap["oracle_update_restore_options"].([]interface{})) > 0 {
		OracleUpdateRestoreOptionsModel, err := resourceIbmRecoveryMapToRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions(modelMap["oracle_update_restore_options"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleUpdateRestoreOptions = OracleUpdateRestoreOptionsModel
	}
	if modelMap["skip_clone_nid"] != nil {
		model.SkipCloneNid = core.BoolPtr(modelMap["skip_clone_nid"].(bool))
	}
	if modelMap["no_filename_check"] != nil {
		model.NoFilenameCheck = core.BoolPtr(modelMap["no_filename_check"].(bool))
	}
	if modelMap["new_name_clause"] != nil && modelMap["new_name_clause"].(string) != "" {
		model.NewNameClause = core.StringPtr(modelMap["new_name_clause"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToOracleDbChannel(modelMap map[string]interface{}) (*backuprecoveryv1.OracleDbChannel, error) {
	model := &backuprecoveryv1.OracleDbChannel{}
	if modelMap["archive_log_retention_days"] != nil {
		model.ArchiveLogRetentionDays = core.Int64Ptr(int64(modelMap["archive_log_retention_days"].(int)))
	}
	if modelMap["archive_log_retention_hours"] != nil {
		model.ArchiveLogRetentionHours = core.Int64Ptr(int64(modelMap["archive_log_retention_hours"].(int)))
	}
	if modelMap["credentials"] != nil && len(modelMap["credentials"].([]interface{})) > 0 {
		CredentialsModel, err := resourceIbmRecoveryMapToCredentials(modelMap["credentials"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Credentials = CredentialsModel
	}
	if modelMap["database_unique_name"] != nil && modelMap["database_unique_name"].(string) != "" {
		model.DatabaseUniqueName = core.StringPtr(modelMap["database_unique_name"].(string))
	}
	if modelMap["database_uuid"] != nil && modelMap["database_uuid"].(string) != "" {
		model.DatabaseUUID = core.StringPtr(modelMap["database_uuid"].(string))
	}
	if modelMap["default_channel_count"] != nil {
		model.DefaultChannelCount = core.Int64Ptr(int64(modelMap["default_channel_count"].(int)))
	}
	if modelMap["database_node_list"] != nil {
		databaseNodeList := []backuprecoveryv1.OracleDatabaseHost{}
		for _, databaseNodeListItem := range modelMap["database_node_list"].([]interface{}) {
			databaseNodeListItemModel, err := resourceIbmRecoveryMapToOracleDatabaseHost(databaseNodeListItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			databaseNodeList = append(databaseNodeList, *databaseNodeListItemModel)
		}
		model.DatabaseNodeList = databaseNodeList
	}
	if modelMap["max_host_count"] != nil {
		model.MaxHostCount = core.Int64Ptr(int64(modelMap["max_host_count"].(int)))
	}
	if modelMap["enable_dg_primary_backup"] != nil {
		model.EnableDgPrimaryBackup = core.BoolPtr(modelMap["enable_dg_primary_backup"].(bool))
	}
	if modelMap["rman_backup_type"] != nil && modelMap["rman_backup_type"].(string) != "" {
		model.RmanBackupType = core.StringPtr(modelMap["rman_backup_type"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToCredentials(modelMap map[string]interface{}) (*backuprecoveryv1.Credentials, error) {
	model := &backuprecoveryv1.Credentials{}
	model.Username = core.StringPtr(modelMap["username"].(string))
	model.Password = core.StringPtr(modelMap["password"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToOracleDatabaseHost(modelMap map[string]interface{}) (*backuprecoveryv1.OracleDatabaseHost, error) {
	model := &backuprecoveryv1.OracleDatabaseHost{}
	if modelMap["host_id"] != nil && modelMap["host_id"].(string) != "" {
		model.HostID = core.StringPtr(modelMap["host_id"].(string))
	}
	if modelMap["channel_count"] != nil {
		model.ChannelCount = core.Int64Ptr(int64(modelMap["channel_count"].(int)))
	}
	if modelMap["port"] != nil {
		model.Port = core.Int64Ptr(int64(modelMap["port"].(int)))
	}
	if modelMap["sbt_host_params"] != nil && len(modelMap["sbt_host_params"].([]interface{})) > 0 {
		SbtHostParamsModel, err := resourceIbmRecoveryMapToOracleSbtHostParams(modelMap["sbt_host_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SbtHostParams = SbtHostParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToOracleSbtHostParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleSbtHostParams, error) {
	model := &backuprecoveryv1.OracleSbtHostParams{}
	if modelMap["sbt_library_path"] != nil && modelMap["sbt_library_path"].(string) != "" {
		model.SbtLibraryPath = core.StringPtr(modelMap["sbt_library_path"].(string))
	}
	if modelMap["view_fs_path"] != nil && modelMap["view_fs_path"].(string) != "" {
		model.ViewFsPath = core.StringPtr(modelMap["view_fs_path"].(string))
	}
	if modelMap["vip_list"] != nil {
		vipList := []string{}
		for _, vipListItem := range modelMap["vip_list"].([]interface{}) {
			vipList = append(vipList, vipListItem.(string))
		}
		model.VipList = vipList
	}
	if modelMap["vlan_info_list"] != nil {
		vlanInfoList := []backuprecoveryv1.OracleVlanInfo{}
		for _, vlanInfoListItem := range modelMap["vlan_info_list"].([]interface{}) {
			vlanInfoListItemModel, err := resourceIbmRecoveryMapToOracleVlanInfo(vlanInfoListItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			vlanInfoList = append(vlanInfoList, *vlanInfoListItemModel)
		}
		model.VlanInfoList = vlanInfoList
	}
	return model, nil
}

func resourceIbmRecoveryMapToOracleVlanInfo(modelMap map[string]interface{}) (*backuprecoveryv1.OracleVlanInfo, error) {
	model := &backuprecoveryv1.OracleVlanInfo{}
	if modelMap["ip_list"] != nil {
		ipList := []string{}
		for _, ipListItem := range modelMap["ip_list"].([]interface{}) {
			ipList = append(ipList, ipListItem.(string))
		}
		model.IpList = ipList
	}
	if modelMap["gateway"] != nil && modelMap["gateway"].(string) != "" {
		model.Gateway = core.StringPtr(modelMap["gateway"].(string))
	}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["subnet_ip"] != nil && modelMap["subnet_ip"].(string) != "" {
		model.SubnetIp = core.StringPtr(modelMap["subnet_ip"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToCommonOracleAppSourceConfigGranularRestoreInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonOracleAppSourceConfigGranularRestoreInfo, error) {
	model := &backuprecoveryv1.CommonOracleAppSourceConfigGranularRestoreInfo{}
	if modelMap["granularity_type"] != nil && modelMap["granularity_type"].(string) != "" {
		model.GranularityType = core.StringPtr(modelMap["granularity_type"].(string))
	}
	if modelMap["pdb_restore_params"] != nil && len(modelMap["pdb_restore_params"].([]interface{})) > 0 {
		PdbRestoreParamsModel, err := resourceIbmRecoveryMapToRecoverOracleGranularRestoreInfoPdbRestoreParams(modelMap["pdb_restore_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PdbRestoreParams = PdbRestoreParamsModel
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleGranularRestoreInfoPdbRestoreParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleGranularRestoreInfoPdbRestoreParams, error) {
	model := &backuprecoveryv1.RecoverOracleGranularRestoreInfoPdbRestoreParams{}
	if modelMap["drop_duplicate_pdb"] != nil {
		model.DropDuplicatePDB = core.BoolPtr(modelMap["drop_duplicate_pdb"].(bool))
	}
	if modelMap["pdb_objects"] != nil {
		pdbObjects := []backuprecoveryv1.OraclePdbObjectInfo{}
		for _, pdbObjectsItem := range modelMap["pdb_objects"].([]interface{}) {
			pdbObjectsItemModel, err := resourceIbmRecoveryMapToOraclePdbObjectInfo(pdbObjectsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			pdbObjects = append(pdbObjects, *pdbObjectsItemModel)
		}
		model.PdbObjects = pdbObjects
	}
	if modelMap["restore_to_existing_cdb"] != nil {
		model.RestoreToExistingCdb = core.BoolPtr(modelMap["restore_to_existing_cdb"].(bool))
	}
	if modelMap["rename_pdb_map"] != nil {
		renamePdbMap := []backuprecoveryv1.KeyValuePair{}
		for _, renamePdbMapItem := range modelMap["rename_pdb_map"].([]interface{}) {
			renamePdbMapItemModel, err := resourceIbmRecoveryMapToKeyValuePair(renamePdbMapItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			renamePdbMap = append(renamePdbMap, *renamePdbMapItemModel)
		}
		model.RenamePdbMap = renamePdbMap
	}
	if modelMap["include_in_restore"] != nil {
		model.IncludeInRestore = core.BoolPtr(modelMap["include_in_restore"].(bool))
	}
	return model, nil
}

func resourceIbmRecoveryMapToOraclePdbObjectInfo(modelMap map[string]interface{}) (*backuprecoveryv1.OraclePdbObjectInfo, error) {
	model := &backuprecoveryv1.OraclePdbObjectInfo{}
	model.DbID = core.StringPtr(modelMap["db_id"].(string))
	model.DbName = core.StringPtr(modelMap["db_name"].(string))
	return model, nil
}

func resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleArchiveLogInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonOracleAppSourceConfigOracleArchiveLogInfo, error) {
	model := &backuprecoveryv1.CommonOracleAppSourceConfigOracleArchiveLogInfo{}
	if modelMap["range_type"] != nil && modelMap["range_type"].(string) != "" {
		model.RangeType = core.StringPtr(modelMap["range_type"].(string))
	}
	if modelMap["range_info_vec"] != nil {
		rangeInfoVec := []backuprecoveryv1.OracleRangeMetaInfo{}
		for _, rangeInfoVecItem := range modelMap["range_info_vec"].([]interface{}) {
			rangeInfoVecItemModel, err := resourceIbmRecoveryMapToOracleRangeMetaInfo(rangeInfoVecItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			rangeInfoVec = append(rangeInfoVec, *rangeInfoVecItemModel)
		}
		model.RangeInfoVec = rangeInfoVec
	}
	if modelMap["archive_log_restore_dest"] != nil && modelMap["archive_log_restore_dest"].(string) != "" {
		model.ArchiveLogRestoreDest = core.StringPtr(modelMap["archive_log_restore_dest"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToOracleRangeMetaInfo(modelMap map[string]interface{}) (*backuprecoveryv1.OracleRangeMetaInfo, error) {
	model := &backuprecoveryv1.OracleRangeMetaInfo{}
	if modelMap["start_of_range"] != nil {
		model.StartOfRange = core.Int64Ptr(int64(modelMap["start_of_range"].(int)))
	}
	if modelMap["end_of_range"] != nil {
		model.EndOfRange = core.Int64Ptr(int64(modelMap["end_of_range"].(int)))
	}
	if modelMap["protection_group_id"] != nil && modelMap["protection_group_id"].(string) != "" {
		model.ProtectionGroupID = core.StringPtr(modelMap["protection_group_id"].(string))
	}
	if modelMap["reset_log_id"] != nil {
		model.ResetLogID = core.Int64Ptr(int64(modelMap["reset_log_id"].(int)))
	}
	if modelMap["incarnation_id"] != nil {
		model.IncarnationID = core.Int64Ptr(int64(modelMap["incarnation_id"].(int)))
	}
	if modelMap["thread_id"] != nil {
		model.ThreadID = core.Int64Ptr(int64(modelMap["thread_id"].(int)))
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleRecoveryValidationInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonOracleAppSourceConfigOracleRecoveryValidationInfo, error) {
	model := &backuprecoveryv1.CommonOracleAppSourceConfigOracleRecoveryValidationInfo{}
	if modelMap["create_dummy_instance"] != nil {
		model.CreateDummyInstance = core.BoolPtr(modelMap["create_dummy_instance"].(bool))
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonOracleAppSourceConfigRestoreSpfileOrPfileInfo(modelMap map[string]interface{}) (*backuprecoveryv1.CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo, error) {
	model := &backuprecoveryv1.CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo{}
	if modelMap["should_restore_spfile_or_pfile"] != nil {
		model.ShouldRestoreSpfileOrPfile = core.BoolPtr(modelMap["should_restore_spfile_or_pfile"].(bool))
	}
	if modelMap["file_location"] != nil && modelMap["file_location"].(string) != "" {
		model.FileLocation = core.StringPtr(modelMap["file_location"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleNewTargetDatabaseConfigRedoLogConfig(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigRedoLogConfig, error) {
	model := &backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigRedoLogConfig{}
	if modelMap["num_groups"] != nil {
		model.NumGroups = core.Int64Ptr(int64(modelMap["num_groups"].(int)))
	}
	if modelMap["member_prefix"] != nil && modelMap["member_prefix"].(string) != "" {
		model.MemberPrefix = core.StringPtr(modelMap["member_prefix"].(string))
	}
	if modelMap["size_m_bytes"] != nil {
		model.SizeMBytes = core.Int64Ptr(int64(modelMap["size_m_bytes"].(int)))
	}
	if modelMap["group_members"] != nil {
		groupMembers := []string{}
		for _, groupMembersItem := range modelMap["group_members"].([]interface{}) {
			groupMembers = append(groupMembers, groupMembersItem.(string))
		}
		model.GroupMembers = groupMembers
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions, error) {
	model := &backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions{}
	if modelMap["delay_secs"] != nil {
		model.DelaySecs = core.Int64Ptr(int64(modelMap["delay_secs"].(int)))
	}
	if modelMap["target_path_vec"] != nil {
		targetPathVec := []string{}
		for _, targetPathVecItem := range modelMap["target_path_vec"].([]interface{}) {
			targetPathVec = append(targetPathVec, targetPathVecItem.(string))
		}
		model.TargetPathVec = targetPathVec
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleAppNewSourceConfigRecoverViewParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverViewParams, error) {
	model := &backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverViewParams{}
	if modelMap["restore_time_usecs"] != nil {
		model.RestoreTimeUsecs = core.Int64Ptr(int64(modelMap["restore_time_usecs"].(int)))
	}
	if modelMap["db_channels"] != nil {
		dbChannels := []backuprecoveryv1.OracleDbChannel{}
		for _, dbChannelsItem := range modelMap["db_channels"].([]interface{}) {
			dbChannelsItemModel, err := resourceIbmRecoveryMapToOracleDbChannel(dbChannelsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			dbChannels = append(dbChannels, *dbChannelsItemModel)
		}
		model.DbChannels = dbChannels
	}
	if modelMap["recovery_mode"] != nil {
		model.RecoveryMode = core.BoolPtr(modelMap["recovery_mode"].(bool))
	}
	if modelMap["shell_evironment_vars"] != nil {
		shellEvironmentVars := []backuprecoveryv1.KeyValuePair{}
		for _, shellEvironmentVarsItem := range modelMap["shell_evironment_vars"].([]interface{}) {
			shellEvironmentVarsItemModel, err := resourceIbmRecoveryMapToKeyValuePair(shellEvironmentVarsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			shellEvironmentVars = append(shellEvironmentVars, *shellEvironmentVarsItemModel)
		}
		model.ShellEvironmentVars = shellEvironmentVars
	}
	if modelMap["granular_restore_info"] != nil && len(modelMap["granular_restore_info"].([]interface{})) > 0 {
		GranularRestoreInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigGranularRestoreInfo(modelMap["granular_restore_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.GranularRestoreInfo = GranularRestoreInfoModel
	}
	if modelMap["oracle_archive_log_info"] != nil && len(modelMap["oracle_archive_log_info"].([]interface{})) > 0 {
		OracleArchiveLogInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleArchiveLogInfo(modelMap["oracle_archive_log_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleArchiveLogInfo = OracleArchiveLogInfoModel
	}
	if modelMap["oracle_recovery_validation_info"] != nil && len(modelMap["oracle_recovery_validation_info"].([]interface{})) > 0 {
		OracleRecoveryValidationInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleRecoveryValidationInfo(modelMap["oracle_recovery_validation_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleRecoveryValidationInfo = OracleRecoveryValidationInfoModel
	}
	if modelMap["restore_spfile_or_pfile_info"] != nil && len(modelMap["restore_spfile_or_pfile_info"].([]interface{})) > 0 {
		RestoreSpfileOrPfileInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigRestoreSpfileOrPfileInfo(modelMap["restore_spfile_or_pfile_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RestoreSpfileOrPfileInfo = RestoreSpfileOrPfileInfoModel
	}
	if modelMap["use_scn_for_restore"] != nil {
		model.UseScnForRestore = core.BoolPtr(modelMap["use_scn_for_restore"].(bool))
	}
	if modelMap["view_mount_path"] != nil && modelMap["view_mount_path"].(string) != "" {
		model.ViewMountPath = core.StringPtr(modelMap["view_mount_path"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryMapToCommonRecoverOracleAppTargetParamsOriginalSourceConfig(modelMap map[string]interface{}) (*backuprecoveryv1.CommonRecoverOracleAppTargetParamsOriginalSourceConfig, error) {
	model := &backuprecoveryv1.CommonRecoverOracleAppTargetParamsOriginalSourceConfig{}
	if modelMap["restore_time_usecs"] != nil {
		model.RestoreTimeUsecs = core.Int64Ptr(int64(modelMap["restore_time_usecs"].(int)))
	}
	if modelMap["db_channels"] != nil {
		dbChannels := []backuprecoveryv1.OracleDbChannel{}
		for _, dbChannelsItem := range modelMap["db_channels"].([]interface{}) {
			dbChannelsItemModel, err := resourceIbmRecoveryMapToOracleDbChannel(dbChannelsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			dbChannels = append(dbChannels, *dbChannelsItemModel)
		}
		model.DbChannels = dbChannels
	}
	if modelMap["recovery_mode"] != nil {
		model.RecoveryMode = core.BoolPtr(modelMap["recovery_mode"].(bool))
	}
	if modelMap["shell_evironment_vars"] != nil {
		shellEvironmentVars := []backuprecoveryv1.KeyValuePair{}
		for _, shellEvironmentVarsItem := range modelMap["shell_evironment_vars"].([]interface{}) {
			shellEvironmentVarsItemModel, err := resourceIbmRecoveryMapToKeyValuePair(shellEvironmentVarsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			shellEvironmentVars = append(shellEvironmentVars, *shellEvironmentVarsItemModel)
		}
		model.ShellEvironmentVars = shellEvironmentVars
	}
	if modelMap["granular_restore_info"] != nil && len(modelMap["granular_restore_info"].([]interface{})) > 0 {
		GranularRestoreInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigGranularRestoreInfo(modelMap["granular_restore_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.GranularRestoreInfo = GranularRestoreInfoModel
	}
	if modelMap["oracle_archive_log_info"] != nil && len(modelMap["oracle_archive_log_info"].([]interface{})) > 0 {
		OracleArchiveLogInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleArchiveLogInfo(modelMap["oracle_archive_log_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleArchiveLogInfo = OracleArchiveLogInfoModel
	}
	if modelMap["oracle_recovery_validation_info"] != nil && len(modelMap["oracle_recovery_validation_info"].([]interface{})) > 0 {
		OracleRecoveryValidationInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigOracleRecoveryValidationInfo(modelMap["oracle_recovery_validation_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleRecoveryValidationInfo = OracleRecoveryValidationInfoModel
	}
	if modelMap["restore_spfile_or_pfile_info"] != nil && len(modelMap["restore_spfile_or_pfile_info"].([]interface{})) > 0 {
		RestoreSpfileOrPfileInfoModel, err := resourceIbmRecoveryMapToCommonOracleAppSourceConfigRestoreSpfileOrPfileInfo(modelMap["restore_spfile_or_pfile_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RestoreSpfileOrPfileInfo = RestoreSpfileOrPfileInfoModel
	}
	if modelMap["use_scn_for_restore"] != nil {
		model.UseScnForRestore = core.BoolPtr(modelMap["use_scn_for_restore"].(bool))
	}
	if modelMap["roll_forward_log_path_vec"] != nil {
		rollForwardLogPathVec := []string{}
		for _, rollForwardLogPathVecItem := range modelMap["roll_forward_log_path_vec"].([]interface{}) {
			rollForwardLogPathVec = append(rollForwardLogPathVec, rollForwardLogPathVecItem.(string))
		}
		model.RollForwardLogPathVec = rollForwardLogPathVec
	}
	if modelMap["attempt_complete_recovery"] != nil {
		model.AttemptCompleteRecovery = core.BoolPtr(modelMap["attempt_complete_recovery"].(bool))
	}
	if modelMap["roll_forward_time_msecs"] != nil {
		model.RollForwardTimeMsecs = core.Int64Ptr(int64(modelMap["roll_forward_time_msecs"].(int)))
	}
	if modelMap["stop_active_passive"] != nil {
		model.StopActivePassive = core.BoolPtr(modelMap["stop_active_passive"].(bool))
	}
	return model, nil
}

func resourceIbmRecoveryMapToRecoverOracleAppParamsVlanConfig(modelMap map[string]interface{}) (*backuprecoveryv1.RecoverOracleAppParamsVlanConfig, error) {
	model := &backuprecoveryv1.RecoverOracleAppParamsVlanConfig{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	return model, nil
}

func resourceIbmRecoveryRecoverPhysicalParamsToMap(model *backuprecoveryv1.RecoverPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := resourceIbmRecoveryCommonRecoverObjectSnapshotParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	modelMap["recovery_action"] = model.RecoveryAction
	if model.RecoverVolumeParams != nil {
		recoverVolumeParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsRecoverVolumeParamsToMap(model.RecoverVolumeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_volume_params"] = []map[string]interface{}{recoverVolumeParamsMap}
	}
	if model.MountVolumeParams != nil {
		mountVolumeParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsMountVolumeParamsToMap(model.MountVolumeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["mount_volume_params"] = []map[string]interface{}{mountVolumeParamsMap}
	}
	if model.RecoverFileAndFolderParams != nil {
		recoverFileAndFolderParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(model.RecoverFileAndFolderParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_file_and_folder_params"] = []map[string]interface{}{recoverFileAndFolderParamsMap}
	}
	if model.DownloadFileAndFolderParams != nil {
		downloadFileAndFolderParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(model.DownloadFileAndFolderParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["download_file_and_folder_params"] = []map[string]interface{}{downloadFileAndFolderParamsMap}
	}
	if model.SystemRecoveryParams != nil {
		systemRecoveryParamsMap, err := resourceIbmRecoveryRecoverPhysicalParamsSystemRecoveryParamsToMap(model.SystemRecoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_recovery_params"] = []map[string]interface{}{systemRecoveryParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonRecoverObjectSnapshotParamsToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParams) (map[string]interface{}, error) {
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
		objectInfoMap, err := resourceIbmRecoveryCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
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
		archivalTargetInfoMap, err := resourceIbmRecoveryCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
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

func resourceIbmRecoveryCommonRecoverObjectSnapshotParamsObjectInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo) (map[string]interface{}, error) {
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

func resourceIbmRecoveryCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := resourceIbmRecoveryArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := resourceIbmRecoveryOracleTiersToMap(model.OracleTiering)
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

func resourceIbmRecoveryOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := resourceIbmRecoveryOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func resourceIbmRecoveryOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func resourceIbmRecoveryRecoverPhysicalParamsRecoverVolumeParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := resourceIbmRecoveryRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	mountTargetMap, err := resourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeMountTargetToMap(model.MountTarget)
	if err != nil {
		return modelMap, err
	}
	modelMap["mount_target"] = []map[string]interface{}{mountTargetMap}
	volumeMapping := []map[string]interface{}{}
	for _, volumeMappingItem := range model.VolumeMapping {
		volumeMappingItemMap, err := resourceIbmRecoveryRecoverVolumeMappingToMap(&volumeMappingItem)
		if err != nil {
			return modelMap, err
		}
		volumeMapping = append(volumeMapping, volumeMappingItemMap)
	}
	modelMap["volume_mapping"] = volumeMapping
	if model.ForceUnmountVolume != nil {
		modelMap["force_unmount_volume"] = model.ForceUnmountVolume
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := resourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeMountTargetToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverVolumeMappingToMap(model *backuprecoveryv1.RecoverVolumeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_volume_guid"] = model.SourceVolumeGuid
	modelMap["destination_volume_guid"] = model.DestinationVolumeGuid
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = model.InterfaceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalParamsMountVolumeParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := resourceIbmRecoveryMountPhysicalVolumeParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryMountPhysicalVolumeParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mount_to_original_target"] = model.MountToOriginalTarget
	if model.OriginalTargetConfig != nil {
		originalTargetConfigMap, err := resourceIbmRecoveryPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(model.OriginalTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_target_config"] = []map[string]interface{}{originalTargetConfigMap}
	}
	if model.NewTargetConfig != nil {
		newTargetConfigMap, err := resourceIbmRecoveryPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(model.NewTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["new_target_config"] = []map[string]interface{}{newTargetConfigMap}
	}
	if model.ReadOnlyMount != nil {
		modelMap["read_only_mount"] = model.ReadOnlyMount
	}
	if model.VolumeNames != nil {
		modelMap["volume_names"] = model.VolumeNames
	}
	if model.MountedVolumeMapping != nil {
		mountedVolumeMapping := []map[string]interface{}{}
		for _, mountedVolumeMappingItem := range model.MountedVolumeMapping {
			mountedVolumeMappingItemMap, err := resourceIbmRecoveryMountedVolumeMappingToMap(&mountedVolumeMappingItem)
			if err != nil {
				return modelMap, err
			}
			mountedVolumeMapping = append(mountedVolumeMapping, mountedVolumeMappingItemMap)
		}
		modelMap["mounted_volume_mapping"] = mountedVolumeMapping
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := resourceIbmRecoveryPhysicalTargetParamsForMountVolumeVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ServerCredentials != nil {
		serverCredentialsMap, err := resourceIbmRecoveryPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(model.ServerCredentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["server_credentials"] = []map[string]interface{}{serverCredentialsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(model *backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	mountTargetMap, err := resourceIbmRecoveryRecoverTargetToMap(model.MountTarget)
	if err != nil {
		return modelMap, err
	}
	modelMap["mount_target"] = []map[string]interface{}{mountTargetMap}
	if model.ServerCredentials != nil {
		serverCredentialsMap, err := resourceIbmRecoveryPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(model.ServerCredentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["server_credentials"] = []map[string]interface{}{serverCredentialsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverTargetToMap(model *backuprecoveryv1.RecoverTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ParentSourceID != nil {
		modelMap["parent_source_id"] = flex.IntValue(model.ParentSourceID)
	}
	if model.ParentSourceName != nil {
		modelMap["parent_source_name"] = model.ParentSourceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(model *backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func resourceIbmRecoveryMountedVolumeMappingToMap(model *backuprecoveryv1.MountedVolumeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.OriginalVolume != nil {
		modelMap["original_volume"] = model.OriginalVolume
	}
	if model.MountedVolume != nil {
		modelMap["mounted_volume"] = model.MountedVolume
	}
	if model.FileSystemType != nil {
		modelMap["file_system_type"] = model.FileSystemType
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForMountVolumeVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = model.InterfaceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	filesAndFolders := []map[string]interface{}{}
	for _, filesAndFoldersItem := range model.FilesAndFolders {
		filesAndFoldersItemMap, err := resourceIbmRecoveryCommonRecoverFileAndFolderInfoToMap(&filesAndFoldersItem)
		if err != nil {
			return modelMap, err
		}
		filesAndFolders = append(filesAndFolders, filesAndFoldersItemMap)
	}
	modelMap["files_and_folders"] = filesAndFolders
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := resourceIbmRecoveryRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonRecoverFileAndFolderInfoToMap(model *backuprecoveryv1.CommonRecoverFileAndFolderInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["absolute_path"] = model.AbsolutePath
	if model.DestinationDir != nil {
		modelMap["destination_dir"] = model.DestinationDir
	}
	if model.IsDirectory != nil {
		modelMap["is_directory"] = model.IsDirectory
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Messages != nil {
		modelMap["messages"] = model.Messages
	}
	if model.IsViewFileRecovery != nil {
		modelMap["is_view_file_recovery"] = model.IsViewFileRecovery
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	recoverTargetMap, err := resourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(model.RecoverTarget)
	if err != nil {
		return modelMap, err
	}
	modelMap["recover_target"] = []map[string]interface{}{recoverTargetMap}
	if model.RestoreToOriginalPaths != nil {
		modelMap["restore_to_original_paths"] = model.RestoreToOriginalPaths
	}
	if model.OverwriteExisting != nil {
		modelMap["overwrite_existing"] = model.OverwriteExisting
	}
	if model.AlternateRestoreDirectory != nil {
		modelMap["alternate_restore_directory"] = model.AlternateRestoreDirectory
	}
	if model.PreserveAttributes != nil {
		modelMap["preserve_attributes"] = model.PreserveAttributes
	}
	if model.PreserveTimestamps != nil {
		modelMap["preserve_timestamps"] = model.PreserveTimestamps
	}
	if model.PreserveAcls != nil {
		modelMap["preserve_acls"] = model.PreserveAcls
	}
	if model.ContinueOnError != nil {
		modelMap["continue_on_error"] = model.ContinueOnError
	}
	if model.SaveSuccessFiles != nil {
		modelMap["save_success_files"] = model.SaveSuccessFiles
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := resourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	if model.RestoreEntityType != nil {
		modelMap["restore_entity_type"] = model.RestoreEntityType
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ParentSourceID != nil {
		modelMap["parent_source_id"] = flex.IntValue(model.ParentSourceID)
	}
	if model.ParentSourceName != nil {
		modelMap["parent_source_name"] = model.ParentSourceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = model.InterfaceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FilesAndFolders != nil {
		filesAndFolders := []map[string]interface{}{}
		for _, filesAndFoldersItem := range model.FilesAndFolders {
			filesAndFoldersItemMap, err := resourceIbmRecoveryCommonRecoverFileAndFolderInfoToMap(&filesAndFoldersItem)
			if err != nil {
				return modelMap, err
			}
			filesAndFolders = append(filesAndFolders, filesAndFoldersItemMap)
		}
		modelMap["files_and_folders"] = filesAndFolders
	}
	if model.DownloadFilePath != nil {
		modelMap["download_file_path"] = model.DownloadFilePath
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverPhysicalParamsSystemRecoveryParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FullNasPath != nil {
		modelMap["full_nas_path"] = model.FullNasPath
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleParamsToMap(model *backuprecoveryv1.RecoverOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := resourceIbmRecoveryRecoverOracleDbSnapshotParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	modelMap["recovery_action"] = model.RecoveryAction
	if model.RecoverAppParams != nil {
		recoverAppParamsMap, err := resourceIbmRecoveryRecoverOracleParamsRecoverAppParamsToMap(model.RecoverAppParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_app_params"] = []map[string]interface{}{recoverAppParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleDbSnapshotParamsToMap(model *backuprecoveryv1.RecoverOracleDbSnapshotParams) (map[string]interface{}, error) {
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
		objectInfoMap, err := resourceIbmRecoveryCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
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
		archivalTargetInfoMap, err := resourceIbmRecoveryCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
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
	if model.InstantRecoveryInfo != nil {
		instantRecoveryInfoMap, err := resourceIbmRecoveryRecoverOracleDbSnapshotParamsInstantRecoveryInfoToMap(model.InstantRecoveryInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["instant_recovery_info"] = []map[string]interface{}{instantRecoveryInfoMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleDbSnapshotParamsInstantRecoveryInfoToMap(model *backuprecoveryv1.RecoverOracleDbSnapshotParamsInstantRecoveryInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
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
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleParamsRecoverAppParamsToMap(model *backuprecoveryv1.RecoverOracleParamsRecoverAppParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.OracleTargetParams != nil {
		oracleTargetParamsMap, err := resourceIbmRecoveryRecoverOracleAppParamsOracleTargetParamsToMap(model.OracleTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_target_params"] = []map[string]interface{}{oracleTargetParamsMap}
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := resourceIbmRecoveryRecoverOracleAppParamsVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleAppParamsOracleTargetParamsToMap(model *backuprecoveryv1.RecoverOracleAppParamsOracleTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["recover_to_new_source"] = model.RecoverToNewSource
	if model.NewSourceConfig != nil {
		newSourceConfigMap, err := resourceIbmRecoveryCommonRecoverOracleAppTargetParamsNewSourceConfigToMap(model.NewSourceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["new_source_config"] = []map[string]interface{}{newSourceConfigMap}
	}
	if model.OriginalSourceConfig != nil {
		originalSourceConfigMap, err := resourceIbmRecoveryCommonRecoverOracleAppTargetParamsOriginalSourceConfigToMap(model.OriginalSourceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_source_config"] = []map[string]interface{}{originalSourceConfigMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonRecoverOracleAppTargetParamsNewSourceConfigToMap(model *backuprecoveryv1.CommonRecoverOracleAppTargetParamsNewSourceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	hostMap, err := resourceIbmRecoveryRecoverOracleAppNewSourceConfigHostToMap(model.Host)
	if err != nil {
		return modelMap, err
	}
	modelMap["host"] = []map[string]interface{}{hostMap}
	if model.RecoveryTarget != nil {
		modelMap["recovery_target"] = model.RecoveryTarget
	}
	if model.RecoverDatabaseParams != nil {
		recoverDatabaseParamsMap, err := resourceIbmRecoveryRecoverOracleAppNewSourceConfigRecoverDatabaseParamsToMap(model.RecoverDatabaseParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_database_params"] = []map[string]interface{}{recoverDatabaseParamsMap}
	}
	if model.RecoverViewParams != nil {
		recoverViewParamsMap, err := resourceIbmRecoveryRecoverOracleAppNewSourceConfigRecoverViewParamsToMap(model.RecoverViewParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_view_params"] = []map[string]interface{}{recoverViewParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleAppNewSourceConfigHostToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigHost) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleAppNewSourceConfigRecoverDatabaseParamsToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverDatabaseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := resourceIbmRecoveryOracleDbChannelToMap(&dbChannelsItem)
			if err != nil {
				return modelMap, err
			}
			dbChannels = append(dbChannels, dbChannelsItemMap)
		}
		modelMap["db_channels"] = dbChannels
	}
	if model.RecoveryMode != nil {
		modelMap["recovery_mode"] = model.RecoveryMode
	}
	if model.ShellEvironmentVars != nil {
		shellEvironmentVars := []map[string]interface{}{}
		for _, shellEvironmentVarsItem := range model.ShellEvironmentVars {
			shellEvironmentVarsItemMap, err := resourceIbmRecoveryKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["restore_spfile_or_pfile_info"] = []map[string]interface{}{restoreSpfileOrPfileInfoMap}
	}
	if model.UseScnForRestore != nil {
		modelMap["use_scn_for_restore"] = model.UseScnForRestore
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	if model.OracleBaseFolder != nil {
		modelMap["oracle_base_folder"] = model.OracleBaseFolder
	}
	if model.OracleHomeFolder != nil {
		modelMap["oracle_home_folder"] = model.OracleHomeFolder
	}
	if model.DbFilesDestination != nil {
		modelMap["db_files_destination"] = model.DbFilesDestination
	}
	if model.DbConfigFilePath != nil {
		modelMap["db_config_file_path"] = model.DbConfigFilePath
	}
	if model.EnableArchiveLogMode != nil {
		modelMap["enable_archive_log_mode"] = model.EnableArchiveLogMode
	}
	if model.PfileParameterMap != nil {
		pfileParameterMap := []map[string]interface{}{}
		for _, pfileParameterMapItem := range model.PfileParameterMap {
			pfileParameterMapItemMap, err := resourceIbmRecoveryKeyValuePairToMap(&pfileParameterMapItem)
			if err != nil {
				return modelMap, err
			}
			pfileParameterMap = append(pfileParameterMap, pfileParameterMapItemMap)
		}
		modelMap["pfile_parameter_map"] = pfileParameterMap
	}
	if model.BctFilePath != nil {
		modelMap["bct_file_path"] = model.BctFilePath
	}
	if model.NumTempfiles != nil {
		modelMap["num_tempfiles"] = flex.IntValue(model.NumTempfiles)
	}
	if model.RedoLogConfig != nil {
		redoLogConfigMap, err := resourceIbmRecoveryRecoverOracleNewTargetDatabaseConfigRedoLogConfigToMap(model.RedoLogConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["redo_log_config"] = []map[string]interface{}{redoLogConfigMap}
	}
	if model.IsMultiStageRestore != nil {
		modelMap["is_multi_stage_restore"] = model.IsMultiStageRestore
	}
	if model.OracleUpdateRestoreOptions != nil {
		oracleUpdateRestoreOptionsMap, err := resourceIbmRecoveryRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptionsToMap(model.OracleUpdateRestoreOptions)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_update_restore_options"] = []map[string]interface{}{oracleUpdateRestoreOptionsMap}
	}
	if model.SkipCloneNid != nil {
		modelMap["skip_clone_nid"] = model.SkipCloneNid
	}
	if model.NoFilenameCheck != nil {
		modelMap["no_filename_check"] = model.NoFilenameCheck
	}
	if model.NewNameClause != nil {
		modelMap["new_name_clause"] = model.NewNameClause
	}
	return modelMap, nil
}

func resourceIbmRecoveryOracleDbChannelToMap(model *backuprecoveryv1.OracleDbChannel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchiveLogRetentionDays != nil {
		modelMap["archive_log_retention_days"] = flex.IntValue(model.ArchiveLogRetentionDays)
	}
	if model.ArchiveLogRetentionHours != nil {
		modelMap["archive_log_retention_hours"] = flex.IntValue(model.ArchiveLogRetentionHours)
	}
	if model.Credentials != nil {
		credentialsMap, err := resourceIbmRecoveryCredentialsToMap(model.Credentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["credentials"] = []map[string]interface{}{credentialsMap}
	}
	if model.DatabaseUniqueName != nil {
		modelMap["database_unique_name"] = model.DatabaseUniqueName
	}
	if model.DatabaseUUID != nil {
		modelMap["database_uuid"] = model.DatabaseUUID
	}
	if model.DefaultChannelCount != nil {
		modelMap["default_channel_count"] = flex.IntValue(model.DefaultChannelCount)
	}
	if model.DatabaseNodeList != nil {
		databaseNodeList := []map[string]interface{}{}
		for _, databaseNodeListItem := range model.DatabaseNodeList {
			databaseNodeListItemMap, err := resourceIbmRecoveryOracleDatabaseHostToMap(&databaseNodeListItem)
			if err != nil {
				return modelMap, err
			}
			databaseNodeList = append(databaseNodeList, databaseNodeListItemMap)
		}
		modelMap["database_node_list"] = databaseNodeList
	}
	if model.MaxHostCount != nil {
		modelMap["max_host_count"] = flex.IntValue(model.MaxHostCount)
	}
	if model.EnableDgPrimaryBackup != nil {
		modelMap["enable_dg_primary_backup"] = model.EnableDgPrimaryBackup
	}
	if model.RmanBackupType != nil {
		modelMap["rman_backup_type"] = model.RmanBackupType
	}
	return modelMap, nil
}

func resourceIbmRecoveryCredentialsToMap(model *backuprecoveryv1.Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func resourceIbmRecoveryOracleDatabaseHostToMap(model *backuprecoveryv1.OracleDatabaseHost) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.HostID != nil {
		modelMap["host_id"] = model.HostID
	}
	if model.ChannelCount != nil {
		modelMap["channel_count"] = flex.IntValue(model.ChannelCount)
	}
	if model.Port != nil {
		modelMap["port"] = flex.IntValue(model.Port)
	}
	if model.SbtHostParams != nil {
		sbtHostParamsMap, err := resourceIbmRecoveryOracleSbtHostParamsToMap(model.SbtHostParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sbt_host_params"] = []map[string]interface{}{sbtHostParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryOracleSbtHostParamsToMap(model *backuprecoveryv1.OracleSbtHostParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SbtLibraryPath != nil {
		modelMap["sbt_library_path"] = model.SbtLibraryPath
	}
	if model.ViewFsPath != nil {
		modelMap["view_fs_path"] = model.ViewFsPath
	}
	if model.VipList != nil {
		modelMap["vip_list"] = model.VipList
	}
	if model.VlanInfoList != nil {
		vlanInfoList := []map[string]interface{}{}
		for _, vlanInfoListItem := range model.VlanInfoList {
			vlanInfoListItemMap, err := resourceIbmRecoveryOracleVlanInfoToMap(&vlanInfoListItem)
			if err != nil {
				return modelMap, err
			}
			vlanInfoList = append(vlanInfoList, vlanInfoListItemMap)
		}
		modelMap["vlan_info_list"] = vlanInfoList
	}
	return modelMap, nil
}

func resourceIbmRecoveryOracleVlanInfoToMap(model *backuprecoveryv1.OracleVlanInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IpList != nil {
		modelMap["ip_list"] = model.IpList
	}
	if model.Gateway != nil {
		modelMap["gateway"] = model.Gateway
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.SubnetIp != nil {
		modelMap["subnet_ip"] = model.SubnetIp
	}
	return modelMap, nil
}

func resourceIbmRecoveryKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmRecoveryCommonOracleAppSourceConfigGranularRestoreInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigGranularRestoreInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GranularityType != nil {
		modelMap["granularity_type"] = model.GranularityType
	}
	if model.PdbRestoreParams != nil {
		pdbRestoreParamsMap, err := resourceIbmRecoveryRecoverOracleGranularRestoreInfoPdbRestoreParamsToMap(model.PdbRestoreParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["pdb_restore_params"] = []map[string]interface{}{pdbRestoreParamsMap}
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleGranularRestoreInfoPdbRestoreParamsToMap(model *backuprecoveryv1.RecoverOracleGranularRestoreInfoPdbRestoreParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DropDuplicatePDB != nil {
		modelMap["drop_duplicate_pdb"] = model.DropDuplicatePDB
	}
	if model.PdbObjects != nil {
		pdbObjects := []map[string]interface{}{}
		for _, pdbObjectsItem := range model.PdbObjects {
			pdbObjectsItemMap, err := resourceIbmRecoveryOraclePdbObjectInfoToMap(&pdbObjectsItem)
			if err != nil {
				return modelMap, err
			}
			pdbObjects = append(pdbObjects, pdbObjectsItemMap)
		}
		modelMap["pdb_objects"] = pdbObjects
	}
	if model.RestoreToExistingCdb != nil {
		modelMap["restore_to_existing_cdb"] = model.RestoreToExistingCdb
	}
	if model.RenamePdbMap != nil {
		renamePdbMap := []map[string]interface{}{}
		for _, renamePdbMapItem := range model.RenamePdbMap {
			renamePdbMapItemMap, err := resourceIbmRecoveryKeyValuePairToMap(&renamePdbMapItem)
			if err != nil {
				return modelMap, err
			}
			renamePdbMap = append(renamePdbMap, renamePdbMapItemMap)
		}
		modelMap["rename_pdb_map"] = renamePdbMap
	}
	if model.IncludeInRestore != nil {
		modelMap["include_in_restore"] = model.IncludeInRestore
	}
	return modelMap, nil
}

func resourceIbmRecoveryOraclePdbObjectInfoToMap(model *backuprecoveryv1.OraclePdbObjectInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["db_id"] = model.DbID
	modelMap["db_name"] = model.DbName
	return modelMap, nil
}

func resourceIbmRecoveryCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigOracleArchiveLogInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RangeType != nil {
		modelMap["range_type"] = model.RangeType
	}
	if model.RangeInfoVec != nil {
		rangeInfoVec := []map[string]interface{}{}
		for _, rangeInfoVecItem := range model.RangeInfoVec {
			rangeInfoVecItemMap, err := resourceIbmRecoveryOracleRangeMetaInfoToMap(&rangeInfoVecItem)
			if err != nil {
				return modelMap, err
			}
			rangeInfoVec = append(rangeInfoVec, rangeInfoVecItemMap)
		}
		modelMap["range_info_vec"] = rangeInfoVec
	}
	if model.ArchiveLogRestoreDest != nil {
		modelMap["archive_log_restore_dest"] = model.ArchiveLogRestoreDest
	}
	return modelMap, nil
}

func resourceIbmRecoveryOracleRangeMetaInfoToMap(model *backuprecoveryv1.OracleRangeMetaInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StartOfRange != nil {
		modelMap["start_of_range"] = flex.IntValue(model.StartOfRange)
	}
	if model.EndOfRange != nil {
		modelMap["end_of_range"] = flex.IntValue(model.EndOfRange)
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ResetLogID != nil {
		modelMap["reset_log_id"] = flex.IntValue(model.ResetLogID)
	}
	if model.IncarnationID != nil {
		modelMap["incarnation_id"] = flex.IntValue(model.IncarnationID)
	}
	if model.ThreadID != nil {
		modelMap["thread_id"] = flex.IntValue(model.ThreadID)
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigOracleRecoveryValidationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreateDummyInstance != nil {
		modelMap["create_dummy_instance"] = model.CreateDummyInstance
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ShouldRestoreSpfileOrPfile != nil {
		modelMap["should_restore_spfile_or_pfile"] = model.ShouldRestoreSpfileOrPfile
	}
	if model.FileLocation != nil {
		modelMap["file_location"] = model.FileLocation
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleNewTargetDatabaseConfigRedoLogConfigToMap(model *backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigRedoLogConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NumGroups != nil {
		modelMap["num_groups"] = flex.IntValue(model.NumGroups)
	}
	if model.MemberPrefix != nil {
		modelMap["member_prefix"] = model.MemberPrefix
	}
	if model.SizeMBytes != nil {
		modelMap["size_m_bytes"] = flex.IntValue(model.SizeMBytes)
	}
	if model.GroupMembers != nil {
		modelMap["group_members"] = model.GroupMembers
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptionsToMap(model *backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DelaySecs != nil {
		modelMap["delay_secs"] = flex.IntValue(model.DelaySecs)
	}
	if model.TargetPathVec != nil {
		modelMap["target_path_vec"] = model.TargetPathVec
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleAppNewSourceConfigRecoverViewParamsToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverViewParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := resourceIbmRecoveryOracleDbChannelToMap(&dbChannelsItem)
			if err != nil {
				return modelMap, err
			}
			dbChannels = append(dbChannels, dbChannelsItemMap)
		}
		modelMap["db_channels"] = dbChannels
	}
	if model.RecoveryMode != nil {
		modelMap["recovery_mode"] = model.RecoveryMode
	}
	if model.ShellEvironmentVars != nil {
		shellEvironmentVars := []map[string]interface{}{}
		for _, shellEvironmentVarsItem := range model.ShellEvironmentVars {
			shellEvironmentVarsItemMap, err := resourceIbmRecoveryKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["restore_spfile_or_pfile_info"] = []map[string]interface{}{restoreSpfileOrPfileInfoMap}
	}
	if model.UseScnForRestore != nil {
		modelMap["use_scn_for_restore"] = model.UseScnForRestore
	}
	if model.ViewMountPath != nil {
		modelMap["view_mount_path"] = model.ViewMountPath
	}
	return modelMap, nil
}

func resourceIbmRecoveryCommonRecoverOracleAppTargetParamsOriginalSourceConfigToMap(model *backuprecoveryv1.CommonRecoverOracleAppTargetParamsOriginalSourceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := resourceIbmRecoveryOracleDbChannelToMap(&dbChannelsItem)
			if err != nil {
				return modelMap, err
			}
			dbChannels = append(dbChannels, dbChannelsItemMap)
		}
		modelMap["db_channels"] = dbChannels
	}
	if model.RecoveryMode != nil {
		modelMap["recovery_mode"] = model.RecoveryMode
	}
	if model.ShellEvironmentVars != nil {
		shellEvironmentVars := []map[string]interface{}{}
		for _, shellEvironmentVarsItem := range model.ShellEvironmentVars {
			shellEvironmentVarsItemMap, err := resourceIbmRecoveryKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := resourceIbmRecoveryCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["restore_spfile_or_pfile_info"] = []map[string]interface{}{restoreSpfileOrPfileInfoMap}
	}
	if model.UseScnForRestore != nil {
		modelMap["use_scn_for_restore"] = model.UseScnForRestore
	}
	if model.RollForwardLogPathVec != nil {
		modelMap["roll_forward_log_path_vec"] = model.RollForwardLogPathVec
	}
	if model.AttemptCompleteRecovery != nil {
		modelMap["attempt_complete_recovery"] = model.AttemptCompleteRecovery
	}
	if model.RollForwardTimeMsecs != nil {
		modelMap["roll_forward_time_msecs"] = flex.IntValue(model.RollForwardTimeMsecs)
	}
	if model.StopActivePassive != nil {
		modelMap["stop_active_passive"] = model.StopActivePassive
	}
	return modelMap, nil
}

func resourceIbmRecoveryRecoverOracleAppParamsVlanConfigToMap(model *backuprecoveryv1.RecoverOracleAppParamsVlanConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = model.InterfaceName
	}
	return modelMap, nil
}

func resourceIbmRecoveryTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmRecoveryCreationInfoToMap(model *backuprecoveryv1.CreationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UserName != nil {
		modelMap["user_name"] = model.UserName
	}
	return modelMap, nil
}

func resourceIbmRecoveryRetrieveArchiveTaskToMap(model *backuprecoveryv1.RetrieveArchiveTask) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TaskUid != nil {
		modelMap["task_uid"] = model.TaskUid
	}
	if model.UptierExpiryTimes != nil {
		modelMap["uptier_expiry_times"] = model.UptierExpiryTimes
	}
	return modelMap, nil
}
