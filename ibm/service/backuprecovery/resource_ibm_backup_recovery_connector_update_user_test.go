// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorUpdateUserBasic(t *testing.T) {

	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConnectorUpdateUserConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorUpdateUserConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_update_user" "backup_recovery_connector_update_user_instance" {}
	`)
}
