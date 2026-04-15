// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryProtectionRunProgress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryProtectionRunProgressRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the tenant accessing the cluster.",
			},
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
									"aws_tiering": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies aws tiers.",
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
										Computed:    true,
										Description: "Specifies Azure tiers.",
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
										Computed:    true,
										Description: "Specifies the cloud platform to enable tiering.",
									},
									"google_tiering": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies Google tiers.",
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
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
								},
							},
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
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
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_file_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
											},
										},
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backup_file_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
														},
													},
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
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
			"local_run": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the progress of a local backup run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
								},
							},
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
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
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_file_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
											},
										},
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backup_file_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
														},
													},
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
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
						"aws_target_config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the configuration for adding AWS as repilcation target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the AWS Replication target.",
									},
									"region": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.",
									},
									"region_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the source id of the AWS protection source registered on IBM cluster.",
									},
								},
							},
						},
						"azure_target_config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the configuration for adding Azure as replication target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the Azure Replication target.",
									},
									"resource_group": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies id of the Azure resource group used to filter regions in UI.",
									},
									"resource_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the Azure resource group used to filter regions in UI.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the source id of the Azure protection source registered on IBM cluster.",
									},
									"storage_account": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies id of the storage account of Azure replication target which will contain storage container.",
									},
									"storage_account_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the storage account of Azure replication target which will contain storage container.",
									},
									"storage_container": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies id of the storage container of Azure Replication target.",
									},
									"storage_container_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the storage container of Azure Replication target.",
									},
									"storage_resource_group": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies id of the storage resource group of Azure Replication target.",
									},
									"storage_resource_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the storage resource group of Azure Replication target.",
									},
								},
							},
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"expected_remaining_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats within progress.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
								},
							},
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the current status of the progress task.",
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
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"expected_remaining_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats within progress.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_file_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
											},
										},
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the progress task.",
									},
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies progress for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"expected_remaining_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).",
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
												"stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats within progress.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backup_file_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.",
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
														},
													},
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the current status of the progress task.",
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

