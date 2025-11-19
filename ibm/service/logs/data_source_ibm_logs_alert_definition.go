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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsAlertDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsAlertDefinitionRead,

		Schema: map[string]*schema.Schema{
			"logs_alert_definition_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Alert definition ID.",
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
				Description: "The previous or old alert ID.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the alert definition.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A detailed description of what the alert monitors and when it triggers.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the alert is currently active and monitoring. If true, alert is active.",
			},
			"priority": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The priority of the alert definition.",
			},
			"active_on": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Defining when the alert is active.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day_of_week": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Days of the week when the alert is active.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Start time of the alert activity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hours": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The hour of the day in 24-hour format. Must be an integer between 0 and 23.",
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Minute of the hour of the day. Must be an integer between 0 and 59.",
									},
								},
							},
						},
						"end_time": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Start time of the alert activity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hours": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The hour of the day in 24-hour format. Must be an integer between 0 and 23.",
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Minute of the hour of the day. Must be an integer between 0 and 59.",
									},
								},
							},
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alert type.",
			},
			"group_by_keys": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Keys used to group and aggregate alert data.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"incidents_settings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Incident creation and management settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notify_on": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicate if the alert should be triggered or triggered and resolved.",
						},
						"minutes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The time in minutes before the alert can be triggered again.",
						},
					},
				},
			},
			"notification_group": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Primary notification group for alert events.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_by_keys": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Group the alerts by these keys.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"webhooks": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The settings for webhooks associated with the alert definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notify_on": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicate if the alert should be triggered or triggered and resolved.",
									},
									"integration": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The integration type for webhook notifications.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"integration_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The integration ID for the notification.",
												},
											},
										},
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
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
				Computed:    true,
				Description: "Labels used to identify and categorize the alert entity.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"phantom_mode": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the alert is in phantom mode (creating incidents or not).",
			},
			"deleted": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the alert has been marked as deleted.",
			},
			"logs_immediate": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for immediate log-based alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"logs_threshold": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for the log-based threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Should trigger the alert when undetected values are detected. If true, alert is triggered.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the threshold alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the threshold alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time window defined for an alert to be triggered.",
															},
														},
													},
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "The condition type for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"logs_ratio_threshold": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for the log-based ratio threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"numerator": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"numerator_alias": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The alias for the numerator filter, used for display purposes.",
						},
						"denominator": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"denominator_alias": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The alias for the denominator filter, used for display purposes.",
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the ratio alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the ratio alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_ratio_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
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
										Computed:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "The condition type for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"group_by_for": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The group by settings for the numerator and denominator filters.",
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Should trigger the alert when undetected values are detected. If true, alert is triggered.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"ignore_infinity": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determine whether to ignore an infinity result or not. If true, alert is not triggered. When the value of second query is 0, the result of the ratio will be infinity.",
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"logs_time_relative_threshold": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for time-relative log threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the time-relative alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the time-relative alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"compared_to": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The time frame to compare the current value against.",
												},
											},
										},
									},
									"override": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"ignore_infinity": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Ignore infinity values in the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Should trigger the alert when undetected values are detected. If true, alert is triggered.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"metric_threshold": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for metric-based threshold alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match metric entries for the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"promql": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The filter is a PromQL expression.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the metric threshold alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the metric threshold alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"for_over_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The percentage of values that must exceed the threshold to trigger the alert.",
												},
												"of_the_last": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time window as a specific value.",
															},
															"metric_time_window_dynamic_duration": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
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
										Computed:    true,
										Description: "The override settings for the alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "The type of the alert condition.",
						},
						"undetected_values_management": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration for handling the undetected values in the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trigger_undetected_values": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Should trigger the alert when undetected values are detected. If true, alert is triggered.",
									},
									"auto_retire_timeframe": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timeframe for auto-retiring the alert when undetected values are detected.",
									},
								},
							},
						},
						"missing_values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration for handling missing values in the alert. Only one of `replace_with_zero` or `min_non_null_value_pct` is supported.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"replace_with_zero": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If set to true, missing values will be replaced with zero.",
									},
									"min_non_null_values_pct": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "If set, specifies the minimum percentage of non-null values required for the alert to be triggered.",
									},
								},
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
					},
				},
			},
			"flow": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for flow alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stages": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The definition of stages of the flow alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timeframe_ms": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timeframe for the flow alert in milliseconds.",
									},
									"timeframe_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of timeframe for the flow alert.",
									},
									"flow_stages_groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The definition of groups in the flow alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"groups": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The definition of an array of groups with alerts and logical operation among those alerts in the flow alert.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alert_defs": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The alert definitions for the flow stage group.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The alert definition ID.",
																		},
																		"not": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Whether or not to negate the alert definition. If true, flow checks for the negate condition of the respective alert.",
																		},
																	},
																},
															},
															"next_op": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The logical operation to apply to the next stage.",
															},
															"alerts_op": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
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
							Computed:    true,
							Description: "Whether to enforce suppression for the flow alert.",
						},
					},
				},
			},
			"logs_anomaly": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for the log-based anomaly detection alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the log anomaly alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the anomaly alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"minimum_threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time window defined for an alert to be triggered.",
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
							Computed:    true,
							Description: "The condition type for the alert.",
						},
						"notification_payload_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The notification payload filter to specify which fields are included in the notification.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"anomaly_alert_settings": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The anomaly alert settings configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"percentage_of_deviation": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The percentage of deviation from the baseline when the alert is triggered.",
									},
								},
							},
						},
					},
				},
			},
			"metric_anomaly": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for metric-based anomaly detection alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match metric entries for the alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"promql": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The filter is a PromQL expression.",
									},
								},
							},
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the metric anomaly alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for the metric anomaly alert.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The threshold value for the alert condition.",
												},
												"for_over_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The percentage of the metric values that must exceed the threshold to trigger the alert.",
												},
												"of_the_last": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the alert condition.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time window as a specific value.",
															},
															"metric_time_window_dynamic_duration": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time window as a dynamic value.",
															},
														},
													},
												},
												"min_non_null_values_pct": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
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
							Computed:    true,
							Description: "The condition type for the alert.",
						},
						"evaluation_delay_ms": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The delay in milliseconds before evaluating the alert condition.",
						},
						"anomaly_alert_settings": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The anomaly alert settings configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"percentage_of_deviation": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The percentage of deviation from the baseline when the alert is triggered.",
									},
								},
							},
						},
					},
				},
			},
			"logs_new_value": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for alerts triggered by new log values.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The condition rules for the log new value alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for detecting new values in logs.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"keypath_to_track": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The keypath to track for new values.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for detecting new values.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_new_value_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
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
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"logs_unique_count": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration for alerts based on unique log value counts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logs_filter": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter to match log entries for immediate alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"simple_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A simple filter that uses a Lucene query and label filters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lucene_query": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Lucene query to filter logs.",
												},
												"label_filters": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The label filters to filter logs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"application_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by application names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"subsystem_name": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by subsystem names.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value used to filter the label.",
																		},
																		"operation": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The operation to perform on the label value.",
																		},
																	},
																},
															},
															"severities": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Filter by log severities.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Rules defining the conditions for the unique count alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition for detecting unique counts in logs.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_unique_count": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The maximum unique count for the alert condition.",
												},
												"time_window": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The time window for the unique count alert.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"logs_unique_value_time_window_specific_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
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
							Computed:    true,
							Description: "The filter to specify which fields are included in the notification payload.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"max_unique_count_per_group_by_key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The maximum unique count per group by key.",
						},
						"unique_count_keypath": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The keypath in the logs to be used for unique count.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsAlertDefinitionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	getAlertDefOptions := &logsv0.GetAlertDefOptions{}

	getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(d.Get("logs_alert_definition_id").(string))))

	alertDefinitionIntf, _, err := logsClient.GetAlertDefWithContext(context, getAlertDefOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAlertDefWithContext failed: %s", err.Error()), "(Data) ibm_logs_alert_definition", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)

	d.SetId(fmt.Sprintf("%s", *getAlertDefOptions.ID))

	if err = d.Set("instance_id", instanceId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if err = d.Set("region", region); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
	}

	if !core.IsNil(alertDefinition.CreatedTime) {
		if err = d.Set("created_time", flex.DateTimeToString(alertDefinition.CreatedTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_time: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-created_time").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.UpdatedTime) {
		if err = d.Set("updated_time", flex.DateTimeToString(alertDefinition.UpdatedTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting updated_time: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-updated_time").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.AlertVersionID) {
		if err = d.Set("alert_version_id", flex.Stringify(alertDefinition.AlertVersionID)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alert_version_id: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-alert_version_id").GetDiag()
		}
	}

	if err = d.Set("name", alertDefinition.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-name").GetDiag()
	}

	if !core.IsNil(alertDefinition.Description) {
		if err = d.Set("description", alertDefinition.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.Enabled) {
		if err = d.Set("enabled", alertDefinition.Enabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enabled: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-enabled").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.Priority) {
		if err = d.Set("priority", alertDefinition.Priority); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting priority: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-priority").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.ActiveOn) {
		activeOn := []map[string]interface{}{}
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(alertDefinition.ActiveOn)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "active_on-to-map").GetDiag()
		}
		activeOn = append(activeOn, activeOnMap)
		if err = d.Set("active_on", activeOn); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting active_on: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-active_on").GetDiag()
		}
	}

	if err = d.Set("type", alertDefinition.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-type").GetDiag()
	}

	if !core.IsNil(alertDefinition.GroupByKeys) {
		groupByKeys := []interface{}{}
		for _, groupByKeysItem := range alertDefinition.GroupByKeys {
			groupByKeys = append(groupByKeys, groupByKeysItem)
		}
		if err = d.Set("group_by_keys", groupByKeys); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting group_by_keys: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-group_by_keys").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.IncidentsSettings) {
		incidentsSettings := []map[string]interface{}{}
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(alertDefinition.IncidentsSettings)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "incidents_settings-to-map").GetDiag()
		}
		incidentsSettings = append(incidentsSettings, incidentsSettingsMap)
		if err = d.Set("incidents_settings", incidentsSettings); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting incidents_settings: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-incidents_settings").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.NotificationGroup) {
		notificationGroup := []map[string]interface{}{}
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(alertDefinition.NotificationGroup)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "notification_group-to-map").GetDiag()
		}
		notificationGroup = append(notificationGroup, notificationGroupMap)
		if err = d.Set("notification_group", notificationGroup); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting notification_group: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-notification_group").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.EntityLabels) {
		if err = d.Set("entity_labels", alertDefinition.EntityLabels); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_labels: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-entity_labels").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.PhantomMode) {
		if err = d.Set("phantom_mode", alertDefinition.PhantomMode); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting phantom_mode: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-phantom_mode").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.Deleted) {
		if err = d.Set("deleted", alertDefinition.Deleted); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deleted: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-deleted").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsImmediate) {
		logsImmediate := []map[string]interface{}{}
		logsImmediateMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(alertDefinition.LogsImmediate)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_immediate-to-map").GetDiag()
		}
		logsImmediate = append(logsImmediate, logsImmediateMap)
		if err = d.Set("logs_immediate", logsImmediate); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_immediate: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_immediate").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsThreshold) {
		logsThreshold := []map[string]interface{}{}
		logsThresholdMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(alertDefinition.LogsThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_threshold-to-map").GetDiag()
		}
		logsThreshold = append(logsThreshold, logsThresholdMap)
		if err = d.Set("logs_threshold", logsThreshold); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_threshold: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_threshold").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsRatioThreshold) {
		logsRatioThreshold := []map[string]interface{}{}
		logsRatioThresholdMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(alertDefinition.LogsRatioThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_ratio_threshold-to-map").GetDiag()
		}
		logsRatioThreshold = append(logsRatioThreshold, logsRatioThresholdMap)
		if err = d.Set("logs_ratio_threshold", logsRatioThreshold); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_ratio_threshold: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_ratio_threshold").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsTimeRelativeThreshold) {
		logsTimeRelativeThreshold := []map[string]interface{}{}
		logsTimeRelativeThresholdMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(alertDefinition.LogsTimeRelativeThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_time_relative_threshold-to-map").GetDiag()
		}
		logsTimeRelativeThreshold = append(logsTimeRelativeThreshold, logsTimeRelativeThresholdMap)
		if err = d.Set("logs_time_relative_threshold", logsTimeRelativeThreshold); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_time_relative_threshold: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_time_relative_threshold").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.MetricThreshold) {
		metricThreshold := []map[string]interface{}{}
		metricThresholdMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(alertDefinition.MetricThreshold)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "metric_threshold-to-map").GetDiag()
		}
		metricThreshold = append(metricThreshold, metricThresholdMap)
		if err = d.Set("metric_threshold", metricThreshold); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metric_threshold: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-metric_threshold").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.Flow) {
		flow := []map[string]interface{}{}
		flowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(alertDefinition.Flow)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "flow-to-map").GetDiag()
		}
		flow = append(flow, flowMap)
		if err = d.Set("flow", flow); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting flow: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-flow").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsAnomaly) {
		logsAnomaly := []map[string]interface{}{}
		logsAnomalyMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(alertDefinition.LogsAnomaly)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_anomaly-to-map").GetDiag()
		}
		logsAnomaly = append(logsAnomaly, logsAnomalyMap)
		if err = d.Set("logs_anomaly", logsAnomaly); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_anomaly: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_anomaly").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.MetricAnomaly) {
		metricAnomaly := []map[string]interface{}{}
		metricAnomalyMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(alertDefinition.MetricAnomaly)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "metric_anomaly-to-map").GetDiag()
		}
		metricAnomaly = append(metricAnomaly, metricAnomalyMap)
		if err = d.Set("metric_anomaly", metricAnomaly); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metric_anomaly: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-metric_anomaly").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsNewValue) {
		logsNewValue := []map[string]interface{}{}
		logsNewValueMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(alertDefinition.LogsNewValue)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_new_value-to-map").GetDiag()
		}
		logsNewValue = append(logsNewValue, logsNewValueMap)
		if err = d.Set("logs_new_value", logsNewValue); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_new_value: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_new_value").GetDiag()
		}
	}

	if !core.IsNil(alertDefinition.LogsUniqueCount) {
		logsUniqueCount := []map[string]interface{}{}
		logsUniqueCountMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(alertDefinition.LogsUniqueCount)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definition", "read", "logs_unique_count-to-map").GetDiag()
		}
		logsUniqueCount = append(logsUniqueCount, logsUniqueCountMap)
		if err = d.Set("logs_unique_count", logsUniqueCount); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting logs_unique_count: %s", err), "(Data) ibm_logs_alert_definition", "read", "set-logs_unique_count").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(model *logsv0.ApisAlertDefinitionActivitySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	startTimeMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model *logsv0.ApisAlertDefinitionTimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hours != nil {
		modelMap["hours"] = flex.IntValue(model.Hours)
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefIncidentSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(model *logsv0.ApisAlertDefinitionAlertDefNotificationGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GroupByKeys != nil {
		modelMap["group_by_keys"] = model.GroupByKeys
	}
	if model.Webhooks != nil {
		webhooks := []map[string]interface{}{}
		for _, webhooksItem := range model.Webhooks {
			webhooksItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(&webhooksItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			webhooks = append(webhooks, webhooksItemMap)
		}
		modelMap["webhooks"] = webhooks
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefWebhooksSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	integrationMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model.Integration)
	if err != nil {
		return modelMap, err
	}
	modelMap["integration"] = []map[string]interface{}{integrationMap}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model logsv0.ApisAlertDefinitionIntegrationTypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID); ok {
		return DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID))
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model *logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IntegrationID != nil {
		modelMap["integration_id"] = flex.IntValue(model.IntegrationID)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(model *logsv0.ApisAlertDefinitionLogsImmediateType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model *logsv0.ApisAlertDefinitionLogsFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SimpleFilter != nil {
		simpleFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model.SimpleFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["simple_filter"] = []map[string]interface{}{simpleFilterMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model *logsv0.ApisAlertDefinitionLogsSimpleFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LuceneQuery != nil {
		modelMap["lucene_query"] = *model.LuceneQuery
	}
	if model.LabelFilters != nil {
		labelFiltersMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model.LabelFilters)
		if err != nil {
			return modelMap, err
		}
		modelMap["label_filters"] = []map[string]interface{}{labelFiltersMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model *logsv0.ApisAlertDefinitionLabelFilters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ApplicationName != nil {
		applicationName := []map[string]interface{}{}
		for _, applicationNameItem := range model.ApplicationName {
			applicationNameItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(&applicationNameItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			applicationName = append(applicationName, applicationNameItemMap)
		}
		modelMap["application_name"] = applicationName
	}
	if model.SubsystemName != nil {
		subsystemName := []map[string]interface{}{}
		for _, subsystemNameItem := range model.SubsystemName {
			subsystemNameItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(&subsystemNameItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			subsystemName = append(subsystemName, subsystemNameItemMap)
		}
		modelMap["subsystem_name"] = subsystemName
	}
	if model.Severities != nil {
		modelMap["severities"] = model.Severities
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(model *logsv0.ApisAlertDefinitionLabelFilterType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	modelMap["operation"] = *model.Operation
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model *logsv0.ApisAlertDefinitionUndetectedValuesManagement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["trigger_undetected_values"] = *model.TriggerUndetectedValues
	modelMap["auto_retire_timeframe"] = *model.AutoRetireTimeframe
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(model *logsv0.ApisAlertDefinitionLogsThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model *logsv0.ApisAlertDefinitionLogsThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_time_window_specific_value"] = *model.LogsTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model *logsv0.ApisAlertDefinitionAlertDefOverride) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["priority"] = *model.Priority
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsRatioThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	numeratorMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.Numerator)
	if err != nil {
		return modelMap, err
	}
	modelMap["numerator"] = []map[string]interface{}{numeratorMap}
	if model.NumeratorAlias != nil {
		modelMap["numerator_alias"] = *model.NumeratorAlias
	}
	denominatorMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.Denominator)
	if err != nil {
		return modelMap, err
	}
	modelMap["denominator"] = []map[string]interface{}{denominatorMap}
	if model.DenominatorAlias != nil {
		modelMap["denominator_alias"] = *model.DenominatorAlias
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(&rulesItem) // #nosec G601
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
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(model *logsv0.ApisAlertDefinitionLogsRatioRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model *logsv0.ApisAlertDefinitionLogsRatioCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsRatioTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_ratio_time_window_specific_value"] = *model.LogsRatioTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(&rulesItem) // #nosec G601
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
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["compared_to"] = *model.ComparedTo
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(model *logsv0.ApisAlertDefinitionMetricThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	missingValuesMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model.MissingValues)
	if err != nil {
		return modelMap, err
	}
	modelMap["missing_values"] = []map[string]interface{}{missingValuesMap}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model *logsv0.ApisAlertDefinitionMetricFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["promql"] = *model.Promql
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(model *logsv0.ApisAlertDefinitionMetricThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model *logsv0.ApisAlertDefinitionMetricThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	ofTheLastMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model logsv0.ApisAlertDefinitionMetricTimeWindowIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue); ok {
		return DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration); ok {
		return DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration))
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowSpecificValue != nil {
		modelMap["metric_time_window_specific_value"] = *model.MetricTimeWindowSpecificValue
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowDynamicDuration != nil {
		modelMap["metric_time_window_dynamic_duration"] = *model.MetricTimeWindowDynamicDuration
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model logsv0.ApisAlertDefinitionMetricMissingValuesIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero); ok {
		return DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct); ok {
		return DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct))
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplaceWithZero != nil {
		modelMap["replace_with_zero"] = *model.ReplaceWithZero
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MinNonNullValuesPct != nil {
		modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(model *logsv0.ApisAlertDefinitionFlowType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	stages := []map[string]interface{}{}
	for _, stagesItem := range model.Stages {
		stagesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(&stagesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(model *logsv0.ApisAlertDefinitionFlowStages) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["timeframe_ms"] = *model.TimeframeMs
	modelMap["timeframe_type"] = *model.TimeframeType
	flowStagesGroupsMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model.FlowStagesGroups)
	if err != nil {
		return modelMap, err
	}
	modelMap["flow_stages_groups"] = []map[string]interface{}{flowStagesGroupsMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroups) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	groups := []map[string]interface{}{}
	for _, groupsItem := range model.Groups {
		groupsItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(&groupsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		groups = append(groups, groupsItemMap)
	}
	modelMap["groups"] = groups
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	alertDefs := []map[string]interface{}{}
	for _, alertDefsItem := range model.AlertDefs {
		alertDefsItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(&alertDefsItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.Not != nil {
		modelMap["not"] = *model.Not
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(&rulesItem) // #nosec G601
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
		anomalyAlertSettingsMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["minimum_threshold"] = *model.MinimumThreshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model *logsv0.ApisAlertDefinitionAnomalyAlertSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PercentageOfDeviation != nil {
		modelMap["percentage_of_deviation"] = flex.Float64Value(model.PercentageOfDeviation)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(&rulesItem) // #nosec G601
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
		anomalyAlertSettingsMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	if model.ForOverPct != nil {
		modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	}
	ofTheLastMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(model *logsv0.ApisAlertDefinitionLogsNewValueType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(model *logsv0.ApisAlertDefinitionLogsNewValueRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model *logsv0.ApisAlertDefinitionLogsNewValueCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["keypath_to_track"] = *model.KeypathToTrack
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsNewValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_new_value_time_window_specific_value"] = *model.LogsNewValueTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["max_unique_count"] = *model.MaxUniqueCount
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_unique_value_time_window_specific_value"] = *model.LogsUniqueValueTimeWindowSpecificValue
	return modelMap, nil
}
