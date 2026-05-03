// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmBackupRecoveryProtectionSourceRefreshBasic(t *testing.T) {
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionSourceRefreshConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_source_refresh.backup_recovery_protection_source_refresh_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_backup_recovery_protection_source_refresh.backup_recovery_protection_source_refresh_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionSourceRefreshConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_protection_source_refresh" "backup_recovery_source_registration_instance" {
			x_ibm_tenant_id = "%v"
			backup_recovery_protection_source_id =     72126
		}
	`, xIbmTenantID)
}
