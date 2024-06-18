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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv0"
)

func ResourceIbmCommonSourceRegistrationRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCommonSourceRegistrationRequestCreate,
		ReadContext:   resourceIbmCommonSourceRegistrationRequestRead,
		UpdateContext: resourceIbmCommonSourceRegistrationRequestUpdate,
		DeleteContext: resourceIbmCommonSourceRegistrationRequestDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"environment": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the environment type of the Protection Source.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A user specified name for this source.",
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
				Description: "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user.",
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
		},
	}
}

func ResourceIbmCommonSourceRegistrationRequestValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_common_source_registration_request", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCommonSourceRegistrationRequestCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	registerProtectionSourceOptions := &backuprecoveryv0.RegisterProtectionSourceOptions{}

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
		var connections []backuprecoveryv0.ConnectionConfig
		for _, v := range d.Get("connections").([]interface{}) {
			value := v.(map[string]interface{})
			connectionsItem, err := resourceIbmCommonSourceRegistrationRequestMapToConnectionConfig(value)
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
		var advancedConfigs []backuprecoveryv0.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			advancedConfigsItem, err := resourceIbmCommonSourceRegistrationRequestMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, *advancedConfigsItem)
		}
		registerProtectionSourceOptions.SetAdvancedConfigs(advancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		physicalParamsModel, err := resourceIbmCommonSourceRegistrationRequestMapToPhysicalParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		registerProtectionSourceOptions.SetPhysicalParams(physicalParamsModel)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		oracleParamsModel, err := resourceIbmCommonSourceRegistrationRequestMapToOracleParams(d.Get("oracle_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		registerProtectionSourceOptions.SetOracleParams(oracleParamsModel)
	}

	commonSourceRegistrationReponseParams, response, err := backupRecoveryClient.RegisterProtectionSourceWithContext(context, registerProtectionSourceOptions)
	if err != nil {
		log.Printf("[DEBUG] RegisterProtectionSourceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("RegisterProtectionSourceWithContext failed %s\n%s", err, response))
	}

	d.SetId(strconv.Itoa(int(*commonSourceRegistrationReponseParams.ID)))

	return resourceIbmCommonSourceRegistrationRequestRead(context, d, meta)
}

func resourceIbmCommonSourceRegistrationRequestRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions := &backuprecoveryv0.GetProtectionSourceRegistrationOptions{}

	// getProtectionSourceRegistrationOptions.SetID(d.Id())

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionSourceRegistrationOptions.SetID(int64(id))

	commonSourceRegistrationReponseParams, response, err := backupRecoveryClient.GetProtectionSourceRegistrationWithContext(context, getProtectionSourceRegistrationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionSourceRegistrationWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("environment", commonSourceRegistrationReponseParams.Environment); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.Name) {
		if err = d.Set("name", commonSourceRegistrationReponseParams.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}

	if !core.IsNil(commonSourceRegistrationReponseParams.ConnectionID) {
		if err = d.Set("connection_id", flex.IntValue(commonSourceRegistrationReponseParams.ConnectionID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connection_id: %s", err))
		}
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.Connections) {
		connections := []map[string]interface{}{}
		for _, connectionsItem := range commonSourceRegistrationReponseParams.Connections {
			connectionsItemMap, err := resourceIbmCommonSourceRegistrationRequestConnectionConfigToMap(&connectionsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			connections = append(connections, connectionsItemMap)
		}
		if err = d.Set("connections", connections); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connections: %s", err))
		}
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.ConnectorGroupID) {
		if err = d.Set("connector_group_id", flex.IntValue(commonSourceRegistrationReponseParams.ConnectorGroupID)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting connector_group_id: %s", err))
		}
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.AdvancedConfigs) {
		advancedConfigs := []map[string]interface{}{}
		for _, advancedConfigsItem := range commonSourceRegistrationReponseParams.AdvancedConfigs {
			advancedConfigsItemMap, err := resourceIbmCommonSourceRegistrationRequestKeyValuePairToMap(&advancedConfigsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			advancedConfigs = append(advancedConfigs, advancedConfigsItemMap)
		}
		if err = d.Set("advanced_configs", advancedConfigs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting advanced_configs: %s", err))
		}
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.PhysicalParams) {
		physicalParamsMap, err := resourceIbmCommonSourceRegistrationRequestPhysicalParamsToMap(commonSourceRegistrationReponseParams.PhysicalParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("physical_params", []map[string]interface{}{physicalParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting physical_params: %s", err))
		}
	}
	if !core.IsNil(commonSourceRegistrationReponseParams.OracleParams) {
		oracleParamsMap, err := resourceIbmCommonSourceRegistrationRequestOracleParamsToMap(commonSourceRegistrationReponseParams.OracleParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("oracle_params", []map[string]interface{}{oracleParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting oracle_params: %s", err))
		}
	}

	return nil
}

func resourceIbmCommonSourceRegistrationRequestUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProtectionSourceRegistrationOptions := &backuprecoveryv0.UpdateProtectionSourceRegistrationOptions{}

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	updateProtectionSourceRegistrationOptions.SetID(int64(id))

	updateProtectionSourceRegistrationOptions.SetEnvironment(d.Get("environment").(string))
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
		var newConnections []backuprecoveryv0.ConnectionConfig
		for _, v := range d.Get("connections").([]interface{}) {
			value := v.(map[string]interface{})
			newConnectionsItem, err := resourceIbmCommonSourceRegistrationRequestMapToConnectionConfig(value)
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
		var newAdvancedConfigs []backuprecoveryv0.KeyValuePair
		for _, v := range d.Get("advanced_configs").([]interface{}) {
			value := v.(map[string]interface{})
			newAdvancedConfigsItem, err := resourceIbmCommonSourceRegistrationRequestMapToKeyValuePair(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newAdvancedConfigs = append(newAdvancedConfigs, *newAdvancedConfigsItem)
		}
		updateProtectionSourceRegistrationOptions.SetAdvancedConfigs(newAdvancedConfigs)
	}
	if _, ok := d.GetOk("physical_params"); ok {
		newPhysicalParams, err := resourceIbmCommonSourceRegistrationRequestMapToPhysicalParams(d.Get("physical_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProtectionSourceRegistrationOptions.SetPhysicalParams(newPhysicalParams)
	}
	if _, ok := d.GetOk("oracle_params"); ok {
		newOracleParams, err := resourceIbmCommonSourceRegistrationRequestMapToOracleParams(d.Get("oracle_params.0").(map[string]interface{}))
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

	return resourceIbmCommonSourceRegistrationRequestRead(context, d, meta)
}

func resourceIbmCommonSourceRegistrationRequestDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV0()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProtectionSourceRegistrationOptions := &backuprecoveryv0.DeleteProtectionSourceRegistrationOptions{}

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

func resourceIbmCommonSourceRegistrationRequestMapToConnectionConfig(modelMap map[string]interface{}) (*backuprecoveryv0.ConnectionConfig, error) {
	model := &backuprecoveryv0.ConnectionConfig{}
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

func resourceIbmCommonSourceRegistrationRequestMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv0.KeyValuePair, error) {
	model := &backuprecoveryv0.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmCommonSourceRegistrationRequestMapToPhysicalParams(modelMap map[string]interface{}) (*backuprecoveryv0.PhysicalParams, error) {
	model := &backuprecoveryv0.PhysicalParams{}
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

func resourceIbmCommonSourceRegistrationRequestMapToOracleParams(modelMap map[string]interface{}) (*backuprecoveryv0.OracleParams, error) {
	model := &backuprecoveryv0.OracleParams{}
	if modelMap["database_entity_info"] != nil && len(modelMap["database_entity_info"].([]interface{})) > 0 {
		DatabaseEntityInfoModel, err := resourceIbmCommonSourceRegistrationRequestMapToDatabaseEntityInfo(modelMap["database_entity_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DatabaseEntityInfo = DatabaseEntityInfoModel
	}
	if modelMap["host_info"] != nil && len(modelMap["host_info"].([]interface{})) > 0 {
		HostInfoModel, err := resourceIbmCommonSourceRegistrationRequestMapToHostInformation(modelMap["host_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.HostInfo = HostInfoModel
	}
	return model, nil
}

func resourceIbmCommonSourceRegistrationRequestMapToDatabaseEntityInfo(modelMap map[string]interface{}) (*backuprecoveryv0.DatabaseEntityInfo, error) {
	model := &backuprecoveryv0.DatabaseEntityInfo{}
	if modelMap["container_database_info"] != nil {
		containerDatabaseInfo := []backuprecoveryv0.PluggableDatabaseInfo{}
		for _, containerDatabaseInfoItem := range modelMap["container_database_info"].([]interface{}) {
			containerDatabaseInfoItemModel, err := resourceIbmCommonSourceRegistrationRequestMapToPluggableDatabaseInfo(containerDatabaseInfoItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, *containerDatabaseInfoItemModel)
		}
		model.ContainerDatabaseInfo = containerDatabaseInfo
	}
	if modelMap["data_guard_info"] != nil && len(modelMap["data_guard_info"].([]interface{})) > 0 {
		DataGuardInfoModel, err := resourceIbmCommonSourceRegistrationRequestMapToDataGuardInfo(modelMap["data_guard_info"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataGuardInfo = DataGuardInfoModel
	}
	return model, nil
}

func resourceIbmCommonSourceRegistrationRequestMapToPluggableDatabaseInfo(modelMap map[string]interface{}) (*backuprecoveryv0.PluggableDatabaseInfo, error) {
	model := &backuprecoveryv0.PluggableDatabaseInfo{}
	if modelMap["database_id"] != nil && modelMap["database_id"].(string) != "" {
		model.DatabaseID = core.StringPtr(modelMap["database_id"].(string))
	}
	if modelMap["database_name"] != nil && modelMap["database_name"].(string) != "" {
		model.DatabaseName = core.StringPtr(modelMap["database_name"].(string))
	}
	return model, nil
}

func resourceIbmCommonSourceRegistrationRequestMapToDataGuardInfo(modelMap map[string]interface{}) (*backuprecoveryv0.DataGuardInfo, error) {
	model := &backuprecoveryv0.DataGuardInfo{}
	if modelMap["role"] != nil && modelMap["role"].(string) != "" {
		model.Role = core.StringPtr(modelMap["role"].(string))
	}
	if modelMap["standby_type"] != nil && modelMap["standby_type"].(string) != "" {
		model.StandbyType = core.StringPtr(modelMap["standby_type"].(string))
	}
	return model, nil
}

func resourceIbmCommonSourceRegistrationRequestMapToHostInformation(modelMap map[string]interface{}) (*backuprecoveryv0.HostInformation, error) {
	model := &backuprecoveryv0.HostInformation{}
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

func resourceIbmCommonSourceRegistrationRequestConnectionConfigToMap(model *backuprecoveryv0.ConnectionConfig) (map[string]interface{}, error) {
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

func resourceIbmCommonSourceRegistrationRequestKeyValuePairToMap(model *backuprecoveryv0.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmCommonSourceRegistrationRequestPhysicalParamsToMap(model *backuprecoveryv0.PhysicalParams) (map[string]interface{}, error) {
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

func resourceIbmCommonSourceRegistrationRequestOracleParamsToMap(model *backuprecoveryv0.OracleParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseEntityInfo != nil {
		databaseEntityInfoMap, err := resourceIbmCommonSourceRegistrationRequestDatabaseEntityInfoToMap(model.DatabaseEntityInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["database_entity_info"] = []map[string]interface{}{databaseEntityInfoMap}
	}
	if model.HostInfo != nil {
		hostInfoMap, err := resourceIbmCommonSourceRegistrationRequestHostInformationToMap(model.HostInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["host_info"] = []map[string]interface{}{hostInfoMap}
	}
	return modelMap, nil
}

func resourceIbmCommonSourceRegistrationRequestDatabaseEntityInfoToMap(model *backuprecoveryv0.DatabaseEntityInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ContainerDatabaseInfo != nil {
		containerDatabaseInfo := []map[string]interface{}{}
		for _, containerDatabaseInfoItem := range model.ContainerDatabaseInfo {
			containerDatabaseInfoItemMap, err := resourceIbmCommonSourceRegistrationRequestPluggableDatabaseInfoToMap(&containerDatabaseInfoItem)
			if err != nil {
				return modelMap, err
			}
			containerDatabaseInfo = append(containerDatabaseInfo, containerDatabaseInfoItemMap)
		}
		modelMap["container_database_info"] = containerDatabaseInfo
	}
	if model.DataGuardInfo != nil {
		dataGuardInfoMap, err := resourceIbmCommonSourceRegistrationRequestDataGuardInfoToMap(model.DataGuardInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_guard_info"] = []map[string]interface{}{dataGuardInfoMap}
	}
	return modelMap, nil
}

func resourceIbmCommonSourceRegistrationRequestPluggableDatabaseInfoToMap(model *backuprecoveryv0.PluggableDatabaseInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DatabaseID != nil {
		modelMap["database_id"] = model.DatabaseID
	}
	if model.DatabaseName != nil {
		modelMap["database_name"] = model.DatabaseName
	}
	return modelMap, nil
}

func resourceIbmCommonSourceRegistrationRequestDataGuardInfoToMap(model *backuprecoveryv0.DataGuardInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Role != nil {
		modelMap["role"] = model.Role
	}
	if model.StandbyType != nil {
		modelMap["standby_type"] = model.StandbyType
	}
	return modelMap, nil
}

func resourceIbmCommonSourceRegistrationRequestHostInformationToMap(model *backuprecoveryv0.HostInformation) (map[string]interface{}, error) {
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
