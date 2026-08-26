// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func DataSourceIbmBrsMigrationWorkload() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationWorkloadRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration service workload ID (wl-{uuid4} format).",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable name for this workload.",
			},
			"volume_ownership_map": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Server-computed map of `vol-*` volume_id → owning payload_id. Only populated when the same `source.volume_id` appears in two or more payloads. Null in all other cases.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current lifecycle state of the workload.",
			},
			"payloads": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of source-to-destination payload mappings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service payload ID (pl-{uuid4} format).",
						},
						"source_host_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service ID of the source host.",
						},
						"destination_host_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration service ID of the destination host.",
						},
						"data_specs": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "One or more data specifications for this payload.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_format": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "How this data unit is structured. When `source.volume_id` is a registered `vol-*` ID, the server derives this from the volume registry — pass `raw` as a default if unsure.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Source data location and format for this mapping.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"volume_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Migration service vol-* ID from POST /migrations/{migration_id}/volumes.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Filesystem type, database engine, or raw block device type marker.",
												},
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "OS mount path, database data directory path, or block device path.",
												},
											},
										},
									},
									"destination": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Destination data location and format for this mapping.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"volume_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Migration service vol-* ID from POST /migrations/{migration_id}/volumes.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Filesystem type, database engine, or raw block device type marker.",
												},
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "OS mount path, database data directory path, or block device path.",
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
			"schedule": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Populated once setup completes. Null while settingUp or in created/failed state.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable name for the schedule / protection policy.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional description of the schedule.",
						},
						"backup_policy": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the backup schedule and retentions of a protection policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"regular": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the incremental and full backup schedule settings with retention. When `incremental` or `full` is provided, `retention` must also be provided — the BRS API returns a 400 if a schedule is present without retention. When neither schedule is supplied, a default daily incremental with 30-day retention is applied automatically by the service.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"incremental": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies incremental backup schedule settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"schedule": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Defines how frequently incremental backup runs are started. The sub-schedule field matching `unit` must be provided: `minutes` → `minute_schedule`, `hours` → `hour_schedule`, `days` → `day_schedule`, `weeks` → `week_schedule` (with at least one day), `months` → `month_schedule`, `years` → `year_schedule`. The BRS API returns a 400 if the matching sub-schedule is absent.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Frequency unit for the incremental schedule.",
																		},
																		"minute_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "A simple multiplier schedule — repeat every N units.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"frequency": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "How many units between each run.",
																					},
																				},
																			},
																		},
																		"hour_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "A simple multiplier schedule — repeat every N units.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"frequency": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "How many units between each run.",
																					},
																				},
																			},
																		},
																		"day_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "A simple multiplier schedule — repeat every N units.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"frequency": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "How many units between each run.",
																					},
																				},
																			},
																		},
																		"week_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on specific days of the week. At least one day must be specified.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week on which runs are started. At least one day is required.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},
																		"month_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. \"first Sunday\"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"week_of_month": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Which week of the month the run falls on. Use with `day_of_week`.",
																					},
																					"day_of_month": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.",
																					},
																				},
																			},
																		},
																		"year_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on the first or last day of the year.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_year": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Day of the year when the run starts.",
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
													Computed:    true,
													Description: "Specifies full backup schedule settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"schedule": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Defines when full backup runs are started. The sub-schedule field matching `unit` must be provided: `days` → `day_schedule`, `weeks` → `week_schedule` (with at least one day), `months` → `month_schedule`, `years` → `year_schedule`. The BRS API returns a 400 if the matching sub-schedule is absent.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"unit": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Frequency unit for the full backup schedule.",
																		},
																		"day_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "A simple multiplier schedule — repeat every N units.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"frequency": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "How many units between each run.",
																					},
																				},
																			},
																		},
																		"week_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on specific days of the week. At least one day must be specified.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week on which runs are started. At least one day is required.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},
																		"month_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. \"first Sunday\"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"week_of_month": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Which week of the month the run falls on. Use with `day_of_week`.",
																					},
																					"day_of_month": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.",
																					},
																				},
																			},
																		},
																		"year_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Runs on the first or last day of the year.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_year": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Day of the year when the run starts.",
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
													Computed:    true,
													Description: "Specifies how long backup snapshots are retained.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Retention time unit.",
															},
															"duration": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Number of `unit`s to retain the snapshot.",
															},
														},
													},
												},
											},
										},
									},
									"log": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies log backup schedule settings for database workloads.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Defines how frequently log backup runs are started.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Frequency unit for the log backup schedule.",
															},
															"minute_schedule": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "A simple multiplier schedule — repeat every N units.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "How many units between each run.",
																		},
																	},
																},
															},
															"hour_schedule": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "A simple multiplier schedule — repeat every N units.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"frequency": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "How many units between each run.",
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
						"blackout_window": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of blackout windows during which new runs will not be started.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Day of the week the blackout applies to.",
									},
									"start_time": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a time of day used in scheduling and blackout windows.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hour": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Hour of the day (0-23).",
												},
												"minute": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minute of the hour (0-59).",
												},
												"timezone": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "IANA time zone name. Defaults to `America/Los_Angeles` if omitted.",
												},
											},
										},
									},
									"end_time": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a time of day used in scheduling and blackout windows.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hour": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Hour of the day (0-23).",
												},
												"minute": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minute of the hour (0-59).",
												},
												"timezone": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "IANA time zone name. Defaults to `America/Los_Angeles` if omitted.",
												},
											},
										},
									},
								},
							},
						},
						"extended_retention": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Additional retention policies applied on top of regular retention.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Frequency at which the extended retention rule applies.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Schedule unit for extended retention.",
												},
												"frequency": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Multiplier applied to the unit.",
												},
											},
										},
									},
									"retention": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies how long backup snapshots are retained.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Retention time unit.",
												},
												"duration": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of `unit`s to retain the snapshot.",
												},
											},
										},
									},
									"run_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The backup run type this extended retention applies to.",
									},
								},
							},
						},
						"retry_options": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Controls how failed runs are retried.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retries": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of retry attempts before the run is marked failed.",
									},
									"retry_interval_mins": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Minutes to wait between retry attempts.",
									},
								},
							},
						},
					},
				},
			},
			"scheduling_error": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Non-empty when `state` is `failed` due to async setup failure.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this workload was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last update to this workload.",
			},
			"completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` finished.",
			},
		},
	}
}

