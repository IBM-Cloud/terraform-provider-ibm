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

func DataSourceIbmSearchObjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSearchObjectsRead,

		Schema: map[string]*schema.Schema{
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"search_string": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria.",
			},
			"environments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the environment type to filter objects.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protection_types": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the protection type to filter objects.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"source_uuids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_protected": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned.",
			},
			"is_deleted": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned.",
			},
			"last_run_status_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of region ids. Only records from clusters having these region ids will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"storage_domain_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of storage domain ids. Format is clusterId:clusterIncarnationId:storageDomainId. Only objects having protection in these storage domains will be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_deleted_objects": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated.",
			},
			"pagination_cookie": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the pagination cookie with which subsequent parts of the response can be fetched.",
			},
			"count": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the number of objects to be fetched for the specified pagination cookie.",
			},
			"must_have_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies tags which must be all present in the document.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"might_have_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"must_have_snapshot_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies snapshot tags which must be all present in the document.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"might_have_snapshot_tag_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"objects": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Objects.",
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
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies tag applied to the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Id of tag applied to the object.",
									},
								},
							},
						},
						"snapshot_tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies snapshot tags applied to the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Id of tag applied to the object.",
									},
									"run_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies runs the tags are applied to.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
								},
							},
						},
						"object_protection_infos": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the object info on each cluster.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the object id.",
									},
									"source_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the source id.",
									},
									"view_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the view id for the object.",
									},
									"region_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the region id where this object belongs to.",
									},
									"cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the cluster id where this object belongs to.",
									},
									"cluster_incarnation_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the cluster incarnation id where this object belongs to.",
									},
									"tenant_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of Tenants the object belongs to.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_deleted": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the object is deleted. Deleted objects can't be protected but can be recovered or unprotected.",
									},
									"protection_groups": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a list of protection groups protecting this object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group name.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection group id.",
												},
												"protection_env_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the protection type of the job if any.",
												},
												"policy_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the policy name for this group.",
												},
												"policy_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the policy id for this group.",
												},
												"storage_domain_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the storage domain id of this group. Format is clusterId:clusterIncarnationId:storageDomainId.",
												},
												"last_backup_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last local back up run.",
												},
												"last_archival_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last archival run.",
												},
												"last_replication_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last replication run.",
												},
												"last_run_sla_violated": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if the sla is violated in last run.",
												},
											},
										},
									},
									"object_backup_configuration": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies a list of object protections.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"policy_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the policy name for this group.",
												},
												"policy_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the policy id for this protection.",
												},
												"storage_domain_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the storage domain id of this protection. Format is clusterId:clusterIncarnationId:storageDomainId.",
												},
												"last_backup_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last local back up run.",
												},
												"last_archival_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last archival run.",
												},
												"last_replication_run_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the status of last replication run.",
												},
												"last_run_sla_violated": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Specifies if the sla is violated in last run.",
												},
											},
										},
									},
									"last_run_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the status of the object's last protection run.",
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

func dataSourceIbmSearchObjectsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	searchObjectsOptions := &backuprecoveryv1.SearchObjectsOptions{}

	if _, ok := d.GetOk("request_initiator_type"); ok {
		searchObjectsOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("search_string"); ok {
		searchObjectsOptions.SetSearchString(d.Get("search_string").(string))
	}
	if _, ok := d.GetOk("environments"); ok {
		var environments []string
		for _, v := range d.Get("environments").([]interface{}) {
			environmentsItem := v.(string)
			environments = append(environments, environmentsItem)
		}
		searchObjectsOptions.SetEnvironments(environments)
	}
	if _, ok := d.GetOk("protection_types"); ok {
		var protectionTypes []string
		for _, v := range d.Get("protection_types").([]interface{}) {
			protectionTypesItem := v.(string)
			protectionTypes = append(protectionTypes, protectionTypesItem)
		}
		searchObjectsOptions.SetProtectionTypes(protectionTypes)
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		searchObjectsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		searchObjectsOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("protection_group_ids"); ok {
		var protectionGroupIds []string
		for _, v := range d.Get("protection_group_ids").([]interface{}) {
			protectionGroupIdsItem := v.(string)
			protectionGroupIds = append(protectionGroupIds, protectionGroupIdsItem)
		}
		searchObjectsOptions.SetProtectionGroupIds(protectionGroupIds)
	}
	if _, ok := d.GetOk("object_ids"); ok {
		var objectIds []int64
		for _, v := range d.Get("object_ids").([]interface{}) {
			objectIdsItem := int64(v.(int))
			objectIds = append(objectIds, objectIdsItem)
		}
		searchObjectsOptions.SetObjectIds(objectIds)
	}
	if _, ok := d.GetOk("os_types"); ok {
		var osTypes []string
		for _, v := range d.Get("os_types").([]interface{}) {
			osTypesItem := v.(string)
			osTypes = append(osTypes, osTypesItem)
		}
		searchObjectsOptions.SetOsTypes(osTypes)
	}
	if _, ok := d.GetOk("source_ids"); ok {
		var sourceIds []int64
		for _, v := range d.Get("source_ids").([]interface{}) {
			sourceIdsItem := int64(v.(int))
			sourceIds = append(sourceIds, sourceIdsItem)
		}
		searchObjectsOptions.SetSourceIds(sourceIds)
	}
	if _, ok := d.GetOk("source_uuids"); ok {
		var sourceUUIDs []string
		for _, v := range d.Get("source_uuids").([]interface{}) {
			sourceUUIDsItem := v.(string)
			sourceUUIDs = append(sourceUUIDs, sourceUUIDsItem)
		}
		searchObjectsOptions.SetSourceUUIDs(sourceUUIDs)
	}
	if _, ok := d.GetOk("is_protected"); ok {
		searchObjectsOptions.SetIsProtected(d.Get("is_protected").(bool))
	}
	if _, ok := d.GetOk("is_deleted"); ok {
		searchObjectsOptions.SetIsDeleted(d.Get("is_deleted").(bool))
	}
	if _, ok := d.GetOk("last_run_status_list"); ok {
		var lastRunStatusList []string
		for _, v := range d.Get("last_run_status_list").([]interface{}) {
			lastRunStatusListItem := v.(string)
			lastRunStatusList = append(lastRunStatusList, lastRunStatusListItem)
		}
		searchObjectsOptions.SetLastRunStatusList(lastRunStatusList)
	}
	if _, ok := d.GetOk("region_ids"); ok {
		var regionIds []string
		for _, v := range d.Get("region_ids").([]interface{}) {
			regionIdsItem := v.(string)
			regionIds = append(regionIds, regionIdsItem)
		}
		searchObjectsOptions.SetRegionIds(regionIds)
	}
	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []string
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			clusterIdentifiersItem := v.(string)
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItem)
		}
		searchObjectsOptions.SetClusterIdentifiers(clusterIdentifiers)
	}
	if _, ok := d.GetOk("storage_domain_ids"); ok {
		var storageDomainIds []string
		for _, v := range d.Get("storage_domain_ids").([]interface{}) {
			storageDomainIdsItem := v.(string)
			storageDomainIds = append(storageDomainIds, storageDomainIdsItem)
		}
		searchObjectsOptions.SetStorageDomainIds(storageDomainIds)
	}
	if _, ok := d.GetOk("include_deleted_objects"); ok {
		searchObjectsOptions.SetIncludeDeletedObjects(d.Get("include_deleted_objects").(bool))
	}
	if _, ok := d.GetOk("pagination_cookie"); ok {
		searchObjectsOptions.SetPaginationCookie(d.Get("pagination_cookie").(string))
	}
	if _, ok := d.GetOk("count"); ok {
		searchObjectsOptions.SetCount(int64(d.Get("count").(int)))
	}
	if _, ok := d.GetOk("must_have_tag_ids"); ok {
		var mustHaveTagIds []string
		for _, v := range d.Get("must_have_tag_ids").([]interface{}) {
			mustHaveTagIdsItem := v.(string)
			mustHaveTagIds = append(mustHaveTagIds, mustHaveTagIdsItem)
		}
		searchObjectsOptions.SetMustHaveTagIds(mustHaveTagIds)
	}
	if _, ok := d.GetOk("might_have_tag_ids"); ok {
		var mightHaveTagIds []string
		for _, v := range d.Get("might_have_tag_ids").([]interface{}) {
			mightHaveTagIdsItem := v.(string)
			mightHaveTagIds = append(mightHaveTagIds, mightHaveTagIdsItem)
		}
		searchObjectsOptions.SetMightHaveTagIds(mightHaveTagIds)
	}
	if _, ok := d.GetOk("must_have_snapshot_tag_ids"); ok {
		var mustHaveSnapshotTagIds []string
		for _, v := range d.Get("must_have_snapshot_tag_ids").([]interface{}) {
			mustHaveSnapshotTagIdsItem := v.(string)
			mustHaveSnapshotTagIds = append(mustHaveSnapshotTagIds, mustHaveSnapshotTagIdsItem)
		}
		searchObjectsOptions.SetMustHaveSnapshotTagIds(mustHaveSnapshotTagIds)
	}
	if _, ok := d.GetOk("might_have_snapshot_tag_ids"); ok {
		var mightHaveSnapshotTagIds []string
		for _, v := range d.Get("might_have_snapshot_tag_ids").([]interface{}) {
			mightHaveSnapshotTagIdsItem := v.(string)
			mightHaveSnapshotTagIds = append(mightHaveSnapshotTagIds, mightHaveSnapshotTagIdsItem)
		}
		searchObjectsOptions.SetMightHaveSnapshotTagIds(mightHaveSnapshotTagIds)
	}

	objectsSearchResponseBody, response, err := backupRecoveryClient.SearchObjectsWithContext(context, searchObjectsOptions)
	if err != nil {
		log.Printf("[DEBUG] SearchObjectsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("SearchObjectsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmSearchObjectsID(d))

	objects := []map[string]interface{}{}
	if objectsSearchResponseBody.Objects != nil {
		for _, modelItem := range objectsSearchResponseBody.Objects {
			modelMap, err := dataSourceIbmSearchObjectsSearchObjectToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			objects = append(objects, modelMap)
		}
	}
	if err = d.Set("objects", objects); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting objects %s", err))
	}

	return nil
}

