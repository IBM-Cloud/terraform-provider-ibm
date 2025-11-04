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

func DataSourceIbmBackupRecoveryManagerGetAlertsSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetAlertsSummaryRead,

		Schema: map[string]*schema.Schema{
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
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user's organization should be used to compute summary.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TenantIds contains ids of the tenants for which alerts are to be used to compute summary.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"states_list": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"x_scope_identifier": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field uniquely represents a service        instance. Please specify the values as \"service-instance-id: <value>\".",
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

func dataSourceIbmBackupRecoveryManagerGetAlertsSummaryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts_summary", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_get_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_get_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	getAlertSummaryOptions := &backuprecoveryv1.GetAlertSummaryOptions{}

	if _, ok := d.GetOk("start_time_usecs"); ok {
		getAlertSummaryOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getAlertSummaryOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		getAlertSummaryOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getAlertSummaryOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("states_list"); ok {
		var statesList []string
		for _, v := range d.Get("states_list").([]interface{}) {
			statesListItem := v.(string)
			statesList = append(statesList, statesListItem)
		}
		getAlertSummaryOptions.SetStatesList(statesList)
	}
	if _, ok := d.GetOk("x_scope_identifier"); ok {
		getAlertSummaryOptions.SetXScopeIdentifier(d.Get("x_scope_identifier").(string))
	}

	alertsSummaryResponse, _, err := managementApiClient.GetAlertSummaryWithContext(context, getAlertSummaryOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAlertSummaryWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_alerts_summary", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetAlertsSummaryID(d))

	if !core.IsNil(alertsSummaryResponse.AlertsSummary) {
		alertsSummary := []map[string]interface{}{}
		for _, alertsSummaryItem := range alertsSummaryResponse.AlertsSummary {
			alertsSummaryItemMap, err := DataSourceIbmBackupRecoveryManagerGetAlertsSummaryAlertGroupSummaryToMap(&alertsSummaryItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_alerts_summary", "read", "alerts_summary-to-map").GetDiag()
			}
			alertsSummary = append(alertsSummary, alertsSummaryItemMap)
		}
		if err = d.Set("alerts_summary", alertsSummary); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting alerts_summary: %s", err), "(Data) ibm_backup_recovery_manager_get_alerts_summary", "read", "set-alerts_summary").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetAlertsSummaryID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetAlertsSummaryID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetAlertsSummaryAlertGroupSummaryToMap(model *backuprecoveryv1.AlertGroupSummary) (map[string]interface{}, error) {
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
