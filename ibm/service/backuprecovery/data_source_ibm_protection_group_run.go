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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmProtectionGroupRun() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionGroupRunRead,

		Schema: map[string]*schema.Schema{
			"protection_group_run_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique id of the Protection Group.",
			},
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique run id of the Protection Group run.",
			},
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which the run is to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default.",
			},
			"include_object_details": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned.",
			},
			"use_cached_data": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.",
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
	}
}

func dataSourceIbmProtectionGroupRunRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionGroupRunOptions := &backuprecoveryv1.GetProtectionGroupRunOptions{}

	getProtectionGroupRunOptions.SetID(d.Get("protection_group_run_id").(string))
	getProtectionGroupRunOptions.SetRunID(d.Get("run_id").(string))
	if _, ok := d.GetOk("request_initiator_type"); ok {
		getProtectionGroupRunOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getProtectionGroupRunOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getProtectionGroupRunOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("include_object_details"); ok {
		getProtectionGroupRunOptions.SetIncludeObjectDetails(d.Get("include_object_details").(bool))
	}
	if _, ok := d.GetOk("use_cached_data"); ok {
		getProtectionGroupRunOptions.SetUseCachedData(d.Get("use_cached_data").(bool))
	}

	protectionGroupRun, response, err := backupRecoveryClient.GetProtectionGroupRunWithContext(context, getProtectionGroupRunOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionGroupRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionGroupRunWithContext failed %s\n%s", err, response))
	}

	d.SetId(*protectionGroupRun.ID)

	if err = d.Set("protection_group_instance_id", flex.IntValue(protectionGroupRun.ProtectionGroupInstanceID)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting protection_group_instance_id: %s", err))
	}

	if err = d.Set("protection_group_id", protectionGroupRun.ProtectionGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting protection_group_id: %s", err))
	}

	if err = d.Set("is_replication_run", protectionGroupRun.IsReplicationRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting is_replication_run: %s", err))
	}

	originClusterIdentifier := []map[string]interface{}{}
	if protectionGroupRun.OriginClusterIdentifier != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunClusterIdentifierToMap(protectionGroupRun.OriginClusterIdentifier)
		if err != nil {
			return diag.FromErr(err)
		}
		originClusterIdentifier = append(originClusterIdentifier, modelMap)
	}
	if err = d.Set("origin_cluster_identifier", originClusterIdentifier); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting origin_cluster_identifier %s", err))
	}

	if err = d.Set("origin_protection_group_id", protectionGroupRun.OriginProtectionGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting origin_protection_group_id: %s", err))
	}

	if err = d.Set("protection_group_name", protectionGroupRun.ProtectionGroupName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting protection_group_name: %s", err))
	}

	if err = d.Set("is_local_snapshots_deleted", protectionGroupRun.IsLocalSnapshotsDeleted); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting is_local_snapshots_deleted: %s", err))
	}

	objects := []map[string]interface{}{}
	if protectionGroupRun.Objects != nil {
		for _, modelItem := range protectionGroupRun.Objects {
			modelMap, err := dataSourceIbmProtectionGroupRunObjectRunResultToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			objects = append(objects, modelMap)
		}
	}
	if err = d.Set("objects", objects); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting objects %s", err))
	}

	localBackupInfo := []map[string]interface{}{}
	if protectionGroupRun.LocalBackupInfo != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunBackupRunSummaryToMap(protectionGroupRun.LocalBackupInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		localBackupInfo = append(localBackupInfo, modelMap)
	}
	if err = d.Set("local_backup_info", localBackupInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting local_backup_info %s", err))
	}

	originalBackupInfo := []map[string]interface{}{}
	if protectionGroupRun.OriginalBackupInfo != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunBackupRunSummaryToMap(protectionGroupRun.OriginalBackupInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		originalBackupInfo = append(originalBackupInfo, modelMap)
	}
	if err = d.Set("original_backup_info", originalBackupInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting original_backup_info %s", err))
	}

	replicationInfo := []map[string]interface{}{}
	if protectionGroupRun.ReplicationInfo != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunReplicationRunSummaryToMap(protectionGroupRun.ReplicationInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		replicationInfo = append(replicationInfo, modelMap)
	}
	if err = d.Set("replication_info", replicationInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting replication_info %s", err))
	}

	archivalInfo := []map[string]interface{}{}
	if protectionGroupRun.ArchivalInfo != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunArchivalRunSummaryToMap(protectionGroupRun.ArchivalInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		archivalInfo = append(archivalInfo, modelMap)
	}
	if err = d.Set("archival_info", archivalInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting archival_info %s", err))
	}

	cloudSpinInfo := []map[string]interface{}{}
	if protectionGroupRun.CloudSpinInfo != nil {
		modelMap, err := dataSourceIbmProtectionGroupRunCloudSpinRunSummaryToMap(protectionGroupRun.CloudSpinInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		cloudSpinInfo = append(cloudSpinInfo, modelMap)
	}
	if err = d.Set("cloud_spin_info", cloudSpinInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cloud_spin_info %s", err))
	}

	if err = d.Set("on_legal_hold", protectionGroupRun.OnLegalHold); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting on_legal_hold: %s", err))
	}

	permissions := []map[string]interface{}{}
	if protectionGroupRun.Permissions != nil {
		for _, modelItem := range protectionGroupRun.Permissions {
			modelMap, err := dataSourceIbmProtectionGroupRunTenantToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			permissions = append(permissions, modelMap)
		}
	}
	if err = d.Set("permissions", permissions); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting permissions %s", err))
	}

	if err = d.Set("is_cloud_archival_direct", protectionGroupRun.IsCloudArchivalDirect); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting is_cloud_archival_direct: %s", err))
	}

	if err = d.Set("has_local_snapshot", protectionGroupRun.HasLocalSnapshot); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting has_local_snapshot: %s", err))
	}

	if err = d.Set("environment", protectionGroupRun.Environment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
	}

	if err = d.Set("externally_triggered_backup_tag", protectionGroupRun.ExternallyTriggeredBackupTag); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting externally_triggered_backup_tag: %s", err))
	}

	return nil
}

