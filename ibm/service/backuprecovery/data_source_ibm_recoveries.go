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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmRecoveries() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRecoveriesRead,

		Schema: map[string]*schema.Schema{
			"ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter Recoveries for given ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"return_only_child_recoveries": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the organizations for which recoveries are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned.",
			},
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds.",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds.",
			},
			"storage_domain_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter by Storage Domain id. Only recoveries writing data to this Storage Domain will be returned.",
			},
			"snapshot_target_type": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the snapshot's target type from which recovery has been performed.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"archival_target_type": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snapshot_environments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recovery_actions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recoveries": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies list of Recoveries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the Recovery.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Recovery.",
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
						"snapshot_environment": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of snapshot environment for which the Recovery was performed.",
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
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"is_multi_stage_restore": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case.",
						},
						"physical_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the recovery options specific to Physical environment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"objects": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of Recover Object parameters. For recovering files, specifies the object contains the file to recover.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"snapshot_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the snapshot id.",
												},
												"point_in_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.",
												},
												"protection_group_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group id of the object snapshot.",
												},
												"protection_group_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group name of the object snapshot.",
												},
												"snapshot_creation_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.",
												},
												"object_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the information about the object for which the snapshot is taken.",
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
													Computed:    true,
													Description: "Specifies the archival target information if the snapshot is an archival snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"target_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the archival target ID.",
															},
															"archival_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.",
															},
															"target_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival target name.",
															},
															"target_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival target type.",
															},
															"usage_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the usage type for the target.",
															},
															"ownership_context": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the ownership context for the target.",
															},
															"tier_settings": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the tier info for archival.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cloud_platform": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the cloud platform to enable tiering.",
																		},
																		"oracle_tiering": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies Oracle tiers.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"tiers": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
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
																			Computed:    true,
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
													Computed:    true,
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
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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
										Computed:    true,
										Description: "Specifies the type of recover action to be performed.",
									},
									"recover_volume_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters to recover Physical Volumes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
												},
												"physical_target_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the params for recovering to a physical target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mount_target": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the target entity where the volumes are being mounted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
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
																Computed:    true,
																Description: "Specifies the mapping from source volumes to destination volumes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"source_volume_guid": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the guid of the source volume.",
																		},
																		"destination_volume_guid": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the guid of the destination volume.",
																		},
																	},
																},
															},
															"force_unmount_volume": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false.",
															},
															"vlan_config": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
																		},
																		"disable_vlan": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
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
										Computed:    true,
										Description: "Specifies the parameters to mount Physical Volumes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
												},
												"physical_target_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the params for recovering to a physical target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mount_to_original_target": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.",
															},
															"original_target_config": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the configuration for mounting to the original target.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"server_credentials": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies credentials to access the target server. This is required if the server is of Linux OS.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"username": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the username to access target entity.",
																					},
																					"password": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
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
																Computed:    true,
																Description: "Specifies the configuration for mounting to a new target.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mount_target": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the target entity to recover to.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
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
																			Computed:    true,
																			Description: "Specifies credentials to access the target server. This is required if the server is of Linux OS.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"username": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the username to access target entity.",
																					},
																					"password": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
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
																Computed:    true,
																Description: "Specifies whether to perform a read-only mount. Default is false.",
															},
															"volume_names": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
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
																Computed:    true,
																Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
																		},
																		"disable_vlan": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
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
										Computed:    true,
										Description: "Specifies the parameters to perform a file and folder recovery.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"files_and_folders": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the information about the files and folders to be recovered.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"absolute_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the absolute path to the file or folder.",
															},
															"destination_dir": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the destination directory where the file/directory was copied.",
															},
															"is_directory": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
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
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_view_file_recovery": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specify if the recovery is of type view file/folder.",
															},
														},
													},
												},
												"target_environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
												},
												"physical_target_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the parameters to recover to a Physical target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recover_target": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the target entity where the volumes are being mounted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
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
																Computed:    true,
																Description: "If this is true, then files will be restored to original paths.",
															},
															"overwrite_existing": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to overwrite existing file/folder during recovery.",
															},
															"alternate_restore_directory": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the directory path where restore should happen if restore_to_original_paths is set to false.",
															},
															"preserve_attributes": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to preserve file/folder attributes during recovery.",
															},
															"preserve_timestamps": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether to preserve the original time stamps.",
															},
															"preserve_acls": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether to preserve the ACLs of the original file.",
															},
															"continue_on_error": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.",
															},
															"save_success_files": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to save success files or not. Default value is false.",
															},
															"vlan_config": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
																		},
																		"disable_vlan": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
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
																Computed:    true,
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
										Computed:    true,
										Description: "Specifies the parameters to download files and folders.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"files_and_folders": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the info about the files and folders to be recovered.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"absolute_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the absolute path to the file or folder.",
															},
															"destination_dir": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the destination directory where the file/directory was copied.",
															},
															"is_directory": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
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
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_view_file_recovery": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specify if the recovery is of type view file/folder.",
															},
														},
													},
												},
												"download_file_path": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the path location to download the files and folders.",
												},
											},
										},
									},
									"system_recovery_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters to perform a system recovery.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"full_nas_path": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "Specifies the recovery options specific to oracle environment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"objects": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of parameters for list of objects to be recovered.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"snapshot_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the snapshot id.",
												},
												"point_in_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.",
												},
												"protection_group_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group id of the object snapshot.",
												},
												"protection_group_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group name of the object snapshot.",
												},
												"snapshot_creation_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.",
												},
												"object_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the information about the object for which the snapshot is taken.",
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
													Computed:    true,
													Description: "Specifies the archival target information if the snapshot is an archival snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"target_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the archival target ID.",
															},
															"archival_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.",
															},
															"target_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival target name.",
															},
															"target_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the archival target type.",
															},
															"usage_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the usage type for the target.",
															},
															"ownership_context": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the ownership context for the target.",
															},
															"tier_settings": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the tier info for archival.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cloud_platform": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the cloud platform to enable tiering.",
																		},
																		"oracle_tiering": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies Oracle tiers.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"tiers": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
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
																			Computed:    true,
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
													Computed:    true,
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
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"bytes_restored": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specify the total bytes restored.",
												},
												"instant_recovery_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
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
										Computed:    true,
										Description: "Specifies the type of recover action to be performed.",
									},
									"recover_app_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters to recover Oracle databases.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment of the recovery target. The corresponding params below must be filled out.",
												},
												"oracle_target_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the params for recovering to a oracle host. Provided oracle backup should be recovered to same type of target host. For Example: If you have oracle backup taken from a physical host then that should be recovered to physical host only.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"recover_to_new_source": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the parameter whether the recovery should be performed to a new source or an original Source Target.",
															},
															"new_source_config": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"host": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
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
																			Computed:    true,
																			Description: "Specifies if recovery target is a database or a view.",
																		},
																		"recover_database_params": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies recovery parameters when recovering to a database.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"restore_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
																					},
																					"db_channels": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"archive_log_retention_days": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																								},
																								"archive_log_retention_hours": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																								},
																								"credentials": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the object to hold username and password.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"username": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the username to access target entity.",
																											},
																											"password": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the password to access target entity.",
																											},
																										},
																									},
																								},
																								"database_unique_name": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the unique Name of the database.",
																								},
																								"database_uuid": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																								},
																								"default_channel_count": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																								},
																								"database_node_list": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"host_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the id of the database host from which backup is allowed.",
																											},
																											"channel_count": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																											},
																											"port": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies the port where the Database is listening.",
																											},
																											"sbt_host_params": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies details about capturing Oracle SBT host info.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"sbt_library_path": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies the path of sbt library.",
																														},
																														"view_fs_path": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies the Cohesity view path.",
																														},
																														"vip_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the list of Cohesity primary VIPs.",
																															Elem: &schema.Schema{
																																Type: schema.TypeString,
																															},
																														},
																														"vlan_info_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the Vlan information for Cohesity cluster.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"ip_list": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the list of Ips in this VLAN.",
																																		Elem: &schema.Schema{
																																			Type: schema.TypeString,
																																		},
																																	},
																																	"gateway": &schema.Schema{
																																		Type:        schema.TypeString,
																																		Computed:    true,
																																		Description: "Specifies the gateway of this VLAN.",
																																	},
																																	"id": &schema.Schema{
																																		Type:        schema.TypeInt,
																																		Computed:    true,
																																		Description: "Specifies the Id of this VLAN.",
																																	},
																																	"subnet_ip": &schema.Schema{
																																		Type:        schema.TypeString,
																																		Computed:    true,
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
																									Computed:    true,
																									Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																								},
																								"enable_dg_primary_backup": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																								},
																								"rman_backup_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of Oracle RMAN backup requested.",
																								},
																							},
																						},
																					},
																					"recovery_mode": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if database should be left in recovery mode.",
																					},
																					"shell_evironment_vars": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"key": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "key.",
																								},
																								"value": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "value.",
																								},
																							},
																						},
																					},
																					"granular_restore_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies information about list of objects (PDBs) to restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"granularity_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies type of granular restore.",
																								},
																								"pdb_restore_params": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies information about the list of pdbs to be restored.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"drop_duplicate_pdb": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																											},
																											"pdb_objects": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies list of PDB objects to restore.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"db_id": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies pluggable database id.",
																														},
																														"db_name": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies name of the DB.",
																														},
																													},
																												},
																											},
																											"restore_to_existing_cdb": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if pdbs should be restored to an existing CDB.",
																											},
																											"rename_pdb_map": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the new PDB name mapping to existing PDBs.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"key": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "key.",
																														},
																														"value": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "value.",
																														},
																													},
																												},
																											},
																											"include_in_restore": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
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
																						Computed:    true,
																						Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"range_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of range.",
																								},
																								"range_info_vec": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies an array of oracle restore ranges.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"start_of_range": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																											},
																											"end_of_range": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																											},
																											"protection_group_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																											},
																											"reset_log_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																											},
																											"incarnation_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																											},
																											"thread_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																											},
																										},
																									},
																								},
																								"archive_log_restore_dest": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies destination where archive logs are to be restored.",
																								},
																							},
																						},
																					},
																					"oracle_recovery_validation_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies parameters related to Oracle Recovery Validation.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"create_dummy_instance": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																								},
																							},
																						},
																					},
																					"restore_spfile_or_pfile_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies parameters related to spfile/pfile restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"should_restore_spfile_or_pfile": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether to restore spfile/pfile or skip it.",
																								},
																								"file_location": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																								},
																							},
																						},
																					},
																					"use_scn_for_restore": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether database recovery performed should use scn value or not.",
																					},
																					"database_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.",
																					},
																					"oracle_base_folder": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the oracle base folder at selected host.",
																					},
																					"oracle_home_folder": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the oracle home folder at selected host.",
																					},
																					"db_files_destination": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the location to restore database files.",
																					},
																					"db_config_file_path": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the config file path on selected host which configures the restored database.",
																					},
																					"enable_archive_log_mode": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies archive log mode for oracle restore.",
																					},
																					"pfile_parameter_map": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies a key value pair for pfile parameters.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"key": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "key.",
																								},
																								"value": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "value.",
																								},
																							},
																						},
																					},
																					"bct_file_path": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies BCT file path.",
																					},
																					"num_tempfiles": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies no. of tempfiles to be used for the recovered database.",
																					},
																					"redo_log_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies redo log config.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"num_groups": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies no. of redo log groups.",
																								},
																								"member_prefix": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies Log member name prefix.",
																								},
																								"size_m_bytes": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies Size of the member in MB.",
																								},
																								"group_members": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies list of members of this redo log group.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"is_multi_stage_restore": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether this task is a multistage restore task. If set, we migrate the DB after clone completes.",
																					},
																					"oracle_update_restore_options": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the parameters that are needed for updating oracle restore options.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"delay_secs": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies when the migration of the oracle instance should be started after successful recovery.",
																								},
																								"target_path_vec": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the target paths to be used for DB migration.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"skip_clone_nid": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Whether or not to skip the nid step in Oracle Clone workflow. Applicable to both smart and old clone workflow.",
																					},
																					"no_filename_check": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether to validate filenames or not in Oracle alternate restore workflow.",
																					},
																					"new_name_clause": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies newname clause for db files which allows user to have full control on how their database files can be renamed during the oracle alternate restore workflow.",
																					},
																				},
																			},
																		},
																		"recover_view_params": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies recovery parameters when recovering to a view.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"restore_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
																					},
																					"db_channels": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"archive_log_retention_days": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																								},
																								"archive_log_retention_hours": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																								},
																								"credentials": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the object to hold username and password.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"username": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the username to access target entity.",
																											},
																											"password": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the password to access target entity.",
																											},
																										},
																									},
																								},
																								"database_unique_name": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the unique Name of the database.",
																								},
																								"database_uuid": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																								},
																								"default_channel_count": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																								},
																								"database_node_list": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"host_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the id of the database host from which backup is allowed.",
																											},
																											"channel_count": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																											},
																											"port": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies the port where the Database is listening.",
																											},
																											"sbt_host_params": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies details about capturing Oracle SBT host info.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"sbt_library_path": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies the path of sbt library.",
																														},
																														"view_fs_path": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies the Cohesity view path.",
																														},
																														"vip_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the list of Cohesity primary VIPs.",
																															Elem: &schema.Schema{
																																Type: schema.TypeString,
																															},
																														},
																														"vlan_info_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the Vlan information for Cohesity cluster.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"ip_list": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the list of Ips in this VLAN.",
																																		Elem: &schema.Schema{
																																			Type: schema.TypeString,
																																		},
																																	},
																																	"gateway": &schema.Schema{
																																		Type:        schema.TypeString,
																																		Computed:    true,
																																		Description: "Specifies the gateway of this VLAN.",
																																	},
																																	"id": &schema.Schema{
																																		Type:        schema.TypeInt,
																																		Computed:    true,
																																		Description: "Specifies the Id of this VLAN.",
																																	},
																																	"subnet_ip": &schema.Schema{
																																		Type:        schema.TypeString,
																																		Computed:    true,
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
																									Computed:    true,
																									Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																								},
																								"enable_dg_primary_backup": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																								},
																								"rman_backup_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of Oracle RMAN backup requested.",
																								},
																							},
																						},
																					},
																					"recovery_mode": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if database should be left in recovery mode.",
																					},
																					"shell_evironment_vars": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"key": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "key.",
																								},
																								"value": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "value.",
																								},
																							},
																						},
																					},
																					"granular_restore_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies information about list of objects (PDBs) to restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"granularity_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies type of granular restore.",
																								},
																								"pdb_restore_params": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies information about the list of pdbs to be restored.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"drop_duplicate_pdb": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																											},
																											"pdb_objects": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies list of PDB objects to restore.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"db_id": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies pluggable database id.",
																														},
																														"db_name": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies name of the DB.",
																														},
																													},
																												},
																											},
																											"restore_to_existing_cdb": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if pdbs should be restored to an existing CDB.",
																											},
																											"rename_pdb_map": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the new PDB name mapping to existing PDBs.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"key": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "key.",
																														},
																														"value": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "value.",
																														},
																													},
																												},
																											},
																											"include_in_restore": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
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
																						Computed:    true,
																						Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"range_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of range.",
																								},
																								"range_info_vec": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies an array of oracle restore ranges.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"start_of_range": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																											},
																											"end_of_range": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																											},
																											"protection_group_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																											},
																											"reset_log_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																											},
																											"incarnation_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																											},
																											"thread_id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																											},
																										},
																									},
																								},
																								"archive_log_restore_dest": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies destination where archive logs are to be restored.",
																								},
																							},
																						},
																					},
																					"oracle_recovery_validation_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies parameters related to Oracle Recovery Validation.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"create_dummy_instance": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																								},
																							},
																						},
																					},
																					"restore_spfile_or_pfile_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies parameters related to spfile/pfile restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"should_restore_spfile_or_pfile": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether to restore spfile/pfile or skip it.",
																								},
																								"file_location": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																								},
																							},
																						},
																					},
																					"use_scn_for_restore": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether database recovery performed should use scn value or not.",
																					},
																					"view_mount_path": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
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
																Computed:    true,
																Description: "Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"restore_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.",
																		},
																		"db_channels": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"archive_log_retention_days": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n.",
																					},
																					"archive_log_retention_hours": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k.",
																					},
																					"credentials": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the object to hold username and password.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"username": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the username to access target entity.",
																								},
																								"password": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the password to access target entity.",
																								},
																							},
																						},
																					},
																					"database_unique_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the unique Name of the database.",
																					},
																					"database_uuid": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id.",
																					},
																					"default_channel_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item's channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same.",
																					},
																					"database_node_list": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the Node info from where we are allowed to take the backup/restore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"host_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the id of the database host from which backup is allowed.",
																								},
																								"channel_count": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host.",
																								},
																								"port": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the port where the Database is listening.",
																								},
																								"sbt_host_params": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies details about capturing Oracle SBT host info.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"sbt_library_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the path of sbt library.",
																											},
																											"view_fs_path": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the Cohesity view path.",
																											},
																											"vip_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the list of Cohesity primary VIPs.",
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"vlan_info_list": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the Vlan information for Cohesity cluster.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"ip_list": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the list of Ips in this VLAN.",
																															Elem: &schema.Schema{
																																Type: schema.TypeString,
																															},
																														},
																														"gateway": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Specifies the gateway of this VLAN.",
																														},
																														"id": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the Id of this VLAN.",
																														},
																														"subnet_ip": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
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
																						Computed:    true,
																						Description: "Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList.",
																					},
																					"enable_dg_primary_backup": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether the database having the Primary role within Data Guard configuration is to be backed up.",
																					},
																					"rman_backup_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of Oracle RMAN backup requested.",
																					},
																				},
																			},
																		},
																		"recovery_mode": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if database should be left in recovery mode.",
																		},
																		"shell_evironment_vars": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies key value pairs of shell variables which defines the restore shell environment.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "key.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "value.",
																					},
																				},
																			},
																		},
																		"granular_restore_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies information about list of objects (PDBs) to restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"granularity_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies type of granular restore.",
																					},
																					"pdb_restore_params": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies information about the list of pdbs to be restored.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"drop_duplicate_pdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies if the PDB should be ignored if a PDB already exists with same name.",
																								},
																								"pdb_objects": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies list of PDB objects to restore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"db_id": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies pluggable database id.",
																											},
																											"db_name": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies name of the DB.",
																											},
																										},
																									},
																								},
																								"restore_to_existing_cdb": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies if pdbs should be restored to an existing CDB.",
																								},
																								"rename_pdb_map": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the new PDB name mapping to existing PDBs.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"key": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "key.",
																											},
																											"value": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "value.",
																											},
																										},
																									},
																								},
																								"include_in_restore": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
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
																			Computed:    true,
																			Description: "Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"range_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of range.",
																					},
																					"range_info_vec": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies an array of oracle restore ranges.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"start_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies starting value of the range in time (usecs), SCN or sequence no.",
																								},
																								"end_of_range": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies ending value of the range in time (usecs), SCN or sequence no.",
																								},
																								"protection_group_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies id of the Protection Group corresponding to this oracle range.",
																								},
																								"reset_log_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"incarnation_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no.",
																								},
																								"thread_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies thread id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type sequence no.",
																								},
																							},
																						},
																					},
																					"archive_log_restore_dest": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies destination where archive logs are to be restored.",
																					},
																				},
																			},
																		},
																		"oracle_recovery_validation_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies parameters related to Oracle Recovery Validation.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"create_dummy_instance": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.",
																					},
																				},
																			},
																		},
																		"restore_spfile_or_pfile_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies parameters related to spfile/pfile restore.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"should_restore_spfile_or_pfile": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether to restore spfile/pfile or skip it.",
																					},
																					"file_location": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs.",
																					},
																				},
																			},
																		},
																		"use_scn_for_restore": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether database recovery performed should use scn value or not.",
																		},
																		"roll_forward_log_path_vec": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "List of archive logs to apply on Database after overwrite restore.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"attempt_complete_recovery": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Whether or not this is a complete recovery attempt.",
																		},
																		"roll_forward_time_msecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "UTC time in msecs till which we have to roll-forward the database.",
																		},
																		"stop_active_passive": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
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
													Computed:    true,
													Description: "Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.",
															},
															"disable_vlan": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
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
					},
				},
			},
		},
	}
}

func dataSourceIbmRecoveriesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRecoveriesOptions := &backuprecoveryv1.GetRecoveriesOptions{}

	if _, ok := d.GetOk("ids"); ok {
		var ids []string
		for _, v := range d.Get("ids").([]interface{}) {
			idsItem := v.(string)
			ids = append(ids, idsItem)
		}
		getRecoveriesOptions.SetIds(ids)
	}
	if _, ok := d.GetOk("return_only_child_recoveries"); ok {
		getRecoveriesOptions.SetReturnOnlyChildRecoveries(d.Get("return_only_child_recoveries").(bool))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getRecoveriesOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getRecoveriesOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("start_time_usecs"); ok {
		getRecoveriesOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getRecoveriesOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("storage_domain_id"); ok {
		getRecoveriesOptions.SetStorageDomainID(int64(d.Get("storage_domain_id").(int)))
	}
	if _, ok := d.GetOk("snapshot_target_type"); ok {
		var snapshotTargetType []string
		for _, v := range d.Get("snapshot_target_type").([]interface{}) {
			snapshotTargetTypeItem := v.(string)
			snapshotTargetType = append(snapshotTargetType, snapshotTargetTypeItem)
		}
		getRecoveriesOptions.SetSnapshotTargetType(snapshotTargetType)
	}
	if _, ok := d.GetOk("archival_target_type"); ok {
		var archivalTargetType []string
		for _, v := range d.Get("archival_target_type").([]interface{}) {
			archivalTargetTypeItem := v.(string)
			archivalTargetType = append(archivalTargetType, archivalTargetTypeItem)
		}
		getRecoveriesOptions.SetArchivalTargetType(archivalTargetType)
	}
	if _, ok := d.GetOk("snapshot_environments"); ok {
		var snapshotEnvironments []string
		for _, v := range d.Get("snapshot_environments").([]interface{}) {
			snapshotEnvironmentsItem := v.(string)
			snapshotEnvironments = append(snapshotEnvironments, snapshotEnvironmentsItem)
		}
		getRecoveriesOptions.SetSnapshotEnvironments(snapshotEnvironments)
	}
	if _, ok := d.GetOk("status"); ok {
		var status []string
		for _, v := range d.Get("status").([]interface{}) {
			statusItem := v.(string)
			status = append(status, statusItem)
		}
		getRecoveriesOptions.SetStatus(status)
	}
	if _, ok := d.GetOk("recovery_actions"); ok {
		var recoveryActions []string
		for _, v := range d.Get("recovery_actions").([]interface{}) {
			recoveryActionsItem := v.(string)
			recoveryActions = append(recoveryActions, recoveryActionsItem)
		}
		getRecoveriesOptions.SetRecoveryActions(recoveryActions)
	}

	recoveriesResponse, response, err := backupRecoveryClient.GetRecoveriesWithContext(context, getRecoveriesOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRecoveriesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRecoveriesWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRecoveriesID(d))

	recoveries := []map[string]interface{}{}
	if recoveriesResponse.Recoveries != nil {
		for _, modelItem := range recoveriesResponse.Recoveries {
			modelMap, err := dataSourceIbmRecoveriesRecoveryToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			recoveries = append(recoveries, modelMap)
		}
	}
	if err = d.Set("recoveries", recoveries); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting recoveries %s", err))
	}

	return nil
}

