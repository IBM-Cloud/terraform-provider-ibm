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

func TestAccIbmBackupRecoveryManagerGetReportsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetReportsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_reports.backup_recovery_manager_get_reports_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetReportsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_reports" "backup_recovery_manager_get_reports_instance" {
			ids = [ "ids" ]
			userContext = "IBMBaaS"
		}
	`)
}

func TestDataSourceIbmBackupRecoveryManagerGetReportsReportToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["category"] = "Protection"
		model["component_ids"] = []string{"testString"}
		model["description"] = "testString"
		model["id"] = "testString"
		model["supported_user_contexts"] = []string{"IBMBaaS"}
		model["title"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Report)
	model.Category = core.StringPtr("Protection")
	model.ComponentIds = []string{"testString"}
	model.Description = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.SupportedUserContexts = []string{"IBMBaaS"}
	model.Title = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetReportsReportToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
