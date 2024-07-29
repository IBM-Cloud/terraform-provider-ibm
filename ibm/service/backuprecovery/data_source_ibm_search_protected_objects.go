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

func DataSourceIbmSearchProtectedObjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSearchProtectedObjectsRead,

		Schema: map[string]*schema.Schema{
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"search_string": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria.",
			},
			"environments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the environment type to filter objects.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snapshot_actions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_action_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which objects are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will include Objects which belongs to all tenants which the current user has permission to see.",
			},
			"protection_group_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of Object ids to filter.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"storage_domain_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the Storage Domain ids to filter objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"sub_result_size": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the size of objects to be fetched for a single subresult.",
			},
			"filter_snapshot_from_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value.",
			},
			"filter_snapshot_to_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value.",
			},
			"os_types": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the operating system types to filter objects on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"run_instance_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"cdp_protected_only": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to only return the CDP protected objects.",
			},
			"region_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of region ids. Only records from clusters having these region ids will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_cached_data": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source.",
			},
			"objects": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Protected Objects.",
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
						"protection_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the count and size of protected and unprotected objects for the size.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the environment of the object.",
									},
									"protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of the protected leaf objects.",
									},
									"unprotected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of the unprotected leaf objects.",
									},
									"deleted_protected_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of protected leaf objects which were deleted from the source after being protected.",
									},
									"protected_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the protected logical size in bytes.",
									},
									"unprotected_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the unprotected logical size in bytes.",
									},
								},
							},
						},
						"permissions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of users, groups and users that have permissions for a given object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the id of the object.",
									},
									"users": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of users which has the permissions to the object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the user.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the sid of the user.",
												},
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the domain of the user.",
												},
											},
										},
									},
									"groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the list of user groups which has permissions to the object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the user group.",
												},
												"sid": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the sid of the user group.",
												},
												"domain": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the domain of the user group.",
												},
											},
										},
									},
									"tenant": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a tenant object.",
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
								},
							},
						},
						"oracle_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the parameters for Oracle object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database_entity_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Object details about Oracle database entity info.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"container_database_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the list of Pluggable databases within a container database.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"database_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the database Id of the Pluggable DB.",
															},
															"database_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the Pluggable DB.",
															},
														},
													},
												},
												"data_guard_info": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Dataguard info about Oracle database.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"role": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the role of the Oracle DataGuard database.",
															},
															"standby_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the type of the standby oracle database.",
															},
														},
													},
												},
											},
										},
									},
									"host_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the id of the host object.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the name of the host object.",
												},
												"environment": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the environment of the object.",
												},
											},
										},
									},
								},
							},
						},
						"physical_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the parameters for Physical object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_system_backup": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies if system backup was enabled for the source in a particular run.",
									},
								},
							},
						},
						"source_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the Source Object information.",
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
						"latest_snapshots_info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the latest snapshot information for every Protection Group for a given object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_snapshot_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the local snapshot information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"snapshot_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the id of the local snapshot for the object.",
												},
												"logical_size_bytes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the logical size of this snapshot in bytes.",
												},
											},
										},
									},
									"archival_snapshots_info": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the archival snapshots information.",
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
												"snapshot_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the id of the archival snapshot for the object.",
												},
												"logical_size_bytes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the logical size of this snapshot in bytes.",
												},
											},
										},
									},
									"indexing_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the indexing status of objects in this snapshot.<br> 'InProgress' indicates the indexing is in progress.<br> 'Done' indicates indexing is done.<br> 'NoIndex' indicates indexing is not applicable.<br> 'Error' indicates indexing failed with error.",
									},
									"protection_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies id of the Protection Group.",
									},
									"protection_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the Protection Group.",
									},
									"run_instance_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the instance id of the protection run which create the snapshot.",
									},
									"source_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the source protection group id in case of replication.",
									},
									"storage_domain_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the Storage Domain id where the backup data of Object is present.",
									},
									"storage_domain_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of Storage Domain id where the backup data of Object is present.",
									},
									"protection_run_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the id of Protection Group Run.",
									},
									"run_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of protection run created this snapshot.",
									},
									"protection_run_start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of Protection Group Run in Unix timestamp epoch in microseconds.",
									},
									"protection_run_end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of Protection Group Run in Unix timestamp epoch in microseconds.",
									},
								},
							},
						},
					},
				},
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the metadata information about the Protection Groups, Protection Policy etc., for search result.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unique_protection_group_identifiers": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of unique Protection Group identifiers for all the Objects returned in the response.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protection_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Protection Group id.",
									},
									"protection_group_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Protection Group name.",
									},
								},
							},
						},
					},
				},
			},
			"num_results": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the total number of search results which matches the search criteria.",
			},
		},
	}
}

func dataSourceIbmSearchProtectedObjectsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	searchProtectedObjectsOptions := &backuprecoveryv1.SearchProtectedObjectsOptions{}

	if _, ok := d.GetOk("request_initiator_type"); ok {
		searchProtectedObjectsOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("search_string"); ok {
		searchProtectedObjectsOptions.SetSearchString(d.Get("search_string").(string))
	}
	if _, ok := d.GetOk("environments"); ok {
		var environments []string
		for _, v := range d.Get("environments").([]interface{}) {
			environmentsItem := v.(string)
			environments = append(environments, environmentsItem)
		}
		searchProtectedObjectsOptions.SetEnvironments(environments)
	}
	if _, ok := d.GetOk("snapshot_actions"); ok {
		var snapshotActions []string
		for _, v := range d.Get("snapshot_actions").([]interface{}) {
			snapshotActionsItem := v.(string)
			snapshotActions = append(snapshotActions, snapshotActionsItem)
		}
		searchProtectedObjectsOptions.SetSnapshotActions(snapshotActions)
	}
	if _, ok := d.GetOk("object_action_key"); ok {
		searchProtectedObjectsOptions.SetObjectActionKey(d.Get("object_action_key").(string))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		searchProtectedObjectsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		searchProtectedObjectsOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("protection_group_ids"); ok {
		var protectionGroupIds []string
		for _, v := range d.Get("protection_group_ids").([]interface{}) {
			protectionGroupIdsItem := v.(string)
			protectionGroupIds = append(protectionGroupIds, protectionGroupIdsItem)
		}
		searchProtectedObjectsOptions.SetProtectionGroupIds(protectionGroupIds)
	}
	if _, ok := d.GetOk("object_ids"); ok {
		var objectIds []int64
		for _, v := range d.Get("object_ids").([]interface{}) {
			objectIdsItem := int64(v.(int))
			objectIds = append(objectIds, objectIdsItem)
		}
		searchProtectedObjectsOptions.SetObjectIds(objectIds)
	}
	if _, ok := d.GetOk("storage_domain_ids"); ok {
		var storageDomainIds []int64
		for _, v := range d.Get("storage_domain_ids").([]interface{}) {
			storageDomainIdsItem := int64(v.(int))
			storageDomainIds = append(storageDomainIds, storageDomainIdsItem)
		}
		searchProtectedObjectsOptions.SetStorageDomainIds(storageDomainIds)
	}
	if _, ok := d.GetOk("sub_result_size"); ok {
		searchProtectedObjectsOptions.SetSubResultSize(int64(d.Get("sub_result_size").(int)))
	}
	if _, ok := d.GetOk("filter_snapshot_from_usecs"); ok {
		searchProtectedObjectsOptions.SetFilterSnapshotFromUsecs(int64(d.Get("filter_snapshot_from_usecs").(int)))
	}
	if _, ok := d.GetOk("filter_snapshot_to_usecs"); ok {
		searchProtectedObjectsOptions.SetFilterSnapshotToUsecs(int64(d.Get("filter_snapshot_to_usecs").(int)))
	}
	if _, ok := d.GetOk("os_types"); ok {
		var osTypes []string
		for _, v := range d.Get("os_types").([]interface{}) {
			osTypesItem := v.(string)
			osTypes = append(osTypes, osTypesItem)
		}
		searchProtectedObjectsOptions.SetOsTypes(osTypes)
	}
	if _, ok := d.GetOk("source_ids"); ok {
		var sourceIds []int64
		for _, v := range d.Get("source_ids").([]interface{}) {
			sourceIdsItem := int64(v.(int))
			sourceIds = append(sourceIds, sourceIdsItem)
		}
		searchProtectedObjectsOptions.SetSourceIds(sourceIds)
	}
	if _, ok := d.GetOk("run_instance_ids"); ok {
		var runInstanceIds []int64
		for _, v := range d.Get("run_instance_ids").([]interface{}) {
			runInstanceIdsItem := int64(v.(int))
			runInstanceIds = append(runInstanceIds, runInstanceIdsItem)
		}
		searchProtectedObjectsOptions.SetRunInstanceIds(runInstanceIds)
	}
	if _, ok := d.GetOk("cdp_protected_only"); ok {
		searchProtectedObjectsOptions.SetCdpProtectedOnly(d.Get("cdp_protected_only").(bool))
	}
	if _, ok := d.GetOk("region_ids"); ok {
		var regionIds []string
		for _, v := range d.Get("region_ids").([]interface{}) {
			regionIdsItem := v.(string)
			regionIds = append(regionIds, regionIdsItem)
		}
		searchProtectedObjectsOptions.SetRegionIds(regionIds)
	}
	if _, ok := d.GetOk("use_cached_data"); ok {
		searchProtectedObjectsOptions.SetUseCachedData(d.Get("use_cached_data").(bool))
	}

	protectedObjectsSearchResponse, response, err := backupRecoveryClient.SearchProtectedObjectsWithContext(context, searchProtectedObjectsOptions)
	if err != nil {
		log.Printf("[DEBUG] SearchProtectedObjectsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("SearchProtectedObjectsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmSearchProtectedObjectsID(d))

	objects := []map[string]interface{}{}
	if protectedObjectsSearchResponse.Objects != nil {
		for _, modelItem := range protectedObjectsSearchResponse.Objects {
			modelMap, err := dataSourceIbmSearchProtectedObjectsProtectedObjectToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			objects = append(objects, modelMap)
		}
	}
	if err = d.Set("objects", objects); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting objects %s", err))
	}

	metadata := []map[string]interface{}{}
	if protectedObjectsSearchResponse.Metadata != nil {
		modelMap, err := dataSourceIbmSearchProtectedObjectsProtectedObjectsSearchResponseMetadataToMap(protectedObjectsSearchResponse.Metadata)
		if err != nil {
			return diag.FromErr(err)
		}
		metadata = append(metadata, modelMap)
	}
	if err = d.Set("metadata", metadata); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting metadata %s", err))
	}

	if err = d.Set("num_results", flex.IntValue(protectedObjectsSearchResponse.NumResults)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting num_results: %s", err))
	}

	return nil
}

