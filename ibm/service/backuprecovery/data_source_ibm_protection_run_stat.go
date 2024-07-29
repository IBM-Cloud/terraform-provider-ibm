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

func DataSourceIbmProtectionRunStat() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionRunStatRead,

		Schema: map[string]*schema.Schema{
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique run id of the Protection Run.",
			},
			"objects": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the objects whose stats will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run stats will not be returned and only the stats of the specified objects will be returned.",
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
				Description: "Specifies the time after which the stats task starts in Unix epoch Timestamp(in microseconds).",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the time before which the stats task ends in Unix epoch Timestamp(in microseconds).",
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
			"run_task_path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the task path of the run or object run. This is applicable only if stats of a protection group with one or more object is required. If provided this will be used to fetch stats details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format.",
			},
			"object_task_paths": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch stats details directly without looking actuall task path of the object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"local_run": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the stats of a local backup run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_generic_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats which are generic for all adapters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_errors": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of errors for this run.",
									},
									"remaining_data_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the amount of data which has to be ingested in bytes.",
									},
									"data_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the amount of data which has been ingested in bytes.",
									},
									"data_ingestion_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which data is being ingested in bytes per minute.",
									},
									"estimated_backup_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time in which backup should finish in minutes.",
									},
									"error_classes": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Divides the errors into classes for better understanding for the user.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"nas_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats which are specific for NAS adapter.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"files_discovered": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have already been discovered.",
									},
									"file_discovery_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being discovered in files per minute.",
									},
									"file_discovery_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time taken for file discovery.",
									},
									"files_analyzed": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have been analyzed.",
									},
									"file_analysis_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being analyzed in files per minute.",
									},
									"files_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have been ingested.",
									},
									"file_ingestion_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being ingested in files per minute.",
									},
									"walker_run_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the run time for directory walker in seconds.",
									},
								},
							},
						},
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies stats for objects.",
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
									"backup_generic_stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats which are generic for all adapters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"num_errors": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of errors for this run.",
												},
												"remaining_data_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the amount of data which has to be ingested in bytes.",
												},
												"data_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the amount of data which has been ingested in bytes.",
												},
												"data_ingestion_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which data is being ingested in bytes per minute.",
												},
												"estimated_backup_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time in which backup should finish in minutes.",
												},
												"error_classes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Divides the errors into classes for better understanding for the user.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"nas_stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats which are specific for NAS adapter.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"files_discovered": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have already been discovered.",
												},
												"file_discovery_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being discovered in files per minute.",
												},
												"file_discovery_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time taken for file discovery.",
												},
												"files_analyzed": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have been analyzed.",
												},
												"file_analysis_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being analyzed in files per minute.",
												},
												"files_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have been ingested.",
												},
												"file_ingestion_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being ingested in files per minute.",
												},
												"walker_run_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the run time for directory walker in seconds.",
												},
											},
										},
									},
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies stats for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_generic_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats which are generic for all adapters.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"num_errors": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of errors for this run.",
															},
															"remaining_data_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the amount of data which has to be ingested in bytes.",
															},
															"data_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the amount of data which has been ingested in bytes.",
															},
															"data_ingestion_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which data is being ingested in bytes per minute.",
															},
															"estimated_backup_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time in which backup should finish in minutes.",
															},
															"error_classes": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Divides the errors into classes for better understanding for the user.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"class_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"nas_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats which are specific for NAS adapter.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"files_discovered": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have already been discovered.",
															},
															"file_discovery_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being discovered in files per minute.",
															},
															"file_discovery_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time taken for file discovery.",
															},
															"files_analyzed": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have been analyzed.",
															},
															"file_analysis_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being analyzed in files per minute.",
															},
															"files_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have been ingested.",
															},
															"file_ingestion_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being ingested in files per minute.",
															},
															"walker_run_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the run time for directory walker in seconds.",
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
				Description: "Stats for the archival run.",
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
						"backup_generic_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats which are generic for all adapters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_errors": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of errors for this run.",
									},
									"remaining_data_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the amount of data which has to be ingested in bytes.",
									},
									"data_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the amount of data which has been ingested in bytes.",
									},
									"data_ingestion_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which data is being ingested in bytes per minute.",
									},
									"estimated_backup_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time in which backup should finish in minutes.",
									},
									"error_classes": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Divides the errors into classes for better understanding for the user.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"nas_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats which are specific for NAS adapter.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"files_discovered": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have already been discovered.",
									},
									"file_discovery_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being discovered in files per minute.",
									},
									"file_discovery_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the time taken for file discovery.",
									},
									"files_analyzed": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have been analyzed.",
									},
									"file_analysis_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being analyzed in files per minute.",
									},
									"files_ingested": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of files which have been ingested.",
									},
									"file_ingestion_rate": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the rate at which files are being ingested in files per minute.",
									},
									"walker_run_time": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the run time for directory walker in seconds.",
									},
								},
							},
						},
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies stats for objects.",
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
									"backup_generic_stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats which are generic for all adapters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"num_errors": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of errors for this run.",
												},
												"remaining_data_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the amount of data which has to be ingested in bytes.",
												},
												"data_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the amount of data which has been ingested in bytes.",
												},
												"data_ingestion_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which data is being ingested in bytes per minute.",
												},
												"estimated_backup_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time in which backup should finish in minutes.",
												},
												"error_classes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Divides the errors into classes for better understanding for the user.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"class_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"nas_stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the stats which are specific for NAS adapter.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"files_discovered": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have already been discovered.",
												},
												"file_discovery_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being discovered in files per minute.",
												},
												"file_discovery_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the time taken for file discovery.",
												},
												"files_analyzed": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have been analyzed.",
												},
												"file_analysis_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being analyzed in files per minute.",
												},
												"files_ingested": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of files which have been ingested.",
												},
												"file_ingestion_rate": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the rate at which files are being ingested in files per minute.",
												},
												"walker_run_time": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the run time for directory walker in seconds.",
												},
											},
										},
									},
									"failed_attempts": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies stats for failed attempts of this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_generic_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats which are generic for all adapters.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"num_errors": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of errors for this run.",
															},
															"remaining_data_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the amount of data which has to be ingested in bytes.",
															},
															"data_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the amount of data which has been ingested in bytes.",
															},
															"data_ingestion_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which data is being ingested in bytes per minute.",
															},
															"estimated_backup_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time in which backup should finish in minutes.",
															},
															"error_classes": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Divides the errors into classes for better understanding for the user.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"class_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"nas_stats": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the stats which are specific for NAS adapter.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"files_discovered": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have already been discovered.",
															},
															"file_discovery_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being discovered in files per minute.",
															},
															"file_discovery_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the time taken for file discovery.",
															},
															"files_analyzed": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have been analyzed.",
															},
															"file_analysis_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being analyzed in files per minute.",
															},
															"files_ingested": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of files which have been ingested.",
															},
															"file_ingestion_rate": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the rate at which files are being ingested in files per minute.",
															},
															"walker_run_time": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the run time for directory walker in seconds.",
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

func dataSourceIbmProtectionRunStatRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionRunStatsOptions := &backuprecoveryv1.GetProtectionRunStatsOptions{}

	getProtectionRunStatsOptions.SetRunID(d.Get("run_id").(string))
	if _, ok := d.GetOk("objects"); ok {
		var objects []int64
		for _, v := range d.Get("objects").([]interface{}) {
			objectsItem := int64(v.(int))
			objects = append(objects, objectsItem)
		}
		getProtectionRunStatsOptions.SetObjects(objects)
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getProtectionRunStatsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getProtectionRunStatsOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("include_finished_tasks"); ok {
		getProtectionRunStatsOptions.SetIncludeFinishedTasks(d.Get("include_finished_tasks").(bool))
	}
	if _, ok := d.GetOk("start_time_usecs"); ok {
		getProtectionRunStatsOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getProtectionRunStatsOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("max_tasks_num"); ok {
		getProtectionRunStatsOptions.SetMaxTasksNum(int64(d.Get("max_tasks_num").(int)))
	}
	if _, ok := d.GetOk("exclude_object_details"); ok {
		getProtectionRunStatsOptions.SetExcludeObjectDetails(d.Get("exclude_object_details").(bool))
	}
	if _, ok := d.GetOk("run_task_path"); ok {
		getProtectionRunStatsOptions.SetRunTaskPath(d.Get("run_task_path").(string))
	}
	if _, ok := d.GetOk("object_task_paths"); ok {
		var objectTaskPaths []string
		for _, v := range d.Get("object_task_paths").([]interface{}) {
			objectTaskPathsItem := v.(string)
			objectTaskPaths = append(objectTaskPaths, objectTaskPathsItem)
		}
		getProtectionRunStatsOptions.SetObjectTaskPaths(objectTaskPaths)
	}

	getProtectionRunStats, response, err := backupRecoveryClient.GetProtectionRunStatsWithContext(context, getProtectionRunStatsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionRunStatsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionRunStatsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmProtectionRunStatID(d))

	localRun := []map[string]interface{}{}
	if getProtectionRunStats.LocalRun != nil {
		modelMap, err := dataSourceIbmProtectionRunStatBackupRunStatsInfoToMap(getProtectionRunStats.LocalRun)
		if err != nil {
			return diag.FromErr(err)
		}
		localRun = append(localRun, modelMap)
	}
	if err = d.Set("local_run", localRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting local_run %s", err))
	}

	archivalRun := []map[string]interface{}{}
	if getProtectionRunStats.ArchivalRun != nil {
		for _, modelItem := range getProtectionRunStats.ArchivalRun {
			modelMap, err := dataSourceIbmProtectionRunStatArchivalTargetStatsInfoToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			archivalRun = append(archivalRun, modelMap)
		}
	}
	if err = d.Set("archival_run", archivalRun); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting archival_run %s", err))
	}

	return nil
}

