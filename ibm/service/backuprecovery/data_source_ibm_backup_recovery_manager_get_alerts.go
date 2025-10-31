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

func DataSourceIbmBackupRecoveryManagerGetAlerts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetAlertsRead,

		Schema: map[string]*schema.Schema{
			"alert_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_types": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert types.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"alert_categories": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert categories.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_states": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert states.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_severities": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert severity types.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alert_type_buckets": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by list of alert type buckets.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies start time Unix epoch time in microseconds to filter alerts by.",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies end time Unix epoch time in microseconds to filter alerts by.",
			},
			"max_alerts": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies maximum number of alerts to return.The default value is 100 and maximum allowed value is 1000.",
			},
			"property_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies name of the property to filter alerts by.",
			},
			"property_value": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies value of the property to filter alerts by.",
			},
			"alert_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies name of alert to filter alerts by.",
			},
			"resolution_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies alert resolution ids to filter alerts by.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
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
			"all_under_hierarchy": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter by objects of all the tenants under the hierarchy of the logged in user's organization.",
			},
			"x_scope_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field uniquely represents a service        instance. Please specify the values as \"service-instance-id: <value>\".",
			},
			"alerts": &schema.Schema{
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
						"event_source": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies source where the event occurred.",
						},
						"first_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies Unix epoch Timestamp (in microseconds) of the first occurrence of the Alert.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies unique id of the alert.",
						},
						"label_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the labels for which this alert has been raised.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"latest_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies Unix epoch Timestamp (in microseconds) of the most recent occurrence of the Alert.",
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
						"resolution_details": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies information about the Alert Resolution.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resolution_details": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies detailed notes about the Resolution.",
									},
									"resolution_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the unique resolution id assigned in management console.",
									},
									"resolution_summary": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies short description about the Resolution.",
									},
									"timestamp_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies unix epoch timestamp (in microseconds) when the Alert was resolved.",
									},
									"user_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies name of the IBM Cluster user who resolved the Alerts.",
									},
								},
							},
						},
						"resolution_id_string": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resolution Id String.",
						},
						"resolved_timestamp_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies Unix epoch Timestamps in microseconds when alert is resolved.",
						},
						"severity": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the alert severity.",
						},
						"suppression_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies unique id generated when the Alert is suppressed by the admin.",
						},
						"tenant_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the tenants for which this alert has been raised.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func dataSourceIbmBackupRecoveryManagerGetAlertsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAlertsOptions := &backuprecoveryv1.GetAlertsOptions{}

	if _, ok := d.GetOk("alert_ids"); ok {
		var alertIds []string
		for _, v := range d.Get("alert_ids").([]interface{}) {
			alertIdsItem := v.(string)
			alertIds = append(alertIds, alertIdsItem)
		}
		getAlertsOptions.SetAlertIds(alertIds)
	}
	if _, ok := d.GetOk("alert_types"); ok {
		var alertTypes []int64
		for _, v := range d.Get("alert_types").([]interface{}) {
			alertTypesItem := int64(v.(int))
			alertTypes = append(alertTypes, alertTypesItem)
		}
		getAlertsOptions.SetAlertTypes(alertTypes)
	}
	if _, ok := d.GetOk("alert_categories"); ok {
		var alertCategories []string
		for _, v := range d.Get("alert_categories").([]interface{}) {
			alertCategoriesItem := v.(string)
			alertCategories = append(alertCategories, alertCategoriesItem)
		}
		getAlertsOptions.SetAlertCategories(alertCategories)
	}
	if _, ok := d.GetOk("alert_states"); ok {
		var alertStates []string
		for _, v := range d.Get("alert_states").([]interface{}) {
			alertStatesItem := v.(string)
			alertStates = append(alertStates, alertStatesItem)
		}
		getAlertsOptions.SetAlertStates(alertStates)
	}
	if _, ok := d.GetOk("alert_severities"); ok {
		var alertSeverities []string
		for _, v := range d.Get("alert_severities").([]interface{}) {
			alertSeveritiesItem := v.(string)
			alertSeverities = append(alertSeverities, alertSeveritiesItem)
		}
		getAlertsOptions.SetAlertSeverities(alertSeverities)
	}
	if _, ok := d.GetOk("alert_type_buckets"); ok {
		var alertTypeBuckets []string
		for _, v := range d.Get("alert_type_buckets").([]interface{}) {
			alertTypeBucketsItem := v.(string)
			alertTypeBuckets = append(alertTypeBuckets, alertTypeBucketsItem)
		}
		getAlertsOptions.SetAlertTypeBuckets(alertTypeBuckets)
	}
	if _, ok := d.GetOk("start_time_usecs"); ok {
		getAlertsOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getAlertsOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("max_alerts"); ok {
		getAlertsOptions.SetMaxAlerts(int64(d.Get("max_alerts").(int)))
	}
	if _, ok := d.GetOk("property_key"); ok {
		getAlertsOptions.SetPropertyKey(d.Get("property_key").(string))
	}
	if _, ok := d.GetOk("property_value"); ok {
		getAlertsOptions.SetPropertyValue(d.Get("property_value").(string))
	}
	if _, ok := d.GetOk("alert_name"); ok {
		getAlertsOptions.SetAlertName(d.Get("alert_name").(string))
	}
	if _, ok := d.GetOk("resolution_ids"); ok {
		var resolutionIds []int64
		for _, v := range d.Get("resolution_ids").([]interface{}) {
			resolutionIdsItem := int64(v.(int))
			resolutionIds = append(resolutionIds, resolutionIdsItem)
		}
		getAlertsOptions.SetResolutionIds(resolutionIds)
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getAlertsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("all_under_hierarchy"); ok {
		getAlertsOptions.SetAllUnderHierarchy(d.Get("all_under_hierarchy").(bool))
	}
	if _, ok := d.GetOk("x_scope_identifier"); ok {
		getAlertsOptions.SetXScopeIdentifier(d.Get("x_scope_identifier").(string))
	}

	alertList, _, err := managementSreApiClient.GetAlertsWithContext(context, getAlertsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAlertsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_alerts", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetAlertsID(d))

	alerts := []map[string]interface{}{}
	for _, alertsItem := range alertList.Alerts {
		alertsItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsAlertInfoToMap(&alertsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts", "read", "alerts-to-map").GetDiag()
		}
		alerts = append(alerts, alertsItemMap)
	}
	if err = d.Set("alerts", alerts); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alerts: %s", err), "(Data) ibm_backup_recovery_manager_get_alerts", "read", "set-alerts").GetDiag()
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetAlertsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetAlertsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetAlertsAlertInfoToMap(model *backuprecoveryv1.AlertInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AlertCategory != nil {
		modelMap["alert_category"] = *model.AlertCategory
	}
	if model.AlertCode != nil {
		modelMap["alert_code"] = *model.AlertCode
	}
	if model.AlertDocument != nil {
		alertDocumentMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsAlertDocumentToMap(model.AlertDocument)
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
	if model.EventSource != nil {
		modelMap["event_source"] = *model.EventSource
	}
	if model.FirstTimestampUsecs != nil {
		modelMap["first_timestamp_usecs"] = flex.IntValue(model.FirstTimestampUsecs)
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.LabelIds != nil {
		modelMap["label_ids"] = model.LabelIds
	}
	if model.LatestTimestampUsecs != nil {
		modelMap["latest_timestamp_usecs"] = flex.IntValue(model.LatestTimestampUsecs)
	}
	if model.PropertyList != nil {
		propertyList := []map[string]interface{}{}
		for _, propertyListItem := range model.PropertyList {
			propertyListItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsLabelToMap(&propertyListItem) // #nosec G601
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
	if model.ResolutionDetails != nil {
		resolutionDetailsMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsAlertResolutionDetailsToMap(model.ResolutionDetails)
		if err != nil {
			return modelMap, err
		}
		modelMap["resolution_details"] = []map[string]interface{}{resolutionDetailsMap}
	}
	if model.ResolutionIdString != nil {
		modelMap["resolution_id_string"] = *model.ResolutionIdString
	}
	if model.ResolvedTimestampUsecs != nil {
		modelMap["resolved_timestamp_usecs"] = flex.IntValue(model.ResolvedTimestampUsecs)
	}
	if model.Severity != nil {
		modelMap["severity"] = *model.Severity
	}
	if model.SuppressionID != nil {
		modelMap["suppression_id"] = flex.IntValue(model.SuppressionID)
	}
	if model.TenantIds != nil {
		modelMap["tenant_ids"] = model.TenantIds
	}
	if model.Vaults != nil {
		vaults := []map[string]interface{}{}
		for _, vaultsItem := range model.Vaults {
			vaultsItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsVaultToMap(&vaultsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			vaults = append(vaults, vaultsItemMap)
		}
		modelMap["vaults"] = vaults
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetAlertsAlertDocumentToMap(model *backuprecoveryv1.AlertDocument) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryManagerGetAlertsLabelToMap(model *backuprecoveryv1.Label) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = *model.Key
	modelMap["value"] = *model.Value
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetAlertsAlertResolutionDetailsToMap(model *backuprecoveryv1.AlertResolutionDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResolutionDetails != nil {
		modelMap["resolution_details"] = *model.ResolutionDetails
	}
	if model.ResolutionID != nil {
		modelMap["resolution_id"] = flex.IntValue(model.ResolutionID)
	}
	if model.ResolutionSummary != nil {
		modelMap["resolution_summary"] = *model.ResolutionSummary
	}
	if model.TimestampUsecs != nil {
		modelMap["timestamp_usecs"] = flex.IntValue(model.TimestampUsecs)
	}
	if model.UserName != nil {
		modelMap["user_name"] = *model.UserName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetAlertsVaultToMap(model *backuprecoveryv1.Vault) (map[string]interface{}, error) {
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
