// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorGetUsersDataSourceBasic(t *testing.T) {
	sessionNameCookie := "MTczNzk4MTgxMHxEWDhFQVFMX2dBQUJFQUVRQUFELUFZWF9nQUFKQm5OMGNtbHVad3dLQUFoMWMyVnlibUZ0WlFaemRISnBibWNNQndBRllXUnRhVzRHYzNSeWFXNW5EQWdBQm14dlkyRnNaUVp6ZEhKcGJtY01Cd0FGWlc0dGRYTUdjM1J5YVc1bkRBc0FDV0YxZEdndGRIbHdaUVp6ZEhKcGJtY01Bd0FCTVFaemRISnBibWNNQ2dBSWRYTmxjaTF6YVdRR2MzUnlhVzVuRENBQUhsTXRNUzB4TURBdE1qRXRNVEEyTnpJM05ESXRNemszTVRBeE5URXRNUVp6ZEhKcGJtY01Dd0FKYzJsa2N5MW9ZWE5vQm5OMGNtbHVad3d0QUN0MGIwSlBjV0ZLV1VkUldFNU1RWHBsWTNsTk9IVkxUbTl0YzB4UFZXZHNWVkJTTmpBMU1tUjJZa2hqQm5OMGNtbHVad3dIQUFWeWIyeGxjd1p6ZEhKcGJtY01FQUFPUTA5SVJWTkpWRmxmUVVSTlNVNEdjM1J5YVc1bkRCRUFEMlY0Y0dseVlYUnBiMjR0ZEdsdFpRWnpkSEpwYm1jTURBQUtNVGN6T0RBMk9ESXhNQVp6ZEhKcGJtY01EQUFLYVc0dFkyeDFjM1JsY2dSaWIyOXNBZ0lBQVFaemRISnBibWNNQ0FBR1pHOXRZV2x1Qm5OMGNtbHVad3dIQUFWTVQwTkJUQT09fPjantRCGr1LJDC_-6kSdwW-sLcWRqWdy8RAMldMkc5n"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorGetUsersDataSourceConfigBasic(sessionNameCookie),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_get_users.backup_recovery_connector_get_users_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_connector_get_users.backup_recovery_connector_get_users_instance", "users.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorGetUsersDataSourceConfigBasic(sessionNameCookie string) string {
	return fmt.Sprintf(`

		data "ibm_backup_recovery_connector_get_users" "backup_recovery_connector_get_users_instance" {
			session_name_cookie = "%s"
			usernames = ["admin"]
		}
	`, sessionNameCookie)
}
