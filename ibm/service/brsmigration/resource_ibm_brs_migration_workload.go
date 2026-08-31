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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func ResourceIbmBrsMigrationWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmBrsMigrationWorkloadCreate,
		ReadContext:     resourceIbmBrsMigrationWorkloadRead,
		DeleteContext:   resourceIbmBrsMigrationWorkloadDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_workload", "migration_id"),
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_workload", "name"),
				Description: "Human-readable name for this workload.",
			},
			"payloads": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "List of source-to-destination payload mappings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Migration service payload ID (pl-{uuid4} format).",
						},
						"source_host_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Migration service ID of the source host.",
						},
						"destination_host_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Migration service ID of the destination host.",
						},
						"data_specs": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "One or more data specifications for this payload.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_format": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "How this data unit is structured. When `source.volume_id` is a registered `vol-*` ID, the server derives this from the volume registry — pass `raw` as a default if unsure.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Source data location and format for this mapping.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"volume_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Migration service vol-* ID from POST /migrations/{migration_id}/volumes.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Filesystem type, database engine, or raw block device type marker.",
												},
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "OS mount path, database data directory path, or block device path.",
												},
											},
										},
									},
									"destination": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Destination data location and format for this mapping.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"volume_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Migration service vol-* ID from POST /migrations/{migration_id}/volumes.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Filesystem type, database engine, or raw block device type marker.",
												},
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
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
			"volume_ownership_map": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Server-computed map of `vol-*` volume_id → owning payload_id. Only populated when the same `source.volume_id` appears in two or more payloads. Null in all other cases.",
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current lifecycle state of the workload.",
			},
			"schedule": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Populated once setup completes. Null while settingUp or in created/failed state.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Human-readable name for the schedule / protection policy.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
													Optional:    true,
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
																			Optional:    true,
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
																			Optional:    true,
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
																			Optional:    true,
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
																			Optional:    true,
																			Computed:    true,
																			Description: "Runs on specific days of the week. At least one day must be specified.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week on which runs are started. At least one day is required.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																				},
																			},
																		},
																		"month_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Description: "Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. \"first Sunday\"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																					"week_of_month": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Which week of the month the run falls on. Use with `day_of_week`.",
																					},
																					"day_of_month": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Computed:    true,
																						Description: "Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.",
																					},
																				},
																			},
																		},
																		"year_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
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
													Optional:    true,
													Computed:    true,
													Description: "Specifies full backup schedule settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"schedule": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
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
																			Optional:    true,
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
																			Optional:    true,
																			Computed:    true,
																			Description: "Runs on specific days of the week. At least one day must be specified.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week on which runs are started. At least one day is required.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																				},
																			},
																		},
																		"month_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Computed:    true,
																			Description: "Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. \"first Sunday\"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_of_week": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.",
																						Elem:        &schema.Schema{Type: schema.TypeString},
																					},
																					"week_of_month": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Which week of the month the run falls on. Use with `day_of_week`.",
																					},
																					"day_of_month": &schema.Schema{
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Computed:    true,
																						Description: "Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.",
																					},
																				},
																			},
																		},
																		"year_schedule": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
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
													Optional:    true,
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
										Optional:    true,
										Computed:    true,
										Description: "Specifies log backup schedule settings for database workloads.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedule": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
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
																Optional:    true,
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
																Optional:    true,
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
													Optional:    true,
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
													Optional:    true,
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
													Optional:    true,
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
							Optional:    true,
							Computed:    true,
							Description: "Controls how failed runs are retried.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retries": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of retry attempts before the run is marked failed.",
									},
									"retry_interval_mins": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
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
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Migration service workload ID (wl-{uuid4} format).",
			},
		},
	}
}

func ResourceIbmBrsMigrationWorkloadValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "migration_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^mgr-[0-9a-f-]{36}$`,
			MinValueLength:             40,
			MaxValueLength:             40,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^.+$`,
			MinValueLength:             1,
			MaxValueLength:             1024,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_brs_migration_workload", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBrsMigrationWorkloadCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createWorkloadOptions := &brsmigrationv1.CreateWorkloadOptions{}

	createWorkloadOptions.SetMigrationID(d.Get("migration_id").(string))
	var payloads []brsmigrationv1.WorkloadPayloadMappingInput
	for _, v := range d.Get("payloads").([]interface{}) {
		value := v.(map[string]interface{})
		payloadsItem, err := ResourceIbmBrsMigrationWorkloadMapToWorkloadPayloadMappingInput(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "create", "parse-payloads").GetDiag()
		}
		payloads = append(payloads, *payloadsItem)
	}
	createWorkloadOptions.SetPayloads(payloads)
	if _, ok := d.GetOk("name"); ok {
		createWorkloadOptions.SetName(d.Get("name").(string))
	}

	workload, _, err := brsMigrationClient.CreateWorkloadWithContext(context, createWorkloadOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateWorkloadWithContext failed: %s", err.Error()), "ibm_brs_migration_workload", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createWorkloadOptions.MigrationID, *workload.ID))

	return resourceIbmBrsMigrationWorkloadRead(context, d, meta)
}

func resourceIbmBrsMigrationWorkloadRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getWorkloadOptions := &brsmigrationv1.GetWorkloadOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "sep-id-parts").GetDiag()
	}

	getWorkloadOptions.SetMigrationID(parts[0])
	getWorkloadOptions.SetWorkloadID(parts[1])

	workload, response, err := brsMigrationClient.GetWorkloadWithContext(context, getWorkloadOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetWorkloadWithContext failed: %s", err.Error()), "ibm_brs_migration_workload", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(workload.Name) {
		if err = d.Set("name", workload.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-name").GetDiag()
		}
	}
	payloads := []map[string]interface{}{}
	for _, payloadsItem := range workload.Payloads {
		payloadsItemMap, err := ResourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(&payloadsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "payloads-to-map").GetDiag()
		}
		payloads = append(payloads, payloadsItemMap)
	}
	if err = d.Set("payloads", payloads); err != nil {
		err = fmt.Errorf("Error setting payloads: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-payloads").GetDiag()
	}
	if !core.IsNil(workload.VolumeOwnershipMap) {
		volumeOwnershipMap := make(map[string]string)
		for k, v := range workload.VolumeOwnershipMap {
			volumeOwnershipMap[k] = string(v)
		}
		if err = d.Set("volume_ownership_map", volumeOwnershipMap); err != nil {
			err = fmt.Errorf("Error setting volume_ownership_map: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-volume_ownership_map").GetDiag()
		}
	}
	if err = d.Set("state", workload.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-state").GetDiag()
	}
	if !core.IsNil(workload.Schedule) {
		scheduleMap, err := ResourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(workload.Schedule)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "schedule-to-map").GetDiag()
		}
		if err = d.Set("schedule", []map[string]interface{}{scheduleMap}); err != nil {
			err = fmt.Errorf("Error setting schedule: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-schedule").GetDiag()
		}
	}
	if !core.IsNil(workload.SchedulingError) {
		if err = d.Set("scheduling_error", workload.SchedulingError); err != nil {
			err = fmt.Errorf("Error setting scheduling_error: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-scheduling_error").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(workload.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(workload.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(workload.UpdatedAt)); err != nil {
			err = fmt.Errorf("Error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-updated_at").GetDiag()
		}
	}
	if !core.IsNil(workload.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(workload.CompletedAt)); err != nil {
			err = fmt.Errorf("Error setting completed_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-completed_at").GetDiag()
		}
	}
	if err = d.Set("workload_id", workload.ID); err != nil {
		err = fmt.Errorf("Error setting workload_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "read", "set-workload_id").GetDiag()
	}

	return nil
}

func resourceIbmBrsMigrationWorkloadDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteWorkloadOptions := &brsmigrationv1.DeleteWorkloadOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload", "delete", "sep-id-parts").GetDiag()
	}

	deleteWorkloadOptions.SetMigrationID(parts[0])
	deleteWorkloadOptions.SetWorkloadID(parts[1])

	_, err = brsMigrationClient.DeleteWorkloadWithContext(context, deleteWorkloadOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteWorkloadWithContext failed: %s", err.Error()), "ibm_brs_migration_workload", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmBrsMigrationWorkloadMapToWorkloadPayloadMappingInput(modelMap map[string]interface{}) (*brsmigrationv1.WorkloadPayloadMappingInput, error) {
	model := &brsmigrationv1.WorkloadPayloadMappingInput{}
	model.SourceHostID = core.StringPtr(modelMap["source_host_id"].(string))
	model.DestinationHostID = core.StringPtr(modelMap["destination_host_id"].(string))
	dataSpecs := []brsmigrationv1.DataSpec{}
	for _, dataSpecsItem := range modelMap["data_specs"].([]interface{}) {
		dataSpecsItemModel, err := ResourceIbmBrsMigrationWorkloadMapToDataSpec(dataSpecsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		dataSpecs = append(dataSpecs, *dataSpecsItemModel)
	}
	model.DataSpecs = dataSpecs
	return model, nil
}

func ResourceIbmBrsMigrationWorkloadMapToDataSpec(modelMap map[string]interface{}) (*brsmigrationv1.DataSpec, error) {
	model := &brsmigrationv1.DataSpec{}
	model.DataFormat = core.StringPtr(modelMap["data_format"].(string))
	SourceModel, err := ResourceIbmBrsMigrationWorkloadMapToDataPayload(modelMap["source"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Source = SourceModel
	DestinationModel, err := ResourceIbmBrsMigrationWorkloadMapToDataPayload(modelMap["destination"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Destination = DestinationModel
	return model, nil
}

func ResourceIbmBrsMigrationWorkloadMapToDataPayload(modelMap map[string]interface{}) (*brsmigrationv1.DataPayload, error) {
	model := &brsmigrationv1.DataPayload{}
	if modelMap["volume_id"] != nil && modelMap["volume_id"].(string) != "" {
		model.VolumeID = core.StringPtr(modelMap["volume_id"].(string))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	return model, nil
}

func ResourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(model *brsmigrationv1.WorkloadPayloadMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["source_host_id"] = *model.SourceHostID
	modelMap["destination_host_id"] = *model.DestinationHostID
	dataSpecs := []map[string]interface{}{}
	for _, dataSpecsItem := range model.DataSpecs {
		dataSpecsItemMap, err := ResourceIbmBrsMigrationWorkloadDataSpecToMap(&dataSpecsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		dataSpecs = append(dataSpecs, dataSpecsItemMap)
	}
	modelMap["data_specs"] = dataSpecs
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadDataSpecToMap(model *brsmigrationv1.DataSpec) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["data_format"] = *model.DataFormat
	sourceMap, err := ResourceIbmBrsMigrationWorkloadDataPayloadToMap(model.Source)
	if err != nil {
		return modelMap, err
	}
	modelMap["source"] = []map[string]interface{}{sourceMap}
	destinationMap, err := ResourceIbmBrsMigrationWorkloadDataPayloadToMap(model.Destination)
	if err != nil {
		return modelMap, err
	}
	modelMap["destination"] = []map[string]interface{}{destinationMap}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadDataPayloadToMap(model *brsmigrationv1.DataPayload) (map[string]interface{}, error) {
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

func ResourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(model *brsmigrationv1.WorkloadSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	backupPolicyMap, err := ResourceIbmBrsMigrationWorkloadBackupPolicyToMap(model.BackupPolicy)
	if err != nil {
		return modelMap, err
	}
	modelMap["backup_policy"] = []map[string]interface{}{backupPolicyMap}
	blackoutWindow := []map[string]interface{}{}
	for _, blackoutWindowItem := range model.BlackoutWindow {
		blackoutWindowItemMap, err := ResourceIbmBrsMigrationWorkloadBlackoutWindowToMap(&blackoutWindowItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		blackoutWindow = append(blackoutWindow, blackoutWindowItemMap)
	}
	modelMap["blackout_window"] = blackoutWindow
	extendedRetention := []map[string]interface{}{}
	for _, extendedRetentionItem := range model.ExtendedRetention {
		extendedRetentionItemMap, err := ResourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(&extendedRetentionItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		extendedRetention = append(extendedRetention, extendedRetentionItemMap)
	}
	modelMap["extended_retention"] = extendedRetention
	if model.RetryOptions != nil {
		retryOptionsMap, err := ResourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(model.RetryOptions)
		if err != nil {
			return modelMap, err
		}
		modelMap["retry_options"] = []map[string]interface{}{retryOptionsMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadBackupPolicyToMap(model *brsmigrationv1.BackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	regularMap, err := ResourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(model.Regular)
	if err != nil {
		return modelMap, err
	}
	modelMap["regular"] = []map[string]interface{}{regularMap}
	if model.Log != nil {
		logMap, err := ResourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(model.Log)
		if err != nil {
			return modelMap, err
		}
		modelMap["log"] = []map[string]interface{}{logMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(model *brsmigrationv1.RegularBackupPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Incremental != nil {
		incrementalMap, err := ResourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(model.Incremental)
		if err != nil {
			return modelMap, err
		}
		modelMap["incremental"] = []map[string]interface{}{incrementalMap}
	}
	if model.Full != nil {
		fullMap, err := ResourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(model.Full)
		if err != nil {
			return modelMap, err
		}
		modelMap["full"] = []map[string]interface{}{fullMap}
	}
	if model.Retention != nil {
		retentionMap, err := ResourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(model *brsmigrationv1.RegularBackupPolicyIncremental) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(model *brsmigrationv1.IncrementalSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	if model.DaySchedule != nil {
		dayScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := ResourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(model *brsmigrationv1.IncrementalScheduleMinuteSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(model *brsmigrationv1.IncrementalScheduleHourSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(model *brsmigrationv1.IncrementalScheduleDaySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(model *brsmigrationv1.IncrementalScheduleWeekSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(model *brsmigrationv1.IncrementalScheduleMonthSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	modelMap["week_of_month"] = *model.WeekOfMonth
	if model.DayOfMonth != nil {
		modelMap["day_of_month"] = flex.IntValue(model.DayOfMonth)
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(model *brsmigrationv1.IncrementalScheduleYearSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_year"] = *model.DayOfYear
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(model *brsmigrationv1.RegularBackupPolicyFull) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Schedule != nil {
		scheduleMap, err := ResourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(model.Schedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(model *brsmigrationv1.FullBackupPolicySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.DaySchedule != nil {
		dayScheduleMap, err := ResourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(model.DaySchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_schedule"] = []map[string]interface{}{dayScheduleMap}
	}
	if model.WeekSchedule != nil {
		weekScheduleMap, err := ResourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(model.WeekSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["week_schedule"] = []map[string]interface{}{weekScheduleMap}
	}
	if model.MonthSchedule != nil {
		monthScheduleMap, err := ResourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(model.MonthSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["month_schedule"] = []map[string]interface{}{monthScheduleMap}
	}
	if model.YearSchedule != nil {
		yearScheduleMap, err := ResourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(model.YearSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["year_schedule"] = []map[string]interface{}{yearScheduleMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(model *brsmigrationv1.FullScheduleDaySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(model *brsmigrationv1.FullScheduleWeekSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(model *brsmigrationv1.FullScheduleMonthSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	modelMap["week_of_month"] = *model.WeekOfMonth
	if model.DayOfMonth != nil {
		modelMap["day_of_month"] = flex.IntValue(model.DayOfMonth)
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(model *brsmigrationv1.FullScheduleYearSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_year"] = *model.DayOfYear
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(model *brsmigrationv1.RegularBackupPolicyRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(model *brsmigrationv1.BackupPolicyLog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Schedule != nil {
		scheduleMap, err := ResourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(model.Schedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(model *brsmigrationv1.LogBackupPolicySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.MinuteSchedule != nil {
		minuteScheduleMap, err := ResourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(model.MinuteSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["minute_schedule"] = []map[string]interface{}{minuteScheduleMap}
	}
	if model.HourSchedule != nil {
		hourScheduleMap, err := ResourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(model.HourSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["hour_schedule"] = []map[string]interface{}{hourScheduleMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(model *brsmigrationv1.LogScheduleMinuteSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(model *brsmigrationv1.LogScheduleHourSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["frequency"] = flex.IntValue(model.Frequency)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadBlackoutWindowToMap(model *brsmigrationv1.BlackoutWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day"] = *model.Day
	startTimeMap, err := ResourceIbmBrsMigrationWorkloadTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := ResourceIbmBrsMigrationWorkloadTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadTimeOfDayToMap(model *brsmigrationv1.TimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["hour"] = flex.IntValue(model.Hour)
	modelMap["minute"] = flex.IntValue(model.Minute)
	if model.Timezone != nil {
		modelMap["timezone"] = *model.Timezone
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(model *brsmigrationv1.ExtendedRetentionPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scheduleMap, err := ResourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(model.Schedule)
	if err != nil {
		return modelMap, err
	}
	modelMap["schedule"] = []map[string]interface{}{scheduleMap}
	retentionMap, err := ResourceIbmBrsMigrationWorkloadRetentionToMap(model.Retention)
	if err != nil {
		return modelMap, err
	}
	modelMap["retention"] = []map[string]interface{}{retentionMap}
	modelMap["run_type"] = *model.RunType
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(model *brsmigrationv1.ExtendedRetentionSchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	if model.Frequency != nil {
		modelMap["frequency"] = flex.IntValue(model.Frequency)
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRetentionToMap(model *brsmigrationv1.Retention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = *model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(model *brsmigrationv1.WorkloadScheduleRetryOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Retries != nil {
		modelMap["retries"] = flex.IntValue(model.Retries)
	}
	if model.RetryIntervalMins != nil {
		modelMap["retry_interval_mins"] = flex.IntValue(model.RetryIntervalMins)
	}
	return modelMap, nil
}