func dataSourceIbmProtectionGroupRunClusterIdentifierToMap(model *backuprecoveryv1.ClusterIdentifier) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunObjectRunResultToMap(model *backuprecoveryv1.ObjectRunResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Object != nil {
		objectMap, err := dataSourceIbmProtectionGroupRunObjectSummaryToMap(model.Object)
		if err != nil {
			return modelMap, err
		}
		modelMap["object"] = []map[string]interface{}{objectMap}
	}
	if model.LocalSnapshotInfo != nil {
		localSnapshotInfoMap, err := dataSourceIbmProtectionGroupRunBackupRunToMap(model.LocalSnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_snapshot_info"] = []map[string]interface{}{localSnapshotInfoMap}
	}
	if model.OriginalBackupInfo != nil {
		originalBackupInfoMap, err := dataSourceIbmProtectionGroupRunBackupRunToMap(model.OriginalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_backup_info"] = []map[string]interface{}{originalBackupInfoMap}
	}
	if model.ReplicationInfo != nil {
		replicationInfoMap, err := dataSourceIbmProtectionGroupRunReplicationRunToMap(model.ReplicationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["replication_info"] = []map[string]interface{}{replicationInfoMap}
	}
	if model.ArchivalInfo != nil {
		archivalInfoMap, err := dataSourceIbmProtectionGroupRunArchivalRunToMap(model.ArchivalInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_info"] = []map[string]interface{}{archivalInfoMap}
	}
	if model.CloudSpinInfo != nil {
		cloudSpinInfoMap, err := dataSourceIbmProtectionGroupRunCloudSpinRunToMap(model.CloudSpinInfo)
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

func dataSourceIbmProtectionGroupRunObjectSummaryToMap(model *backuprecoveryv1.ObjectSummary) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunBackupRunToMap(model *backuprecoveryv1.BackupRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SnapshotInfo != nil {
		snapshotInfoMap, err := dataSourceIbmProtectionGroupRunSnapshotInfoToMap(model.SnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["snapshot_info"] = []map[string]interface{}{snapshotInfoMap}
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := dataSourceIbmProtectionGroupRunBackupAttemptToMap(&failedAttemptsItem)
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunSnapshotInfoToMap(model *backuprecoveryv1.SnapshotInfo) (map[string]interface{}, error) {
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
		statsMap, err := dataSourceIbmProtectionGroupRunBackupDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunBackupDataStatsToMap(model *backuprecoveryv1.BackupDataStats) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model *backuprecoveryv1.DataLockConstraints) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Mode != nil {
		modelMap["mode"] = model.Mode
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunBackupAttemptToMap(model *backuprecoveryv1.BackupAttempt) (map[string]interface{}, error) {
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
		statsMap, err := dataSourceIbmProtectionGroupRunBackupDataStatsToMap(model.Stats)
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

func dataSourceIbmProtectionGroupRunReplicationRunToMap(model *backuprecoveryv1.ReplicationRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunReplicationTargetResultToMap(model *backuprecoveryv1.ReplicationTargetResult) (map[string]interface{}, error) {
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
		statsMap, err := dataSourceIbmProtectionGroupRunReplicationDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model.DataLockConstraints)
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

func dataSourceIbmProtectionGroupRunReplicationDataStatsToMap(model *backuprecoveryv1.ReplicationDataStats) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunArchivalRunToMap(model *backuprecoveryv1.ArchivalRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunArchivalTargetResultToMap(model *backuprecoveryv1.ArchivalTargetResult) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmProtectionGroupRunArchivalTargetTierInfoToMap(model.TierSettings)
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
		statsMap, err := dataSourceIbmProtectionGroupRunArchivalDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	if model.WormProperties != nil {
		wormPropertiesMap, err := dataSourceIbmProtectionGroupRunWormPropertiesToMap(model.WormProperties)
		if err != nil {
			return modelMap, err
		}
		modelMap["worm_properties"] = []map[string]interface{}{wormPropertiesMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmProtectionGroupRunOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmProtectionGroupRunOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmProtectionGroupRunOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunArchivalDataStatsToMap(model *backuprecoveryv1.ArchivalDataStats) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunWormPropertiesToMap(model *backuprecoveryv1.WormProperties) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionGroupRunCloudSpinRunToMap(model *backuprecoveryv1.CloudSpinRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunCloudSpinTargetResultToMap(model *backuprecoveryv1.CloudSpinTargetResult) (map[string]interface{}, error) {
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
		statsMap, err := dataSourceIbmProtectionGroupRunCloudSpinDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model.DataLockConstraints)
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

func dataSourceIbmProtectionGroupRunCloudSpinDataStatsToMap(model *backuprecoveryv1.CloudSpinDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PhysicalBytesTransferred != nil {
		modelMap["physical_bytes_transferred"] = flex.IntValue(model.PhysicalBytesTransferred)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunBackupRunSummaryToMap(model *backuprecoveryv1.BackupRunSummary) (map[string]interface{}, error) {
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
		localSnapshotStatsMap, err := dataSourceIbmProtectionGroupRunBackupDataStatsToMap(model.LocalSnapshotStats)
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
		dataLockConstraintsMap, err := dataSourceIbmProtectionGroupRunDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunReplicationRunSummaryToMap(model *backuprecoveryv1.ReplicationRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunArchivalRunSummaryToMap(model *backuprecoveryv1.ArchivalRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunCloudSpinRunSummaryToMap(model *backuprecoveryv1.CloudSpinRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := dataSourceIbmProtectionGroupRunCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func dataSourceIbmProtectionGroupRunTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}
