// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func ResourceIbmLogsAlertDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsAlertDefinitionCreate,
		ReadContext:   resourceIbmLogsAlertDefinitionRead,
		UpdateContext: resourceIbmLogsAlertDefinitionUpdate,
		DeleteContext: resourceIbmLogsAlertDefinitionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_alert_definition", "name"),
				Description:  "The name of the alert definition.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_alert_definition", "description"),
				Description:  "A detailed description of what the alert monitors and when it triggers.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the alert is currently active and monitoring.",
			},
			"priority": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_alert_definition", "priority"),
				Description:  "The priority of the alert definition.",
			},
			"active_on": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Defining when the alert is active.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day_of_week": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Days of the week when the alert is active.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Start time of the alert activity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hours": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Hours of day in 24 hour format. Should be from 0 to 23.",
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minutes of hour of day. Must be from 0 to 59.",
									},
								},
							},
						},
						"end_time": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Start time of the alert activity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hours": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Hours of day in 24 hour format. Should be from 0 to 23.",
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minutes of hour of day. Must be from 0 to 59.",
									},
								},
							},
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_alert_definition", "type"),
				Description:  "Alert type.",
			},
			"group_by_keys": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Keys used to group and aggregate alert data.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"incidents_settings": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Incident creation and management settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notify_on": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The condition to notify about the alert.",
						},
						"minutes": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The time in minutes before the alert can be retriggered.",
						},
					},
				},
			},
			"notification_group": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Primary notification group for alert events.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_by_keys": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The keys to group the alerts by.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"webhooks": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The settings for webhooks associated with the alert definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notify_on": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The condition to notify about the alert.",
									},
									"integration": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The integration type for webhook notifications.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"integration_id": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The integration ID for the notification.",
												},
											},
										},
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The time in minutes before the notification is sent.",
									},
								},
							},
						},
					},
				},
			},
			"entity_labels": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Labels used to identify and categorize the alert entity.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"phantom_mode": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the alert is in phantom mode (creating incidents or not).",
			},
			"deleted": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the alert has been marked as deleted.",
			},
			"logs_immediate": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for immediate log-based alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"logs_threshold": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for log-based threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Should trigger the alert when undetected values are detected.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the threshold alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the threshold alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "A time window defined by a specific value.",
															},
														},
													},
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The priority of the alert definition.",
												},
											},
										},
									},
								},
							},
						},
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of condition for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"logs_ratio_threshold": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for log-based ratio threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"numerator": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"numerator_alias": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alias for the numerator filter, used for display purposes.",
						},
						"denominator": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"denominator_alias": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alias for the denominator filter, used for display purposes.",
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the ratio alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the ratio alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_ratio_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the time window for the ratio alert.",
															},
														},
													},
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The priority of the alert definition.",
												},
											},
										},
									},
								},
							},
						},
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of condition for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"group_by_for": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The group by settings for the numerator and denominator filters.",
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Should trigger the alert when undetected values are detected.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"ignore_infinity": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The configuration for ignoring infinity values in the ratio.",
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"logs_time_relative_threshold": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for time-relative log threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the time-relative alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the time-relative alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"compared_to": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The time frame to compare the current value against.",
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The priority of the alert definition.",
												},
											},
										},
									},
								},
							},
						},
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"ignore_infinity": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Ignore infinity values in the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Should trigger the alert when undetected values are detected.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"metric_threshold": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for metric-based threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_filter": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The filter to match metric entries for the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"promql": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The filter is a PromQL expression.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the metric threshold alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the metric threshold alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"for_over_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The percentage of values that must exceed the threshold to trigger the alert.",
												},
												"of_the_last": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The time window as a specific value.",
															},
															"metric_time_window_dynamic_duration": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The time window as a dynamic value.",
															},
														},
													},
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The priority of the alert definition.",
												},
											},
										},
									},
								},
							},
						},
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of the alert condition.",
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Should trigger the alert when undetected values are detected.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"missing_values": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Configuration for handling missing values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"replace_with_zero": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to true, missing values will be replaced with zero.",
									},
									"min_non_null_values_pct": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "If set, specifies the minimum percentage of non-null values required for the alert to be triggered.",
									},
								},
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"flow": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for flow-based alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stages": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The stages of the flow alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeframe_ms": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The timeframe for the flow alert in milliseconds.",
									},
									"timeframe_type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of timeframe for the flow alert.",
									},
									"flow_stages_groups": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "Flow stages groups.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"groups": &schema.Schema{
													Type:        schema.TypeList,
													Required:    true,
													Description: "The groups of stages in the flow alert.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alert_defs": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "The alert definitions for the flow stage group.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The alert definition ID.",
																		},
																		"not": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Whether to negate the alert definition or not.",
																		},
																	},
																},
															},
															"next_op": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "The logical operation to apply to the next stage.",
															},
															"alerts_op": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "The logical operation to apply to the alerts in the group.",
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
						"enforce_suppression": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enforce suppression for the flow alert.",
						},
					},
				},
			},
			"logs_anomaly": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for log-based anomaly detection alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the log anomaly alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the anomaly alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"minimum_threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "A time window defined by a specific value.",
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
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of condition for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The notification payload filter to specify which fields to include in the notification.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"anomaly_alert_settings": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The anomaly alert settings configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"percentage_of_deviation": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Description: "The percentage of deviation from the baseline for triggering the alert.",
									},
								},
							},
						},
					},
				},
			},
			"metric_anomaly": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for metric-based anomaly detection alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_filter": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The filter to match metric entries for the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"promql": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The filter is a PromQL expression.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the metric anomaly alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for the metric anomaly alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Required:    true,
													Description: "The threshold value for the alert condition.",
												},
												"for_over_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The percentage of the metric values that must exceed the threshold to trigger the alert.",
												},
												"of_the_last": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The time window as a specific value.",
															},
															"metric_time_window_dynamic_duration": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The time window as a dynamic value.",
															},
														},
													},
												},
												"min_non_null_values_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The percentage of non-null values required to trigger the alert.",
												},
											},
										},
									},
								},
							},
						},
						"condition_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of condition for the alert.",
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"anomaly_alert_settings": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The anomaly alert settings configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"percentage_of_deviation": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Description: "The percentage of deviation from the baseline for triggering the alert.",
									},
								},
							},
						},
					},
				},
			},
			"logs_new_value": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for alerts triggered by new log values.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the log new value alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for detecting new values in logs.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"keypath_to_track": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The keypath to track for new values.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for detecting new values.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_new_value_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "A time window defined by a specific value.",
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
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"logs_unique_count": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration for alerts based on unique log value counts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "/ The value of the label to filter by.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Required:    true,
																Description: "Filter by log severities.",
																Elem:        &schema.Schema{Type: schema.TypeString},
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The rules for the log unique count alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										MinItems:    1,
										MaxItems:    1,
										Required:    true,
										Description: "The condition for detecting unique counts in logs.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_unique_count": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "The maximum unique count for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													MinItems:    1,
													MaxItems:    1,
													Required:    true,
													Description: "The time window for the unique count alert.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_unique_value_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "A time window defined by a specific value.",
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
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The filter to specify which fields to include in the notification payload.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"max_unique_count_per_group_by_key": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The maximum unique count per group by key.",
						},
						"unique_count_keypath": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The keypath in the logs to be used for unique count.",
						},
					},
				},
			},
			"created_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the alert definition was created.",
			},
			"updated_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the alert definition was last updated.",
			},
			"alert_version_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The old alert ID.",
			},
			"alert_def_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alert Definition Id.",
			},
		},
	}
}

func ResourceIbmLogsAlertDefinitionValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$`,
			MinValueLength:             1,
			MaxValueLength:             4096,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$`,
			MinValueLength:             1,
			MaxValueLength:             4096,
		},
		validate.ValidateSchema{
			Identifier:                 "priority",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "p1, p2, p3, p4, p5_or_unspecified",
		},
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "flow, logs_anomaly, logs_immediate_or_unspecified, logs_new_value, logs_ratio_threshold, logs_threshold, logs_time_relative_threshold, logs_unique_count, metric_anomaly, metric_threshold",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_alert_definition", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsAlertDefinitionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	bodyModelMap := map[string]interface{}{}
	createAlertDefOptions := &logsv0.CreateAlertDefOptions{}

	bodyModelMap["name"] = d.Get("name")
	if _, ok := d.GetOk("description"); ok {
		bodyModelMap["description"] = d.Get("description")
	}
	if _, ok := d.GetOkExists("enabled"); ok {
		bodyModelMap["enabled"] = d.Get("enabled")
	}
	if _, ok := d.GetOk("priority"); ok {
		bodyModelMap["priority"] = d.Get("priority")
	}
	if _, ok := d.GetOk("active_on"); ok {
		bodyModelMap["active_on"] = d.Get("active_on")
	}
	bodyModelMap["type"] = d.Get("type")
	bodyModelMap["group_by_keys"] = d.Get("group_by_keys")
	if _, ok := d.GetOk("incidents_settings"); ok {
		bodyModelMap["incidents_settings"] = d.Get("incidents_settings")
	}
	if _, ok := d.GetOk("notification_group"); ok {
		bodyModelMap["notification_group"] = d.Get("notification_group")
	}
	if _, ok := d.GetOk("entity_labels"); ok {
		bodyModelMap["entity_labels"] = d.Get("entity_labels")
	}
	if _, ok := d.GetOk("phantom_mode"); ok {
		bodyModelMap["phantom_mode"] = d.Get("phantom_mode")
	}
	if _, ok := d.GetOk("deleted"); ok {
		bodyModelMap["deleted"] = d.Get("deleted")
	}
	if _, ok := d.GetOk("logs_immediate"); ok {
		bodyModelMap["logs_immediate"] = d.Get("logs_immediate")
	}
	if _, ok := d.GetOk("logs_threshold"); ok {
		bodyModelMap["logs_threshold"] = d.Get("logs_threshold")
	}
	if _, ok := d.GetOk("logs_ratio_threshold"); ok {
		bodyModelMap["logs_ratio_threshold"] = d.Get("logs_ratio_threshold")
	}
	if _, ok := d.GetOk("logs_time_relative_threshold"); ok {
		bodyModelMap["logs_time_relative_threshold"] = d.Get("logs_time_relative_threshold")
	}
	if _, ok := d.GetOk("metric_threshold"); ok {
		bodyModelMap["metric_threshold"] = d.Get("metric_threshold")
	}
	if _, ok := d.GetOk("flow"); ok {
		bodyModelMap["flow"] = d.Get("flow")
	}
	if _, ok := d.GetOk("logs_anomaly"); ok {
		bodyModelMap["logs_anomaly"] = d.Get("logs_anomaly")
	}
	if _, ok := d.GetOk("metric_anomaly"); ok {
		bodyModelMap["metric_anomaly"] = d.Get("metric_anomaly")
	}
	if _, ok := d.GetOk("logs_new_value"); ok {
		bodyModelMap["logs_new_value"] = d.Get("logs_new_value")
	}
	if _, ok := d.GetOk("logs_unique_count"); ok {
		bodyModelMap["logs_unique_count"] = d.Get("logs_unique_count")
	}
	convertedModel, err := ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "create", "parse-request-body").GetDiag()
	}
	createAlertDefOptions.AlertDefinitionPrototype = convertedModel

	alertDefinitionIntf, _, err := logsClient.CreateAlertDefWithContext(context, createAlertDefOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAlertDefWithContext failed: %s", err.Error()), "ibm_logs_alert_definition", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)
	alertDefId := fmt.Sprintf("%s/%s/%s", region, instanceId, *alertDefinition.ID)
	d.SetId(alertDefId)

	return resourceIbmLogsAlertDefinitionRead(context, d, meta)
}

func resourceIbmLogsAlertDefinitionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, region, instanceId, alertDefId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}
	getAlertDefOptions := &logsv0.GetAlertDefOptions{}

	getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(alertDefId)))

	alertDefinitionIntf, response, err := logsClient.GetAlertDefWithContext(context, getAlertDefOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAlertDefWithContext failed: %s", err.Error()), "ibm_logs_alert_definition", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)

	if err = d.Set("alert_def_id", alertDefId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting alert_def_id: %s", err))
	}
	if err = d.Set("instance_id", instanceId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if err = d.Set("region", region); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
	}

	if err = d.Set("name", alertDefinition.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-name").GetDiag()
	}
	if !core.IsNil(alertDefinition.Description) {
		if err = d.Set("description", alertDefinition.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.Enabled) {
		if err = d.Set("enabled", alertDefinition.Enabled); err != nil {
			err = fmt.Errorf("Error setting enabled: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-enabled").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.Priority) {
		if err = d.Set("priority", alertDefinition.Priority); err != nil {
			err = fmt.Errorf("Error setting priority: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-priority").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.ActiveOn) {
		activeOnMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(alertDefinition.ActiveOn)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "active_on-to-map").GetDiag()
		}
		if err = d.Set("active_on", []map[string]interface{}{activeOnMap}); err != nil {
			err = fmt.Errorf("Error setting active_on: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-active_on").GetDiag()
		}
	}
	if err = d.Set("type", alertDefinition.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-type").GetDiag()
	}
	if err = d.Set("group_by_keys", alertDefinition.GroupByKeys); err != nil {
		err = fmt.Errorf("Error setting group_by_keys: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-group_by_keys").GetDiag()
	}
	if !core.IsNil(alertDefinition.IncidentsSettings) {
		incidentsSettingsMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(alertDefinition.IncidentsSettings)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "incidents_settings-to-map").GetDiag()
		}
		if err = d.Set("incidents_settings", []map[string]interface{}{incidentsSettingsMap}); err != nil {
			err = fmt.Errorf("Error setting incidents_settings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-incidents_settings").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.NotificationGroup) {
		notificationGroupMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(alertDefinition.NotificationGroup)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "notification_group-to-map").GetDiag()
		}
		if err = d.Set("notification_group", []map[string]interface{}{notificationGroupMap}); err != nil {
			err = fmt.Errorf("Error setting notification_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-notification_group").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.EntityLabels) {
		entityLabels := make(map[string]string)
		for k, v := range alertDefinition.EntityLabels {
			entityLabels[k] = string(v)
		}
		if err = d.Set("entity_labels", entityLabels); err != nil {
			err = fmt.Errorf("Error setting entity_labels: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-entity_labels").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.PhantomMode) {
		if err = d.Set("phantom_mode", alertDefinition.PhantomMode); err != nil {
			err = fmt.Errorf("Error setting phantom_mode: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-phantom_mode").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.Deleted) {
		if err = d.Set("deleted", alertDefinition.Deleted); err != nil {
			err = fmt.Errorf("Error setting deleted: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-deleted").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsImmediate) {
		logsImmediateMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(alertDefinition.LogsImmediate)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_immediate-to-map").GetDiag()
		}
		if err = d.Set("logs_immediate", []map[string]interface{}{logsImmediateMap}); err != nil {
			err = fmt.Errorf("Error setting logs_immediate: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_immediate").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsThreshold) {
		logsThresholdMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(alertDefinition.LogsThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_threshold-to-map").GetDiag()
		}
		if err = d.Set("logs_threshold", []map[string]interface{}{logsThresholdMap}); err != nil {
			err = fmt.Errorf("Error setting logs_threshold: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_threshold").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsRatioThreshold) {
		logsRatioThresholdMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(alertDefinition.LogsRatioThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_ratio_threshold-to-map").GetDiag()
		}
		if err = d.Set("logs_ratio_threshold", []map[string]interface{}{logsRatioThresholdMap}); err != nil {
			err = fmt.Errorf("Error setting logs_ratio_threshold: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_ratio_threshold").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsTimeRelativeThreshold) {
		logsTimeRelativeThresholdMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(alertDefinition.LogsTimeRelativeThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_time_relative_threshold-to-map").GetDiag()
		}
		if err = d.Set("logs_time_relative_threshold", []map[string]interface{}{logsTimeRelativeThresholdMap}); err != nil {
			err = fmt.Errorf("Error setting logs_time_relative_threshold: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_time_relative_threshold").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.MetricThreshold) {
		metricThresholdMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(alertDefinition.MetricThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "metric_threshold-to-map").GetDiag()
		}
		if err = d.Set("metric_threshold", []map[string]interface{}{metricThresholdMap}); err != nil {
			err = fmt.Errorf("Error setting metric_threshold: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-metric_threshold").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.Flow) {
		flowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(alertDefinition.Flow)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "flow-to-map").GetDiag()
		}
		if err = d.Set("flow", []map[string]interface{}{flowMap}); err != nil {
			err = fmt.Errorf("Error setting flow: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-flow").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsAnomaly) {
		logsAnomalyMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(alertDefinition.LogsAnomaly)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_anomaly-to-map").GetDiag()
		}
		if err = d.Set("logs_anomaly", []map[string]interface{}{logsAnomalyMap}); err != nil {
			err = fmt.Errorf("Error setting logs_anomaly: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_anomaly").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.MetricAnomaly) {
		metricAnomalyMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(alertDefinition.MetricAnomaly)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "metric_anomaly-to-map").GetDiag()
		}
		if err = d.Set("metric_anomaly", []map[string]interface{}{metricAnomalyMap}); err != nil {
			err = fmt.Errorf("Error setting metric_anomaly: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-metric_anomaly").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsNewValue) {
		logsNewValueMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(alertDefinition.LogsNewValue)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_new_value-to-map").GetDiag()
		}
		if err = d.Set("logs_new_value", []map[string]interface{}{logsNewValueMap}); err != nil {
			err = fmt.Errorf("Error setting logs_new_value: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_new_value").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.LogsUniqueCount) {
		logsUniqueCountMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(alertDefinition.LogsUniqueCount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "logs_unique_count-to-map").GetDiag()
		}
		if err = d.Set("logs_unique_count", []map[string]interface{}{logsUniqueCountMap}); err != nil {
			err = fmt.Errorf("Error setting logs_unique_count: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-logs_unique_count").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.CreatedTime) {
		if err = d.Set("created_time", flex.DateTimeToString(alertDefinition.CreatedTime)); err != nil {
			err = fmt.Errorf("Error setting created_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-created_time").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.UpdatedTime) {
		if err = d.Set("updated_time", flex.DateTimeToString(alertDefinition.UpdatedTime)); err != nil {
			err = fmt.Errorf("Error setting updated_time: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-updated_time").GetDiag()
		}
	}
	if !core.IsNil(alertDefinition.AlertVersionID) {
		if err = d.Set("alert_version_id", alertDefinition.AlertVersionID); err != nil {
			err = fmt.Errorf("Error setting alert_version_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "read", "set-alert_version_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmLogsAlertDefinitionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, _, _, alertdefId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}
	replaceAlertDefOptions := &logsv0.ReplaceAlertDefOptions{}

	replaceAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(alertdefId)))

	hasChange := false
	bodyModelMap := map[string]interface{}{}

	if d.HasChange("name") ||
		d.HasChange("description") ||
		d.HasChange("enabled") ||
		d.HasChange("priority") ||
		d.HasChange("active_on") ||
		d.HasChange("type") ||
		d.HasChange("group_by_keys") ||
		d.HasChange("incidents_settings") ||
		d.HasChange("notification_group") ||
		d.HasChange("entity_labels") ||
		d.HasChange("phantom_mode") ||
		d.HasChange("deleted") ||
		d.HasChange("logs_immediate") ||
		d.HasChange("logs_threshold") ||
		d.HasChange("logs_ratio_threshold") ||
		d.HasChange("logs_time_relative_threshold") ||
		d.HasChange("metric_threshold") ||
		d.HasChange("flow") ||
		d.HasChange("logs_anomaly") ||
		d.HasChange("metric_anomaly") ||
		d.HasChange("logs_new_value") ||
		d.HasChange("logs_unique_count") {

		bodyModelMap["name"] = d.Get("name")
		if _, ok := d.GetOk("description"); ok {
			bodyModelMap["description"] = d.Get("description")
		}
		if _, ok := d.GetOkExists("enabled"); ok {
			bodyModelMap["enabled"] = d.Get("enabled")
		}
		if _, ok := d.GetOk("priority"); ok {
			bodyModelMap["priority"] = d.Get("priority")
		}
		if _, ok := d.GetOk("active_on"); ok {
			bodyModelMap["active_on"] = d.Get("active_on")
		}
		bodyModelMap["type"] = d.Get("type")
		bodyModelMap["group_by_keys"] = d.Get("group_by_keys")
		if _, ok := d.GetOk("incidents_settings"); ok {
			bodyModelMap["incidents_settings"] = d.Get("incidents_settings")
		}
		if _, ok := d.GetOk("notification_group"); ok {
			bodyModelMap["notification_group"] = d.Get("notification_group")
		}
		if _, ok := d.GetOk("entity_labels"); ok {
			bodyModelMap["entity_labels"] = d.Get("entity_labels")
		}
		if _, ok := d.GetOk("phantom_mode"); ok {
			bodyModelMap["phantom_mode"] = d.Get("phantom_mode")
		}
		if _, ok := d.GetOk("deleted"); ok {
			bodyModelMap["deleted"] = d.Get("deleted")
		}
		if _, ok := d.GetOk("logs_immediate"); ok {
			bodyModelMap["logs_immediate"] = d.Get("logs_immediate")
		}
		if _, ok := d.GetOk("logs_threshold"); ok {
			bodyModelMap["logs_threshold"] = d.Get("logs_threshold")
		}
		if _, ok := d.GetOk("logs_ratio_threshold"); ok {
			bodyModelMap["logs_ratio_threshold"] = d.Get("logs_ratio_threshold")
		}
		if _, ok := d.GetOk("logs_time_relative_threshold"); ok {
			bodyModelMap["logs_time_relative_threshold"] = d.Get("logs_time_relative_threshold")
		}
		if _, ok := d.GetOk("metric_threshold"); ok {
			bodyModelMap["metric_threshold"] = d.Get("metric_threshold")
		}
		if _, ok := d.GetOk("flow"); ok {
			bodyModelMap["flow"] = d.Get("flow")
		}
		if _, ok := d.GetOk("logs_anomaly"); ok {
			bodyModelMap["logs_anomaly"] = d.Get("logs_anomaly")
		}
		if _, ok := d.GetOk("metric_anomaly"); ok {
			bodyModelMap["metric_anomaly"] = d.Get("metric_anomaly")
		}
		if _, ok := d.GetOk("logs_new_value"); ok {
			bodyModelMap["logs_new_value"] = d.Get("logs_new_value")
		}
		if _, ok := d.GetOk("logs_unique_count"); ok {
			bodyModelMap["logs_unique_count"] = d.Get("logs_unique_count")
		}
		convertedModel, err := ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototype(bodyModelMap)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "create", "parse-request-body").GetDiag()
		}
		replaceAlertDefOptions.SetAlertDefinitionPrototype(convertedModel)
		hasChange = true
	}
	if hasChange {
		_, _, err = logsClient.ReplaceAlertDefWithContext(context, replaceAlertDefOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceAlertDefWithContext failed: %s", err.Error()), "ibm_logs_alert_definition", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmLogsAlertDefinitionRead(context, d, meta)
}

func resourceIbmLogsAlertDefinitionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_alert_definition", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	logsClient, _, _, alertDefId, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		return diag.FromErr(err)
	}
	deleteAlertDefOptions := &logsv0.DeleteAlertDefOptions{}

	deleteAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(alertDefId)))

	_, err = logsClient.DeleteAlertDefWithContext(context, deleteAlertDefOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAlertDefWithContext failed: %s", err.Error()), "ibm_logs_alert_definition", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionActivitySchedule, error) {
	model := &logsv0.ApisAlertDefinitionActivitySchedule{}
	dayOfWeek := []string{}
	for _, dayOfWeekItem := range modelMap["day_of_week"].([]interface{}) {
		dayOfWeek = append(dayOfWeek, dayOfWeekItem.(string))
	}
	model.DayOfWeek = dayOfWeek
	StartTimeModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionTimeOfDay(modelMap["start_time"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.StartTime = StartTimeModel
	EndTimeModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionTimeOfDay(modelMap["end_time"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.EndTime = EndTimeModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionTimeOfDay(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionTimeOfDay, error) {
	model := &logsv0.ApisAlertDefinitionTimeOfDay{}
	if modelMap["hours"] != nil {
		model.Hours = core.Int64Ptr(int64(modelMap["hours"].(int)))
	}
	if modelMap["minutes"] != nil {
		model.Minutes = core.Int64Ptr(int64(modelMap["minutes"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionAlertDefIncidentSettings, error) {
	model := &logsv0.ApisAlertDefinitionAlertDefIncidentSettings{}
	if modelMap["notify_on"] != nil && modelMap["notify_on"].(string) != "" {
		model.NotifyOn = core.StringPtr(modelMap["notify_on"].(string))
	}
	if modelMap["minutes"] != nil {
		model.Minutes = core.Int64Ptr(int64(modelMap["minutes"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionAlertDefNotificationGroup, error) {
	model := &logsv0.ApisAlertDefinitionAlertDefNotificationGroup{}
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	webhooks := []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{}
	for _, webhooksItem := range modelMap["webhooks"].([]interface{}) {
		webhooksItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefWebhooksSettings(webhooksItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		webhooks = append(webhooks, *webhooksItemModel)
	}
	model.Webhooks = webhooks
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefWebhooksSettings(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionAlertDefWebhooksSettings, error) {
	model := &logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{}
	if modelMap["notify_on"] != nil && modelMap["notify_on"].(string) != "" {
		model.NotifyOn = core.StringPtr(modelMap["notify_on"].(string))
	}
	IntegrationModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationType(modelMap["integration"].([]interface{}))
	if err != nil {
		return model, err
	}
	model.Integration = IntegrationModel
	if modelMap["minutes"] != nil && modelMap["minutes"] != 0 { // manual change: as tf by default takes int value as 0 if not provided by user.
		model.Minutes = core.Int64Ptr(int64(modelMap["minutes"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationType(modelMap []interface{}) (logsv0.ApisAlertDefinitionIntegrationTypeIntf, error) {
	model := &logsv0.ApisAlertDefinitionIntegrationType{}
	if len(modelMap) > 0 && modelMap[0] != nil {
		modelMapElement := modelMap[0].(map[string]interface{})
		if modelMapElement["integration_id"] != nil {
			model.IntegrationID = core.Int64Ptr(int64(modelMapElement["integration_id"].(int)))
		}
	}

	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID, error) {
	model := &logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID{}
	if modelMap["integration_id"] != nil {
		model.IntegrationID = core.Int64Ptr(int64(modelMap["integration_id"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsImmediateType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsImmediateType, error) {
	model := &logsv0.ApisAlertDefinitionLogsImmediateType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsFilter, error) {
	model := &logsv0.ApisAlertDefinitionLogsFilter{}
	if modelMap["simple_filter"] != nil && len(modelMap["simple_filter"].([]interface{})) > 0 {
		SimpleFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsSimpleFilter(modelMap["simple_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SimpleFilter = SimpleFilterModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsSimpleFilter(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsSimpleFilter, error) {
	model := &logsv0.ApisAlertDefinitionLogsSimpleFilter{}
	if modelMap["lucene_query"] != nil && modelMap["lucene_query"].(string) != "" {
		model.LuceneQuery = core.StringPtr(modelMap["lucene_query"].(string))
	}
	if modelMap["label_filters"] != nil && len(modelMap["label_filters"].([]interface{})) > 0 {
		LabelFiltersModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilters(modelMap["label_filters"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LabelFilters = LabelFiltersModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilters(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLabelFilters, error) {
	model := &logsv0.ApisAlertDefinitionLabelFilters{}
	applicationName := []logsv0.ApisAlertDefinitionLabelFilterType{}
	for _, applicationNameItem := range modelMap["application_name"].([]interface{}) {
		applicationNameItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilterType(applicationNameItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		applicationName = append(applicationName, *applicationNameItemModel)
	}
	model.ApplicationName = applicationName
	subsystemName := []logsv0.ApisAlertDefinitionLabelFilterType{}
	for _, subsystemNameItem := range modelMap["subsystem_name"].([]interface{}) {
		subsystemNameItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilterType(subsystemNameItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		subsystemName = append(subsystemName, *subsystemNameItemModel)
	}
	model.SubsystemName = subsystemName
	severities := []string{}
	for _, severitiesItem := range modelMap["severities"].([]interface{}) {
		severities = append(severities, severitiesItem.(string))
	}
	model.Severities = severities
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilterType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLabelFilterType, error) {
	model := &logsv0.ApisAlertDefinitionLabelFilterType{}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	model.Operation = core.StringPtr(modelMap["operation"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsThresholdType, error) {
	model := &logsv0.ApisAlertDefinitionLogsThresholdType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	if modelMap["undetected_values_management"] != nil && len(modelMap["undetected_values_management"].([]interface{})) > 0 {
		UndetectedValuesManagementModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(modelMap["undetected_values_management"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.UndetectedValuesManagement = UndetectedValuesManagementModel
	}
	rules := []logsv0.ApisAlertDefinitionLogsThresholdRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionUndetectedValuesManagement, error) {
	model := &logsv0.ApisAlertDefinitionUndetectedValuesManagement{}
	model.TriggerUndetectedValues = core.BoolPtr(modelMap["trigger_undetected_values"].(bool))
	model.AutoRetireTimeframe = core.StringPtr(modelMap["auto_retire_timeframe"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsThresholdRule, error) {
	model := &logsv0.ApisAlertDefinitionLogsThresholdRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	OverrideModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(modelMap["override"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Override = OverrideModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsThresholdCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsThresholdCondition{}
	model.Threshold = core.Float64Ptr(modelMap["threshold"].(float64))
	TimeWindowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeWindow(modelMap["time_window"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.TimeWindow = TimeWindowModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeWindow(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsTimeWindow, error) {
	model := &logsv0.ApisAlertDefinitionLogsTimeWindow{}
	model.LogsTimeWindowSpecificValue = core.StringPtr(modelMap["logs_time_window_specific_value"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionAlertDefOverride, error) {
	model := &logsv0.ApisAlertDefinitionAlertDefOverride{}
	model.Priority = core.StringPtr(modelMap["priority"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioThresholdType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsRatioThresholdType, error) {
	model := &logsv0.ApisAlertDefinitionLogsRatioThresholdType{}
	NumeratorModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["numerator"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Numerator = NumeratorModel
	if modelMap["numerator_alias"] != nil && modelMap["numerator_alias"].(string) != "" {
		model.NumeratorAlias = core.StringPtr(modelMap["numerator_alias"].(string))
	}
	DenominatorModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["denominator"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Denominator = DenominatorModel
	if modelMap["denominator_alias"] != nil && modelMap["denominator_alias"].(string) != "" {
		model.DenominatorAlias = core.StringPtr(modelMap["denominator_alias"].(string))
	}
	rules := []logsv0.ApisAlertDefinitionLogsRatioRules{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioRules(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	model.GroupByFor = core.StringPtr(modelMap["group_by_for"].(string))
	if modelMap["undetected_values_management"] != nil && len(modelMap["undetected_values_management"].([]interface{})) > 0 {
		UndetectedValuesManagementModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(modelMap["undetected_values_management"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.UndetectedValuesManagement = UndetectedValuesManagementModel
	}
	if modelMap["ignore_infinity"] != nil {
		model.IgnoreInfinity = core.BoolPtr(modelMap["ignore_infinity"].(bool))
	}
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioRules(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsRatioRules, error) {
	model := &logsv0.ApisAlertDefinitionLogsRatioRules{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	OverrideModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(modelMap["override"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Override = OverrideModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsRatioCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsRatioCondition{}
	model.Threshold = core.Float64Ptr(modelMap["threshold"].(float64))
	TimeWindowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioTimeWindow(modelMap["time_window"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.TimeWindow = TimeWindowModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioTimeWindow(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsRatioTimeWindow, error) {
	model := &logsv0.ApisAlertDefinitionLogsRatioTimeWindow{}
	model.LogsRatioTimeWindowSpecificValue = core.StringPtr(modelMap["logs_ratio_time_window_specific_value"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeThresholdType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType, error) {
	model := &logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	rules := []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["ignore_infinity"] != nil {
		model.IgnoreInfinity = core.BoolPtr(modelMap["ignore_infinity"].(bool))
	}
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	if modelMap["undetected_values_management"] != nil && len(modelMap["undetected_values_management"].([]interface{})) > 0 {
		UndetectedValuesManagementModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(modelMap["undetected_values_management"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.UndetectedValuesManagement = UndetectedValuesManagementModel
	}
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsTimeRelativeRule, error) {
	model := &logsv0.ApisAlertDefinitionLogsTimeRelativeRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	OverrideModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(modelMap["override"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Override = OverrideModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsTimeRelativeCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsTimeRelativeCondition{}
	model.Threshold = core.Float64Ptr(modelMap["threshold"].(float64))
	model.ComparedTo = core.StringPtr(modelMap["compared_to"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricThresholdType, error) {
	model := &logsv0.ApisAlertDefinitionMetricThresholdType{}
	MetricFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricFilter(modelMap["metric_filter"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.MetricFilter = MetricFilterModel
	rules := []logsv0.ApisAlertDefinitionMetricThresholdRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["undetected_values_management"] != nil && len(modelMap["undetected_values_management"].([]interface{})) > 0 {
		UndetectedValuesManagementModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(modelMap["undetected_values_management"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.UndetectedValuesManagement = UndetectedValuesManagementModel
	}
	MissingValuesModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValues(modelMap["missing_values"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.MissingValues = MissingValuesModel
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricFilter(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricFilter, error) {
	model := &logsv0.ApisAlertDefinitionMetricFilter{}
	model.Promql = core.StringPtr(modelMap["promql"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricThresholdRule, error) {
	model := &logsv0.ApisAlertDefinitionMetricThresholdRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	OverrideModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(modelMap["override"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Override = OverrideModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricThresholdCondition, error) {
	model := &logsv0.ApisAlertDefinitionMetricThresholdCondition{}
	model.Threshold = core.Float64Ptr(modelMap["threshold"].(float64))
	model.ForOverPct = core.Int64Ptr(int64(modelMap["for_over_pct"].(int)))
	OfTheLastModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindow(modelMap["of_the_last"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.OfTheLast = OfTheLastModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindow(modelMap map[string]interface{}) (logsv0.ApisAlertDefinitionMetricTimeWindowIntf, error) {
	model := &logsv0.ApisAlertDefinitionMetricTimeWindow{}
	if modelMap["metric_time_window_specific_value"] != nil && modelMap["metric_time_window_specific_value"].(string) != "" {
		model.MetricTimeWindowSpecificValue = core.StringPtr(modelMap["metric_time_window_specific_value"].(string))
	}
	if modelMap["metric_time_window_dynamic_duration"] != nil && modelMap["metric_time_window_dynamic_duration"].(string) != "" {
		model.MetricTimeWindowDynamicDuration = core.StringPtr(modelMap["metric_time_window_dynamic_duration"].(string))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue, error) {
	model := &logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue{}
	if modelMap["metric_time_window_specific_value"] != nil && modelMap["metric_time_window_specific_value"].(string) != "" {
		model.MetricTimeWindowSpecificValue = core.StringPtr(modelMap["metric_time_window_specific_value"].(string))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration, error) {
	model := &logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration{}
	if modelMap["metric_time_window_dynamic_duration"] != nil && modelMap["metric_time_window_dynamic_duration"].(string) != "" {
		model.MetricTimeWindowDynamicDuration = core.StringPtr(modelMap["metric_time_window_dynamic_duration"].(string))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValues(modelMap map[string]interface{}) (logsv0.ApisAlertDefinitionMetricMissingValuesIntf, error) {
	model := &logsv0.ApisAlertDefinitionMetricMissingValues{}
	if modelMap["replace_with_zero"] != nil {
		model.ReplaceWithZero = core.BoolPtr(modelMap["replace_with_zero"].(bool))
	}
	if modelMap["min_non_null_values_pct"] != nil {
		model.MinNonNullValuesPct = core.Int64Ptr(int64(modelMap["min_non_null_values_pct"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero, error) {
	model := &logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero{}
	if modelMap["replace_with_zero"] != nil {
		model.ReplaceWithZero = core.BoolPtr(modelMap["replace_with_zero"].(bool))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct, error) {
	model := &logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct{}
	if modelMap["min_non_null_values_pct"] != nil {
		model.MinNonNullValuesPct = core.Int64Ptr(int64(modelMap["min_non_null_values_pct"].(int)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionFlowType, error) {
	model := &logsv0.ApisAlertDefinitionFlowType{}
	stages := []logsv0.ApisAlertDefinitionFlowStages{}
	for _, stagesItem := range modelMap["stages"].([]interface{}) {
		stagesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStages(stagesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		stages = append(stages, *stagesItemModel)
	}
	model.Stages = stages
	if modelMap["enforce_suppression"] != nil {
		model.EnforceSuppression = core.BoolPtr(modelMap["enforce_suppression"].(bool))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStages(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionFlowStages, error) {
	model := &logsv0.ApisAlertDefinitionFlowStages{}
	model.TimeframeMs = core.StringPtr(modelMap["timeframe_ms"].(string))
	model.TimeframeType = core.StringPtr(modelMap["timeframe_type"].(string))
	FlowStagesGroupsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroups(modelMap["flow_stages_groups"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.FlowStagesGroups = FlowStagesGroupsModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroups(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionFlowStagesGroups, error) {
	model := &logsv0.ApisAlertDefinitionFlowStagesGroups{}
	groups := []logsv0.ApisAlertDefinitionFlowStagesGroup{}
	for _, groupsItem := range modelMap["groups"].([]interface{}) {
		groupsItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroup(groupsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		groups = append(groups, *groupsItemModel)
	}
	model.Groups = groups
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroup(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionFlowStagesGroup, error) {
	model := &logsv0.ApisAlertDefinitionFlowStagesGroup{}
	alertDefs := []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{}
	for _, alertDefsItem := range modelMap["alert_defs"].([]interface{}) {
		alertDefsItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroupsAlertDefs(alertDefsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		alertDefs = append(alertDefs, *alertDefsItemModel)
	}
	model.AlertDefs = alertDefs
	model.NextOp = core.StringPtr(modelMap["next_op"].(string))
	model.AlertsOp = core.StringPtr(modelMap["alerts_op"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroupsAlertDefs(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs, error) {
	model := &logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{}
	model.ID = core.UUIDPtr(strfmt.UUID(modelMap["id"].(string)))
	if modelMap["not"] != nil {
		model.Not = core.BoolPtr(modelMap["not"].(bool))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsAnomalyType, error) {
	model := &logsv0.ApisAlertDefinitionLogsAnomalyType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	rules := []logsv0.ApisAlertDefinitionLogsAnomalyRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	if modelMap["anomaly_alert_settings"] != nil && len(modelMap["anomaly_alert_settings"].([]interface{})) > 0 {
		AnomalyAlertSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAnomalyAlertSettings(modelMap["anomaly_alert_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.AnomalyAlertSettings = AnomalyAlertSettingsModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsAnomalyRule, error) {
	model := &logsv0.ApisAlertDefinitionLogsAnomalyRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsAnomalyCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsAnomalyCondition{}
	model.MinimumThreshold = core.Float64Ptr(modelMap["minimum_threshold"].(float64))
	TimeWindowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeWindow(modelMap["time_window"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.TimeWindow = TimeWindowModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAnomalyAlertSettings(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionAnomalyAlertSettings, error) {
	model := &logsv0.ApisAlertDefinitionAnomalyAlertSettings{}
	if modelMap["percentage_of_deviation"] != nil {
		model.PercentageOfDeviation = core.Float32Ptr(float32(modelMap["percentage_of_deviation"].(float64)))
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricAnomalyType, error) {
	model := &logsv0.ApisAlertDefinitionMetricAnomalyType{}
	MetricFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricFilter(modelMap["metric_filter"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.MetricFilter = MetricFilterModel
	rules := []logsv0.ApisAlertDefinitionMetricAnomalyRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	model.ConditionType = core.StringPtr(modelMap["condition_type"].(string))
	if modelMap["evaluation_delay_ms"] != nil {
		model.EvaluationDelayMs = core.Int64Ptr(int64(modelMap["evaluation_delay_ms"].(int)))
	}
	if modelMap["anomaly_alert_settings"] != nil && len(modelMap["anomaly_alert_settings"].([]interface{})) > 0 {
		AnomalyAlertSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAnomalyAlertSettings(modelMap["anomaly_alert_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.AnomalyAlertSettings = AnomalyAlertSettingsModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricAnomalyRule, error) {
	model := &logsv0.ApisAlertDefinitionMetricAnomalyRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionMetricAnomalyCondition, error) {
	model := &logsv0.ApisAlertDefinitionMetricAnomalyCondition{}
	model.Threshold = core.Float64Ptr(modelMap["threshold"].(float64))
	if modelMap["for_over_pct"] != nil {
		model.ForOverPct = core.Int64Ptr(int64(modelMap["for_over_pct"].(int)))
	}
	OfTheLastModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindow(modelMap["of_the_last"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.OfTheLast = OfTheLastModel
	model.MinNonNullValuesPct = core.Int64Ptr(int64(modelMap["min_non_null_values_pct"].(int)))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsNewValueType, error) {
	model := &logsv0.ApisAlertDefinitionLogsNewValueType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	rules := []logsv0.ApisAlertDefinitionLogsNewValueRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsNewValueRule, error) {
	model := &logsv0.ApisAlertDefinitionLogsNewValueRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsNewValueCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsNewValueCondition{}
	model.KeypathToTrack = core.StringPtr(modelMap["keypath_to_track"].(string))
	TimeWindowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueTimeWindow(modelMap["time_window"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.TimeWindow = TimeWindowModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueTimeWindow(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsNewValueTimeWindow, error) {
	model := &logsv0.ApisAlertDefinitionLogsNewValueTimeWindow{}
	model.LogsNewValueTimeWindowSpecificValue = core.StringPtr(modelMap["logs_new_value_time_window_specific_value"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountType(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsUniqueCountType, error) {
	model := &logsv0.ApisAlertDefinitionLogsUniqueCountType{}
	if modelMap["logs_filter"] != nil && len(modelMap["logs_filter"].([]interface{})) > 0 {
		LogsFilterModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(modelMap["logs_filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsFilter = LogsFilterModel
	}
	rules := []logsv0.ApisAlertDefinitionLogsUniqueCountRule{}
	for _, rulesItem := range modelMap["rules"].([]interface{}) {
		rulesItemModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountRule(rulesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		rules = append(rules, *rulesItemModel)
	}
	model.Rules = rules
	if modelMap["notification_payload_filter"] != nil {
		notificationPayloadFilter := []string{}
		for _, notificationPayloadFilterItem := range modelMap["notification_payload_filter"].([]interface{}) {
			notificationPayloadFilter = append(notificationPayloadFilter, notificationPayloadFilterItem.(string))
		}
		model.NotificationPayloadFilter = notificationPayloadFilter
	}
	if modelMap["max_unique_count_per_group_by_key"] != nil && modelMap["max_unique_count_per_group_by_key"].(string) != "" {
		model.MaxUniqueCountPerGroupByKey = core.StringPtr(modelMap["max_unique_count_per_group_by_key"].(string))
	}
	model.UniqueCountKeypath = core.StringPtr(modelMap["unique_count_keypath"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountRule(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsUniqueCountRule, error) {
	model := &logsv0.ApisAlertDefinitionLogsUniqueCountRule{}
	ConditionModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountCondition(modelMap["condition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Condition = ConditionModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountCondition(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsUniqueCountCondition, error) {
	model := &logsv0.ApisAlertDefinitionLogsUniqueCountCondition{}
	model.MaxUniqueCount = core.StringPtr(modelMap["max_unique_count"].(string))
	TimeWindowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueValueTimeWindow(modelMap["time_window"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.TimeWindow = TimeWindowModel
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueValueTimeWindow(modelMap map[string]interface{}) (*logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow, error) {
	model := &logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow{}
	model.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr(modelMap["logs_unique_value_time_window_specific_value"].(string))
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototype(modelMap map[string]interface{}) (logsv0.AlertDefinitionPrototypeIntf, error) {
	model := &logsv0.AlertDefinitionPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_immediate"] != nil && len(modelMap["logs_immediate"].([]interface{})) > 0 {
		LogsImmediateModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsImmediateType(modelMap["logs_immediate"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsImmediate = LogsImmediateModel
	}
	if modelMap["logs_threshold"] != nil && len(modelMap["logs_threshold"].([]interface{})) > 0 {
		LogsThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdType(modelMap["logs_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsThreshold = LogsThresholdModel
	}
	if modelMap["logs_ratio_threshold"] != nil && len(modelMap["logs_ratio_threshold"].([]interface{})) > 0 {
		LogsRatioThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioThresholdType(modelMap["logs_ratio_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsRatioThreshold = LogsRatioThresholdModel
	}
	if modelMap["logs_time_relative_threshold"] != nil && len(modelMap["logs_time_relative_threshold"].([]interface{})) > 0 {
		LogsTimeRelativeThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeThresholdType(modelMap["logs_time_relative_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsTimeRelativeThreshold = LogsTimeRelativeThresholdModel
	}
	if modelMap["metric_threshold"] != nil && len(modelMap["metric_threshold"].([]interface{})) > 0 {
		MetricThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdType(modelMap["metric_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MetricThreshold = MetricThresholdModel
	}
	if modelMap["flow"] != nil && len(modelMap["flow"].([]interface{})) > 0 {
		FlowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowType(modelMap["flow"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Flow = FlowModel
	}
	if modelMap["logs_anomaly"] != nil && len(modelMap["logs_anomaly"].([]interface{})) > 0 {
		LogsAnomalyModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyType(modelMap["logs_anomaly"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsAnomaly = LogsAnomalyModel
	}
	if modelMap["metric_anomaly"] != nil && len(modelMap["metric_anomaly"].([]interface{})) > 0 {
		MetricAnomalyModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyType(modelMap["metric_anomaly"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MetricAnomaly = MetricAnomalyModel
	}
	if modelMap["logs_new_value"] != nil && len(modelMap["logs_new_value"].([]interface{})) > 0 {
		LogsNewValueModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueType(modelMap["logs_new_value"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsNewValue = LogsNewValueModel
	}
	if modelMap["logs_unique_count"] != nil && len(modelMap["logs_unique_count"].([]interface{})) > 0 {
		LogsUniqueCountModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountType(modelMap["logs_unique_count"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsUniqueCount = LogsUniqueCountModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_immediate"] != nil && len(modelMap["logs_immediate"].([]interface{})) > 0 {
		LogsImmediateModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsImmediateType(modelMap["logs_immediate"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsImmediate = LogsImmediateModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_threshold"] != nil && len(modelMap["logs_threshold"].([]interface{})) > 0 {
		LogsThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdType(modelMap["logs_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsThreshold = LogsThresholdModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_ratio_threshold"] != nil && len(modelMap["logs_ratio_threshold"].([]interface{})) > 0 {
		LogsRatioThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioThresholdType(modelMap["logs_ratio_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsRatioThreshold = LogsRatioThresholdModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_time_relative_threshold"] != nil && len(modelMap["logs_time_relative_threshold"].([]interface{})) > 0 {
		LogsTimeRelativeThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeThresholdType(modelMap["logs_time_relative_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsTimeRelativeThreshold = LogsTimeRelativeThresholdModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["metric_threshold"] != nil && len(modelMap["metric_threshold"].([]interface{})) > 0 {
		MetricThresholdModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdType(modelMap["metric_threshold"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MetricThreshold = MetricThresholdModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["flow"] != nil && len(modelMap["flow"].([]interface{})) > 0 {
		FlowModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowType(modelMap["flow"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Flow = FlowModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_anomaly"] != nil && len(modelMap["logs_anomaly"].([]interface{})) > 0 {
		LogsAnomalyModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyType(modelMap["logs_anomaly"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsAnomaly = LogsAnomalyModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["metric_anomaly"] != nil && len(modelMap["metric_anomaly"].([]interface{})) > 0 {
		MetricAnomalyModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyType(modelMap["metric_anomaly"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.MetricAnomaly = MetricAnomalyModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_new_value"] != nil && len(modelMap["logs_new_value"].([]interface{})) > 0 {
		LogsNewValueModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueType(modelMap["logs_new_value"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsNewValue = LogsNewValueModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount(modelMap map[string]interface{}) (*logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount, error) {
	model := &logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["priority"] != nil && modelMap["priority"].(string) != "" {
		model.Priority = core.StringPtr(modelMap["priority"].(string))
	}
	if modelMap["active_on"] != nil && len(modelMap["active_on"].([]interface{})) > 0 {
		ActiveOnModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(modelMap["active_on"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActiveOn = ActiveOnModel
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	groupByKeys := []string{}
	for _, groupByKeysItem := range modelMap["group_by_keys"].([]interface{}) {
		groupByKeys = append(groupByKeys, groupByKeysItem.(string))
	}
	model.GroupByKeys = groupByKeys
	if modelMap["incidents_settings"] != nil && len(modelMap["incidents_settings"].([]interface{})) > 0 {
		IncidentsSettingsModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(modelMap["incidents_settings"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.IncidentsSettings = IncidentsSettingsModel
	}
	if modelMap["notification_group"] != nil && len(modelMap["notification_group"].([]interface{})) > 0 {
		NotificationGroupModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(modelMap["notification_group"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.NotificationGroup = NotificationGroupModel
	}
	if modelMap["entity_labels"] != nil {
		// TODO: handle EntityLabels, map with entry type 'string'
	}
	if modelMap["phantom_mode"] != nil {
		model.PhantomMode = core.BoolPtr(modelMap["phantom_mode"].(bool))
	}
	if modelMap["deleted"] != nil {
		model.Deleted = core.BoolPtr(modelMap["deleted"].(bool))
	}
	if modelMap["logs_unique_count"] != nil && len(modelMap["logs_unique_count"].([]interface{})) > 0 {
		LogsUniqueCountModel, err := ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountType(modelMap["logs_unique_count"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LogsUniqueCount = LogsUniqueCountModel
	}
	return model, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(model *logsv0.ApisAlertDefinitionActivitySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	startTimeMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model *logsv0.ApisAlertDefinitionTimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hours != nil {
		modelMap["hours"] = flex.IntValue(model.Hours)
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefIncidentSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(model *logsv0.ApisAlertDefinitionAlertDefNotificationGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["group_by_keys"] = model.GroupByKeys
	webhooks := []map[string]interface{}{}
	for _, webhooksItem := range model.Webhooks {
		webhooksItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(&webhooksItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		webhooks = append(webhooks, webhooksItemMap)
	}
	modelMap["webhooks"] = webhooks
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefWebhooksSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	integrationMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model.Integration)
	if err != nil {
		return modelMap, err
	}
	modelMap["integration"] = []map[string]interface{}{integrationMap}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model logsv0.ApisAlertDefinitionIntegrationTypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID); ok {
		return ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionIntegrationType); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.ApisAlertDefinitionIntegrationType)
		if model.IntegrationID != nil {
			modelMap["integration_id"] = flex.IntValue(model.IntegrationID)
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.ApisAlertDefinitionIntegrationTypeIntf subtype encountered")
	}
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model *logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IntegrationID != nil {
		modelMap["integration_id"] = flex.IntValue(model.IntegrationID)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(model *logsv0.ApisAlertDefinitionLogsImmediateType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model *logsv0.ApisAlertDefinitionLogsFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SimpleFilter != nil {
		simpleFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model.SimpleFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["simple_filter"] = []map[string]interface{}{simpleFilterMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model *logsv0.ApisAlertDefinitionLogsSimpleFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LuceneQuery != nil {
		modelMap["lucene_query"] = *model.LuceneQuery
	}
	if model.LabelFilters != nil {
		labelFiltersMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model.LabelFilters)
		if err != nil {
			return modelMap, err
		}
		modelMap["label_filters"] = []map[string]interface{}{labelFiltersMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model *logsv0.ApisAlertDefinitionLabelFilters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	applicationName := []map[string]interface{}{}
	for _, applicationNameItem := range model.ApplicationName {
		applicationNameItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(&applicationNameItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		applicationName = append(applicationName, applicationNameItemMap)
	}
	modelMap["application_name"] = applicationName
	subsystemName := []map[string]interface{}{}
	for _, subsystemNameItem := range model.SubsystemName {
		subsystemNameItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(&subsystemNameItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		subsystemName = append(subsystemName, subsystemNameItemMap)
	}
	modelMap["subsystem_name"] = subsystemName
	modelMap["severities"] = model.Severities
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(model *logsv0.ApisAlertDefinitionLabelFilterType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	modelMap["operation"] = *model.Operation
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model *logsv0.ApisAlertDefinitionUndetectedValuesManagement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["trigger_undetected_values"] = *model.TriggerUndetectedValues
	modelMap["auto_retire_timeframe"] = *model.AutoRetireTimeframe
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(model *logsv0.ApisAlertDefinitionLogsThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model *logsv0.ApisAlertDefinitionLogsThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_time_window_specific_value"] = *model.LogsTimeWindowSpecificValue
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model *logsv0.ApisAlertDefinitionAlertDefOverride) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["priority"] = *model.Priority
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsRatioThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	numeratorMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.Numerator)
	if err != nil {
		return modelMap, err
	}
	modelMap["numerator"] = []map[string]interface{}{numeratorMap}
	if model.NumeratorAlias != nil {
		modelMap["numerator_alias"] = *model.NumeratorAlias
	}
	denominatorMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.Denominator)
	if err != nil {
		return modelMap, err
	}
	modelMap["denominator"] = []map[string]interface{}{denominatorMap}
	if model.DenominatorAlias != nil {
		modelMap["denominator_alias"] = *model.DenominatorAlias
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	modelMap["group_by_for"] = *model.GroupByFor
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	if model.IgnoreInfinity != nil {
		modelMap["ignore_infinity"] = *model.IgnoreInfinity
	}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(model *logsv0.ApisAlertDefinitionLogsRatioRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model *logsv0.ApisAlertDefinitionLogsRatioCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsRatioTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_ratio_time_window_specific_value"] = *model.LogsRatioTimeWindowSpecificValue
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.IgnoreInfinity != nil {
		modelMap["ignore_infinity"] = *model.IgnoreInfinity
	}
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["compared_to"] = *model.ComparedTo
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(model *logsv0.ApisAlertDefinitionMetricThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	missingValuesMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model.MissingValues)
	if err != nil {
		return modelMap, err
	}
	modelMap["missing_values"] = []map[string]interface{}{missingValuesMap}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model *logsv0.ApisAlertDefinitionMetricFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["promql"] = *model.Promql
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(model *logsv0.ApisAlertDefinitionMetricThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model *logsv0.ApisAlertDefinitionMetricThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	ofTheLastMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model logsv0.ApisAlertDefinitionMetricTimeWindowIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue); ok {
		return ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration); ok {
		return ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindow); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.ApisAlertDefinitionMetricTimeWindow)
		if model.MetricTimeWindowSpecificValue != nil {
			modelMap["metric_time_window_specific_value"] = *model.MetricTimeWindowSpecificValue
		}
		if model.MetricTimeWindowDynamicDuration != nil {
			modelMap["metric_time_window_dynamic_duration"] = *model.MetricTimeWindowDynamicDuration
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.ApisAlertDefinitionMetricTimeWindowIntf subtype encountered")
	}
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowSpecificValue != nil {
		modelMap["metric_time_window_specific_value"] = *model.MetricTimeWindowSpecificValue
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowDynamicDuration != nil {
		modelMap["metric_time_window_dynamic_duration"] = *model.MetricTimeWindowDynamicDuration
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model logsv0.ApisAlertDefinitionMetricMissingValuesIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero); ok {
		return ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct); ok {
		return ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValues); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.ApisAlertDefinitionMetricMissingValues)
		if model.ReplaceWithZero != nil {
			modelMap["replace_with_zero"] = *model.ReplaceWithZero
		}
		if model.MinNonNullValuesPct != nil {
			modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.ApisAlertDefinitionMetricMissingValuesIntf subtype encountered")
	}
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplaceWithZero != nil {
		modelMap["replace_with_zero"] = *model.ReplaceWithZero
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MinNonNullValuesPct != nil {
		modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(model *logsv0.ApisAlertDefinitionFlowType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	stages := []map[string]interface{}{}
	for _, stagesItem := range model.Stages {
		stagesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(&stagesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		stages = append(stages, stagesItemMap)
	}
	modelMap["stages"] = stages
	if model.EnforceSuppression != nil {
		modelMap["enforce_suppression"] = *model.EnforceSuppression
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(model *logsv0.ApisAlertDefinitionFlowStages) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["timeframe_ms"] = *model.TimeframeMs
	modelMap["timeframe_type"] = *model.TimeframeType
	flowStagesGroupsMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model.FlowStagesGroups)
	if err != nil {
		return modelMap, err
	}
	modelMap["flow_stages_groups"] = []map[string]interface{}{flowStagesGroupsMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroups) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	groups := []map[string]interface{}{}
	for _, groupsItem := range model.Groups {
		groupsItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(&groupsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		groups = append(groups, groupsItemMap)
	}
	modelMap["groups"] = groups
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	alertDefs := []map[string]interface{}{}
	for _, alertDefsItem := range model.AlertDefs {
		alertDefsItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(&alertDefsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		alertDefs = append(alertDefs, alertDefsItemMap)
	}
	modelMap["alert_defs"] = alertDefs
	modelMap["next_op"] = *model.NextOp
	modelMap["alerts_op"] = *model.AlertsOp
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.Not != nil {
		modelMap["not"] = *model.Not
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	if model.AnomalyAlertSettings != nil {
		anomalyAlertSettingsMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["minimum_threshold"] = *model.MinimumThreshold
	timeWindowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model *logsv0.ApisAlertDefinitionAnomalyAlertSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PercentageOfDeviation != nil {
		modelMap["percentage_of_deviation"] = flex.Float64Value(model.PercentageOfDeviation)
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	if model.AnomalyAlertSettings != nil {
		anomalyAlertSettingsMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	if model.ForOverPct != nil {
		modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	}
	ofTheLastMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(model *logsv0.ApisAlertDefinitionLogsNewValueType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(model *logsv0.ApisAlertDefinitionLogsNewValueRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model *logsv0.ApisAlertDefinitionLogsNewValueCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["keypath_to_track"] = *model.KeypathToTrack
	timeWindowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsNewValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_new_value_time_window_specific_value"] = *model.LogsNewValueTimeWindowSpecificValue
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	if model.NotificationPayloadFilter != nil {
		modelMap["notification_payload_filter"] = model.NotificationPayloadFilter
	}
	if model.MaxUniqueCountPerGroupByKey != nil {
		modelMap["max_unique_count_per_group_by_key"] = *model.MaxUniqueCountPerGroupByKey
	}
	modelMap["unique_count_keypath"] = *model.UniqueCountKeypath
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["max_unique_count"] = *model.MaxUniqueCount
	timeWindowMap, err := ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_unique_value_time_window_specific_value"] = *model.LogsUniqueValueTimeWindowSpecificValue
	return modelMap, nil
}