// dataSourceIbmSearchObjectsID returns a reasonable ID for the list.
func dataSourceIbmSearchObjectsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSearchObjectsSearchObjectToMap(model *backuprecoveryv1.SearchObject) (map[string]interface{}, error) {
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
			protectionStatsItemMap, err := dataSourceIbmSearchObjectsObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := dataSourceIbmSearchObjectsPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmSearchObjectsSearchObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmSearchObjectsSearchObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.Tags != nil {
		tags := []map[string]interface{}{}
		for _, tagsItem := range model.Tags {
			tagsItemMap, err := dataSourceIbmSearchObjectsTagInfoToMap(&tagsItem)
			if err != nil {
				return modelMap, err
			}
			tags = append(tags, tagsItemMap)
		}
		modelMap["tags"] = tags
	}
	if model.SnapshotTags != nil {
		snapshotTags := []map[string]interface{}{}
		for _, snapshotTagsItem := range model.SnapshotTags {
			snapshotTagsItemMap, err := dataSourceIbmSearchObjectsSnapshotTagInfoToMap(&snapshotTagsItem)
			if err != nil {
				return modelMap, err
			}
			snapshotTags = append(snapshotTags, snapshotTagsItemMap)
		}
		modelMap["snapshot_tags"] = snapshotTags
	}
	if model.SourceInfo != nil {
		sourceInfoMap, err := dataSourceIbmSearchObjectsSearchObjectSourceInfoToMap(model.SourceInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["source_info"] = []map[string]interface{}{sourceInfoMap}
	}
	if model.ObjectProtectionInfos != nil {
		objectProtectionInfos := []map[string]interface{}{}
		for _, objectProtectionInfosItem := range model.ObjectProtectionInfos {
			objectProtectionInfosItemMap, err := dataSourceIbmSearchObjectsObjectProtectionInfoToMap(&objectProtectionInfosItem)
			if err != nil {
				return modelMap, err
			}
			objectProtectionInfos = append(objectProtectionInfos, objectProtectionInfosItemMap)
		}
		modelMap["object_protection_infos"] = objectProtectionInfos
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
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

func dataSourceIbmSearchObjectsPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := dataSourceIbmSearchObjectsUserToMap(&usersItem)
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
			groupsItemMap, err := dataSourceIbmSearchObjectsGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := dataSourceIbmSearchObjectsTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
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

func dataSourceIbmSearchObjectsGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
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

func dataSourceIbmSearchObjectsTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsSearchObjectOracleParamsToMap(model *backuprecoveryv1.SearchObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmSearchObjectsDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmSearchObjectsHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := dataSourceIbmSearchObjectsPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := dataSourceIbmSearchObjectsDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
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

func dataSourceIbmSearchObjectsSearchObjectPhysicalParamsToMap(model *backuprecoveryv1.SearchObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsTagInfoToMap(model *backuprecoveryv1.TagInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["tag_id"] = model.TagID
	return modelMap, nil
}

func dataSourceIbmSearchObjectsSnapshotTagInfoToMap(model *backuprecoveryv1.SnapshotTagInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["tag_id"] = model.TagID
	if model.RunIds != nil {
		modelMap["run_ids"] = model.RunIds
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsSearchObjectSourceInfoToMap(model *backuprecoveryv1.SearchObjectSourceInfo) (map[string]interface{}, error) {
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
			protectionStatsItemMap, err := dataSourceIbmSearchObjectsObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := dataSourceIbmSearchObjectsPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmSearchObjectsSearchObjectSourceInfoOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmSearchObjectsSearchObjectSourceInfoPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsSearchObjectSourceInfoOracleParamsToMap(model *backuprecoveryv1.SearchObjectSourceInfoOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmSearchObjectsDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmSearchObjectsHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsSearchObjectSourceInfoPhysicalParamsToMap(model *backuprecoveryv1.SearchObjectSourceInfoPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsObjectProtectionInfoToMap(model *backuprecoveryv1.ObjectProtectionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.SourceID != nil {
		modelMap["source_id"] = flex.IntValue(model.SourceID)
	}
	if model.ViewID != nil {
		modelMap["view_id"] = flex.IntValue(model.ViewID)
	}
	if model.RegionID != nil {
		modelMap["region_id"] = model.RegionID
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.TenantIds != nil {
		modelMap["tenant_ids"] = model.TenantIds
	}
	if model.IsDeleted != nil {
		modelMap["is_deleted"] = model.IsDeleted
	}
	if model.ProtectionGroups != nil {
		protectionGroups := []map[string]interface{}{}
		for _, protectionGroupsItem := range model.ProtectionGroups {
			protectionGroupsItemMap, err := dataSourceIbmSearchObjectsObjectProtectionGroupSummaryToMap(&protectionGroupsItem)
			if err != nil {
				return modelMap, err
			}
			protectionGroups = append(protectionGroups, protectionGroupsItemMap)
		}
		modelMap["protection_groups"] = protectionGroups
	}
	if model.ObjectBackupConfiguration != nil {
		objectBackupConfiguration := []map[string]interface{}{}
		for _, objectBackupConfigurationItem := range model.ObjectBackupConfiguration {
			objectBackupConfigurationItemMap, err := dataSourceIbmSearchObjectsProtectionSummaryToMap(&objectBackupConfigurationItem)
			if err != nil {
				return modelMap, err
			}
			objectBackupConfiguration = append(objectBackupConfiguration, objectBackupConfigurationItemMap)
		}
		modelMap["object_backup_configuration"] = objectBackupConfiguration
	}
	if model.LastRunStatus != nil {
		modelMap["last_run_status"] = model.LastRunStatus
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsObjectProtectionGroupSummaryToMap(model *backuprecoveryv1.ObjectProtectionGroupSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.ProtectionEnvType != nil {
		modelMap["protection_env_type"] = model.ProtectionEnvType
	}
	if model.PolicyName != nil {
		modelMap["policy_name"] = model.PolicyName
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = model.PolicyID
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = model.StorageDomainID
	}
	if model.LastBackupRunStatus != nil {
		modelMap["last_backup_run_status"] = model.LastBackupRunStatus
	}
	if model.LastArchivalRunStatus != nil {
		modelMap["last_archival_run_status"] = model.LastArchivalRunStatus
	}
	if model.LastReplicationRunStatus != nil {
		modelMap["last_replication_run_status"] = model.LastReplicationRunStatus
	}
	if model.LastRunSlaViolated != nil {
		modelMap["last_run_sla_violated"] = model.LastRunSlaViolated
	}
	return modelMap, nil
}

func dataSourceIbmSearchObjectsProtectionSummaryToMap(model *backuprecoveryv1.ProtectionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PolicyName != nil {
		modelMap["policy_name"] = model.PolicyName
	}
	if model.PolicyID != nil {
		modelMap["policy_id"] = model.PolicyID
	}
	if model.StorageDomainID != nil {
		modelMap["storage_domain_id"] = model.StorageDomainID
	}
	if model.LastBackupRunStatus != nil {
		modelMap["last_backup_run_status"] = model.LastBackupRunStatus
	}
	if model.LastArchivalRunStatus != nil {
		modelMap["last_archival_run_status"] = model.LastArchivalRunStatus
	}
	if model.LastReplicationRunStatus != nil {
		modelMap["last_replication_run_status"] = model.LastReplicationRunStatus
	}
	if model.LastRunSlaViolated != nil {
		modelMap["last_run_sla_violated"] = model.LastRunSlaViolated
	}
	return modelMap, nil
}
