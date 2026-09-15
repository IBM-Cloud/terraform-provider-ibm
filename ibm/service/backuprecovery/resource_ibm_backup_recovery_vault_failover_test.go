// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryVaultFailoverBasic(t *testing.T) {
	vaultID := os.Getenv("IBMCLOUD_BACKUP_RECOVERY_VAULT_ID")
	tenantID := os.Getenv("IBMCLOUD_BACKUP_RECOVERY_TENANT_ID")

	if vaultID == "" {
		t.Skip("IBMCLOUD_BACKUP_RECOVERY_VAULT_ID must be set for this acceptance test")
	}
	if tenantID == "" {
		t.Skip("IBMCLOUD_BACKUP_RECOVERY_TENANT_ID must be set for this acceptance test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmBackupRecoveryVaultFailoverConfigBasic(vaultID, tenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_failover.backup_recovery_vault_failover_instance", "x_ibm_tenant_id", tenantID),
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_failover.backup_recovery_vault_failover_instance", "cloud_type", "ibm"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_vault_failover.backup_recovery_vault_failover_instance", "uid"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryVaultFailoverConfigBasic(vaultID string, tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_vault_failover" "backup_recovery_vault_failover_instance" {
			x_ibm_tenant_id = "%s"
			cloud_type      = "ibm"
			failover_request_params {
				vault_id = %s
			}
		}
	`, tenantID, vaultID)
}
