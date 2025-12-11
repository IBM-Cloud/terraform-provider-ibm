// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
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
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerGetManagementAlerts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetManagementAlertsRead,

		Schema: map[string]*schema.Schema{
			"alert_id_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_state_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert states.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_type_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert types.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"alert_severity_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert severity types.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of region ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of cluster ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_date_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the start time of the alerts to be returned. All the alerts returned are raised after the specified start time. This value should be in Unix timestamp epoch in microseconds.",
			},
			"end_date_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the end time of the alerts to be returned. All the alerts returned are raised before the specified end time. This value should be in Unix timestamp epoch in microseconds.",
			},
			"max_alerts": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies maximum number of alerts to return.",
			},
			"alert_category_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert categories.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by tenant ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_type_bucket_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert type buckets.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_property_key_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of the alert property keys to query.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_property_value_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of the alert property value, multiple values for one key should be joined by '|'.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies name of alert to filter alerts by.",
			},
			"service_instance_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies services instance ids to filter alerts for IBM customers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alerts_list": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_category": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the alert category.",
						},
						"alert_code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies a unique code that categorizes the Alert, for example: CE00200014, where CE stands for IBM Error, the alert state next 3 digits is the id of the Alert Category (e.g. 002 for 'kNode') and the last 5 digits is the id of the Alert Type (e.g. 00014 for 'kNodeHighCpuUsage').",
						},
						"alert_document": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the fields of alert document.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_cause": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the cause of alert.",
									},
									"alert_description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the description of alert.",
									},
									"alert_help_text": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the help text for alert.",
									},
									"alert_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of alert.",
									},
									"alert_summary": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Short description for the alert.",
									},
								},
							},
						},
						"alert_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the alert state.",
						},
						"alert_type": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the alert type.",
						},
						"alert_type_bucket": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the Alert type bucket.",
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Id of the cluster which the alert is associated.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of cluster which alert is raised from.",
						},
						"dedup_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the dedup count of alert.",
						},
						"dedup_timestamps": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies Unix epoch Timestamps (in microseconds) for the last 25 occurrences of duplicated Alerts that are stored with the original/primary Alert. Alerts are grouped into one Alert if the Alerts are the same type, are reporting on the same Object and occur within one hour. 'dedupCount' always reports the total count of duplicated Alerts even if there are more than 25 occurrences. For example, if there are 100 occurrences of this Alert, dedupTimestamps stores the timestamps of the last 25 occurrences and dedupCount equals 100.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"first_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SpeSpecifies Unix epoch Timestamp (in microseconds) of the first occurrence of the Alert.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies unique id of the alert.",
						},
						"latest_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SpeSpecifies Unix epoch Timestamp (in microseconds) of the most recent occurrence of the Alert.",
						},
						"property_list": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of property key and values associated with alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Key of the Label.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the Label, multiple values should be joined by '|'.",
									},
								},
							},
						},
						"region_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the region id of the alert.",
						},
						"resolution_id_string": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the resolution id of the alert if its resolved.",
						},
						"service_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of the serrvice instance which the alert is associated.",
						},
						"severity": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the alert severity.",
						},
						"vaults": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies information about vaults where source object associated with alert is vaulted. This could be empty if alert is not related to any source object or it is not vaulted.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"global_vault_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies Global vault id.",
									},
									"region_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies id of region where vault resides.",
									},
									"region_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of region where vault resides.",
									},
									"vault_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of vault.",
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

