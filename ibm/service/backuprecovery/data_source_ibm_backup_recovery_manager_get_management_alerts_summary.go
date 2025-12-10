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

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryRead,

		Schema: map[string]*schema.Schema{
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day.",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time.",
			},
			"states_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alerts_summary": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies a list of alerts summary grouped by category.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Category of alerts by which summary is grouped.",
						},
						"critical_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies count of critical alerts.",
						},
						"info_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies count of info alerts.",
						},
						"total_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies count of total alerts.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type/bucket that this alert category belongs to.",
						},
						"warning_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies count of warning alerts.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_management_alerts_summary", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_get_management_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_get_management_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	getManagementAlertsSummaryOptions := &backuprecoveryv1.GetManagementAlertsSummaryOptions{}

	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []string
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			clusterIdentifiersItem := v.(string)
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItem)
		}
		getManagementAlertsSummaryOptions.SetClusterIdentifiers(clusterIdentifiers)
	}
	if _, ok := d.GetOk("start_time_usecs"); ok {
		getManagementAlertsSummaryOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getManagementAlertsSummaryOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("states_list"); ok {
		var statesList []string
		for _, v := range d.Get("states_list").([]interface{}) {
			statesListItem := v.(string)
			statesList = append(statesList, statesListItem)
		}
		getManagementAlertsSummaryOptions.SetStatesList(statesList)
	}

	alertsManagementSummaryResponse, _, err := managementApiClient.GetManagementAlertsSummaryWithContext(context, getManagementAlertsSummaryOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetManagementAlertsSummaryWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_management_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryID(d))

	if !core.IsNil(alertsManagementSummaryResponse.AlertsSummary) {
		alertsSummary := []map[string]interface{}{}
		for _, alertsSummaryItem := range alertsManagementSummaryResponse.AlertsSummary {
			alertsSummaryItemMap, err := DataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryAlertGroupSummaryToMap(&alertsSummaryItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_management_alerts_summary", "read", "alerts_summary-to-map").GetDiag()
			}
			alertsSummary = append(alertsSummary, alertsSummaryItemMap)
		}
		if err = d.Set("alerts_summary", alertsSummary); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alerts_summary: %s", err), "(Data) ibm_backup_recovery_manager_get_management_alerts_summary", "read", "set-alerts_summary").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetManagementAlertsSummaryAlertGroupSummaryToMap(model *backuprecoveryv1.AlertGroupSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Category != nil {
		modelMap["category"] = *model.Category
	}
	if model.CriticalCount != nil {
		modelMap["critical_count"] = flex.IntValue(model.CriticalCount)
	}
	if model.InfoCount != nil {
		modelMap["info_count"] = flex.IntValue(model.InfoCount)
	}
	if model.TotalCount != nil {
		modelMap["total_count"] = flex.IntValue(model.TotalCount)
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.WarningCount != nil {
		modelMap["warning_count"] = flex.IntValue(model.WarningCount)
	}
	return modelMap, nil
}
