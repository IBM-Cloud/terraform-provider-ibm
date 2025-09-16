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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsAlertDefinitions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsAlertDefinitionsRead,

		Schema: map[string]*schema.Schema{
			"alert_definitions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of alert definitions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is the alert definition's persistent ID (does not change on replace), AKA UniqueIdentifier.",
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
							Description: "Whether the alert is currently active and monitoring.",
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
													Description: "Hours of day in 24 hour format. Should be from 0 to 23.",
												},
												"minutes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minutes of hour of day. Must be from 0 to 59.",
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
													Description: "Hours of day in 24 hour format. Should be from 0 to 23.",
												},
												"minutes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Minutes of hour of day. Must be from 0 to 59.",
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
										Description: "The condition to notify about the alert.",
									},
									"minutes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The time in minutes before the alert can be retriggered.",
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
										Description: "The keys to group the alerts by.",
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
													Description: "The condition to notify about the alert.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The filter to specify which fields to include in the notification payload.",
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
							Description: "Configuration for log-based threshold alerts.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
													Description: "Should trigger the alert when undetected values are detected.",
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
										Description: "The rules for the threshold alert.",
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
										Description: "The type of condition for the alert.",
									},
									"notification_payload_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The filter to specify which fields to include in the notification payload.",
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
							Description: "Configuration for log-based ratio threshold alerts.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The rules for the ratio alert.",
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
										Description: "The type of condition for the alert.",
									},
									"notification_payload_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The filter to specify which fields to include in the notification payload.",
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
													Description: "Should trigger the alert when undetected values are detected.",
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
										Description: "The configuration for ignoring infinity values in the ratio.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The rules for the time-relative alert.",
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
										Description: "The filter to specify which fields to include in the notification payload.",
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
													Description: "Should trigger the alert when undetected values are detected.",
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
										Description: "The rules for the metric threshold alert.",
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
													Description: "Should trigger the alert when undetected values are detected.",
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
										Description: "Configuration for handling missing values in the alert.",
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
							Description: "Configuration for flow-based alerts.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stages": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The stages of the flow alert.",
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
													Description: "Flow stages groups.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"groups": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The groups of stages in the flow alert.",
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
																						Description: "Whether to negate the alert definition or not.",
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
							Description: "Configuration for log-based anomaly detection alerts.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The rules for the log anomaly alert.",
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
										Computed:    true,
										Description: "The type of condition for the alert.",
									},
									"notification_payload_filter": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The notification payload filter to specify which fields to include in the notification.",
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
										Description: "The rules for the metric anomaly alert.",
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
										Description: "The type of condition for the alert.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The rules for the log new value alert.",
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
										Description: "The filter to specify which fields to include in the notification payload.",
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
																						Description: "/ The value of the label to filter by.",
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
																						Description: "/ The value of the label to filter by.",
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
										Description: "The rules for the log unique count alert.",
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
										Description: "The filter to specify which fields to include in the notification payload.",
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
				},
			},
		},
	}
}

func dataSourceIbmLogsAlertDefinitionsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definitions", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	listAlertDefsOptions := &logsv0.ListAlertDefsOptions{}

	alertDefinitionCollection, _, err := logsClient.ListAlertDefsWithContext(context, listAlertDefsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListAlertDefsWithContext failed: %s", err.Error()), "(Data) ibm_logs_alert_definitions", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmLogsAlertDefinitionsID(d))
	if err = d.Set("instance_id", instanceId); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if err = d.Set("region", region); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
	}

	alertDefinitions := []map[string]interface{}{}
	for _, alertDefinitionsItem := range alertDefinitionCollection.AlertDefinitions {
		alertDefinitionsItemMap, err := DataSourceIbmLogsAlertDefinitionsAlertDefinitionToMap(alertDefinitionsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_alert_definitions", "read", "alert_definitions-to-map").GetDiag()
		}
		alertDefinitions = append(alertDefinitions, alertDefinitionsItemMap)
	}
	if err = d.Set("alert_definitions", alertDefinitions); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alert_definitions: %s", err), "(Data) ibm_logs_alert_definitions", "read", "set-alert_definitions").GetDiag()
	}

	return nil
}

