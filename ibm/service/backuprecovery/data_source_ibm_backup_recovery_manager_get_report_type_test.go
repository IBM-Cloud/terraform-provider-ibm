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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmBackupRecoveryManagerGetReportTypeDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetReportTypeDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_report_type.backup_recovery_manager_get_report_type_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_report_type.backup_recovery_manager_get_report_type_instance", "report_type"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetReportTypeDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_report_type" "backup_recovery_manager_get_report_type_instance" {
			reportType = "Failures"
		}
	`)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportTypeReportTypeAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["data_type"] = "Bool"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ReportTypeAttribute)
	model.DataType = core.StringPtr("Bool")
	model.Name = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportTypeReportTypeAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
