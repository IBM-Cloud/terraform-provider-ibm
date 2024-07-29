// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmSourceRegistration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSourceRegistrationRead,

		Schema: map[string]*schema.Schema{
			"source_registration_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the id of the Protection Source registration.",
			},
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"source_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of top level source object discovered after the registration.",
			},
			"source_info": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies information about an object.",
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
			"environment": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the environment type of the Protection Source.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user specified name for this source.",
			},
			"connection_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field.",
			},
			"connections": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specfies the list of connections for the source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the id of the connection.",
						},
						"entity_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the entity id of the source. The source can a non-root entity.",
						},
						"connector_group_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the connector group id of connector groups.",
						},
					},
				},
			},
			"connector_group_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the connector group id of connector groups.",
			},
			"advanced_configs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the advanced configuration for a protection source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "key.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "value.",
						},
					},
				},
			},
			"authentication_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress.",
			},
			"registration_time_msecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the time when the source was registered in milliseconds.",
			},
			"last_refreshed_time_msecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the time when the source was last refreshed in milliseconds.",
			},
			"physical_params": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Physical Params params.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the endpoint IPaddress, URL or hostname of the physical host.",
						},
						"force_register": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster.",
						},
						"host_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of host.",
						},
						"physical_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of physical server.",
						},
						"applications": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of applications to be registered with Physical Source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"oracle_params": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Physical Params params.",
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
		},
	}
}

func dataSourceIbmSourceRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

	getProtectionSourceRegistrationOptions.SetID(int64(d.Get("source_registration_id").(int)))
	if _, ok := d.GetOk("request_initiator_type"); ok {
		getProtectionSourceRegistrationOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}

	sourceRegistrationReponseParams, response, err := backupRecoveryClient.GetProtectionSourceRegistrationWithContext(context, getProtectionSourceRegistrationOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(strconv.Itoa(int(*sourceRegistrationReponseParams.ID)))

	if err = d.Set("source_id", flex.IntValue(sourceRegistrationReponseParams.SourceID)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_id: %s", err))
	}

	sourceInfo := []map[string]interface{}{}
	if sourceRegistrationReponseParams.SourceInfo != nil {
		modelMap, err := dataSourceIbmSourceRegistrationObjectToMap(sourceRegistrationReponseParams.SourceInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		sourceInfo = append(sourceInfo, modelMap)
	}
	if err = d.Set("source_info", sourceInfo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_info %s", err))
	}

	if err = d.Set("environment", sourceRegistrationReponseParams.Environment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
	}

	if err = d.Set("name", sourceRegistrationReponseParams.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("connection_id", flex.IntValue(sourceRegistrationReponseParams.ConnectionID)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting connection_id: %s", err))
	}

	connections := []map[string]interface{}{}
	if sourceRegistrationReponseParams.Connections != nil {
		for _, modelItem := range sourceRegistrationReponseParams.Connections {
			modelMap, err := dataSourceIbmSourceRegistrationConnectionConfigToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			connections = append(connections, modelMap)
		}
	}
	if err = d.Set("connections", connections); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting connections %s", err))
	}

	if err = d.Set("connector_group_id", flex.IntValue(sourceRegistrationReponseParams.ConnectorGroupID)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting connector_group_id: %s", err))
	}

	advancedConfigs := []map[string]interface{}{}
	if sourceRegistrationReponseParams.AdvancedConfigs != nil {
		for _, modelItem := range sourceRegistrationReponseParams.AdvancedConfigs {
			modelMap, err := dataSourceIbmSourceRegistrationKeyValuePairToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, modelMap)
		}
	}
	if err = d.Set("advanced_configs", advancedConfigs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting advanced_configs %s", err))
	}

	if err = d.Set("authentication_status", sourceRegistrationReponseParams.AuthenticationStatus); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting authentication_status: %s", err))
	}

	if err = d.Set("registration_time_msecs", flex.IntValue(sourceRegistrationReponseParams.RegistrationTimeMsecs)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting registration_time_msecs: %s", err))
	}

	if err = d.Set("last_refreshed_time_msecs", flex.IntValue(sourceRegistrationReponseParams.LastRefreshedTimeMsecs)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_refreshed_time_msecs: %s", err))
	}

	physicalParams := []map[string]interface{}{}
	if sourceRegistrationReponseParams.PhysicalParams != nil {
		modelMap, err := dataSourceIbmSourceRegistrationPhysicalParamsToMap(sourceRegistrationReponseParams.PhysicalParams)
		if err != nil {
			return diag.FromErr(err)
		}
		physicalParams = append(physicalParams, modelMap)
	}
	if err = d.Set("physical_params", physicalParams); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting physical_params %s", err))
	}

	oracleParams := []map[string]interface{}{}
	if sourceRegistrationReponseParams.OracleParams != nil {
		modelMap, err := dataSourceIbmSourceRegistrationOracleParamsToMap(sourceRegistrationReponseParams.OracleParams)
		if err != nil {
			return diag.FromErr(err)
		}
		oracleParams = append(oracleParams, modelMap)
	}
	if err = d.Set("oracle_params", oracleParams); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting oracle_params %s", err))
	}

	return nil
}

func dataSourceIbmSourceRegistrationObjectToMap(model *backuprecoveryv1.Object) (map[string]interface{}, error) {
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
			protectionStatsItemMap, err := dataSourceIbmSourceRegistrationObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := dataSourceIbmSourceRegistrationPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmSourceRegistrationObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmSourceRegistrationObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
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

func dataSourceIbmSourceRegistrationPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := dataSourceIbmSourceRegistrationUserToMap(&usersItem)
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
			groupsItemMap, err := dataSourceIbmSourceRegistrationGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := dataSourceIbmSourceRegistrationTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
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

func dataSourceIbmSourceRegistrationGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
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

func dataSourceIbmSourceRegistrationTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationObjectOracleParamsToMap(model *backuprecoveryv1.ObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmSourceRegistrationDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmSourceRegistrationHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := dataSourceIbmSourceRegistrationPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := dataSourceIbmSourceRegistrationDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
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

func dataSourceIbmSourceRegistrationObjectPhysicalParamsToMap(model *backuprecoveryv1.ObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationConnectionConfigToMap(model *backuprecoveryv1.ConnectionConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConnectionID != nil {
		modelMap["connection_id"] = flex.IntValue(model.ConnectionID)
	}
	if model.EntityID != nil {
		modelMap["entity_id"] = flex.IntValue(model.EntityID)
	}
	if model.ConnectorGroupID != nil {
		modelMap["connector_group_id"] = flex.IntValue(model.ConnectorGroupID)
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationPhysicalParamsToMap(model *backuprecoveryv1.PhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["endpoint"] = model.Endpoint
	if model.ForceRegister != nil {
		modelMap["force_register"] = model.ForceRegister
	}
	if model.HostType != nil {
		modelMap["host_type"] = model.HostType
	}
	if model.PhysicalType != nil {
		modelMap["physical_type"] = model.PhysicalType
	}
	if model.Applications != nil {
		modelMap["applications"] = model.Applications
	}
	return modelMap, nil
}

func dataSourceIbmSourceRegistrationOracleParamsToMap(model *backuprecoveryv1.OracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmSourceRegistrationDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmSourceRegistrationHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}
