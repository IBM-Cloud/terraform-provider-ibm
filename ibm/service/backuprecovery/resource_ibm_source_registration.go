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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmSourceRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSourceRegistrationCreate,
		ReadContext:   resourceIbmSourceRegistrationRead,
		UpdateContext: resourceIbmSourceRegistrationUpdate,
		DeleteContext: resourceIbmSourceRegistrationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ValidateFunc: validate.InvokeValidator("ibm_source_registration", "environment"),
				Description: "Specifies the environment type of the Protection Source.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user specified name for this source.",
			},
			"is_internal_encrypted": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if credentials are encrypted by internal key.",
			},
			"encryption_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the key that user has encrypted the credential with.",
			},
			"connection_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field.",
			},
			"connections": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specfies the list of connections for the source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the id of the connection.",
						},
						"entity_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the entity id of the source. The source can a non-root entity.",
						},
						"connector_group_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the connector group id of connector groups.",
						},
					},
				},
			},
			"connector_group_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the connector group id of connector groups.",
			},
			"advanced_configs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the advanced configuration for a protection source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "key.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "value.",
						},
					},
				},
			},
			"physical_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Physical Params params.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the endpoint IPaddress, URL or hostname of the physical host.",
						},
						"force_register": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster.",
						},
						"host_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the type of host.",
						},
						"physical_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the type of physical server.",
						},
						"applications": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the list of applications to be registered with Physical Source.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"oracle_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Physical Params params.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_entity_info": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Object details about Oracle database entity info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"container_database_info": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies the list of Pluggable databases within a container database.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the database Id of the Pluggable DB.",
												},
												"database_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the name of the Pluggable DB.",
												},
											},
										},
									},
									"data_guard_info": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Dataguard info about Oracle database.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"role": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies the role of the Oracle DataGuard database.",
												},
												"standby_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
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
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the id of the host object.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the name of the host object.",
									},
									"environment": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the environment of the object.",
									},
								},
							},
						},
					},
				},
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
		},
	}
}

func ResourceIbmSourceRegistrationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "environment",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kOracle, kPhysical, kSQL",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_source_registration", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSourceRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	registerProtectionSourceOptions := &backuprecoveryv1.RegisterProtectionSourceOptions{}

	registerProtectionSourceOptions.SetEnvironment(d.Get("environment").(string))
	if _, ok := d.GetOk("name"); ok {
		registerProtectionSourceOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("is_internal_encrypted"); ok {
		registerProtectionSourceOptions.SetIsInternalEncrypted(d.Get("is_internal_encrypted").(bool))
	}
	if _, ok := d.GetOk("encryption_key"); ok {
		registerProtectionSourceOptions.SetEncryptionKey(d.Get("encryption_key").(string))
	}
	if _, ok := d.GetOk("connection_id"); ok {
		registerProtectionSourceOptions.SetConnectionID(int64(d.Get("connection_id").(int)))
	}
	if _, ok := d.GetOk("connections"); ok {
		var connections []backuprecoveryv1.ConnectionConfig
		for _, v := range d.Get("connections").([]interface{}) {
			value := v.(map[string]interface{})
			connectionsItem, err := resourceIbmSourceRegistrationMapToConnectionConfig(value)
			if err != nil {
				return diag.FromErr(err)
			}
			connections = append(connections, *connectionsItem)
		}
		registerProtectionSourceOptions.SetConnections(connections)
	}
	if _, ok := d.GetOk("connector_group_id"); ok {
		registerProtectionSourceOptions.SetConnectorGroupID(int64(d.Get("connector_group_id").(int)))
	}
	if _, ok := d.GetOk("advanced_configs"); ok {
		var advancedConfigs []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			advancedConfigsItem, err := resourceIbmSourceRegistrationMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, *advancedConfigsItem)
		}
		registerProtectionSourceOptions.SetAdvancedConfigs(advancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		physicalParamsModel, err := resourceIbmSourceRegistrationMapToPhysicalParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		registerProtectionSourceOptions.SetPhysicalParams(physicalParamsModel)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		oracleParamsModel, err := resourceIbmSourceRegistrationMapToOracleParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		registerProtectionSourceOptions.SetOracleParams(oracleParamsModel)
	}

	sourceRegistrationReponseParams, response, err := backupRecoveryClient.RegisterProtectionSourceWithContext(context, registerProtectionSourceOptions)
	if err != nil {
		log.Printf("[DEBUG] RegisterProtectionSourceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("RegisterProtectionSourceWithContext failed %s\n%s", err, response))
	}

	d.SetId(strconv.Itoa(int(*sourceRegistrationReponseParams.ID)))

	return resourceIbmSourceRegistrationRead(context, d, meta)
}

func resourceIbmSourceRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

	//getProtectionSourceRegistrationOptions.SetID(d.Id())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions.SetID(int64(id))

	sourceRegistrationReponseParams, response, err := backupRecoveryClient.GetProtectionSourceRegistrationWithContext(context, getProtectionSourceRegistrationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("environment", sourceRegistrationReponseParams.Environment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
	}
	if !core.IsNil(sourceRegistrationReponseParams.Name) {
		if err = d.Set("name", sourceRegistrationReponseParams.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}

	if !core.IsNil(sourceRegistrationReponseParams.ConnectionID) {
		if err = d.Set("connection_id", flex.IntValue(sourceRegistrationReponseParams.ConnectionID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connection_id: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.Connections) {
		connections := []map[string]interface{}{}
		for _, connectionsItem := range sourceRegistrationReponseParams.Connections {
			connectionsItemMap, err := resourceIbmSourceRegistrationConnectionConfigToMap(&connectionsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			connections = append(connections, connectionsItemMap)
		}
		if err = d.Set("connections", connections); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connections: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.ConnectorGroupID) {
		if err = d.Set("connector_group_id", flex.IntValue(sourceRegistrationReponseParams.ConnectorGroupID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connector_group_id: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.AdvancedConfigs) {
		advancedConfigs := []map[string]interface{}{}
		for _, advancedConfigsItem := range sourceRegistrationReponseParams.AdvancedConfigs {
			advancedConfigsItemMap, err := resourceIbmSourceRegistrationKeyValuePairToMap(&advancedConfigsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, advancedConfigsItemMap)
		}
		if err = d.Set("advanced_configs", advancedConfigs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting advanced_configs: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.PhysicalParams) {
		physicalParamsMap, err := resourceIbmSourceRegistrationPhysicalParamsToMap(sourceRegistrationReponseParams.PhysicalParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("physical_params", []map[string]interface{}{physicalParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting physical_params: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.OracleParams) {
		oracleParamsMap, err := resourceIbmSourceRegistrationOracleParamsToMap(sourceRegistrationReponseParams.OracleParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("oracle_params", []map[string]interface{}{oracleParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting oracle_params: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.SourceID) {
		if err = d.Set("source_id", flex.IntValue(sourceRegistrationReponseParams.SourceID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting source_id: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.SourceInfo) {
		sourceInfoMap, err := resourceIbmSourceRegistrationObjectToMap(sourceRegistrationReponseParams.SourceInfo)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("source_info", []map[string]interface{}{sourceInfoMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting source_info: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.AuthenticationStatus) {
		if err = d.Set("authentication_status", sourceRegistrationReponseParams.AuthenticationStatus); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting authentication_status: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.RegistrationTimeMsecs) {
		if err = d.Set("registration_time_msecs", flex.IntValue(sourceRegistrationReponseParams.RegistrationTimeMsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting registration_time_msecs: %s", err))
		}
	}
	if !core.IsNil(sourceRegistrationReponseParams.LastRefreshedTimeMsecs) {
		if err = d.Set("last_refreshed_time_msecs", flex.IntValue(sourceRegistrationReponseParams.LastRefreshedTimeMsecs)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_refreshed_time_msecs: %s", err))
		}
	}

	return nil
}

func resourceIbmSourceRegistrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProtectionSourceRegistrationOptions := &backuprecoveryv1.UpdateProtectionSourceRegistrationOptions{}

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	updateProtectionSourceRegistrationOptions.SetID(int64(id))

	updateProtectionSourceRegistrationOptions.SetEnvironment(d.Get("environment").(string))

	// if _, ok := d.GetOk("id"); ok {
	// 	updateProtectionSourceRegistrationOptions.SetID(d.Get("id").(string))
	// }
	if _, ok := d.GetOk("name"); ok {
		updateProtectionSourceRegistrationOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("is_internal_encrypted"); ok {
		updateProtectionSourceRegistrationOptions.SetIsInternalEncrypted(d.Get("is_internal_encrypted").(bool))
	}
	if _, ok := d.GetOk("encryption_key"); ok {
		updateProtectionSourceRegistrationOptions.SetEncryptionKey(d.Get("encryption_key").(string))
	}
	if _, ok := d.GetOk("connection_id"); ok {
		updateProtectionSourceRegistrationOptions.SetConnectionID(int64(d.Get("connection_id").(int)))
	}
	if _, ok := d.GetOk("connections"); ok {
		var newConnections []backuprecoveryv1.ConnectionConfig
		for _, v := range d.Get("connections").([]interface{}) {
			value := v.(map[string]interface{})
			newConnectionsItem, err := resourceIbmSourceRegistrationMapToConnectionConfig(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newConnections = append(newConnections, *newConnectionsItem)
		}
		updateProtectionSourceRegistrationOptions.SetConnections(newConnections)
	}
	if _, ok := d.GetOk("connector_group_id"); ok {
		updateProtectionSourceRegistrationOptions.SetConnectorGroupID(int64(d.Get("connector_group_id").(int)))
	}
	if _, ok := d.GetOk("advanced_configs"); ok {
		var newAdvancedConfigs []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			newAdvancedConfigsItem, err := resourceIbmSourceRegistrationMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newAdvancedConfigs = append(newAdvancedConfigs, *newAdvancedConfigsItem)
		}
		updateProtectionSourceRegistrationOptions.SetAdvancedConfigs(newAdvancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		newPhysicalParams, err := resourceIbmSourceRegistrationMapToPhysicalParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionSourceRegistrationOptions.SetPhysicalParams(newPhysicalParams)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		newOracleParams, err := resourceIbmSourceRegistrationMapToOracleParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionSourceRegistrationOptions.SetOracleParams(newOracleParams)
	}

	_, response, err := backupRecoveryClient.UpdateProtectionSourceRegistrationWithContext(context, updateProtectionSourceRegistrationOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateProtectionSourceRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateProtectionSourceRegistrationWithContext failed %s\n%s", err, response))
	}

	return resourceIbmSourceRegistrationRead(context, d, meta)
}

func resourceIbmSourceRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionSourceRegistrationOptions := &backuprecoveryv1.DeleteProtectionSourceRegistrationOptions{}

	// deleteProtectionSourceRegistrationOptions.SetID(d.Id())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionSourceRegistrationOptions.SetID(int64(id))

	response, err := backupRecoveryClient.DeleteProtectionSourceRegistrationWithContext(context, deleteProtectionSourceRegistrationOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteProtectionSourceRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteProtectionSourceRegistrationWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmSourceRegistrationMapToConnectionConfig(modelMap map[string]interface{}) (*backuprecoveryv1.ConnectionConfig, error) {
	model := &backuprecoveryv1.ConnectionConfig{}
	if modelMap["connection_id"] != nil {
		model.ConnectionID = core.Int64Ptr(int64(modelMap["connection_id"].(int)))
	}
	if modelMap["entity_id"] != nil {
		model.EntityID = core.Int64Ptr(int64(modelMap["entity_id"].(int)))
	}
	if modelMap["connector_group_id"] != nil {
		model.ConnectorGroupID = core.Int64Ptr(int64(modelMap["connector_group_id"].(int)))
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmSourceRegistrationMapToPhysicalParams(modelMap map[string]interface{}) (*backuprecoveryv1.PhysicalParams, error) {
	model := &backuprecoveryv1.PhysicalParams{}
	model.Endpoint = core.StringPtr(modelMap["endpoint"].(string))
	if modelMap["force_register"] != nil {
		model.ForceRegister = core.BoolPtr(modelMap["force_register"].(bool))
	}
	if modelMap["host_type"] != nil && modelMap["host_type"].(string) != "" {
		model.HostType = core.StringPtr(modelMap["host_type"].(string))
	}
	if modelMap["physical_type"] != nil && modelMap["physical_type"].(string) != "" {
		model.PhysicalType = core.StringPtr(modelMap["physical_type"].(string))
	}
	if modelMap["applications"] != nil {
		applications := []string{}
		for _, applicationsItem := range modelMap["applications"].([]interface{}) {
			applications = append(applications, applicationsItem.(string))
		}
		model.Applications = applications
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToOracleParams(modelMap map[string]interface{}) (*backuprecoveryv1.OracleParams, error) {
	model := &backuprecoveryv1.OracleParams{}
	if modelMap["database_entity_info"] != nil && len(modelMap["database_entity_info"].([]interface{})) > 0 {
		DatabaseEntityInfoModel, err := resourceIbmSourceRegistrationMapToDatabaseEntityInfo(modelMap["database_entity_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DatabaseEntityInfo = DatabaseEntityInfoModel
	}
	if modelMap["host_info"] != nil && len(modelMap["host_info"].([]interface{})) > 0 {
		HostInfoModel, err := resourceIbmSourceRegistrationMapToHostInformation(modelMap["host_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.HostInfo = HostInfoModel
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToDatabaseEntityInfo(modelMap map[string]interface{}) (*backuprecoveryv1.DatabaseEntityInfo, error) {
	model := &backuprecoveryv1.DatabaseEntityInfo{}
	if modelMap["container_database_info"] != nil {
		containerDatabaseInfo := []backuprecoveryv1.PluggableDatabaseInfo{}
		for _, containerDatabaseInfoItem := range modelMap["container_database_info"].([]interface{}) {
			containerDatabaseInfoItemModel, err := resourceIbmSourceRegistrationMapToPluggableDatabaseInfo(containerDatabaseInfoItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, *containerDatabaseInfoItemModel)
		}
		model.ContainerDatabaseInfo = containerDatabaseInfo
	}
	if modelMap["data_guard_info"] != nil && len(modelMap["data_guard_info"].([]interface{})) > 0 {
		DataGuardInfoModel, err := resourceIbmSourceRegistrationMapToDataGuardInfo(modelMap["data_guard_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataGuardInfo = DataGuardInfoModel
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToPluggableDatabaseInfo(modelMap map[string]interface{}) (*backuprecoveryv1.PluggableDatabaseInfo, error) {
	model := &backuprecoveryv1.PluggableDatabaseInfo{}
	if modelMap["database_id"] != nil && modelMap["database_id"].(string) != "" {
		model.DatabaseID = core.StringPtr(modelMap["database_id"].(string))
	}
	if modelMap["database_name"] != nil && modelMap["database_name"].(string) != "" {
		model.DatabaseName = core.StringPtr(modelMap["database_name"].(string))
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToDataGuardInfo(modelMap map[string]interface{}) (*backuprecoveryv1.DataGuardInfo, error) {
	model := &backuprecoveryv1.DataGuardInfo{}
	if modelMap["role"] != nil && modelMap["role"].(string) != "" {
		model.Role = core.StringPtr(modelMap["role"].(string))
	}
	if modelMap["standby_type"] != nil && modelMap["standby_type"].(string) != "" {
		model.StandbyType = core.StringPtr(modelMap["standby_type"].(string))
	}
	return model, nil
}

func resourceIbmSourceRegistrationMapToHostInformation(modelMap map[string]interface{}) (*backuprecoveryv1.HostInformation, error) {
	model := &backuprecoveryv1.HostInformation{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["environment"] != nil && modelMap["environment"].(string) != "" {
		model.Environment = core.StringPtr(modelMap["environment"].(string))
	}
	return model, nil
}

func resourceIbmSourceRegistrationConnectionConfigToMap(model *backuprecoveryv1.ConnectionConfig) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmSourceRegistrationPhysicalParamsToMap(model *backuprecoveryv1.PhysicalParams) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationOracleParamsToMap(model *backuprecoveryv1.OracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := resourceIbmSourceRegistrationDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := resourceIbmSourceRegistrationHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationDatabaseEntityInfoToMap(model *backuprecoveryv1.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := resourceIbmSourceRegistrationPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := resourceIbmSourceRegistrationDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationPluggableDatabaseInfoToMap(model *backuprecoveryv1.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationDataGuardInfoToMap(model *backuprecoveryv1.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationHostInformationToMap(model *backuprecoveryv1.HostInformation) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationObjectToMap(model *backuprecoveryv1.Object) (map[string]interface{}, error) {
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
			protectionStatsItemMap, err := resourceIbmSourceRegistrationObjectProtectionStatsSummaryToMap(&protectionStatsItem)
			if err != nil {
				return modelMap, err
			}
			protectionStats = append(protectionStats, protectionStatsItemMap)
		}
		modelMap["protection_stats"] = protectionStats
	}
	if model.Permissions != nil {
		permissionsMap, err := resourceIbmSourceRegistrationPermissionInfoToMap(model.Permissions)
		if err != nil {
			return modelMap, err
		}
		modelMap["permissions"] = []map[string]interface{}{permissionsMap}
	}
	if model.OracleParams != nil {
		oracleParamsMap, err := resourceIbmSourceRegistrationObjectOracleParamsToMap(model.OracleParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["oracle_params"] = []map[string]interface{}{oracleParamsMap}
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := resourceIbmSourceRegistrationObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationObjectProtectionStatsSummaryToMap(model *backuprecoveryv1.ObjectProtectionStatsSummary) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationPermissionInfoToMap(model *backuprecoveryv1.PermissionInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ObjectID != nil {
		modelMap["object_id"] = flex.IntValue(model.ObjectID)
	}
	if model.Users != nil {
		users := []map[string]interface{}{}
		for _, usersItem := range model.Users {
			usersItemMap, err := resourceIbmSourceRegistrationUserToMap(&usersItem)
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
			groupsItemMap, err := resourceIbmSourceRegistrationGroupToMap(&groupsItem)
			if err != nil {
				return modelMap, err
			}
			groups = append(groups, groupsItemMap)
		}
		modelMap["groups"] = groups
	}
	if model.Tenant != nil {
		tenantMap, err := resourceIbmSourceRegistrationTenantToMap(model.Tenant)
		if err != nil {
			return modelMap, err
		}
		modelMap["tenant"] = []map[string]interface{}{tenantMap}
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationUserToMap(model *backuprecoveryv1.User) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationGroupToMap(model *backuprecoveryv1.Group) (map[string]interface{}, error) {
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

func resourceIbmSourceRegistrationTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationObjectOracleParamsToMap(model *backuprecoveryv1.ObjectOracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := resourceIbmSourceRegistrationDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := resourceIbmSourceRegistrationHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func resourceIbmSourceRegistrationObjectPhysicalParamsToMap(model *backuprecoveryv1.ObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnableSystemBackup != nil {
		modelMap["enable_system_backup"] = model.EnableSystemBackup
	}
	return modelMap, nil
}
