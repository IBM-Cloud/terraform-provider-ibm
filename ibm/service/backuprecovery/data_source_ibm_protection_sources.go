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

func DataSourceIbmProtectionSources() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionSourcesRead,

		Schema: map[string]*schema.Schema{
			"request_initiator_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which Sources are to be returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned.",
			},
			"include_source_credentials": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key.",
			},
			"encryption_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"sources": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Protection Sources.",
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
						"last_refreshed_time": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time at which the data about this protection source was last refreshed.",
						},
						"registration_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Id of the registration as part of which this source was discovered.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmProtectionSourcesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourcesOptions := &backuprecoveryv1.GetProtectionSourcesOptions{}

	if _, ok := d.GetOk("request_initiator_type"); ok {
		getProtectionSourcesOptions.SetRequestInitiatorType(d.Get("request_initiator_type").(string))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getProtectionSourcesOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getProtectionSourcesOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("include_source_credentials"); ok {
		getProtectionSourcesOptions.SetIncludeSourceCredentials(d.Get("include_source_credentials").(bool))
	}
	if _, ok := d.GetOk("encryption_key"); ok {
		getProtectionSourcesOptions.SetEncryptionKey(d.Get("encryption_key").(string))
	}

	protectionSources, response, err := backupRecoveryClient.GetProtectionSourcesWithContext(context, getProtectionSourcesOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionSourcesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionSourcesWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmProtectionSourcesID(d))

	sources := []map[string]interface{}{}
	if protectionSources.Sources != nil {
		for _, modelItem := range protectionSources.Sources {
			modelMap, err := dataSourceIbmProtectionSourcesSourceToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			sources = append(sources, modelMap)
		}
	}
	if err = d.Set("sources", sources); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sources %s", err))
	}

	return nil
}

// dataSourceIbmProtectionSourcesID returns a reasonable ID for the list.
func dataSourceIbmProtectionSourcesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmProtectionSourcesSourceToMap(model *backuprecoveryv1.Source) (map[string]interface{}, error) {
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
			protectionStatsItemMap, err := dataSourceIbmProtectionSourcesObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := dataSourceIbmProtectionSourcesPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := dataSourceIbmProtectionSourcesObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := dataSourceIbmProtectionSourcesObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	if model.LastRefreshedTime != nil {
		modelMap["last_refreshed_time"] = flex.IntValue(model.LastRefreshedTime)
	}
	if model.RegistrationID != nil {
		modelMap["registration_id"] = flex.IntValue(model.RegistrationID)
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionSourcesPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := dataSourceIbmProtectionSourcesUserToMap(&usersItem)
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
			groupsItemMap, err := dataSourceIbmProtectionSourcesGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := dataSourceIbmProtectionSourcesTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionSourcesGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionSourcesTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesObjectOracleParamsToMap(model *backuprecoveryv1.ObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := dataSourceIbmProtectionSourcesDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := dataSourceIbmProtectionSourcesHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := dataSourceIbmProtectionSourcesPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := dataSourceIbmProtectionSourcesDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func dataSourceIbmProtectionSourcesHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
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

func dataSourceIbmProtectionSourcesObjectPhysicalParamsToMap(model *backuprecoveryv1.ObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}
