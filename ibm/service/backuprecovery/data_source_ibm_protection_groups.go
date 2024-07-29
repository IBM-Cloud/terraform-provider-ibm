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

func DataSourceIbmProtectionGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionGroupsRead,

		Schema: map[string]*schema.Schema{
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by a list of Protection Group ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"names": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by a list of Protection Group names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"policy_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"storage_domain_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter by Storage Domain id. Only Protection Groups writing data to this Storage Domain will be returned.",
			},
			"include_groups_with_datalock_only": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to only return Protection Groups with a datalock.",
			},
			"environments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by environment types such as 'kVMware', 'kView', etc. Only Protection Groups protecting the specified environment types are returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_active": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive.",
			},
			"is_deleted": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases.",
			},
			"is_paused": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned.",
			},
			"last_run_local_backup_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_run_replication_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_run_archival_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_run_cloud_spin_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_run_any_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_last_run_sla_violated": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, return Protection Groups for which last run SLA was violated.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which objects are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned.",
			},
			"include_last_run_info": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will include last run info. If it is false or not specified, the last run info won't be returned.",
			},
			"prune_excluded_source_ids": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user.",
			},
			"prune_source_ids": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will exclude the list of source IDs within the group specified.",
			},
			"use_cached_data": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.",
			},
			"protection_groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Protection Groups which were returned by the request.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the ID of the Protection Group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Protection Group.",
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the cluster ID.",
						},
						"region_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the region ID.",
						},
						"policy_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc.",
						},
						"priority": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the priority of the Protection Group.",
						},
						"storage_domain_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the Storage Domain (View Box) ID where this Protection Group writes data.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies a description of the Protection Group.",
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the time of day. Used for scheduling purposes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hour": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the hour of the day (0-23).",
									},
									"minute": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the minute of the hour (0-59).",
									},
									"time_zone": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.",
									},
								},
							},
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.",
						},
						"last_modified_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.",
						},
						"alert_policy": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies a policy for alerting users of the status of a Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_run_status": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the run status for which the user would like to receive alerts.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"alert_targets": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a list of targets to receive the alerts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"email_address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies an email address to receive an alert.",
												},
												"language": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the language of the delivery target. Default value is 'en-us'.",
												},
												"recipient_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the recipient type of email recipient. Default value is 'kTo'.",
												},
											},
										},
									},
									"raise_object_level_failure_alert": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether object level alerts are raised for backup failures after the backup run.",
									},
									"raise_object_level_failure_alert_after_last_attempt": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether object level alerts are raised for backup failures after last backup attempt.",
									},
									"raise_object_level_failure_alert_after_each_attempt": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether object level alerts are raised for backup failures after each backup attempt.",
									},
								},
							},
						},
						"sla": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the SLA parameters for this Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of run this rule should apply to.",
									},
									"sla_minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of minutes allotted to a run of the specified type before SLA is considered violated.",
									},
								},
							},
						},
						"qos_policy": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies whether the Protection Group will be written to HDD or SSD.",
						},
						"abort_in_blackouts": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false.",
						},
						"pause_in_blackouts": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true.",
						},
						"is_active": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the Protection Group is active or not.",
						},
						"is_deleted": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the Protection Group has been deleted.",
						},
						"is_paused": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted.",
						},
						"environment": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the environment of the Protection Group.",
						},
						"last_run": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the parameters which are common between Protection Group runs of all Protection Groups.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the ID of the Protection Group run.",
									},
									"protection_group_instance_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Protection Group instance Id. This field will be removed later.",
									},
									"protection_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ProtectionGroupId to which this run belongs.",
									},
									"is_replication_run": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if this protection run is a replication run.",
									},
									"origin_cluster_identifier": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the information about a cluster.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the id of the cluster.",
												},
												"cluster_incarnation_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the incarnation id of the cluster.",
												},
												"cluster_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the cluster.",
												},
											},
										},
									},
									"origin_protection_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ProtectionGroupId to which this run belongs on the primary cluster if this run is a replication run.",
									},
									"protection_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the Protection Group to which this run belongs.",
									},
									"is_local_snapshots_deleted": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if snapshots for this run has been deleted.",
									},
									"objects": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Snapahot, replication, archival results for each object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"object": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Object Summary.",
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
												"local_snapshot_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies information about backup run for an object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"snapshot_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Snapshot info for an object.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"snapshot_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of snapshot.",
																		},
																		"status_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A message decribing the status. This will be populated currently only for kWaitingForOlderBackupRun status.",
																		},
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"admitted_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"permit_grant_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.",
																		},
																		"queue_duration_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.",
																		},
																		"snapshot_creation_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about local snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical size of object(s) in bytes.",
																					},
																					"bytes_written": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total size of data in bytes written after taking backup.",
																					},
																					"bytes_read": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical bytes read for creating the snapshot.",
																					},
																				},
																			},
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for backup of the object.",
																		},
																		"indexing_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for the indexing of documents in an object.",
																		},
																		"stats_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Stats task for an object.",
																		},
																		"warnings": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies a list of warning messages.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"is_manually_deleted": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the snapshot is deleted manually.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
																		},
																		"total_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The total number of file and directory entities visited in this backup. Only applicable to file based backups.",
																		},
																		"backup_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
																		},
																		"data_lock_constraints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the dataLock constraints for local or target snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"failed_attempts": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Failed backup attempts for an object.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"admitted_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"permit_grant_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.",
																		},
																		"queue_duration_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.",
																		},
																		"snapshot_creation_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of the attempt for an object. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Pausing' indicates that the ongoing run is in the process of being paused. 'Resuming' indicates that the already paused run is in the process of being running again. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about local snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical size of object(s) in bytes.",
																					},
																					"bytes_written": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total size of data in bytes written after taking backup.",
																					},
																					"bytes_read": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical bytes read for creating the snapshot.",
																					},
																				},
																			},
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for an object..",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A message about the error if encountered while performing backup.",
																		},
																	},
																},
															},
														},
													},
												},
												"original_backup_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies information about backup run for an object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"snapshot_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Snapshot info for an object.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"snapshot_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of snapshot.",
																		},
																		"status_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A message decribing the status. This will be populated currently only for kWaitingForOlderBackupRun status.",
																		},
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"admitted_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"permit_grant_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.",
																		},
																		"queue_duration_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.",
																		},
																		"snapshot_creation_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about local snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical size of object(s) in bytes.",
																					},
																					"bytes_written": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total size of data in bytes written after taking backup.",
																					},
																					"bytes_read": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical bytes read for creating the snapshot.",
																					},
																				},
																			},
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for backup of the object.",
																		},
																		"indexing_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for the indexing of documents in an object.",
																		},
																		"stats_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Stats task for an object.",
																		},
																		"warnings": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies a list of warning messages.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"is_manually_deleted": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the snapshot is deleted manually.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
																		},
																		"total_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The total number of file and directory entities visited in this backup. Only applicable to file based backups.",
																		},
																		"backup_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
																		},
																		"data_lock_constraints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the dataLock constraints for local or target snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"failed_attempts": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Failed backup attempts for an object.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"admitted_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"permit_grant_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.",
																		},
																		"queue_duration_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.",
																		},
																		"snapshot_creation_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of the attempt for an object. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Pausing' indicates that the ongoing run is in the process of being paused. 'Resuming' indicates that the already paused run is in the process of being running again. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about local snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical size of object(s) in bytes.",
																					},
																					"bytes_written": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total size of data in bytes written after taking backup.",
																					},
																					"bytes_read": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical bytes read for creating the snapshot.",
																					},
																				},
																			},
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for an object..",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "A message about the error if encountered while performing backup.",
																		},
																	},
																},
															},
														},
													},
												},
												"replication_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies information about replication run for an object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"replication_target_results": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Replication result for a target.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cluster_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the id of the cluster.",
																		},
																		"cluster_incarnation_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the incarnation id of the cluster.",
																		},
																		"cluster_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the cluster.",
																		},
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"queued_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of the replication for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Message about the replication run.",
																		},
																		"percentage_completed": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the progress in percentage.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about replication data.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the total logical size in bytes.",
																					},
																					"logical_bytes_transferred": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the total logical bytes transferred.",
																					},
																					"physical_bytes_transferred": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the total physical bytes transferred.",
																					},
																				},
																			},
																		},
																		"is_manually_deleted": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the snapshot is deleted manually.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
																		},
																		"replication_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Task UID for a replication protection run. This is for tasks that are replicated from another cluster.",
																		},
																		"entries_changed": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the number of metadata actions completed during the protection run.",
																		},
																		"is_in_bound": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false.",
																		},
																		"data_lock_constraints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the dataLock constraints for local or target snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																					},
																				},
																			},
																		},
																		"on_legal_hold": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies the legal hold status for a replication target.",
																		},
																		"multi_object_replication": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether view based replication was used. In this case, the view containing all objects is replicated as a whole instead of replicating on a per object basis.",
																		},
																	},
																},
															},
														},
													},
												},
												"archival_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies information about archival run for an object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archival_target_results": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Archival result for an archival target.",
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
																		"run_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.",
																		},
																		"is_sla_violated": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Indicated if SLA has been violated for this run.",
																		},
																		"snapshot_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Snapshot id for a successful snapshot. This field will not be set if the archival Run fails to take the snapshot.",
																		},
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.",
																		},
																		"queued_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time when the archival is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"is_incremental": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Whether this is an incremental archive. If set to true, this is an incremental archive, otherwise this is a full archive.",
																		},
																		"is_forever_incremental": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Whether this is forever incremental or not.",
																		},
																		"is_cad_archive": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Whether this is CAD archive or not.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of the replication run for an archival target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Message about the archival run.",
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task id for archival.",
																		},
																		"stats_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Run Stats task id for archival.",
																		},
																		"indexing_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task for indexing.",
																		},
																		"successful_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of objects for which backup was successful.",
																		},
																		"failed_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of objects for which backup failed.",
																		},
																		"cancelled_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of objects for which backup was cancelled.",
																		},
																		"successful_app_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of app objects for which backup was successful.",
																		},
																		"failed_app_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of app objects for which backup failed.",
																		},
																		"cancelled_app_objects_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the count of app objects for which backup was cancelled.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about archival data.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"logical_size_bytes": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the logicalSizeBytes.",
																					},
																					"bytes_read": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies total logical bytes read for creating the snapshot.",
																					},
																					"logical_bytes_transferred": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the logical bytes transferred.",
																					},
																					"physical_bytes_transferred": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the physical bytes transferred.",
																					},
																					"avg_logical_transfer_rate_bps": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the average rate of transfer in bytes per second.",
																					},
																					"file_walk_done": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether the file system walk is done. Only applicable to file based backups.",
																					},
																					"total_file_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.",
																					},
																					"backup_file_count": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
																					},
																				},
																			},
																		},
																		"is_manually_deleted": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the snapshot is deleted manually.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																		},
																		"data_lock_constraints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the dataLock constraints for local or target snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																					},
																				},
																			},
																		},
																		"on_legal_hold": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies the legal hold status for a archival target.",
																		},
																		"worm_properties": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the WORM related properties for this archive.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"is_archive_worm_compliant": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether this archive run is WORM compliant.",
																					},
																					"worm_non_compliance_reason": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies reason of archive not being worm compliant.",
																					},
																					"worm_expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the time at which the WORM protection expires.",
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
												"cloud_spin_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies information about Cloud Spin run for an object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cloud_spin_target_results": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Cloud Spin result for a target.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the unique id of the cloud spin entity.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the already added cloud spin target.",
																		},
																		"start_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"end_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Status of the Cloud Spin for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Message about the Cloud Spin run.",
																		},
																		"stats": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies statistics about Cloud Spin data.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"physical_bytes_transferred": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the physical bytes transferred.",
																					},
																				},
																			},
																		},
																		"is_manually_deleted": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the snapshot is deleted manually.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
																		},
																		"cloudspin_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Task ID for a CloudSpin protection run.",
																		},
																		"progress_task_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Progress monitor task id for Cloud Spin run.",
																		},
																		"data_lock_constraints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the dataLock constraints for local or target snapshot.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"expiry_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																					},
																				},
																			},
																		},
																		"on_legal_hold": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies the legal hold status for a cloud spin target.",
																		},
																	},
																},
															},
														},
													},
												},
												"on_legal_hold": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if object's snapshot is on legal hold.",
												},
											},
										},
									},
									"local_backup_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies summary information about local snapshot run across all objects.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"run_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. 'kStorageArraySnapshot' indicates storage array snapshot backup.",
												},
												"is_sla_violated": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicated if SLA has been violated for this run.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of backup run in Unix epoch Timestamp(in microseconds).",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of backup run in Unix epoch Timestamp(in microseconds).",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of the backup run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
												},
												"messages": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Message about the backup run.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"successful_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup was successful.",
												},
												"failed_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup failed.",
												},
												"cancelled_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup was cancelled.",
												},
												"successful_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup was successful.",
												},
												"failed_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup failed.",
												},
												"cancelled_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup was cancelled.",
												},
												"local_snapshot_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies statistics about local snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logical_size_bytes": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total logical size of object(s) in bytes.",
															},
															"bytes_written": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total size of data in bytes written after taking backup.",
															},
															"bytes_read": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total logical bytes read for creating the snapshot.",
															},
														},
													},
												},
												"indexing_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Progress monitor task for indexing.",
												},
												"progress_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Progress monitor task id for local backup run.",
												},
												"stats_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Stats task id for local backup run.",
												},
												"data_lock": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "This field is deprecated. Use DataLockConstraints field instead.",
												},
												"local_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Task ID for a local protection run.",
												},
												"data_lock_constraints": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the dataLock constraints for local or target snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"expiry_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
															},
														},
													},
												},
											},
										},
									},
									"original_backup_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies summary information about local snapshot run across all objects.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"run_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. 'kStorageArraySnapshot' indicates storage array snapshot backup.",
												},
												"is_sla_violated": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicated if SLA has been violated for this run.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of backup run in Unix epoch Timestamp(in microseconds).",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of backup run in Unix epoch Timestamp(in microseconds).",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of the backup run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
												},
												"messages": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Message about the backup run.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"successful_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup was successful.",
												},
												"failed_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup failed.",
												},
												"cancelled_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of objects for which backup was cancelled.",
												},
												"successful_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup was successful.",
												},
												"failed_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup failed.",
												},
												"cancelled_app_objects_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the count of app objects for which backup was cancelled.",
												},
												"local_snapshot_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies statistics about local snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logical_size_bytes": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total logical size of object(s) in bytes.",
															},
															"bytes_written": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total size of data in bytes written after taking backup.",
															},
															"bytes_read": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies total logical bytes read for creating the snapshot.",
															},
														},
													},
												},
												"indexing_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Progress monitor task for indexing.",
												},
												"progress_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Progress monitor task id for local backup run.",
												},
												"stats_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Stats task id for local backup run.",
												},
												"data_lock": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "This field is deprecated. Use DataLockConstraints field instead.",
												},
												"local_task_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Task ID for a local protection run.",
												},
												"data_lock_constraints": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the dataLock constraints for local or target snapshot.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"expiry_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
															},
														},
													},
												},
											},
										},
									},
									"replication_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies summary information about replication run.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"replication_target_results": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Replication results for each replication target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the id of the cluster.",
															},
															"cluster_incarnation_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the incarnation id of the cluster.",
															},
															"cluster_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the cluster.",
															},
															"start_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"end_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"queued_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status of the replication for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
															},
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Message about the replication run.",
															},
															"percentage_completed": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the progress in percentage.",
															},
															"stats": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies statistics about replication data.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"logical_size_bytes": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the total logical size in bytes.",
																		},
																		"logical_bytes_transferred": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the total logical bytes transferred.",
																		},
																		"physical_bytes_transferred": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the total physical bytes transferred.",
																		},
																	},
																},
															},
															"is_manually_deleted": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the snapshot is deleted manually.",
															},
															"expiry_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
															},
															"replication_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Task UID for a replication protection run. This is for tasks that are replicated from another cluster.",
															},
															"entries_changed": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of metadata actions completed during the protection run.",
															},
															"is_in_bound": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false.",
															},
															"data_lock_constraints": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the dataLock constraints for local or target snapshot.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																		},
																	},
																},
															},
															"on_legal_hold": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the legal hold status for a replication target.",
															},
															"multi_object_replication": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether view based replication was used. In this case, the view containing all objects is replicated as a whole instead of replicating on a per object basis.",
															},
														},
													},
												},
											},
										},
									},
									"archival_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies summary information about archival run.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"archival_target_results": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Archival results for each archival target.",
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
															"run_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.",
															},
															"is_sla_violated": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Indicated if SLA has been violated for this run.",
															},
															"snapshot_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Snapshot id for a successful snapshot. This field will not be set if the archival Run fails to take the snapshot.",
															},
															"start_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the start time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.",
															},
															"end_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the end time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.",
															},
															"queued_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time when the archival is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"is_incremental": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether this is an incremental archive. If set to true, this is an incremental archive, otherwise this is a full archive.",
															},
															"is_forever_incremental": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether this is forever incremental or not.",
															},
															"is_cad_archive": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether this is CAD archive or not.",
															},
															"status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status of the replication run for an archival target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
															},
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Message about the archival run.",
															},
															"progress_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Progress monitor task id for archival.",
															},
															"stats_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Run Stats task id for archival.",
															},
															"indexing_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Progress monitor task for indexing.",
															},
															"successful_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of objects for which backup was successful.",
															},
															"failed_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of objects for which backup failed.",
															},
															"cancelled_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of objects for which backup was cancelled.",
															},
															"successful_app_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of app objects for which backup was successful.",
															},
															"failed_app_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of app objects for which backup failed.",
															},
															"cancelled_app_objects_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the count of app objects for which backup was cancelled.",
															},
															"stats": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies statistics about archival data.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"logical_size_bytes": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the logicalSizeBytes.",
																		},
																		"bytes_read": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies total logical bytes read for creating the snapshot.",
																		},
																		"logical_bytes_transferred": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the logical bytes transferred.",
																		},
																		"physical_bytes_transferred": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the physical bytes transferred.",
																		},
																		"avg_logical_transfer_rate_bps": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the average rate of transfer in bytes per second.",
																		},
																		"file_walk_done": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the file system walk is done. Only applicable to file based backups.",
																		},
																		"total_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.",
																		},
																		"backup_file_count": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
																		},
																	},
																},
															},
															"is_manually_deleted": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the snapshot is deleted manually.",
															},
															"expiry_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
															},
															"data_lock_constraints": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the dataLock constraints for local or target snapshot.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																		},
																	},
																},
															},
															"on_legal_hold": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the legal hold status for a archival target.",
															},
															"worm_properties": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the WORM related properties for this archive.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"is_archive_worm_compliant": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether this archive run is WORM compliant.",
																		},
																		"worm_non_compliance_reason": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies reason of archive not being worm compliant.",
																		},
																		"worm_expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the time at which the WORM protection expires.",
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
									"cloud_spin_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies summary information about cloud spin run.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud_spin_target_results": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Cloud Spin results for each Cloud Spin target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the unique id of the cloud spin entity.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the already added cloud spin target.",
															},
															"start_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"end_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.",
															},
															"status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status of the Cloud Spin for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
															},
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Message about the Cloud Spin run.",
															},
															"stats": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies statistics about Cloud Spin data.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"physical_bytes_transferred": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the physical bytes transferred.",
																		},
																	},
																},
															},
															"is_manually_deleted": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the snapshot is deleted manually.",
															},
															"expiry_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.",
															},
															"cloudspin_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Task ID for a CloudSpin protection run.",
															},
															"progress_task_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Progress monitor task id for Cloud Spin run.",
															},
															"data_lock_constraints": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the dataLock constraints for local or target snapshot.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"expiry_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).",
																		},
																	},
																},
															},
															"on_legal_hold": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the legal hold status for a cloud spin target.",
															},
														},
													},
												},
											},
										},
									},
									"on_legal_hold": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if the Protection Run is on legal hold.",
									},
									"permissions": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of tenants that have permissions for this protection group run.",
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
									"is_cloud_archival_direct": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the run is a CAD run if cloud archive direct feature is enabled. If this field is true, the primary backup copy will only be available at the given archived location.",
									},
									"has_local_snapshot": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the run has a local snapshot. For cloud retrieved runs there may not be local snapshots.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the environment of the Protection Group.",
									},
									"externally_triggered_backup_tag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag of externally triggered backup job.",
									},
								},
							},
						},
						"permissions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of tenants that have permissions for this protection group.",
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
						"is_protect_once": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group.",
						},
						"missing_entities": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Information about missing entities.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the ID of the object.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the object.",
									},
									"parent_source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the id of the parent source of the object.",
									},
									"parent_source_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the parent source of the object.",
									},
								},
							},
						},
						"invalid_entities": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Information about invalid entities. An entity will be considered invalid if it is part of an active protection group but has lost compatibility for the given backup type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the ID of the object.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the object.",
									},
									"parent_source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the id of the parent source of the object.",
									},
									"parent_source_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the parent source of the object.",
									},
								},
							},
						},
						"num_protected_objects": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of protected objects of the Protection Group.",
						},
						"advanced_configs": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the advanced configuration for a protection job.",
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
						"physical_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protection_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the Physical Protection Group type.",
									},
									"volume_protection_type_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters which are specific to Volume based physical Protection Groups.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"objects": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the ID of the object protected.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the object protected.",
															},
															"volume_guids": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the list of GUIDs of volumes protected. If empty, then all volumes will be protected by default.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"enable_system_backup": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether or not to take a system backup. Applicable only for windows sources.",
															},
														},
													},
												},
												"indexing_policy": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enable_indexing": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.",
															},
															"include_paths": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"exclude_paths": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"perform_source_side_deduplication": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether or not to perform source side deduplication on this Protection Group.",
												},
												"quiesce": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.",
												},
												"continue_on_quiesce_failure": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to continue backing up on quiesce failure.",
												},
												"incremental_backup_after_restart": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether or not to perform an incremental backup after the server restarts. This is applicable to windows environments.",
												},
												"pre_post_script": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the params for pre and post scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"pre_script": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the common params for PreBackup scripts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the absolute path to the script on the remote host.",
																		},
																		"params": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
																		},
																		"timeout_secs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
																		},
																		"is_active": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the script should be enabled, default value set to true.",
																		},
																		"continue_on_error": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
																		},
																	},
																},
															},
															"post_script": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the common params for PostBackup scripts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the absolute path to the script on the remote host.",
																		},
																		"params": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
																		},
																		"timeout_secs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
																		},
																		"is_active": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the script should be enabled, default value set to true.",
																		},
																	},
																},
															},
														},
													},
												},
												"dedup_exclusion_source_ids": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies ids of sources for which deduplication has to be disabled.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"excluded_vss_writers": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies writer names which should be excluded from physical volume based backups.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"cobmr_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to take a CoBMR backup.",
												},
											},
										},
									},
									"file_protection_type_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters which are specific to Physical related Protection Groups.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"objects": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the list of objects protected by this Protection Group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the ID of the object protected.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the object protected.",
															},
															"file_paths": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies a list of file paths to be protected by this Protection Group.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"included_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies a path to be included on the source. All paths under this path will be included unless they are specifically mentioned in excluded paths.",
																		},
																		"excluded_paths": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies a set of paths nested under the include path which should be excluded from the Protection Group.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"skip_nested_volumes": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether to skip any nested volumes (both local and network) that are mounted under include path. Applicable only for windows sources.",
																		},
																	},
																},
															},
															"uses_path_level_skip_nested_volume_setting": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether path level or object level skip nested volume setting will be used.",
															},
															"nested_volume_types_to_skip": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies mount types of nested volumes to be skipped.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"follow_nas_symlink_target": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to follow NAS target pointed by symlink for windows sources.",
															},
															"metadata_file_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the path of metadatafile on source. This file contains absolute paths of files that needs to be backed up on the same source.",
															},
														},
													},
												},
												"indexing_policy": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enable_indexing": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.",
															},
															"include_paths": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"exclude_paths": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"perform_source_side_deduplication": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether or not to perform source side deduplication on this Protection Group.",
												},
												"perform_brick_based_deduplication": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether or not to perform brick based deduplication on this Protection Group.",
												},
												"task_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the timeouts for all the objects inside this Protection Group, for both full and incremental backups.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"timeout_mins": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the timeout in mins.",
															},
															"backup_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The scheduled backup type(kFull, kRegular etc.).",
															},
														},
													},
												},
												"quiesce": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.",
												},
												"continue_on_quiesce_failure": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to continue backing up on quiesce failure.",
												},
												"cobmr_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to take CoBMR backup.",
												},
												"pre_post_script": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the params for pre and post scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"pre_script": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the common params for PreBackup scripts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the absolute path to the script on the remote host.",
																		},
																		"params": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
																		},
																		"timeout_secs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
																		},
																		"is_active": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the script should be enabled, default value set to true.",
																		},
																		"continue_on_error": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
																		},
																	},
																},
															},
															"post_script": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the common params for PostBackup scripts.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the absolute path to the script on the remote host.",
																		},
																		"params": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
																		},
																		"timeout_secs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
																		},
																		"is_active": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the script should be enabled, default value set to true.",
																		},
																	},
																},
															},
														},
													},
												},
												"dedup_exclusion_source_ids": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies ids of sources for which deduplication has to be disabled.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"global_exclude_paths": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies global exclude filters which are applied to all sources in a job.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"global_exclude_fs": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies global exclude filesystems which are applied to all sources in a job.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ignorable_errors": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Errors to be ignored in error db.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allow_parallel_runs": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether or not this job can have parallel runs.",
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
							Description: "Specifies the parameters to create Oracle Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"objects": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of object ids to be protected.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"source_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the id of the host on which databases are hosted.",
												},
												"source_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the host on which databases are hosted.",
												},
												"db_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the properties of the Oracle databases.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"database_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the id of the Oracle database.",
															},
															"database_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the Oracle database.",
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
														},
													},
												},
											},
										},
									},
									"persist_mountpoints": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. Defaults to true if value is null to handle the backward compatibility for the upgrade case.",
									},
									"vlan_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies VLAN params associated with the backup/restore operation.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
												},
												"disable_vlan": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
												},
												"interface_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
												},
											},
										},
									},
									"pre_post_script": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the params for pre and post scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_script": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the common params for PreBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the script should be enabled, default value set to true.",
															},
															"continue_on_error": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
															},
														},
													},
												},
												"post_script": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the common params for PostBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the script should be enabled, default value set to true.",
															},
														},
													},
												},
											},
										},
									},
									"log_auto_kill_timeout_secs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Time in seconds after which the log backup of the database in given backup job should be auto-killed.",
									},
									"incr_auto_kill_timeout_secs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Time in seconds after which the incremental backup of the database in given backup job should be auto-killed.",
									},
									"full_auto_kill_timeout_secs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Time in seconds after which the full backup of the database in given backup job should be auto-killed.",
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

func dataSourceIbmProtectionGroupsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionGroupsOptions := &backuprecoveryv1.GetProtectionGroupsOptions{}

	if _, ok := d.GetOk("request_initiator_type"); ok {
		getProtectionGroupsOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("ids"); ok {
		var ids []string
		for _, v := range d.Get("ids").([]interface{}) {
			idsItem := v.(string)
			ids = append(ids, idsItem)
		}
		getProtectionGroupsOptions.SetIds(ids)
	}
	if _, ok := d.GetOk("names"); ok {
		var names []string
		for _, v := range d.Get("names").([]interface{}) {
			namesItem := v.(string)
			names = append(names, namesItem)
		}
		getProtectionGroupsOptions.SetNames(names)
	}
	if _, ok := d.GetOk("policy_ids"); ok {
		var policyIds []string
		for _, v := range d.Get("policy_ids").([]interface{}) {
			policyIdsItem := v.(string)
			policyIds = append(policyIds, policyIdsItem)
		}
		getProtectionGroupsOptions.SetPolicyIds(policyIds)
	}
	if _, ok := d.GetOk("storage_domain_id"); ok {
		getProtectionGroupsOptions.SetStorageDomainID(int64(d.Get("storage_domain_id").(int)))
	}
	if _, ok := d.GetOk("include_groups_with_datalock_only"); ok {
		getProtectionGroupsOptions.SetIncludeGroupsWithDatalockOnly(d.Get("include_groups_with_datalock_only").(bool))
	}
	if _, ok := d.GetOk("environments"); ok {
		var environments []string
		for _, v := range d.Get("environments").([]interface{}) {
			environmentsItem := v.(string)
			environments = append(environments, environmentsItem)
		}
		getProtectionGroupsOptions.SetEnvironments(environments)
	}
	if _, ok := d.GetOk("is_active"); ok {
		getProtectionGroupsOptions.SetIsActive(d.Get("is_active").(bool))
	}
	if _, ok := d.GetOk("is_deleted"); ok {
		getProtectionGroupsOptions.SetIsDeleted(d.Get("is_deleted").(bool))
	}
	if _, ok := d.GetOk("is_paused"); ok {
		getProtectionGroupsOptions.SetIsPaused(d.Get("is_paused").(bool))
	}
	if _, ok := d.GetOk("last_run_local_backup_status"); ok {
		var lastRunLocalBackupStatus []string
		for _, v := range d.Get("last_run_local_backup_status").([]interface{}) {
			lastRunLocalBackupStatusItem := v.(string)
			lastRunLocalBackupStatus = append(lastRunLocalBackupStatus, lastRunLocalBackupStatusItem)
		}
		getProtectionGroupsOptions.SetLastRunLocalBackupStatus(lastRunLocalBackupStatus)
	}
	if _, ok := d.GetOk("last_run_replication_status"); ok {
		var lastRunReplicationStatus []string
		for _, v := range d.Get("last_run_replication_status").([]interface{}) {
			lastRunReplicationStatusItem := v.(string)
			lastRunReplicationStatus = append(lastRunReplicationStatus, lastRunReplicationStatusItem)
		}
		getProtectionGroupsOptions.SetLastRunReplicationStatus(lastRunReplicationStatus)
	}
	if _, ok := d.GetOk("last_run_archival_status"); ok {
		var lastRunArchivalStatus []string
		for _, v := range d.Get("last_run_archival_status").([]interface{}) {
			lastRunArchivalStatusItem := v.(string)
			lastRunArchivalStatus = append(lastRunArchivalStatus, lastRunArchivalStatusItem)
		}
		getProtectionGroupsOptions.SetLastRunArchivalStatus(lastRunArchivalStatus)
	}
	if _, ok := d.GetOk("last_run_cloud_spin_status"); ok {
		var lastRunCloudSpinStatus []string
		for _, v := range d.Get("last_run_cloud_spin_status").([]interface{}) {
			lastRunCloudSpinStatusItem := v.(string)
			lastRunCloudSpinStatus = append(lastRunCloudSpinStatus, lastRunCloudSpinStatusItem)
		}
		getProtectionGroupsOptions.SetLastRunCloudSpinStatus(lastRunCloudSpinStatus)
	}
	if _, ok := d.GetOk("last_run_any_status"); ok {
		var lastRunAnyStatus []string
		for _, v := range d.Get("last_run_any_status").([]interface{}) {
			lastRunAnyStatusItem := v.(string)
			lastRunAnyStatus = append(lastRunAnyStatus, lastRunAnyStatusItem)
		}
		getProtectionGroupsOptions.SetLastRunAnyStatus(lastRunAnyStatus)
	}
	if _, ok := d.GetOk("is_last_run_sla_violated"); ok {
		getProtectionGroupsOptions.SetIsLastRunSlaViolated(d.Get("is_last_run_sla_violated").(bool))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getProtectionGroupsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getProtectionGroupsOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("include_last_run_info"); ok {
		getProtectionGroupsOptions.SetIncludeLastRunInfo(d.Get("include_last_run_info").(bool))
	}
	if _, ok := d.GetOk("prune_excluded_source_ids"); ok {
		getProtectionGroupsOptions.SetPruneExcludedSourceIds(d.Get("prune_excluded_source_ids").(bool))
	}
	if _, ok := d.GetOk("prune_source_ids"); ok {
		getProtectionGroupsOptions.SetPruneSourceIds(d.Get("prune_source_ids").(bool))
	}
	if _, ok := d.GetOk("use_cached_data"); ok {
		getProtectionGroupsOptions.SetUseCachedData(d.Get("use_cached_data").(bool))
	}

	protectionGroupsResponse, response, err := backupRecoveryClient.GetProtectionGroupsWithContext(context, getProtectionGroupsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionGroupsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionGroupsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmProtectionGroupsID(d))

	protectionGroups := []map[string]interface{}{}
	if protectionGroupsResponse.ProtectionGroups != nil {
		for _, modelItem := range protectionGroupsResponse.ProtectionGroups {
			modelMap, err := dataSourceIbmProtectionGroupsProtectionGroupResponseToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			protectionGroups = append(protectionGroups, modelMap)
		}
	}
	if err = d.Set("protection_groups", protectionGroups); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting protection_groups %s", err))
	}

	return nil
}

