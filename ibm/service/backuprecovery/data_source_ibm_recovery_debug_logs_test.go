// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRecoveryDebugLogsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDebugLogsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery_debug_logs.recovery_debug_logs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery_debug_logs.recovery_debug_logs_instance", "recovery_debug_logs_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryDebugLogsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_recovery_debug_logs" "recovery_debug_logs_instance" {
			id = "id"
		}
	`)
}
