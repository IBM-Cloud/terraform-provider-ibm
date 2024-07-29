// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRunErrorReportDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRunErrorReportDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_run_error_report.run_error_report_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_run_error_report.run_error_report_instance", "run_error_report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_run_error_report.run_error_report_instance", "run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_run_error_report.run_error_report_instance", "object_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRunErrorReportDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_run_error_report" "run_error_report_instance" {
			id = "id"
			runId = "runId"
			objectId = "objectId"
		}
	`)
}
