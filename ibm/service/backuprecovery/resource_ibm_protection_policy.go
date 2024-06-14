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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv0"
)

func ResourceIbmProtectionPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProtectionPolicyCreate,
		ReadContext:   resourceIbmProtectionPolicyRead,
		UpdateContext: resourceIbmProtectionPolicyUpdate,
		DeleteContext: resourceIbmProtectionPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the name of the Protection Policy.",
			},
			"backup_policy": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Specifies the backup schedule and retentions of a Protection Policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regular": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Specifies the Incremental and Full policy settings and also the common Retention policy settings.\".",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"incremental": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies incremental backup settings for a Protection Group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies settings that defines how frequent backup will be performed for a Protection Group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies how often to start new runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field. <br>'Days' specifies that Protection Group run starts periodically after certain number of days specified in 'frequency' field. <br>'Week' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Month' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.",
															},
															"minute_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
																		},
																	},
																},
															},
															"hour_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
																		},
																	},
																},
															},
															"day_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
																		},
																	},
																},
															},
															"week_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Required:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"month_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"week_of_month": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.",
																		},
																		"day_of_month": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.",
																		},
																	},
																},
															},
															"year_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_year": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.",
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
									"full": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies full backup settings for a Protection Group. Currently, full backup settings can be specified by using either of 'schedule' or 'schdulesAndRetentions' field. Using 'schdulesAndRetentions' is recommended when multiple full backups need to be configured. If full and incremental backup has common retention then only setting 'schedule' is recommended.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that defines how frequent full backup will be performed for a Protection Group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.",
															},
															"day_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
																		},
																	},
																},
															},
															"week_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Required:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"month_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"week_of_month": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.",
																		},
																		"day_of_month": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.",
																		},
																	},
																},
															},
															"year_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_year": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.",
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
									"full_backups": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies multiple schedules and retentions for full backup. Specify either of the 'full' or 'fullBackups' values. Its recommended to use 'fullBaackups' value since 'full' will be deprecated after few releases.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies settings that defines how frequent full backup will be performed for a Protection Group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.",
															},
															"day_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
																		},
																	},
																},
															},
															"week_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Required:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"month_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_week": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"week_of_month": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.",
																		},
																		"day_of_month": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.",
																		},
																	},
																},
															},
															"year_schedule": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"day_of_year": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.",
																		},
																	},
																},
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
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
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"primary_backup_target": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the primary backup target settings for regular backups. If the backup target field is not specified then backup will be taken locally on the Cohesity cluster.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "Local",
													Description: "Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster.",
												},
												"archival_target_settings": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the primary archival settings. Mainly used for cloud direct archive (CAD) policy where primary backup is stored on archival target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"target_id": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the Archival target id to take primary backup.",
															},
															"target_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the Archival target name where Snapshots are copied.",
															},
															"tier_settings": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].",
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
																									Optional:    true,
																									Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																								},
																								"move_after": &schema.Schema{
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																								},
																								"tier_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Required:    true,
																									Description: "Specifies the Oracle tier types.",
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
										},
									},
								},
							},
						},
						"log": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies log backup settings for a Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies settings that defines how frequent log backup will be performed for a Protection Group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.",
												},
												"minute_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
												"hour_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
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
						"bmr": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the BMR schedule in case of physical source protection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies settings that defines how frequent bmr backup will be performed for a Protection Group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies how often to start new runs of a Protection Group. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week.",
												},
												"day_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
												"week_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_week": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
												"month_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_week": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"week_of_month": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.",
															},
															"day_of_month": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.",
															},
														},
													},
												},
												"year_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_year": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.",
															},
														},
													},
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
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
						"cdp": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies CDP (Continious Data Protection) backup settings for a Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a CDP backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a CDP backup measured in minutes or hours.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a cdp backup retention.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
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
						"storage_array_snapshot": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies storage snapshot managment backup settings for a Protection Group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies settings that defines how frequent Storage Snapshot Management backup will be performed for a Protection Group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.",
												},
												"minute_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
												"hour_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
												"day_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.",
															},
														},
													},
												},
												"week_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_week": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
												"month_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_week": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"week_of_month": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.",
															},
															"day_of_month": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.",
															},
														},
													},
												},
												"year_schedule": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_year": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.",
															},
														},
													},
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
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
						"run_timeouts": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the backup timeouts for different type of runs(kFull, kRegular etc.).",
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
					},
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the description of the Protection Policy.",
			},
			"blackout_window": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies a day in the week when no new Protection Group Runs should be started such as 'Sunday'. Specifies a day in a week such as 'Sunday', 'Monday', etc.",
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
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
						"end_time": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
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
						"config_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.",
						},
					},
				},
			},
			"extended_retention": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Specifies a schedule frequency and schedule unit for Extended Retentions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"unit": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.",
									},
									"frequency": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.",
									},
								},
							},
						},
						"retention": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Specifies the retention of a backup.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"unit": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
									},
									"duration": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
									},
									"data_lock_config": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
												},
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
												},
												"enable_worm_on_external_target": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
												},
											},
										},
									},
								},
							},
						},
						"run_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.",
						},
						"config_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.",
						},
					},
				},
			},
			"remote_target_policy": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Specifies the replication, archival and cloud spin targets of Protection Policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"replication_targets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_on_run_success": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
									},
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
									},
									"run_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
									"log_retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"target_type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the type of target to which replication need to be performed.",
									},
									"remote_target_config": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the configuration for adding remote cluster as repilcation target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_id": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the cluster id of the target replication cluster.",
												},
												"cluster_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the cluster name of the target replication cluster.",
												},
											},
										},
									},
								},
							},
						},
						"archival_targets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_on_run_success": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
									},
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
									},
									"run_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
									"log_retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"target_id": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies the Archival target to copy the Snapshots to.",
									},
									"target_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the Archival target name where Snapshots are copied.",
									},
									"target_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the Archival target type where Snapshots are copied.",
									},
									"tier_settings": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].",
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
																			Optional:    true,
																			Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																		},
																		"move_after": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																		},
																		"tier_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the Oracle tier types.",
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
									"extended_retention": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule frequency and schedule unit for Extended Retentions.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.",
												},
											},
										},
									},
								},
							},
						},
						"cloud_spin_targets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_on_run_success": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
									},
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
									},
									"run_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
									"log_retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"target": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the unique id of the cloud spin entity.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the already added cloud spin target.",
												},
											},
										},
									},
								},
							},
						},
						"onprem_deploy_targets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_on_run_success": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
									},
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
									},
									"run_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
									"log_retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"params": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies the unique id of the onprem entity.",
												},
												"restore_v_mware_params": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the parameters for a VMware recovery target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"target_vm_folder_id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the folder ID where the VMs should be created.",
															},
															"target_data_store_id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the folder where the restore datastore should be created.",
															},
															"enable_copy_recovery": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether to perform copy recovery or not.",
															},
															"resource_pool_id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies if the restore is to alternate location.",
															},
															"datastore_ids": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Specifies Datastore Ids, if the restore is to alternate location.",
																Elem:        &schema.Schema{Type: schema.TypeInt},
															},
															"overwrite_existing_vm": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether to overwrite the VM at the target location.",
															},
															"power_off_and_rename_existing_vm": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether to power off and mark the VM at the target location as deprecated.",
															},
															"attempt_differential_restore": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether to attempt differential restore.",
															},
															"is_on_prem_deploy": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether a task in on prem deploy or not.",
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
						"rpaas_targets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_on_run_success": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
									},
									"backup_run_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
									},
									"run_timeouts": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
									"log_retention": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"target_id": &schema.Schema{
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies the RPaaS target to copy the Snapshots.",
									},
									"target_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the RPaaS target name where Snapshots are copied.",
									},
									"target_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the RPaaS target type where Snapshots are copied.",
									},
								},
							},
						},
					},
				},
			},
			"cascaded_targets_config": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the source cluster id from where the remote operations will be performed to the next set of remote targets.",
						},
						"remote_targets": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Specifies the replication, archival and cloud spin targets of Protection Policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"replication_targets": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"copy_on_run_success": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
												},
												"backup_run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
												},
												"run_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
												"log_retention": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"target_type": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specifies the type of target to which replication need to be performed.",
												},
												"remote_target_config": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the configuration for adding remote cluster as repilcation target.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_id": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the cluster id of the target replication cluster.",
															},
															"cluster_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the cluster name of the target replication cluster.",
															},
														},
													},
												},
											},
										},
									},
									"archival_targets": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"copy_on_run_success": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
												},
												"backup_run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
												},
												"run_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
												"log_retention": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"target_id": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the Archival target to copy the Snapshots to.",
												},
												"target_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Archival target name where Snapshots are copied.",
												},
												"target_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Archival target type where Snapshots are copied.",
												},
												"tier_settings": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].",
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
																						Optional:    true,
																						Description: "Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.",
																					},
																					"move_after": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Specifies the time period after which the backup will be moved from current tier to next tier.",
																					},
																					"tier_type": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Specifies the Oracle tier types.",
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
												"extended_retention": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"schedule": &schema.Schema{
																Type:        schema.TypeList,
																MinItems:    1,
																MaxItems:    1,
																Required:    true,
																Description: "Specifies a schedule frequency and schedule unit for Extended Retentions.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.",
																		},
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.",
																		},
																	},
																},
															},
															"retention": &schema.Schema{
																Type:        schema.TypeList,
																MinItems:    1,
																MaxItems:    1,
																Required:    true,
																Description: "Specifies the retention of a backup.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
																		},
																		"data_lock_config": &schema.Schema{
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mode": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																					},
																					"unit": &schema.Schema{
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																					},
																					"duration": &schema.Schema{
																						Type:        schema.TypeInt,
																						Required:    true,
																						Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																					},
																					"enable_worm_on_external_target": &schema.Schema{
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"run_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.",
															},
															"config_id": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.",
															},
														},
													},
												},
											},
										},
									},
									"cloud_spin_targets": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"copy_on_run_success": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
												},
												"backup_run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
												},
												"run_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
												"log_retention": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"target": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the unique id of the cloud spin entity.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the already added cloud spin target.",
															},
														},
													},
												},
											},
										},
									},
									"onprem_deploy_targets": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"copy_on_run_success": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
												},
												"backup_run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
												},
												"run_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
												"log_retention": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"params": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies the unique id of the onprem entity.",
															},
															"restore_v_mware_params": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies the parameters for a VMware recovery target.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"target_vm_folder_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the folder ID where the VMs should be created.",
																		},
																		"target_data_store_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies the folder where the restore datastore should be created.",
																		},
																		"enable_copy_recovery": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to perform copy recovery or not.",
																		},
																		"resource_pool_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Specifies if the restore is to alternate location.",
																		},
																		"datastore_ids": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Specifies Datastore Ids, if the restore is to alternate location.",
																			Elem:        &schema.Schema{Type: schema.TypeInt},
																		},
																		"overwrite_existing_vm": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to overwrite the VM at the target location.",
																		},
																		"power_off_and_rename_existing_vm": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to power off and mark the VM at the target location as deprecated.",
																		},
																		"attempt_differential_restore": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether to attempt differential restore.",
																		},
																		"is_on_prem_deploy": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether a task in on prem deploy or not.",
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
									"rpaas_targets": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.",
															},
															"frequency": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.",
															},
														},
													},
												},
												"retention": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"copy_on_run_success": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.",
												},
												"config_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.",
												},
												"backup_run_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.",
												},
												"run_timeouts": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).",
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
												"log_retention": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies the retention of a backup.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
															},
															"data_lock_config": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mode": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
																		},
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
																		},
																		"duration": &schema.Schema{
																			Type:        schema.TypeInt,
																			Required:    true,
																			Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
																		},
																		"enable_worm_on_external_target": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
																		},
																	},
																},
															},
														},
													},
												},
												"target_id": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the RPaaS target to copy the Snapshots.",
												},
												"target_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the RPaaS target name where Snapshots are copied.",
												},
												"target_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the RPaaS target type where Snapshots are copied.",
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
			"retry_options": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Retry Options of a Protection Policy when a Protection Group run fails.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retries": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails.",
						},
						"retry_interval_mins": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the number of minutes before retrying a failed Protection Group.",
						},
					},
				},
			},
			"data_lock": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field is now deprecated. Please use the DataLockConfig in the backup retention.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases.",
			},
			"is_cbs_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature.",
			},
			"last_modification_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.",
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the parent policy template id to which the policy is linked to.",
			},
		},
	}
}

func ResourceIbmProtectionPolicyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "data_lock",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "Administrative, Compliance",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_protection_policy", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProtectionPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	createProtectionPolicyOptions := &backuprecoveryv0.CreateProtectionPolicyOptions{}

	createProtectionPolicyOptions.SetName(d.Get("name").(string))
	backupPolicyModel, err := resourceIbmProtectionPolicyMapToBackupPolicy(d.Get("backup_policy.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createProtectionPolicyOptions.SetBackupPolicy(backupPolicyModel)
	if _, ok := d.GetOk("description"); ok {
		createProtectionPolicyOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("blackout_window"); ok {
		var blackoutWindow []backuprecoveryv0.BlackoutWindow
		for _, v := range d.Get("blackout_window").([]interface{}) {
			value := v.(map[string]interface{})
			blackoutWindowItem, err := resourceIbmProtectionPolicyMapToBlackoutWindow(value)
			if err != nil {
				return diag.FromErr(err)
			}
			blackoutWindow = append(blackoutWindow, *blackoutWindowItem)
		}
		createProtectionPolicyOptions.SetBlackoutWindow(blackoutWindow)
	}
	if _, ok := d.GetOk("extended_retention"); ok {
		var extendedRetention []backuprecoveryv0.ExtendedRetentionPolicy
		for _, v := range d.Get("extended_retention").([]interface{}) {
			value := v.(map[string]interface{})
			extendedRetentionItem, err := resourceIbmProtectionPolicyMapToExtendedRetentionPolicy(value)
			if err != nil {
				return diag.FromErr(err)
			}
			extendedRetention = append(extendedRetention, *extendedRetentionItem)
		}
		createProtectionPolicyOptions.SetExtendedRetention(extendedRetention)
	}
	if _, ok := d.GetOk("remote_target_policy"); ok {
		remoteTargetPolicyModel, err := resourceIbmProtectionPolicyMapToTargetsConfiguration(d.Get("remote_target_policy.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionPolicyOptions.SetRemoteTargetPolicy(remoteTargetPolicyModel)
	}
	if _, ok := d.GetOk("cascaded_targets_config"); ok {
		var cascadedTargetsConfig []backuprecoveryv0.CascadedTargetConfiguration
		for _, v := range d.Get("cascaded_targets_config").([]interface{}) {
			value := v.(map[string]interface{})
			cascadedTargetsConfigItem, err := resourceIbmProtectionPolicyMapToCascadedTargetConfiguration(value)
			if err != nil {
				return diag.FromErr(err)
			}
			cascadedTargetsConfig = append(cascadedTargetsConfig, *cascadedTargetsConfigItem)
		}
		createProtectionPolicyOptions.SetCascadedTargetsConfig(cascadedTargetsConfig)
	}
	if _, ok := d.GetOk("retry_options"); ok {
		retryOptionsModel, err := resourceIbmProtectionPolicyMapToRetryOptions(d.Get("retry_options.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionPolicyOptions.SetRetryOptions(retryOptionsModel)
	}
	if _, ok := d.GetOk("data_lock"); ok {
		createProtectionPolicyOptions.SetDataLock(d.Get("data_lock").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		createProtectionPolicyOptions.SetVersion(int64(d.Get("version").(int)))
	}
	if _, ok := d.GetOk("is_cbs_enabled"); ok {
		createProtectionPolicyOptions.SetIsCBSEnabled(d.Get("is_cbs_enabled").(bool))
	}
	if _, ok := d.GetOk("last_modification_time_usecs"); ok {
		createProtectionPolicyOptions.SetLastModificationTimeUsecs(int64(d.Get("last_modification_time_usecs").(int)))
	}
	if _, ok := d.GetOk("template_id"); ok {
		createProtectionPolicyOptions.SetTemplateID(d.Get("template_id").(string))
	}

	protectionPolicyResponse, response, err := backupRecoveryClient.CreateProtectionPolicyWithContext(context, createProtectionPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateProtectionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateProtectionPolicyWithContext failed %s\n%s", err, response))
	}

	d.SetId(*protectionPolicyResponse.ID)

	return resourceIbmProtectionPolicyRead(context, d, meta)
}

func resourceIbmProtectionPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionPolicyByIdOptions := &backuprecoveryv0.GetProtectionPolicyByIdOptions{}

	getProtectionPolicyByIdOptions.SetID(d.Id())

	protectionPolicyResponse, response, err := backupRecoveryClient.GetProtectionPolicyByIDWithContext(context, getProtectionPolicyByIdOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProtectionPolicyByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionPolicyByIDWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", protectionPolicyResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	backupPolicyMap, err := resourceIbmProtectionPolicyBackupPolicyToMap(protectionPolicyResponse.BackupPolicy)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("backup_policy", []map[string]interface{}{backupPolicyMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting backup_policy: %s", err))
	}
	if !core.IsNil(protectionPolicyResponse.Description) {
		if err = d.Set("description", protectionPolicyResponse.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.BlackoutWindow) {
		blackoutWindow := []map[string]interface{}{}
		for _, blackoutWindowItem := range protectionPolicyResponse.BlackoutWindow {
			blackoutWindowItemMap, err := resourceIbmProtectionPolicyBlackoutWindowToMap(&blackoutWindowItem)
			if err != nil {
				return diag.FromErr(err)
			}
			blackoutWindow = append(blackoutWindow, blackoutWindowItemMap)
		}
		if err = d.Set("blackout_window", blackoutWindow); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting blackout_window: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.ExtendedRetention) {
		extendedRetention := []map[string]interface{}{}
		for _, extendedRetentionItem := range protectionPolicyResponse.ExtendedRetention {
			extendedRetentionItemMap, err := resourceIbmProtectionPolicyExtendedRetentionPolicyToMap(&extendedRetentionItem)
			if err != nil {
				return diag.FromErr(err)
			}
			extendedRetention = append(extendedRetention, extendedRetentionItemMap)
		}
		if err = d.Set("extended_retention", extendedRetention); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting extended_retention: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.RemoteTargetPolicy) {
		remoteTargetPolicyMap, err := resourceIbmProtectionPolicyTargetsConfigurationToMap(protectionPolicyResponse.RemoteTargetPolicy)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("remote_target_policy", []map[string]interface{}{remoteTargetPolicyMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting remote_target_policy: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.CascadedTargetsConfig) {
		cascadedTargetsConfig := []map[string]interface{}{}
		for _, cascadedTargetsConfigItem := range protectionPolicyResponse.CascadedTargetsConfig {
			cascadedTargetsConfigItemMap, err := resourceIbmProtectionPolicyCascadedTargetConfigurationToMap(&cascadedTargetsConfigItem)
			if err != nil {
				return diag.FromErr(err)
			}
			cascadedTargetsConfig = append(cascadedTargetsConfig, cascadedTargetsConfigItemMap)
		}
		if err = d.Set("cascaded_targets_config", cascadedTargetsConfig); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cascaded_targets_config: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.RetryOptions) {
		retryOptionsMap, err := resourceIbmProtectionPolicyRetryOptionsToMap(protectionPolicyResponse.RetryOptions)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("retry_options", []map[string]interface{}{retryOptionsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting retry_options: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.DataLock) {
		if err = d.Set("data_lock", protectionPolicyResponse.DataLock); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting data_lock: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.Version) {
		if err = d.Set("version", flex.IntValue(protectionPolicyResponse.Version)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.IsCBSEnabled) {
		if err = d.Set("is_cbs_enabled", protectionPolicyResponse.IsCBSEnabled); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_cbs_enabled: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.LastModificationTimeUsecs) {
		if err = d.Set("last_modification_time_usecs", flex.IntValue(protectionPolicyResponse.LastModificationTimeUsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_modification_time_usecs: %s", err))
		}
	}
	if !core.IsNil(protectionPolicyResponse.TemplateID) {
		if err = d.Set("template_id", protectionPolicyResponse.TemplateID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting template_id: %s", err))
		}
	}

	return nil
}

func resourceIbmProtectionPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProtectionPolicyOptions := &backuprecoveryv0.UpdateProtectionPolicyOptions{}

	updateProtectionPolicyOptions.SetID(d.Id())
	updateProtectionPolicyOptions.SetName(d.Get("name").(string))
	newBackupPolicy, err := resourceIbmProtectionPolicyMapToBackupPolicy(d.Get("backup_policy.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	updateProtectionPolicyOptions.SetBackupPolicy(newBackupPolicy)
	if _, ok := d.GetOk("id"); ok {
		updateProtectionPolicyOptions.SetID(d.Get("id").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		updateProtectionPolicyOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("blackout_window"); ok {
		var newBlackoutWindow []backuprecoveryv0.BlackoutWindow
		for _, v := range d.Get("blackout_window").([]interface{}) {
			value := v.(map[string]interface{})
			newBlackoutWindowItem, err := resourceIbmProtectionPolicyMapToBlackoutWindow(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newBlackoutWindow = append(newBlackoutWindow, *newBlackoutWindowItem)
		}
		updateProtectionPolicyOptions.SetBlackoutWindow(newBlackoutWindow)
	}
	if _, ok := d.GetOk("extended_retention"); ok {
		var newExtendedRetention []backuprecoveryv0.ExtendedRetentionPolicy
		for _, v := range d.Get("extended_retention").([]interface{}) {
			value := v.(map[string]interface{})
			newExtendedRetentionItem, err := resourceIbmProtectionPolicyMapToExtendedRetentionPolicy(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newExtendedRetention = append(newExtendedRetention, *newExtendedRetentionItem)
		}
		updateProtectionPolicyOptions.SetExtendedRetention(newExtendedRetention)
	}
	if _, ok := d.GetOk("remote_target_policy"); ok {
		newRemoteTargetPolicy, err := resourceIbmProtectionPolicyMapToTargetsConfiguration(d.Get("remote_target_policy.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionPolicyOptions.SetRemoteTargetPolicy(newRemoteTargetPolicy)
	}
	if _, ok := d.GetOk("cascaded_targets_config"); ok {
		var newCascadedTargetsConfig []backuprecoveryv0.CascadedTargetConfiguration
		for _, v := range d.Get("cascaded_targets_config").([]interface{}) {
			value := v.(map[string]interface{})
			newCascadedTargetsConfigItem, err := resourceIbmProtectionPolicyMapToCascadedTargetConfiguration(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newCascadedTargetsConfig = append(newCascadedTargetsConfig, *newCascadedTargetsConfigItem)
		}
		updateProtectionPolicyOptions.SetCascadedTargetsConfig(newCascadedTargetsConfig)
	}
	if _, ok := d.GetOk("retry_options"); ok {
		newRetryOptions, err := resourceIbmProtectionPolicyMapToRetryOptions(d.Get("retry_options.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionPolicyOptions.SetRetryOptions(newRetryOptions)
	}
	if _, ok := d.GetOk("data_lock"); ok {
		updateProtectionPolicyOptions.SetDataLock(d.Get("data_lock").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		updateProtectionPolicyOptions.SetVersion(int64(d.Get("version").(int)))
	}
	if _, ok := d.GetOk("is_cbs_enabled"); ok {
		updateProtectionPolicyOptions.SetIsCBSEnabled(d.Get("is_cbs_enabled").(bool))
	}
	if _, ok := d.GetOk("last_modification_time_usecs"); ok {
		updateProtectionPolicyOptions.SetLastModificationTimeUsecs(int64(d.Get("last_modification_time_usecs").(int)))
	}
	if _, ok := d.GetOk("template_id"); ok {
		updateProtectionPolicyOptions.SetTemplateID(d.Get("template_id").(string))
	}

	_, response, err := backupRecoveryClient.UpdateProtectionPolicyWithContext(context, updateProtectionPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateProtectionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateProtectionPolicyWithContext failed %s\n%s", err, response))
	}

	return resourceIbmProtectionPolicyRead(context, d, meta)
}

func resourceIbmProtectionPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionPolicyOptions := &backuprecoveryv0.DeleteProtectionPolicyOptions{}

	deleteProtectionPolicyOptions.SetID(d.Id())

	response, err := backupRecoveryClient.DeleteProtectionPolicyWithContext(context, deleteProtectionPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteProtectionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteProtectionPolicyWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmProtectionPolicyMapToBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.BackupPolicy, error) {
	model := &backuprecoveryv0.BackupPolicy{}
	RegularModel, err := resourceIbmProtectionPolicyMapToRegularBackupPolicy(modelMap["regular"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Regular = RegularModel
	if modelMap["log"] != nil && len(modelMap["log"].([]interface{})) > 0 {
		LogModel, err := resourceIbmProtectionPolicyMapToLogBackupPolicy(modelMap["log"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Log = LogModel
	}
	if modelMap["bmr"] != nil && len(modelMap["bmr"].([]interface{})) > 0 {
		BmrModel, err := resourceIbmProtectionPolicyMapToBmrBackupPolicy(modelMap["bmr"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Bmr = BmrModel
	}
	if modelMap["cdp"] != nil && len(modelMap["cdp"].([]interface{})) > 0 {
		CdpModel, err := resourceIbmProtectionPolicyMapToCdpBackupPolicy(modelMap["cdp"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Cdp = CdpModel
	}
	if modelMap["storage_array_snapshot"] != nil && len(modelMap["storage_array_snapshot"].([]interface{})) > 0 {
		StorageArraySnapshotModel, err := resourceIbmProtectionPolicyMapToStorageArraySnapshotBackupPolicy(modelMap["storage_array_snapshot"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.StorageArraySnapshot = StorageArraySnapshotModel
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToRegularBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.RegularBackupPolicy, error) {
	model := &backuprecoveryv0.RegularBackupPolicy{}
	if modelMap["incremental"] != nil && len(modelMap["incremental"].([]interface{})) > 0 {
		IncrementalModel, err := resourceIbmProtectionPolicyMapToIncrementalBackupPolicy(modelMap["incremental"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Incremental = IncrementalModel
	}
	if modelMap["full"] != nil && len(modelMap["full"].([]interface{})) > 0 {
		FullModel, err := resourceIbmProtectionPolicyMapToFullBackupPolicy(modelMap["full"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Full = FullModel
	}
	if modelMap["full_backups"] != nil {
		fullBackups := []backuprecoveryv0.FullScheduleAndRetention{}
		for _, fullBackupsItem := range modelMap["full_backups"].([]interface{}) {
			fullBackupsItemModel, err := resourceIbmProtectionPolicyMapToFullScheduleAndRetention(fullBackupsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			fullBackups = append(fullBackups, *fullBackupsItemModel)
		}
		model.FullBackups = fullBackups
	}
	if modelMap["retention"] != nil && len(modelMap["retention"].([]interface{})) > 0 {
		RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Retention = RetentionModel
	}
	if modelMap["primary_backup_target"] != nil && len(modelMap["primary_backup_target"].([]interface{})) > 0 {
		PrimaryBackupTargetModel, err := resourceIbmProtectionPolicyMapToPrimaryBackupTarget(modelMap["primary_backup_target"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PrimaryBackupTarget = PrimaryBackupTargetModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToIncrementalBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.IncrementalBackupPolicy, error) {
	model := &backuprecoveryv0.IncrementalBackupPolicy{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToIncrementalSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToIncrementalSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.IncrementalSchedule, error) {
	model := &backuprecoveryv0.IncrementalSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["minute_schedule"] != nil && len(modelMap["minute_schedule"].([]interface{})) > 0 {
		MinuteScheduleModel, err := resourceIbmProtectionPolicyMapToMinuteSchedule(modelMap["minute_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MinuteSchedule = MinuteScheduleModel
	}
	if modelMap["hour_schedule"] != nil && len(modelMap["hour_schedule"].([]interface{})) > 0 {
		HourScheduleModel, err := resourceIbmProtectionPolicyMapToHourSchedule(modelMap["hour_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.HourSchedule = HourScheduleModel
	}
	if modelMap["day_schedule"] != nil && len(modelMap["day_schedule"].([]interface{})) > 0 {
		DayScheduleModel, err := resourceIbmProtectionPolicyMapToDaySchedule(modelMap["day_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DaySchedule = DayScheduleModel
	}
	if modelMap["week_schedule"] != nil && len(modelMap["week_schedule"].([]interface{})) > 0 {
		WeekScheduleModel, err := resourceIbmProtectionPolicyMapToWeekSchedule(modelMap["week_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WeekSchedule = WeekScheduleModel
	}
	if modelMap["month_schedule"] != nil && len(modelMap["month_schedule"].([]interface{})) > 0 {
		MonthScheduleModel, err := resourceIbmProtectionPolicyMapToMonthSchedule(modelMap["month_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MonthSchedule = MonthScheduleModel
	}
	if modelMap["year_schedule"] != nil && len(modelMap["year_schedule"].([]interface{})) > 0 {
		YearScheduleModel, err := resourceIbmProtectionPolicyMapToYearSchedule(modelMap["year_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.YearSchedule = YearScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToMinuteSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.MinuteSchedule, error) {
	model := &backuprecoveryv0.MinuteSchedule{}
	model.Frequency = core.Int64Ptr(int64(modelMap["frequency"].(int)))
	return model, nil
}

func resourceIbmProtectionPolicyMapToHourSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.HourSchedule, error) {
	model := &backuprecoveryv0.HourSchedule{}
	model.Frequency = core.Int64Ptr(int64(modelMap["frequency"].(int)))
	return model, nil
}

func resourceIbmProtectionPolicyMapToDaySchedule(modelMap map[string]interface{}) (*backuprecoveryv0.DaySchedule, error) {
	model := &backuprecoveryv0.DaySchedule{}
	model.Frequency = core.Int64Ptr(int64(modelMap["frequency"].(int)))
	return model, nil
}

func resourceIbmProtectionPolicyMapToWeekSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.WeekSchedule, error) {
	model := &backuprecoveryv0.WeekSchedule{}
	dayOfWeek := []string{}
	for _, dayOfWeekItem := range modelMap["day_of_week"].([]interface{}) {
		dayOfWeek = append(dayOfWeek, dayOfWeekItem.(string))
	}
	model.DayOfWeek = dayOfWeek
	return model, nil
}

func resourceIbmProtectionPolicyMapToMonthSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.MonthSchedule, error) {
	model := &backuprecoveryv0.MonthSchedule{}
	if modelMap["day_of_week"] != nil {
		dayOfWeek := []string{}
		for _, dayOfWeekItem := range modelMap["day_of_week"].([]interface{}) {
			dayOfWeek = append(dayOfWeek, dayOfWeekItem.(string))
		}
		model.DayOfWeek = dayOfWeek
	}
	if modelMap["week_of_month"] != nil && modelMap["week_of_month"].(string) != "" {
		model.WeekOfMonth = core.StringPtr(modelMap["week_of_month"].(string))
	}
	if modelMap["day_of_month"] != nil {
		model.DayOfMonth = core.Int64Ptr(int64(modelMap["day_of_month"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToYearSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.YearSchedule, error) {
	model := &backuprecoveryv0.YearSchedule{}
	model.DayOfYear = core.StringPtr(modelMap["day_of_year"].(string))
	return model, nil
}

func resourceIbmProtectionPolicyMapToFullBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.FullBackupPolicy, error) {
	model := &backuprecoveryv0.FullBackupPolicy{}
	if modelMap["schedule"] != nil && len(modelMap["schedule"].([]interface{})) > 0 {
		ScheduleModel, err := resourceIbmProtectionPolicyMapToFullSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Schedule = ScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToFullSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.FullSchedule, error) {
	model := &backuprecoveryv0.FullSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["day_schedule"] != nil && len(modelMap["day_schedule"].([]interface{})) > 0 {
		DayScheduleModel, err := resourceIbmProtectionPolicyMapToDaySchedule(modelMap["day_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DaySchedule = DayScheduleModel
	}
	if modelMap["week_schedule"] != nil && len(modelMap["week_schedule"].([]interface{})) > 0 {
		WeekScheduleModel, err := resourceIbmProtectionPolicyMapToWeekSchedule(modelMap["week_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WeekSchedule = WeekScheduleModel
	}
	if modelMap["month_schedule"] != nil && len(modelMap["month_schedule"].([]interface{})) > 0 {
		MonthScheduleModel, err := resourceIbmProtectionPolicyMapToMonthSchedule(modelMap["month_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MonthSchedule = MonthScheduleModel
	}
	if modelMap["year_schedule"] != nil && len(modelMap["year_schedule"].([]interface{})) > 0 {
		YearScheduleModel, err := resourceIbmProtectionPolicyMapToYearSchedule(modelMap["year_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.YearSchedule = YearScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToFullScheduleAndRetention(modelMap map[string]interface{}) (*backuprecoveryv0.FullScheduleAndRetention, error) {
	model := &backuprecoveryv0.FullScheduleAndRetention{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToFullSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToRetention(modelMap map[string]interface{}) (*backuprecoveryv0.Retention, error) {
	model := &backuprecoveryv0.Retention{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	model.Duration = core.Int64Ptr(int64(modelMap["duration"].(int)))
	if modelMap["data_lock_config"] != nil && len(modelMap["data_lock_config"].([]interface{})) > 0 {
		DataLockConfigModel, err := resourceIbmProtectionPolicyMapToDataLockConfig(modelMap["data_lock_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataLockConfig = DataLockConfigModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToDataLockConfig(modelMap map[string]interface{}) (*backuprecoveryv0.DataLockConfig, error) {
	model := &backuprecoveryv0.DataLockConfig{}
	model.Mode = core.StringPtr(modelMap["mode"].(string))
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	model.Duration = core.Int64Ptr(int64(modelMap["duration"].(int)))
	if modelMap["enable_worm_on_external_target"] != nil {
		model.EnableWormOnExternalTarget = core.BoolPtr(modelMap["enable_worm_on_external_target"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToPrimaryBackupTarget(modelMap map[string]interface{}) (*backuprecoveryv0.PrimaryBackupTarget, error) {
	model := &backuprecoveryv0.PrimaryBackupTarget{}
	if modelMap["target_type"] != nil && modelMap["target_type"].(string) != "" {
		model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	}
	if modelMap["archival_target_settings"] != nil && len(modelMap["archival_target_settings"].([]interface{})) > 0 {
		ArchivalTargetSettingsModel, err := resourceIbmProtectionPolicyMapToPrimaryArchivalTarget(modelMap["archival_target_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ArchivalTargetSettings = ArchivalTargetSettingsModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToPrimaryArchivalTarget(modelMap map[string]interface{}) (*backuprecoveryv0.PrimaryArchivalTarget, error) {
	model := &backuprecoveryv0.PrimaryArchivalTarget{}
	model.TargetID = core.Int64Ptr(int64(modelMap["target_id"].(int)))
	if modelMap["target_name"] != nil && modelMap["target_name"].(string) != "" {
		model.TargetName = core.StringPtr(modelMap["target_name"].(string))
	}
	if modelMap["tier_settings"] != nil && len(modelMap["tier_settings"].([]interface{})) > 0 {
		TierSettingsModel, err := resourceIbmProtectionPolicyMapToTierLevelSettings(modelMap["tier_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TierSettings = TierSettingsModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToTierLevelSettings(modelMap map[string]interface{}) (*backuprecoveryv0.TierLevelSettings, error) {
	model := &backuprecoveryv0.TierLevelSettings{}
	model.CloudPlatform = core.StringPtr(modelMap["cloud_platform"].(string))
	if modelMap["oracle_tiering"] != nil && len(modelMap["oracle_tiering"].([]interface{})) > 0 {
		OracleTieringModel, err := resourceIbmProtectionPolicyMapToOracleTiers(modelMap["oracle_tiering"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.OracleTiering = OracleTieringModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToOracleTiers(modelMap map[string]interface{}) (*backuprecoveryv0.OracleTiers, error) {
	model := &backuprecoveryv0.OracleTiers{}
	tiers := []backuprecoveryv0.OracleTier{}
	for _, tiersItem := range modelMap["tiers"].([]interface{}) {
		tiersItemModel, err := resourceIbmProtectionPolicyMapToOracleTier(tiersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		tiers = append(tiers, *tiersItemModel)
	}
	model.Tiers = tiers
	return model, nil
}

func resourceIbmProtectionPolicyMapToOracleTier(modelMap map[string]interface{}) (*backuprecoveryv0.OracleTier, error) {
	model := &backuprecoveryv0.OracleTier{}
	if modelMap["move_after_unit"] != nil && modelMap["move_after_unit"].(string) != "" {
		model.MoveAfterUnit = core.StringPtr(modelMap["move_after_unit"].(string))
	}
	if modelMap["move_after"] != nil {
		model.MoveAfter = core.Int64Ptr(int64(modelMap["move_after"].(int)))
	}
	model.TierType = core.StringPtr(modelMap["tier_type"].(string))
	return model, nil
}

func resourceIbmProtectionPolicyMapToLogBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.LogBackupPolicy, error) {
	model := &backuprecoveryv0.LogBackupPolicy{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToLogSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToLogSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.LogSchedule, error) {
	model := &backuprecoveryv0.LogSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["minute_schedule"] != nil && len(modelMap["minute_schedule"].([]interface{})) > 0 {
		MinuteScheduleModel, err := resourceIbmProtectionPolicyMapToMinuteSchedule(modelMap["minute_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MinuteSchedule = MinuteScheduleModel
	}
	if modelMap["hour_schedule"] != nil && len(modelMap["hour_schedule"].([]interface{})) > 0 {
		HourScheduleModel, err := resourceIbmProtectionPolicyMapToHourSchedule(modelMap["hour_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.HourSchedule = HourScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToBmrBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.BmrBackupPolicy, error) {
	model := &backuprecoveryv0.BmrBackupPolicy{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToBmrSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToBmrSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.BmrSchedule, error) {
	model := &backuprecoveryv0.BmrSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["day_schedule"] != nil && len(modelMap["day_schedule"].([]interface{})) > 0 {
		DayScheduleModel, err := resourceIbmProtectionPolicyMapToDaySchedule(modelMap["day_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DaySchedule = DayScheduleModel
	}
	if modelMap["week_schedule"] != nil && len(modelMap["week_schedule"].([]interface{})) > 0 {
		WeekScheduleModel, err := resourceIbmProtectionPolicyMapToWeekSchedule(modelMap["week_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WeekSchedule = WeekScheduleModel
	}
	if modelMap["month_schedule"] != nil && len(modelMap["month_schedule"].([]interface{})) > 0 {
		MonthScheduleModel, err := resourceIbmProtectionPolicyMapToMonthSchedule(modelMap["month_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MonthSchedule = MonthScheduleModel
	}
	if modelMap["year_schedule"] != nil && len(modelMap["year_schedule"].([]interface{})) > 0 {
		YearScheduleModel, err := resourceIbmProtectionPolicyMapToYearSchedule(modelMap["year_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.YearSchedule = YearScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToCdpBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.CdpBackupPolicy, error) {
	model := &backuprecoveryv0.CdpBackupPolicy{}
	RetentionModel, err := resourceIbmProtectionPolicyMapToCdpRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToCdpRetention(modelMap map[string]interface{}) (*backuprecoveryv0.CdpRetention, error) {
	model := &backuprecoveryv0.CdpRetention{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	model.Duration = core.Int64Ptr(int64(modelMap["duration"].(int)))
	if modelMap["data_lock_config"] != nil && len(modelMap["data_lock_config"].([]interface{})) > 0 {
		DataLockConfigModel, err := resourceIbmProtectionPolicyMapToDataLockConfig(modelMap["data_lock_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataLockConfig = DataLockConfigModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToStorageArraySnapshotBackupPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.StorageArraySnapshotBackupPolicy, error) {
	model := &backuprecoveryv0.StorageArraySnapshotBackupPolicy{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToStorageArraySnapshotSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToStorageArraySnapshotSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.StorageArraySnapshotSchedule, error) {
	model := &backuprecoveryv0.StorageArraySnapshotSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["minute_schedule"] != nil && len(modelMap["minute_schedule"].([]interface{})) > 0 {
		MinuteScheduleModel, err := resourceIbmProtectionPolicyMapToMinuteSchedule(modelMap["minute_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MinuteSchedule = MinuteScheduleModel
	}
	if modelMap["hour_schedule"] != nil && len(modelMap["hour_schedule"].([]interface{})) > 0 {
		HourScheduleModel, err := resourceIbmProtectionPolicyMapToHourSchedule(modelMap["hour_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.HourSchedule = HourScheduleModel
	}
	if modelMap["day_schedule"] != nil && len(modelMap["day_schedule"].([]interface{})) > 0 {
		DayScheduleModel, err := resourceIbmProtectionPolicyMapToDaySchedule(modelMap["day_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DaySchedule = DayScheduleModel
	}
	if modelMap["week_schedule"] != nil && len(modelMap["week_schedule"].([]interface{})) > 0 {
		WeekScheduleModel, err := resourceIbmProtectionPolicyMapToWeekSchedule(modelMap["week_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WeekSchedule = WeekScheduleModel
	}
	if modelMap["month_schedule"] != nil && len(modelMap["month_schedule"].([]interface{})) > 0 {
		MonthScheduleModel, err := resourceIbmProtectionPolicyMapToMonthSchedule(modelMap["month_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MonthSchedule = MonthScheduleModel
	}
	if modelMap["year_schedule"] != nil && len(modelMap["year_schedule"].([]interface{})) > 0 {
		YearScheduleModel, err := resourceIbmProtectionPolicyMapToYearSchedule(modelMap["year_schedule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.YearSchedule = YearScheduleModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToCancellationTimeoutParams(modelMap map[string]interface{}) (*backuprecoveryv0.CancellationTimeoutParams, error) {
	model := &backuprecoveryv0.CancellationTimeoutParams{}
	if modelMap["timeout_mins"] != nil {
		model.TimeoutMins = core.Int64Ptr(int64(modelMap["timeout_mins"].(int)))
	}
	if modelMap["backup_type"] != nil && modelMap["backup_type"].(string) != "" {
		model.BackupType = core.StringPtr(modelMap["backup_type"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToBlackoutWindow(modelMap map[string]interface{}) (*backuprecoveryv0.BlackoutWindow, error) {
	model := &backuprecoveryv0.BlackoutWindow{}
	model.Day = core.StringPtr(modelMap["day"].(string))
	StartTimeModel, err := resourceIbmProtectionPolicyMapToTimeOfDay(modelMap["start_time"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.StartTime = StartTimeModel
	EndTimeModel, err := resourceIbmProtectionPolicyMapToTimeOfDay(modelMap["end_time"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.EndTime = EndTimeModel
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToTimeOfDay(modelMap map[string]interface{}) (*backuprecoveryv0.TimeOfDay, error) {
	model := &backuprecoveryv0.TimeOfDay{}
	model.Hour = core.Int64Ptr(int64(modelMap["hour"].(int)))
	model.Minute = core.Int64Ptr(int64(modelMap["minute"].(int)))
	if modelMap["time_zone"] != nil && modelMap["time_zone"].(string) != "" {
		model.TimeZone = core.StringPtr(modelMap["time_zone"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToExtendedRetentionPolicy(modelMap map[string]interface{}) (*backuprecoveryv0.ExtendedRetentionPolicy, error) {
	model := &backuprecoveryv0.ExtendedRetentionPolicy{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToExtendedRetentionSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["run_type"] != nil && modelMap["run_type"].(string) != "" {
		model.RunType = core.StringPtr(modelMap["run_type"].(string))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToExtendedRetentionSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.ExtendedRetentionSchedule, error) {
	model := &backuprecoveryv0.ExtendedRetentionSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["frequency"] != nil {
		model.Frequency = core.Int64Ptr(int64(modelMap["frequency"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToTargetsConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.TargetsConfiguration, error) {
	model := &backuprecoveryv0.TargetsConfiguration{}
	if modelMap["replication_targets"] != nil {
		replicationTargets := []backuprecoveryv0.ReplicationTargetConfiguration{}
		for _, replicationTargetsItem := range modelMap["replication_targets"].([]interface{}) {
			replicationTargetsItemModel, err := resourceIbmProtectionPolicyMapToReplicationTargetConfiguration(replicationTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			replicationTargets = append(replicationTargets, *replicationTargetsItemModel)
		}
		model.ReplicationTargets = replicationTargets
	}
	if modelMap["archival_targets"] != nil {
		archivalTargets := []backuprecoveryv0.ArchivalTargetConfiguration{}
		for _, archivalTargetsItem := range modelMap["archival_targets"].([]interface{}) {
			archivalTargetsItemModel, err := resourceIbmProtectionPolicyMapToArchivalTargetConfiguration(archivalTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			archivalTargets = append(archivalTargets, *archivalTargetsItemModel)
		}
		model.ArchivalTargets = archivalTargets
	}
	if modelMap["cloud_spin_targets"] != nil {
		cloudSpinTargets := []backuprecoveryv0.CloudSpinTargetConfiguration{}
		for _, cloudSpinTargetsItem := range modelMap["cloud_spin_targets"].([]interface{}) {
			cloudSpinTargetsItemModel, err := resourceIbmProtectionPolicyMapToCloudSpinTargetConfiguration(cloudSpinTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			cloudSpinTargets = append(cloudSpinTargets, *cloudSpinTargetsItemModel)
		}
		model.CloudSpinTargets = cloudSpinTargets
	}
	if modelMap["onprem_deploy_targets"] != nil {
		onpremDeployTargets := []backuprecoveryv0.OnpremDeployTargetConfiguration{}
		for _, onpremDeployTargetsItem := range modelMap["onprem_deploy_targets"].([]interface{}) {
			onpremDeployTargetsItemModel, err := resourceIbmProtectionPolicyMapToOnpremDeployTargetConfiguration(onpremDeployTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			onpremDeployTargets = append(onpremDeployTargets, *onpremDeployTargetsItemModel)
		}
		model.OnpremDeployTargets = onpremDeployTargets
	}
	if modelMap["rpaas_targets"] != nil {
		rpaasTargets := []backuprecoveryv0.RpaasTargetConfiguration{}
		for _, rpaasTargetsItem := range modelMap["rpaas_targets"].([]interface{}) {
			rpaasTargetsItemModel, err := resourceIbmProtectionPolicyMapToRpaasTargetConfiguration(rpaasTargetsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			rpaasTargets = append(rpaasTargets, *rpaasTargetsItemModel)
		}
		model.RpaasTargets = rpaasTargets
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToReplicationTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.ReplicationTargetConfiguration, error) {
	model := &backuprecoveryv0.ReplicationTargetConfiguration{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToTargetSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["copy_on_run_success"] != nil {
		model.CopyOnRunSuccess = core.BoolPtr(modelMap["copy_on_run_success"].(bool))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	if modelMap["log_retention"] != nil && len(modelMap["log_retention"].([]interface{})) > 0 {
		LogRetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["log_retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogRetention = LogRetentionModel
	}
	model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	if modelMap["remote_target_config"] != nil && len(modelMap["remote_target_config"].([]interface{})) > 0 {
		RemoteTargetConfigModel, err := resourceIbmProtectionPolicyMapToRemoteTargetConfig(modelMap["remote_target_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RemoteTargetConfig = RemoteTargetConfigModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToTargetSchedule(modelMap map[string]interface{}) (*backuprecoveryv0.TargetSchedule, error) {
	model := &backuprecoveryv0.TargetSchedule{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	if modelMap["frequency"] != nil {
		model.Frequency = core.Int64Ptr(int64(modelMap["frequency"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToRemoteTargetConfig(modelMap map[string]interface{}) (*backuprecoveryv0.RemoteTargetConfig, error) {
	model := &backuprecoveryv0.RemoteTargetConfig{}
	model.ClusterID = core.Int64Ptr(int64(modelMap["cluster_id"].(int)))
	if modelMap["cluster_name"] != nil && modelMap["cluster_name"].(string) != "" {
		model.ClusterName = core.StringPtr(modelMap["cluster_name"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToArchivalTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.ArchivalTargetConfiguration, error) {
	model := &backuprecoveryv0.ArchivalTargetConfiguration{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToTargetSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["copy_on_run_success"] != nil {
		model.CopyOnRunSuccess = core.BoolPtr(modelMap["copy_on_run_success"].(bool))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	if modelMap["log_retention"] != nil && len(modelMap["log_retention"].([]interface{})) > 0 {
		LogRetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["log_retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogRetention = LogRetentionModel
	}
	model.TargetID = core.Int64Ptr(int64(modelMap["target_id"].(int)))
	if modelMap["target_name"] != nil && modelMap["target_name"].(string) != "" {
		model.TargetName = core.StringPtr(modelMap["target_name"].(string))
	}
	if modelMap["target_type"] != nil && modelMap["target_type"].(string) != "" {
		model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	}
	if modelMap["tier_settings"] != nil && len(modelMap["tier_settings"].([]interface{})) > 0 {
		TierSettingsModel, err := resourceIbmProtectionPolicyMapToTierLevelSettings(modelMap["tier_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TierSettings = TierSettingsModel
	}
	if modelMap["extended_retention"] != nil {
		extendedRetention := []backuprecoveryv0.ExtendedRetentionPolicy{}
		for _, extendedRetentionItem := range modelMap["extended_retention"].([]interface{}) {
			extendedRetentionItemModel, err := resourceIbmProtectionPolicyMapToExtendedRetentionPolicy(extendedRetentionItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			extendedRetention = append(extendedRetention, *extendedRetentionItemModel)
		}
		model.ExtendedRetention = extendedRetention
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToCloudSpinTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.CloudSpinTargetConfiguration, error) {
	model := &backuprecoveryv0.CloudSpinTargetConfiguration{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToTargetSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["copy_on_run_success"] != nil {
		model.CopyOnRunSuccess = core.BoolPtr(modelMap["copy_on_run_success"].(bool))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	if modelMap["log_retention"] != nil && len(modelMap["log_retention"].([]interface{})) > 0 {
		LogRetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["log_retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogRetention = LogRetentionModel
	}
	TargetModel, err := resourceIbmProtectionPolicyMapToCloudSpinTarget(modelMap["target"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Target = TargetModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToCloudSpinTarget(modelMap map[string]interface{}) (*backuprecoveryv0.CloudSpinTarget, error) {
	model := &backuprecoveryv0.CloudSpinTarget{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToOnpremDeployTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.OnpremDeployTargetConfiguration, error) {
	model := &backuprecoveryv0.OnpremDeployTargetConfiguration{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToTargetSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["copy_on_run_success"] != nil {
		model.CopyOnRunSuccess = core.BoolPtr(modelMap["copy_on_run_success"].(bool))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	if modelMap["log_retention"] != nil && len(modelMap["log_retention"].([]interface{})) > 0 {
		LogRetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["log_retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogRetention = LogRetentionModel
	}
	if modelMap["params"] != nil && len(modelMap["params"].([]interface{})) > 0 {
		ParamsModel, err := resourceIbmProtectionPolicyMapToOnpremDeployParams(modelMap["params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Params = ParamsModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToOnpremDeployParams(modelMap map[string]interface{}) (*backuprecoveryv0.OnpremDeployParams, error) {
	model := &backuprecoveryv0.OnpremDeployParams{}
	if modelMap["id"] != nil {
		model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	}
	if modelMap["restore_v_mware_params"] != nil && len(modelMap["restore_v_mware_params"].([]interface{})) > 0 {
		RestoreVMwareParamsModel, err := resourceIbmProtectionPolicyMapToRestoreVMwareVMParams(modelMap["restore_v_mware_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RestoreVMwareParams = RestoreVMwareParamsModel
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToRestoreVMwareVMParams(modelMap map[string]interface{}) (*backuprecoveryv0.RestoreVMwareVMParams, error) {
	model := &backuprecoveryv0.RestoreVMwareVMParams{}
	if modelMap["target_vm_folder_id"] != nil {
		model.TargetVMFolderID = core.Int64Ptr(int64(modelMap["target_vm_folder_id"].(int)))
	}
	if modelMap["target_data_store_id"] != nil {
		model.TargetDataStoreID = core.Int64Ptr(int64(modelMap["target_data_store_id"].(int)))
	}
	if modelMap["enable_copy_recovery"] != nil {
		model.EnableCopyRecovery = core.BoolPtr(modelMap["enable_copy_recovery"].(bool))
	}
	if modelMap["resource_pool_id"] != nil {
		model.ResourcePoolID = core.Int64Ptr(int64(modelMap["resource_pool_id"].(int)))
	}
	if modelMap["datastore_ids"] != nil {
		datastoreIds := []int64{}
		for _, datastoreIdsItem := range modelMap["datastore_ids"].([]interface{}) {
			datastoreIds = append(datastoreIds, int64(datastoreIdsItem.(int)))
		}
		model.DatastoreIds = datastoreIds
	}
	if modelMap["overwrite_existing_vm"] != nil {
		model.OverwriteExistingVm = core.BoolPtr(modelMap["overwrite_existing_vm"].(bool))
	}
	if modelMap["power_off_and_rename_existing_vm"] != nil {
		model.PowerOffAndRenameExistingVm = core.BoolPtr(modelMap["power_off_and_rename_existing_vm"].(bool))
	}
	if modelMap["attempt_differential_restore"] != nil {
		model.AttemptDifferentialRestore = core.BoolPtr(modelMap["attempt_differential_restore"].(bool))
	}
	if modelMap["is_on_prem_deploy"] != nil {
		model.IsOnPremDeploy = core.BoolPtr(modelMap["is_on_prem_deploy"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToRpaasTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.RpaasTargetConfiguration, error) {
	model := &backuprecoveryv0.RpaasTargetConfiguration{}
	ScheduleModel, err := resourceIbmProtectionPolicyMapToTargetSchedule(modelMap["schedule"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Schedule = ScheduleModel
	RetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Retention = RetentionModel
	if modelMap["copy_on_run_success"] != nil {
		model.CopyOnRunSuccess = core.BoolPtr(modelMap["copy_on_run_success"].(bool))
	}
	if modelMap["config_id"] != nil && modelMap["config_id"].(string) != "" {
		model.ConfigID = core.StringPtr(modelMap["config_id"].(string))
	}
	if modelMap["backup_run_type"] != nil && modelMap["backup_run_type"].(string) != "" {
		model.BackupRunType = core.StringPtr(modelMap["backup_run_type"].(string))
	}
	if modelMap["run_timeouts"] != nil {
		runTimeouts := []backuprecoveryv0.CancellationTimeoutParams{}
		for _, runTimeoutsItem := range modelMap["run_timeouts"].([]interface{}) {
			runTimeoutsItemModel, err := resourceIbmProtectionPolicyMapToCancellationTimeoutParams(runTimeoutsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			runTimeouts = append(runTimeouts, *runTimeoutsItemModel)
		}
		model.RunTimeouts = runTimeouts
	}
	if modelMap["log_retention"] != nil && len(modelMap["log_retention"].([]interface{})) > 0 {
		LogRetentionModel, err := resourceIbmProtectionPolicyMapToRetention(modelMap["log_retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogRetention = LogRetentionModel
	}
	model.TargetID = core.Int64Ptr(int64(modelMap["target_id"].(int)))
	if modelMap["target_name"] != nil && modelMap["target_name"].(string) != "" {
		model.TargetName = core.StringPtr(modelMap["target_name"].(string))
	}
	if modelMap["target_type"] != nil && modelMap["target_type"].(string) != "" {
		model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	}
	return model, nil
}

func resourceIbmProtectionPolicyMapToCascadedTargetConfiguration(modelMap map[string]interface{}) (*backuprecoveryv0.CascadedTargetConfiguration, error) {
	model := &backuprecoveryv0.CascadedTargetConfiguration{}
	model.SourceClusterID = core.Int64Ptr(int64(modelMap["source_cluster_id"].(int)))
	RemoteTargetsModel, err := resourceIbmProtectionPolicyMapToTargetsConfiguration(modelMap["remote_targets"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.RemoteTargets = RemoteTargetsModel
	return model, nil
}

func resourceIbmProtectionPolicyMapToRetryOptions(modelMap map[string]interface{}) (*backuprecoveryv0.RetryOptions, error) {
	model := &backuprecoveryv0.RetryOptions{}
	if modelMap["retries"] != nil {
		model.Retries = core.Int64Ptr(int64(modelMap["retries"].(int)))
	}
	if modelMap["retry_interval_mins"] != nil {
		model.RetryIntervalMins = core.Int64Ptr(int64(modelMap["retry_interval_mins"].(int)))
	}
	return model, nil
}

func resourceIbmProtectionPolicyBackupPolicyToMap(model *backuprecoveryv0.BackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	regularMap, err := resourceIbmProtectionPolicyRegularBackupPolicyToMap(model.Regular)
	if err != nil {
		return modelMap, err
	}
	modelMap["regular"] = []map[string]interface{}{regularMap}
	if model.Log != nil {
		logMap, err := resourceIbmProtectionPolicyLogBackupPolicyToMap(model.Log)
		if err != nil {
			return modelMap, err
		}
		modelMap["log"] = []map[string]interface{}{logMap}
	}
	if model.Bmr != nil {
		bmrMap, err := resourceIbmProtectionPolicyBmrBackupPolicyToMap(model.Bmr)
		if err != nil {
			return modelMap, err
		}
		modelMap["bmr"] = []map[string]interface{}{bmrMap}
	}
	if model.Cdp != nil {
		cdpMap, err := resourceIbmProtectionPolicyCdpBackupPolicyToMap(model.Cdp)
		if err != nil {
			return modelMap, err
		}
		modelMap["cdp"] = []map[string]interface{}{cdpMap}
	}
	if model.StorageArraySnapshot != nil {
		storageArraySnapshotMap, err := resourceIbmProtectionPolicyStorageArraySnapshotBackupPolicyToMap(model.StorageArraySnapshot)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_array_snapshot"] = []map[string]interface{}{storageArraySnapshotMap}
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRegularBackupPolicyToMap(model *backuprecoveryv0.RegularBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Incremental != nil {
		incrementalMap, err := resourceIbmProtectionPolicyIncrementalBackupPolicyToMap(model.Incremental)
		if err != nil {
			return modelMap, err
		}
		modelMap["incremental"] = []map[string]interface{}{incrementalMap}
	}
	if model.Full != nil {
		fullMap, err := resourceIbmProtectionPolicyFullBackupPolicyToMap(model.Full)
		if err != nil {
			return modelMap, err
		}
		modelMap["full"] = []map[string]interface{}{fullMap}
	}
	if model.FullBackups != nil {
		fullBackups := []map[string]interface{}{}
		for _, fullBackupsItem := range model.FullBackups {
			fullBackupsItemMap, err := resourceIbmProtectionPolicyFullScheduleAndRetentionToMap(&fullBackupsItem)
			if err != nil {
				return modelMap, err
			}
			fullBackups = append(fullBackups, fullBackupsItemMap)
		}
		modelMap["full_backups"] = fullBackups
	}
	if model.Retention != nil {
		retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	if model.PrimaryBackupTarget != nil {
		primaryBackupTargetMap, err := resourceIbmProtectionPolicyPrimaryBackupTargetToMap(model.PrimaryBackupTarget)
		if err != nil {
			return modelMap, err
		}
		modelMap["primary_backup_target"] = []map[string]interface{}{primaryBackupTargetMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyIncrementalBackupPolicyToMap(model *backuprecoveryv0.IncrementalBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyIncrementalScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyIncrementalScheduleToMap(model *backuprecoveryv0.IncrementalSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := resourceIbmProtectionPolicyMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := resourceIbmProtectionPolicyHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	if model.DaySchedule != nil {
		dayScheduleMap, err := resourceIbmProtectionPolicyDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := resourceIbmProtectionPolicyWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := resourceIbmProtectionPolicyMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := resourceIbmProtectionPolicyYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyMinuteScheduleToMap(model *backuprecoveryv0.MinuteSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func resourceIbmProtectionPolicyHourScheduleToMap(model *backuprecoveryv0.HourSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func resourceIbmProtectionPolicyDayScheduleToMap(model *backuprecoveryv0.DaySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func resourceIbmProtectionPolicyWeekScheduleToMap(model *backuprecoveryv0.WeekSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	return modelMap, nil
}

func resourceIbmProtectionPolicyMonthScheduleToMap(model *backuprecoveryv0.MonthSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DayOfWeek != nil {
		modelMap["day_of_week"] = model.DayOfWeek
	}
	if model.WeekOfMonth != nil {
		modelMap["week_of_month"] = model.WeekOfMonth
	}
	if model.DayOfMonth != nil {
		modelMap["day_of_month"] = flex.IntValue(model.DayOfMonth)
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyYearScheduleToMap(model *backuprecoveryv0.YearSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_year"] = model.DayOfYear
	return modelMap, nil
}

func resourceIbmProtectionPolicyFullBackupPolicyToMap(model *backuprecoveryv0.FullBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Schedule != nil {
		scheduleMap, err := resourceIbmProtectionPolicyFullScheduleToMap(model.Schedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyFullScheduleToMap(model *backuprecoveryv0.FullSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.DaySchedule != nil {
		dayScheduleMap, err := resourceIbmProtectionPolicyDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := resourceIbmProtectionPolicyWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := resourceIbmProtectionPolicyMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := resourceIbmProtectionPolicyYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyFullScheduleAndRetentionToMap(model *backuprecoveryv0.FullScheduleAndRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyFullScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRetentionToMap(model *backuprecoveryv0.Retention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	if model.DataLockConfig != nil {
		dataLockConfigMap, err := resourceIbmProtectionPolicyDataLockConfigToMap(model.DataLockConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_config"] = []map[string]interface{}{dataLockConfigMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyDataLockConfigToMap(model *backuprecoveryv0.DataLockConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mode"] = model.Mode
	modelMap["unit"] = model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	if model.EnableWormOnExternalTarget != nil {
		modelMap["enable_worm_on_external_target"] = model.EnableWormOnExternalTarget
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyPrimaryBackupTargetToMap(model *backuprecoveryv0.PrimaryBackupTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetType != nil {
		modelMap["target_type"] = model.TargetType
	}
	if model.ArchivalTargetSettings != nil {
		archivalTargetSettingsMap, err := resourceIbmProtectionPolicyPrimaryArchivalTargetToMap(model.ArchivalTargetSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["archival_target_settings"] = []map[string]interface{}{archivalTargetSettingsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyPrimaryArchivalTargetToMap(model *backuprecoveryv0.PrimaryArchivalTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_id"] = flex.IntValue(model.TargetID)
	if model.TargetName != nil {
		modelMap["target_name"] = model.TargetName
	}
	if model.TierSettings != nil {
		tierSettingsMap, err := resourceIbmProtectionPolicyTierLevelSettingsToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyTierLevelSettingsToMap(model *backuprecoveryv0.TierLevelSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := resourceIbmProtectionPolicyOracleTiersToMap(model.OracleTiering)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_tiering"] = []map[string]interface{}{oracleTieringMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyOracleTiersToMap(model *backuprecoveryv0.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := resourceIbmProtectionPolicyOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func resourceIbmProtectionPolicyOracleTierToMap(model *backuprecoveryv0.OracleTier) (map[string]interface{}, error) {
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

func resourceIbmProtectionPolicyLogBackupPolicyToMap(model *backuprecoveryv0.LogBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyLogScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyLogScheduleToMap(model *backuprecoveryv0.LogSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := resourceIbmProtectionPolicyMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := resourceIbmProtectionPolicyHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyBmrBackupPolicyToMap(model *backuprecoveryv0.BmrBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyBmrScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyBmrScheduleToMap(model *backuprecoveryv0.BmrSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.DaySchedule != nil {
		dayScheduleMap, err := resourceIbmProtectionPolicyDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := resourceIbmProtectionPolicyWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := resourceIbmProtectionPolicyMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := resourceIbmProtectionPolicyYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCdpBackupPolicyToMap(model *backuprecoveryv0.CdpBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	retentionMap, err := resourceIbmProtectionPolicyCdpRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCdpRetentionToMap(model *backuprecoveryv0.CdpRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	if model.DataLockConfig != nil {
		dataLockConfigMap, err := resourceIbmProtectionPolicyDataLockConfigToMap(model.DataLockConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_config"] = []map[string]interface{}{dataLockConfigMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyStorageArraySnapshotBackupPolicyToMap(model *backuprecoveryv0.StorageArraySnapshotBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyStorageArraySnapshotScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyStorageArraySnapshotScheduleToMap(model *backuprecoveryv0.StorageArraySnapshotSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := resourceIbmProtectionPolicyMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := resourceIbmProtectionPolicyHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	if model.DaySchedule != nil {
		dayScheduleMap, err := resourceIbmProtectionPolicyDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := resourceIbmProtectionPolicyWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := resourceIbmProtectionPolicyMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := resourceIbmProtectionPolicyYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(model *backuprecoveryv0.CancellationTimeoutParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TimeoutMins != nil {
		modelMap["timeout_mins"] = flex.IntValue(model.TimeoutMins)
	}
	if model.BackupType != nil {
		modelMap["backup_type"] = model.BackupType
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyBlackoutWindowToMap(model *backuprecoveryv0.BlackoutWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day"] = model.Day
	startTimeMap, err := resourceIbmProtectionPolicyTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := resourceIbmProtectionPolicyTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyTimeOfDayToMap(model *backuprecoveryv0.TimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["hour"] = flex.IntValue(model.Hour)
	modelMap["minute"] = flex.IntValue(model.Minute)
	if model.TimeZone != nil {
		modelMap["time_zone"] = model.TimeZone
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyExtendedRetentionPolicyToMap(model *backuprecoveryv0.ExtendedRetentionPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyExtendedRetentionScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.RunType != nil {
		modelMap["run_type"] = model.RunType
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyExtendedRetentionScheduleToMap(model *backuprecoveryv0.ExtendedRetentionSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.Frequency != nil {
		modelMap["frequency"] = flex.IntValue(model.Frequency)
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyTargetsConfigurationToMap(model *backuprecoveryv0.TargetsConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplicationTargets != nil {
		replicationTargets := []map[string]interface{}{}
		for _, replicationTargetsItem := range model.ReplicationTargets {
			replicationTargetsItemMap, err := resourceIbmProtectionPolicyReplicationTargetConfigurationToMap(&replicationTargetsItem)
			if err != nil {
				return modelMap, err
			}
			replicationTargets = append(replicationTargets, replicationTargetsItemMap)
		}
		modelMap["replication_targets"] = replicationTargets
	}
	if model.ArchivalTargets != nil {
		archivalTargets := []map[string]interface{}{}
		for _, archivalTargetsItem := range model.ArchivalTargets {
			archivalTargetsItemMap, err := resourceIbmProtectionPolicyArchivalTargetConfigurationToMap(&archivalTargetsItem)
			if err != nil {
				return modelMap, err
			}
			archivalTargets = append(archivalTargets, archivalTargetsItemMap)
		}
		modelMap["archival_targets"] = archivalTargets
	}
	if model.CloudSpinTargets != nil {
		cloudSpinTargets := []map[string]interface{}{}
		for _, cloudSpinTargetsItem := range model.CloudSpinTargets {
			cloudSpinTargetsItemMap, err := resourceIbmProtectionPolicyCloudSpinTargetConfigurationToMap(&cloudSpinTargetsItem)
			if err != nil {
				return modelMap, err
			}
			cloudSpinTargets = append(cloudSpinTargets, cloudSpinTargetsItemMap)
		}
		modelMap["cloud_spin_targets"] = cloudSpinTargets
	}
	if model.OnpremDeployTargets != nil {
		onpremDeployTargets := []map[string]interface{}{}
		for _, onpremDeployTargetsItem := range model.OnpremDeployTargets {
			onpremDeployTargetsItemMap, err := resourceIbmProtectionPolicyOnpremDeployTargetConfigurationToMap(&onpremDeployTargetsItem)
			if err != nil {
				return modelMap, err
			}
			onpremDeployTargets = append(onpremDeployTargets, onpremDeployTargetsItemMap)
		}
		modelMap["onprem_deploy_targets"] = onpremDeployTargets
	}
	if model.RpaasTargets != nil {
		rpaasTargets := []map[string]interface{}{}
		for _, rpaasTargetsItem := range model.RpaasTargets {
			rpaasTargetsItemMap, err := resourceIbmProtectionPolicyRpaasTargetConfigurationToMap(&rpaasTargetsItem)
			if err != nil {
				return modelMap, err
			}
			rpaasTargets = append(rpaasTargets, rpaasTargetsItemMap)
		}
		modelMap["rpaas_targets"] = rpaasTargets
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyReplicationTargetConfigurationToMap(model *backuprecoveryv0.ReplicationTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyTargetScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.CopyOnRunSuccess != nil {
		modelMap["copy_on_run_success"] = model.CopyOnRunSuccess
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	if model.LogRetention != nil {
		logRetentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.LogRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_retention"] = []map[string]interface{}{logRetentionMap}
	}
	modelMap["target_type"] = model.TargetType
	if model.RemoteTargetConfig != nil {
		remoteTargetConfigMap, err := resourceIbmProtectionPolicyRemoteTargetConfigToMap(model.RemoteTargetConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["remote_target_config"] = []map[string]interface{}{remoteTargetConfigMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyTargetScheduleToMap(model *backuprecoveryv0.TargetSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	if model.Frequency != nil {
		modelMap["frequency"] = flex.IntValue(model.Frequency)
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRemoteTargetConfigToMap(model *backuprecoveryv0.RemoteTargetConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	if model.ClusterName != nil {
		modelMap["cluster_name"] = model.ClusterName
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyArchivalTargetConfigurationToMap(model *backuprecoveryv0.ArchivalTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyTargetScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.CopyOnRunSuccess != nil {
		modelMap["copy_on_run_success"] = model.CopyOnRunSuccess
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	if model.LogRetention != nil {
		logRetentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.LogRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_retention"] = []map[string]interface{}{logRetentionMap}
	}
	modelMap["target_id"] = flex.IntValue(model.TargetID)
	if model.TargetName != nil {
		modelMap["target_name"] = model.TargetName
	}
	if model.TargetType != nil {
		modelMap["target_type"] = model.TargetType
	}
	if model.TierSettings != nil {
		tierSettingsMap, err := resourceIbmProtectionPolicyTierLevelSettingsToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.ExtendedRetention != nil {
		extendedRetention := []map[string]interface{}{}
		for _, extendedRetentionItem := range model.ExtendedRetention {
			extendedRetentionItemMap, err := resourceIbmProtectionPolicyExtendedRetentionPolicyToMap(&extendedRetentionItem)
			if err != nil {
				return modelMap, err
			}
			extendedRetention = append(extendedRetention, extendedRetentionItemMap)
		}
		modelMap["extended_retention"] = extendedRetention
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCloudSpinTargetConfigurationToMap(model *backuprecoveryv0.CloudSpinTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyTargetScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.CopyOnRunSuccess != nil {
		modelMap["copy_on_run_success"] = model.CopyOnRunSuccess
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	if model.LogRetention != nil {
		logRetentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.LogRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_retention"] = []map[string]interface{}{logRetentionMap}
	}
	targetMap, err := resourceIbmProtectionPolicyCloudSpinTargetToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCloudSpinTargetToMap(model *backuprecoveryv0.CloudSpinTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyOnpremDeployTargetConfigurationToMap(model *backuprecoveryv0.OnpremDeployTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyTargetScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.CopyOnRunSuccess != nil {
		modelMap["copy_on_run_success"] = model.CopyOnRunSuccess
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	if model.LogRetention != nil {
		logRetentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.LogRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_retention"] = []map[string]interface{}{logRetentionMap}
	}
	if model.Params != nil {
		paramsMap, err := resourceIbmProtectionPolicyOnpremDeployParamsToMap(model.Params)
		if err != nil {
			return modelMap, err
		}
		modelMap["params"] = []map[string]interface{}{paramsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyOnpremDeployParamsToMap(model *backuprecoveryv0.OnpremDeployParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.RestoreVMwareParams != nil {
		restoreVMwareParamsMap, err := resourceIbmProtectionPolicyRestoreVMwareVMParamsToMap(model.RestoreVMwareParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["restore_v_mware_params"] = []map[string]interface{}{restoreVMwareParamsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRestoreVMwareVMParamsToMap(model *backuprecoveryv0.RestoreVMwareVMParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetVMFolderID != nil {
		modelMap["target_vm_folder_id"] = flex.IntValue(model.TargetVMFolderID)
	}
	if model.TargetDataStoreID != nil {
		modelMap["target_data_store_id"] = flex.IntValue(model.TargetDataStoreID)
	}
	if model.EnableCopyRecovery != nil {
		modelMap["enable_copy_recovery"] = model.EnableCopyRecovery
	}
	if model.ResourcePoolID != nil {
		modelMap["resource_pool_id"] = flex.IntValue(model.ResourcePoolID)
	}
	if model.DatastoreIds != nil {
		modelMap["datastore_ids"] = model.DatastoreIds
	}
	if model.OverwriteExistingVm != nil {
		modelMap["overwrite_existing_vm"] = model.OverwriteExistingVm
	}
	if model.PowerOffAndRenameExistingVm != nil {
		modelMap["power_off_and_rename_existing_vm"] = model.PowerOffAndRenameExistingVm
	}
	if model.AttemptDifferentialRestore != nil {
		modelMap["attempt_differential_restore"] = model.AttemptDifferentialRestore
	}
	if model.IsOnPremDeploy != nil {
		modelMap["is_on_prem_deploy"] = model.IsOnPremDeploy
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRpaasTargetConfigurationToMap(model *backuprecoveryv0.RpaasTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := resourceIbmProtectionPolicyTargetScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	if model.CopyOnRunSuccess != nil {
		modelMap["copy_on_run_success"] = model.CopyOnRunSuccess
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.BackupRunType != nil {
		modelMap["backup_run_type"] = model.BackupRunType
	}
	if model.RunTimeouts != nil {
		runTimeouts := []map[string]interface{}{}
		for _, runTimeoutsItem := range model.RunTimeouts {
			runTimeoutsItemMap, err := resourceIbmProtectionPolicyCancellationTimeoutParamsToMap(&runTimeoutsItem)
			if err != nil {
				return modelMap, err
			}
			runTimeouts = append(runTimeouts, runTimeoutsItemMap)
		}
		modelMap["run_timeouts"] = runTimeouts
	}
	if model.LogRetention != nil {
		logRetentionMap, err := resourceIbmProtectionPolicyRetentionToMap(model.LogRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_retention"] = []map[string]interface{}{logRetentionMap}
	}
	modelMap["target_id"] = flex.IntValue(model.TargetID)
	if model.TargetName != nil {
		modelMap["target_name"] = model.TargetName
	}
	if model.TargetType != nil {
		modelMap["target_type"] = model.TargetType
	}
	return modelMap, nil
}

func resourceIbmProtectionPolicyCascadedTargetConfigurationToMap(model *backuprecoveryv0.CascadedTargetConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_cluster_id"] = flex.IntValue(model.SourceClusterID)
	remoteTargetsMap, err := resourceIbmProtectionPolicyTargetsConfigurationToMap(model.RemoteTargets)
	if err != nil {
		return modelMap, err
	}
	modelMap["remote_targets"] = []map[string]interface{}{remoteTargetsMap}
	return modelMap, nil
}

func resourceIbmProtectionPolicyRetryOptionsToMap(model *backuprecoveryv0.RetryOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Retries != nil {
		modelMap["retries"] = flex.IntValue(model.Retries)
	}
	if model.RetryIntervalMins != nil {
		modelMap["retry_interval_mins"] = flex.IntValue(model.RetryIntervalMins)
	}
	return modelMap, nil
}