// dataSourceIbmSearchProtectedObjectsID returns a reasonable ID for the list.
func dataSourceIbmSearchProtectedObjectsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSearchProtectedObjectsProtectedObjectToMap(model *backuprecoveryv1.ProtectedObject) (map[string]interface{}, error) {
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
	if model.ProtectionStats != nil {
		protectionStats := []map[string]interface{}{}
		for _, protectionStatsItem := range model.ProtectionStats {
			protectionStatsItemMap, err := dataSourceIbmSearchProtectedObjectsObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := dataSourceIbmSearchProtectedObjectsPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmSearchProtectedObjectsProtectedObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmSearchProtectedObjectsProtectedObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.SourceInfo != nil {
		sourceInfoMap, err := dataSourceIbmSearchProtectedObjectsProtectedObjectSourceInfoToMap(model.SourceInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["source_info"] = []map[string]interface{}{sourceInfoMap}
	}
	if model.LatestSnapshotsInfo != nil {
		latestSnapshotsInfo := []map[string]interface{}{}
		for _, latestSnapshotsInfoItem := range model.LatestSnapshotsInfo {
			latestSnapshotsInfoItemMap, err := dataSourceIbmSearchProtectedObjectsObjectSnapshotsInfoToMap(&latestSnapshotsInfoItem)
			if err != nil {
				return modelMap, err
			}
			latestSnapshotsInfo = append(latestSnapshotsInfo, latestSnapshotsInfoItemMap)
		}
		modelMap["latest_snapshots_info"] = latestSnapshotsInfo
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.ProtectedCount != nil {
		modelMap["protected_count"] = flex.IntValue(model.ProtectedCount)
	}
	if model.UnprotectedCount != nil {
		modelMap["unprotected_count"] = flex.IntValue(model.UnprotectedCount)
	}
	if model.DeletedProtectedCount != nil {
		modelMap["deleted_protected_count"] = flex.IntValue(model.DeletedProtectedCount)
	}
	if model.ProtectedSizeBytes != nil {
		modelMap["protected_size_bytes"] = flex.IntValue(model.ProtectedSizeBytes)
	}
	if model.UnprotectedSizeBytes != nil {
		modelMap["unprotected_size_bytes"] = flex.IntValue(model.UnprotectedSizeBytes)
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := dataSourceIbmSearchProtectedObjectsUserToMap(&usersItem)
			if err != nil {
				return modelMap, err
			}
			users = append(users, usersItemMap)
		}
		modelMap["users"] = users
	}
	if model.Groups != nil {
		groups := []map[string]interface{}{}
		for _, groupsItem := range model.Groups {
			groupsItemMap, err := dataSourceIbmSearchProtectedObjectsGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := dataSourceIbmSearchProtectedObjectsTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = model.Domain
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Sid != nil {
		modelMap["sid"] = model.Sid
	}
	if model.Domain != nil {
		modelMap["domain"] = model.Domain
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsProtectedObjectOracleParamsToMap(model *backuprecoveryv1.ProtectedObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmSearchProtectedObjectsDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmSearchProtectedObjectsHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := dataSourceIbmSearchProtectedObjectsPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := dataSourceIbmSearchProtectedObjectsDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsProtectedObjectPhysicalParamsToMap(model *backuprecoveryv1.ProtectedObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsProtectedObjectSourceInfoToMap(model *backuprecoveryv1.ProtectedObjectSourceInfo) (map[string]interface{}, error) {
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

func dataSourceIbmSearchProtectedObjectsObjectSnapshotsInfoToMap(model *backuprecoveryv1.ObjectSnapshotsInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LocalSnapshotInfo != nil {
		localSnapshotInfoMap, err := dataSourceIbmSearchProtectedObjectsObjectSnapshotsInfoLocalSnapshotInfoToMap(model.LocalSnapshotInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["local_snapshot_info"] = []map[string]interface{}{localSnapshotInfoMap}
	}
	if model.ArchivalSnapshotsInfo != nil {
		archivalSnapshotsInfo := []map[string]interface{}{}
		for _, archivalSnapshotsInfoItem := range model.ArchivalSnapshotsInfo {
			archivalSnapshotsInfoItemMap, err := dataSourceIbmSearchProtectedObjectsObjectArchivalSnapshotInfoToMap(&archivalSnapshotsInfoItem)
			if err != nil {
				return modelMap, err
			}
			archivalSnapshotsInfo = append(archivalSnapshotsInfo, archivalSnapshotsInfoItemMap)
		}
		modelMap["archival_snapshots_info"] = archivalSnapshotsInfo
	}
	if model.IndexingStatus != nil {
		modelMap["indexing_status"] = model.IndexingStatus
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.RunInstanceID != nil {
		modelMap["run_instance_id"] = flex.IntValue(model.RunInstanceID)
	}
	if model.SourceGroupID != nil {
		modelMap["source_group_id"] = model.SourceGroupID
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = flex.IntValue(model.StorageDomainID)
	}
	if model.StorageDomainName != nil {
		modelMap["storage_domain_name"] = model.StorageDomainName
	}
	if model.ProtectionRunID != nil {
		modelMap["protection_run_id"] = model.ProtectionRunID
	}
	if model.RunType != nil {
		modelMap["run_type"] = model.RunType
	}
	if model.ProtectionRunStartTimeUsecs != nil {
		modelMap["protection_run_start_time_usecs"] = flex.IntValue(model.ProtectionRunStartTimeUsecs)
	}
	if model.ProtectionRunEndTimeUsecs != nil {
		modelMap["protection_run_end_time_usecs"] = flex.IntValue(model.ProtectionRunEndTimeUsecs)
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsObjectSnapshotsInfoLocalSnapshotInfoToMap(model *backuprecoveryv1.ObjectSnapshotsInfoLocalSnapshotInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SnapshotID != nil {
		modelMap["snapshot_id"] = model.SnapshotID
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsObjectArchivalSnapshotInfoToMap(model *backuprecoveryv1.ObjectArchivalSnapshotInfo) (map[string]interface{}, error) {
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
		tierSettingsMap, err := dataSourceIbmSearchProtectedObjectsArchivalTargetTierInfoToMap(model.TierSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["tier_settings"] = []map[string]interface{}{tierSettingsMap}
	}
	if model.SnapshotID != nil {
		modelMap["snapshot_id"] = model.SnapshotID
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsArchivalTargetTierInfoToMap(model *backuprecoveryv1.ArchivalTargetTierInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cloud_platform"] = model.CloudPlatform
	if model.OracleTiering != nil {
		oracleTieringMap, err := dataSourceIbmSearchProtectedObjectsOracleTiersToMap(model.OracleTiering)
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

func dataSourceIbmSearchProtectedObjectsOracleTiersToMap(model *backuprecoveryv1.OracleTiers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	tiers := []map[string]interface{}{}
	for _, tiersItem := range model.Tiers {
		tiersItemMap, err := dataSourceIbmSearchProtectedObjectsOracleTierToMap(&tiersItem)
		if err != nil {
			return modelMap, err
		}
		tiers = append(tiers, tiersItemMap)
	}
	modelMap["tiers"] = tiers
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsOracleTierToMap(model *backuprecoveryv1.OracleTier) (map[string]interface{}, error) {
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

func dataSourceIbmSearchProtectedObjectsProtectedObjectsSearchResponseMetadataToMap(model *backuprecoveryv1.ProtectedObjectsSearchResponseMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UniqueProtectionGroupIdentifiers != nil {
		uniqueProtectionGroupIdentifiers := []map[string]interface{}{}
		for _, uniqueProtectionGroupIdentifiersItem := range model.UniqueProtectionGroupIdentifiers {
			uniqueProtectionGroupIdentifiersItemMap, err := dataSourceIbmSearchProtectedObjectsProtectionGroupIdentifierToMap(&uniqueProtectionGroupIdentifiersItem)
			if err != nil {
				return modelMap, err
			}
			uniqueProtectionGroupIdentifiers = append(uniqueProtectionGroupIdentifiers, uniqueProtectionGroupIdentifiersItemMap)
		}
		modelMap["unique_protection_group_identifiers"] = uniqueProtectionGroupIdentifiers
	}
	return modelMap, nil
}

func dataSourceIbmSearchProtectedObjectsProtectionGroupIdentifierToMap(model *backuprecoveryv1.ProtectionGroupIdentifier) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	return modelMap, nil
}