// dataSourceIbmRecoveriesID returns a reasonable ID for the list.
func dataSourceIbmRecoveriesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmRecoveriesRecoveryToMap(model *backuprecoveryv1.Recovery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.SnapshotEnvironment != nil {
		modelMap["snapshot_environment"] = model.SnapshotEnvironment
	}
	if model.RecoveryAction != nil {
		modelMap["recovery_action"] = model.RecoveryAction
	}
	if model.Permissions != nil {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range model.Permissions {
			permissionsItemMap, err := dataSourceIbmRecoveriesTenantToMap(&permissionsItem)
			if err != nil {
				return modelMap, err
			}
			permissions = append(permissions, permissionsItemMap)
		}
		modelMap["permissions"] = permissions
	}
	if model.CreationInfo != nil {
		creationInfoMap, err := dataSourceIbmRecoveriesCreationInfoToMap(model.CreationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["creation_info"] = []map[string]interface{}{creationInfoMap}
	}
	if model.CanTearDown != nil {
		modelMap["can_tear_down"] = model.CanTearDown
	}
	if model.TearDownStatus != nil {
		modelMap["tear_down_status"] = model.TearDownStatus
	}
	if model.TearDownMessage != nil {
		modelMap["tear_down_message"] = model.TearDownMessage
	}
	if model.Messages != nil {
		modelMap["messages"] = model.Messages
	}
	if model.IsParentRecovery != nil {
		modelMap["is_parent_recovery"] = model.IsParentRecovery
	}
	if model.ParentRecoveryID != nil {
		modelMap["parent_recovery_id"] = model.ParentRecoveryID
	}
	if model.RetrieveArchiveTasks != nil {
		retrieveArchiveTasks := []map[string]interface{}{}
		for _, retrieveArchiveTasksItem := range model.RetrieveArchiveTasks {
			retrieveArchiveTasksItemMap, err := dataSourceIbmRecoveriesRetrieveArchiveTaskToMap(&retrieveArchiveTasksItem)
			if err != nil {
				return modelMap, err
			}
			retrieveArchiveTasks = append(retrieveArchiveTasks, retrieveArchiveTasksItemMap)
		}
		modelMap["retrieve_archive_tasks"] = retrieveArchiveTasks
	}
	if model.IsMultiStageRestore != nil {
		modelMap["is_multi_stage_restore"] = model.IsMultiStageRestore
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmRecoveriesRecoverOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesCreationInfoToMap(model *backuprecoveryv1.CreationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UserName != nil {
		modelMap["user_name"] = model.UserName
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRetrieveArchiveTaskToMap(model *backuprecoveryv1.RetrieveArchiveTask) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TaskUid != nil {
		modelMap["task_uid"] = model.TaskUid
	}
	if model.UptierExpiryTimes != nil {
		modelMap["uptier_expiry_times"] = model.UptierExpiryTimes
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverPhysicalParamsToMap(model *backuprecoveryv1.RecoverPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	modelMap["recovery_action"] = model.RecoveryAction
	if model.RecoverVolumeParams != nil {
		recoverVolumeParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsRecoverVolumeParamsToMap(model.RecoverVolumeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_volume_params"] = []map[string]interface{}{recoverVolumeParamsMap}
	}
	if model.MountVolumeParams != nil {
		mountVolumeParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsMountVolumeParamsToMap(model.MountVolumeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["mount_volume_params"] = []map[string]interface{}{mountVolumeParamsMap}
	}
	if model.RecoverFileAndFolderParams != nil {
		recoverFileAndFolderParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(model.RecoverFileAndFolderParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_file_and_folder_params"] = []map[string]interface{}{recoverFileAndFolderParamsMap}
	}
	if model.DownloadFileAndFolderParams != nil {
		downloadFileAndFolderParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(model.DownloadFileAndFolderParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["download_file_and_folder_params"] = []map[string]interface{}{downloadFileAndFolderParamsMap}
	}
	if model.SystemRecoveryParams != nil {
		systemRecoveryParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalParamsSystemRecoveryParamsToMap(model.SystemRecoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_recovery_params"] = []map[string]interface{}{systemRecoveryParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParams) (map[string]interface{}, error) {
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
		objectInfoMap, err := dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
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
		archivalTargetInfoMap, err := dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
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

func dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsObjectInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model *backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmRecoveriesArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmRecoveriesOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmRecoveriesOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmRecoveriesOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmRecoveriesOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverPhysicalParamsRecoverVolumeParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	mountTargetMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverVolumeMountTargetToMap(model.MountTarget)
	if err != nil {
		return modelMap, err
	}
	modelMap["mount_target"] = []map[string]interface{}{mountTargetMap}
	volumeMapping := []map[string]interface{}{}
	for _, volumeMappingItem := range model.VolumeMapping {
		volumeMappingItemMap, err := dataSourceIbmRecoveriesRecoverVolumeMappingToMap(&volumeMappingItem)
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
		vlanConfigMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverVolumeMountTargetToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverVolumeMappingToMap(model *backuprecoveryv1.RecoverVolumeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_volume_guid"] = model.SourceVolumeGuid
	modelMap["destination_volume_guid"] = model.DestinationVolumeGuid
	return modelMap, nil
}

func dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverPhysicalParamsMountVolumeParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := dataSourceIbmRecoveriesMountPhysicalVolumeParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesMountPhysicalVolumeParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mount_to_original_target"] = model.MountToOriginalTarget
	if model.OriginalTargetConfig != nil {
		originalTargetConfigMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(model.OriginalTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_target_config"] = []map[string]interface{}{originalTargetConfigMap}
	}
	if model.NewTargetConfig != nil {
		newTargetConfigMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(model.NewTargetConfig)
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
			mountedVolumeMappingItemMap, err := dataSourceIbmRecoveriesMountedVolumeMappingToMap(&mountedVolumeMappingItem)
			if err != nil {
				return modelMap, err
			}
			mountedVolumeMapping = append(mountedVolumeMapping, mountedVolumeMappingItemMap)
		}
		modelMap["mounted_volume_mapping"] = mountedVolumeMapping
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ServerCredentials != nil {
		serverCredentialsMap, err := dataSourceIbmRecoveriesPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(model.ServerCredentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["server_credentials"] = []map[string]interface{}{serverCredentialsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(model *backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	mountTargetMap, err := dataSourceIbmRecoveriesRecoverTargetToMap(model.MountTarget)
	if err != nil {
		return modelMap, err
	}
	modelMap["mount_target"] = []map[string]interface{}{mountTargetMap}
	if model.ServerCredentials != nil {
		serverCredentialsMap, err := dataSourceIbmRecoveriesPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(model.ServerCredentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["server_credentials"] = []map[string]interface{}{serverCredentialsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverTargetToMap(model *backuprecoveryv1.RecoverTarget) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(model *backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func dataSourceIbmRecoveriesMountedVolumeMappingToMap(model *backuprecoveryv1.MountedVolumeMapping) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesPhysicalTargetParamsForMountVolumeVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	filesAndFolders := []map[string]interface{}{}
	for _, filesAndFoldersItem := range model.FilesAndFolders {
		filesAndFoldersItemMap, err := dataSourceIbmRecoveriesCommonRecoverFileAndFolderInfoToMap(&filesAndFoldersItem)
		if err != nil {
			return modelMap, err
		}
		filesAndFolders = append(filesAndFolders, filesAndFoldersItemMap)
	}
	modelMap["files_and_folders"] = filesAndFolders
	modelMap["target_environment"] = model.TargetEnvironment
	if model.PhysicalTargetParams != nil {
		physicalTargetParamsMap, err := dataSourceIbmRecoveriesRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(model.PhysicalTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_target_params"] = []map[string]interface{}{physicalTargetParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonRecoverFileAndFolderInfoToMap(model *backuprecoveryv1.CommonRecoverFileAndFolderInfo) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(model *backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	recoverTargetMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(model.RecoverTarget)
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
		vlanConfigMap, err := dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(model.VlanConfig)
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

func dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(model *backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FilesAndFolders != nil {
		filesAndFolders := []map[string]interface{}{}
		for _, filesAndFoldersItem := range model.FilesAndFolders {
			filesAndFoldersItemMap, err := dataSourceIbmRecoveriesCommonRecoverFileAndFolderInfoToMap(&filesAndFoldersItem)
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

func dataSourceIbmRecoveriesRecoverPhysicalParamsSystemRecoveryParamsToMap(model *backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FullNasPath != nil {
		modelMap["full_nas_path"] = model.FullNasPath
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleParamsToMap(model *backuprecoveryv1.RecoverOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := dataSourceIbmRecoveriesRecoverOracleDbSnapshotParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	modelMap["recovery_action"] = model.RecoveryAction
	if model.RecoverAppParams != nil {
		recoverAppParamsMap, err := dataSourceIbmRecoveriesRecoverOracleParamsRecoverAppParamsToMap(model.RecoverAppParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_app_params"] = []map[string]interface{}{recoverAppParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleDbSnapshotParamsToMap(model *backuprecoveryv1.RecoverOracleDbSnapshotParams) (map[string]interface{}, error) {
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
		objectInfoMap, err := dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsObjectInfoToMap(model.ObjectInfo)
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
		archivalTargetInfoMap, err := dataSourceIbmRecoveriesCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model.ArchivalTargetInfo)
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
		instantRecoveryInfoMap, err := dataSourceIbmRecoveriesRecoverOracleDbSnapshotParamsInstantRecoveryInfoToMap(model.InstantRecoveryInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["instant_recovery_info"] = []map[string]interface{}{instantRecoveryInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleDbSnapshotParamsInstantRecoveryInfoToMap(model *backuprecoveryv1.RecoverOracleDbSnapshotParamsInstantRecoveryInfo) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverOracleParamsRecoverAppParamsToMap(model *backuprecoveryv1.RecoverOracleParamsRecoverAppParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_environment"] = model.TargetEnvironment
	if model.OracleTargetParams != nil {
		oracleTargetParamsMap, err := dataSourceIbmRecoveriesRecoverOracleAppParamsOracleTargetParamsToMap(model.OracleTargetParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_target_params"] = []map[string]interface{}{oracleTargetParamsMap}
	}
	if model.VlanConfig != nil {
		vlanConfigMap, err := dataSourceIbmRecoveriesRecoverOracleAppParamsVlanConfigToMap(model.VlanConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_config"] = []map[string]interface{}{vlanConfigMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleAppParamsOracleTargetParamsToMap(model *backuprecoveryv1.RecoverOracleAppParamsOracleTargetParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["recover_to_new_source"] = model.RecoverToNewSource
	if model.NewSourceConfig != nil {
		newSourceConfigMap, err := dataSourceIbmRecoveriesCommonRecoverOracleAppTargetParamsNewSourceConfigToMap(model.NewSourceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["new_source_config"] = []map[string]interface{}{newSourceConfigMap}
	}
	if model.OriginalSourceConfig != nil {
		originalSourceConfigMap, err := dataSourceIbmRecoveriesCommonRecoverOracleAppTargetParamsOriginalSourceConfigToMap(model.OriginalSourceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_source_config"] = []map[string]interface{}{originalSourceConfigMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonRecoverOracleAppTargetParamsNewSourceConfigToMap(model *backuprecoveryv1.CommonRecoverOracleAppTargetParamsNewSourceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	hostMap, err := dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigHostToMap(model.Host)
	if err != nil {
		return modelMap, err
	}
	modelMap["host"] = []map[string]interface{}{hostMap}
	if model.RecoveryTarget != nil {
		modelMap["recovery_target"] = model.RecoveryTarget
	}
	if model.RecoverDatabaseParams != nil {
		recoverDatabaseParamsMap, err := dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigRecoverDatabaseParamsToMap(model.RecoverDatabaseParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_database_params"] = []map[string]interface{}{recoverDatabaseParamsMap}
	}
	if model.RecoverViewParams != nil {
		recoverViewParamsMap, err := dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigRecoverViewParamsToMap(model.RecoverViewParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["recover_view_params"] = []map[string]interface{}{recoverViewParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigHostToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigHost) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigRecoverDatabaseParamsToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverDatabaseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := dataSourceIbmRecoveriesOracleDbChannelToMap(&dbChannelsItem)
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
			shellEvironmentVarsItemMap, err := dataSourceIbmRecoveriesKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
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
			pfileParameterMapItemMap, err := dataSourceIbmRecoveriesKeyValuePairToMap(&pfileParameterMapItem)
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
		redoLogConfigMap, err := dataSourceIbmRecoveriesRecoverOracleNewTargetDatabaseConfigRedoLogConfigToMap(model.RedoLogConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["redo_log_config"] = []map[string]interface{}{redoLogConfigMap}
	}
	if model.IsMultiStageRestore != nil {
		modelMap["is_multi_stage_restore"] = model.IsMultiStageRestore
	}
	if model.OracleUpdateRestoreOptions != nil {
		oracleUpdateRestoreOptionsMap, err := dataSourceIbmRecoveriesRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptionsToMap(model.OracleUpdateRestoreOptions)
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

func dataSourceIbmRecoveriesOracleDbChannelToMap(model *backuprecoveryv1.OracleDbChannel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchiveLogRetentionDays != nil {
		modelMap["archive_log_retention_days"] = flex.IntValue(model.ArchiveLogRetentionDays)
	}
	if model.ArchiveLogRetentionHours != nil {
		modelMap["archive_log_retention_hours"] = flex.IntValue(model.ArchiveLogRetentionHours)
	}
	if model.Credentials != nil {
		credentialsMap, err := dataSourceIbmRecoveriesCredentialsToMap(model.Credentials)
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
			databaseNodeListItemMap, err := dataSourceIbmRecoveriesOracleDatabaseHostToMap(&databaseNodeListItem)
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

func dataSourceIbmRecoveriesCredentialsToMap(model *backuprecoveryv1.Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func dataSourceIbmRecoveriesOracleDatabaseHostToMap(model *backuprecoveryv1.OracleDatabaseHost) (map[string]interface{}, error) {
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
		sbtHostParamsMap, err := dataSourceIbmRecoveriesOracleSbtHostParamsToMap(model.SbtHostParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sbt_host_params"] = []map[string]interface{}{sbtHostParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesOracleSbtHostParamsToMap(model *backuprecoveryv1.OracleSbtHostParams) (map[string]interface{}, error) {
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
			vlanInfoListItemMap, err := dataSourceIbmRecoveriesOracleVlanInfoToMap(&vlanInfoListItem)
			if err != nil {
				return modelMap, err
			}
			vlanInfoList = append(vlanInfoList, vlanInfoListItemMap)
		}
		modelMap["vlan_info_list"] = vlanInfoList
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesOracleVlanInfoToMap(model *backuprecoveryv1.OracleVlanInfo) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonOracleAppSourceConfigGranularRestoreInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigGranularRestoreInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GranularityType != nil {
		modelMap["granularity_type"] = model.GranularityType
	}
	if model.PdbRestoreParams != nil {
		pdbRestoreParamsMap, err := dataSourceIbmRecoveriesRecoverOracleGranularRestoreInfoPdbRestoreParamsToMap(model.PdbRestoreParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["pdb_restore_params"] = []map[string]interface{}{pdbRestoreParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleGranularRestoreInfoPdbRestoreParamsToMap(model *backuprecoveryv1.RecoverOracleGranularRestoreInfoPdbRestoreParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DropDuplicatePDB != nil {
		modelMap["drop_duplicate_pdb"] = model.DropDuplicatePDB
	}
	if model.PdbObjects != nil {
		pdbObjects := []map[string]interface{}{}
		for _, pdbObjectsItem := range model.PdbObjects {
			pdbObjectsItemMap, err := dataSourceIbmRecoveriesOraclePdbObjectInfoToMap(&pdbObjectsItem)
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
			renamePdbMapItemMap, err := dataSourceIbmRecoveriesKeyValuePairToMap(&renamePdbMapItem)
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

func dataSourceIbmRecoveriesOraclePdbObjectInfoToMap(model *backuprecoveryv1.OraclePdbObjectInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["db_id"] = model.DbID
	modelMap["db_name"] = model.DbName
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigOracleArchiveLogInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RangeType != nil {
		modelMap["range_type"] = model.RangeType
	}
	if model.RangeInfoVec != nil {
		rangeInfoVec := []map[string]interface{}{}
		for _, rangeInfoVecItem := range model.RangeInfoVec {
			rangeInfoVecItemMap, err := dataSourceIbmRecoveriesOracleRangeMetaInfoToMap(&rangeInfoVecItem)
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

func dataSourceIbmRecoveriesOracleRangeMetaInfoToMap(model *backuprecoveryv1.OracleRangeMetaInfo) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigOracleRecoveryValidationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreateDummyInstance != nil {
		modelMap["create_dummy_instance"] = model.CreateDummyInstance
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model *backuprecoveryv1.CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ShouldRestoreSpfileOrPfile != nil {
		modelMap["should_restore_spfile_or_pfile"] = model.ShouldRestoreSpfileOrPfile
	}
	if model.FileLocation != nil {
		modelMap["file_location"] = model.FileLocation
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleNewTargetDatabaseConfigRedoLogConfigToMap(model *backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigRedoLogConfig) (map[string]interface{}, error) {
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

func dataSourceIbmRecoveriesRecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptionsToMap(model *backuprecoveryv1.RecoverOracleNewTargetDatabaseConfigOracleUpdateRestoreOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DelaySecs != nil {
		modelMap["delay_secs"] = flex.IntValue(model.DelaySecs)
	}
	if model.TargetPathVec != nil {
		modelMap["target_path_vec"] = model.TargetPathVec
	}
	return modelMap, nil
}

func dataSourceIbmRecoveriesRecoverOracleAppNewSourceConfigRecoverViewParamsToMap(model *backuprecoveryv1.RecoverOracleAppNewSourceConfigRecoverViewParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := dataSourceIbmRecoveriesOracleDbChannelToMap(&dbChannelsItem)
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
			shellEvironmentVarsItemMap, err := dataSourceIbmRecoveriesKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
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

func dataSourceIbmRecoveriesCommonRecoverOracleAppTargetParamsOriginalSourceConfigToMap(model *backuprecoveryv1.CommonRecoverOracleAppTargetParamsOriginalSourceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestoreTimeUsecs != nil {
		modelMap["restore_time_usecs"] = flex.IntValue(model.RestoreTimeUsecs)
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := dataSourceIbmRecoveriesOracleDbChannelToMap(&dbChannelsItem)
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
			shellEvironmentVarsItemMap, err := dataSourceIbmRecoveriesKeyValuePairToMap(&shellEvironmentVarsItem)
			if err != nil {
				return modelMap, err
			}
			shellEvironmentVars = append(shellEvironmentVars, shellEvironmentVarsItemMap)
		}
		modelMap["shell_evironment_vars"] = shellEvironmentVars
	}
	if model.GranularRestoreInfo != nil {
		granularRestoreInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigGranularRestoreInfoToMap(model.GranularRestoreInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["granular_restore_info"] = []map[string]interface{}{granularRestoreInfoMap}
	}
	if model.OracleArchiveLogInfo != nil {
		oracleArchiveLogInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleArchiveLogInfoToMap(model.OracleArchiveLogInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_archive_log_info"] = []map[string]interface{}{oracleArchiveLogInfoMap}
	}
	if model.OracleRecoveryValidationInfo != nil {
		oracleRecoveryValidationInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigOracleRecoveryValidationInfoToMap(model.OracleRecoveryValidationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_recovery_validation_info"] = []map[string]interface{}{oracleRecoveryValidationInfoMap}
	}
	if model.RestoreSpfileOrPfileInfo != nil {
		restoreSpfileOrPfileInfoMap, err := dataSourceIbmRecoveriesCommonOracleAppSourceConfigRestoreSpfileOrPfileInfoToMap(model.RestoreSpfileOrPfileInfo)
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

func dataSourceIbmRecoveriesRecoverOracleAppParamsVlanConfigToMap(model *backuprecoveryv1.RecoverOracleAppParamsVlanConfig) (map[string]interface{}, error) {
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
