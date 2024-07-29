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

func DataSourceIbmProtectionRunProgress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionRunProgressRead,

		Schema: map[string]*schema.Schema{
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique run id of the Protection Run.",
			},
			"objects": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
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
			"include_finished_tasks": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to return finished tasks. By default only active tasks are returned.",
			},
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds).",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds).",
			},
			"max_tasks_num": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the maximum number of tasks to return.",
			},
			"exclude_object_details": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to return objects. By default all the task tree are returned.",
			},
			"include_event_logs": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to include event logs.",
			},
			"max_log_level": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true.",
			},
			"run_task_path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format.",
			},
			"object_task_paths": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"local_run": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the progress of a local backup run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
						},
						"percentage_completed": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Specifies the current completed percentage of the progress task.",
						},
						"start_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"events": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the event log created for progress Task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the log message describing the current event.",
									},
									"occured_at_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
									},
								},
							},
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies progress for objects.",
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
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"percentage_completed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Specifies the current completed percentage of the progress task.",
									},
									"start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"events": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the event log created for progress Task.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the log message describing the current event.",
												},
												"occured_at_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
												},
											},
										},
									},
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
												},
												"percentage_completed": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Specifies the current completed percentage of the progress task.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"events": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the event log created for progress Task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the log message describing the current event.",
															},
															"occured_at_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
															},
														},
													},
												},
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"archival_run": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Progress for the archival run.",
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
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
						},
						"percentage_completed": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Specifies the current completed percentage of the progress task.",
						},
						"start_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"events": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the event log created for progress Task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the log message describing the current event.",
									},
									"occured_at_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
									},
								},
							},
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies progress for objects.",
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
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"percentage_completed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Specifies the current completed percentage of the progress task.",
									},
									"start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"events": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the event log created for progress Task.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the log message describing the current event.",
												},
												"occured_at_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
												},
											},
										},
									},
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
												},
												"percentage_completed": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Specifies the current completed percentage of the progress task.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"events": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the event log created for progress Task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the log message describing the current event.",
															},
															"occured_at_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
															},
														},
													},
												},
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"replication_run": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Progress for the replication run.",
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
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
						},
						"percentage_completed": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Specifies the current completed percentage of the progress task.",
						},
						"start_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
						},
						"events": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the event log created for progress Task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the log message describing the current event.",
									},
									"occured_at_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
									},
								},
							},
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies progress for objects.",
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
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"percentage_completed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Specifies the current completed percentage of the progress task.",
									},
									"start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
									},
									"events": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the event log created for progress Task.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the log message describing the current event.",
												},
												"occured_at_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
												},
											},
										},
									},
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
												},
												"percentage_completed": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Specifies the current completed percentage of the progress task.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
												},
												"events": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the event log created for progress Task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the log message describing the current event.",
															},
															"occured_at_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).",
															},
														},
													},
												},
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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

func dataSourceIbmProtectionRunProgressRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionRunProgressOptions := &backuprecoveryv1.GetProtectionRunProgressOptions{}

	getProtectionRunProgressOptions.SetRunID(d.Get("run_id").(string))
	if _, ok := d.GetOk("objects"); ok {
		var objects []int64
		for _, v := range d.Get("objects").([]interface{}) {
			objectsItem := int64(v.(int))
			objects = append(objects, objectsItem)
		}
		getProtectionRunProgressOptions.SetObjects(objects)
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getProtectionRunProgressOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getProtectionRunProgressOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("include_finished_tasks"); ok {
		getProtectionRunProgressOptions.SetIncludeFinishedTasks(d.Get("include_finished_tasks").(bool))
	}
	if _, ok := d.GetOk("start_time_usecs"); ok {
		getProtectionRunProgressOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getProtectionRunProgressOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("max_tasks_num"); ok {
		getProtectionRunProgressOptions.SetMaxTasksNum(int64(d.Get("max_tasks_num").(int)))
	}
	if _, ok := d.GetOk("exclude_object_details"); ok {
		getProtectionRunProgressOptions.SetExcludeObjectDetails(d.Get("exclude_object_details").(bool))
	}
	if _, ok := d.GetOk("include_event_logs"); ok {
		getProtectionRunProgressOptions.SetIncludeEventLogs(d.Get("include_event_logs").(bool))
	}
	if _, ok := d.GetOk("max_log_level"); ok {
		getProtectionRunProgressOptions.SetMaxLogLevel(int64(d.Get("max_log_level").(int)))
	}
	if _, ok := d.GetOk("run_task_path"); ok {
		getProtectionRunProgressOptions.SetRunTaskPath(d.Get("run_task_path").(string))
	}
	if _, ok := d.GetOk("object_task_paths"); ok {
		var objectTaskPaths []string
		for _, v := range d.Get("object_task_paths").([]interface{}) {
			objectTaskPathsItem := v.(string)
			objectTaskPaths = append(objectTaskPaths, objectTaskPathsItem)
		}
		getProtectionRunProgressOptions.SetObjectTaskPaths(objectTaskPaths)
	}

	getProtectionRunProgress, response, err := backupRecoveryClient.GetProtectionRunProgressWithContext(context, getProtectionRunProgressOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionRunProgressWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionRunProgressWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmProtectionRunProgressID(d))

	localRun := []map[string]interface{}{}
	if getProtectionRunProgress.LocalRun != nil {
		modelMap, err := dataSourceIbmProtectionRunProgressBackupRunProgressInfoToMap(getProtectionRunProgress.LocalRun)
		if err != nil {
			return diag.FromErr(err)
		}
		localRun = append(localRun, modelMap)
	}
	if err = d.Set("local_run", localRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting local_run %s", err))
	}

	archivalRun := []map[string]interface{}{}
	if getProtectionRunProgress.ArchivalRun != nil {
		for _, modelItem := range getProtectionRunProgress.ArchivalRun {
			modelMap, err := dataSourceIbmProtectionRunProgressArchivalTargetProgressInfoToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			archivalRun = append(archivalRun, modelMap)
		}
	}
	if err = d.Set("archival_run", archivalRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting archival_run %s", err))
	}

	replicationRun := []map[string]interface{}{}
	if getProtectionRunProgress.ReplicationRun != nil {
		for _, modelItem := range getProtectionRunProgress.ReplicationRun {
			modelMap, err := dataSourceIbmProtectionRunProgressReplicationTargetProgressInfoToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			replicationRun = append(replicationRun, modelMap)
		}
	}
	if err = d.Set("replication_run", replicationRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting replication_run %s", err))
	}

	return nil
}

// dataSourceIbmProtectionRunProgressID returns a reasonable ID for the list.
func dataSourceIbmProtectionRunProgressID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmProtectionRunProgressBackupRunProgressInfoToMap(model *backuprecoveryv1.BackupRunProgressInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskEventToMap(&eventsItem)
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionRunProgressObjectProgressInfoToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressProgressTaskEventToMap(model *backuprecoveryv1.ProgressTaskEvent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.OccuredAtUsecs != nil {
		modelMap["occured_at_usecs"] = flex.IntValue(model.OccuredAtUsecs)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressProgressStatsToMap(model *backuprecoveryv1.ProgressStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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

func dataSourceIbmProtectionRunProgressObjectProgressInfoToMap(model *backuprecoveryv1.ObjectProgressInfo) (map[string]interface{}, error) {
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
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskEventToMap(&eventsItem)
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskInfoToMap(&failedAttemptsItem)
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressProgressTaskInfoToMap(model *backuprecoveryv1.ProgressTaskInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskEventToMap(&eventsItem)
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressArchivalTargetProgressInfoToMap(model *backuprecoveryv1.ArchivalTargetProgressInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmProtectionRunProgressArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskEventToMap(&eventsItem)
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionRunProgressObjectProgressInfoToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmProtectionRunProgressOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmProtectionRunProgressOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmProtectionRunProgressOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmProtectionRunProgressOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionRunProgressReplicationTargetProgressInfoToMap(model *backuprecoveryv1.ReplicationTargetProgressInfo) (map[string]interface{}, error) {
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
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := dataSourceIbmProtectionRunProgressProgressTaskEventToMap(&eventsItem)
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.Stats != nil {
		statsMap, err := dataSourceIbmProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionRunProgressObjectProgressInfoToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}
