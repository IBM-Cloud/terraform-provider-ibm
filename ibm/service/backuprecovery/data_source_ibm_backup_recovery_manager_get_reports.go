// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
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

func DataSourceIbmBackupRecoveryManagerGetReports() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetReportsRead,

		Schema: map[string]*schema.Schema{
			"ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the ids of reports to fetch.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_context": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the user context to filter reports.",
			},
			"reports": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies list of reports.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies categoty of the Report.",
						},
						"component_ids": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the list of component ids in the Report.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies description of the Report.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the report.",
						},
						"supported_user_contexts": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies all the supported user contexts for this report.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"title": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the title of the report.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetReportsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_reports", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getReportsOptions := &backuprecoveryv1.GetReportsOptions{}

	if _, ok := d.GetOk("ids"); ok {
		var ids []string
		for _, v := range d.Get("ids").([]interface{}) {
			idsItem := v.(string)
			ids = append(ids, idsItem)
		}
		getReportsOptions.SetIds(ids)
	}
	if _, ok := d.GetOk("user_context"); ok {
		getReportsOptions.SetUserContext(d.Get("user_context").(string))
	}

	reports, _, err := backupRecoveryClient.GetReportsWithContext(context, getReportsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReportsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_reports", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetReportsID(d))

	if !core.IsNil(reports.Reports) {
		reportsMap := []map[string]interface{}{}
		for _, reportsItem := range reports.Reports {
			reportsItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportsReportToMap(&reportsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_reports", "read", "reports-to-map").GetDiag()
			}
			reportsMap = append(reportsMap, reportsItemMap)
		}
		if err = d.Set("reports", reportsMap); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reports: %s", err), "(Data) ibm_backup_recovery_manager_get_reports", "read", "set-reports").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetReportsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetReportsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetReportsReportToMap(model *backuprecoveryv1.Report) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Category != nil {
		modelMap["category"] = *model.Category
	}
	if model.ComponentIds != nil {
		modelMap["component_ids"] = model.ComponentIds
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.SupportedUserContexts != nil {
		modelMap["supported_user_contexts"] = model.SupportedUserContexts
	}
	if model.Title != nil {
		modelMap["title"] = *model.Title
	}
	return modelMap, nil
}
