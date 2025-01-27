// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorGetUsersDataSourceBasic(t *testing.T) {
	userDetailsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorGetUsersDataSourceConfigBasic(userDetailsXIBMTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_get_users.backup_recovery_connector_get_users_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_get_users.backup_recovery_connector_get_users_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorGetUsersDataSourceConfigBasic(userDetailsXIBMTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_update_user" "backup_recovery_connector_update_user_instance" {
			x_ibm_tenant_id = "%s"
		}

		data "ibm_backup_recovery_connector_get_users" "backup_recovery_connector_get_users_instance" {
			x_ibm_tenant_id = ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance.x_ibm_tenant_id
			tenant_ids = [ "tenantIds" ]
			all_under_hierarchy = true
			usernames = [ "usernames" ]
			email_addresses = [ "emailAddresses" ]
			domain = ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance.domain
			partial_match = true
		}
	`, userDetailsXIBMTenantID)
}
