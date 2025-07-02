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

func DataSourceIbmBackupRecoveryManagerGetReportType() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetReportTypeRead,

		Schema: map[string]*schema.Schema{
			"report_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the report type.",
			},
			"attributes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the attribute name.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the data type of the attribute.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the attribute name.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetReportTypeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_report_type", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getReportTypeOptions := &backuprecoveryv1.GetReportTypeOptions{}

	getReportTypeOptions.SetReportType(d.Get("report_type").(string))

	reportTypeAttributes, _, err := backupRecoveryClient.GetReportTypeWithContext(context, getReportTypeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReportTypeWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_report_type", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetReportTypeID(d))

	if !core.IsNil(reportTypeAttributes.Attributes) {
		attributes := []map[string]interface{}{}
		for _, attributesItem := range reportTypeAttributes.Attributes {
			attributesItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportTypeReportTypeAttributeToMap(&attributesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_report_type", "read", "attributes-to-map").GetDiag()
			}
			attributes = append(attributes, attributesItemMap)
		}
		if err = d.Set("attributes", attributes); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting attributes: %s", err), "(Data) ibm_backup_recovery_manager_get_report_type", "read", "set-attributes").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetReportTypeID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetReportTypeID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetReportTypeReportTypeAttributeToMap(model *backuprecoveryv1.ReportTypeAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DataType != nil {
		modelMap["data_type"] = *model.DataType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
