// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryManagerExportReportDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerExportReportDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_export_report.backup_recovery_manager_export_report_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_export_report.backup_recovery_manager_export_report_instance", "backup_recovery_manager_export_report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerExportReportDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_export_report" "backup_recovery_manager_export_report_instance" {
			id = "id"
			async = true
			filters = [ { attribute="attribute", filter_type="In", in_filter_params={ attribute_data_type="Bool", attribute_labels=[ "attributeLabels" ], bool_filter_values=[ true ], int32_filter_values=[ 1 ], int64_filter_values=[ 1 ], string_filter_values=[ "stringFilterValues" ] }, range_filter_params={ lower_bound=1, upper_bound=1 }, systems_filter_params={ system_ids=[ "systemIds" ], system_names=[ "systemNames" ] }, time_range_filter_params={ date_range="Last1Hour", duration_hours=1, lower_bound=1, upper_bound=1 } } ]
			layout = "layout"
			reportFormat = "XLS"
			timezone = "timezone"
		}
	`)
}