// dataSourceIbmLogsAlertDefinitionsID returns a reasonable ID for the list.
func dataSourceIbmLogsAlertDefinitionsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionToMap(model logsv0.AlertDefinitionIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediateToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThresholdToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThresholdToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThresholdToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThresholdToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlowToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomalyToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomalyToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValueToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue))
	} else if _, ok := model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount); ok {
		return DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCountToMap(model.(*logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount))
	} else if _, ok := model.(*logsv0.AlertDefinition); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.AlertDefinition)
		modelMap["id"] = model.ID.String()
		if model.CreatedTime != nil {
			modelMap["created_time"] = model.CreatedTime.String()
		}
		if model.UpdatedTime != nil {
			modelMap["updated_time"] = model.UpdatedTime.String()
		}
		if model.AlertVersionID != nil {
			modelMap["alert_version_id"] = model.AlertVersionID.String()
		}
		modelMap["name"] = *model.Name
		if model.Description != nil {
			modelMap["description"] = *model.Description
		}
		if model.Enabled != nil {
			modelMap["enabled"] = *model.Enabled
		}
		if model.Priority != nil {
			modelMap["priority"] = *model.Priority
		}
		if model.ActiveOn != nil {
			activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
			if err != nil {
				return modelMap, err
			}
			modelMap["active_on"] = []map[string]interface{}{activeOnMap}
		}
		modelMap["type"] = *model.Type
		modelMap["group_by_keys"] = model.GroupByKeys
		if model.IncidentsSettings != nil {
			incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
			if err != nil {
				return modelMap, err
			}
			modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
		}
		if model.NotificationGroup != nil {
			notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
			if err != nil {
				return modelMap, err
			}
			modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
		}
		if model.EntityLabels != nil {
			entityLabels := make(map[string]interface{})
			for k, v := range model.EntityLabels {
				entityLabels[k] = flex.Stringify(v)
			}
			modelMap["entity_labels"] = entityLabels
		}
		if model.PhantomMode != nil {
			modelMap["phantom_mode"] = *model.PhantomMode
		}
		if model.Deleted != nil {
			modelMap["deleted"] = *model.Deleted
		}
		if model.LogsImmediate != nil {
			logsImmediateMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsImmediateTypeToMap(model.LogsImmediate)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_immediate"] = []map[string]interface{}{logsImmediateMap}
		}
		if model.LogsThreshold != nil {
			logsThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdTypeToMap(model.LogsThreshold)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_threshold"] = []map[string]interface{}{logsThresholdMap}
		}
		if model.LogsRatioThreshold != nil {
			logsRatioThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioThresholdTypeToMap(model.LogsRatioThreshold)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_ratio_threshold"] = []map[string]interface{}{logsRatioThresholdMap}
		}
		if model.LogsTimeRelativeThreshold != nil {
			logsTimeRelativeThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model.LogsTimeRelativeThreshold)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_time_relative_threshold"] = []map[string]interface{}{logsTimeRelativeThresholdMap}
		}
		if model.MetricThreshold != nil {
			metricThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdTypeToMap(model.MetricThreshold)
			if err != nil {
				return modelMap, err
			}
			modelMap["metric_threshold"] = []map[string]interface{}{metricThresholdMap}
		}
		if model.Flow != nil {
			flowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowTypeToMap(model.Flow)
			if err != nil {
				return modelMap, err
			}
			modelMap["flow"] = []map[string]interface{}{flowMap}
		}
		if model.LogsAnomaly != nil {
			logsAnomalyMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyTypeToMap(model.LogsAnomaly)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_anomaly"] = []map[string]interface{}{logsAnomalyMap}
		}
		if model.MetricAnomaly != nil {
			metricAnomalyMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyTypeToMap(model.MetricAnomaly)
			if err != nil {
				return modelMap, err
			}
			modelMap["metric_anomaly"] = []map[string]interface{}{metricAnomalyMap}
		}
		if model.LogsNewValue != nil {
			logsNewValueMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTypeToMap(model.LogsNewValue)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_new_value"] = []map[string]interface{}{logsNewValueMap}
		}
		if model.LogsUniqueCount != nil {
			logsUniqueCountMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountTypeToMap(model.LogsUniqueCount)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_unique_count"] = []map[string]interface{}{logsUniqueCountMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.AlertDefinitionIntf subtype encountered")
	}
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model *logsv0.ApisAlertDefinitionActivitySchedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["day_of_week"] = model.DayOfWeek
	startTimeMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionTimeOfDayToMap(model.StartTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	endTimeMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionTimeOfDayToMap(model.EndTime)
	if err != nil {
		return modelMap, err
	}
	modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionTimeOfDayToMap(model *logsv0.ApisAlertDefinitionTimeOfDay) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hours != nil {
		modelMap["hours"] = flex.IntValue(model.Hours)
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefIncidentSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model *logsv0.ApisAlertDefinitionAlertDefNotificationGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["group_by_keys"] = model.GroupByKeys
	webhooks := []map[string]interface{}{}
	for _, webhooksItem := range model.Webhooks {
		webhooksItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefWebhooksSettingsToMap(&webhooksItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		webhooks = append(webhooks, webhooksItemMap)
	}
	modelMap["webhooks"] = webhooks
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefWebhooksSettingsToMap(model *logsv0.ApisAlertDefinitionAlertDefWebhooksSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotifyOn != nil {
		modelMap["notify_on"] = *model.NotifyOn
	}
	integrationMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeToMap(model.Integration)
	if err != nil {
		return modelMap, err
	}
	modelMap["integration"] = []map[string]interface{}{integrationMap}
	if model.Minutes != nil {
		modelMap["minutes"] = flex.IntValue(model.Minutes)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeToMap(model logsv0.ApisAlertDefinitionIntegrationTypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID); ok {
		return DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model.(*logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID))
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model *logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IntegrationID != nil {
		modelMap["integration_id"] = flex.IntValue(model.IntegrationID)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsImmediateTypeToMap(model *logsv0.ApisAlertDefinitionLogsImmediateType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model *logsv0.ApisAlertDefinitionLogsFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SimpleFilter != nil {
		simpleFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsSimpleFilterToMap(model.SimpleFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["simple_filter"] = []map[string]interface{}{simpleFilterMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsSimpleFilterToMap(model *logsv0.ApisAlertDefinitionLogsSimpleFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LuceneQuery != nil {
		modelMap["lucene_query"] = *model.LuceneQuery
	}
	if model.LabelFilters != nil {
		labelFiltersMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFiltersToMap(model.LabelFilters)
		if err != nil {
			return modelMap, err
		}
		modelMap["label_filters"] = []map[string]interface{}{labelFiltersMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFiltersToMap(model *logsv0.ApisAlertDefinitionLabelFilters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	applicationName := []map[string]interface{}{}
	for _, applicationNameItem := range model.ApplicationName {
		applicationNameItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFilterTypeToMap(&applicationNameItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		applicationName = append(applicationName, applicationNameItemMap)
	}
	modelMap["application_name"] = applicationName
	subsystemName := []map[string]interface{}{}
	for _, subsystemNameItem := range model.SubsystemName {
		subsystemNameItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFilterTypeToMap(&subsystemNameItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		subsystemName = append(subsystemName, subsystemNameItemMap)
	}
	modelMap["subsystem_name"] = subsystemName
	modelMap["severities"] = model.Severities
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFilterTypeToMap(model *logsv0.ApisAlertDefinitionLabelFilterType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	modelMap["operation"] = *model.Operation
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model *logsv0.ApisAlertDefinitionUndetectedValuesManagement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["trigger_undetected_values"] = *model.TriggerUndetectedValues
	modelMap["auto_retire_timeframe"] = *model.AutoRetireTimeframe
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdRuleToMap(model *logsv0.ApisAlertDefinitionLogsThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdConditionToMap(model *logsv0.ApisAlertDefinitionLogsThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_time_window_specific_value"] = *model.LogsTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model *logsv0.ApisAlertDefinitionAlertDefOverride) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["priority"] = *model.Priority
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsRatioThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	numeratorMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.Numerator)
	if err != nil {
		return modelMap, err
	}
	modelMap["numerator"] = []map[string]interface{}{numeratorMap}
	if model.NumeratorAlias != nil {
		modelMap["numerator_alias"] = *model.NumeratorAlias
	}
	denominatorMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.Denominator)
	if err != nil {
		return modelMap, err
	}
	modelMap["denominator"] = []map[string]interface{}{denominatorMap}
	if model.DenominatorAlias != nil {
		modelMap["denominator_alias"] = *model.DenominatorAlias
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioRulesToMap(&rulesItem) // #nosec G601
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
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioRulesToMap(model *logsv0.ApisAlertDefinitionLogsRatioRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioConditionToMap(model *logsv0.ApisAlertDefinitionLogsRatioCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsRatioTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_ratio_time_window_specific_value"] = *model.LogsRatioTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeRuleToMap(&rulesItem) // #nosec G601
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
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeRuleToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeConditionToMap(model *logsv0.ApisAlertDefinitionLogsTimeRelativeCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["compared_to"] = *model.ComparedTo
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdTypeToMap(model *logsv0.ApisAlertDefinitionMetricThresholdType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["condition_type"] = *model.ConditionType
	if model.UndetectedValuesManagement != nil {
		undetectedValuesManagementMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model.UndetectedValuesManagement)
		if err != nil {
			return modelMap, err
		}
		modelMap["undetected_values_management"] = []map[string]interface{}{undetectedValuesManagementMap}
	}
	missingValuesMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesToMap(model.MissingValues)
	if err != nil {
		return modelMap, err
	}
	modelMap["missing_values"] = []map[string]interface{}{missingValuesMap}
	if model.EvaluationDelayMs != nil {
		modelMap["evaluation_delay_ms"] = flex.IntValue(model.EvaluationDelayMs)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricFilterToMap(model *logsv0.ApisAlertDefinitionMetricFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["promql"] = *model.Promql
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdRuleToMap(model *logsv0.ApisAlertDefinitionMetricThresholdRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	overrideMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model.Override)
	if err != nil {
		return modelMap, err
	}
	modelMap["override"] = []map[string]interface{}{overrideMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdConditionToMap(model *logsv0.ApisAlertDefinitionMetricThresholdCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	ofTheLastMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowToMap(model logsv0.ApisAlertDefinitionMetricTimeWindowIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue); ok {
		return DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration); ok {
		return DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model.(*logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration))
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowSpecificValue != nil {
		modelMap["metric_time_window_specific_value"] = *model.MetricTimeWindowSpecificValue
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetricTimeWindowDynamicDuration != nil {
		modelMap["metric_time_window_dynamic_duration"] = *model.MetricTimeWindowDynamicDuration
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesToMap(model logsv0.ApisAlertDefinitionMetricMissingValuesIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero); ok {
		return DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero))
	} else if _, ok := model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct); ok {
		return DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model.(*logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct))
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ReplaceWithZero != nil {
		modelMap["replace_with_zero"] = *model.ReplaceWithZero
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MinNonNullValuesPct != nil {
		modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowTypeToMap(model *logsv0.ApisAlertDefinitionFlowType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	stages := []map[string]interface{}{}
	for _, stagesItem := range model.Stages {
		stagesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesToMap(&stagesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesToMap(model *logsv0.ApisAlertDefinitionFlowStages) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["timeframe_ms"] = *model.TimeframeMs
	modelMap["timeframe_type"] = *model.TimeframeType
	flowStagesGroupsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsToMap(model.FlowStagesGroups)
	if err != nil {
		return modelMap, err
	}
	modelMap["flow_stages_groups"] = []map[string]interface{}{flowStagesGroupsMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroups) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	groups := []map[string]interface{}{}
	for _, groupsItem := range model.Groups {
		groupsItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupToMap(&groupsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		groups = append(groups, groupsItemMap)
	}
	modelMap["groups"] = groups
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	alertDefs := []map[string]interface{}{}
	for _, alertDefsItem := range model.AlertDefs {
		alertDefsItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(&alertDefsItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model *logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.Not != nil {
		modelMap["not"] = *model.Not
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyRuleToMap(&rulesItem) // #nosec G601
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
		anomalyAlertSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionLogsAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["minimum_threshold"] = *model.MinimumThreshold
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAnomalyAlertSettingsToMap(model *logsv0.ApisAlertDefinitionAnomalyAlertSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PercentageOfDeviation != nil {
		modelMap["percentage_of_deviation"] = flex.Float64Value(model.PercentageOfDeviation)
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyTypeToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	metricFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricFilterToMap(model.MetricFilter)
	if err != nil {
		return modelMap, err
	}
	modelMap["metric_filter"] = []map[string]interface{}{metricFilterMap}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyRuleToMap(&rulesItem) // #nosec G601
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
		anomalyAlertSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAnomalyAlertSettingsToMap(model.AnomalyAlertSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["anomaly_alert_settings"] = []map[string]interface{}{anomalyAlertSettingsMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyRuleToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyConditionToMap(model *logsv0.ApisAlertDefinitionMetricAnomalyCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["threshold"] = *model.Threshold
	if model.ForOverPct != nil {
		modelMap["for_over_pct"] = flex.IntValue(model.ForOverPct)
	}
	ofTheLastMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowToMap(model.OfTheLast)
	if err != nil {
		return modelMap, err
	}
	modelMap["of_the_last"] = []map[string]interface{}{ofTheLastMap}
	modelMap["min_non_null_values_pct"] = flex.IntValue(model.MinNonNullValuesPct)
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTypeToMap(model *logsv0.ApisAlertDefinitionLogsNewValueType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueRuleToMap(model *logsv0.ApisAlertDefinitionLogsNewValueRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueConditionToMap(model *logsv0.ApisAlertDefinitionLogsNewValueCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["keypath_to_track"] = *model.KeypathToTrack
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsNewValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_new_value_time_window_specific_value"] = *model.LogsNewValueTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountTypeToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogsFilter != nil {
		logsFilterMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model.LogsFilter)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_filter"] = []map[string]interface{}{logsFilterMap}
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountRuleToMap(&rulesItem) // #nosec G601
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

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountRuleToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountRule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	conditionMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountConditionToMap(model.Condition)
	if err != nil {
		return modelMap, err
	}
	modelMap["condition"] = []map[string]interface{}{conditionMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountConditionToMap(model *logsv0.ApisAlertDefinitionLogsUniqueCountCondition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["max_unique_count"] = *model.MaxUniqueCount
	timeWindowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model.TimeWindow)
	if err != nil {
		return modelMap, err
	}
	modelMap["time_window"] = []map[string]interface{}{timeWindowMap}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model *logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["logs_unique_value_time_window_specific_value"] = *model.LogsUniqueValueTimeWindowSpecificValue
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediateToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsImmediate != nil {
		logsImmediateMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsImmediateTypeToMap(model.LogsImmediate)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_immediate"] = []map[string]interface{}{logsImmediateMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThresholdToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsThreshold != nil {
		logsThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdTypeToMap(model.LogsThreshold)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_threshold"] = []map[string]interface{}{logsThresholdMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThresholdToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsRatioThreshold != nil {
		logsRatioThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioThresholdTypeToMap(model.LogsRatioThreshold)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_ratio_threshold"] = []map[string]interface{}{logsRatioThresholdMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThresholdToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsTimeRelativeThreshold != nil {
		logsTimeRelativeThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model.LogsTimeRelativeThreshold)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_time_relative_threshold"] = []map[string]interface{}{logsTimeRelativeThresholdMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThresholdToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.MetricThreshold != nil {
		metricThresholdMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdTypeToMap(model.MetricThreshold)
		if err != nil {
			return modelMap, err
		}
		modelMap["metric_threshold"] = []map[string]interface{}{metricThresholdMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlowToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.Flow != nil {
		flowMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowTypeToMap(model.Flow)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow"] = []map[string]interface{}{flowMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomalyToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsAnomaly != nil {
		logsAnomalyMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyTypeToMap(model.LogsAnomaly)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_anomaly"] = []map[string]interface{}{logsAnomalyMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomalyToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.MetricAnomaly != nil {
		metricAnomalyMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyTypeToMap(model.MetricAnomaly)
		if err != nil {
			return modelMap, err
		}
		modelMap["metric_anomaly"] = []map[string]interface{}{metricAnomalyMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValueToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsNewValue != nil {
		logsNewValueMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTypeToMap(model.LogsNewValue)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_new_value"] = []map[string]interface{}{logsNewValueMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCountToMap(model *logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.CreatedTime != nil {
		modelMap["created_time"] = model.CreatedTime.String()
	}
	if model.UpdatedTime != nil {
		modelMap["updated_time"] = model.UpdatedTime.String()
	}
	if model.AlertVersionID != nil {
		modelMap["alert_version_id"] = model.AlertVersionID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.ActiveOn != nil {
		activeOnMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model.ActiveOn)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_on"] = []map[string]interface{}{activeOnMap}
	}
	modelMap["type"] = *model.Type
	modelMap["group_by_keys"] = model.GroupByKeys
	if model.IncidentsSettings != nil {
		incidentsSettingsMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model.IncidentsSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["incidents_settings"] = []map[string]interface{}{incidentsSettingsMap}
	}
	if model.NotificationGroup != nil {
		notificationGroupMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model.NotificationGroup)
		if err != nil {
			return modelMap, err
		}
		modelMap["notification_group"] = []map[string]interface{}{notificationGroupMap}
	}
	if model.EntityLabels != nil {
		entityLabels := make(map[string]interface{})
		for k, v := range model.EntityLabels {
			entityLabels[k] = flex.Stringify(v)
		}
		modelMap["entity_labels"] = entityLabels
	}
	if model.PhantomMode != nil {
		modelMap["phantom_mode"] = *model.PhantomMode
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.LogsUniqueCount != nil {
		logsUniqueCountMap, err := DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountTypeToMap(model.LogsUniqueCount)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_unique_count"] = []map[string]interface{}{logsUniqueCountMap}
	}
	return modelMap, nil
}