func dataSourceIbmBackupRecoveryProtectionRunProgressRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_protection_run_progress", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProtectionRunProgressOptions := &backuprecoveryv1.GetProtectionRunProgressOptions{}

	getProtectionRunProgressOptions.XIBMTenantID = (d.Get("x_ibm_tenant_id").(*string))
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

	getProtectionRunProgressBody, _, err := backupRecoveryClient.GetProtectionRunProgressWithContext(context, getProtectionRunProgressOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProtectionRunProgressWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_protection_run_progress", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryProtectionRunProgressID(d))

	if !core.IsNil(getProtectionRunProgressBody.ArchivalRun) {
		archivalRun := []map[string]interface{}{}
		for _, archivalRunItem := range getProtectionRunProgressBody.ArchivalRun {
			archivalRunItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetProgressInfoToMap(&archivalRunItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_protection_run_progress", "read", "archival_run-to-map").GetDiag()
			}
			archivalRun = append(archivalRun, archivalRunItemMap)
		}
		if err = d.Set("archival_run", archivalRun); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting archival_run: %s", err), "(Data) ibm_backup_recovery_protection_run_progress", "read", "set-archival_run").GetDiag()
		}
	}

	if !core.IsNil(getProtectionRunProgressBody.LocalRun) {
		localRun := []map[string]interface{}{}
		localRunMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressBackupRunProgressInfoToMap(getProtectionRunProgressBody.LocalRun)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_protection_run_progress", "read", "local_run-to-map").GetDiag()
		}
		localRun = append(localRun, localRunMap)
		if err = d.Set("local_run", localRun); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting local_run: %s", err), "(Data) ibm_backup_recovery_protection_run_progress", "read", "set-local_run").GetDiag()
		}
	}

	if !core.IsNil(getProtectionRunProgressBody.ReplicationRun) {
		replicationRun := []map[string]interface{}{}
		for _, replicationRunItem := range getProtectionRunProgressBody.ReplicationRun {
			replicationRunItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressReplicationTargetProgressInfoToMap(&replicationRunItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_protection_run_progress", "read", "replication_run-to-map").GetDiag()
			}
			replicationRun = append(replicationRun, replicationRunItemMap)
		}
		if err = d.Set("replication_run", replicationRun); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting replication_run: %s", err), "(Data) ibm_backup_recovery_protection_run_progress", "read", "set-replication_run").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryProtectionRunProgressID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryProtectionRunProgressID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetProgressInfoToMap(model *backuprecoveryv1.ArchivalTargetProgressInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(&eventsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(&objectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AwsTiering != nil {
		awsTieringMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAWSTiersToMap(model.AwsTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["aws_tiering"] = []map[string]interface{}{awsTieringMap}
	}
	if model.AzureTiering != nil {
		azureTieringMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAzureTiersToMap(model.AzureTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["azure_tiering"] = []map[string]interface{}{azureTieringMap}
	}
	if model.CloudPlatform != nil {
		modelMap["cloud_platform"] = *model.CloudPlatform
	}
	if model.GoogleTiering != nil {
		googleTieringMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTiersToMap(model.GoogleTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["google_tiering"] = []map[string]interface{}{googleTieringMap}
	}
	if model.OracleTiering != nil {
		oracleTieringMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressOracleTiersToMap(model.OracleTiering)
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

func DataSourceIbmBackupRecoveryProtectionRunProgressAWSTiersToMap(model *backuprecoveryv1.AWSTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAWSTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressAWSTierToMap(model *backuprecoveryv1.AWSTier) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryProtectionRunProgressAzureTiersToMap(model *backuprecoveryv1.AzureTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Tiers != nil {
		tiers := []map[string]interface{}{}
		for _, tiersItem := range model.Tiers {
			tiersItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAzureTierToMap(&tiersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			tiers = append(tiers, tiersItemMap)
		}
		modelMap["tiers"] = tiers
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressAzureTierToMap(model *backuprecoveryv1.AzureTier) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTiersToMap(model *backuprecoveryv1.GoogleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTierToMap(model *backuprecoveryv1.GoogleTier) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryProtectionRunProgressOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressOracleTierToMap(&tiersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(model *backuprecoveryv1.ProgressTaskEvent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.OccuredAtUsecs != nil {
		modelMap["occured_at_usecs"] = flex.IntValue(model.OccuredAtUsecs)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model *backuprecoveryv1.ProgressStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupFileCount != nil {
		modelMap["backup_file_count"] = flex.IntValue(model.BackupFileCount)
	}
	if model.FileWalkDone != nil {
		modelMap["file_walk_done"] = *model.FileWalkDone
	}
	if model.TotalFileCount != nil {
		modelMap["total_file_count"] = flex.IntValue(model.TotalFileCount)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(model *backuprecoveryv1.ObjectProgressInfo) (map[string]interface{}, error) {
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
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(&eventsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskInfoToMap(&failedAttemptsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskInfoToMap(model *backuprecoveryv1.ProgressTaskInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(&eventsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressBackupRunProgressInfoToMap(model *backuprecoveryv1.BackupRunProgressInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(&eventsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(&objectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressReplicationTargetProgressInfoToMap(model *backuprecoveryv1.ReplicationTargetProgressInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = *model.ClusterName
	}
	if model.AwsTargetConfig != nil {
		awsTargetConfigMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAWSTargetConfigToMap(model.AwsTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["aws_target_config"] = []map[string]interface{}{awsTargetConfigMap}
	}
	if model.AzureTargetConfig != nil {
		azureTargetConfigMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressAzureTargetConfigToMap(model.AzureTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["azure_target_config"] = []map[string]interface{}{azureTargetConfigMap}
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Events != nil {
		events := []map[string]interface{}{}
		for _, eventsItem := range model.Events {
			eventsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(&eventsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			events = append(events, eventsItemMap)
		}
		modelMap["events"] = events
	}
	if model.ExpectedRemainingTimeUsecs != nil {
		modelMap["expected_remaining_time_usecs"] = flex.IntValue(model.ExpectedRemainingTimeUsecs)
	}
	if model.PercentageCompleted != nil {
		modelMap["percentage_completed"] = flex.Float64Value(model.PercentageCompleted)
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := DataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(&objectsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressAWSTargetConfigToMap(model *backuprecoveryv1.AWSTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["region"] = flex.IntValue(model.Region)
	if model.RegionName != nil {
		modelMap["region_name"] = *model.RegionName
	}
	modelMap["source_id"] = flex.IntValue(model.SourceID)
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProtectionRunProgressAzureTargetConfigToMap(model *backuprecoveryv1.AzureTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = flex.IntValue(model.ResourceGroup)
	}
	if model.ResourceGroupName != nil {
		modelMap["resource_group_name"] = *model.ResourceGroupName
	}
	modelMap["source_id"] = flex.IntValue(model.SourceID)
	if model.StorageAccount != nil {
		modelMap["storage_account"] = flex.IntValue(model.StorageAccount)
	}
	if model.StorageAccountName != nil {
		modelMap["storage_account_name"] = *model.StorageAccountName
	}
	if model.StorageContainer != nil {
		modelMap["storage_container"] = flex.IntValue(model.StorageContainer)
	}
	if model.StorageContainerName != nil {
		modelMap["storage_container_name"] = *model.StorageContainerName
	}
	if model.StorageResourceGroup != nil {
		modelMap["storage_resource_group"] = flex.IntValue(model.StorageResourceGroup)
	}
	if model.StorageResourceGroupName != nil {
		modelMap["storage_resource_group_name"] = *model.StorageResourceGroupName
	}
	return modelMap, nil
}
