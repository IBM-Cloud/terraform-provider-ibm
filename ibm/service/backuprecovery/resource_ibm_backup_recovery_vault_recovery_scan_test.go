// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryVaultRecoveryScanBasic(t *testing.T) {
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	cloudType := "ibm"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryVaultRecoveryScanConfigBasic(xIbmTenantID, cloudType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance", "cloud_type", cloudType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryVaultRecoveryScanConfigBasic(xIbmTenantID string, cloudType string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_vault_recovery_scan" "backup_recovery_vault_recovery_scan_instance" {
			x_ibm_tenant_id = "%s"
			cloud_type = "%s"
		}
	`, xIbmTenantID, cloudType)
}