func dataSourceIbmBrsMigrationWorkloadRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getWorkloadOptions := &brsmigrationv2.GetWorkloadOptions{}

	getWorkloadOptions.SetMigrationID(d.Get("migration_id").(string))
	getWorkloadOptions.SetWorkloadID(d.Get("workload_id").(string))

	workload, _, err := brsMigrationClient.GetWorkloadWithContext(context, getWorkloadOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetWorkloadWithContext failed: %s", err.Error()), "(Data) ibm_brs_migration_workload", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getWorkloadOptions.MigrationID, *getWorkloadOptions.WorkloadID))

	if !core.IsNil(workload.Name) {
		if err = d.Set("name", workload.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(workload.VolumeOwnershipMap) {
		if err = d.Set("volume_ownership_map", workload.VolumeOwnershipMap); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting volume_ownership_map: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-volume_ownership_map").GetDiag()
		}
	}

	if err = d.Set("state", workload.State); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting state: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-state").GetDiag()
	}

	payloads := []map[string]interface{}{}
	for _, payloadsItem := range workload.Payloads {
		payloadsItemMap, err := DataSourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(&payloadsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload", "read", "payloads-to-map").GetDiag()
		}
		payloads = append(payloads, payloadsItemMap)
	}
	if err = d.Set("payloads", payloads); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting payloads: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-payloads").GetDiag()
	}

	if !core.IsNil(workload.Schedule) {
		schedule := []map[string]interface{}{}
		scheduleMap, err := DataSourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(workload.Schedule)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload", "read", "schedule-to-map").GetDiag()
		}
		schedule = append(schedule, scheduleMap)
		if err = d.Set("schedule", schedule); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting schedule: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-schedule").GetDiag()
		}
	}

	if !core.IsNil(workload.SchedulingError) {
		if err = d.Set("scheduling_error", workload.SchedulingError); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting scheduling_error: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-scheduling_error").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(workload.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(workload.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(workload.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting updated_at: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-updated_at").GetDiag()
		}
	}

	if !core.IsNil(workload.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(workload.CompletedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting completed_at: %s", err), "(Data) ibm_brs_migration_workload", "read", "set-completed_at").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(model *brsmigrationv2.WorkloadPayloadMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["source_host_id"] = *model.SourceHostID
	modelMap["destination_host_id"] = *model.DestinationHostID
	dataSpecs := []map[string]interface{}{}
	for _, dataSpecsItem := range model.DataSpecs {
		dataSpecsItemMap, err := DataSourceIbmBrsMigrationWorkloadDataSpecToMap(&dataSpecsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		dataSpecs = append(dataSpecs, dataSpecsItemMap)
	}
	modelMap["data_specs"] = dataSpecs
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadDataSpecToMap(model *brsmigrationv2.DataSpec) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["data_format"] = *model.DataFormat
	sourceMap, err := DataSourceIbmBrsMigrationWorkloadDataPayloadToMap(model.Source)
	if err != nil {
		return modelMap, err
	}
	modelMap["source"] = []map[string]interface{}{sourceMap}
	destinationMap, err := DataSourceIbmBrsMigrationWorkloadDataPayloadToMap(model.Destination)
	if err != nil {
		return modelMap, err
	}
	modelMap["destination"] = []map[string]interface{}{destinationMap}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadDataPayloadToMap(model *brsmigrationv2.DataPayload) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VolumeID != nil {
		modelMap["volume_id"] = *model.VolumeID
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(model *brsmigrationv2.WorkloadSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	backupPolicyMap, err := DataSourceIbmBrsMigrationWorkloadBackupPolicyToMap(model.BackupPolicy)
	if err != nil {
		return modelMap, err
	}
	modelMap["backup_policy"] = []map[string]interface{}{backupPolicyMap}
	blackoutWindow := []map[string]interface{}{}
	for _, blackoutWindowItem := range model.BlackoutWindow {
		blackoutWindowItemMap, err := DataSourceIbmBrsMigrationWorkloadBlackoutWindowToMap(&blackoutWindowItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		blackoutWindow = append(blackoutWindow, blackoutWindowItemMap)
	}
	modelMap["blackout_window"] = blackoutWindow
	extendedRetention := []map[string]interface{}{}
	for _, extendedRetentionItem := range model.ExtendedRetention {
		extendedRetentionItemMap, err := DataSourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(&extendedRetentionItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		extendedRetention = append(extendedRetention, extendedRetentionItemMap)
	}
	modelMap["extended_retention"] = extendedRetention
	if model.RetryOptions != nil {
		retryOptionsMap, err := DataSourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(model.RetryOptions)
		if err != nil {
			return modelMap, err
		}
		modelMap["retry_options"] = []map[string]interface{}{retryOptionsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadBackupPolicyToMap(model *brsmigrationv2.BackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	regularMap, err := DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(model.Regular)
	if err != nil {
		return modelMap, err
	}
	modelMap["regular"] = []map[string]interface{}{regularMap}
	if model.Log != nil {
		logMap, err := DataSourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(model.Log)
		if err != nil {
			return modelMap, err
		}
		modelMap["log"] = []map[string]interface{}{logMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(model *brsmigrationv2.RegularBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Incremental != nil {
		incrementalMap, err := DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(model.Incremental)
		if err != nil {
			return modelMap, err
		}
		modelMap["incremental"] = []map[string]interface{}{incrementalMap}
	}
	if model.Full != nil {
		fullMap, err := DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(model.Full)
		if err != nil {
			return modelMap, err
		}
		modelMap["full"] = []map[string]interface{}{fullMap}
	}
	if model.Retention != nil {
		retentionMap, err := DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(model *brsmigrationv2.RegularBackupPolicyIncremental) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(model *brsmigrationv2.IncrementalSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	if model.DaySchedule != nil {
		dayScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := DataSourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(model *brsmigrationv2.IncrementalScheduleMinuteSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(model *brsmigrationv2.IncrementalScheduleHourSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(model *brsmigrationv2.IncrementalScheduleDaySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(model *brsmigrationv2.IncrementalScheduleWeekSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(model *brsmigrationv2.IncrementalScheduleMonthSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	modelMap["week_of_month"] = *model.WeekOfMonth
	if model.DayOfMonth != nil {
		modelMap["day_of_month"] = flex.IntValue(model.DayOfMonth)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(model *brsmigrationv2.IncrementalScheduleYearSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_year"] = *model.DayOfYear
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(model *brsmigrationv2.RegularBackupPolicyFull) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Schedule != nil {
		scheduleMap, err := DataSourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(model.Schedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(model *brsmigrationv2.FullBackupPolicySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.DaySchedule != nil {
		dayScheduleMap, err := DataSourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := DataSourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := DataSourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := DataSourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(model *brsmigrationv2.FullScheduleDaySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(model *brsmigrationv2.FullScheduleWeekSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(model *brsmigrationv2.FullScheduleMonthSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	modelMap["week_of_month"] = *model.WeekOfMonth
	if model.DayOfMonth != nil {
		modelMap["day_of_month"] = flex.IntValue(model.DayOfMonth)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(model *brsmigrationv2.FullScheduleYearSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_year"] = *model.DayOfYear
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(model *brsmigrationv2.RegularBackupPolicyRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(model *brsmigrationv2.BackupPolicyLog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Schedule != nil {
		scheduleMap, err := DataSourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(model.Schedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(model *brsmigrationv2.LogBackupPolicySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := DataSourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := DataSourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(model *brsmigrationv2.LogScheduleMinuteSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(model *brsmigrationv2.LogScheduleHourSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadBlackoutWindowToMap(model *brsmigrationv2.BlackoutWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day"] = *model.Day
	startTimeMap, err := DataSourceIbmBrsMigrationWorkloadTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := DataSourceIbmBrsMigrationWorkloadTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadTimeOfDayToMap(model *brsmigrationv2.TimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["hour"] = flex.IntValue(model.Hour)
	modelMap["minute"] = flex.IntValue(model.Minute)
	if model.Timezone != nil {
		modelMap["timezone"] = *model.Timezone
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(model *brsmigrationv2.ExtendedRetentionPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := DataSourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := DataSourceIbmBrsMigrationWorkloadRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	modelMap["run_type"] = *model.RunType
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(model *brsmigrationv2.ExtendedRetentionSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.Frequency != nil {
		modelMap["frequency"] = flex.IntValue(model.Frequency)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRetentionToMap(model *brsmigrationv2.Retention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(model *brsmigrationv2.WorkloadScheduleRetryOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Retries != nil {
		modelMap["retries"] = flex.IntValue(model.Retries)
	}
	if model.RetryIntervalMins != nil {
		modelMap["retry_interval_mins"] = flex.IntValue(model.RetryIntervalMins)
	}
	return modelMap, nil
}
