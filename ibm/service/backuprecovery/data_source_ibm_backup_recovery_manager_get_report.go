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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerGetReport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetReportRead,

		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of the report.",
			},
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
	}
}

func dataSourceIbmBackupRecoveryManagerGetReportRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_report", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getReportByIdOptions := &backuprecoveryv1.GetReportByIdOptions{}

	getReportByIdOptions.SetID(d.Get("report_id").(string))

	report, _, err := backupRecoveryClient.GetReportByIDWithContext(context, getReportByIdOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReportByIDWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_report", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*report.ID)

	if !core.IsNil(report.Category) {
		if err = d.Set("category", report.Category); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting category: %s", err), "(Data) ibm_backup_recovery_manager_get_report", "read", "set-category").GetDiag()
		}
	}

	if !core.IsNil(report.ComponentIds) {
		componentIds := []interface{}{}
		for _, componentIdsItem := range report.ComponentIds {
			componentIds = append(componentIds, componentIdsItem)
		}
		if err = d.Set("component_ids", componentIds); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting component_ids: %s", err), "(Data) ibm_backup_recovery_manager_get_report", "read", "set-component_ids").GetDiag()
		}
	}

	if !core.IsNil(report.Description) {
		if err = d.Set("description", report.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_backup_recovery_manager_get_report", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(report.SupportedUserContexts) {
		supportedUserContexts := []interface{}{}
		for _, supportedUserContextsItem := range report.SupportedUserContexts {
			supportedUserContexts = append(supportedUserContexts, supportedUserContextsItem)
		}
		if err = d.Set("supported_user_contexts", supportedUserContexts); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting supported_user_contexts: %s", err), "(Data) ibm_backup_recovery_manager_get_report", "read", "set-supported_user_contexts").GetDiag()
		}
	}

	if !core.IsNil(report.Title) {
		if err = d.Set("title", report.Title); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting title: %s", err), "(Data) ibm_backup_recovery_manager_get_report", "read", "set-title").GetDiag()
		}
	}

	return nil
}
