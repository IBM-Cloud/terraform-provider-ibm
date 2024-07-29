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

func ResourceIbmProtectionGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProtectionGroupCreate,
		ReadContext:   resourceIbmProtectionGroupRead,
		UpdateContext: resourceIbmProtectionGroupUpdate,
		DeleteContext: resourceIbmProtectionGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the name of the Protection Group.",
			},
			"policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc.",
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//ValidateFunc: validate.InvokeValidator("ibm_protection_group", "priority"),
				Description: "Specifies the priority of the Protection Group.",
			},
			"storage_domain_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the Storage Domain (View Box) ID where this Protection Group writes data.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies a description of the Protection Group.",
			},
			"start_time": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies the time of day. Used for scheduling purposes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hour": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the hour of the day (0-23).",
						},
						"minute": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the minute of the hour (0-59).",
						},
						"time_zone": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "America/Los_Angeles",
							Description: "Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.",
						},
					},
				},
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.",
			},
			"last_modified_timestamp_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.",
			},
			"alert_policy": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies a policy for alerting users of the status of a Protection Group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_run_status": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Specifies the run status for which the user would like to receive alerts.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"alert_targets": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of targets to receive the alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email_address": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies an email address to receive an alert.",
									},
									"language": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the language of the delivery target. Default value is 'en-us'.",
									},
									"recipient_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the recipient type of email recipient. Default value is 'kTo'.",
									},
								},
							},
						},
						"raise_object_level_failure_alert": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether object level alerts are raised for backup failures after the backup run.",
						},
						"raise_object_level_failure_alert_after_last_attempt": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether object level alerts are raised for backup failures after last backup attempt.",
						},
						"raise_object_level_failure_alert_after_each_attempt": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Specifies whether object level alerts are raised for backup failures after each backup attempt.",
						},
					},
				},
			},
			"sla": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the SLA parameters for this Protection Group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_run_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the type of run this rule should apply to.",
						},
						"sla_minutes": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the number of minutes allotted to a run of the specified type before SLA is considered violated.",
						},
					},
				},
			},
			"qos_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ValidateFunc: validate.InvokeValidator("ibm_protection_group", "qos_policy"),
				Description: "Specifies whether the Protection Group will be written to HDD or SSD.",
			},
			"abort_in_blackouts": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false.",
			},
			"pause_in_blackouts": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true.",
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_protection_group_request", "environment"),
				Description: "Specifies the environment type of the Protection Group.",
			},
			"is_paused": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted.",
			},
			"advanced_configs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the advanced configuration for a protection job.",
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
			"physical_params": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protection_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the Physical Protection Group type.",
						},
						"volume_protection_type_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters which are specific to Volume based physical Protection Groups.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"objects": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the ID of the object protected.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the name of the object protected.",
												},
												"volume_guids": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the list of GUIDs of volumes protected. If empty, then all volumes will be protected by default.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"enable_system_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether or not to take a system backup. Applicable only for windows sources.",
												},
											},
										},
									},
									"indexing_policy": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable_indexing": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.",
												},
												"include_paths": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"exclude_paths": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"perform_source_side_deduplication": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether or not to perform source side deduplication on this Protection Group.",
									},
									"quiesce": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.",
									},
									"continue_on_quiesce_failure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether to continue backing up on quiesce failure.",
									},
									"incremental_backup_after_restart": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether or not to perform an incremental backup after the server restarts. This is applicable to windows environments.",
									},
									"pre_post_script": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the params for pre and post scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_script": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the common params for PreBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether the script should be enabled, default value set to true.",
															},
															"continue_on_error": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
															},
														},
													},
												},
												"post_script": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the common params for PostBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
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
										Optional:    true,
										Description: "Specifies ids of sources for which deduplication has to be disabled.",
										Elem:        &schema.Schema{Type: schema.TypeInt},
									},
									"excluded_vss_writers": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies writer names which should be excluded from physical volume based backups.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"cobmr_backup": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether to take a CoBMR backup.",
									},
								},
							},
						},
						"file_protection_type_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters which are specific to Physical related Protection Groups.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"objects": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "Specifies the list of objects protected by this Protection Group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the ID of the object protected.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the object protected.",
												},
												"file_paths": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies a list of file paths to be protected by this Protection Group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"included_path": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies a path to be included on the source. All paths under this path will be included unless they are specifically mentioned in excluded paths.",
															},
															"excluded_paths": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies a set of paths nested under the include path which should be excluded from the Protection Group.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"skip_nested_volumes": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether to skip any nested volumes (both local and network) that are mounted under include path. Applicable only for windows sources.",
															},
														},
													},
												},
												"uses_path_level_skip_nested_volume_setting": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether path level or object level skip nested volume setting will be used.",
												},
												"nested_volume_types_to_skip": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies mount types of nested volumes to be skipped.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"follow_nas_symlink_target": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether to follow NAS target pointed by symlink for windows sources.",
												},
												"metadata_file_path": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the path of metadatafile on source. This file contains absolute paths of files that needs to be backed up on the same source.",
												},
											},
										},
									},
									"indexing_policy": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable_indexing": &schema.Schema{
													Type:        schema.TypeBool,
													Required:    true,
													Description: "Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.",
												},
												"include_paths": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"exclude_paths": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"perform_source_side_deduplication": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether or not to perform source side deduplication on this Protection Group.",
									},
									"perform_brick_based_deduplication": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether or not to perform brick based deduplication on this Protection Group.",
									},
									"task_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the timeouts for all the objects inside this Protection Group, for both full and incremental backups.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"timeout_mins": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the timeout in mins.",
												},
												"backup_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The scheduled backup type(kFull, kRegular etc.).",
												},
											},
										},
									},
									"quiesce": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.",
									},
									"continue_on_quiesce_failure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether to continue backing up on quiesce failure.",
									},
									"cobmr_backup": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies whether to take CoBMR backup.",
									},
									"pre_post_script": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the params for pre and post scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_script": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the common params for PreBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether the script should be enabled, default value set to true.",
															},
															"continue_on_error": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
															},
														},
													},
												},
												"post_script": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the common params for PostBackup scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the absolute path to the script on the remote host.",
															},
															"params": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
															},
															"timeout_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
															},
															"is_active": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
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
										Optional:    true,
										Description: "Specifies ids of sources for which deduplication has to be disabled.",
										Elem:        &schema.Schema{Type: schema.TypeInt},
									},
									"global_exclude_paths": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies global exclude filters which are applied to all sources in a job.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"global_exclude_fs": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies global exclude filesystems which are applied to all sources in a job.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"ignorable_errors": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the Errors to be ignored in error db.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"allow_parallel_runs": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
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
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies the parameters to create Oracle Protection Group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"objects": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Specifies the list of object ids to be protected.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies the id of the host on which databases are hosted.",
									},
									"source_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the host on which databases are hosted.",
									},
									"db_params": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the properties of the Oracle databases.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the id of the Oracle database.",
												},
												"database_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the Oracle database.",
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
											},
										},
									},
								},
							},
						},
						"persist_mountpoints": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. Defaults to true if value is null to handle the backward compatibility for the upgrade case.",
						},
						"vlan_params": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies VLAN params associated with the backup/restore operation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_id": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
									},
									"disable_vlan": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
									},
									"interface_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
									},
								},
							},
						},
						"pre_post_script": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the params for pre and post scripts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pre_script": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the common params for PreBackup scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the absolute path to the script on the remote host.",
												},
												"params": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
												},
												"timeout_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether the script should be enabled, default value set to true.",
												},
												"continue_on_error": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.",
												},
											},
										},
									},
									"post_script": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the common params for PostBackup scripts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the absolute path to the script on the remote host.",
												},
												"params": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: \"database=myDatabase user=me\".",
												},
												"timeout_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.",
												},
												"is_active": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
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
							Optional:    true,
							Description: "Time in seconds after which the log backup of the database in given backup job should be auto-killed.",
						},
						"incr_auto_kill_timeout_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Time in seconds after which the incremental backup of the database in given backup job should be auto-killed.",
						},
						"full_auto_kill_timeout_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Time in seconds after which the full backup of the database in given backup job should be auto-killed.",
						},
					},
				},
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
																Elem:        &schema.Schema{Type: schema.TypeString},
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
																Elem:        &schema.Schema{Type: schema.TypeString},
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
										Elem:        &schema.Schema{Type: schema.TypeString},
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
										Elem:        &schema.Schema{Type: schema.TypeString},
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
		},
	}
}

func ResourceIbmProtectionGroupValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "priority",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "kHigh, kLow, kMedium",
		},
		validate.ValidateSchema{
			Identifier:                 "qos_policy",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "kBackupAll, kBackupHDD, kBackupSSD, kTestAndDevHigh",
		},
		validate.ValidateSchema{
			Identifier:                 "environment",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kOracle, kPhysical, kSQL",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_protection_group", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProtectionGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createProtectionGroupOptions := &backuprecoveryv1.CreateProtectionGroupOptions{}

	createProtectionGroupOptions.SetName(d.Get("name").(string))
	createProtectionGroupOptions.SetPolicyID(d.Get("policy_id").(string))
	createProtectionGroupOptions.SetEnvironment(d.Get("environment").(string))
	if _, ok := d.GetOk("priority"); ok {
		createProtectionGroupOptions.SetPriority(d.Get("priority").(string))
	}
	if _, ok := d.GetOk("storage_domain_id"); ok {
		createProtectionGroupOptions.SetStorageDomainID(int64(d.Get("storage_domain_id").(int)))
	}
	if _, ok := d.GetOk("description"); ok {
		createProtectionGroupOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("start_time"); ok {
		startTimeModel, err := resourceIbmProtectionGroupMapToTimeOfDay(d.Get("start_time.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupOptions.SetStartTime(startTimeModel)
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		createProtectionGroupOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("last_modified_timestamp_usecs"); ok {
		createProtectionGroupOptions.SetLastModifiedTimestampUsecs(int64(d.Get("last_modified_timestamp_usecs").(int)))
	}
	if _, ok := d.GetOk("alert_policy"); ok {
		alertPolicyModel, err := resourceIbmProtectionGroupMapToProtectionGroupAlertingPolicy(d.Get("alert_policy.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupOptions.SetAlertPolicy(alertPolicyModel)
	}
	if _, ok := d.GetOk("sla"); ok {
		var sla []backuprecoveryv1.SlaRule
		for _, v := range d.Get("sla").([]interface{}) {
			value := v.(map[string]interface{})
			slaItem, err := resourceIbmProtectionGroupMapToSlaRule(value)
			if err != nil {
				return diag.FromErr(err)
			}
			sla = append(sla, *slaItem)
		}
		createProtectionGroupOptions.SetSla(sla)
	}
	if _, ok := d.GetOk("qos_policy"); ok {
		createProtectionGroupOptions.SetQosPolicy(d.Get("qos_policy").(string))
	}
	if _, ok := d.GetOk("abort_in_blackouts"); ok {
		createProtectionGroupOptions.SetAbortInBlackouts(d.Get("abort_in_blackouts").(bool))
	}
	if _, ok := d.GetOk("pause_in_blackouts"); ok {
		createProtectionGroupOptions.SetPauseInBlackouts(d.Get("pause_in_blackouts").(bool))
	}
	if _, ok := d.GetOk("is_paused"); ok {
		createProtectionGroupOptions.SetIsPaused(d.Get("is_paused").(bool))
	}
	if _, ok := d.GetOk("advanced_configs"); ok {
		var advancedConfigs []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			advancedConfigsItem, err := resourceIbmProtectionGroupMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, *advancedConfigsItem)
		}
		createProtectionGroupOptions.SetAdvancedConfigs(advancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		physicalParamsModel, err := resourceIbmProtectionGroupMapToPhysicalProtectionGroupParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupOptions.SetPhysicalParams(physicalParamsModel)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		oracleParamsModel, err := resourceIbmProtectionGroupMapToOracleProtectionGroupParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupOptions.SetOracleParams(oracleParamsModel)
	}

	protectionGroupResponse, response, err := backupRecoveryClient.CreateProtectionGroupWithContext(context, createProtectionGroupOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateProtectionGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateProtectionGroupWithContext failed %s\n%s", err, response))
	}

	d.SetId(*protectionGroupResponse.ID)

	return resourceIbmProtectionGroupRead(context, d, meta)
}

func resourceIbmProtectionGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

	getProtectionGroupByIdOptions.SetID(d.Id())

	protectionGroupResponse, response, err := backupRecoveryClient.GetProtectionGroupByIDWithContext(context, getProtectionGroupByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProtectionGroupByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionGroupByIDWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", protectionGroupResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("policy_id", protectionGroupResponse.PolicyID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting policy_id: %s", err))
	}
	if !core.IsNil(protectionGroupResponse.Priority) {
		if err = d.Set("priority", protectionGroupResponse.Priority); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting priority: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.StorageDomainID) {
		if err = d.Set("storage_domain_id", flex.IntValue(protectionGroupResponse.StorageDomainID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting storage_domain_id: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.Description) {
		if err = d.Set("description", protectionGroupResponse.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.StartTime) {
		startTimeMap, err := resourceIbmProtectionGroupTimeOfDayToMap(protectionGroupResponse.StartTime)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("start_time", []map[string]interface{}{startTimeMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting start_time: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.EndTimeUsecs) {
		if err = d.Set("end_time_usecs", flex.IntValue(protectionGroupResponse.EndTimeUsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting end_time_usecs: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.LastModifiedTimestampUsecs) {
		if err = d.Set("last_modified_timestamp_usecs", flex.IntValue(protectionGroupResponse.LastModifiedTimestampUsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_modified_timestamp_usecs: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.AlertPolicy) {
		alertPolicyMap, err := resourceIbmProtectionGroupProtectionGroupAlertingPolicyToMap(protectionGroupResponse.AlertPolicy)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("alert_policy", []map[string]interface{}{alertPolicyMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting alert_policy: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.Sla) {
		sla := []map[string]interface{}{}
		for _, slaItem := range protectionGroupResponse.Sla {
			slaItemMap, err := resourceIbmProtectionGroupSlaRuleToMap(&slaItem)
			if err != nil {
				return diag.FromErr(err)
			}
			sla = append(sla, slaItemMap)
		}
		if err = d.Set("sla", sla); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting sla: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.QosPolicy) {
		if err = d.Set("qos_policy", protectionGroupResponse.QosPolicy); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting qos_policy: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.AbortInBlackouts) {
		if err = d.Set("abort_in_blackouts", protectionGroupResponse.AbortInBlackouts); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting abort_in_blackouts: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.PauseInBlackouts) {
		if err = d.Set("pause_in_blackouts", protectionGroupResponse.PauseInBlackouts); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting pause_in_blackouts: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.IsPaused) {
		if err = d.Set("is_paused", protectionGroupResponse.IsPaused); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_paused: %s", err))
		}
	}
	if err = d.Set("environment", protectionGroupResponse.Environment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
	}
	if !core.IsNil(protectionGroupResponse.AdvancedConfigs) {
		advancedConfigs := []map[string]interface{}{}
		for _, advancedConfigsItem := range protectionGroupResponse.AdvancedConfigs {
			advancedConfigsItemMap, err := resourceIbmProtectionGroupKeyValuePairToMap(&advancedConfigsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, advancedConfigsItemMap)
		}
		if err = d.Set("advanced_configs", advancedConfigs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting advanced_configs: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.PhysicalParams) {
		physicalParamsMap, err := resourceIbmProtectionGroupPhysicalProtectionGroupParamsToMap(protectionGroupResponse.PhysicalParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("physical_params", []map[string]interface{}{physicalParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting physical_params: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.OracleParams) {
		oracleParamsMap, err := resourceIbmProtectionGroupOracleProtectionGroupParamsToMap(protectionGroupResponse.OracleParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("oracle_params", []map[string]interface{}{oracleParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting oracle_params: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.ClusterID) {
		if err = d.Set("cluster_id", protectionGroupResponse.ClusterID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cluster_id: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.RegionID) {
		if err = d.Set("region_id", protectionGroupResponse.RegionID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting region_id: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.IsActive) {
		if err = d.Set("is_active", protectionGroupResponse.IsActive); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_active: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.IsDeleted) {
		if err = d.Set("is_deleted", protectionGroupResponse.IsDeleted); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_deleted: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.LastRun) {
		lastRunMap, err := resourceIbmProtectionGroupProtectionGroupRunToMap(protectionGroupResponse.LastRun)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("last_run", []map[string]interface{}{lastRunMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_run: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.Permissions) {
		permissions := []map[string]interface{}{}
		for _, permissionsItem := range protectionGroupResponse.Permissions {
			permissionsItemMap, err := resourceIbmProtectionGroupTenantToMap(&permissionsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			permissions = append(permissions, permissionsItemMap)
		}
		if err = d.Set("permissions", permissions); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting permissions: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.IsProtectOnce) {
		if err = d.Set("is_protect_once", protectionGroupResponse.IsProtectOnce); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_protect_once: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.MissingEntities) {
		missingEntities := []map[string]interface{}{}
		for _, missingEntitiesItem := range protectionGroupResponse.MissingEntities {
			missingEntitiesItemMap, err := resourceIbmProtectionGroupMissingEntityParamsToMap(&missingEntitiesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			missingEntities = append(missingEntities, missingEntitiesItemMap)
		}
		if err = d.Set("missing_entities", missingEntities); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting missing_entities: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.InvalidEntities) {
		invalidEntities := []map[string]interface{}{}
		for _, invalidEntitiesItem := range protectionGroupResponse.InvalidEntities {
			invalidEntitiesItemMap, err := resourceIbmProtectionGroupMissingEntityParamsToMap(&invalidEntitiesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			invalidEntities = append(invalidEntities, invalidEntitiesItemMap)
		}
		if err = d.Set("invalid_entities", invalidEntities); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting invalid_entities: %s", err))
		}
	}
	if !core.IsNil(protectionGroupResponse.NumProtectedObjects) {
		if err = d.Set("num_protected_objects", flex.IntValue(protectionGroupResponse.NumProtectedObjects)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting num_protected_objects: %s", err))
		}
	}

	return nil
}

func resourceIbmProtectionGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProtectionGroupOptions := &backuprecoveryv1.UpdateProtectionGroupOptions{}

	updateProtectionGroupOptions.SetID(d.Id())
	updateProtectionGroupOptions.SetName(d.Get("name").(string))
	updateProtectionGroupOptions.SetPolicyID(d.Get("policy_id").(string))
	updateProtectionGroupOptions.SetEnvironment(d.Get("environment").(string))
	// if _, ok := d.GetOk("id"); ok {
	// 	updateProtectionGroupOptions.SetID(d.Get("id").(string))
	// }
	if _, ok := d.GetOk("priority"); ok {
		updateProtectionGroupOptions.SetPriority(d.Get("priority").(string))
	}
	if _, ok := d.GetOk("storage_domain_id"); ok {
		updateProtectionGroupOptions.SetStorageDomainID(int64(d.Get("storage_domain_id").(int)))
	}
	if _, ok := d.GetOk("description"); ok {
		updateProtectionGroupOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("start_time"); ok {
		newStartTime, err := resourceIbmProtectionGroupMapToTimeOfDay(d.Get("start_time.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionGroupOptions.SetStartTime(newStartTime)
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		updateProtectionGroupOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("last_modified_timestamp_usecs"); ok {
		updateProtectionGroupOptions.SetLastModifiedTimestampUsecs(int64(d.Get("last_modified_timestamp_usecs").(int)))
	}
	if _, ok := d.GetOk("alert_policy"); ok {
		newAlertPolicy, err := resourceIbmProtectionGroupMapToProtectionGroupAlertingPolicy(d.Get("alert_policy.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionGroupOptions.SetAlertPolicy(newAlertPolicy)
	}
	if _, ok := d.GetOk("sla"); ok {
		var newSla []backuprecoveryv1.SlaRule
		for _, v := range d.Get("sla").([]interface{}) {
			value := v.(map[string]interface{})
			newSlaItem, err := resourceIbmProtectionGroupMapToSlaRule(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newSla = append(newSla, *newSlaItem)
		}
		updateProtectionGroupOptions.SetSla(newSla)
	}
	if _, ok := d.GetOk("qos_policy"); ok {
		updateProtectionGroupOptions.SetQosPolicy(d.Get("qos_policy").(string))
	}
	if _, ok := d.GetOk("abort_in_blackouts"); ok {
		updateProtectionGroupOptions.SetAbortInBlackouts(d.Get("abort_in_blackouts").(bool))
	}
	if _, ok := d.GetOk("pause_in_blackouts"); ok {
		updateProtectionGroupOptions.SetPauseInBlackouts(d.Get("pause_in_blackouts").(bool))
	}
	if _, ok := d.GetOk("is_paused"); ok {
		updateProtectionGroupOptions.SetIsPaused(d.Get("is_paused").(bool))
	}
	if _, ok := d.GetOk("advanced_configs"); ok {
		var newAdvancedConfigs []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			newAdvancedConfigsItem, err := resourceIbmProtectionGroupMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newAdvancedConfigs = append(newAdvancedConfigs, *newAdvancedConfigsItem)
		}
		updateProtectionGroupOptions.SetAdvancedConfigs(newAdvancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		newPhysicalParams, err := resourceIbmProtectionGroupMapToPhysicalProtectionGroupParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionGroupOptions.SetPhysicalParams(newPhysicalParams)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		newOracleParams, err := resourceIbmProtectionGroupMapToOracleProtectionGroupParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionGroupOptions.SetOracleParams(newOracleParams)
	}

	_, response, err := backupRecoveryClient.UpdateProtectionGroupWithContext(context, updateProtectionGroupOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateProtectionGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateProtectionGroupWithContext failed %s\n%s", err, response))
	}

	return resourceIbmProtectionGroupRead(context, d, meta)
}

func resourceIbmProtectionGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionGroupOptions := &backuprecoveryv1.DeleteProtectionGroupOptions{}

	deleteProtectionGroupOptions.SetID(d.Id())

	response, err := backupRecoveryClient.DeleteProtectionGroupWithContext(context, deleteProtectionGroupOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteProtectionGroupWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteProtectionGroupWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmProtectionGroupMapToTimeOfDay(modelMap map[string]interface{}) (*backuprecoveryv1.TimeOfDay, error) {
	model := &backuprecoveryv1.TimeOfDay{}
	model.Hour = core.Int64Ptr(int64(modelMap["hour"].(int)))
	model.Minute = core.Int64Ptr(int64(modelMap["minute"].(int)))
	if modelMap["time_zone"] != nil && modelMap["time_zone"].(string) != "" {
		model.TimeZone = core.StringPtr(modelMap["time_zone"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToProtectionGroupAlertingPolicy(modelMap map[string]interface{}) (*backuprecoveryv1.ProtectionGroupAlertingPolicy, error) {
	model := &backuprecoveryv1.ProtectionGroupAlertingPolicy{}
	backupRunStatus := []string{}
	for _, backupRunStatusItem := range modelMap["backup_run_status"].([]interface{}) {
		backupRunStatus = append(backupRunStatus, backupRunStatusItem.(string))
	}
	model.BackupRunStatus = backupRunStatus
	if modelMap["alert_targets"] != nil {
		alertTargets := []backuprecoveryv1.AlertTarget{}
		for _, alertTargetsItem := range modelMap["alert_targets"].([]interface{}) {
			alertTargetsItemModel, err := resourceIbmProtectionGroupMapToAlertTarget(alertTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			alertTargets = append(alertTargets, *alertTargetsItemModel)
		}
		model.AlertTargets = alertTargets
	}
	if modelMap["raise_object_level_failure_alert"] != nil {
		model.RaiseObjectLevelFailureAlert = core.BoolPtr(modelMap["raise_object_level_failure_alert"].(bool))
	}
	if modelMap["raise_object_level_failure_alert_after_last_attempt"] != nil {
		model.RaiseObjectLevelFailureAlertAfterLastAttempt = core.BoolPtr(modelMap["raise_object_level_failure_alert_after_last_attempt"].(bool))
	}
	if modelMap["raise_object_level_failure_alert_after_each_attempt"] != nil {
		model.RaiseObjectLevelFailureAlertAfterEachAttempt = core.BoolPtr(modelMap["raise_object_level_failure_alert_after_each_attempt"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToAlertTarget(modelMap map[string]interface{}) (*backuprecoveryv1.AlertTarget, error) {
	model := &backuprecoveryv1.AlertTarget{}
	model.EmailAddress = core.StringPtr(modelMap["email_address"].(string))
	if modelMap["language"] != nil && modelMap["language"].(string) != "" {
		model.Language = core.StringPtr(modelMap["language"].(string))
	}
	if modelMap["recipient_type"] != nil && modelMap["recipient_type"].(string) != "" {
		model.RecipientType = core.StringPtr(modelMap["recipient_type"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToSlaRule(modelMap map[string]interface{}) (*backuprecoveryv1.SlaRule, error) {
	model := &backuprecoveryv1.SlaRule{}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["sla_minutes"] != nil {
		model.SlaMinutes = core.Int64Ptr(int64(modelMap["sla_minutes"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalProtectionGroupParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalProtectionGroupParams, error) {
	model := &backuprecoveryv1.PhysicalProtectionGroupParams{}
	model.ProtectionType = core.StringPtr(modelMap["protection_type"].(string))
	if modelMap["volume_protection_type_params"] != nil && len(modelMap["volume_protection_type_params"].([]interface{})) > 0 {
		VolumeProtectionTypeParamsModel, err := resourceIbmProtectionGroupMapToPhysicalVolumeProtectionGroupParams(modelMap["volume_protection_type_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VolumeProtectionTypeParams = VolumeProtectionTypeParamsModel
	}
	if modelMap["file_protection_type_params"] != nil && len(modelMap["file_protection_type_params"].([]interface{})) > 0 {
		FileProtectionTypeParamsModel, err := resourceIbmProtectionGroupMapToPhysicalFileProtectionGroupParams(modelMap["file_protection_type_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FileProtectionTypeParams = FileProtectionTypeParamsModel
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalVolumeProtectionGroupParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalVolumeProtectionGroupParams, error) {
	model := &backuprecoveryv1.PhysicalVolumeProtectionGroupParams{}
	objects := []backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{}
	for _, objectsItem := range modelMap["objects"].([]interface{}) {
		objectsItemModel, err := resourceIbmProtectionGroupMapToPhysicalVolumeProtectionGroupObjectParams(objectsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		objects = append(objects, *objectsItemModel)
	}
	model.Objects = objects
	if modelMap["indexing_policy"] != nil && len(modelMap["indexing_policy"].([]interface{})) > 0 {
		IndexingPolicyModel, err := resourceIbmProtectionGroupMapToIndexingPolicy(modelMap["indexing_policy"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IndexingPolicy = IndexingPolicyModel
	}
	if modelMap["perform_source_side_deduplication"] != nil {
		model.PerformSourceSideDeduplication = core.BoolPtr(modelMap["perform_source_side_deduplication"].(bool))
	}
	if modelMap["quiesce"] != nil {
		model.Quiesce = core.BoolPtr(modelMap["quiesce"].(bool))
	}
	if modelMap["continue_on_quiesce_failure"] != nil {
		model.ContinueOnQuiesceFailure = core.BoolPtr(modelMap["continue_on_quiesce_failure"].(bool))
	}
	if modelMap["incremental_backup_after_restart"] != nil {
		model.IncrementalBackupAfterRestart = core.BoolPtr(modelMap["incremental_backup_after_restart"].(bool))
	}
	if modelMap["pre_post_script"] != nil && len(modelMap["pre_post_script"].([]interface{})) > 0 {
		PrePostScriptModel, err := resourceIbmProtectionGroupMapToPrePostScriptParams(modelMap["pre_post_script"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PrePostScript = PrePostScriptModel
	}
	if modelMap["dedup_exclusion_source_ids"] != nil {
		dedupExclusionSourceIds := []int64{}
		for _, dedupExclusionSourceIdsItem := range modelMap["dedup_exclusion_source_ids"].([]interface{}) {
			dedupExclusionSourceIds = append(dedupExclusionSourceIds, int64(dedupExclusionSourceIdsItem.(int)))
		}
		model.DedupExclusionSourceIds = dedupExclusionSourceIds
	}
	if modelMap["excluded_vss_writers"] != nil {
		excludedVssWriters := []string{}
		for _, excludedVssWritersItem := range modelMap["excluded_vss_writers"].([]interface{}) {
			excludedVssWriters = append(excludedVssWriters, excludedVssWritersItem.(string))
		}
		model.ExcludedVssWriters = excludedVssWriters
	}
	if modelMap["cobmr_backup"] != nil {
		model.CobmrBackup = core.BoolPtr(modelMap["cobmr_backup"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalVolumeProtectionGroupObjectParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams, error) {
	model := &backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["volume_guids"] != nil {
		volumeGuids := []string{}
		for _, volumeGuidsItem := range modelMap["volume_guids"].([]interface{}) {
			volumeGuids = append(volumeGuids, volumeGuidsItem.(string))
		}
		model.VolumeGuids = volumeGuids
	}
	if modelMap["enable_system_backup"] != nil {
		model.EnableSystemBackup = core.BoolPtr(modelMap["enable_system_backup"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToIndexingPolicy(modelMap map[string]interface{}) (*backuprecoveryv1.IndexingPolicy, error) {
	model := &backuprecoveryv1.IndexingPolicy{}
	model.EnableIndexing = core.BoolPtr(modelMap["enable_indexing"].(bool))
	if modelMap["include_paths"] != nil {
		includePaths := []string{}
		for _, includePathsItem := range modelMap["include_paths"].([]interface{}) {
			includePaths = append(includePaths, includePathsItem.(string))
		}
		model.IncludePaths = includePaths
	}
	if modelMap["exclude_paths"] != nil {
		excludePaths := []string{}
		for _, excludePathsItem := range modelMap["exclude_paths"].([]interface{}) {
			excludePaths = append(excludePaths, excludePathsItem.(string))
		}
		model.ExcludePaths = excludePaths
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPrePostScriptParams(modelMap map[string]interface{}) (*backuprecoveryv1.PrePostScriptParams, error) {
	model := &backuprecoveryv1.PrePostScriptParams{}
	if modelMap["pre_script"] != nil && len(modelMap["pre_script"].([]interface{})) > 0 {
		PreScriptModel, err := resourceIbmProtectionGroupMapToCommonPreBackupScriptParams(modelMap["pre_script"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PreScript = PreScriptModel
	}
	if modelMap["post_script"] != nil && len(modelMap["post_script"].([]interface{})) > 0 {
		PostScriptModel, err := resourceIbmProtectionGroupMapToCommonPostBackupScriptParams(modelMap["post_script"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PostScript = PostScriptModel
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToCommonPreBackupScriptParams(modelMap map[string]interface{}) (*backuprecoveryv1.CommonPreBackupScriptParams, error) {
	model := &backuprecoveryv1.CommonPreBackupScriptParams{}
	model.Path = core.StringPtr(modelMap["path"].(string))
	if modelMap["params"] != nil && modelMap["params"].(string) != "" {
		model.Params = core.StringPtr(modelMap["params"].(string))
	}
	if modelMap["timeout_secs"] != nil {
		model.TimeoutSecs = core.Int64Ptr(int64(modelMap["timeout_secs"].(int)))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	if modelMap["continue_on_error"] != nil {
		model.ContinueOnError = core.BoolPtr(modelMap["continue_on_error"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToCommonPostBackupScriptParams(modelMap map[string]interface{}) (*backuprecoveryv1.CommonPostBackupScriptParams, error) {
	model := &backuprecoveryv1.CommonPostBackupScriptParams{}
	model.Path = core.StringPtr(modelMap["path"].(string))
	if modelMap["params"] != nil && modelMap["params"].(string) != "" {
		model.Params = core.StringPtr(modelMap["params"].(string))
	}
	if modelMap["timeout_secs"] != nil {
		model.TimeoutSecs = core.Int64Ptr(int64(modelMap["timeout_secs"].(int)))
	}
	if modelMap["is_active"] != nil {
		model.IsActive = core.BoolPtr(modelMap["is_active"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalFileProtectionGroupParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalFileProtectionGroupParams, error) {
	model := &backuprecoveryv1.PhysicalFileProtectionGroupParams{}
	objects := []backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{}
	for _, objectsItem := range modelMap["objects"].([]interface{}) {
		objectsItemModel, err := resourceIbmProtectionGroupMapToPhysicalFileProtectionGroupObjectParams(objectsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		objects = append(objects, *objectsItemModel)
	}
	model.Objects = objects
	if modelMap["indexing_policy"] != nil && len(modelMap["indexing_policy"].([]interface{})) > 0 {
		IndexingPolicyModel, err := resourceIbmProtectionGroupMapToIndexingPolicy(modelMap["indexing_policy"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IndexingPolicy = IndexingPolicyModel
	}
	if modelMap["perform_source_side_deduplication"] != nil {
		model.PerformSourceSideDeduplication = core.BoolPtr(modelMap["perform_source_side_deduplication"].(bool))
	}
	if modelMap["perform_brick_based_deduplication"] != nil {
		model.PerformBrickBasedDeduplication = core.BoolPtr(modelMap["perform_brick_based_deduplication"].(bool))
	}
	if modelMap["task_timeouts"] != nil {
		taskTimeouts := []backuprecoveryv1.CancellationTimeoutParams{}
		for _, taskTimeoutsItem := range modelMap["task_timeouts"].([]interface{}) {
			taskTimeoutsItemModel, err := resourceIbmProtectionGroupMapToCancellationTimeoutParams(taskTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			taskTimeouts = append(taskTimeouts, *taskTimeoutsItemModel)
		}
		model.TaskTimeouts = taskTimeouts
	}
	if modelMap["quiesce"] != nil {
		model.Quiesce = core.BoolPtr(modelMap["quiesce"].(bool))
	}
	if modelMap["continue_on_quiesce_failure"] != nil {
		model.ContinueOnQuiesceFailure = core.BoolPtr(modelMap["continue_on_quiesce_failure"].(bool))
	}
	if modelMap["cobmr_backup"] != nil {
		model.CobmrBackup = core.BoolPtr(modelMap["cobmr_backup"].(bool))
	}
	if modelMap["pre_post_script"] != nil && len(modelMap["pre_post_script"].([]interface{})) > 0 {
		PrePostScriptModel, err := resourceIbmProtectionGroupMapToPrePostScriptParams(modelMap["pre_post_script"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PrePostScript = PrePostScriptModel
	}
	if modelMap["dedup_exclusion_source_ids"] != nil {
		dedupExclusionSourceIds := []int64{}
		for _, dedupExclusionSourceIdsItem := range modelMap["dedup_exclusion_source_ids"].([]interface{}) {
			dedupExclusionSourceIds = append(dedupExclusionSourceIds, int64(dedupExclusionSourceIdsItem.(int)))
		}
		model.DedupExclusionSourceIds = dedupExclusionSourceIds
	}
	if modelMap["global_exclude_paths"] != nil {
		globalExcludePaths := []string{}
		for _, globalExcludePathsItem := range modelMap["global_exclude_paths"].([]interface{}) {
			globalExcludePaths = append(globalExcludePaths, globalExcludePathsItem.(string))
		}
		model.GlobalExcludePaths = globalExcludePaths
	}
	if modelMap["global_exclude_fs"] != nil {
		globalExcludeFs := []string{}
		for _, globalExcludeFsItem := range modelMap["global_exclude_fs"].([]interface{}) {
			globalExcludeFs = append(globalExcludeFs, globalExcludeFsItem.(string))
		}
		model.GlobalExcludeFS = globalExcludeFs
	}
	if modelMap["ignorable_errors"] != nil {
		ignorableErrors := []string{}
		for _, ignorableErrorsItem := range modelMap["ignorable_errors"].([]interface{}) {
			ignorableErrors = append(ignorableErrors, ignorableErrorsItem.(string))
		}
		model.IgnorableErrors = ignorableErrors
	}
	if modelMap["allow_parallel_runs"] != nil {
		model.AllowParallelRuns = core.BoolPtr(modelMap["allow_parallel_runs"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalFileProtectionGroupObjectParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalFileProtectionGroupObjectParams, error) {
	model := &backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["file_paths"] != nil {
		filePaths := []backuprecoveryv1.PhysicalFileBackupPathParams{}
		for _, filePathsItem := range modelMap["file_paths"].([]interface{}) {
			filePathsItemModel, err := resourceIbmProtectionGroupMapToPhysicalFileBackupPathParams(filePathsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			filePaths = append(filePaths, *filePathsItemModel)
		}
		model.FilePaths = filePaths
	}
	if modelMap["uses_path_level_skip_nested_volume_setting"] != nil {
		model.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(modelMap["uses_path_level_skip_nested_volume_setting"].(bool))
	}
	if modelMap["nested_volume_types_to_skip"] != nil {
		nestedVolumeTypesToSkip := []string{}
		for _, nestedVolumeTypesToSkipItem := range modelMap["nested_volume_types_to_skip"].([]interface{}) {
			nestedVolumeTypesToSkip = append(nestedVolumeTypesToSkip, nestedVolumeTypesToSkipItem.(string))
		}
		model.NestedVolumeTypesToSkip = nestedVolumeTypesToSkip
	}
	if modelMap["follow_nas_symlink_target"] != nil {
		model.FollowNasSymlinkTarget = core.BoolPtr(modelMap["follow_nas_symlink_target"].(bool))
	}
	if modelMap["metadata_file_path"] != nil && modelMap["metadata_file_path"].(string) != "" {
		model.MetadataFilePath = core.StringPtr(modelMap["metadata_file_path"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToPhysicalFileBackupPathParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalFileBackupPathParams, error) {
	model := &backuprecoveryv1.PhysicalFileBackupPathParams{}
	model.IncludedPath = core.StringPtr(modelMap["included_path"].(string))
	if modelMap["excluded_paths"] != nil {
		excludedPaths := []string{}
		for _, excludedPathsItem := range modelMap["excluded_paths"].([]interface{}) {
			excludedPaths = append(excludedPaths, excludedPathsItem.(string))
		}
		model.ExcludedPaths = excludedPaths
	}
	if modelMap["skip_nested_volumes"] != nil {
		model.SkipNestedVolumes = core.BoolPtr(modelMap["skip_nested_volumes"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToCancellationTimeoutParams(modelMap map[string]interface{}) (*backuprecoveryv1.CancellationTimeoutParams, error) {
	model := &backuprecoveryv1.CancellationTimeoutParams{}
	if modelMap["timeout_mins"] != nil {
		model.TimeoutMins = core.Int64Ptr(int64(modelMap["timeout_mins"].(int)))
	}
	if modelMap["backup_type"] != nil && modelMap["backup_type"].(string) != "" {
		model.BackupType = core.StringPtr(modelMap["backup_type"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleProtectionGroupParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleProtectionGroupParams, error) {
	model := &backuprecoveryv1.OracleProtectionGroupParams{}
	objects := []backuprecoveryv1.OracleProtectionGroupObjectParams{}
	for _, objectsItem := range modelMap["objects"].([]interface{}) {
		objectsItemModel, err := resourceIbmProtectionGroupMapToOracleProtectionGroupObjectParams(objectsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		objects = append(objects, *objectsItemModel)
	}
	model.Objects = objects
	if modelMap["persist_mountpoints"] != nil {
		model.PersistMountpoints = core.BoolPtr(modelMap["persist_mountpoints"].(bool))
	}
	if modelMap["vlan_params"] != nil && len(modelMap["vlan_params"].([]interface{})) > 0 {
		VlanParamsModel, err := resourceIbmProtectionGroupMapToVlanParams(modelMap["vlan_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.VlanParams = VlanParamsModel
	}
	if modelMap["pre_post_script"] != nil && len(modelMap["pre_post_script"].([]interface{})) > 0 {
		PrePostScriptModel, err := resourceIbmProtectionGroupMapToPrePostScriptParams(modelMap["pre_post_script"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PrePostScript = PrePostScriptModel
	}
	if modelMap["log_auto_kill_timeout_secs"] != nil {
		model.LogAutoKillTimeoutSecs = core.Int64Ptr(int64(modelMap["log_auto_kill_timeout_secs"].(int)))
	}
	if modelMap["incr_auto_kill_timeout_secs"] != nil {
		model.IncrAutoKillTimeoutSecs = core.Int64Ptr(int64(modelMap["incr_auto_kill_timeout_secs"].(int)))
	}
	if modelMap["full_auto_kill_timeout_secs"] != nil {
		model.FullAutoKillTimeoutSecs = core.Int64Ptr(int64(modelMap["full_auto_kill_timeout_secs"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleProtectionGroupObjectParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleProtectionGroupObjectParams, error) {
	model := &backuprecoveryv1.OracleProtectionGroupObjectParams{}
	model.SourceID = core.Int64Ptr(int64(modelMap["source_id"].(int)))
	if modelMap["source_name"] != nil && modelMap["source_name"].(string) != "" {
		model.SourceName = core.StringPtr(modelMap["source_name"].(string))
	}
	if modelMap["db_params"] != nil {
		dbParams := []backuprecoveryv1.OracleProtectionGroupDbParams{}
		for _, dbParamsItem := range modelMap["db_params"].([]interface{}) {
			dbParamsItemModel, err := resourceIbmProtectionGroupMapToOracleProtectionGroupDbParams(dbParamsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			dbParams = append(dbParams, *dbParamsItemModel)
		}
		model.DbParams = dbParams
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleProtectionGroupDbParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleProtectionGroupDbParams, error) {
	model := &backuprecoveryv1.OracleProtectionGroupDbParams{}
	if modelMap["database_id"] != nil {
		model.DatabaseID = core.Int64Ptr(int64(modelMap["database_id"].(int)))
	}
	if modelMap["database_name"] != nil && modelMap["database_name"].(string) != "" {
		model.DatabaseName = core.StringPtr(modelMap["database_name"].(string))
	}
	if modelMap["db_channels"] != nil {
		dbChannels := []backuprecoveryv1.OracleDbChannel{}
		for _, dbChannelsItem := range modelMap["db_channels"].([]interface{}) {
			dbChannelsItemModel, err := resourceIbmProtectionGroupMapToOracleDbChannel(dbChannelsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			dbChannels = append(dbChannels, *dbChannelsItemModel)
		}
		model.DbChannels = dbChannels
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleDbChannel(modelMap map[string]interface{}) (*backuprecoveryv1.OracleDbChannel, error) {
	model := &backuprecoveryv1.OracleDbChannel{}
	if modelMap["archive_log_retention_days"] != nil {
		model.ArchiveLogRetentionDays = core.Int64Ptr(int64(modelMap["archive_log_retention_days"].(int)))
	}
	if modelMap["archive_log_retention_hours"] != nil {
		model.ArchiveLogRetentionHours = core.Int64Ptr(int64(modelMap["archive_log_retention_hours"].(int)))
	}
	if modelMap["credentials"] != nil && len(modelMap["credentials"].([]interface{})) > 0 {
		CredentialsModel, err := resourceIbmProtectionGroupMapToCredentials(modelMap["credentials"].([]interface{})[0].(map[string]interface{}))
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
			databaseNodeListItemModel, err := resourceIbmProtectionGroupMapToOracleDatabaseHost(databaseNodeListItem.(map[string]interface{}))
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

func resourceIbmProtectionGroupMapToCredentials(modelMap map[string]interface{}) (*backuprecoveryv1.Credentials, error) {
	model := &backuprecoveryv1.Credentials{}
	model.Username = core.StringPtr(modelMap["username"].(string))
	model.Password = core.StringPtr(modelMap["password"].(string))
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleDatabaseHost(modelMap map[string]interface{}) (*backuprecoveryv1.OracleDatabaseHost, error) {
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
		SbtHostParamsModel, err := resourceIbmProtectionGroupMapToOracleSbtHostParams(modelMap["sbt_host_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SbtHostParams = SbtHostParamsModel
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleSbtHostParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleSbtHostParams, error) {
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
			vlanInfoListItemModel, err := resourceIbmProtectionGroupMapToOracleVlanInfo(vlanInfoListItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			vlanInfoList = append(vlanInfoList, *vlanInfoListItemModel)
		}
		model.VlanInfoList = vlanInfoList
	}
	return model, nil
}

func resourceIbmProtectionGroupMapToOracleVlanInfo(modelMap map[string]interface{}) (*backuprecoveryv1.OracleVlanInfo, error) {
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

func resourceIbmProtectionGroupMapToVlanParams(modelMap map[string]interface{}) (*backuprecoveryv1.VlanParams, error) {
	model := &backuprecoveryv1.VlanParams{}
	if modelMap["vlan_id"] != nil {
		model.VlanID = core.Int64Ptr(int64(modelMap["vlan_id"].(int)))
	}
	if modelMap["disable_vlan"] != nil {
		model.DisableVlan = core.BoolPtr(modelMap["disable_vlan"].(bool))
	}
	if modelMap["interface_name"] != nil && modelMap["interface_name"].(string) != "" {
		model.InterfaceName = core.StringPtr(modelMap["interface_name"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupTimeOfDayToMap(model *backuprecoveryv1.TimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["hour"] = flex.IntValue(model.Hour)
	modelMap["minute"] = flex.IntValue(model.Minute)
	if model.TimeZone != nil {
		modelMap["time_zone"] = model.TimeZone
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupProtectionGroupAlertingPolicyToMap(model *backuprecoveryv1.ProtectionGroupAlertingPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["backup_run_status"] = model.BackupRunStatus
	if model.AlertTargets != nil {
		alertTargets := []map[string]interface{}{}
		for _, alertTargetsItem := range model.AlertTargets {
			alertTargetsItemMap, err := resourceIbmProtectionGroupAlertTargetToMap(&alertTargetsItem)
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

func resourceIbmProtectionGroupAlertTargetToMap(model *backuprecoveryv1.AlertTarget) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupSlaRuleToMap(model *backuprecoveryv1.SlaRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.SlaMinutes != nil {
		modelMap["sla_minutes"] = flex.IntValue(model.SlaMinutes)
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmProtectionGroupPhysicalProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["protection_type"] = model.ProtectionType
	if model.VolumeProtectionTypeParams != nil {
		volumeProtectionTypeParamsMap, err := resourceIbmProtectionGroupPhysicalVolumeProtectionGroupParamsToMap(model.VolumeProtectionTypeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["volume_protection_type_params"] = []map[string]interface{}{volumeProtectionTypeParamsMap}
	}
	if model.FileProtectionTypeParams != nil {
		fileProtectionTypeParamsMap, err := resourceIbmProtectionGroupPhysicalFileProtectionGroupParamsToMap(model.FileProtectionTypeParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["file_protection_type_params"] = []map[string]interface{}{fileProtectionTypeParamsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupPhysicalVolumeProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalVolumeProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := resourceIbmProtectionGroupPhysicalVolumeProtectionGroupObjectParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	if model.IndexingPolicy != nil {
		indexingPolicyMap, err := resourceIbmProtectionGroupIndexingPolicyToMap(model.IndexingPolicy)
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
		prePostScriptMap, err := resourceIbmProtectionGroupPrePostScriptParamsToMap(model.PrePostScript)
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

func resourceIbmProtectionGroupPhysicalVolumeProtectionGroupObjectParamsToMap(model *backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupIndexingPolicyToMap(model *backuprecoveryv1.IndexingPolicy) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupPrePostScriptParamsToMap(model *backuprecoveryv1.PrePostScriptParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PreScript != nil {
		preScriptMap, err := resourceIbmProtectionGroupCommonPreBackupScriptParamsToMap(model.PreScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["pre_script"] = []map[string]interface{}{preScriptMap}
	}
	if model.PostScript != nil {
		postScriptMap, err := resourceIbmProtectionGroupCommonPostBackupScriptParamsToMap(model.PostScript)
		if err != nil {
			return modelMap, err
		}
		modelMap["post_script"] = []map[string]interface{}{postScriptMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupCommonPreBackupScriptParamsToMap(model *backuprecoveryv1.CommonPreBackupScriptParams) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupCommonPostBackupScriptParamsToMap(model *backuprecoveryv1.CommonPostBackupScriptParams) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupPhysicalFileProtectionGroupParamsToMap(model *backuprecoveryv1.PhysicalFileProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := resourceIbmProtectionGroupPhysicalFileProtectionGroupObjectParamsToMap(&objectsItem)
		if err != nil {
			return modelMap, err
		}
		objects = append(objects, objectsItemMap)
	}
	modelMap["objects"] = objects
	if model.IndexingPolicy != nil {
		indexingPolicyMap, err := resourceIbmProtectionGroupIndexingPolicyToMap(model.IndexingPolicy)
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
			taskTimeoutsItemMap, err := resourceIbmProtectionGroupCancellationTimeoutParamsToMap(&taskTimeoutsItem)
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
		prePostScriptMap, err := resourceIbmProtectionGroupPrePostScriptParamsToMap(model.PrePostScript)
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

func resourceIbmProtectionGroupPhysicalFileProtectionGroupObjectParamsToMap(model *backuprecoveryv1.PhysicalFileProtectionGroupObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.FilePaths != nil {
		filePaths := []map[string]interface{}{}
		for _, filePathsItem := range model.FilePaths {
			filePathsItemMap, err := resourceIbmProtectionGroupPhysicalFileBackupPathParamsToMap(&filePathsItem)
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

func resourceIbmProtectionGroupPhysicalFileBackupPathParamsToMap(model *backuprecoveryv1.PhysicalFileBackupPathParams) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupCancellationTimeoutParamsToMap(model *backuprecoveryv1.CancellationTimeoutParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TimeoutMins != nil {
		modelMap["timeout_mins"] = flex.IntValue(model.TimeoutMins)
	}
	if model.BackupType != nil {
		modelMap["backup_type"] = model.BackupType
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleProtectionGroupParamsToMap(model *backuprecoveryv1.OracleProtectionGroupParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	objects := []map[string]interface{}{}
	for _, objectsItem := range model.Objects {
		objectsItemMap, err := resourceIbmProtectionGroupOracleProtectionGroupObjectParamsToMap(&objectsItem)
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
		vlanParamsMap, err := resourceIbmProtectionGroupVlanParamsToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	if model.PrePostScript != nil {
		prePostScriptMap, err := resourceIbmProtectionGroupPrePostScriptParamsToMap(model.PrePostScript)
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

func resourceIbmProtectionGroupOracleProtectionGroupObjectParamsToMap(model *backuprecoveryv1.OracleProtectionGroupObjectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_id"] = flex.IntValue(model.SourceID)
	if model.SourceName != nil {
		modelMap["source_name"] = model.SourceName
	}
	if model.DbParams != nil {
		dbParams := []map[string]interface{}{}
		for _, dbParamsItem := range model.DbParams {
			dbParamsItemMap, err := resourceIbmProtectionGroupOracleProtectionGroupDbParamsToMap(&dbParamsItem)
			if err != nil {
				return modelMap, err
			}
			dbParams = append(dbParams, dbParamsItemMap)
		}
		modelMap["db_params"] = dbParams
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleProtectionGroupDbParamsToMap(model *backuprecoveryv1.OracleProtectionGroupDbParams) (map[string]interface{}, error) {
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
			dbChannelsItemMap, err := resourceIbmProtectionGroupOracleDbChannelToMap(&dbChannelsItem)
			if err != nil {
				return modelMap, err
			}
			dbChannels = append(dbChannels, dbChannelsItemMap)
		}
		modelMap["db_channels"] = dbChannels
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleDbChannelToMap(model *backuprecoveryv1.OracleDbChannel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchiveLogRetentionDays != nil {
		modelMap["archive_log_retention_days"] = flex.IntValue(model.ArchiveLogRetentionDays)
	}
	if model.ArchiveLogRetentionHours != nil {
		modelMap["archive_log_retention_hours"] = flex.IntValue(model.ArchiveLogRetentionHours)
	}
	if model.Credentials != nil {
		credentialsMap, err := resourceIbmProtectionGroupCredentialsToMap(model.Credentials)
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
			databaseNodeListItemMap, err := resourceIbmProtectionGroupOracleDatabaseHostToMap(&databaseNodeListItem)
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

func resourceIbmProtectionGroupCredentialsToMap(model *backuprecoveryv1.Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = model.Username
	modelMap["password"] = model.Password
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleDatabaseHostToMap(model *backuprecoveryv1.OracleDatabaseHost) (map[string]interface{}, error) {
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
		sbtHostParamsMap, err := resourceIbmProtectionGroupOracleSbtHostParamsToMap(model.SbtHostParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sbt_host_params"] = []map[string]interface{}{sbtHostParamsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleSbtHostParamsToMap(model *backuprecoveryv1.OracleSbtHostParams) (map[string]interface{}, error) {
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
			vlanInfoListItemMap, err := resourceIbmProtectionGroupOracleVlanInfoToMap(&vlanInfoListItem)
			if err != nil {
				return modelMap, err
			}
			vlanInfoList = append(vlanInfoList, vlanInfoListItemMap)
		}
		modelMap["vlan_info_list"] = vlanInfoList
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleVlanInfoToMap(model *backuprecoveryv1.OracleVlanInfo) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupVlanParamsToMap(model *backuprecoveryv1.VlanParams) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupProtectionGroupRunToMap(model *backuprecoveryv1.ProtectionGroupRun) (map[string]interface{}, error) {
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
		originClusterIdentifierMap, err := resourceIbmProtectionGroupClusterIdentifierToMap(model.OriginClusterIdentifier)
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
			objectsItemMap, err := resourceIbmProtectionGroupObjectRunResultToMap(&objectsItem)
			if err != nil {
				return modelMap, err
			}
			objects = append(objects, objectsItemMap)
		}
		modelMap["objects"] = objects
	}
	if model.LocalBackupInfo != nil {
		localBackupInfoMap, err := resourceIbmProtectionGroupBackupRunSummaryToMap(model.LocalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_backup_info"] = []map[string]interface{}{localBackupInfoMap}
	}
	if model.OriginalBackupInfo != nil {
		originalBackupInfoMap, err := resourceIbmProtectionGroupBackupRunSummaryToMap(model.OriginalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_backup_info"] = []map[string]interface{}{originalBackupInfoMap}
	}
	if model.ReplicationInfo != nil {
		replicationInfoMap, err := resourceIbmProtectionGroupReplicationRunSummaryToMap(model.ReplicationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["replication_info"] = []map[string]interface{}{replicationInfoMap}
	}
	if model.ArchivalInfo != nil {
		archivalInfoMap, err := resourceIbmProtectionGroupArchivalRunSummaryToMap(model.ArchivalInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_info"] = []map[string]interface{}{archivalInfoMap}
	}
	if model.CloudSpinInfo != nil {
		cloudSpinInfoMap, err := resourceIbmProtectionGroupCloudSpinRunSummaryToMap(model.CloudSpinInfo)
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
			permissionsItemMap, err := resourceIbmProtectionGroupTenantToMap(&permissionsItem)
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

func resourceIbmProtectionGroupClusterIdentifierToMap(model *backuprecoveryv1.ClusterIdentifier) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupObjectRunResultToMap(model *backuprecoveryv1.ObjectRunResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Object != nil {
		objectMap, err := resourceIbmProtectionGroupObjectSummaryToMap(model.Object)
		if err != nil {
			return modelMap, err
		}
		modelMap["object"] = []map[string]interface{}{objectMap}
	}
	if model.LocalSnapshotInfo != nil {
		localSnapshotInfoMap, err := resourceIbmProtectionGroupBackupRunToMap(model.LocalSnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_snapshot_info"] = []map[string]interface{}{localSnapshotInfoMap}
	}
	if model.OriginalBackupInfo != nil {
		originalBackupInfoMap, err := resourceIbmProtectionGroupBackupRunToMap(model.OriginalBackupInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["original_backup_info"] = []map[string]interface{}{originalBackupInfoMap}
	}
	if model.ReplicationInfo != nil {
		replicationInfoMap, err := resourceIbmProtectionGroupReplicationRunToMap(model.ReplicationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["replication_info"] = []map[string]interface{}{replicationInfoMap}
	}
	if model.ArchivalInfo != nil {
		archivalInfoMap, err := resourceIbmProtectionGroupArchivalRunToMap(model.ArchivalInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_info"] = []map[string]interface{}{archivalInfoMap}
	}
	if model.CloudSpinInfo != nil {
		cloudSpinInfoMap, err := resourceIbmProtectionGroupCloudSpinRunToMap(model.CloudSpinInfo)
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

func resourceIbmProtectionGroupObjectSummaryToMap(model *backuprecoveryv1.ObjectSummary) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupBackupRunToMap(model *backuprecoveryv1.BackupRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SnapshotInfo != nil {
		snapshotInfoMap, err := resourceIbmProtectionGroupSnapshotInfoToMap(model.SnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["snapshot_info"] = []map[string]interface{}{snapshotInfoMap}
	}
	if model.FailedAttempts != nil {
		failedAttempts := []map[string]interface{}{}
		for _, failedAttemptsItem := range model.FailedAttempts {
			failedAttemptsItemMap, err := resourceIbmProtectionGroupBackupAttemptToMap(&failedAttemptsItem)
			if err != nil {
				return modelMap, err
			}
			failedAttempts = append(failedAttempts, failedAttemptsItemMap)
		}
		modelMap["failed_attempts"] = failedAttempts
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupSnapshotInfoToMap(model *backuprecoveryv1.SnapshotInfo) (map[string]interface{}, error) {
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
		statsMap, err := resourceIbmProtectionGroupBackupDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := resourceIbmProtectionGroupDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupBackupDataStatsToMap(model *backuprecoveryv1.BackupDataStats) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupDataLockConstraintsToMap(model *backuprecoveryv1.DataLockConstraints) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Mode != nil {
		modelMap["mode"] = model.Mode
	}
	if model.ExpiryTimeUsecs != nil {
		modelMap["expiry_time_usecs"] = flex.IntValue(model.ExpiryTimeUsecs)
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupBackupAttemptToMap(model *backuprecoveryv1.BackupAttempt) (map[string]interface{}, error) {
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
		statsMap, err := resourceIbmProtectionGroupBackupDataStatsToMap(model.Stats)
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

func resourceIbmProtectionGroupReplicationRunToMap(model *backuprecoveryv1.ReplicationRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := resourceIbmProtectionGroupReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupReplicationTargetResultToMap(model *backuprecoveryv1.ReplicationTargetResult) (map[string]interface{}, error) {
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
		statsMap, err := resourceIbmProtectionGroupReplicationDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := resourceIbmProtectionGroupDataLockConstraintsToMap(model.DataLockConstraints)
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

func resourceIbmProtectionGroupReplicationDataStatsToMap(model *backuprecoveryv1.ReplicationDataStats) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupArchivalRunToMap(model *backuprecoveryv1.ArchivalRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := resourceIbmProtectionGroupArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupArchivalTargetResultToMap(model *backuprecoveryv1.ArchivalTargetResult) (map[string]interface{}, error) {
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
		tierSettingsMap, err := resourceIbmProtectionGroupArchivalTargetTierInfoToMap(model.TierSettings)
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
		statsMap, err := resourceIbmProtectionGroupArchivalDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := resourceIbmProtectionGroupDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	if model.OnLegalHold != nil {
		modelMap["on_legal_hold"] = model.OnLegalHold
	}
	if model.WormProperties != nil {
		wormPropertiesMap, err := resourceIbmProtectionGroupWormPropertiesToMap(model.WormProperties)
		if err != nil {
			return modelMap, err
		}
		modelMap["worm_properties"] = []map[string]interface{}{wormPropertiesMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := resourceIbmProtectionGroupOracleTiersToMap(model.OracleTiering)
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

func resourceIbmProtectionGroupOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := resourceIbmProtectionGroupOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func resourceIbmProtectionGroupOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupArchivalDataStatsToMap(model *backuprecoveryv1.ArchivalDataStats) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupWormPropertiesToMap(model *backuprecoveryv1.WormProperties) (map[string]interface{}, error) {
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

func resourceIbmProtectionGroupCloudSpinRunToMap(model *backuprecoveryv1.CloudSpinRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := resourceIbmProtectionGroupCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupCloudSpinTargetResultToMap(model *backuprecoveryv1.CloudSpinTargetResult) (map[string]interface{}, error) {
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
		statsMap, err := resourceIbmProtectionGroupCloudSpinDataStatsToMap(model.Stats)
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
		dataLockConstraintsMap, err := resourceIbmProtectionGroupDataLockConstraintsToMap(model.DataLockConstraints)
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

func resourceIbmProtectionGroupCloudSpinDataStatsToMap(model *backuprecoveryv1.CloudSpinDataStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PhysicalBytesTransferred != nil {
		modelMap["physical_bytes_transferred"] = flex.IntValue(model.PhysicalBytesTransferred)
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupBackupRunSummaryToMap(model *backuprecoveryv1.BackupRunSummary) (map[string]interface{}, error) {
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
		localSnapshotStatsMap, err := resourceIbmProtectionGroupBackupDataStatsToMap(model.LocalSnapshotStats)
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
		dataLockConstraintsMap, err := resourceIbmProtectionGroupDataLockConstraintsToMap(model.DataLockConstraints)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupReplicationRunSummaryToMap(model *backuprecoveryv1.ReplicationRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargetResults != nil {
		replicationTargetResults := []map[string]interface{}{}
		for _, replicationTargetResultsItem := range model.ReplicationTargetResults {
			replicationTargetResultsItemMap, err := resourceIbmProtectionGroupReplicationTargetResultToMap(&replicationTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargetResults = append(replicationTargetResults, replicationTargetResultsItemMap)
		}
		modelMap["replication_target_results"] = replicationTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupArchivalRunSummaryToMap(model *backuprecoveryv1.ArchivalRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ArchivalTargetResults != nil {
		archivalTargetResults := []map[string]interface{}{}
		for _, archivalTargetResultsItem := range model.ArchivalTargetResults {
			archivalTargetResultsItemMap, err := resourceIbmProtectionGroupArchivalTargetResultToMap(&archivalTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargetResults = append(archivalTargetResults, archivalTargetResultsItemMap)
		}
		modelMap["archival_target_results"] = archivalTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupCloudSpinRunSummaryToMap(model *backuprecoveryv1.CloudSpinRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CloudSpinTargetResults != nil {
		cloudSpinTargetResults := []map[string]interface{}{}
		for _, cloudSpinTargetResultsItem := range model.CloudSpinTargetResults {
			cloudSpinTargetResultsItemMap, err := resourceIbmProtectionGroupCloudSpinTargetResultToMap(&cloudSpinTargetResultsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargetResults = append(cloudSpinTargetResults, cloudSpinTargetResultsItemMap)
		}
		modelMap["cloud_spin_target_results"] = cloudSpinTargetResults
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupMissingEntityParamsToMap(model *backuprecoveryv1.MissingEntityParams) (map[string]interface{}, error) {
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