func dataSourceIbmBackupRecoveryManagerGetManagementAlertsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_management_alerts", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_get_management_alerts", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_get_management_alerts", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	getManagementAlertsOptions := &backuprecoveryv1.GetManagementAlertsOptions{}

	if _, ok := d.GetOk("alert_id_list"); ok {
		var alertIdList []string
		for _, v := range d.Get("alert_id_list").([]interface{}) {
			alertIdListItem := v.(string)
			alertIdList = append(alertIdList, alertIdListItem)
		}
		getManagementAlertsOptions.SetAlertIdList(alertIdList)
	}
	if _, ok := d.GetOk("alert_state_list"); ok {
		var alertStateList []string
		for _, v := range d.Get("alert_state_list").([]interface{}) {
			alertStateListItem := v.(string)
			alertStateList = append(alertStateList, alertStateListItem)
		}
		getManagementAlertsOptions.SetAlertStateList(alertStateList)
	}
	if _, ok := d.GetOk("alert_type_list"); ok {
		var alertTypeList []int64
		for _, v := range d.Get("alert_type_list").([]interface{}) {
			alertTypeListItem := int64(v.(int))
			alertTypeList = append(alertTypeList, alertTypeListItem)
		}
		getManagementAlertsOptions.SetAlertTypeList(alertTypeList)
	}
	if _, ok := d.GetOk("alert_severity_list"); ok {
		var alertSeverityList []string
		for _, v := range d.Get("alert_severity_list").([]interface{}) {
			alertSeverityListItem := v.(string)
			alertSeverityList = append(alertSeverityList, alertSeverityListItem)
		}
		getManagementAlertsOptions.SetAlertSeverityList(alertSeverityList)
	}
	if _, ok := d.GetOk("region_ids"); ok {
		var regionIds []string
		for _, v := range d.Get("region_ids").([]interface{}) {
			regionIdsItem := v.(string)
			regionIds = append(regionIds, regionIdsItem)
		}
		getManagementAlertsOptions.SetRegionIds(regionIds)
	}
	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []string
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			clusterIdentifiersItem := v.(string)
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItem)
		}
		getManagementAlertsOptions.SetClusterIdentifiers(clusterIdentifiers)
	}
	if _, ok := d.GetOk("start_date_usecs"); ok {
		getManagementAlertsOptions.SetStartDateUsecs(int64(d.Get("start_date_usecs").(int)))
	}
	if _, ok := d.GetOk("end_date_usecs"); ok {
		getManagementAlertsOptions.SetEndDateUsecs(int64(d.Get("end_date_usecs").(int)))
	}
	if _, ok := d.GetOk("max_alerts"); ok {
		getManagementAlertsOptions.SetMaxAlerts(int64(d.Get("max_alerts").(int)))
	}
	if _, ok := d.GetOk("alert_category_list"); ok {
		var alertCategoryList []string
		for _, v := range d.Get("alert_category_list").([]interface{}) {
			alertCategoryListItem := v.(string)
			alertCategoryList = append(alertCategoryList, alertCategoryListItem)
		}
		getManagementAlertsOptions.SetAlertCategoryList(alertCategoryList)
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getManagementAlertsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("alert_type_bucket_list"); ok {
		var alertTypeBucketList []string
		for _, v := range d.Get("alert_type_bucket_list").([]interface{}) {
			alertTypeBucketListItem := v.(string)
			alertTypeBucketList = append(alertTypeBucketList, alertTypeBucketListItem)
		}
		getManagementAlertsOptions.SetAlertTypeBucketList(alertTypeBucketList)
	}
	if _, ok := d.GetOk("alert_property_key_list"); ok {
		var alertPropertyKeyList []string
		for _, v := range d.Get("alert_property_key_list").([]interface{}) {
			alertPropertyKeyListItem := v.(string)
			alertPropertyKeyList = append(alertPropertyKeyList, alertPropertyKeyListItem)
		}
		getManagementAlertsOptions.SetAlertPropertyKeyList(alertPropertyKeyList)
	}
	if _, ok := d.GetOk("alert_property_value_list"); ok {
		var alertPropertyValueList []string
		for _, v := range d.Get("alert_property_value_list").([]interface{}) {
			alertPropertyValueListItem := v.(string)
			alertPropertyValueList = append(alertPropertyValueList, alertPropertyValueListItem)
		}
		getManagementAlertsOptions.SetAlertPropertyValueList(alertPropertyValueList)
	}
	if _, ok := d.GetOk("alert_name"); ok {
		getManagementAlertsOptions.SetAlertName(d.Get("alert_name").(string))
	}
	if _, ok := d.GetOk("service_instance_ids"); ok {
		var serviceInstanceIds []string
		for _, v := range d.Get("service_instance_ids").([]interface{}) {
			serviceInstanceIdsItem := v.(string)
			serviceInstanceIds = append(serviceInstanceIds, serviceInstanceIdsItem)
		}
		getManagementAlertsOptions.SetServiceInstanceIds(serviceInstanceIds)
	}

	alertsList, _, err := managementApiClient.GetManagementAlertsWithContext(context, getManagementAlertsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetManagementAlertsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_management_alerts", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetManagementAlertsID(d))

	alertsListResult := []map[string]interface{}{}
	for _, alertsListItem := range alertsList.AlertsList {
		alertsListItemMap, err := DataSourceIbmBackupRecoveryManagerGetManagementAlertsAlertToMap(&alertsListItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_management_alerts", "read", "alerts_list-to-map").GetDiag()
		}
		alertsListResult = append(alertsListResult, alertsListItemMap)
	}
	if err = d.Set("alerts_list", alertsListResult); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alerts_list: %s", err), "(Data) ibm_backup_recovery_manager_get_management_alerts", "read", "set-alerts_list").GetDiag()
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetManagementAlertsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetManagementAlertsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsAlertToMap(model *backuprecoveryv1.Alert) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AlertCategory != nil {
		modelMap["alert_category"] = *model.AlertCategory
	}
	if model.AlertCode != nil {
		modelMap["alert_code"] = *model.AlertCode
	}
	if model.AlertDocument != nil {
		alertDocumentMap, err := DataSourceIbmBackupRecoveryManagerGetManagementAlertsAlertDocumentToMap(model.AlertDocument)
		if err != nil {
			return modelMap, err
		}
		modelMap["alert_document"] = []map[string]interface{}{alertDocumentMap}
	}
	if model.AlertState != nil {
		modelMap["alert_state"] = *model.AlertState
	}
	if model.AlertType != nil {
		modelMap["alert_type"] = flex.IntValue(model.AlertType)
	}
	if model.AlertTypeBucket != nil {
		modelMap["alert_type_bucket"] = *model.AlertTypeBucket
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = *model.ClusterName
	}
	if model.DedupCount != nil {
		modelMap["dedup_count"] = flex.IntValue(model.DedupCount)
	}
	if model.DedupTimestamps != nil {
		modelMap["dedup_timestamps"] = model.DedupTimestamps
	}
	if model.FirstTimestampUsecs != nil {
		modelMap["first_timestamp_usecs"] = flex.IntValue(model.FirstTimestampUsecs)
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.LatestTimestampUsecs != nil {
		modelMap["latest_timestamp_usecs"] = flex.IntValue(model.LatestTimestampUsecs)
	}
	if model.PropertyList != nil {
		propertyList := []map[string]interface{}{}
		for _, propertyListItem := range model.PropertyList {
			propertyListItemMap, err := DataSourceIbmBackupRecoveryManagerGetManagementAlertsLabelToMap(&propertyListItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			propertyList = append(propertyList, propertyListItemMap)
		}
		modelMap["property_list"] = propertyList
	}
	if model.RegionID != nil {
		modelMap["region_id"] = *model.RegionID
	}
	if model.ResolutionIdString != nil {
		modelMap["resolution_id_string"] = *model.ResolutionIdString
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = *model.ServiceInstanceID
	}
	if model.Severity != nil {
		modelMap["severity"] = *model.Severity
	}
	if model.Vaults != nil {
		vaults := []map[string]interface{}{}
		for _, vaultsItem := range model.Vaults {
			vaultsItemMap, err := DataSourceIbmBackupRecoveryManagerGetManagementAlertsVaultToMap(&vaultsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			vaults = append(vaults, vaultsItemMap)
		}
		modelMap["vaults"] = vaults
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsAlertDocumentToMap(model *backuprecoveryv1.AlertDocument) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AlertCause != nil {
		modelMap["alert_cause"] = *model.AlertCause
	}
	if model.AlertDescription != nil {
		modelMap["alert_description"] = *model.AlertDescription
	}
	if model.AlertHelpText != nil {
		modelMap["alert_help_text"] = *model.AlertHelpText
	}
	if model.AlertName != nil {
		modelMap["alert_name"] = *model.AlertName
	}
	if model.AlertSummary != nil {
		modelMap["alert_summary"] = *model.AlertSummary
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsLabelToMap(model *backuprecoveryv1.Label) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = *model.Key
	modelMap["value"] = *model.Value
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsVaultToMap(model *backuprecoveryv1.Vault) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.GlobalVaultID != nil {
		modelMap["global_vault_id"] = *model.GlobalVaultID
	}
	if model.RegionID != nil {
		modelMap["region_id"] = *model.RegionID
	}
	if model.RegionName != nil {
		modelMap["region_name"] = *model.RegionName
	}
	if model.VaultName != nil {
		modelMap["vault_name"] = *model.VaultName
	}
	return modelMap, nil
}