// dataSourceIbmProtectionRunStatID returns a reasonable ID for the list.
func dataSourceIbmProtectionRunStatID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmProtectionRunStatBackupRunStatsInfoToMap(model *backuprecoveryv1.BackupRunStatsInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupGenericStats != nil {
		backupGenericStatsMap, err := dataSourceIbmProtectionRunStatBackupGenericStatsToMap(model.BackupGenericStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["backup_generic_stats"] = []map[string]interface{}{backupGenericStatsMap}
	}
	if model.NasStats != nil {
		nasStatsMap, err := dataSourceIbmProtectionRunStatBackupNasStatsToMap(model.NasStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_stats"] = []map[string]interface{}{nasStatsMap}
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionRunStatObjectStatsInfoToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatBackupGenericStatsToMap(model *backuprecoveryv1.BackupGenericStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NumErrors != nil {
		modelMap["num_errors"] = flex.IntValue(model.NumErrors)
	}
	if model.RemainingDataIngested != nil {
		modelMap["remaining_data_ingested"] = flex.IntValue(model.RemainingDataIngested)
	}
	if model.DataIngested != nil {
		modelMap["data_ingested"] = flex.IntValue(model.DataIngested)
	}
	if model.DataIngestionRate != nil {
		modelMap["data_ingestion_rate"] = flex.IntValue(model.DataIngestionRate)
	}
	if model.EstimatedBackupTime != nil {
		modelMap["estimated_backup_time"] = flex.IntValue(model.EstimatedBackupTime)
	}
	if model.ErrorClasses != nil {
		errorClasses := []map[string]interface{}{}
		for _, errorClassesItem := range model.ErrorClasses {
			errorClassesItemMap, err := dataSourceIbmProtectionRunStatErrorClassToMap(&errorClassesItem)
			if err != nil {
				return modelMap, err
			}
			errorClasses = append(errorClasses, errorClassesItemMap)
		}
		modelMap["error_classes"] = errorClasses
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatErrorClassToMap(model *backuprecoveryv1.ErrorClass) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClassName != nil {
		modelMap["class_name"] = model.ClassName
	}
	if model.Count != nil {
		modelMap["count"] = flex.IntValue(model.Count)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatBackupNasStatsToMap(model *backuprecoveryv1.BackupNasStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FilesDiscovered != nil {
		modelMap["files_discovered"] = flex.IntValue(model.FilesDiscovered)
	}
	if model.FileDiscoveryRate != nil {
		modelMap["file_discovery_rate"] = flex.IntValue(model.FileDiscoveryRate)
	}
	if model.FileDiscoveryTime != nil {
		modelMap["file_discovery_time"] = flex.IntValue(model.FileDiscoveryTime)
	}
	if model.FilesAnalyzed != nil {
		modelMap["files_analyzed"] = flex.IntValue(model.FilesAnalyzed)
	}
	if model.FileAnalysisRate != nil {
		modelMap["file_analysis_rate"] = flex.IntValue(model.FileAnalysisRate)
	}
	if model.FilesIngested != nil {
		modelMap["files_ingested"] = flex.IntValue(model.FilesIngested)
	}
	if model.FileIngestionRate != nil {
		modelMap["file_ingestion_rate"] = flex.IntValue(model.FileIngestionRate)
	}
	if model.WalkerRunTime != nil {
		modelMap["walker_run_time"] = flex.IntValue(model.WalkerRunTime)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatObjectStatsInfoToMap(model *backuprecoveryv1.ObjectStatsInfo) (map[string]interface{}, error) {
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
	if model.BackupGenericStats != nil {
		backupGenericStatsMap, err := dataSourceIbmProtectionRunStatBackupGenericStatsToMap(model.BackupGenericStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["backup_generic_stats"] = []map[string]interface{}{backupGenericStatsMap}
	}
	if model.NasStats != nil {
		nasStatsMap, err := dataSourceIbmProtectionRunStatBackupNasStatsToMap(model.NasStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_stats"] = []map[string]interface{}{nasStatsMap}
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := dataSourceIbmProtectionRunStatStatsTaskInfoToMap(&failedAttemptsItem)
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatStatsTaskInfoToMap(model *backuprecoveryv1.StatsTaskInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupGenericStats != nil {
		backupGenericStatsMap, err := dataSourceIbmProtectionRunStatBackupGenericStatsToMap(model.BackupGenericStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["backup_generic_stats"] = []map[string]interface{}{backupGenericStatsMap}
	}
	if model.NasStats != nil {
		nasStatsMap, err := dataSourceIbmProtectionRunStatBackupNasStatsToMap(model.NasStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_stats"] = []map[string]interface{}{nasStatsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatArchivalTargetStatsInfoToMap(model *backuprecoveryv1.ArchivalTargetStatsInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmProtectionRunStatArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.BackupGenericStats != nil {
		backupGenericStatsMap, err := dataSourceIbmProtectionRunStatBackupGenericStatsToMap(model.BackupGenericStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["backup_generic_stats"] = []map[string]interface{}{backupGenericStatsMap}
	}
	if model.NasStats != nil {
		nasStatsMap, err := dataSourceIbmProtectionRunStatBackupNasStatsToMap(model.NasStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_stats"] = []map[string]interface{}{nasStatsMap}
	}
	if model.Objects != nil {
		objects := []map[string]interface{}{}
		for _, objectsItem := range model.Objects {
			objectsItemMap, err := dataSourceIbmProtectionRunStatObjectStatsInfoToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmProtectionRunStatOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmProtectionRunStatOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmProtectionRunStatOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmProtectionRunStatOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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
