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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerGetAlertsResolution() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetAlertsResolutionRead,

		Schema: map[string]*schema.Schema{
			"max_resolutions": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the max number of Resolutions to be returned, from the latest created to the earliest created.",
			},
			"resolution_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies Alert Resolution Name to query.",
			},
			"resolution_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies Alert Resolution id to query.",
			},
			"alert_resolutions_list": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of alert resolutions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies account id of the user who create the resolution.",
						},
						"created_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies unix epoch timestamp (in microseconds) when the resolution is created.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the full description about the Resolution.",
						},
						"external_key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the external key assigned outside of management console, with the form of \"clusterid:resolutionid\".",
						},
						"resolution_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique reslution id assigned in management console.",
						},
						"resolution_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique name of the resolution.",
						},
						"resolved_alerts": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Id of the alert.",
									},
									"alert_id_str": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Alert Id with string format.",
									},
									"alert_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the alert being resolved.",
									},
									"cluster_id": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Id of the cluster which the alert is associated.",
									},
									"first_timestamp_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "First occurrence of the alert.",
									},
									"resolved_time_usec": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"service_instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of the service instance which the alert is associated.",
									},
								},
							},
						},
						"silence_minutes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the time duration (in minutes) for silencing alerts. If unspecified or set zero, a silence rule will be created with default expiry time. No silence rule will be created if value < 0.",
						},
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies tenant id of the user who create the resolution.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetAlertsResolutionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts_resolution", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_get_alerts_resolution", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_get_alerts_resolution", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	getManagementAlertResolutionOptions := &backuprecoveryv1.GetManagementAlertResolutionOptions{}

	getManagementAlertResolutionOptions.SetMaxResolutions(int64(d.Get("max_resolutions").(int)))
	if _, ok := d.GetOk("resolution_name"); ok {
		getManagementAlertResolutionOptions.SetResolutionName(d.Get("resolution_name").(string))
	}
	if _, ok := d.GetOk("resolution_id"); ok {
		getManagementAlertResolutionOptions.SetResolutionID(d.Get("resolution_id").(string))
	}

	alertResolutionsList, _, err := managementApiClient.GetManagementAlertResolutionWithContext(context, getManagementAlertResolutionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetManagementAlertResolutionWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_alerts_resolution", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetAlertsResolutionID(d))

	if !core.IsNil(alertResolutionsList.AlertResolutionsList) {
		alertResolutionsListResult := []map[string]interface{}{}
		for _, alertResolutionsListItem := range alertResolutionsList.AlertResolutionsList {
			alertResolutionsListItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsResolutionAlertResolutionToMap(&alertResolutionsListItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts_resolution", "read", "alert_resolutions_list-to-map").GetDiag()
			}
			alertResolutionsListResult = append(alertResolutionsListResult, alertResolutionsListItemMap)
		}
		if err = d.Set("alert_resolutions_list", alertResolutionsListResult); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alert_resolutions_list: %s", err), "(Data) ibm_backup_recovery_manager_get_alerts_resolution", "read", "set-alert_resolutions_list").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetAlertsResolutionID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetAlertsResolutionID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetAlertsResolutionAlertResolutionToMap(model *backuprecoveryv1.AlertResolution) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.CreatedTimeUsecs != nil {
		modelMap["created_time_usecs"] = flex.IntValue(model.CreatedTimeUsecs)
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.ExternalKey != nil {
		modelMap["external_key"] = *model.ExternalKey
	}
	if model.ResolutionID != nil {
		modelMap["resolution_id"] = *model.ResolutionID
	}
	if model.ResolutionName != nil {
		modelMap["resolution_name"] = *model.ResolutionName
	}
	if model.ResolvedAlerts != nil {
		resolvedAlerts := []map[string]interface{}{}
		for _, resolvedAlertsItem := range model.ResolvedAlerts {
			resolvedAlertsItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsResolutionResolvedAlertInfoToMap(&resolvedAlertsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resolvedAlerts = append(resolvedAlerts, resolvedAlertsItemMap)
		}
		modelMap["resolved_alerts"] = resolvedAlerts
	}
	if model.SilenceMinutes != nil {
		modelMap["silence_minutes"] = flex.IntValue(model.SilenceMinutes)
	}
	if model.TenantID != nil {
		modelMap["tenant_id"] = *model.TenantID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetAlertsResolutionResolvedAlertInfoToMap(model *backuprecoveryv1.ResolvedAlertInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AlertID != nil {
		modelMap["alert_id"] = flex.IntValue(model.AlertID)
	}
	if model.AlertIdStr != nil {
		modelMap["alert_id_str"] = *model.AlertIdStr
	}
	if model.AlertName != nil {
		modelMap["alert_name"] = *model.AlertName
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.FirstTimestampUsecs != nil {
		modelMap["first_timestamp_usecs"] = flex.IntValue(model.FirstTimestampUsecs)
	}
	if model.ResolvedTimeUsec != nil {
		modelMap["resolved_time_usec"] = flex.IntValue(model.ResolvedTimeUsec)
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = *model.ServiceInstanceID
	}
	return modelMap, nil
}
