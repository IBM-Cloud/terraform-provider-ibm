// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
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

func DataSourceIbmBackupRecoveryRegistrationInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryRegistrationInfoRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"environments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Return only the registered root nodes whose Ids are given in the list.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"include_entity_permission_info": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If specified, then a list of entities with permissions assigned to them are returned.",
			},
			"sids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter the registered root nodes for the sids given in the list.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_source_credentials": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If specified, then crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied 'encryption_key'.",
			},
			"encryption_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Key to be used to encrypt the source credential. If include_source_credentials is set to true this key must be specified.",
			},
			"include_applications_tree_info": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to return applications tree info or not.",
			},
			"prune_non_critical_info": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to prune non critical info within entities. Incase of VMs, virtual disk information will be pruned. Incase of Office365, metadata about user entities will be pruned. This can be used to limit the size of the response by caller.",
			},
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of the request. Possible values are UIUser and UIAuto, which means the request is triggered by user or is an auto refresh request. Services like magneto will use this to determine the priority of the requests, so that it can more intelligently handle overload situations by prioritizing higher priority requests.",
			},
			"use_cached_data": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether we can serve the GET request to the read replica cache. setting this to true ensures that the API request is served to the read replica. setting this to false will serve the request to the master.",
			},
			"include_external_metadata": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if entity external metadata should be included within the response to get entity hierarchy call.",
			},
			"maintenance_status": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the maintenance status of a source 'UnderMaintenance' indicates the source is currently under maintenance. 'ScheduledMaintenance' indicates the source is scheduled for maintenance. 'NotConfigured' indicates maintenance is not configured on the source.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which objects are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"all_under_hierarchy": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.",
			},
			"root_nodes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the registration, protection and permission information of either all or a subset of registered Protection Sources matching the filter parameters. overrideDescription: true.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"applications": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of applications hierarchy registered on this node. Specifies the application type and the list of instances of the application objects. For example for SQL Server, this list provides the SQL Server instances running on a VM or a Physical Server.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_tree_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Application Server and the subtrees below them. Specifies the application subtree used to store additional application level Objects. Different environments use the subtree to store application level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM or a Physical Server. overrideDescription: true.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the connection id of the tenant.",
												},
												"connector_group_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the connector group id of the connector groups.",
												},
												"custom_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the user provided custom name of the Protection Source.",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment (such as 'kVMware' or 'kSQL') where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies an id of the Protection Source.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a name of the Protection Source.",
												},
												"parent_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies an id of the parent of the Protection Source.",
												},
												"physical_protection_source": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies a Protection Source in a Physical environment.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"agents": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifiles the agents running on the Physical Protection Source and the status information.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cbmr_version": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the version if Cristie BMR product is installed on the host.",
																		},
																		"file_cbt_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "CBT version and service state info.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"file_version": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Subcomponent version. The interpretation of the version is based on operating system.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"build_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"major_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"minor_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"revision_num": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"is_installed": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Indicates whether the cbt driver is installed.",
																					},
																					"reboot_status": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Indicates whether host is rebooted post VolCBT installation.",
																					},
																					"service_state": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Structure to Hold Service Status.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"state": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"host_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the host type where the agent is running. This is only set for persistent agents.",
																		},
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the agent's id.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the agent's name.",
																		},
																		"oracle_multi_node_channel_supported": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether oracle multi node multi channel is supported or not.",
																		},
																		"registration_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies information about a registered Source.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the parameters required to establish a connection with a particular environment.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"connection_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.",
																								},
																								"connector_group_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.",
																								},
																								"endpoint": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).",
																								},
																								"environment": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																								},
																								"id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.",
																								},
																								"version": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.",
																								},
																							},
																						},
																					},
																					"allowed_ip_addresses": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"authentication_error_message": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.",
																					},
																					"authentication_status": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster.",
																					},
																					"blacklisted_ip_addresses": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "This field is deprecated. Use DeniedIpAddresses instead.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"denied_ip_addresses": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"environments": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"is_db_authenticated": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if application entity dbAuthenticated or not.",
																					},
																					"is_storage_array_snapshot_enabled": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if this source entity has enabled storage array snapshot or not.",
																					},
																					"link_vms_across_vcenter": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.",
																					},
																					"minimum_free_space_gb": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
																					},
																					"minimum_free_space_percent": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
																					},
																					"password": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies password of the username to access the target source.",
																					},
																					"physical_params": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"applications": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"password": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies password of the username to access the target source.",
																								},
																								"throttling_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the source side throttling configuration.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"cpu_throttling_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the Throttling Configuration Parameters.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"fixed_threshold": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																														},
																														"pattern_type": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																														},
																														"throttling_windows": &schema.Schema{
																															Type:     schema.TypeList,
																															Computed: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"day_time_window": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Window Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"end_time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the Day Time Parameters.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"day": &schema.Schema{
																																								Type:        schema.TypeString,
																																								Computed:    true,
																																								Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																							},
																																							"time": &schema.Schema{
																																								Type:        schema.TypeList,
																																								Computed:    true,
																																								Description: "Specifies the time in hours and minutes.",
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"hour": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the hour of this time.",
																																										},
																																										"minute": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the minute of this time.",
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																				"start_time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the Day Time Parameters.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"day": &schema.Schema{
																																								Type:        schema.TypeString,
																																								Computed:    true,
																																								Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																							},
																																							"time": &schema.Schema{
																																								Type:        schema.TypeList,
																																								Computed:    true,
																																								Description: "Specifies the time in hours and minutes.",
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"hour": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the hour of this time.",
																																										},
																																										"minute": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the minute of this time.",
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
																																	"threshold": &schema.Schema{
																																		Type:        schema.TypeInt,
																																		Computed:    true,
																																		Description: "Throttling threshold applicable in the window.",
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																											"network_throttling_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the Throttling Configuration Parameters.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"fixed_threshold": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																														},
																														"pattern_type": &schema.Schema{
																															Type:        schema.TypeString,
																															Computed:    true,
																															Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																														},
																														"throttling_windows": &schema.Schema{
																															Type:     schema.TypeList,
																															Computed: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"day_time_window": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Window Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"end_time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the Day Time Parameters.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"day": &schema.Schema{
																																								Type:        schema.TypeString,
																																								Computed:    true,
																																								Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																							},
																																							"time": &schema.Schema{
																																								Type:        schema.TypeList,
																																								Computed:    true,
																																								Description: "Specifies the time in hours and minutes.",
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"hour": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the hour of this time.",
																																										},
																																										"minute": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the minute of this time.",
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																				"start_time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the Day Time Parameters.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"day": &schema.Schema{
																																								Type:        schema.TypeString,
																																								Computed:    true,
																																								Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																							},
																																							"time": &schema.Schema{
																																								Type:        schema.TypeList,
																																								Computed:    true,
																																								Description: "Specifies the time in hours and minutes.",
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"hour": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the hour of this time.",
																																										},
																																										"minute": &schema.Schema{
																																											Type:        schema.TypeInt,
																																											Computed:    true,
																																											Description: "Specifies the minute of this time.",
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
																																	"threshold": &schema.Schema{
																																		Type:        schema.TypeInt,
																																		Computed:    true,
																																		Description: "Throttling threshold applicable in the window.",
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
																								"username": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies username to access the target source.",
																								},
																							},
																						},
																					},
																					"progress_monitor_path": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Captures the current progress and pulse details w.r.t to either the registration or refresh.",
																					},
																					"refresh_error_message": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.",
																					},
																					"refresh_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.",
																					},
																					"registered_apps_info": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies information of the applications registered on this protection source.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"authentication_error_message": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.",
																								},
																								"authentication_status": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.",
																								},
																								"environment": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																								},
																								"host_settings_check_results": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"check_type": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.",
																											},
																											"result_type": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.",
																											},
																											"user_message": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies a descriptive message for failed/warning types.",
																											},
																										},
																									},
																								},
																								"refresh_error_message": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.",
																								},
																							},
																						},
																					},
																					"registration_time_usecs": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.",
																					},
																					"subnets": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"component": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Component that has reserved the subnet.",
																								},
																								"description": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Description of the subnet.",
																								},
																								"id": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "ID of the subnet.",
																								},
																								"ip": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies either an IPv6 address or an IPv4 address.",
																								},
																								"netmask_bits": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "netmaskBits.",
																								},
																								"netmask_ip4": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.",
																								},
																								"nfs_access": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Component that has reserved the subnet.",
																								},
																								"nfs_all_squash": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.",
																								},
																								"nfs_root_squash": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether clients from this subnet can mount as root on NFS.",
																								},
																								"s3_access": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
																								},
																								"smb_access": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
																								},
																								"tenant_id": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the unique id of the tenant.",
																								},
																							},
																						},
																					},
																					"throttling_policy": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the throttling policy for a registered Protection Source.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"enforce_max_streams": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
																								},
																								"enforce_registered_source_max_backups": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
																								},
																								"is_enabled": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
																								},
																								"latency_thresholds": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"active_task_msecs": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
																											},
																											"new_task_msecs": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
																											},
																										},
																									},
																								},
																								"max_concurrent_streams": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
																								},
																								"nas_source_params": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
																											},
																											"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
																											},
																											"max_parallel_read_write_full_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
																											},
																											"max_parallel_read_write_incremental_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
																											},
																										},
																									},
																								},
																								"registered_source_max_concurrent_backups": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
																								},
																								"storage_array_snapshot_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies Storage Array Snapshot Configuration.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"is_max_snapshots_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																											},
																											"is_max_space_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																											},
																											"storage_array_snapshot_max_space_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies Storage Array Snapshot Max Space Config.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"max_snapshot_space_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Max number of storage snapshots allowed per volume/lun.",
																														},
																													},
																												},
																											},
																											"storage_array_snapshot_throttling_policies": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies throttling policies configured for individual volume/lun.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"id": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the volume id of the storage array snapshot config.",
																														},
																														"is_max_snapshots_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																														},
																														"is_max_space_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																														},
																														"max_snapshot_config": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies Storage Array Snapshot Max Snapshots Config.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"max_snapshots": &schema.Schema{
																																		Type:        schema.TypeFloat,
																																		Computed:    true,
																																		Description: "Max number of storage snapshots allowed per volume/lun.",
																																	},
																																},
																															},
																														},
																														"max_space_config": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies Storage Array Snapshot Max Space Config.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"max_snapshot_space_percentage": &schema.Schema{
																																		Type:        schema.TypeFloat,
																																		Computed:    true,
																																		Description: "Max number of storage snapshots allowed per volume/lun.",
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
																					"throttling_policy_overrides": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"datastore_id": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Specifies the Protection Source id of the Datastore.",
																								},
																								"datastore_name": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the display name of the Datastore.",
																								},
																								"throttling_policy": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the throttling policy for a registered Protection Source.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"enforce_max_streams": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
																											},
																											"enforce_registered_source_max_backups": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
																											},
																											"is_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
																											},
																											"latency_thresholds": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"active_task_msecs": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
																														},
																														"new_task_msecs": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
																														},
																													},
																												},
																											},
																											"max_concurrent_streams": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
																											},
																											"nas_source_params": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
																														},
																														"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
																														},
																														"max_parallel_read_write_full_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
																														},
																														"max_parallel_read_write_incremental_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
																														},
																													},
																												},
																											},
																											"registered_source_max_concurrent_backups": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
																											},
																											"storage_array_snapshot_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies Storage Array Snapshot Configuration.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"is_max_snapshots_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																														},
																														"is_max_space_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																														},
																														"storage_array_snapshot_max_space_config": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies Storage Array Snapshot Max Space Config.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"max_snapshot_space_percentage": &schema.Schema{
																																		Type:        schema.TypeFloat,
																																		Computed:    true,
																																		Description: "Max number of storage snapshots allowed per volume/lun.",
																																	},
																																},
																															},
																														},
																														"storage_array_snapshot_throttling_policies": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies throttling policies configured for individual volume/lun.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"id": &schema.Schema{
																																		Type:        schema.TypeInt,
																																		Computed:    true,
																																		Description: "Specifies the volume id of the storage array snapshot config.",
																																	},
																																	"is_max_snapshots_config_enabled": &schema.Schema{
																																		Type:        schema.TypeBool,
																																		Computed:    true,
																																		Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																																	},
																																	"is_max_space_config_enabled": &schema.Schema{
																																		Type:        schema.TypeBool,
																																		Computed:    true,
																																		Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																																	},
																																	"max_snapshot_config": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies Storage Array Snapshot Max Snapshots Config.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"max_snapshots": &schema.Schema{
																																					Type:        schema.TypeFloat,
																																					Computed:    true,
																																					Description: "Max number of storage snapshots allowed per volume/lun.",
																																				},
																																			},
																																		},
																																	},
																																	"max_space_config": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies Storage Array Snapshot Max Space Config.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"max_snapshot_space_percentage": &schema.Schema{
																																					Type:        schema.TypeFloat,
																																					Computed:    true,
																																					Description: "Max number of storage snapshots allowed per volume/lun.",
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
																					"use_o_auth_for_exchange_online": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether OAuth should be used for authentication in case of Exchange Online.",
																					},
																					"use_vm_bios_uuid": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if registered vCenter is using BIOS UUID to track virtual machines.",
																					},
																					"user_messages": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"username": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies username to access the target source.",
																					},
																					"vlan_params": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the VLAN configuration for Recovery.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"vlan": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																								},
																								"disable_vlan": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.",
																								},
																								"interface_name": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																								},
																							},
																						},
																					},
																					"warning_messages": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},
																		"source_side_dedup_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether source side dedup is enabled or not.",
																		},
																		"status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the agent status. Specifies the status of the agent running on a physical source.",
																		},
																		"status_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies additional details about the agent status.",
																		},
																		"upgradability": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the upgradability of the agent running on the physical server. Specifies the upgradability of the agent running on the physical server.",
																		},
																		"upgrade_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the status of the upgrade of the agent on a physical server. Specifies the status of the upgrade of the agent on a physical server.",
																		},
																		"upgrade_status_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade.",
																		},
																		"version": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the version of the Agent software.",
																		},
																		"vol_cbt_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "CBT version and service state info.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"file_version": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Subcomponent version. The interpretation of the version is based on operating system.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"build_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"major_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"minor_ver": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																								"revision_num": &schema.Schema{
																									Type:     schema.TypeFloat,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"is_installed": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Indicates whether the cbt driver is installed.",
																					},
																					"reboot_status": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Indicates whether host is rebooted post VolCBT installation.",
																					},
																					"service_state": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Structure to Hold Service Status.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"state": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
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
															"cluster_source_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of cluster resource this source represents.",
															},
															"host_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the hostname.",
															},
															"host_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the environment type for the host.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies an id for an object that is unique across Cohesity Clusters. The id is composite of all the ids listed below.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cluster_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the Cohesity Cluster id where the object was created.",
																		},
																		"cluster_incarnation_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies an id for the Cohesity Cluster that is generated when a Cohesity Cluster is initially created.",
																		},
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies a unique id assigned to an object (such as a Job) by the Cohesity Cluster.",
																		},
																	},
																},
															},
															"is_proxy_host": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the physical host is a proxy host.",
															},
															"memory_size_bytes": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the total memory on the host in bytes.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a human readable name of the Protection Source.",
															},
															"networking_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the struct containing information about network addresses configured on the given box. This is needed for dealing with Windows/Oracle Cluster resources that we discover and protect automatically.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"resource_vec": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of resources on the system that are accessible by an IP address.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"endpoints": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "The endpoints by which the resource is accessible.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"fqdn": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The Fully Qualified Domain Name.",
																								},
																								"ipv4_addr": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The IPv4 address.",
																								},
																								"ipv6_addr": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The IPv6 address.",
																								},
																							},
																						},
																					},
																					"type": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The type of the resource.",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"num_processors": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the number of processors on the host.",
															},
															"os_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a human readable name of the OS of the Protection Source.",
															},
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of managed Object in a Physical Protection Source. 'kGroup' indicates the EH container.",
															},
															"vcs_version": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies cluster version for VCS host.",
															},
															"volumes": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Array of Physical Volumes. Specifies the volumes available on the physical host. These fields are populated only for the kPhysicalHost type.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"device_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the path to the device that hosts the volume locally.",
																		},
																		"guid": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies an id for the Physical Volume.",
																		},
																		"is_boot_volume": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the volume is boot volume.",
																		},
																		"is_extended_attributes_supported": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether this volume supports extended attributes (like ACLs) when performing file backups.",
																		},
																		"is_protected": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if a volume is protected by a Job.",
																		},
																		"is_shared_volume": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the volume is shared volume.",
																		},
																		"label": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies a volume label that can be used for displaying additional identifying information about a volume.",
																		},
																		"logical_size_bytes": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the logical size of the volume in bytes that is not reduced by change-block tracking, compression and deduplication.",
																		},
																		"mount_points": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the mount points where the volume is mounted, for example- 'C:', '/mnt/foo' etc.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"mount_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies mount type of volume e.g. nfs, autofs, ext4 etc.",
																		},
																		"network_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the full path to connect to the network attached volume. For example, (IP or hostname):/path/to/share for NFS volumes).",
																		},
																		"used_size_bytes": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the size used by the volume in bytes.",
																		},
																	},
																},
															},
															"vsswriters": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"is_writer_excluded": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If true, the writer will be excluded by default.",
																		},
																		"writer_name": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies the name of the writer.",
																		},
																	},
																},
															},
														},
													},
												},
												"kubernetes_protection_source": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies a Protection Source in Kubernetes environment.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"datamover_image_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the location of Datamover image in private registry.",
															},
															"datamover_service_type": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies Type of service to be deployed for communication with DataMover pods. Currently, LoadBalancer and NodePort are supported. [default = kNodePort].",
															},
															"datamover_upgradability": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies if the deployed Datamover image needs to be upgraded for this kubernetes entity.",
															},
															"default_vlan_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies VLAN parameters for the restore operation.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"disable_vlan": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.",
																		},
																		"interface_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																		},
																		"vlan": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																		},
																	},
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies an optional description of the object.",
															},
															"distribution": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the entity in a Kubernetes environment. Determines the K8s distribution. kIKS, kROKS.",
															},
															"init_container_image_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the location of the image for init containers.",
															},
															"label_attributes": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the list of label attributes of this source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the Cohesity id of the K8s label.",
																		},
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the appended key and value of the K8s label.",
																		},
																		"uuid": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies Kubernetes Unique Identifier (UUID) of the K8s label.",
																		},
																	},
																},
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a unique name of the Protection Source.",
															},
															"priority_class_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the pritority class name during registration.",
															},
															"resource_annotation_list": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies resource Annotations information provided during registration.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Key for label.",
																		},
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Value for label.",
																		},
																	},
																},
															},
															"resource_label_list": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies resource labels information provided during registration.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Key for label.",
																		},
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Value for label.",
																		},
																	},
																},
															},
															"san_field": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the SAN field for agent certificate.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"service_annotations": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"value": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"storage_class": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies storage class information of source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies name of storage class.",
																		},
																		"provisioner": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "specifies provisioner of storage class.",
																		},
																	},
																},
															},
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the entity in a Kubernetes environment. Specifies the type of a Kubernetes Protection Source. 'kCluster' indicates a Kubernetes Cluster. 'kNamespace' indicates a namespace in a Kubernetes Cluster. 'kService' indicates a service running on a Kubernetes Cluster.",
															},
															"uuid": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the UUID of the object.",
															},
															"velero_aws_plugin_image_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the location of Velero AWS plugin image in private registry.",
															},
															"velero_image_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the location of Velero image in private registry.",
															},
															"velero_openshift_plugin_image_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the location of the image for openshift plugin container.",
															},
															"velero_upgradability": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies if the deployed Velero image needs to be upgraded for this kubernetes entity.",
															},
															"vlan_info_vec": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies VLAN information provided during registration.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"service_annotations": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the service annotation key value.",
																					},
																					"value": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the service annotation value.",
																					},
																				},
																			},
																		},
																		"vlan_params": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies VLAN params associated with the backup/restore operation.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"disable_vlan": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
																					},
																					"interface_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
																					},
																					"vlan_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
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
												"sql_protection_source": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object representing one SQL Server instance or database.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_available_for_vss_backup": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the database is marked as available for backup according to the SQL Server VSS writer. This may be false if either the state of the databases is not online, or if the VSS writer is not online. This field is set only for type 'kDatabase'.",
															},
															"created_timestamp": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the time when the database was created. It is displayed in the timezone of the SQL server on which this database is running.",
															},
															"database_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the database name of the SQL Protection Source, if the type is database.",
															},
															"db_aag_entity_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the AAG entity id if the database is part of an AAG. This field is set only for type 'kDatabase'.",
															},
															"db_aag_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the AAG if the database is part of an AAG. This field is set only for type 'kDatabase'.",
															},
															"db_compatibility_level": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the versions of SQL server that the database is compatible with.",
															},
															"db_file_groups": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the information about the set of file groups for this db on the host. This is only set if the type is kDatabase.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"db_files": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the last known information about the set of database files on the host. This field is set only for type 'kDatabase'.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"file_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. 'kRows' refers to a data file 'kLog' refers to a log file 'kFileStream' refers to a directory containing FILESTREAM data 'kNotSupportedType' is for information purposes only. Not supported. 'kFullText' refers to a full-text catalog.",
																		},
																		"full_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the full path of the database file on the SQL host machine.",
																		},
																		"size_bytes": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the last known size of the database file.",
																		},
																	},
																},
															},
															"db_owner_username": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the database owner.",
															},
															"default_database_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the default path for data files for DBs in an instance.",
															},
															"default_log_location": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the default path for log files for DBs in an instance.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies a unique id for a SQL Protection Source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"created_date_msecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies a unique identifier generated from the date the database is created or renamed. Cohesity uses this identifier in combination with the databaseId to uniquely identify a database.",
																		},
																		"database_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies a unique id of the database but only for the life of the database. SQL Server may reuse database ids. Cohesity uses the createDateMsecs in combination with this databaseId to uniquely identify a database.",
																		},
																		"instance_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies unique id for the SQL Server instance. This id does not change during the life of the instance.",
																		},
																	},
																},
															},
															"is_encrypted": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the database is TDE enabled.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the instance name of the SQL Protection Source.",
															},
															"owner_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the id of the container VM for the SQL Protection Source.",
															},
															"recovery_model": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the Recovery Model for the database in SQL environment. Only meaningful for the 'kDatabase' SQL Protection Source. Specifies the Recovery Model set for the Microsoft SQL Server. 'kSimpleRecoveryModel' indicates the Simple SQL Recovery Model which does not utilize log backups. 'kFullRecoveryModel' indicates the Full SQL Recovery Model which requires log backups and allows recovery to a single point in time. 'kBulkLoggedRecoveryModel' indicates the Bulk Logged SQL Recovery Model which requires log backups and allows high-performance bulk copy operations.",
															},
															"sql_server_db_state": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The state of the database as returned by SQL Server. Indicates the state of the database. The values correspond to the 'state' field in the system table sys.databases. See https://goo.gl/P66XqM. 'kOnline' indicates that database is in online state. 'kRestoring' indicates that database is in restore state. 'kRecovering' indicates that database is in recovery state. 'kRecoveryPending' indicates that database recovery is in pending state. 'kSuspect' indicates that primary filegroup is suspect and may be damaged. 'kEmergency' indicates that manually forced emergency state. 'kOffline' indicates that database is in offline state. 'kCopying' indicates that database is in copying state. 'kOfflineSecondary' indicates that secondary database is in offline state.",
															},
															"sql_server_instance_version": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the Server Instance Version.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"build": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the build.",
																		},
																		"major_version": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the major version.",
																		},
																		"minor_version": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the minor version.",
																		},
																		"revision": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the revision.",
																		},
																		"version_string": &schema.Schema{
																			Type:        schema.TypeFloat,
																			Computed:    true,
																			Description: "Specifies the version string.",
																		},
																	},
																},
															},
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the managed Object in a SQL Protection Source. Examples of SQL Objects include 'kInstance' and 'kDatabase'. 'kInstance' indicates that SQL server instance is being protected. 'kDatabase' indicates that SQL server database is being protected. 'kAAG' indicates that SQL AAG (AlwaysOn Availability Group) is being protected. 'kAAGRootContainer' indicates that SQL AAG's root container is being protected. 'kRootContainer' indicates root container for SQL sources.",
															},
														},
													},
												},
											},
										},
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the environment type of the application such as 'kSQL', 'kExchange' registered on the Protection Source. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment.'kSQL' indicates the SQL Protection Source environment.'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment.'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment.'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
									},
								},
							},
						},
						"entity_permission_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the permission information of entities.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entity_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the entity id.",
									},
									"groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies groups that have access to entity in case of restricted user.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies domain name of the user.",
												},
												"group_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies group name of the group.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies unique Security ID (SID) of the user.",
												},
												"tenant_ids": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the tenants to which the group belongs to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"is_inferred": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the Entity Permission Information is inferred or not. For example, SQL application hosted over vCenter will have inferred entity permission information.",
									},
									"is_registered_by_sp": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether this entity is registered by the SP or not. This will be populated only if the entity is a root entity. Refer to magneto/base/permissions.proto for details.",
									},
									"registering_tenant_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the tenant id that registered this entity. This will be populated only if the entity is a root entity.",
									},
									"tenant": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies struct with basic tenant details.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bifrost_enabled": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if this tenant is bifrost enabled or not.",
												},
												"is_managed_on_helios": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether this tenant is manged on helios.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies name of the tenant.",
												},
												"tenant_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the unique id of the tenant.",
												},
											},
										},
									},
									"users": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies users that have access to entity in case of restricted user.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies domain name of the user.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies unique Security ID (SID) of the user.",
												},
												"tenant_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the tenant to which the user belongs to.",
												},
												"user_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies user name of the user.",
												},
											},
										},
									},
								},
							},
						},
						"logical_size_bytes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the logical size of the Protection Source in bytes.",
						},
						"maintenance_mode_config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the entity metadata for maintenance mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activation_time_intervals": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the absolute intervals where the maintenance schedule is valid, i.e. maintenance_shedule is considered only for these time ranges. (For example, if there is one time range with [now_usecs, now_usecs + 10 days], the action will be done during the maintenance_schedule for the next 10 days.)The start time must be specified. The end time can be -1 which would denote an indefinite maintenance mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the end time of this time range.",
												},
												"start_time_usecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the start time of this time range.",
												},
											},
										},
									},
									"maintenance_schedule": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a schedule for actions to be taken.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"periodic_time_windows": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the time range within the days of the week.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"day_of_the_week": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the week day.",
															},
															"end_time": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the time in hours and minutes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hour": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the hour of this time.",
																		},
																		"minute": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the minute of this time.",
																		},
																	},
																},
															},
															"start_time": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the time in hours and minutes.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hour": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the hour of this time.",
																		},
																		"minute": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the minute of this time.",
																		},
																	},
																},
															},
														},
													},
												},
												"schedule_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of schedule for this ScheduleProto.",
												},
												"time_ranges": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the time ranges in usecs.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"end_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the end time of this time range.",
															},
															"start_time_usecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the start time of this time range.",
															},
														},
													},
												},
												"timezone": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the timezone of the user of this ScheduleProto. The timezones have unique names of the form 'Area/Location'.",
												},
											},
										},
									},
									"user_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "User provided message associated with this maintenance mode.",
									},
									"workflow_intervention_spec_list": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the type of intervention for different workflows when the source goes into maintenance mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"intervention": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the intervention type for ongoing tasks.",
												},
												"workflow_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the workflow type for which an intervention would be needed when maintenance mode begins.",
												},
											},
										},
									},
								},
							},
						},
						"registration_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies registration information for a root node in a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters required to establish a connection with a particular environment.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.",
												},
												"connector_group_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.",
												},
												"endpoint": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment.'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.",
												},
												"version": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.",
												},
											},
										},
									},
									"allowed_ip_addresses": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"authentication_error_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.",
									},
									"authentication_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.",
									},
									"blacklisted_ip_addresses": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "This field is deprecated. Use DeniedIpAddresses instead. deprecated: true.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cassandra_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered cassandra source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cassandra_ports_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object containing information on various Cassandra ports.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"jmx_port": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the Cassandra JMX port.",
															},
															"native_transport_port": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the port for the CQL native transport.",
															},
															"rpc_port": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the Remote Procedure Call (RPC) port for general mechanism for client-server applications.",
															},
															"ssl_storage_port": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the SSL port for encrypted communication.",
															},
															"storage_port": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the TCP port for data. Internally used by Cassandra bulk loader.",
															},
														},
													},
												},
												"cassandra_security_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object containing information on Cassandra security.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cassandra_auth_required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is Cassandra authentication required ?.",
															},
															"cassandra_auth_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cassandra Authentication type. Enum: [PASSWORD KERBEROS LDAP] Specifies the Cassandra auth type.'PASSWORD' 'KERBEROS' 'LDAP'.",
															},
															"cassandra_authorizer": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cassandra Authenticator/Authorizer.",
															},
															"client_encryption": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is Client Encryption enabled for this cluster ?.",
															},
															"dse_authorization": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is DSE Authorization enabled for this cluster ?.",
															},
															"server_encryption_req_client_auth": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is 'Server encryption request client authentication' enabled for this cluster ?.",
															},
															"server_internode_encryption_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "'Server internal node Encryption' type for this cluster.",
															},
														},
													},
												},
												"cassandra_version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Cassandra version.",
												},
												"commit_log_backup_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the commit log archival location for cassandra node.",
												},
												"config_directory": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Directory path containing Config YAML for discovery.",
												},
												"data_centers": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the List of all physical data center or virtual data center. In most cases, the data centers will be listed after discovery operation however, if they are not listed, you must manually type the data center names. Leaving the field blank will disallow data center-specific backup or restore. Entering a subset of all data centers may cause problems in data movement.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"dse_config_directory": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Directory from where DSE specific configuration can be read.",
												},
												"dse_version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "DSE version.",
												},
												"is_dse_authenticator": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether this cluster has DSE Authenticator.",
												},
												"is_dse_tiered_storage": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether this cluster has DSE tiered storage.",
												},
												"is_jmx_auth_enable": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if JMX Authentication enabled in this cluster.",
												},
												"kerberos_principal": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Kerberos Principal for Kerberos connection.",
												},
												"primary_host": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Primary Host for the Cassandra cluster.",
												},
												"seeds": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Seed nodes of this Cassandra cluster.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"solr_nodes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Solr node IP Addresses.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"solr_port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Solr node Port.",
												},
											},
										},
									},
									"couchbase_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered couchbase source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"carrier_direct_port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Carrier direct/sll port.",
												},
												"http_direct_port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the HTTP direct/sll port.",
												},
												"requires_ssl": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether this cluster allows connection through SSL only.",
												},
												"seeds": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Seeds of this Couchbase Cluster.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"denied_ip_addresses": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"environments": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hbase_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered HBase source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hbase_discovery_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object containing information about discovering a Hadoop source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"config_directory": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the configuration directory.",
															},
															"host": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the host IP.",
															},
														},
													},
												},
												"hdfs_entity_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The entity id of the HDFS source for this HBase.",
												},
												"kerberos_principal": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the kerberos principal.",
												},
												"root_data_directory": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the HBase data root directory.",
												},
												"zookeeper_quorum": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the HBase zookeeper quorum.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"hdfs_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered Hdfs source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hadoop_distribution": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Hadoop Distribution. Hadoop distribution. 'CDH' indicates Hadoop distribution type Cloudera. 'HDP' indicates Hadoop distribution type Hortonworks.",
												},
												"hadoop_version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Hadoop version.",
												},
												"hdfs_connection_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Hdfs connection type. Hdfs connection type. 'DFS' indicates Hdfs connection type DFS. 'WEBHDFS' indicates Hdfs connection type WEBHDFS. 'HTTPFSLB' indicates Hdfs connection type HTTPFS_LB. 'HTTPFS' indicates Hdfs connection type HTTPFS.",
												},
												"hdfs_discovery_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object containing information about discovering a Hadoop source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"config_directory": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the configuration directory.",
															},
															"host": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the host IP.",
															},
														},
													},
												},
												"kerberos_principal": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the kerberos principal.",
												},
												"namenode": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Namenode host or Nameservice.",
												},
												"port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Webhdfs Port.",
												},
											},
										},
									},
									"hive_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered Hive source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entity_threshold_exceeded": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if max entity count exceeded for protection source view.",
												},
												"hdfs_entity_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the entity id of the HDFS source for this Hive.",
												},
												"hive_discovery_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an Object containing information about discovering a Hadoop source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"config_directory": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the configuration directory.",
															},
															"host": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the host IP.",
															},
														},
													},
												},
												"kerberos_principal": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the kerberos principal.",
												},
												"metastore": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Hive metastore host.",
												},
												"thrift_port": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Hive metastore thrift Port.",
												},
											},
										},
									},
									"is_db_authenticated": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if application entity dbAuthenticated or not. ex: oracle database.",
									},
									"is_storage_array_snapshot_enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if this source entity has enabled storage array snapshot or not.",
									},
									"isilon_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the Isilon specific Registered Protection Source params. This definition is used to send isilion source params in update protection source params to magneto.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_config_list": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of access zone info in an Isilion Cluster.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dynamic_network_pool_config": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "While caonfiguring the isilon protection source, this is the selected network pool config for the isilon access zone.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"pool_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the Network pool.",
																		},
																		"subnet": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the name of the subnet the network pool belongs to.",
																		},
																		"use_smart_connect": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether to use SmartConnect if available. If true, DNS name for the SmartConnect zone will be used to balance the IPs. Otherwise, pool IPs will be balanced manually.",
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
									"link_vms_across_vcenter": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.",
									},
									"minimum_free_space_gb": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
									},
									"minimum_free_space_percent": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
									},
									"mongodb_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered mongodb source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies whether authentication is configured on this MongoDB cluster. Specifies the type of an MongoDB source entity. 'SCRAM' 'LDAP' 'NONE' 'KERBEROS'.",
												},
												"authenticating_database_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Authenticating Database for this MongoDB cluster.",
												},
												"requires_ssl": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether connection is allowed through SSL only in this cluster.",
												},
												"secondary_node_tag": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "MongoDB Secondary node tag. Required only if 'useSecondaryForBackup' is true. The system will use this to identify the secondary nodes for reading backup data.",
												},
												"seeds": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the seeds of this MongoDB Cluster.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"use_fixed_node_for_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Set this to true if you want the system to peform backups from fixed nodes.",
												},
												"use_secondary_for_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Set this to true if you want the system to peform backups from secondary nodes.",
												},
											},
										},
									},
									"nas_mount_credentials": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the credentials required to mount directories on the NetApp server if given.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the domain in which this credential is valid.",
												},
												"nas_protocol": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protocol used by the NAS server. Specifies the protocol used by a NAS server. 'kNoProtocol' indicates no protocol set. 'kNfs3' indicates NFS v3 protocol. 'kNfs4_1' indicates NFS v4.1 protocol. 'kCifs1' indicates CIFS v1.0 protocol. 'kCifs2' indicates CIFS v2.0 protocol. 'kCifs3' indicates CIFS v3.0 protocol.",
												},
											},
										},
									},
									"o365_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered Office 365 source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"objects_discovery_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the parameters used for discovering the office 365 objects selectively during source registration or refresh.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"discoverable_object_type_list": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the list of object types that will be discovered as part of source registration or refresh.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"sites_discovery_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies discovery params for kSite entities. It should only be populated when the 'DiscoveryParams.discoverableObjectTypeList' includes 'kSites'.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"enable_site_tagging": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the SharePoint Sites will be tagged whether they belong to a group site or teams site.",
																		},
																	},
																},
															},
															"teams_additional_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies additional params for Teams entities. It should only be populated if the 'DiscoveryParams.discoverableObjectTypeList' includes 'kTeams' otherwise this will be ignored.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"allow_posts_backup": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether the Teams posts/conversations will be backed up or not. If this is false or not specified teams' posts backup will not be done.",
																		},
																	},
																},
															},
															"users_discovery_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies discovery params for kUser entities. It should only be populated when the 'DiscoveryParams.discoverableObjectTypeList' includes 'kUsers'.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"allow_chats_backup": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether users' chats should be backed up or not. If this is false or not specified users' chats backup will not be done.",
																		},
																		"discover_users_with_mailbox": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if office 365 users with valid mailboxes should be discovered or not.",
																		},
																		"discover_users_with_onedrive": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if office 365 users with valid Onedrives should be discovered or not.",
																		},
																		"fetch_mailbox_info": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether users' mailbox info including the provisioning status, mailbox type & in-place archival support will be fetched and processed.",
																		},
																		"fetch_one_drive_info": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether users' onedrive info including the provisioning status & storage quota will be fetched and processed.",
																		},
																		"skip_users_without_my_site": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether to skip processing user who have uninitialized OneDrive or are without MySite.",
																		},
																	},
																},
															},
														},
													},
												},
												"csm_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backup_allowed": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the current source allows data backup through M365 Backup Storage APIs. Enabling this, data can be optionally backed up within either Cohesity or MSFT or both depending on the backup configuration.",
															},
														},
													},
												},
											},
										},
									},
									"office365_credentials_list": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Office365 Source Credentials. Specifies credentials needed to authenticate & authorize user for Office365.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the application ID that the registration portal (apps.dev.microsoft.com) assigned.",
												},
												"client_secret": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the application secret that was created in app registration portal.",
												},
												"grant_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the application grant type. eg: For client credentials flow, set this to \"client_credentials\"; For refreshing access-token, set this to \"refresh_token\".",
												},
												"scope": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a space separated list of scopes/permissions for the user. eg: Incase of MS Graph APIs for Office365, scope is set to default: https://graph.microsoft.com/.default.",
												},
												"use_o_auth_for_exchange_online": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "This field is deprecated from here and placed in RegisteredSourceInfo  and ProtectionSourceParameters. deprecated: true.",
												},
											},
										},
									},
									"office365_region": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the region for Office365. Inorder to truly categorize M365 region, clients should not depend upon the endpoint, instead look at this attribute for the same.",
									},
									"office365_service_account_credentials_list": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Office365 Service Account Credentials. Specifies credentials for improving mailbox backup performance for O365.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"username": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the username to access target entity.",
												},
												"password": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the password to access target entity.",
												},
											},
										},
									},
									"password": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies password of the username to access the target source.",
									},
									"physical_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"applications": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"password": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies password of the username to access the target source.",
												},
												"throttling_config": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the source side throttling configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cpu_throttling_config": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"fixed_threshold": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																		},
																		"pattern_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																		},
																		"throttling_windows": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Throttling windows which will be applicable in case of pattern_typec = kScheduleBased.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_time_window": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the Day Time Window Parameters.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"end_time": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Day Time Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"day": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																											},
																											"time": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the time in hours and minutes.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"hour": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the hour of this time.",
																														},
																														"minute": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the minute of this time.",
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																								"start_time": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Day Time Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"day": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																											},
																											"time": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the time in hours and minutes.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"hour": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the hour of this time.",
																														},
																														"minute": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the minute of this time.",
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
																					"threshold": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Throttling threshold applicable in the window.",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"network_throttling_config": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"fixed_threshold": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																		},
																		"pattern_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																		},
																		"throttling_windows": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Throttling windows which will be applicable in case of pattern_typec = kScheduleBased.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"day_time_window": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the Day Time Window Parameters.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"end_time": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Day Time Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"day": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																											},
																											"time": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the time in hours and minutes.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"hour": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the hour of this time.",
																														},
																														"minute": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the minute of this time.",
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																								"start_time": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Day Time Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"day": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																											},
																											"time": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies the time in hours and minutes.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"hour": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the hour of this time.",
																														},
																														"minute": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the minute of this time.",
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
																					"threshold": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Throttling threshold applicable in the window.",
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
												"username": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies username to access the target source.",
												},
											},
										},
									},
									"progress_monitor_path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Captures the current progress and pulse details w.r.t to either the registration or refresh.",
									},
									"refresh_error_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.",
									},
									"refresh_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.",
									},
									"registered_apps_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies information of the applications registered on this protection source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_error_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.",
												},
												"authentication_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
												},
												"host_settings_check_results": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"check_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.",
															},
															"result_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.",
															},
															"user_message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a descriptive message for failed/warning types.",
															},
														},
													},
												},
												"refresh_error_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.",
												},
											},
										},
									},
									"registration_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.",
									},
									"sfdc_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered Salesforce source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"access_token": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Token that will be used in subsequent api requests.",
												},
												"concurrent_api_requests_limit": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the maximum number of concurrent API requests allowed for salesforce.",
												},
												"consumer_key": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Consumer key from the connected app in Sfdc.",
												},
												"consumer_secret": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Consumer secret from the connected app in Sfdc.",
												},
												"daily_api_limit": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Maximum daily api limit.",
												},
												"endpoint": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Sfdc Endpoint URL.",
												},
												"endpoint_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Environment type for salesforce. 'PROD' 'SANDBOX' 'OTHER'.",
												},
												"metadata_endpoint_url": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Metadata endpoint url. All metadata requests must be made to this url.",
												},
												"refresh_token": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Token that will be used to refresh the access token.",
												},
												"soap_endpoint_url": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Soap endpoint url. All soap requests must be made to this url.",
												},
												"use_bulk_api": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "use bulk api if set to true.",
												},
											},
										},
									},
									"subnets": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"component": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Component that has reserved the subnet.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Description of the subnet.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "ID of the subnet.",
												},
												"ip": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies either an IPv6 address or an IPv4 address.",
												},
												"netmask_bits": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "netmaskBits.",
												},
												"netmask_ip4": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.",
												},
												"nfs_access": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Component that has reserved the subnet.",
												},
												"nfs_all_squash": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.",
												},
												"nfs_root_squash": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether clients from this subnet can mount as root on NFS.",
												},
												"s3_access": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
												},
												"smb_access": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
												},
												"tenant_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the unique id of the tenant.",
												},
											},
										},
									},
									"throttling_policy": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the throttling policy for a registered Protection Source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enforce_max_streams": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
												},
												"enforce_registered_source_max_backups": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
												},
												"is_enabled": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
												},
												"latency_thresholds": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"active_task_msecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
															},
															"new_task_msecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
															},
														},
													},
												},
												"max_concurrent_streams": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
												},
												"nas_source_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
															},
															"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
															},
															"max_parallel_read_write_full_percentage": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
															},
															"max_parallel_read_write_incremental_percentage": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
															},
														},
													},
												},
												"registered_source_max_concurrent_backups": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
												},
												"storage_array_snapshot_config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_max_snapshots_config_enabled": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
															},
															"is_max_space_config_enabled": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if the storage array snapshot max space config is enabled or not.",
															},
															"storage_array_snapshot_max_space_config": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_snapshot_space_percentage": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Max number of storage snapshots allowed per volume/lun.",
																		},
																	},
																},
															},
															"storage_array_snapshot_throttling_policies": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies throttling policies configured for individual volume/lun.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the volume id of the storage array snapshot config.",
																		},
																		"is_max_snapshots_config_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																		},
																		"is_max_space_config_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																		},
																		"max_snapshot_config": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"max_snapshots": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Max number of storage snapshots allowed per volume/lun.",
																					},
																				},
																			},
																		},
																		"max_space_config": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"max_snapshot_space_percentage": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Max number of storage snapshots allowed per volume/lun.",
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
									"throttling_policy_overrides": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of Throttling Policy Overrides for Datastores. Specifies a list of Throttling Policy for datastores that override the common throttling policy specified for the registered Protection Source. For datastores not in this list, common policy will still apply.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"datastore_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the Protection Source id of the Datastore.",
												},
												"datastore_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the display name of the Datastore.",
												},
												"throttling_policy": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the throttling policy for a registered Protection Source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enforce_max_streams": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
															},
															"enforce_registered_source_max_backups": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
															},
															"is_enabled": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
															},
															"latency_thresholds": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"active_task_msecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
																		},
																		"new_task_msecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
																		},
																	},
																},
															},
															"max_concurrent_streams": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
															},
															"nas_source_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
																		},
																		"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
																		},
																		"max_parallel_read_write_full_percentage": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
																		},
																		"max_parallel_read_write_incremental_percentage": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
																		},
																	},
																},
															},
															"registered_source_max_concurrent_backups": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
															},
															"storage_array_snapshot_config": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"is_max_snapshots_config_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																		},
																		"is_max_space_config_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																		},
																		"storage_array_snapshot_max_space_config": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"max_snapshot_space_percentage": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Max number of storage snapshots allowed per volume/lun.",
																					},
																				},
																			},
																		},
																		"storage_array_snapshot_throttling_policies": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies throttling policies configured for individual volume/lun.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the volume id of the storage array snapshot config.",
																					},
																					"is_max_snapshots_config_enabled": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																					},
																					"is_max_space_config_enabled": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																					},
																					"max_snapshot_config": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_snapshots": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Max number of storage snapshots allowed per volume/lun.",
																								},
																							},
																						},
																					},
																					"max_space_config": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_snapshot_space_percentage": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "Max number of storage snapshots allowed per volume/lun.",
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
									"uda_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object containing information about a registered Universal Data Adapter source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"capabilities": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"auto_log_backup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"dynamic_config": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the source supports the 'Dynamic Configuration' capability.",
															},
															"entity_support": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Indicates if source has entity capability.",
															},
															"et_log_backup": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the source supports externally triggered log backups.",
															},
															"external_disks": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Only for sources in the cloud. A temporary external disk is provisoned in the cloud and mounted on the control node selected during backup / recovery for dump-sweep workflows that need a local disk to dump data. Prereq - non-mount, AGENT_ON_RIGEL.",
															},
															"full_backup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"incr_backup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"log_backup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"multi_object_restore": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether the source supports restore of multiple objects.",
															},
															"pause_resume_backup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"post_backup_job_script": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Triggers a post backup script on all nodes.",
															},
															"post_restore_job_script": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Triggers a post restore script on all nodes.",
															},
															"pre_backup_job_script": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Make a source call before actual start backup call.",
															},
															"pre_restore_job_script": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Triggers a pre restore script on all nodes.",
															},
															"resource_throttling": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"snapfs_cert": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
												"credentials": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the object to hold username and password.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"username": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the username to access target entity.",
															},
															"password": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the password to access target entity.",
															},
														},
													},
												},
												"et_enable_log_backup_policy": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to enable cohesity policy triggered log backups along with externally triggered backups. Only applicable if etLogBackup capability is true.",
												},
												"et_enable_run_now": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if the user triggered runs are allowed along with externally triggered backups. Only applicable if etLogBackup is true.",
												},
												"host_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment type for the host.",
												},
												"hosts": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of hosts forming the Universal Data Adapter cluster.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"live_data_view": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to use a live view for data backups.",
												},
												"live_log_view": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to use a live view for log backups.",
												},
												"mount_dir": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "This field is deprecated and its value will be ignored. It was used to specify the absolute path on the host where the view would be mounted. This is now controlled by the agent configuration and is common for all the Universal Data Adapter sources. deprecated: true.",
												},
												"mount_view": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to mount a view during the source backup.",
												},
												"script_dir": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Path where various source scripts will be located.",
												},
												"source_args": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Custom arguments which will be provided to the source registration scripts. This is deprecated. Use 'sourceRegistrationArguments' instead.",
												},
												"source_registration_arguments": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies a map of custom arguments to be supplied to the source registration scripts.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"source_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Global app source type.",
												},
											},
										},
									},
									"update_last_backup_details": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if the last backup time and status should be updated for the VMs protected from the vCenter.",
									},
									"use_o_auth_for_exchange_online": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether OAuth should be used for authentication in case of  Exchange Online.",
									},
									"use_vm_bios_uuid": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if registered vCenter is using BIOS UUID to track virtual  machines.",
									},
									"user_messages": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"username": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies username to access the target source.",
									},
									"vlan_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies VLAN parameters for the restore operation.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disable_vlan": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.",
												},
												"interface_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
												},
												"vlan": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
												},
											},
										},
									},
									"warning_messages": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"root_node": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Protection Source for the root node of the Protection Source tree.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connection_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the connection id of the tenant.",
									},
									"connector_group_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the connector group id of the connector groups.",
									},
									"custom_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the user provided custom name of the Protection Source.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the environment (such as 'kVMware' or 'kSQL') where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies an id of the Protection Source.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies a name of the Protection Source.",
									},
									"parent_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies an id of the parent of the Protection Source.",
									},
									"physical_protection_source": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a Protection Source in a Physical environment.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"agents": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifiles the agents running on the Physical Protection Source and the status information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cbmr_version": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the version if Cristie BMR product is installed on the host.",
															},
															"file_cbt_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "CBT version and service state info.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"file_version": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Subcomponent version. The interpretation of the version is based on operating system.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"build_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"major_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"minor_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"revision_num": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"is_installed": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Indicates whether the cbt driver is installed.",
																		},
																		"reboot_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Indicates whether host is rebooted post VolCBT installation.",
																		},
																		"service_state": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Structure to Hold Service Status.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"state": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"host_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the host type where the agent is running. This is only set for persistent agents.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the agent's id.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the agent's name.",
															},
															"oracle_multi_node_channel_supported": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether oracle multi node multi channel is supported or not.",
															},
															"registration_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies information about a registered Source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"access_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the parameters required to establish a connection with a particular environment.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"connection_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.",
																					},
																					"connector_group_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.",
																					},
																					"endpoint": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).",
																					},
																					"environment": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																					},
																					"id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.",
																					},
																					"version": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.",
																					},
																				},
																			},
																		},
																		"allowed_ip_addresses": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"authentication_error_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.",
																		},
																		"authentication_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster.",
																		},
																		"blacklisted_ip_addresses": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "This field is deprecated. Use DeniedIpAddresses instead.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"denied_ip_addresses": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"environments": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"is_db_authenticated": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if application entity dbAuthenticated or not.",
																		},
																		"is_storage_array_snapshot_enabled": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if this source entity has enabled storage array snapshot or not.",
																		},
																		"link_vms_across_vcenter": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.",
																		},
																		"minimum_free_space_gb": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
																		},
																		"minimum_free_space_percent": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.",
																		},
																		"password": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies password of the username to access the target source.",
																		},
																		"physical_params": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"applications": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"password": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies password of the username to access the target source.",
																					},
																					"throttling_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the source side throttling configuration.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"cpu_throttling_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Throttling Configuration Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"fixed_threshold": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																											},
																											"pattern_type": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																											},
																											"throttling_windows": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"day_time_window": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the Day Time Window Parameters.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"end_time": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"day": &schema.Schema{
																																					Type:        schema.TypeString,
																																					Computed:    true,
																																					Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																				},
																																				"time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the time in hours and minutes.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"hour": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the hour of this time.",
																																							},
																																							"minute": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the minute of this time.",
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																	"start_time": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"day": &schema.Schema{
																																					Type:        schema.TypeString,
																																					Computed:    true,
																																					Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																				},
																																				"time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the time in hours and minutes.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"hour": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the hour of this time.",
																																							},
																																							"minute": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the minute of this time.",
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
																														"threshold": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Throttling threshold applicable in the window.",
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																								"network_throttling_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the Throttling Configuration Parameters.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"fixed_threshold": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.",
																											},
																											"pattern_type": &schema.Schema{
																												Type:        schema.TypeString,
																												Computed:    true,
																												Description: "Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.",
																											},
																											"throttling_windows": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"day_time_window": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies the Day Time Window Parameters.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"end_time": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"day": &schema.Schema{
																																					Type:        schema.TypeString,
																																					Computed:    true,
																																					Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																				},
																																				"time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the time in hours and minutes.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"hour": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the hour of this time.",
																																							},
																																							"minute": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the minute of this time.",
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																	"start_time": &schema.Schema{
																																		Type:        schema.TypeList,
																																		Computed:    true,
																																		Description: "Specifies the Day Time Parameters.",
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"day": &schema.Schema{
																																					Type:        schema.TypeString,
																																					Computed:    true,
																																					Description: "Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.",
																																				},
																																				"time": &schema.Schema{
																																					Type:        schema.TypeList,
																																					Computed:    true,
																																					Description: "Specifies the time in hours and minutes.",
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"hour": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the hour of this time.",
																																							},
																																							"minute": &schema.Schema{
																																								Type:        schema.TypeInt,
																																								Computed:    true,
																																								Description: "Specifies the minute of this time.",
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
																														"threshold": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Throttling threshold applicable in the window.",
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
																					"username": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies username to access the target source.",
																					},
																				},
																			},
																		},
																		"progress_monitor_path": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Captures the current progress and pulse details w.r.t to either the registration or refresh.",
																		},
																		"refresh_error_message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.",
																		},
																		"refresh_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.",
																		},
																		"registered_apps_info": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies information of the applications registered on this protection source.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"authentication_error_message": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.",
																					},
																					"authentication_status": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.",
																					},
																					"environment": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.",
																					},
																					"host_settings_check_results": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"check_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.",
																								},
																								"result_type": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.",
																								},
																								"user_message": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "Specifies a descriptive message for failed/warning types.",
																								},
																							},
																						},
																					},
																					"refresh_error_message": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.",
																					},
																				},
																			},
																		},
																		"registration_time_usecs": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.",
																		},
																		"subnets": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"component": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Component that has reserved the subnet.",
																					},
																					"description": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Description of the subnet.",
																					},
																					"id": &schema.Schema{
																						Type:        schema.TypeFloat,
																						Computed:    true,
																						Description: "ID of the subnet.",
																					},
																					"ip": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies either an IPv6 address or an IPv4 address.",
																					},
																					"netmask_bits": &schema.Schema{
																						Type:        schema.TypeFloat,
																						Computed:    true,
																						Description: "netmaskBits.",
																					},
																					"netmask_ip4": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.",
																					},
																					"nfs_access": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Component that has reserved the subnet.",
																					},
																					"nfs_all_squash": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.",
																					},
																					"nfs_root_squash": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether clients from this subnet can mount as root on NFS.",
																					},
																					"s3_access": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
																					},
																					"smb_access": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.",
																					},
																					"tenant_id": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the unique id of the tenant.",
																					},
																				},
																			},
																		},
																		"throttling_policy": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the throttling policy for a registered Protection Source.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"enforce_max_streams": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
																					},
																					"enforce_registered_source_max_backups": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
																					},
																					"is_enabled": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
																					},
																					"latency_thresholds": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"active_task_msecs": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
																								},
																								"new_task_msecs": &schema.Schema{
																									Type:        schema.TypeInt,
																									Computed:    true,
																									Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
																								},
																							},
																						},
																					},
																					"max_concurrent_streams": &schema.Schema{
																						Type:        schema.TypeFloat,
																						Computed:    true,
																						Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
																					},
																					"nas_source_params": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
																								},
																								"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
																								},
																								"max_parallel_read_write_full_percentage": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
																								},
																								"max_parallel_read_write_incremental_percentage": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
																								},
																							},
																						},
																					},
																					"registered_source_max_concurrent_backups": &schema.Schema{
																						Type:        schema.TypeFloat,
																						Computed:    true,
																						Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
																					},
																					"storage_array_snapshot_config": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies Storage Array Snapshot Configuration.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"is_max_snapshots_config_enabled": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																								},
																								"is_max_space_config_enabled": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																								},
																								"storage_array_snapshot_max_space_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies Storage Array Snapshot Max Space Config.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"max_snapshot_space_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Max number of storage snapshots allowed per volume/lun.",
																											},
																										},
																									},
																								},
																								"storage_array_snapshot_throttling_policies": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies throttling policies configured for individual volume/lun.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"id": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "Specifies the volume id of the storage array snapshot config.",
																											},
																											"is_max_snapshots_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																											},
																											"is_max_space_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																											},
																											"max_snapshot_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies Storage Array Snapshot Max Snapshots Config.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"max_snapshots": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Max number of storage snapshots allowed per volume/lun.",
																														},
																													},
																												},
																											},
																											"max_space_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies Storage Array Snapshot Max Space Config.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"max_snapshot_space_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Max number of storage snapshots allowed per volume/lun.",
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
																		"throttling_policy_overrides": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"datastore_id": &schema.Schema{
																						Type:        schema.TypeInt,
																						Computed:    true,
																						Description: "Specifies the Protection Source id of the Datastore.",
																					},
																					"datastore_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the display name of the Datastore.",
																					},
																					"throttling_policy": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Specifies the throttling policy for a registered Protection Source.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"enforce_max_streams": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.",
																								},
																								"enforce_registered_source_max_backups": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.",
																								},
																								"is_enabled": &schema.Schema{
																									Type:        schema.TypeBool,
																									Computed:    true,
																									Description: "Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.",
																								},
																								"latency_thresholds": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"active_task_msecs": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.",
																											},
																											"new_task_msecs": &schema.Schema{
																												Type:        schema.TypeInt,
																												Computed:    true,
																												Description: "If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.",
																											},
																										},
																									},
																								},
																								"max_concurrent_streams": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.",
																								},
																								"nas_source_params": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies the NAS specific source throttling parameters during source registration or during backup of the source.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"max_parallel_metadata_fetch_full_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.",
																											},
																											"max_parallel_metadata_fetch_incremental_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.",
																											},
																											"max_parallel_read_write_full_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent IO during full backup of the source.",
																											},
																											"max_parallel_read_write_incremental_percentage": &schema.Schema{
																												Type:        schema.TypeFloat,
																												Computed:    true,
																												Description: "Specifies the percentage value of maximum concurrent IO during incremental backup of the source.",
																											},
																										},
																									},
																								},
																								"registered_source_max_concurrent_backups": &schema.Schema{
																									Type:        schema.TypeFloat,
																									Computed:    true,
																									Description: "Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.",
																								},
																								"storage_array_snapshot_config": &schema.Schema{
																									Type:        schema.TypeList,
																									Computed:    true,
																									Description: "Specifies Storage Array Snapshot Configuration.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"is_max_snapshots_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																											},
																											"is_max_space_config_enabled": &schema.Schema{
																												Type:        schema.TypeBool,
																												Computed:    true,
																												Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																											},
																											"storage_array_snapshot_max_space_config": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies Storage Array Snapshot Max Space Config.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"max_snapshot_space_percentage": &schema.Schema{
																															Type:        schema.TypeFloat,
																															Computed:    true,
																															Description: "Max number of storage snapshots allowed per volume/lun.",
																														},
																													},
																												},
																											},
																											"storage_array_snapshot_throttling_policies": &schema.Schema{
																												Type:        schema.TypeList,
																												Computed:    true,
																												Description: "Specifies throttling policies configured for individual volume/lun.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"id": &schema.Schema{
																															Type:        schema.TypeInt,
																															Computed:    true,
																															Description: "Specifies the volume id of the storage array snapshot config.",
																														},
																														"is_max_snapshots_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max snapshots config is enabled or not.",
																														},
																														"is_max_space_config_enabled": &schema.Schema{
																															Type:        schema.TypeBool,
																															Computed:    true,
																															Description: "Specifies if the storage array snapshot max space config is enabled or not.",
																														},
																														"max_snapshot_config": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies Storage Array Snapshot Max Snapshots Config.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"max_snapshots": &schema.Schema{
																																		Type:        schema.TypeFloat,
																																		Computed:    true,
																																		Description: "Max number of storage snapshots allowed per volume/lun.",
																																	},
																																},
																															},
																														},
																														"max_space_config": &schema.Schema{
																															Type:        schema.TypeList,
																															Computed:    true,
																															Description: "Specifies Storage Array Snapshot Max Space Config.",
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"max_snapshot_space_percentage": &schema.Schema{
																																		Type:        schema.TypeFloat,
																																		Computed:    true,
																																		Description: "Max number of storage snapshots allowed per volume/lun.",
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
																		"use_o_auth_for_exchange_online": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies whether OAuth should be used for authentication in case of Exchange Online.",
																		},
																		"use_vm_bios_uuid": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Specifies if registered vCenter is using BIOS UUID to track virtual machines.",
																		},
																		"user_messages": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"username": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies username to access the target source.",
																		},
																		"vlan_params": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies the VLAN configuration for Recovery.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"vlan": &schema.Schema{
																						Type:        schema.TypeFloat,
																						Computed:    true,
																						Description: "Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																					},
																					"disable_vlan": &schema.Schema{
																						Type:        schema.TypeBool,
																						Computed:    true,
																						Description: "Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.",
																					},
																					"interface_name": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
																					},
																				},
																			},
																		},
																		"warning_messages": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
															"source_side_dedup_enabled": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether source side dedup is enabled or not.",
															},
															"status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the agent status. Specifies the status of the agent running on a physical source.",
															},
															"status_message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies additional details about the agent status.",
															},
															"upgradability": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the upgradability of the agent running on the physical server. Specifies the upgradability of the agent running on the physical server.",
															},
															"upgrade_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the status of the upgrade of the agent on a physical server. Specifies the status of the upgrade of the agent on a physical server.",
															},
															"upgrade_status_message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade.",
															},
															"version": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the version of the Agent software.",
															},
															"vol_cbt_info": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "CBT version and service state info.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"file_version": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Subcomponent version. The interpretation of the version is based on operating system.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"build_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"major_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"minor_ver": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																					"revision_num": &schema.Schema{
																						Type:     schema.TypeFloat,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"is_installed": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Indicates whether the cbt driver is installed.",
																		},
																		"reboot_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Indicates whether host is rebooted post VolCBT installation.",
																		},
																		"service_state": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Structure to Hold Service Status.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"state": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
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
												"cluster_source_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of cluster resource this source represents.",
												},
												"host_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the hostname.",
												},
												"host_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment type for the host.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an id for an object that is unique across Cohesity Clusters. The id is composite of all the ids listed below.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cluster_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the Cohesity Cluster id where the object was created.",
															},
															"cluster_incarnation_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies an id for the Cohesity Cluster that is generated when a Cohesity Cluster is initially created.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies a unique id assigned to an object (such as a Job) by the Cohesity Cluster.",
															},
														},
													},
												},
												"is_proxy_host": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if the physical host is a proxy host.",
												},
												"memory_size_bytes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total memory on the host in bytes.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a human readable name of the Protection Source.",
												},
												"networking_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the struct containing information about network addresses configured on the given box. This is needed for dealing with Windows/Oracle Cluster resources that we discover and protect automatically.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"resource_vec": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of resources on the system that are accessible by an IP address.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"endpoints": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The endpoints by which the resource is accessible.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"fqdn": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The Fully Qualified Domain Name.",
																					},
																					"ipv4_addr": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The IPv4 address.",
																					},
																					"ipv6_addr": &schema.Schema{
																						Type:        schema.TypeString,
																						Computed:    true,
																						Description: "The IPv6 address.",
																					},
																				},
																			},
																		},
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The type of the resource.",
																		},
																	},
																},
															},
														},
													},
												},
												"num_processors": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of processors on the host.",
												},
												"os_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a human readable name of the OS of the Protection Source.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of managed Object in a Physical Protection Source. 'kGroup' indicates the EH container.",
												},
												"vcs_version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies cluster version for VCS host.",
												},
												"volumes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Array of Physical Volumes. Specifies the volumes available on the physical host. These fields are populated only for the kPhysicalHost type.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"device_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the path to the device that hosts the volume locally.",
															},
															"guid": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies an id for the Physical Volume.",
															},
															"is_boot_volume": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the volume is boot volume.",
															},
															"is_extended_attributes_supported": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether this volume supports extended attributes (like ACLs) when performing file backups.",
															},
															"is_protected": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies if a volume is protected by a Job.",
															},
															"is_shared_volume": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether the volume is shared volume.",
															},
															"label": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a volume label that can be used for displaying additional identifying information about a volume.",
															},
															"logical_size_bytes": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the logical size of the volume in bytes that is not reduced by change-block tracking, compression and deduplication.",
															},
															"mount_points": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies the mount points where the volume is mounted, for example- 'C:', '/mnt/foo' etc.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"mount_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies mount type of volume e.g. nfs, autofs, ext4 etc.",
															},
															"network_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the full path to connect to the network attached volume. For example, (IP or hostname):/path/to/share for NFS volumes).",
															},
															"used_size_bytes": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the size used by the volume in bytes.",
															},
														},
													},
												},
												"vsswriters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_writer_excluded": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If true, the writer will be excluded by default.",
															},
															"writer_name": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies the name of the writer.",
															},
														},
													},
												},
											},
										},
									},
									"kubernetes_protection_source": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a Protection Source in Kubernetes environment.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"datamover_image_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the location of Datamover image in private registry.",
												},
												"datamover_service_type": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies Type of service to be deployed for communication with DataMover pods. Currently, LoadBalancer and NodePort are supported. [default = kNodePort].",
												},
												"datamover_upgradability": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies if the deployed Datamover image needs to be upgraded for this kubernetes entity.",
												},
												"default_vlan_params": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies VLAN parameters for the restore operation.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"disable_vlan": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.",
															},
															"interface_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
															},
															"vlan": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.",
															},
														},
													},
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies an optional description of the object.",
												},
												"distribution": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of the entity in a Kubernetes environment. Determines the K8s distribution. kIKS, kROKS.",
												},
												"init_container_image_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the location of the image for init containers.",
												},
												"label_attributes": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the list of label attributes of this source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the Cohesity id of the K8s label.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the appended key and value of the K8s label.",
															},
															"uuid": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies Kubernetes Unique Identifier (UUID) of the K8s label.",
															},
														},
													},
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies a unique name of the Protection Source.",
												},
												"priority_class_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the pritority class name during registration.",
												},
												"resource_annotation_list": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies resource Annotations information provided during registration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Key for label.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Value for label.",
															},
														},
													},
												},
												"resource_label_list": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies resource labels information provided during registration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Key for label.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Value for label.",
															},
														},
													},
												},
												"san_field": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the SAN field for agent certificate.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"service_annotations": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"storage_class": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies storage class information of source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies name of storage class.",
															},
															"provisioner": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "specifies provisioner of storage class.",
															},
														},
													},
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of the entity in a Kubernetes environment. Specifies the type of a Kubernetes Protection Source. 'kCluster' indicates a Kubernetes Cluster. 'kNamespace' indicates a namespace in a Kubernetes Cluster. 'kService' indicates a service running on a Kubernetes Cluster.",
												},
												"uuid": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the UUID of the object.",
												},
												"velero_aws_plugin_image_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the location of Velero AWS plugin image in private registry.",
												},
												"velero_image_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the location of Velero image in private registry.",
												},
												"velero_openshift_plugin_image_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the location of the image for openshift plugin container.",
												},
												"velero_upgradability": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies if the deployed Velero image needs to be upgraded for this kubernetes entity.",
												},
												"vlan_info_vec": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies VLAN information provided during registration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service_annotations": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the service annotation key value.",
																		},
																		"value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Specifies the service annotation value.",
																		},
																	},
																},
															},
															"vlan_params": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Specifies VLAN params associated with the backup/restore operation.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"disable_vlan": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.",
																		},
																		"interface_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.",
																		},
																		"vlan_id": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.",
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
									"sql_protection_source": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an Object representing one SQL Server instance or database.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_available_for_vss_backup": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether the database is marked as available for backup according to the SQL Server VSS writer. This may be false if either the state of the databases is not online, or if the VSS writer is not online. This field is set only for type 'kDatabase'.",
												},
												"created_timestamp": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the time when the database was created. It is displayed in the timezone of the SQL server on which this database is running.",
												},
												"database_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the database name of the SQL Protection Source, if the type is database.",
												},
												"db_aag_entity_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the AAG entity id if the database is part of an AAG. This field is set only for type 'kDatabase'.",
												},
												"db_aag_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the AAG if the database is part of an AAG. This field is set only for type 'kDatabase'.",
												},
												"db_compatibility_level": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the versions of SQL server that the database is compatible with.",
												},
												"db_file_groups": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the information about the set of file groups for this db on the host. This is only set if the type is kDatabase.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"db_files": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the last known information about the set of database files on the host. This field is set only for type 'kDatabase'.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"file_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. 'kRows' refers to a data file 'kLog' refers to a log file 'kFileStream' refers to a directory containing FILESTREAM data 'kNotSupportedType' is for information purposes only. Not supported. 'kFullText' refers to a full-text catalog.",
															},
															"full_path": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the full path of the database file on the SQL host machine.",
															},
															"size_bytes": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies the last known size of the database file.",
															},
														},
													},
												},
												"db_owner_username": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the database owner.",
												},
												"default_database_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the default path for data files for DBs in an instance.",
												},
												"default_log_location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the default path for log files for DBs in an instance.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies a unique id for a SQL Protection Source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"created_date_msecs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies a unique identifier generated from the date the database is created or renamed. Cohesity uses this identifier in combination with the databaseId to uniquely identify a database.",
															},
															"database_id": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Specifies a unique id of the database but only for the life of the database. SQL Server may reuse database ids. Cohesity uses the createDateMsecs in combination with this databaseId to uniquely identify a database.",
															},
															"instance_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies unique id for the SQL Server instance. This id does not change during the life of the instance.",
															},
														},
													},
												},
												"is_encrypted": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies whether the database is TDE enabled.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the instance name of the SQL Protection Source.",
												},
												"owner_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the id of the container VM for the SQL Protection Source.",
												},
												"recovery_model": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the Recovery Model for the database in SQL environment. Only meaningful for the 'kDatabase' SQL Protection Source. Specifies the Recovery Model set for the Microsoft SQL Server. 'kSimpleRecoveryModel' indicates the Simple SQL Recovery Model which does not utilize log backups. 'kFullRecoveryModel' indicates the Full SQL Recovery Model which requires log backups and allows recovery to a single point in time. 'kBulkLoggedRecoveryModel' indicates the Bulk Logged SQL Recovery Model which requires log backups and allows high-performance bulk copy operations.",
												},
												"sql_server_db_state": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The state of the database as returned by SQL Server. Indicates the state of the database. The values correspond to the 'state' field in the system table sys.databases. See https://goo.gl/P66XqM. 'kOnline' indicates that database is in online state. 'kRestoring' indicates that database is in restore state. 'kRecovering' indicates that database is in recovery state. 'kRecoveryPending' indicates that database recovery is in pending state. 'kSuspect' indicates that primary filegroup is suspect and may be damaged. 'kEmergency' indicates that manually forced emergency state. 'kOffline' indicates that database is in offline state. 'kCopying' indicates that database is in copying state. 'kOfflineSecondary' indicates that secondary database is in offline state.",
												},
												"sql_server_instance_version": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the Server Instance Version.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"build": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the build.",
															},
															"major_version": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the major version.",
															},
															"minor_version": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the minor version.",
															},
															"revision": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the revision.",
															},
															"version_string": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "Specifies the version string.",
															},
														},
													},
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of the managed Object in a SQL Protection Source. Examples of SQL Objects include 'kInstance' and 'kDatabase'. 'kInstance' indicates that SQL server instance is being protected. 'kDatabase' indicates that SQL server database is being protected. 'kAAG' indicates that SQL AAG (AlwaysOn Availability Group) is being protected. 'kAAGRootContainer' indicates that SQL AAG's root container is being protected. 'kRootContainer' indicates root container for SQL sources.",
												},
											},
										},
									},
								},
							},
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the stats of protection for a Protection Source Tree.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are protected under the given entity.",
									},
									"protected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of the protected objects under the given entity.",
									},
									"unprotected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are not protected under the given entity.",
									},
									"unprotected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of the unprotected objects under the given entity.",
									},
								},
							},
						},
						"stats_by_env": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the breakdown of the stats of protection by environment. overrideDescription: true.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of environment of the source object like kSQL etc.  Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders ProtectionSource environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
									},
									"kubernetes_distribution_stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the breakdown of the kubernetes clusters by distribution type.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distribution": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the type of Kuberentes distribution Determines the K8s distribution. kIKS, kROKS.",
												},
												"protected_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of objects that are protected for that distribution.",
												},
												"protected_size": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total size of objects that are protected for that distribution.",
												},
												"total_registered_clusters": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of registered clusters for that distribution.",
												},
												"unprotected_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of objects that are not protected for that distribution.",
												},
												"unprotected_size": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the total size of objects that are not protected for that distribution.",
												},
											},
										},
									},
									"protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are protected under the given entity.",
									},
									"protected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of the protected objects under the given entity.",
									},
									"unprotected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are not protected under the given entity.",
									},
									"unprotected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of the unprotected objects under the given entity.",
									},
								},
							},
						},
						"total_downtiered_size_in_bytes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total bytes downtiered from the source so far.",
						},
						"total_uptiered_size_in_bytes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total bytes uptiered to the source so far.",
						},
					},
				},
			},
			"stats": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the sum of all the stats of protection of Protection Sources and views selected by the query parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protected_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of objects that are protected under the given entity.",
						},
						"protected_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total size of the protected objects under the given entity.",
						},
						"unprotected_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of objects that are not protected under the given entity.",
						},
						"unprotected_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total size of the unprotected objects under the given entity.",
						},
					},
				},
			},
			"stats_by_env": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the breakdown of the stats by environment overrideDescription: true.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of environment of the source object like kSQL etc.  Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders ProtectionSource environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.",
						},
						"kubernetes_distribution_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the breakdown of the kubernetes clusters by distribution type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distribution": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of Kuberentes distribution Determines the K8s distribution. kIKS, kROKS.",
									},
									"protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are protected for that distribution.",
									},
									"protected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of objects that are protected for that distribution.",
									},
									"total_registered_clusters": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of registered clusters for that distribution.",
									},
									"unprotected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects that are not protected for that distribution.",
									},
									"unprotected_size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total size of objects that are not protected for that distribution.",
									},
								},
							},
						},
						"protected_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of objects that are protected under the given entity.",
						},
						"protected_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total size of the protected objects under the given entity.",
						},
						"unprotected_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of objects that are not protected under the given entity.",
						},
						"unprotected_size": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total size of the unprotected objects under the given entity.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryRegistrationInfoRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_registration_info", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listProtectionSourcesRegistrationInfoOptions := &backuprecoveryv1.ListProtectionSourcesRegistrationInfoOptions{}

	listProtectionSourcesRegistrationInfoOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	if _, ok := d.GetOk("environments"); ok {
		var environments []string
		for _, v := range d.Get("environments").([]interface{}) {
			environmentsItem := v.(string)
			environments = append(environments, environmentsItem)
		}
		listProtectionSourcesRegistrationInfoOptions.SetEnvironments(environments)
	}
	if _, ok := d.GetOk("ids"); ok {
		var ids []int64
		for _, v := range d.Get("ids").([]interface{}) {
			idsItem := int64(v.(int))
			ids = append(ids, idsItem)
		}
		listProtectionSourcesRegistrationInfoOptions.SetIds(ids)
	}
	if _, ok := d.GetOk("include_entity_permission_info"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetIncludeEntityPermissionInfo(d.Get("include_entity_permission_info").(bool))
	}
	if _, ok := d.GetOk("sids"); ok {
		var sids []string
		for _, v := range d.Get("sids").([]interface{}) {
			sidsItem := v.(string)
			sids = append(sids, sidsItem)
		}
		listProtectionSourcesRegistrationInfoOptions.SetSids(sids)
	}
	if _, ok := d.GetOk("include_source_credentials"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetIncludeSourceCredentials(d.Get("include_source_credentials").(bool))
	}
	if _, ok := d.GetOk("encryption_key"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetEncryptionKey(d.Get("encryption_key").(string))
	}
	if _, ok := d.GetOk("include_applications_tree_info"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetIncludeApplicationsTreeInfo(d.Get("include_applications_tree_info").(bool))
	}
	if _, ok := d.GetOk("prune_non_critical_info"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetPruneNonCriticalInfo(d.Get("prune_non_critical_info").(bool))
	}
	if _, ok := d.GetOk("request_initiator_type"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("use_cached_data"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetUseCachedData(d.Get("use_cached_data").(bool))
	}
	if _, ok := d.GetOk("include_external_metadata"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetIncludeExternalMetadata(d.Get("include_external_metadata").(bool))
	}
	if _, ok := d.GetOk("maintenance_status"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetMaintenanceStatus(d.Get("maintenance_status").(string))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		listProtectionSourcesRegistrationInfoOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("all_under_hierarchy"); ok {
		listProtectionSourcesRegistrationInfoOptions.SetAllUnderHierarchy(d.Get("all_under_hierarchy").(bool))
	}

	listProtectionSourcesRegistrationInfoResponse, _, err := backupRecoveryClient.ListProtectionSourcesRegistrationInfoWithContext(context, listProtectionSourcesRegistrationInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListProtectionSourcesRegistrationInfoWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_registration_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryRegistrationInfoID(d))

	if !core.IsNil(listProtectionSourcesRegistrationInfoResponse.RootNodes) {
		rootNodes := []map[string]interface{}{}
		for _, rootNodesItem := range listProtectionSourcesRegistrationInfoResponse.RootNodes {
			rootNodesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoToMap(&rootNodesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_registration_info", "read", "root_nodes-to-map").GetDiag()
			}
			rootNodes = append(rootNodes, rootNodesItemMap)
		}
		if err = d.Set("root_nodes", rootNodes); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting root_nodes: %s", err), "(Data) ibm_backup_recovery_registration_info", "read", "set-root_nodes").GetDiag()
		}
	}

	if !core.IsNil(listProtectionSourcesRegistrationInfoResponse.Stats) {
		stats := []map[string]interface{}{}
		statsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoListProtectionSourcesRegistrationInfoResponseStatsToMap(listProtectionSourcesRegistrationInfoResponse.Stats)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_registration_info", "read", "stats-to-map").GetDiag()
		}
		stats = append(stats, statsMap)
		if err = d.Set("stats", stats); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting stats: %s", err), "(Data) ibm_backup_recovery_registration_info", "read", "set-stats").GetDiag()
		}
	}

	if !core.IsNil(listProtectionSourcesRegistrationInfoResponse.StatsByEnv) {
		statsByEnv := []map[string]interface{}{}
		for _, statsByEnvItem := range listProtectionSourcesRegistrationInfoResponse.StatsByEnv {
			statsByEnvItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryByEnvToMap(&statsByEnvItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_registration_info", "read", "stats_by_env-to-map").GetDiag()
			}
			statsByEnv = append(statsByEnv, statsByEnvItemMap)
		}
		if err = d.Set("stats_by_env", statsByEnv); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting stats_by_env: %s", err), "(Data) ibm_backup_recovery_registration_info", "read", "set-stats_by_env").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryRegistrationInfoID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryRegistrationInfoID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoToMap(model *backuprecoveryv1.ProtectionSourceTreeInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Applications != nil {
		applications := []map[string]interface{}{}
		for _, applicationsItem := range model.Applications {
			applicationsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoApplicationInfoToMap(&applicationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			applications = append(applications, applicationsItemMap)
		}
		modelMap["applications"] = applications
	}
	if model.EntityPermissionInfo != nil {
		entityPermissionInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoEntityPermissionInformationToMap(model.EntityPermissionInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["entity_permission_info"] = []map[string]interface{}{entityPermissionInfoMap}
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.MaintenanceModeConfig != nil {
		maintenanceModeConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoMaintenanceModeConfigToMap(model.MaintenanceModeConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["maintenance_mode_config"] = []map[string]interface{}{maintenanceModeConfigMap}
	}
	if model.RegistrationInfo != nil {
		registrationInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceInfoToMap(model.RegistrationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["registration_info"] = []map[string]interface{}{registrationInfoMap}
	}
	if model.RootNode != nil {
		rootNodeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceNodeToMap(model.RootNode)
		if err != nil {
			return modelMap, err
		}
		modelMap["root_node"] = []map[string]interface{}{rootNodeMap}
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	if model.StatsByEnv != nil {
		statsByEnv := []map[string]interface{}{}
		for _, statsByEnvItem := range model.StatsByEnv {
			statsByEnvItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryByEnvToMap(&statsByEnvItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statsByEnv = append(statsByEnv, statsByEnvItemMap)
		}
		modelMap["stats_by_env"] = statsByEnv
	}
	if model.TotalDowntieredSizeInBytes != nil {
		modelMap["total_downtiered_size_in_bytes"] = flex.IntValue(model.TotalDowntieredSizeInBytes)
	}
	if model.TotalUptieredSizeInBytes != nil {
		modelMap["total_uptiered_size_in_bytes"] = flex.IntValue(model.TotalUptieredSizeInBytes)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoApplicationInfoToMap(model *backuprecoveryv1.ApplicationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ApplicationTreeInfo != nil {
		applicationTreeInfo := []map[string]interface{}{}
		for _, applicationTreeInfoItem := range model.ApplicationTreeInfo {
			applicationTreeInfoItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceNodeToMap(&applicationTreeInfoItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			applicationTreeInfo = append(applicationTreeInfo, applicationTreeInfoItemMap)
		}
		modelMap["application_tree_info"] = applicationTreeInfo
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceNodeToMap(model *backuprecoveryv1.ProtectionSourceNode) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConnectionID != nil {
		modelMap["connection_id"] = flex.IntValue(model.ConnectionID)
	}
	if model.ConnectorGroupID != nil {
		modelMap["connector_group_id"] = flex.IntValue(model.ConnectorGroupID)
	}
	if model.CustomName != nil {
		modelMap["custom_name"] = *model.CustomName
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ParentID != nil {
		modelMap["parent_id"] = flex.IntValue(model.ParentID)
	}
	if model.PhysicalProtectionSource != nil {
		physicalProtectionSourceMap, err := DataSourceIbmBackupRecoveryRegistrationInfoPhysicalProtectionSourceToMap(model.PhysicalProtectionSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_protection_source"] = []map[string]interface{}{physicalProtectionSourceMap}
	}
	if model.KubernetesProtectionSource != nil {
		kubernetesProtectionSourceMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKubernetesProtectionSourceToMap(model.KubernetesProtectionSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["kubernetes_protection_source"] = []map[string]interface{}{kubernetesProtectionSourceMap}
	}
	if model.SqlProtectionSource != nil {
		sqlProtectionSourceMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSqlProtectionSourceToMap(model.SqlProtectionSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["sql_protection_source"] = []map[string]interface{}{sqlProtectionSourceMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoPhysicalProtectionSourceToMap(model *backuprecoveryv1.PhysicalProtectionSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Agents != nil {
		agents := []map[string]interface{}{}
		for _, agentsItem := range model.Agents {
			agentsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoAgentInformationToMap(&agentsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			agents = append(agents, agentsItemMap)
		}
		modelMap["agents"] = agents
	}
	if model.ClusterSourceType != nil {
		modelMap["cluster_source_type"] = *model.ClusterSourceType
	}
	if model.HostName != nil {
		modelMap["host_name"] = *model.HostName
	}
	if model.HostType != nil {
		modelMap["host_type"] = *model.HostType
	}
	if model.ID != nil {
		idMap, err := DataSourceIbmBackupRecoveryRegistrationInfoUniqueGlobalIDToMap(model.ID)
		if err != nil {
			return modelMap, err
		}
		modelMap["id"] = []map[string]interface{}{idMap}
	}
	if model.IsProxyHost != nil {
		modelMap["is_proxy_host"] = *model.IsProxyHost
	}
	if model.MemorySizeBytes != nil {
		modelMap["memory_size_bytes"] = flex.IntValue(model.MemorySizeBytes)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.NetworkingInfo != nil {
		networkingInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoNetworkingInformationToMap(model.NetworkingInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["networking_info"] = []map[string]interface{}{networkingInfoMap}
	}
	if model.NumProcessors != nil {
		modelMap["num_processors"] = flex.IntValue(model.NumProcessors)
	}
	if model.OsName != nil {
		modelMap["os_name"] = *model.OsName
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.VcsVersion != nil {
		modelMap["vcs_version"] = *model.VcsVersion
	}
	if model.Volumes != nil {
		volumes := []map[string]interface{}{}
		for _, volumesItem := range model.Volumes {
			volumesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoPhysicalVolumeToMap(&volumesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			volumes = append(volumes, volumesItemMap)
		}
		modelMap["volumes"] = volumes
	}
	if model.Vsswriters != nil {
		vsswriters := []map[string]interface{}{}
		for _, vsswritersItem := range model.Vsswriters {
			vsswritersItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoVssWritersToMap(&vsswritersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			vsswriters = append(vsswriters, vsswritersItemMap)
		}
		modelMap["vsswriters"] = vsswriters
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoAgentInformationToMap(model *backuprecoveryv1.AgentInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CbmrVersion != nil {
		modelMap["cbmr_version"] = *model.CbmrVersion
	}
	if model.FileCbtInfo != nil {
		fileCbtInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCbtInfoToMap(model.FileCbtInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["file_cbt_info"] = []map[string]interface{}{fileCbtInfoMap}
	}
	if model.HostType != nil {
		modelMap["host_type"] = *model.HostType
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.OracleMultiNodeChannelSupported != nil {
		modelMap["oracle_multi_node_channel_supported"] = *model.OracleMultiNodeChannelSupported
	}
	if model.RegistrationInfo != nil {
		registrationInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoAgentRegistrationInfoToMap(model.RegistrationInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["registration_info"] = []map[string]interface{}{registrationInfoMap}
	}
	if model.SourceSideDedupEnabled != nil {
		modelMap["source_side_dedup_enabled"] = *model.SourceSideDedupEnabled
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.Upgradability != nil {
		modelMap["upgradability"] = *model.Upgradability
	}
	if model.UpgradeStatus != nil {
		modelMap["upgrade_status"] = *model.UpgradeStatus
	}
	if model.UpgradeStatusMessage != nil {
		modelMap["upgrade_status_message"] = *model.UpgradeStatusMessage
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	if model.VolCbtInfo != nil {
		volCbtInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCbtInfoToMap(model.VolCbtInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["vol_cbt_info"] = []map[string]interface{}{volCbtInfoMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCbtInfoToMap(model *backuprecoveryv1.CbtInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FileVersion != nil {
		fileVersionMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCbtFileVersionToMap(model.FileVersion)
		if err != nil {
			return modelMap, err
		}
		modelMap["file_version"] = []map[string]interface{}{fileVersionMap}
	}
	if model.IsInstalled != nil {
		modelMap["is_installed"] = *model.IsInstalled
	}
	if model.RebootStatus != nil {
		modelMap["reboot_status"] = *model.RebootStatus
	}
	if model.ServiceState != nil {
		serviceStateMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCbtServiceStateToMap(model.ServiceState)
		if err != nil {
			return modelMap, err
		}
		modelMap["service_state"] = []map[string]interface{}{serviceStateMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCbtFileVersionToMap(model *backuprecoveryv1.CbtFileVersion) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BuildVer != nil {
		modelMap["build_ver"] = *model.BuildVer
	}
	if model.MajorVer != nil {
		modelMap["major_ver"] = *model.MajorVer
	}
	if model.MinorVer != nil {
		modelMap["minor_ver"] = *model.MinorVer
	}
	if model.RevisionNum != nil {
		modelMap["revision_num"] = *model.RevisionNum
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCbtServiceStateToMap(model *backuprecoveryv1.CbtServiceState) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	modelMap["state"] = *model.State
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoAgentRegistrationInfoToMap(model *backuprecoveryv1.AgentRegistrationInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccessInfo != nil {
		accessInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoAgentAccessInfoToMap(model.AccessInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["access_info"] = []map[string]interface{}{accessInfoMap}
	}
	if model.AllowedIpAddresses != nil {
		modelMap["allowed_ip_addresses"] = model.AllowedIpAddresses
	}
	if model.AuthenticationErrorMessage != nil {
		modelMap["authentication_error_message"] = *model.AuthenticationErrorMessage
	}
	if model.AuthenticationStatus != nil {
		modelMap["authentication_status"] = *model.AuthenticationStatus
	}
	if model.BlacklistedIpAddresses != nil {
		modelMap["blacklisted_ip_addresses"] = model.BlacklistedIpAddresses
	}
	if model.DeniedIpAddresses != nil {
		modelMap["denied_ip_addresses"] = model.DeniedIpAddresses
	}
	if model.Environments != nil {
		modelMap["environments"] = model.Environments
	}
	if model.IsDbAuthenticated != nil {
		modelMap["is_db_authenticated"] = *model.IsDbAuthenticated
	}
	if model.IsStorageArraySnapshotEnabled != nil {
		modelMap["is_storage_array_snapshot_enabled"] = *model.IsStorageArraySnapshotEnabled
	}
	if model.LinkVmsAcrossVcenter != nil {
		modelMap["link_vms_across_vcenter"] = *model.LinkVmsAcrossVcenter
	}
	if model.MinimumFreeSpaceGB != nil {
		modelMap["minimum_free_space_gb"] = flex.IntValue(model.MinimumFreeSpaceGB)
	}
	if model.MinimumFreeSpacePercent != nil {
		modelMap["minimum_free_space_percent"] = flex.IntValue(model.MinimumFreeSpacePercent)
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoAgentPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.ProgressMonitorPath != nil {
		modelMap["progress_monitor_path"] = *model.ProgressMonitorPath
	}
	if model.RefreshErrorMessage != nil {
		modelMap["refresh_error_message"] = *model.RefreshErrorMessage
	}
	if model.RefreshTimeUsecs != nil {
		modelMap["refresh_time_usecs"] = flex.IntValue(model.RefreshTimeUsecs)
	}
	if model.RegisteredAppsInfo != nil {
		registeredAppsInfo := []map[string]interface{}{}
		for _, registeredAppsInfoItem := range model.RegisteredAppsInfo {
			registeredAppsInfoItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoRegisteredAppInfoToMap(&registeredAppsInfoItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			registeredAppsInfo = append(registeredAppsInfo, registeredAppsInfoItemMap)
		}
		modelMap["registered_apps_info"] = registeredAppsInfo
	}
	if model.RegistrationTimeUsecs != nil {
		modelMap["registration_time_usecs"] = flex.IntValue(model.RegistrationTimeUsecs)
	}
	if model.Subnets != nil {
		subnets := []map[string]interface{}{}
		for _, subnetsItem := range model.Subnets {
			subnetsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSubnetToMap(&subnetsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			subnets = append(subnets, subnetsItemMap)
		}
		modelMap["subnets"] = subnets
	}
	if model.ThrottlingPolicy != nil {
		throttlingPolicyMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyToMap(model.ThrottlingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_policy"] = []map[string]interface{}{throttlingPolicyMap}
	}
	if model.ThrottlingPolicyOverrides != nil {
		throttlingPolicyOverrides := []map[string]interface{}{}
		for _, throttlingPolicyOverridesItem := range model.ThrottlingPolicyOverrides {
			throttlingPolicyOverridesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverridesToMap(&throttlingPolicyOverridesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			throttlingPolicyOverrides = append(throttlingPolicyOverrides, throttlingPolicyOverridesItemMap)
		}
		modelMap["throttling_policy_overrides"] = throttlingPolicyOverrides
	}
	if model.UseOAuthForExchangeOnline != nil {
		modelMap["use_o_auth_for_exchange_online"] = *model.UseOAuthForExchangeOnline
	}
	if model.UseVmBiosUUID != nil {
		modelMap["use_vm_bios_uuid"] = *model.UseVmBiosUUID
	}
	if model.UserMessages != nil {
		modelMap["user_messages"] = model.UserMessages
	}
	if model.Username != nil {
		modelMap["username"] = *model.Username
	}
	if model.VlanParams != nil {
		vlanParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceVlanConfigToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	if model.WarningMessages != nil {
		modelMap["warning_messages"] = model.WarningMessages
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoAgentAccessInfoToMap(model *backuprecoveryv1.AgentAccessInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConnectionID != nil {
		modelMap["connection_id"] = flex.IntValue(model.ConnectionID)
	}
	if model.ConnectorGroupID != nil {
		modelMap["connector_group_id"] = flex.IntValue(model.ConnectorGroupID)
	}
	if model.Endpoint != nil {
		modelMap["endpoint"] = *model.Endpoint
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoAgentPhysicalParamsToMap(model *backuprecoveryv1.AgentPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Applications != nil {
		modelMap["applications"] = model.Applications
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.ThrottlingConfig != nil {
		throttlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigToMap(model.ThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_config"] = []map[string]interface{}{throttlingConfigMap}
	}
	if model.Username != nil {
		modelMap["username"] = *model.Username
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigToMap(model *backuprecoveryv1.ThrottlingConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CpuThrottlingConfig != nil {
		cpuThrottlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationParamsToMap(model.CpuThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["cpu_throttling_config"] = []map[string]interface{}{cpuThrottlingConfigMap}
	}
	if model.NetworkThrottlingConfig != nil {
		networkThrottlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationParamsToMap(model.NetworkThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["network_throttling_config"] = []map[string]interface{}{networkThrottlingConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationParamsToMap(model *backuprecoveryv1.ThrottlingConfigurationParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FixedThreshold != nil {
		modelMap["fixed_threshold"] = flex.IntValue(model.FixedThreshold)
	}
	if model.PatternType != nil {
		modelMap["pattern_type"] = *model.PatternType
	}
	if model.ThrottlingWindows != nil {
		throttlingWindows := []map[string]interface{}{}
		for _, throttlingWindowsItem := range model.ThrottlingWindows {
			throttlingWindowsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingWindowToMap(&throttlingWindowsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			throttlingWindows = append(throttlingWindows, throttlingWindowsItemMap)
		}
		modelMap["throttling_windows"] = throttlingWindows
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingWindowToMap(model *backuprecoveryv1.ThrottlingWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DayTimeWindow != nil {
		dayTimeWindowMap, err := DataSourceIbmBackupRecoveryRegistrationInfoDayTimeWindowToMap(model.DayTimeWindow)
		if err != nil {
			return modelMap, err
		}
		modelMap["day_time_window"] = []map[string]interface{}{dayTimeWindowMap}
	}
	if model.Threshold != nil {
		modelMap["threshold"] = flex.IntValue(model.Threshold)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoDayTimeWindowToMap(model *backuprecoveryv1.DayTimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndTime != nil {
		endTimeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoDayTimeParamsToMap(model.EndTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	}
	if model.StartTime != nil {
		startTimeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoDayTimeParamsToMap(model.StartTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoDayTimeParamsToMap(model *backuprecoveryv1.DayTimeParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Day != nil {
		modelMap["day"] = *model.Day
	}
	if model.Time != nil {
		timeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(model.Time)
		if err != nil {
			return modelMap, err
		}
		modelMap["time"] = []map[string]interface{}{timeMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(model *backuprecoveryv1.Time) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hour != nil {
		modelMap["hour"] = flex.IntValue(model.Hour)
	}
	if model.Minute != nil {
		modelMap["minute"] = flex.IntValue(model.Minute)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoRegisteredAppInfoToMap(model *backuprecoveryv1.RegisteredAppInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AuthenticationErrorMessage != nil {
		modelMap["authentication_error_message"] = *model.AuthenticationErrorMessage
	}
	if model.AuthenticationStatus != nil {
		modelMap["authentication_status"] = *model.AuthenticationStatus
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.HostSettingsCheckResults != nil {
		hostSettingsCheckResults := []map[string]interface{}{}
		for _, hostSettingsCheckResultsItem := range model.HostSettingsCheckResults {
			hostSettingsCheckResultsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHostSettingsCheckResultToMap(&hostSettingsCheckResultsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			hostSettingsCheckResults = append(hostSettingsCheckResults, hostSettingsCheckResultsItemMap)
		}
		modelMap["host_settings_check_results"] = hostSettingsCheckResults
	}
	if model.RefreshErrorMessage != nil {
		modelMap["refresh_error_message"] = *model.RefreshErrorMessage
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoHostSettingsCheckResultToMap(model *backuprecoveryv1.HostSettingsCheckResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CheckType != nil {
		modelMap["check_type"] = *model.CheckType
	}
	if model.ResultType != nil {
		modelMap["result_type"] = *model.ResultType
	}
	if model.UserMessage != nil {
		modelMap["user_message"] = *model.UserMessage
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSubnetToMap(model *backuprecoveryv1.Subnet) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Component != nil {
		modelMap["component"] = *model.Component
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Ip != nil {
		modelMap["ip"] = *model.Ip
	}
	if model.NetmaskBits != nil {
		modelMap["netmask_bits"] = *model.NetmaskBits
	}
	if model.NetmaskIp4 != nil {
		modelMap["netmask_ip4"] = *model.NetmaskIp4
	}
	if model.NfsAccess != nil {
		modelMap["nfs_access"] = *model.NfsAccess
	}
	if model.NfsAllSquash != nil {
		modelMap["nfs_all_squash"] = *model.NfsAllSquash
	}
	if model.NfsRootSquash != nil {
		modelMap["nfs_root_squash"] = *model.NfsRootSquash
	}
	if model.S3Access != nil {
		modelMap["s3_access"] = *model.S3Access
	}
	if model.SmbAccess != nil {
		modelMap["smb_access"] = *model.SmbAccess
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyToMap(model *backuprecoveryv1.ThrottlingPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnforceMaxStreams != nil {
		modelMap["enforce_max_streams"] = *model.EnforceMaxStreams
	}
	if model.EnforceRegisteredSourceMaxBackups != nil {
		modelMap["enforce_registered_source_max_backups"] = *model.EnforceRegisteredSourceMaxBackups
	}
	if model.IsEnabled != nil {
		modelMap["is_enabled"] = *model.IsEnabled
	}
	if model.LatencyThresholds != nil {
		latencyThresholdsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoLatencyThresholdsToMap(model.LatencyThresholds)
		if err != nil {
			return modelMap, err
		}
		modelMap["latency_thresholds"] = []map[string]interface{}{latencyThresholdsMap}
	}
	if model.MaxConcurrentStreams != nil {
		modelMap["max_concurrent_streams"] = *model.MaxConcurrentStreams
	}
	if model.NasSourceParams != nil {
		nasSourceParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoNasSourceParamsToMap(model.NasSourceParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_source_params"] = []map[string]interface{}{nasSourceParamsMap}
	}
	if model.RegisteredSourceMaxConcurrentBackups != nil {
		modelMap["registered_source_max_concurrent_backups"] = *model.RegisteredSourceMaxConcurrentBackups
	}
	if model.StorageArraySnapshotConfig != nil {
		storageArraySnapshotConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigToMap(model.StorageArraySnapshotConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoLatencyThresholdsToMap(model *backuprecoveryv1.LatencyThresholds) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActiveTaskMsecs != nil {
		modelMap["active_task_msecs"] = flex.IntValue(model.ActiveTaskMsecs)
	}
	if model.NewTaskMsecs != nil {
		modelMap["new_task_msecs"] = flex.IntValue(model.NewTaskMsecs)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoNasSourceParamsToMap(model *backuprecoveryv1.NasSourceParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxParallelMetadataFetchFullPercentage != nil {
		modelMap["max_parallel_metadata_fetch_full_percentage"] = *model.MaxParallelMetadataFetchFullPercentage
	}
	if model.MaxParallelMetadataFetchIncrementalPercentage != nil {
		modelMap["max_parallel_metadata_fetch_incremental_percentage"] = *model.MaxParallelMetadataFetchIncrementalPercentage
	}
	if model.MaxParallelReadWriteFullPercentage != nil {
		modelMap["max_parallel_read_write_full_percentage"] = *model.MaxParallelReadWriteFullPercentage
	}
	if model.MaxParallelReadWriteIncrementalPercentage != nil {
		modelMap["max_parallel_read_write_incremental_percentage"] = *model.MaxParallelReadWriteIncrementalPercentage
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigToMap(model *backuprecoveryv1.StorageArraySnapshotConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsMaxSnapshotsConfigEnabled != nil {
		modelMap["is_max_snapshots_config_enabled"] = *model.IsMaxSnapshotsConfigEnabled
	}
	if model.IsMaxSpaceConfigEnabled != nil {
		modelMap["is_max_space_config_enabled"] = *model.IsMaxSpaceConfigEnabled
	}
	if model.StorageArraySnapshotMaxSpaceConfig != nil {
		storageArraySnapshotMaxSpaceConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigToMap(model.StorageArraySnapshotMaxSpaceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigMap}
	}
	if model.StorageArraySnapshotThrottlingPolicies != nil {
		storageArraySnapshotThrottlingPolicies := []map[string]interface{}{}
		for _, storageArraySnapshotThrottlingPoliciesItem := range model.StorageArraySnapshotThrottlingPolicies {
			storageArraySnapshotThrottlingPoliciesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPoliciesToMap(&storageArraySnapshotThrottlingPoliciesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			storageArraySnapshotThrottlingPolicies = append(storageArraySnapshotThrottlingPolicies, storageArraySnapshotThrottlingPoliciesItemMap)
		}
		modelMap["storage_array_snapshot_throttling_policies"] = storageArraySnapshotThrottlingPolicies
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigToMap(model *backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxSnapshotSpacePercentage != nil {
		modelMap["max_snapshot_space_percentage"] = *model.MaxSnapshotSpacePercentage
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPoliciesToMap(model *backuprecoveryv1.StorageArraySnapshotThrottlingPolicies) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.IsMaxSnapshotsConfigEnabled != nil {
		modelMap["is_max_snapshots_config_enabled"] = *model.IsMaxSnapshotsConfigEnabled
	}
	if model.IsMaxSpaceConfigEnabled != nil {
		modelMap["is_max_space_config_enabled"] = *model.IsMaxSpaceConfigEnabled
	}
	if model.MaxSnapshotConfig != nil {
		maxSnapshotConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoMaxSnapshotConfigToMap(model.MaxSnapshotConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigMap}
	}
	if model.MaxSpaceConfig != nil {
		maxSpaceConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoMaxSpaceConfigToMap(model.MaxSpaceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["max_space_config"] = []map[string]interface{}{maxSpaceConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoMaxSnapshotConfigToMap(model *backuprecoveryv1.MaxSnapshotConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxSnapshots != nil {
		modelMap["max_snapshots"] = *model.MaxSnapshots
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoMaxSpaceConfigToMap(model *backuprecoveryv1.MaxSpaceConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxSnapshotSpacePercentage != nil {
		modelMap["max_snapshot_space_percentage"] = *model.MaxSnapshotSpacePercentage
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverridesToMap(model *backuprecoveryv1.ThrottlingPolicyOverrides) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatastoreID != nil {
		modelMap["datastore_id"] = flex.IntValue(model.DatastoreID)
	}
	if model.DatastoreName != nil {
		modelMap["datastore_name"] = *model.DatastoreName
	}
	if model.ThrottlingPolicy != nil {
		throttlingPolicyMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyToMap(model.ThrottlingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_policy"] = []map[string]interface{}{throttlingPolicyMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceVlanConfigToMap(model *backuprecoveryv1.RegisteredSourceVlanConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Vlan != nil {
		modelMap["vlan"] = *model.Vlan
	}
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = *model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = *model.InterfaceName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoUniqueGlobalIDToMap(model *backuprecoveryv1.UniqueGlobalID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoNetworkingInformationToMap(model *backuprecoveryv1.NetworkingInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceVec != nil {
		resourceVec := []map[string]interface{}{}
		for _, resourceVecItem := range model.ResourceVec {
			resourceVecItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkResourceInformationToMap(&resourceVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resourceVec = append(resourceVec, resourceVecItemMap)
		}
		modelMap["resource_vec"] = resourceVec
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkResourceInformationToMap(model *backuprecoveryv1.ClusterNetworkResourceInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Endpoints != nil {
		endpoints := []map[string]interface{}{}
		for _, endpointsItem := range model.Endpoints {
			endpointsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkingEndpointToMap(&endpointsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			endpoints = append(endpoints, endpointsItemMap)
		}
		modelMap["endpoints"] = endpoints
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkingEndpointToMap(model *backuprecoveryv1.ClusterNetworkingEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Fqdn != nil {
		modelMap["fqdn"] = *model.Fqdn
	}
	if model.Ipv4Addr != nil {
		modelMap["ipv4_addr"] = *model.Ipv4Addr
	}
	if model.Ipv6Addr != nil {
		modelMap["ipv6_addr"] = *model.Ipv6Addr
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoPhysicalVolumeToMap(model *backuprecoveryv1.PhysicalVolume) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DevicePath != nil {
		modelMap["device_path"] = *model.DevicePath
	}
	if model.Guid != nil {
		modelMap["guid"] = *model.Guid
	}
	if model.IsBootVolume != nil {
		modelMap["is_boot_volume"] = *model.IsBootVolume
	}
	if model.IsExtendedAttributesSupported != nil {
		modelMap["is_extended_attributes_supported"] = *model.IsExtendedAttributesSupported
	}
	if model.IsProtected != nil {
		modelMap["is_protected"] = *model.IsProtected
	}
	if model.IsSharedVolume != nil {
		modelMap["is_shared_volume"] = *model.IsSharedVolume
	}
	if model.Label != nil {
		modelMap["label"] = *model.Label
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = *model.LogicalSizeBytes
	}
	if model.MountPoints != nil {
		modelMap["mount_points"] = model.MountPoints
	}
	if model.MountType != nil {
		modelMap["mount_type"] = *model.MountType
	}
	if model.NetworkPath != nil {
		modelMap["network_path"] = *model.NetworkPath
	}
	if model.UsedSizeBytes != nil {
		modelMap["used_size_bytes"] = *model.UsedSizeBytes
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoVssWritersToMap(model *backuprecoveryv1.VssWriters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsWriterExcluded != nil {
		modelMap["is_writer_excluded"] = *model.IsWriterExcluded
	}
	if model.WriterName != nil {
		modelMap["writer_name"] = *model.WriterName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKubernetesProtectionSourceToMap(model *backuprecoveryv1.KubernetesProtectionSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatamoverImageLocation != nil {
		modelMap["datamover_image_location"] = *model.DatamoverImageLocation
	}
	if model.DatamoverServiceType != nil {
		modelMap["datamover_service_type"] = flex.IntValue(model.DatamoverServiceType)
	}
	if model.DatamoverUpgradability != nil {
		modelMap["datamover_upgradability"] = flex.IntValue(model.DatamoverUpgradability)
	}
	if model.DefaultVlanParams != nil {
		defaultVlanParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoVlanParametersToMap(model.DefaultVlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["default_vlan_params"] = []map[string]interface{}{defaultVlanParamsMap}
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Distribution != nil {
		modelMap["distribution"] = *model.Distribution
	}
	if model.InitContainerImageLocation != nil {
		modelMap["init_container_image_location"] = *model.InitContainerImageLocation
	}
	if model.LabelAttributes != nil {
		labelAttributes := []map[string]interface{}{}
		for _, labelAttributesItem := range model.LabelAttributes {
			labelAttributesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKubernetesLabelAttributeToMap(&labelAttributesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			labelAttributes = append(labelAttributes, labelAttributesItemMap)
		}
		modelMap["label_attributes"] = labelAttributes
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.PriorityClassName != nil {
		modelMap["priority_class_name"] = *model.PriorityClassName
	}
	if model.ResourceAnnotationList != nil {
		resourceAnnotationList := []map[string]interface{}{}
		for _, resourceAnnotationListItem := range model.ResourceAnnotationList {
			resourceAnnotationListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoK8sLabelToMap(&resourceAnnotationListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resourceAnnotationList = append(resourceAnnotationList, resourceAnnotationListItemMap)
		}
		modelMap["resource_annotation_list"] = resourceAnnotationList
	}
	if model.ResourceLabelList != nil {
		resourceLabelList := []map[string]interface{}{}
		for _, resourceLabelListItem := range model.ResourceLabelList {
			resourceLabelListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoK8sLabelToMap(&resourceLabelListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resourceLabelList = append(resourceLabelList, resourceLabelListItemMap)
		}
		modelMap["resource_label_list"] = resourceLabelList
	}
	if model.SanField != nil {
		modelMap["san_field"] = model.SanField
	}
	if model.ServiceAnnotations != nil {
		serviceAnnotations := []map[string]interface{}{}
		for _, serviceAnnotationsItem := range model.ServiceAnnotations {
			serviceAnnotationsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoServiceAnnotationsEntryToMap(&serviceAnnotationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			serviceAnnotations = append(serviceAnnotations, serviceAnnotationsItemMap)
		}
		modelMap["service_annotations"] = serviceAnnotations
	}
	if model.StorageClass != nil {
		storageClass := []map[string]interface{}{}
		for _, storageClassItem := range model.StorageClass {
			storageClassItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKubernetesStorageClassInfoToMap(&storageClassItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			storageClass = append(storageClass, storageClassItemMap)
		}
		modelMap["storage_class"] = storageClass
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	if model.VeleroAwsPluginImageLocation != nil {
		modelMap["velero_aws_plugin_image_location"] = *model.VeleroAwsPluginImageLocation
	}
	if model.VeleroImageLocation != nil {
		modelMap["velero_image_location"] = *model.VeleroImageLocation
	}
	if model.VeleroOpenshiftPluginImageLocation != nil {
		modelMap["velero_openshift_plugin_image_location"] = *model.VeleroOpenshiftPluginImageLocation
	}
	if model.VeleroUpgradability != nil {
		modelMap["velero_upgradability"] = *model.VeleroUpgradability
	}
	if model.VlanInfoVec != nil {
		vlanInfoVec := []map[string]interface{}{}
		for _, vlanInfoVecItem := range model.VlanInfoVec {
			vlanInfoVecItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKubernetesVlanInfoToMap(&vlanInfoVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			vlanInfoVec = append(vlanInfoVec, vlanInfoVecItemMap)
		}
		modelMap["vlan_info_vec"] = vlanInfoVec
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoVlanParametersToMap(model *backuprecoveryv1.VlanParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = *model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = *model.InterfaceName
	}
	if model.Vlan != nil {
		modelMap["vlan"] = flex.IntValue(model.Vlan)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKubernetesLabelAttributeToMap(model *backuprecoveryv1.KubernetesLabelAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoK8sLabelToMap(model *backuprecoveryv1.K8sLabel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoServiceAnnotationsEntryToMap(model *backuprecoveryv1.ServiceAnnotationsEntry) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKubernetesStorageClassInfoToMap(model *backuprecoveryv1.KubernetesStorageClassInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Provisioner != nil {
		modelMap["provisioner"] = *model.Provisioner
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKubernetesVlanInfoToMap(model *backuprecoveryv1.KubernetesVlanInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ServiceAnnotations != nil {
		serviceAnnotations := []map[string]interface{}{}
		for _, serviceAnnotationsItem := range model.ServiceAnnotations {
			serviceAnnotationsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKubernetesServiceAnnotationObjectToMap(&serviceAnnotationsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			serviceAnnotations = append(serviceAnnotations, serviceAnnotationsItemMap)
		}
		modelMap["service_annotations"] = serviceAnnotations
	}
	if model.VlanParams != nil {
		vlanParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoVlanParamsToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKubernetesServiceAnnotationObjectToMap(model *backuprecoveryv1.KubernetesServiceAnnotationObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoVlanParamsToMap(model *backuprecoveryv1.VlanParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DisableVlan != nil {
		modelMap["disable_vlan"] = *model.DisableVlan
	}
	if model.InterfaceName != nil {
		modelMap["interface_name"] = *model.InterfaceName
	}
	if model.VlanID != nil {
		modelMap["vlan_id"] = flex.IntValue(model.VlanID)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSqlProtectionSourceToMap(model *backuprecoveryv1.SqlProtectionSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsAvailableForVssBackup != nil {
		modelMap["is_available_for_vss_backup"] = *model.IsAvailableForVssBackup
	}
	if model.CreatedTimestamp != nil {
		modelMap["created_timestamp"] = *model.CreatedTimestamp
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = *model.DatabaseName
	}
	if model.DbAagEntityID != nil {
		modelMap["db_aag_entity_id"] = flex.IntValue(model.DbAagEntityID)
	}
	if model.DbAagName != nil {
		modelMap["db_aag_name"] = *model.DbAagName
	}
	if model.DbCompatibilityLevel != nil {
		modelMap["db_compatibility_level"] = flex.IntValue(model.DbCompatibilityLevel)
	}
	if model.DbFileGroups != nil {
		modelMap["db_file_groups"] = model.DbFileGroups
	}
	if model.DbFiles != nil {
		dbFiles := []map[string]interface{}{}
		for _, dbFilesItem := range model.DbFiles {
			dbFilesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoDatabaseFileInformationToMap(&dbFilesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			dbFiles = append(dbFiles, dbFilesItemMap)
		}
		modelMap["db_files"] = dbFiles
	}
	if model.DbOwnerUsername != nil {
		modelMap["db_owner_username"] = *model.DbOwnerUsername
	}
	if model.DefaultDatabaseLocation != nil {
		modelMap["default_database_location"] = *model.DefaultDatabaseLocation
	}
	if model.DefaultLogLocation != nil {
		modelMap["default_log_location"] = *model.DefaultLogLocation
	}
	if model.ID != nil {
		idMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSQLSourceIDToMap(model.ID)
		if err != nil {
			return modelMap, err
		}
		modelMap["id"] = []map[string]interface{}{idMap}
	}
	if model.IsEncrypted != nil {
		modelMap["is_encrypted"] = *model.IsEncrypted
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.OwnerID != nil {
		modelMap["owner_id"] = flex.IntValue(model.OwnerID)
	}
	if model.RecoveryModel != nil {
		modelMap["recovery_model"] = *model.RecoveryModel
	}
	if model.SqlServerDbState != nil {
		modelMap["sql_server_db_state"] = *model.SqlServerDbState
	}
	if model.SqlServerInstanceVersion != nil {
		sqlServerInstanceVersionMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSQLServerInstanceVersionToMap(model.SqlServerInstanceVersion)
		if err != nil {
			return modelMap, err
		}
		modelMap["sql_server_instance_version"] = []map[string]interface{}{sqlServerInstanceVersionMap}
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoDatabaseFileInformationToMap(model *backuprecoveryv1.DatabaseFileInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FileType != nil {
		modelMap["file_type"] = *model.FileType
	}
	if model.FullPath != nil {
		modelMap["full_path"] = *model.FullPath
	}
	if model.SizeBytes != nil {
		modelMap["size_bytes"] = flex.IntValue(model.SizeBytes)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSQLSourceIDToMap(model *backuprecoveryv1.SQLSourceID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CreatedDateMsecs != nil {
		modelMap["created_date_msecs"] = flex.IntValue(model.CreatedDateMsecs)
	}
	if model.DatabaseID != nil {
		modelMap["database_id"] = flex.IntValue(model.DatabaseID)
	}
	if model.InstanceID != nil {
		modelMap["instance_id"] = *model.InstanceID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSQLServerInstanceVersionToMap(model *backuprecoveryv1.SQLServerInstanceVersion) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Build != nil {
		modelMap["build"] = *model.Build
	}
	if model.MajorVersion != nil {
		modelMap["major_version"] = *model.MajorVersion
	}
	if model.MinorVersion != nil {
		modelMap["minor_version"] = *model.MinorVersion
	}
	if model.Revision != nil {
		modelMap["revision"] = *model.Revision
	}
	if model.VersionString != nil {
		modelMap["version_string"] = *model.VersionString
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoEntityPermissionInformationToMap(model *backuprecoveryv1.EntityPermissionInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EntityID != nil {
		modelMap["entity_id"] = flex.IntValue(model.EntityID)
	}
	if model.Groups != nil {
		groups := []map[string]interface{}{}
		for _, groupsItem := range model.Groups {
			groupsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoGroupInfoToMap(&groupsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.IsInferred != nil {
		modelMap["is_inferred"] = *model.IsInferred
	}
	if model.IsRegisteredBySp != nil {
		modelMap["is_registered_by_sp"] = *model.IsRegisteredBySp
	}
	if model.RegisteringTenantID != nil {
		modelMap["registering_tenant_id"] = *model.RegisteringTenantID
	}
	if model.Tenant != nil {
		tenantMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTenantInfoToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoUserInfoToMap(&usersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			users = append(users, usersItemMap)
		}
		modelMap["users"] = users
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoGroupInfoToMap(model *backuprecoveryv1.GroupInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	if model.GroupName != nil {
		modelMap["group_name"] = *model.GroupName
	}
	if model.Sid != nil {
		modelMap["sid"] = *model.Sid
	}
	if model.TenantIds != nil {
		modelMap["tenant_ids"] = model.TenantIds
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoTenantInfoToMap(model *backuprecoveryv1.TenantInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BifrostEnabled != nil {
		modelMap["bifrost_enabled"] = *model.BifrostEnabled
	}
	if model.IsManagedOnHelios != nil {
		modelMap["is_managed_on_helios"] = *model.IsManagedOnHelios
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoUserInfoToMap(model *backuprecoveryv1.UserInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	if model.Sid != nil {
		modelMap["sid"] = *model.Sid
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	if model.UserName != nil {
		modelMap["user_name"] = *model.UserName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoMaintenanceModeConfigToMap(model *backuprecoveryv1.MaintenanceModeConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActivationTimeIntervals != nil {
		activationTimeIntervals := []map[string]interface{}{}
		for _, activationTimeIntervalsItem := range model.ActivationTimeIntervals {
			activationTimeIntervalsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeRangeUsecsToMap(&activationTimeIntervalsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			activationTimeIntervals = append(activationTimeIntervals, activationTimeIntervalsItemMap)
		}
		modelMap["activation_time_intervals"] = activationTimeIntervals
	}
	if model.MaintenanceSchedule != nil {
		maintenanceScheduleMap, err := DataSourceIbmBackupRecoveryRegistrationInfoScheduleToMap(model.MaintenanceSchedule)
		if err != nil {
			return modelMap, err
		}
		modelMap["maintenance_schedule"] = []map[string]interface{}{maintenanceScheduleMap}
	}
	if model.UserMessage != nil {
		modelMap["user_message"] = *model.UserMessage
	}
	if model.WorkflowInterventionSpecList != nil {
		workflowInterventionSpecList := []map[string]interface{}{}
		for _, workflowInterventionSpecListItem := range model.WorkflowInterventionSpecList {
			workflowInterventionSpecListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoWorkflowInterventionSpecToMap(&workflowInterventionSpecListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			workflowInterventionSpecList = append(workflowInterventionSpecList, workflowInterventionSpecListItemMap)
		}
		modelMap["workflow_intervention_spec_list"] = workflowInterventionSpecList
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoTimeRangeUsecsToMap(model *backuprecoveryv1.TimeRangeUsecs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoScheduleToMap(model *backuprecoveryv1.Schedule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PeriodicTimeWindows != nil {
		periodicTimeWindows := []map[string]interface{}{}
		for _, periodicTimeWindowsItem := range model.PeriodicTimeWindows {
			periodicTimeWindowsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeWindowToMap(&periodicTimeWindowsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			periodicTimeWindows = append(periodicTimeWindows, periodicTimeWindowsItemMap)
		}
		modelMap["periodic_time_windows"] = periodicTimeWindows
	}
	if model.ScheduleType != nil {
		modelMap["schedule_type"] = *model.ScheduleType
	}
	if model.TimeRanges != nil {
		timeRanges := []map[string]interface{}{}
		for _, timeRangesItem := range model.TimeRanges {
			timeRangesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeRangeUsecsToMap(&timeRangesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			timeRanges = append(timeRanges, timeRangesItemMap)
		}
		modelMap["time_ranges"] = timeRanges
	}
	if model.Timezone != nil {
		modelMap["timezone"] = *model.Timezone
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoTimeWindowToMap(model *backuprecoveryv1.TimeWindow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DayOfTheWeek != nil {
		modelMap["day_of_the_week"] = *model.DayOfTheWeek
	}
	if model.EndTime != nil {
		endTimeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(model.EndTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["end_time"] = []map[string]interface{}{endTimeMap}
	}
	if model.StartTime != nil {
		startTimeMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(model.StartTime)
		if err != nil {
			return modelMap, err
		}
		modelMap["start_time"] = []map[string]interface{}{startTimeMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoWorkflowInterventionSpecToMap(model *backuprecoveryv1.WorkflowInterventionSpec) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["intervention"] = *model.Intervention
	modelMap["workflow_type"] = *model.WorkflowType
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceInfoToMap(model *backuprecoveryv1.RegisteredSourceInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccessInfo != nil {
		accessInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoConnectorParametersToMap(model.AccessInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["access_info"] = []map[string]interface{}{accessInfoMap}
	}
	if model.AllowedIpAddresses != nil {
		modelMap["allowed_ip_addresses"] = model.AllowedIpAddresses
	}
	if model.AuthenticationErrorMessage != nil {
		modelMap["authentication_error_message"] = *model.AuthenticationErrorMessage
	}
	if model.AuthenticationStatus != nil {
		modelMap["authentication_status"] = *model.AuthenticationStatus
	}
	if model.BlacklistedIpAddresses != nil {
		modelMap["blacklisted_ip_addresses"] = model.BlacklistedIpAddresses
	}
	if model.CassandraParams != nil {
		cassandraParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCassandraConnectParamsToMap(model.CassandraParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["cassandra_params"] = []map[string]interface{}{cassandraParamsMap}
	}
	if model.CouchbaseParams != nil {
		couchbaseParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCouchbaseConnectParamsToMap(model.CouchbaseParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["couchbase_params"] = []map[string]interface{}{couchbaseParamsMap}
	}
	if model.DeniedIpAddresses != nil {
		modelMap["denied_ip_addresses"] = model.DeniedIpAddresses
	}
	if model.Environments != nil {
		modelMap["environments"] = model.Environments
	}
	if model.HbaseParams != nil {
		hbaseParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHBaseConnectParamsToMap(model.HbaseParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hbase_params"] = []map[string]interface{}{hbaseParamsMap}
	}
	if model.HdfsParams != nil {
		hdfsParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHdfsConnectParamsToMap(model.HdfsParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hdfs_params"] = []map[string]interface{}{hdfsParamsMap}
	}
	if model.HiveParams != nil {
		hiveParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHiveConnectParamsToMap(model.HiveParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hive_params"] = []map[string]interface{}{hiveParamsMap}
	}
	if model.IsDbAuthenticated != nil {
		modelMap["is_db_authenticated"] = *model.IsDbAuthenticated
	}
	if model.IsStorageArraySnapshotEnabled != nil {
		modelMap["is_storage_array_snapshot_enabled"] = *model.IsStorageArraySnapshotEnabled
	}
	if model.IsilonParams != nil {
		isilonParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoRegisteredProtectionSourceIsilonParamsToMap(model.IsilonParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["isilon_params"] = []map[string]interface{}{isilonParamsMap}
	}
	if model.LinkVmsAcrossVcenter != nil {
		modelMap["link_vms_across_vcenter"] = *model.LinkVmsAcrossVcenter
	}
	if model.MinimumFreeSpaceGB != nil {
		modelMap["minimum_free_space_gb"] = flex.IntValue(model.MinimumFreeSpaceGB)
	}
	if model.MinimumFreeSpacePercent != nil {
		modelMap["minimum_free_space_percent"] = flex.IntValue(model.MinimumFreeSpacePercent)
	}
	if model.MongodbParams != nil {
		mongodbParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoMongoDBConnectParamsToMap(model.MongodbParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["mongodb_params"] = []map[string]interface{}{mongodbParamsMap}
	}
	if model.NasMountCredentials != nil {
		nasMountCredentialsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoNASServerCredentialsToMap(model.NasMountCredentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_mount_credentials"] = []map[string]interface{}{nasMountCredentialsMap}
	}
	if model.O365Params != nil {
		o365ParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoO365ConnectParamsToMap(model.O365Params)
		if err != nil {
			return modelMap, err
		}
		modelMap["o365_params"] = []map[string]interface{}{o365ParamsMap}
	}
	if model.Office365CredentialsList != nil {
		office365CredentialsList := []map[string]interface{}{}
		for _, office365CredentialsListItem := range model.Office365CredentialsList {
			office365CredentialsListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoOffice365CredentialsToMap(&office365CredentialsListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			office365CredentialsList = append(office365CredentialsList, office365CredentialsListItemMap)
		}
		modelMap["office365_credentials_list"] = office365CredentialsList
	}
	if model.Office365Region != nil {
		modelMap["office365_region"] = *model.Office365Region
	}
	if model.Office365ServiceAccountCredentialsList != nil {
		office365ServiceAccountCredentialsList := []map[string]interface{}{}
		for _, office365ServiceAccountCredentialsListItem := range model.Office365ServiceAccountCredentialsList {
			office365ServiceAccountCredentialsListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCredentialsToMap(&office365ServiceAccountCredentialsListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			office365ServiceAccountCredentialsList = append(office365ServiceAccountCredentialsList, office365ServiceAccountCredentialsListItemMap)
		}
		modelMap["office365_service_account_credentials_list"] = office365ServiceAccountCredentialsList
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.ProgressMonitorPath != nil {
		modelMap["progress_monitor_path"] = *model.ProgressMonitorPath
	}
	if model.RefreshErrorMessage != nil {
		modelMap["refresh_error_message"] = *model.RefreshErrorMessage
	}
	if model.RefreshTimeUsecs != nil {
		modelMap["refresh_time_usecs"] = flex.IntValue(model.RefreshTimeUsecs)
	}
	if model.RegisteredAppsInfo != nil {
		registeredAppsInfo := []map[string]interface{}{}
		for _, registeredAppsInfoItem := range model.RegisteredAppsInfo {
			registeredAppsInfoItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoRegisteredAppInfoToMap(&registeredAppsInfoItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			registeredAppsInfo = append(registeredAppsInfo, registeredAppsInfoItemMap)
		}
		modelMap["registered_apps_info"] = registeredAppsInfo
	}
	if model.RegistrationTimeUsecs != nil {
		modelMap["registration_time_usecs"] = flex.IntValue(model.RegistrationTimeUsecs)
	}
	if model.SfdcParams != nil {
		sfdcParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSfdcParamsToMap(model.SfdcParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sfdc_params"] = []map[string]interface{}{sfdcParamsMap}
	}
	if model.Subnets != nil {
		subnets := []map[string]interface{}{}
		for _, subnetsItem := range model.Subnets {
			subnetsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSubnetToMap(&subnetsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			subnets = append(subnets, subnetsItemMap)
		}
		modelMap["subnets"] = subnets
	}
	if model.ThrottlingPolicy != nil {
		throttlingPolicyMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyParametersToMap(model.ThrottlingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_policy"] = []map[string]interface{}{throttlingPolicyMap}
	}
	if model.ThrottlingPolicyOverrides != nil {
		throttlingPolicyOverrides := []map[string]interface{}{}
		for _, throttlingPolicyOverridesItem := range model.ThrottlingPolicyOverrides {
			throttlingPolicyOverridesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverrideToMap(&throttlingPolicyOverridesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			throttlingPolicyOverrides = append(throttlingPolicyOverrides, throttlingPolicyOverridesItemMap)
		}
		modelMap["throttling_policy_overrides"] = throttlingPolicyOverrides
	}
	if model.UdaParams != nil {
		udaParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoUdaConnectParamsToMap(model.UdaParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["uda_params"] = []map[string]interface{}{udaParamsMap}
	}
	if model.UpdateLastBackupDetails != nil {
		modelMap["update_last_backup_details"] = *model.UpdateLastBackupDetails
	}
	if model.UseOAuthForExchangeOnline != nil {
		modelMap["use_o_auth_for_exchange_online"] = *model.UseOAuthForExchangeOnline
	}
	if model.UseVmBiosUUID != nil {
		modelMap["use_vm_bios_uuid"] = *model.UseVmBiosUUID
	}
	if model.UserMessages != nil {
		modelMap["user_messages"] = model.UserMessages
	}
	if model.Username != nil {
		modelMap["username"] = *model.Username
	}
	if model.VlanParams != nil {
		vlanParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoVlanParametersToMap(model.VlanParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["vlan_params"] = []map[string]interface{}{vlanParamsMap}
	}
	if model.WarningMessages != nil {
		modelMap["warning_messages"] = model.WarningMessages
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoConnectorParametersToMap(model *backuprecoveryv1.ConnectorParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConnectionID != nil {
		modelMap["connection_id"] = flex.IntValue(model.ConnectionID)
	}
	if model.ConnectorGroupID != nil {
		modelMap["connector_group_id"] = flex.IntValue(model.ConnectorGroupID)
	}
	if model.Endpoint != nil {
		modelMap["endpoint"] = *model.Endpoint
	}
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCassandraConnectParamsToMap(model *backuprecoveryv1.CassandraConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CassandraPortsInfo != nil {
		cassandraPortsInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCassandraPortsInfoToMap(model.CassandraPortsInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["cassandra_ports_info"] = []map[string]interface{}{cassandraPortsInfoMap}
	}
	if model.CassandraSecurityInfo != nil {
		cassandraSecurityInfoMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCassandraSecurityInfoToMap(model.CassandraSecurityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["cassandra_security_info"] = []map[string]interface{}{cassandraSecurityInfoMap}
	}
	if model.CassandraVersion != nil {
		modelMap["cassandra_version"] = *model.CassandraVersion
	}
	if model.CommitLogBackupLocation != nil {
		modelMap["commit_log_backup_location"] = *model.CommitLogBackupLocation
	}
	if model.ConfigDirectory != nil {
		modelMap["config_directory"] = *model.ConfigDirectory
	}
	if model.DataCenters != nil {
		modelMap["data_centers"] = model.DataCenters
	}
	if model.DseConfigDirectory != nil {
		modelMap["dse_config_directory"] = *model.DseConfigDirectory
	}
	if model.DseVersion != nil {
		modelMap["dse_version"] = *model.DseVersion
	}
	if model.IsDseAuthenticator != nil {
		modelMap["is_dse_authenticator"] = *model.IsDseAuthenticator
	}
	if model.IsDseTieredStorage != nil {
		modelMap["is_dse_tiered_storage"] = *model.IsDseTieredStorage
	}
	if model.IsJmxAuthEnable != nil {
		modelMap["is_jmx_auth_enable"] = *model.IsJmxAuthEnable
	}
	if model.KerberosPrincipal != nil {
		modelMap["kerberos_principal"] = *model.KerberosPrincipal
	}
	if model.PrimaryHost != nil {
		modelMap["primary_host"] = *model.PrimaryHost
	}
	if model.Seeds != nil {
		modelMap["seeds"] = model.Seeds
	}
	if model.SolrNodes != nil {
		modelMap["solr_nodes"] = model.SolrNodes
	}
	if model.SolrPort != nil {
		modelMap["solr_port"] = flex.IntValue(model.SolrPort)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCassandraPortsInfoToMap(model *backuprecoveryv1.CassandraPortsInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JmxPort != nil {
		modelMap["jmx_port"] = flex.IntValue(model.JmxPort)
	}
	if model.NativeTransportPort != nil {
		modelMap["native_transport_port"] = flex.IntValue(model.NativeTransportPort)
	}
	if model.RpcPort != nil {
		modelMap["rpc_port"] = flex.IntValue(model.RpcPort)
	}
	if model.SslStoragePort != nil {
		modelMap["ssl_storage_port"] = flex.IntValue(model.SslStoragePort)
	}
	if model.StoragePort != nil {
		modelMap["storage_port"] = flex.IntValue(model.StoragePort)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCassandraSecurityInfoToMap(model *backuprecoveryv1.CassandraSecurityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CassandraAuthRequired != nil {
		modelMap["cassandra_auth_required"] = *model.CassandraAuthRequired
	}
	if model.CassandraAuthType != nil {
		modelMap["cassandra_auth_type"] = *model.CassandraAuthType
	}
	if model.CassandraAuthorizer != nil {
		modelMap["cassandra_authorizer"] = *model.CassandraAuthorizer
	}
	if model.ClientEncryption != nil {
		modelMap["client_encryption"] = *model.ClientEncryption
	}
	if model.DseAuthorization != nil {
		modelMap["dse_authorization"] = *model.DseAuthorization
	}
	if model.ServerEncryptionReqClientAuth != nil {
		modelMap["server_encryption_req_client_auth"] = *model.ServerEncryptionReqClientAuth
	}
	if model.ServerInternodeEncryptionType != nil {
		modelMap["server_internode_encryption_type"] = *model.ServerInternodeEncryptionType
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCouchbaseConnectParamsToMap(model *backuprecoveryv1.CouchbaseConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CarrierDirectPort != nil {
		modelMap["carrier_direct_port"] = flex.IntValue(model.CarrierDirectPort)
	}
	if model.HttpDirectPort != nil {
		modelMap["http_direct_port"] = flex.IntValue(model.HttpDirectPort)
	}
	if model.RequiresSsl != nil {
		modelMap["requires_ssl"] = *model.RequiresSsl
	}
	if model.Seeds != nil {
		modelMap["seeds"] = model.Seeds
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoHBaseConnectParamsToMap(model *backuprecoveryv1.HBaseConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.HbaseDiscoveryParams != nil {
		hbaseDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(model.HbaseDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hbase_discovery_params"] = []map[string]interface{}{hbaseDiscoveryParamsMap}
	}
	if model.HdfsEntityID != nil {
		modelMap["hdfs_entity_id"] = flex.IntValue(model.HdfsEntityID)
	}
	if model.KerberosPrincipal != nil {
		modelMap["kerberos_principal"] = *model.KerberosPrincipal
	}
	if model.RootDataDirectory != nil {
		modelMap["root_data_directory"] = *model.RootDataDirectory
	}
	if model.ZookeeperQuorum != nil {
		modelMap["zookeeper_quorum"] = model.ZookeeperQuorum
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(model *backuprecoveryv1.HadoopDiscoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigDirectory != nil {
		modelMap["config_directory"] = *model.ConfigDirectory
	}
	if model.Host != nil {
		modelMap["host"] = *model.Host
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoHdfsConnectParamsToMap(model *backuprecoveryv1.HdfsConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.HadoopDistribution != nil {
		modelMap["hadoop_distribution"] = *model.HadoopDistribution
	}
	if model.HadoopVersion != nil {
		modelMap["hadoop_version"] = *model.HadoopVersion
	}
	if model.HdfsConnectionType != nil {
		modelMap["hdfs_connection_type"] = *model.HdfsConnectionType
	}
	if model.HdfsDiscoveryParams != nil {
		hdfsDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(model.HdfsDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hdfs_discovery_params"] = []map[string]interface{}{hdfsDiscoveryParamsMap}
	}
	if model.KerberosPrincipal != nil {
		modelMap["kerberos_principal"] = *model.KerberosPrincipal
	}
	if model.Namenode != nil {
		modelMap["namenode"] = *model.Namenode
	}
	if model.Port != nil {
		modelMap["port"] = flex.IntValue(model.Port)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoHiveConnectParamsToMap(model *backuprecoveryv1.HiveConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EntityThresholdExceeded != nil {
		modelMap["entity_threshold_exceeded"] = *model.EntityThresholdExceeded
	}
	if model.HdfsEntityID != nil {
		modelMap["hdfs_entity_id"] = flex.IntValue(model.HdfsEntityID)
	}
	if model.HiveDiscoveryParams != nil {
		hiveDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(model.HiveDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["hive_discovery_params"] = []map[string]interface{}{hiveDiscoveryParamsMap}
	}
	if model.KerberosPrincipal != nil {
		modelMap["kerberos_principal"] = *model.KerberosPrincipal
	}
	if model.Metastore != nil {
		modelMap["metastore"] = *model.Metastore
	}
	if model.ThriftPort != nil {
		modelMap["thrift_port"] = flex.IntValue(model.ThriftPort)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoRegisteredProtectionSourceIsilonParamsToMap(model *backuprecoveryv1.RegisteredProtectionSourceIsilonParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ZoneConfigList != nil {
		zoneConfigList := []map[string]interface{}{}
		for _, zoneConfigListItem := range model.ZoneConfigList {
			zoneConfigListItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoZoneConfigToMap(&zoneConfigListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			zoneConfigList = append(zoneConfigList, zoneConfigListItemMap)
		}
		modelMap["zone_config_list"] = zoneConfigList
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoZoneConfigToMap(model *backuprecoveryv1.ZoneConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DynamicNetworkPoolConfig != nil {
		dynamicNetworkPoolConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoNetworkPoolConfigToMap(model.DynamicNetworkPoolConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["dynamic_network_pool_config"] = []map[string]interface{}{dynamicNetworkPoolConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoNetworkPoolConfigToMap(model *backuprecoveryv1.NetworkPoolConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PoolName != nil {
		modelMap["pool_name"] = *model.PoolName
	}
	if model.Subnet != nil {
		modelMap["subnet"] = *model.Subnet
	}
	if model.UseSmartConnect != nil {
		modelMap["use_smart_connect"] = *model.UseSmartConnect
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoMongoDBConnectParamsToMap(model *backuprecoveryv1.MongoDBConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AuthType != nil {
		modelMap["auth_type"] = *model.AuthType
	}
	if model.AuthenticatingDatabaseName != nil {
		modelMap["authenticating_database_name"] = *model.AuthenticatingDatabaseName
	}
	if model.RequiresSsl != nil {
		modelMap["requires_ssl"] = *model.RequiresSsl
	}
	if model.SecondaryNodeTag != nil {
		modelMap["secondary_node_tag"] = *model.SecondaryNodeTag
	}
	if model.Seeds != nil {
		modelMap["seeds"] = model.Seeds
	}
	if model.UseFixedNodeForBackup != nil {
		modelMap["use_fixed_node_for_backup"] = *model.UseFixedNodeForBackup
	}
	if model.UseSecondaryForBackup != nil {
		modelMap["use_secondary_for_backup"] = *model.UseSecondaryForBackup
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoNASServerCredentialsToMap(model *backuprecoveryv1.NASServerCredentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Domain != nil {
		modelMap["domain"] = *model.Domain
	}
	if model.NasProtocol != nil {
		modelMap["nas_protocol"] = *model.NasProtocol
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoO365ConnectParamsToMap(model *backuprecoveryv1.O365ConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectsDiscoveryParams != nil {
		objectsDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoObjectsDiscoveryParamsToMap(model.ObjectsDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["objects_discovery_params"] = []map[string]interface{}{objectsDiscoveryParamsMap}
	}
	if model.CsmParams != nil {
		csmParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoM365CsmParamsToMap(model.CsmParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["csm_params"] = []map[string]interface{}{csmParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoObjectsDiscoveryParamsToMap(model *backuprecoveryv1.ObjectsDiscoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DiscoverableObjectTypeList != nil {
		modelMap["discoverable_object_type_list"] = model.DiscoverableObjectTypeList
	}
	if model.SitesDiscoveryParams != nil {
		sitesDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSitesDiscoveryParamsToMap(model.SitesDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["sites_discovery_params"] = []map[string]interface{}{sitesDiscoveryParamsMap}
	}
	if model.TeamsAdditionalParams != nil {
		teamsAdditionalParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoTeamsAdditionalParamsToMap(model.TeamsAdditionalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["teams_additional_params"] = []map[string]interface{}{teamsAdditionalParamsMap}
	}
	if model.UsersDiscoveryParams != nil {
		usersDiscoveryParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoUsersDiscoveryParamsToMap(model.UsersDiscoveryParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["users_discovery_params"] = []map[string]interface{}{usersDiscoveryParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSitesDiscoveryParamsToMap(model *backuprecoveryv1.SitesDiscoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSiteTagging != nil {
		modelMap["enable_site_tagging"] = *model.EnableSiteTagging
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoTeamsAdditionalParamsToMap(model *backuprecoveryv1.TeamsAdditionalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AllowPostsBackup != nil {
		modelMap["allow_posts_backup"] = *model.AllowPostsBackup
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoUsersDiscoveryParamsToMap(model *backuprecoveryv1.UsersDiscoveryParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AllowChatsBackup != nil {
		modelMap["allow_chats_backup"] = *model.AllowChatsBackup
	}
	if model.DiscoverUsersWithMailbox != nil {
		modelMap["discover_users_with_mailbox"] = *model.DiscoverUsersWithMailbox
	}
	if model.DiscoverUsersWithOnedrive != nil {
		modelMap["discover_users_with_onedrive"] = *model.DiscoverUsersWithOnedrive
	}
	if model.FetchMailboxInfo != nil {
		modelMap["fetch_mailbox_info"] = *model.FetchMailboxInfo
	}
	if model.FetchOneDriveInfo != nil {
		modelMap["fetch_one_drive_info"] = *model.FetchOneDriveInfo
	}
	if model.SkipUsersWithoutMySite != nil {
		modelMap["skip_users_without_my_site"] = *model.SkipUsersWithoutMySite
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoM365CsmParamsToMap(model *backuprecoveryv1.M365CsmParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BackupAllowed != nil {
		modelMap["backup_allowed"] = *model.BackupAllowed
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoOffice365CredentialsToMap(model *backuprecoveryv1.Office365Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClientID != nil {
		modelMap["client_id"] = *model.ClientID
	}
	if model.ClientSecret != nil {
		modelMap["client_secret"] = *model.ClientSecret
	}
	if model.GrantType != nil {
		modelMap["grant_type"] = *model.GrantType
	}
	if model.Scope != nil {
		modelMap["scope"] = *model.Scope
	}
	if model.UseOAuthForExchangeOnline != nil {
		modelMap["use_o_auth_for_exchange_online"] = *model.UseOAuthForExchangeOnline
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoCredentialsToMap(model *backuprecoveryv1.Credentials) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["username"] = *model.Username
	modelMap["password"] = *model.Password
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoPhysicalParamsToMap(model *backuprecoveryv1.PhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Applications != nil {
		modelMap["applications"] = model.Applications
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.ThrottlingConfig != nil {
		throttlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoSourceThrottlingConfigurationToMap(model.ThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_config"] = []map[string]interface{}{throttlingConfigMap}
	}
	if model.Username != nil {
		modelMap["username"] = *model.Username
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSourceThrottlingConfigurationToMap(model *backuprecoveryv1.SourceThrottlingConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CpuThrottlingConfig != nil {
		cpuThrottlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationToMap(model.CpuThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["cpu_throttling_config"] = []map[string]interface{}{cpuThrottlingConfigMap}
	}
	if model.NetworkThrottlingConfig != nil {
		networkThrottlingConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationToMap(model.NetworkThrottlingConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["network_throttling_config"] = []map[string]interface{}{networkThrottlingConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationToMap(model *backuprecoveryv1.ThrottlingConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FixedThreshold != nil {
		modelMap["fixed_threshold"] = flex.IntValue(model.FixedThreshold)
	}
	if model.PatternType != nil {
		modelMap["pattern_type"] = *model.PatternType
	}
	if model.ThrottlingWindows != nil {
		throttlingWindows := []map[string]interface{}{}
		for _, throttlingWindowsItem := range model.ThrottlingWindows {
			throttlingWindowsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingWindowToMap(&throttlingWindowsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			throttlingWindows = append(throttlingWindows, throttlingWindowsItemMap)
		}
		modelMap["throttling_windows"] = throttlingWindows
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoSfdcParamsToMap(model *backuprecoveryv1.SfdcParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccessToken != nil {
		modelMap["access_token"] = *model.AccessToken
	}
	if model.ConcurrentApiRequestsLimit != nil {
		modelMap["concurrent_api_requests_limit"] = flex.IntValue(model.ConcurrentApiRequestsLimit)
	}
	if model.ConsumerKey != nil {
		modelMap["consumer_key"] = *model.ConsumerKey
	}
	if model.ConsumerSecret != nil {
		modelMap["consumer_secret"] = *model.ConsumerSecret
	}
	if model.DailyApiLimit != nil {
		modelMap["daily_api_limit"] = flex.IntValue(model.DailyApiLimit)
	}
	if model.Endpoint != nil {
		modelMap["endpoint"] = *model.Endpoint
	}
	if model.EndpointType != nil {
		modelMap["endpoint_type"] = *model.EndpointType
	}
	if model.MetadataEndpointURL != nil {
		modelMap["metadata_endpoint_url"] = *model.MetadataEndpointURL
	}
	if model.RefreshToken != nil {
		modelMap["refresh_token"] = *model.RefreshToken
	}
	if model.SoapEndpointURL != nil {
		modelMap["soap_endpoint_url"] = *model.SoapEndpointURL
	}
	if model.UseBulkApi != nil {
		modelMap["use_bulk_api"] = *model.UseBulkApi
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyParametersToMap(model *backuprecoveryv1.ThrottlingPolicyParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnforceMaxStreams != nil {
		modelMap["enforce_max_streams"] = *model.EnforceMaxStreams
	}
	if model.EnforceRegisteredSourceMaxBackups != nil {
		modelMap["enforce_registered_source_max_backups"] = *model.EnforceRegisteredSourceMaxBackups
	}
	if model.IsEnabled != nil {
		modelMap["is_enabled"] = *model.IsEnabled
	}
	if model.LatencyThresholds != nil {
		latencyThresholdsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoLatencyThresholdsToMap(model.LatencyThresholds)
		if err != nil {
			return modelMap, err
		}
		modelMap["latency_thresholds"] = []map[string]interface{}{latencyThresholdsMap}
	}
	if model.MaxConcurrentStreams != nil {
		modelMap["max_concurrent_streams"] = flex.IntValue(model.MaxConcurrentStreams)
	}
	if model.NasSourceParams != nil {
		nasSourceParamsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoNasSourceThrottlingParamsToMap(model.NasSourceParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["nas_source_params"] = []map[string]interface{}{nasSourceParamsMap}
	}
	if model.RegisteredSourceMaxConcurrentBackups != nil {
		modelMap["registered_source_max_concurrent_backups"] = flex.IntValue(model.RegisteredSourceMaxConcurrentBackups)
	}
	if model.StorageArraySnapshotConfig != nil {
		storageArraySnapshotConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigParamsToMap(model.StorageArraySnapshotConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoNasSourceThrottlingParamsToMap(model *backuprecoveryv1.NasSourceThrottlingParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxParallelMetadataFetchFullPercentage != nil {
		modelMap["max_parallel_metadata_fetch_full_percentage"] = flex.IntValue(model.MaxParallelMetadataFetchFullPercentage)
	}
	if model.MaxParallelMetadataFetchIncrementalPercentage != nil {
		modelMap["max_parallel_metadata_fetch_incremental_percentage"] = flex.IntValue(model.MaxParallelMetadataFetchIncrementalPercentage)
	}
	if model.MaxParallelReadWriteFullPercentage != nil {
		modelMap["max_parallel_read_write_full_percentage"] = flex.IntValue(model.MaxParallelReadWriteFullPercentage)
	}
	if model.MaxParallelReadWriteIncrementalPercentage != nil {
		modelMap["max_parallel_read_write_incremental_percentage"] = flex.IntValue(model.MaxParallelReadWriteIncrementalPercentage)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigParamsToMap(model *backuprecoveryv1.StorageArraySnapshotConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsMaxSnapshotsConfigEnabled != nil {
		modelMap["is_max_snapshots_config_enabled"] = *model.IsMaxSnapshotsConfigEnabled
	}
	if model.IsMaxSpaceConfigEnabled != nil {
		modelMap["is_max_space_config_enabled"] = *model.IsMaxSpaceConfigEnabled
	}
	if model.StorageArraySnapshotMaxSpaceConfig != nil {
		storageArraySnapshotMaxSpaceConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigParamsToMap(model.StorageArraySnapshotMaxSpaceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigMap}
	}
	if model.StorageArraySnapshotThrottlingPolicies != nil {
		storageArraySnapshotThrottlingPolicies := []map[string]interface{}{}
		for _, storageArraySnapshotThrottlingPoliciesItem := range model.StorageArraySnapshotThrottlingPolicies {
			storageArraySnapshotThrottlingPoliciesItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPolicyToMap(&storageArraySnapshotThrottlingPoliciesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			storageArraySnapshotThrottlingPolicies = append(storageArraySnapshotThrottlingPolicies, storageArraySnapshotThrottlingPoliciesItemMap)
		}
		modelMap["storage_array_snapshot_throttling_policies"] = storageArraySnapshotThrottlingPolicies
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigParamsToMap(model *backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxSnapshotSpacePercentage != nil {
		modelMap["max_snapshot_space_percentage"] = flex.IntValue(model.MaxSnapshotSpacePercentage)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPolicyToMap(model *backuprecoveryv1.StorageArraySnapshotThrottlingPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = flex.IntValue(model.ID)
	}
	if model.IsMaxSnapshotsConfigEnabled != nil {
		modelMap["is_max_snapshots_config_enabled"] = *model.IsMaxSnapshotsConfigEnabled
	}
	if model.IsMaxSpaceConfigEnabled != nil {
		modelMap["is_max_space_config_enabled"] = *model.IsMaxSpaceConfigEnabled
	}
	if model.MaxSnapshotConfig != nil {
		maxSnapshotConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSnapshotConfigParamsToMap(model.MaxSnapshotConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigMap}
	}
	if model.MaxSpaceConfig != nil {
		maxSpaceConfigMap, err := DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigParamsToMap(model.MaxSpaceConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["max_space_config"] = []map[string]interface{}{maxSpaceConfigMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSnapshotConfigParamsToMap(model *backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxSnapshots != nil {
		modelMap["max_snapshots"] = flex.IntValue(model.MaxSnapshots)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverrideToMap(model *backuprecoveryv1.ThrottlingPolicyOverride) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatastoreID != nil {
		modelMap["datastore_id"] = flex.IntValue(model.DatastoreID)
	}
	if model.DatastoreName != nil {
		modelMap["datastore_name"] = *model.DatastoreName
	}
	if model.ThrottlingPolicy != nil {
		throttlingPolicyMap, err := DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyParametersToMap(model.ThrottlingPolicy)
		if err != nil {
			return modelMap, err
		}
		modelMap["throttling_policy"] = []map[string]interface{}{throttlingPolicyMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoUdaConnectParamsToMap(model *backuprecoveryv1.UdaConnectParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Capabilities != nil {
		capabilitiesMap, err := DataSourceIbmBackupRecoveryRegistrationInfoUdaSourceCapabilitiesToMap(model.Capabilities)
		if err != nil {
			return modelMap, err
		}
		modelMap["capabilities"] = []map[string]interface{}{capabilitiesMap}
	}
	if model.Credentials != nil {
		credentialsMap, err := DataSourceIbmBackupRecoveryRegistrationInfoCredentialsToMap(model.Credentials)
		if err != nil {
			return modelMap, err
		}
		modelMap["credentials"] = []map[string]interface{}{credentialsMap}
	}
	if model.EtEnableLogBackupPolicy != nil {
		modelMap["et_enable_log_backup_policy"] = *model.EtEnableLogBackupPolicy
	}
	if model.EtEnableRunNow != nil {
		modelMap["et_enable_run_now"] = *model.EtEnableRunNow
	}
	if model.HostType != nil {
		modelMap["host_type"] = *model.HostType
	}
	if model.Hosts != nil {
		modelMap["hosts"] = model.Hosts
	}
	if model.LiveDataView != nil {
		modelMap["live_data_view"] = *model.LiveDataView
	}
	if model.LiveLogView != nil {
		modelMap["live_log_view"] = *model.LiveLogView
	}
	if model.MountDir != nil {
		modelMap["mount_dir"] = *model.MountDir
	}
	if model.MountView != nil {
		modelMap["mount_view"] = *model.MountView
	}
	if model.ScriptDir != nil {
		modelMap["script_dir"] = *model.ScriptDir
	}
	if model.SourceArgs != nil {
		modelMap["source_args"] = *model.SourceArgs
	}
	if model.SourceRegistrationArguments != nil {
		sourceRegistrationArguments := []map[string]interface{}{}
		for _, sourceRegistrationArgumentsItem := range model.SourceRegistrationArguments {
			sourceRegistrationArgumentsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoKeyValueStrPairToMap(&sourceRegistrationArgumentsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			sourceRegistrationArguments = append(sourceRegistrationArguments, sourceRegistrationArgumentsItemMap)
		}
		modelMap["source_registration_arguments"] = sourceRegistrationArguments
	}
	if model.SourceType != nil {
		modelMap["source_type"] = *model.SourceType
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoUdaSourceCapabilitiesToMap(model *backuprecoveryv1.UdaSourceCapabilities) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AutoLogBackup != nil {
		modelMap["auto_log_backup"] = *model.AutoLogBackup
	}
	if model.DynamicConfig != nil {
		modelMap["dynamic_config"] = *model.DynamicConfig
	}
	if model.EntitySupport != nil {
		modelMap["entity_support"] = *model.EntitySupport
	}
	if model.EtLogBackup != nil {
		modelMap["et_log_backup"] = *model.EtLogBackup
	}
	if model.ExternalDisks != nil {
		modelMap["external_disks"] = *model.ExternalDisks
	}
	if model.FullBackup != nil {
		modelMap["full_backup"] = *model.FullBackup
	}
	if model.IncrBackup != nil {
		modelMap["incr_backup"] = *model.IncrBackup
	}
	if model.LogBackup != nil {
		modelMap["log_backup"] = *model.LogBackup
	}
	if model.MultiObjectRestore != nil {
		modelMap["multi_object_restore"] = *model.MultiObjectRestore
	}
	if model.PauseResumeBackup != nil {
		modelMap["pause_resume_backup"] = *model.PauseResumeBackup
	}
	if model.PostBackupJobScript != nil {
		modelMap["post_backup_job_script"] = *model.PostBackupJobScript
	}
	if model.PostRestoreJobScript != nil {
		modelMap["post_restore_job_script"] = *model.PostRestoreJobScript
	}
	if model.PreBackupJobScript != nil {
		modelMap["pre_backup_job_script"] = *model.PreBackupJobScript
	}
	if model.PreRestoreJobScript != nil {
		modelMap["pre_restore_job_script"] = *model.PreRestoreJobScript
	}
	if model.ResourceThrottling != nil {
		modelMap["resource_throttling"] = *model.ResourceThrottling
	}
	if model.SnapfsCert != nil {
		modelMap["snapfs_cert"] = *model.SnapfsCert
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoKeyValueStrPairToMap(model *backuprecoveryv1.KeyValueStrPair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = *model.Key
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoStatsToMap(model *backuprecoveryv1.ProtectionSourceTreeInfoStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.ProtectedSize != nil {
		modelMap["protected_size"] = flex.IntValue(model.ProtectedSize)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.UnprotectedSize != nil {
		modelMap["unprotected_size"] = flex.IntValue(model.UnprotectedSize)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryByEnvToMap(model *backuprecoveryv1.ProtectionSummaryByEnv) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Environment != nil {
		modelMap["environment"] = *model.Environment
	}
	if model.KubernetesDistributionStats != nil {
		kubernetesDistributionStats := []map[string]interface{}{}
		for _, kubernetesDistributionStatsItem := range model.KubernetesDistributionStats {
			kubernetesDistributionStatsItemMap, err := DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryForK8sDistributionsToMap(&kubernetesDistributionStatsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			kubernetesDistributionStats = append(kubernetesDistributionStats, kubernetesDistributionStatsItemMap)
		}
		modelMap["kubernetes_distribution_stats"] = kubernetesDistributionStats
	}
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.ProtectedSize != nil {
		modelMap["protected_size"] = flex.IntValue(model.ProtectedSize)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.UnprotectedSize != nil {
		modelMap["unprotected_size"] = flex.IntValue(model.UnprotectedSize)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryForK8sDistributionsToMap(model *backuprecoveryv1.ProtectionSummaryForK8sDistributions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Distribution != nil {
		modelMap["distribution"] = *model.Distribution
	}
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.ProtectedSize != nil {
		modelMap["protected_size"] = flex.IntValue(model.ProtectedSize)
	}
	if model.TotalRegisteredClusters != nil {
		modelMap["total_registered_clusters"] = flex.IntValue(model.TotalRegisteredClusters)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.UnprotectedSize != nil {
		modelMap["unprotected_size"] = flex.IntValue(model.UnprotectedSize)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryRegistrationInfoListProtectionSourcesRegistrationInfoResponseStatsToMap(model *backuprecoveryv1.GetRegistrationInfoResponseStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.ProtectedSize != nil {
		modelMap["protected_size"] = flex.IntValue(model.ProtectedSize)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.UnprotectedSize != nil {
		modelMap["unprotected_size"] = flex.IntValue(model.UnprotectedSize)
	}
	return modelMap, nil
}