// dataSourceIbmProtectionGroupsID returns a reasonable ID for the list.
func dataSourceIbmProtectionGroupsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmProtectionGroupsProtectionGroupResponseToMap(model *backuprecoveryv1.ProtectionGroupResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = model.ClusterID
	}
	if model.RegionID != nil {
		modelMap["region_id"] = model.RegionID
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = model.PolicyID
	}
	if model.Priority != nil {
		modelMap["priority"] = model.Priority
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.StartTime != nil {
		startTimeMap, err := dataSourceIbmProtectionGroupsTimeOfDayToMap(model.StartTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.LastModifiedTimestampUsecs != nil {
		modelMap["last_modified_timestamp_usecs"] = flex.IntValue(model.LastModifiedTimestampUsecs)
	}
	if model.AlertPolicy != nil {
		alertPolicyMap, err := dataSourceIbmProtectionGroupsProtectionGroupAlertingPolicyToMap(model.AlertPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["alert_policy"] = []map[string]interface{}{alertPolicyMap}
	}
	if model.Sla != nil {
		sla := []map[string]interface{}{}
		for _, slaItem := range model.Sla {
			slaItemMap, err := dataSourceIbmProtectionGroupsSlaRuleToMap(&slaItem)
			if err != nil {
				return modelMap, err
			}
			sla = append(sla, slaItemMap)
		}
		modelMap["sla"] = sla
	}
	if model.QosPolicy != nil {
		modelMap["qos_policy"] = model.QosPolicy
	}
	if model.AbortInBlackouts != nil {
		modelMap["abort_in_blackouts"] = model.AbortInBlackouts
	}
	if model.PauseInBlackouts != nil {
		modelMap["pause_in_blackouts"] = model.PauseInBlackouts
	}
	if model.IsActive != nil {
		modelMap["is_active"] = model.IsActive
	}
	if model.IsDeleted != nil {
		modelMap["is_deleted"] = model.IsDeleted
	}
	if model.IsPaused != nil {
		modelMap["is_paused"] = model.IsPaused
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.LastRun != nil {
		lastRunMap, err := dataSourceIbmProtectionGroupsProtectionGroupRunToMap(model.LastRun)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_run"] = []map[string]interface{}{lastRunMap}
	}
	if model.Permissions != nil {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range model.Permissions {
			permissionsItemMap, err := dataSourceIbmProtectionGroupsTenantToMap(&permissionsItem)
			if err != nil {
				return modelMap, err
			}
			permissions = append(permissions, permissionsItemMap)
		}
		modelMap["permissions"] = permissions
	}
	if model.IsProtectOnce != nil {
		modelMap["is_protect_once"] = model.IsProtectOnce
	}
	if model.MissingEntities != nil {
		missingEntities := []map[string]interface{}{}
		for _, missingEntitiesItem := range model.MissingEntities {
			missingEntitiesItemMap, err := dataSourceIbmProtectionGroupsMissingEntityParamsToMap(&missingEntitiesItem)
			if err != nil {
				return modelMap, err
			}
			missingEntities = append(missingEntities, missingEntitiesItemMap)
		}
		modelMap["missing_entities"] = missingEntities
	}
	if model.InvalidEntities != nil {
		invalidEntities := []map[string]interface{}{}
		for _, invalidEntitiesItem := range model.InvalidEntities {
			invalidEntitiesItemMap, err := dataSourceIbmProtectionGroupsMissingEntityParamsToMap(&invalidEntitiesItem)
			if err != nil {
				return modelMap, err
			}
			invalidEntities = append(invalidEntities, invalidEntitiesItemMap)
		}
		modelMap["invalid_entities"] = invalidEntities
	}
	if model.NumProtectedObjects != nil {
		modelMap["num_protected_objects"] = flex.IntValue(model.NumProtectedObjects)
	}
	if model.AdvancedConfigs != nil {
		advancedConfigs := []map[string]interface{}{}
		for _, advancedConfigsItem := range model.AdvancedConfigs {
			advancedConfigsItemMap, err := dataSourceIbmProtectionGroupsKeyValuePairToMap(&advancedConfigsItem)
			if err != nil {
				return modelMap, err
			}
			advancedConfigs = append(advancedConfigs, advancedConfigsItemMap)
		}
		modelMap["advanced_configs"] = advancedConfigs
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmProtectionGroupsPhysicalProtectionGroupParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmProtectionGroupsOracleProtectionGroupParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsTimeOfDayToMap(model *backuprecoveryv1.TimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["hour"] = flex.IntValue(model.Hour)
	modelMap["minute"] = flex.IntValue(model.Minute)
	if model.TimeZone != nil {
		modelMap["time_zone"] = model.TimeZone
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsProtectionGroupAlertingPolicyToMap(model *backuprecoveryv1.ProtectionGroupAlertingPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["backup_run_status"] = model.BackupRunStatus
	if model.AlertTargets != nil {
		alertTargets := []map[string]interface{}{}
		for _, alertTargetsItem := range model.AlertTargets {
			alertTargetsItemMap, err := dataSourceIbmProtectionGroupsAlertTargetToMap(&alertTargetsItem)
			if err != nil {
				return modelMap, err
			}
			alertTargets = append(alertTargets, alertTargetsItemMap)
		}
		modelMap["alert_targets"] = alertTargets
	}
	if model.RaiseObjectLevelFailureAlert != nil {
		modelMap["raise_object_level_failure_alert"] = model.RaiseObjectLevelFailureAlert
	}
	if model.RaiseObjectLevelFailureAlertAfterLastAttempt != nil {
		modelMap["raise_object_level_failure_alert_after_last_attempt"] = model.RaiseObjectLevelFailureAlertAfterLastAttempt
	}
	if model.RaiseObjectLevelFailureAlertAfterEachAttempt != nil {
		modelMap["raise_object_level_failure_alert_after_each_attempt"] = model.RaiseObjectLevelFailureAlertAfterEachAttempt
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsAlertTargetToMap(model *backuprecoveryv1.AlertTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["email_address"] = model.EmailAddress
	if model.Language != nil {
		modelMap["language"] = model.Language
	}
	if model.RecipientType != nil {
		modelMap["recipient_type"] = model.RecipientType
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsSlaRuleToMap(model *backuprecoveryv1.SlaRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.SlaMinutes != nil {
		modelMap["sla_minutes"] = flex.IntValue(model.SlaMinutes)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsProtectionGroupRunToMap(model *backuprecoveryv1.ProtectionGroupRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.ProtectionGroupInstanceID != nil {
		modelMap["protection_group_instance_id"] = flex.IntValue(model.ProtectionGroupInstanceID)
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.IsReplicationRun != nil {
		modelMap["is_replication_run"] = model.IsReplicationRun
	}
	if model.OriginClusterIdentifier != nil {
		originClusterIdentifierMap, err := dataSourceIbmProtectionGroupsClusterIdentifierToMap(model.OriginClusterIdentifier)
		if err != nil {
			return modelMap, err
		}
		modelMap["origin_cluster_identifier"] = []map[string]interface{}{originClusterIdentifierMap}
	}
	if model.OriginProtectionGroupID != nil {
		modelMap["origin_protection_group_id"] = model.OriginProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.IsLocalSnapshotsDeleted != nil {
		modelMap["is_local_snapshots_deleted"] = model.IsLocalSnapshotsDeleted
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionGroupsObjectRunResultToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	if model.LocalBackupInfo != nil {
		localBackupInfoMap, err := dataSourceIbmProtectionGroupsBackupRunSummaryToMap(model.LocalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_backup_info"] = []map[string]interface{}{localBackupInfoMap}
	}
	if model.OriginalBackupInfo != nil {
		originalBackupInfoMap, err := dataSourceIbmProtectionGroupsBackupRunSummaryToMap(model.OriginalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_backup_info"] = []map[string]interface{}{originalBackupInfoMap}
	}
	if model.ReplicationInfo != nil {
		replicationInfoMap, err := dataSourceIbmProtectionGroupsReplicationRunSummaryToMap(model.ReplicationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["replication_info"] = []map[string]interface{}{replicationInfoMap}
	}
	if model.ArchivalInfo != nil {
		archivalInfoMap, err := dataSourceIbmProtectionGroupsArchivalRunSummaryToMap(model.ArchivalInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_info"] = []map[string]interface{}{archivalInfoMap}
	}
	if model.CloudSpinInfo != nil {
		cloudSpinInfoMap, err := dataSourceIbmProtectionGroupsCloudSpinRunSummaryToMap(model.CloudSpinInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["cloud_spin_info"] = []map[string]interface{}{cloudSpinInfoMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	if model.Permissions != nil {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range model.Permissions {
			permissionsItemMap, err := dataSourceIbmProtectionGroupsTenantToMap(&permissionsItem)
			if err != nil {
				return modelMap, err
			}
			permissions = append(permissions, permissionsItemMap)
		}
		modelMap["permissions"] = permissions
	}
	if model.IsCloudArchivalDirect != nil {
		modelMap["is_cloud_archival_direct"] = model.IsCloudArchivalDirect
	}
	if model.HasLocalSnapshot != nil {
		modelMap["has_local_snapshot"] = model.HasLocalSnapshot
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.ExternallyTriggeredBackupTag != nil {
		modelMap["externally_triggered_backup_tag"] = model.ExternallyTriggeredBackupTag
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsClusterIdentifierToMap(model *backuprecoveryv1.ClusterIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = model.ClusterName
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsObjectRunResultToMap(model *backuprecoveryv1.ObjectRunResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Object != nil {
		objectMap, err := dataSourceIbmProtectionGroupsObjectSummaryToMap(model.Object)
		if err != nil {
			return modelMap, err
		}
		modelMap["object"] = []map[string]interface{}{objectMap}
	}
	if model.LocalSnapshotInfo != nil {
		localSnapshotInfoMap, err := dataSourceIbmProtectionGroupsBackupRunToMap(model.LocalSnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_snapshot_info"] = []map[string]interface{}{localSnapshotInfoMap}
	}
	if model.OriginalBackupInfo != nil {
		originalBackupInfoMap, err := dataSourceIbmProtectionGroupsBackupRunToMap(model.OriginalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_backup_info"] = []map[string]interface{}{originalBackupInfoMap}
	}
	if model.ReplicationInfo != nil {
		replicationInfoMap, err := dataSourceIbmProtectionGroupsReplicationRunToMap(model.ReplicationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["replication_info"] = []map[string]interface{}{replicationInfoMap}
	}
	if model.ArchivalInfo != nil {
		archivalInfoMap, err := dataSourceIbmProtectionGroupsArchivalRunToMap(model.ArchivalInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_info"] = []map[string]interface{}{archivalInfoMap}
	}
	if model.CloudSpinInfo != nil {
		cloudSpinInfoMap, err := dataSourceIbmProtectionGroupsCloudSpinRunToMap(model.CloudSpinInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["cloud_spin_info"] = []map[string]interface{}{cloudSpinInfoMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsObjectSummaryToMap(model *backuprecoveryv1.ObjectSummary) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupsBackupRunToMap(model *backuprecoveryv1.BackupRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SnapshotInfo != nil {
		snapshotInfoMap, err := dataSourceIbmProtectionGroupsSnapshotInfoToMap(model.SnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["snapshot_info"] = []map[string]interface{}{snapshotInfoMap}
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := dataSourceIbmProtectionGroupsBackupAttemptToMap(&failedAttemptsItem)
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsSnapshotInfoToMap(model *backuprecoveryv1.SnapshotInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SnapshotID != nil {
		modelMap["snapshot_id"] = model.SnapshotID
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.AdmittedTimeUsecs != nil {
		modelMap["admitted_time_usecs"] = flex.IntValue(model.AdmittedTimeUsecs)
	}
	if model.PermitGrantTimeUsecs != nil {
		modelMap["permit_grant_time_usecs"] = flex.IntValue(model.PermitGrantTimeUsecs)
	}
	if model.QueueDurationUsecs != nil {
		modelMap["queue_duration_usecs"] = flex.IntValue(model.QueueDurationUsecs)
	}
	if model.SnapshotCreationTimeUsecs != nil {
		modelMap["snapshot_creation_time_usecs"] = flex.IntValue(model.SnapshotCreationTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionGroupsBackupDataStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.IndexingTaskID != nil {
		modelMap["indexing_task_id"] = model.IndexingTaskID
	}
	if model.StatsTaskID != nil {
		modelMap["stats_task_id"] = model.StatsTaskID
	}
	if model.Warnings != nil {
		modelMap["warnings"] = model.Warnings
	}
	if model.IsManuallyDeleted != nil {
		modelMap["is_manually_deleted"] = model.IsManuallyDeleted
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	if model.TotalFileCount != nil {
		modelMap["total_file_count"] = flex.IntValue(model.TotalFileCount)
	}
	if model.BackupFileCount != nil {
		modelMap["backup_file_count"] = flex.IntValue(model.BackupFileCount)
	}
	if model.DataLockConstraints != nil {
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsBackupDataStatsToMap(model *backuprecoveryv1.BackupDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.BytesWritten != nil {
		modelMap["bytes_written"] = flex.IntValue(model.BytesWritten)
	}
	if model.BytesRead != nil {
		modelMap["bytes_read"] = flex.IntValue(model.BytesRead)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model *backuprecoveryv1.DataLockConstraints) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Mode != nil {
		modelMap["mode"] = model.Mode
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsBackupAttemptToMap(model *backuprecoveryv1.BackupAttempt) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.AdmittedTimeUsecs != nil {
		modelMap["admitted_time_usecs"] = flex.IntValue(model.AdmittedTimeUsecs)
	}
	if model.PermitGrantTimeUsecs != nil {
		modelMap["permit_grant_time_usecs"] = flex.IntValue(model.PermitGrantTimeUsecs)
	}
	if model.QueueDurationUsecs != nil {
		modelMap["queue_duration_usecs"] = flex.IntValue(model.QueueDurationUsecs)
	}
	if model.SnapshotCreationTimeUsecs != nil {
		modelMap["snapshot_creation_time_usecs"] = flex.IntValue(model.SnapshotCreationTimeUsecs)
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionGroupsBackupDataStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsReplicationRunToMap(model *backuprecoveryv1.ReplicationRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := dataSourceIbmProtectionGroupsReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsReplicationTargetResultToMap(model *backuprecoveryv1.ReplicationTargetResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = model.ClusterName
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.QueuedTimeUsecs != nil {
		modelMap["queued_time_usecs"] = flex.IntValue(model.QueuedTimeUsecs)
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.IntValue(model.PercentageCompleted)
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionGroupsReplicationDataStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.IsManuallyDeleted != nil {
		modelMap["is_manually_deleted"] = model.IsManuallyDeleted
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	if model.ReplicationTaskID != nil {
		modelMap["replication_task_id"] = model.ReplicationTaskID
	}
	if model.EntriesChanged != nil {
		modelMap["entries_changed"] = flex.IntValue(model.EntriesChanged)
	}
	if model.IsInBound != nil {
		modelMap["is_in_bound"] = model.IsInBound
	}
	if model.DataLockConstraints != nil {
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	if model.MultiObjectReplication != nil {
		modelMap["multi_object_replication"] = model.MultiObjectReplication
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsReplicationDataStatsToMap(model *backuprecoveryv1.ReplicationDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.LogicalBytesTransferred != nil {
		modelMap["logical_bytes_transferred"] = flex.IntValue(model.LogicalBytesTransferred)
	}
	if model.PhysicalBytesTransferred != nil {
		modelMap["physical_bytes_transferred"] = flex.IntValue(model.PhysicalBytesTransferred)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsArchivalRunToMap(model *backuprecoveryv1.ArchivalRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := dataSourceIbmProtectionGroupsArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsArchivalTargetResultToMap(model *backuprecoveryv1.ArchivalTargetResult) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmProtectionGroupsArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.RunType != nil {
		modelMap["run_type"] = model.RunType
	}
	if model.IsSlaViolated != nil {
		modelMap["is_sla_violated"] = model.IsSlaViolated
	}
	if model.SnapshotID != nil {
		modelMap["snapshot_id"] = model.SnapshotID
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.QueuedTimeUsecs != nil {
		modelMap["queued_time_usecs"] = flex.IntValue(model.QueuedTimeUsecs)
	}
	if model.IsIncremental != nil {
		modelMap["is_incremental"] = model.IsIncremental
	}
	if model.IsForeverIncremental != nil {
		modelMap["is_forever_incremental"] = model.IsForeverIncremental
	}
	if model.IsCadArchive != nil {
		modelMap["is_cad_archive"] = model.IsCadArchive
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.StatsTaskID != nil {
		modelMap["stats_task_id"] = model.StatsTaskID
	}
	if model.IndexingTaskID != nil {
		modelMap["indexing_task_id"] = model.IndexingTaskID
	}
	if model.SuccessfulObjectsCount != nil {
		modelMap["successful_objects_count"] = flex.IntValue(model.SuccessfulObjectsCount)
	}
	if model.FailedObjectsCount != nil {
		modelMap["failed_objects_count"] = flex.IntValue(model.FailedObjectsCount)
	}
	if model.CancelledObjectsCount != nil {
		modelMap["cancelled_objects_count"] = flex.IntValue(model.CancelledObjectsCount)
	}
	if model.SuccessfulAppObjectsCount != nil {
		modelMap["successful_app_objects_count"] = flex.IntValue(model.SuccessfulAppObjectsCount)
	}
	if model.FailedAppObjectsCount != nil {
		modelMap["failed_app_objects_count"] = flex.IntValue(model.FailedAppObjectsCount)
	}
	if model.CancelledAppObjectsCount != nil {
		modelMap["cancelled_app_objects_count"] = flex.IntValue(model.CancelledAppObjectsCount)
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionGroupsArchivalDataStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.IsManuallyDeleted != nil {
		modelMap["is_manually_deleted"] = model.IsManuallyDeleted
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	if model.DataLockConstraints != nil {
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	if model.WormProperties != nil {
		wormPropertiesMap, err := dataSourceIbmProtectionGroupsWormPropertiesToMap(model.WormProperties)
		if err != nil {
			return modelMap, err
		}
		modelMap["worm_properties"] = []map[string]interface{}{wormPropertiesMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmProtectionGroupsOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmProtectionGroupsOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmProtectionGroupsOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupsArchivalDataStatsToMap(model *backuprecoveryv1.ArchivalDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.BytesRead != nil {
		modelMap["bytes_read"] = flex.IntValue(model.BytesRead)
	}
	if model.LogicalBytesTransferred != nil {
		modelMap["logical_bytes_transferred"] = flex.IntValue(model.LogicalBytesTransferred)
	}
	if model.PhysicalBytesTransferred != nil {
		modelMap["physical_bytes_transferred"] = flex.IntValue(model.PhysicalBytesTransferred)
	}
	if model.AvgLogicalTransferRateBps != nil {
		modelMap["avg_logical_transfer_rate_bps"] = flex.IntValue(model.AvgLogicalTransferRateBps)
	}
	if model.FileWalkDone != nil {
		modelMap["file_walk_done"] = model.FileWalkDone
	}
	if model.TotalFileCount != nil {
		modelMap["total_file_count"] = flex.IntValue(model.TotalFileCount)
	}
	if model.BackupFileCount != nil {
		modelMap["backup_file_count"] = flex.IntValue(model.BackupFileCount)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsWormPropertiesToMap(model *backuprecoveryv1.WormProperties) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsArchiveWormCompliant != nil {
		modelMap["is_archive_worm_compliant"] = model.IsArchiveWormCompliant
	}
	if model.WormNonComplianceReason != nil {
		modelMap["worm_non_compliance_reason"] = model.WormNonComplianceReason
	}
	if model.WormExpiryTimeUsecs != nil {
		modelMap["worm_expiry_time_usecs"] = flex.IntValue(model.WormExpiryTimeUsecs)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCloudSpinRunToMap(model *backuprecoveryv1.CloudSpinRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := dataSourceIbmProtectionGroupsCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCloudSpinTargetResultToMap(model *backuprecoveryv1.CloudSpinTargetResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
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
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionGroupsCloudSpinDataStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.IsManuallyDeleted != nil {
		modelMap["is_manually_deleted"] = model.IsManuallyDeleted
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	if model.CloudspinTaskID != nil {
		modelMap["cloudspin_task_id"] = model.CloudspinTaskID
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.DataLockConstraints != nil {
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCloudSpinDataStatsToMap(model *backuprecoveryv1.CloudSpinDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PhysicalBytesTransferred != nil {
		modelMap["physical_bytes_transferred"] = flex.IntValue(model.PhysicalBytesTransferred)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsBackupRunSummaryToMap(model *backuprecoveryv1.BackupRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RunType != nil {
		modelMap["run_type"] = model.RunType
	}
	if model.IsSlaViolated != nil {
		modelMap["is_sla_violated"] = model.IsSlaViolated
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
	if model.Messages != nil {
		modelMap["messages"] = model.Messages
	}
	if model.SuccessfulObjectsCount != nil {
		modelMap["successful_objects_count"] = flex.IntValue(model.SuccessfulObjectsCount)
	}
	if model.FailedObjectsCount != nil {
		modelMap["failed_objects_count"] = flex.IntValue(model.FailedObjectsCount)
	}
	if model.CancelledObjectsCount != nil {
		modelMap["cancelled_objects_count"] = flex.IntValue(model.CancelledObjectsCount)
	}
	if model.SuccessfulAppObjectsCount != nil {
		modelMap["successful_app_objects_count"] = flex.IntValue(model.SuccessfulAppObjectsCount)
	}
	if model.FailedAppObjectsCount != nil {
		modelMap["failed_app_objects_count"] = flex.IntValue(model.FailedAppObjectsCount)
	}
	if model.CancelledAppObjectsCount != nil {
		modelMap["cancelled_app_objects_count"] = flex.IntValue(model.CancelledAppObjectsCount)
	}
	if model.LocalSnapshotStats != nil {
		localSnapshotStatsMap, err := dataSourceIbmProtectionGroupsBackupDataStatsToMap(model.LocalSnapshotStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_snapshot_stats"] = []map[string]interface{}{localSnapshotStatsMap}
	}
	if model.IndexingTaskID != nil {
		modelMap["indexing_task_id"] = model.IndexingTaskID
	}
	if model.ProgressTaskID != nil {
		modelMap["progress_task_id"] = model.ProgressTaskID
	}
	if model.StatsTaskID != nil {
		modelMap["stats_task_id"] = model.StatsTaskID
	}
	if model.DataLock != nil {
		modelMap["data_lock"] = model.DataLock
	}
	if model.LocalTaskID != nil {
		modelMap["local_task_id"] = model.LocalTaskID
	}
	if model.DataLockConstraints != nil {
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupsDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsReplicationRunSummaryToMap(model *backuprecoveryv1.ReplicationRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := dataSourceIbmProtectionGroupsReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsArchivalRunSummaryToMap(model *backuprecoveryv1.ArchivalRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := dataSourceIbmProtectionGroupsArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCloudSpinRunSummaryToMap(model *backuprecoveryv1.CloudSpinRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := dataSourceIbmProtectionGroupsCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsMissingEntityParamsToMap(model *backuprecoveryv1.MissingEntityParams) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupsKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["protection_type"] = model.ProtectionType
	if model.VolumeProtectionTypeParams != nil {
		volumeProtectionTypeParamsMap, err := dataSourceIbmProtectionGroupsPhysicalVolumeProtectionGroupParamsToMap(model.VolumeProtectionTypeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["volume_protection_type_params"] = []map[string]interface{}{volumeProtectionTypeParamsMap}
	}
	if model.FileProtectionTypeParams != nil {
		fileProtectionTypeParamsMap, err := dataSourceIbmProtectionGroupsPhysicalFileProtectionGroupParamsToMap(model.FileProtectionTypeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["file_protection_type_params"] = []map[string]interface{}{fileProtectionTypeParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalVolumeProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalVolumeProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := dataSourceIbmProtectionGroupsPhysicalVolumeProtectionGroupObjectParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	if model.IndexingPolicy != nil {
		indexingPolicyMap, err := dataSourceIbmProtectionGroupsIndexingPolicyToMap(model.IndexingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["indexing_policy"] = []map[string]interface{}{indexingPolicyMap}
	}
	if model.PerformSourceSideDeduplication != nil {
		modelMap["perform_source_side_deduplication"] = model.PerformSourceSideDeduplication
	}
	if model.Quiesce != nil {
		modelMap["quiesce"] = model.Quiesce
	}
	if model.ContinueOnQuiesceFailure != nil {
		modelMap["continue_on_quiesce_failure"] = model.ContinueOnQuiesceFailure
	}
	if model.IncrementalBackupAfterRestart != nil {
		modelMap["incremental_backup_after_restart"] = model.IncrementalBackupAfterRestart
	}
	if model.PrePostScript != nil {
		prePostScriptMap, err := dataSourceIbmProtectionGroupsPrePostScriptParamsToMap(model.PrePostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["pre_post_script"] = []map[string]interface{}{prePostScriptMap}
	}
	if model.DedupExclusionSourceIds != nil {
		modelMap["dedup_exclusion_source_ids"] = model.DedupExclusionSourceIds
	}
	if model.ExcludedVssWriters != nil {
		modelMap["excluded_vss_writers"] = model.ExcludedVssWriters
	}
	if model.CobmrBackup != nil {
		modelMap["cobmr_backup"] = model.CobmrBackup
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalVolumeProtectionGroupObjectParamsToMap(model *backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.VolumeGuids != nil {
		modelMap["volume_guids"] = model.VolumeGuids
	}
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsIndexingPolicyToMap(model *backuprecoveryv1.IndexingPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["enable_indexing"] = model.EnableIndexing
	if model.IncludePaths != nil {
		modelMap["include_paths"] = model.IncludePaths
	}
	if model.ExcludePaths != nil {
		modelMap["exclude_paths"] = model.ExcludePaths
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPrePostScriptParamsToMap(model *backuprecoveryv1.PrePostScriptParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PreScript != nil {
		preScriptMap, err := dataSourceIbmProtectionGroupsCommonPreBackupScriptParamsToMap(model.PreScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["pre_script"] = []map[string]interface{}{preScriptMap}
	}
	if model.PostScript != nil {
		postScriptMap, err := dataSourceIbmProtectionGroupsCommonPostBackupScriptParamsToMap(model.PostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["post_script"] = []map[string]interface{}{postScriptMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCommonPreBackupScriptParamsToMap(model *backuprecoveryv1.CommonPreBackupScriptParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["path"] = model.Path
	if model.Params != nil {
		modelMap["params"] = model.Params
	}
	if model.TimeoutSecs != nil {
		modelMap["timeout_secs"] = flex.IntValue(model.TimeoutSecs)
	}
	if model.IsActive != nil {
		modelMap["is_active"] = model.IsActive
	}
	if model.ContinueOnError != nil {
		modelMap["continue_on_error"] = model.ContinueOnError
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCommonPostBackupScriptParamsToMap(model *backuprecoveryv1.CommonPostBackupScriptParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["path"] = model.Path
	if model.Params != nil {
		modelMap["params"] = model.Params
	}
	if model.TimeoutSecs != nil {
		modelMap["timeout_secs"] = flex.IntValue(model.TimeoutSecs)
	}
	if model.IsActive != nil {
		modelMap["is_active"] = model.IsActive
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalFileProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalFileProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := dataSourceIbmProtectionGroupsPhysicalFileProtectionGroupObjectParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	if model.IndexingPolicy != nil {
		indexingPolicyMap, err := dataSourceIbmProtectionGroupsIndexingPolicyToMap(model.IndexingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["indexing_policy"] = []map[string]interface{}{indexingPolicyMap}
	}
	if model.PerformSourceSideDeduplication != nil {
		modelMap["perform_source_side_deduplication"] = model.PerformSourceSideDeduplication
	}
	if model.PerformBrickBasedDeduplication != nil {
		modelMap["perform_brick_based_deduplication"] = model.PerformBrickBasedDeduplication
	}
	if model.TaskTimeouts != nil {
		taskTimeouts := []map[string]interface{}{}
		for _, taskTimeoutsItem := range model.TaskTimeouts {
			taskTimeoutsItemMap, err := dataSourceIbmProtectionGroupsCancellationTimeoutParamsToMap(&taskTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			taskTimeouts = append(taskTimeouts, taskTimeoutsItemMap)
		}
		modelMap["task_timeouts"] = taskTimeouts
	}
	if model.Quiesce != nil {
		modelMap["quiesce"] = model.Quiesce
	}
	if model.ContinueOnQuiesceFailure != nil {
		modelMap["continue_on_quiesce_failure"] = model.ContinueOnQuiesceFailure
	}
	if model.CobmrBackup != nil {
		modelMap["cobmr_backup"] = model.CobmrBackup
	}
	if model.PrePostScript != nil {
		prePostScriptMap, err := dataSourceIbmProtectionGroupsPrePostScriptParamsToMap(model.PrePostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["pre_post_script"] = []map[string]interface{}{prePostScriptMap}
	}
	if model.DedupExclusionSourceIds != nil {
		modelMap["dedup_exclusion_source_ids"] = model.DedupExclusionSourceIds
	}
	if model.GlobalExcludePaths != nil {
		modelMap["global_exclude_paths"] = model.GlobalExcludePaths
	}
	if model.GlobalExcludeFS != nil {
		modelMap["global_exclude_fs"] = model.GlobalExcludeFS
	}
	if model.IgnorableErrors != nil {
		modelMap["ignorable_errors"] = model.IgnorableErrors
	}
	if model.AllowParallelRuns != nil {
		modelMap["allow_parallel_runs"] = model.AllowParallelRuns
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalFileProtectionGroupObjectParamsToMap(model *backuprecoveryv1.PhysicalFileProtectionGroupObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.FilePaths != nil {
		filePaths := []map[string]interface{}{}
		for _, filePathsItem := range model.FilePaths {
			filePathsItemMap, err := dataSourceIbmProtectionGroupsPhysicalFileBackupPathParamsToMap(&filePathsItem)
			if err != nil {
				return modelMap, err
			}
			filePaths = append(filePaths, filePathsItemMap)
		}
		modelMap["file_paths"] = filePaths
	}
	if model.UsesPathLevelSkipNestedVolumeSetting != nil {
		modelMap["uses_path_level_skip_nested_volume_setting"] = model.UsesPathLevelSkipNestedVolumeSetting
	}
	if model.NestedVolumeTypesToSkip != nil {
		modelMap["nested_volume_types_to_skip"] = model.NestedVolumeTypesToSkip
	}
	if model.FollowNasSymlinkTarget != nil {
		modelMap["follow_nas_symlink_target"] = model.FollowNasSymlinkTarget
	}
	if model.MetadataFilePath != nil {
		modelMap["metadata_file_path"] = model.MetadataFilePath
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsPhysicalFileBackupPathParamsToMap(model *backuprecoveryv1.PhysicalFileBackupPathParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["included_path"] = model.IncludedPath
	if model.ExcludedPaths != nil {
		modelMap["excluded_paths"] = model.ExcludedPaths
	}
	if model.SkipNestedVolumes != nil {
		modelMap["skip_nested_volumes"] = model.SkipNestedVolumes
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsCancellationTimeoutParamsToMap(model *backuprecoveryv1.CancellationTimeoutParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TimeoutMins != nil {
		modelMap["timeout_mins"] = flex.IntValue(model.TimeoutMins)
	}
	if model.BackupType != nil {
		modelMap["backup_type"] = model.BackupType
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleProtectionGroupParamsToMap(model *backuprecoveryv1.OracleProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := dataSourceIbmProtectionGroupsOracleProtectionGroupObjectParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	if model.PersistMountpoints != nil {
		modelMap["persist_mountpoints"] = model.PersistMountpoints
	}
	if model.VlanParams != nil {
		vlanParamsMap, err := dataSourceIbmProtectionGroupsVlanParamsToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	if model.PrePostScript != nil {
		prePostScriptMap, err := dataSourceIbmProtectionGroupsPrePostScriptParamsToMap(model.PrePostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["pre_post_script"] = []map[string]interface{}{prePostScriptMap}
	}
	if model.LogAutoKillTimeoutSecs != nil {
		modelMap["log_auto_kill_timeout_secs"] = flex.IntValue(model.LogAutoKillTimeoutSecs)
	}
	if model.IncrAutoKillTimeoutSecs != nil {
		modelMap["incr_auto_kill_timeout_secs"] = flex.IntValue(model.IncrAutoKillTimeoutSecs)
	}
	if model.FullAutoKillTimeoutSecs != nil {
		modelMap["full_auto_kill_timeout_secs"] = flex.IntValue(model.FullAutoKillTimeoutSecs)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleProtectionGroupObjectParamsToMap(model *backuprecoveryv1.OracleProtectionGroupObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_id"] = flex.IntValue(model.SourceID)
	if model.SourceName != nil {
		modelMap["source_name"] = model.SourceName
	}
	if model.DbParams != nil {
		dbParams := []map[string]interface{}{}
		for _, dbParamsItem := range model.DbParams {
			dbParamsItemMap, err := dataSourceIbmProtectionGroupsOracleProtectionGroupDbParamsToMap(&dbParamsItem)
			if err != nil {
				return modelMap, err
			}
			dbParams = append(dbParams, dbParamsItemMap)
		}
		modelMap["db_params"] = dbParams
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleProtectionGroupDbParamsToMap(model *backuprecoveryv1.OracleProtectionGroupDbParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = flex.IntValue(model.DatabaseID)
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	if model.DbChannels != nil {
		dbChannels := []map[string]interface{}{}
		for _, dbChannelsItem := range model.DbChannels {
			dbChannelsItemMap, err := dataSourceIbmProtectionGroupsOracleDbChannelToMap(&dbChannelsItem)
			if err != nil {
				return modelMap, err
			}
			dbChannels = append(dbChannels, dbChannelsItemMap)
		}
		modelMap["db_channels"] = dbChannels
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleDbChannelToMap(model *backuprecoveryv1.OracleDbChannel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchiveLogRetentionDays != nil {
		modelMap["archive_log_retention_days"] = flex.IntValue(model.ArchiveLogRetentionDays)
	}
	if model.ArchiveLogRetentionHours != nil {
		modelMap["archive_log_retention_hours"] = flex.IntValue(model.ArchiveLogRetentionHours)
	}
	if model.Credentials != nil {
		credentialsMap, err := dataSourceIbmProtectionGroupsCredentialsToMap(model.Credentials)
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
			databaseNodeListItemMap, err := dataSourceIbmProtectionGroupsOracleDatabaseHostToMap(&databaseNodeListItem)
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

func dataSourceIbmProtectionGroupsCredentialsToMap(model *backuprecoveryv1.Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleDatabaseHostToMap(model *backuprecoveryv1.OracleDatabaseHost) (map[string]interface{}, error) {
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
		sbtHostParamsMap, err := dataSourceIbmProtectionGroupsOracleSbtHostParamsToMap(model.SbtHostParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sbt_host_params"] = []map[string]interface{}{sbtHostParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleSbtHostParamsToMap(model *backuprecoveryv1.OracleSbtHostParams) (map[string]interface{}, error) {
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
			vlanInfoListItemMap, err := dataSourceIbmProtectionGroupsOracleVlanInfoToMap(&vlanInfoListItem)
			if err != nil {
				return modelMap, err
			}
			vlanInfoList = append(vlanInfoList, vlanInfoListItemMap)
		}
		modelMap["vlan_info_list"] = vlanInfoList
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupsOracleVlanInfoToMap(model *backuprecoveryv1.OracleVlanInfo) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupsVlanParamsToMap(model *backuprecoveryv1.VlanParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VlanID != nil {
		modelMap["vlan_id"] = flex.IntValue(model.VlanID)
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = model.InterfaceName
	}
	return modelMap, nil
}
