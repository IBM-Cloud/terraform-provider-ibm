// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorUpdateUserBasic(t *testing.T) {
	sessionNameCookie := "MTczNzk5MzU1MHxEWDhFQVFMX2dBQUJFQUVRQUFELUFZWF9nQUFKQm5OMGNtbHVad3dSQUE5bGVIQnBjbUYwYVc5dUxYUnBiV1VHYzNSeWFXNW5EQXdBQ2pFM016Z3dOems1TlRBR2MzUnlhVzVuREFnQUJtUnZiV0ZwYmdaemRISnBibWNNQndBRlRFOURRVXdHYzNSeWFXNW5EQXNBQ1dGMWRHZ3RkSGx3WlFaemRISnBibWNNQXdBQk1RWnpkSEpwYm1jTUN3QUpjMmxrY3kxb1lYTm9Cbk4wY21sdVp3d3RBQ3QwYjBKUGNXRktXVWRSV0U1TVFYcGxZM2xOT0hWTFRtOXRjMHhQVldkc1ZWQlNOakExTW1SMllraGpCbk4wY21sdVp3d0tBQWgxYzJWeUxYTnBaQVp6ZEhKcGJtY01JQUFlVXkweExURXdNQzB5TVMweE1EWTNNamMwTWkwek9UY3hNREUxTVMweEJuTjBjbWx1Wnd3TUFBcHBiaTFqYkhWemRHVnlCR0p2YjJ3Q0FnQUJCbk4wY21sdVp3d0hBQVZ5YjJ4bGN3WnpkSEpwYm1jTUVBQU9RMDlJUlZOSlZGbGZRVVJOU1U0R2MzUnlhVzVuREFvQUNIVnpaWEp1WVcxbEJuTjBjbWx1Wnd3SEFBVmhaRzFwYmdaemRISnBibWNNQ0FBR2JHOWpZV3hsQm5OMGNtbHVad3dIQUFWbGJpMTFjdz09fAWBv6tiQJ8KsMnNVBbt5YfQowIauhcANmpE592HJH64"
	sid := "S-1-100-21-10672742-39710151-1"
	username := "admin"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConnectorUpdateUserConfigBasic(sessionNameCookie),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance", "username", username),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_update_user.backup_recovery_connector_update_user_instance", "id", fmt.Sprintf("%s:%s", sid, username)),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorUpdateUserConfigBasic(sessionNameCookie string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_update_user" "backup_recovery_connector_update_user_instance" {
			username = "admin"
  			password = "newPassword7"
  			current_password = "cohesity7" 
  			sid = "S-1-100-21-10672742-39710151-1"
  			roles = ["COHESITY_ADMIN"]
  			session_name_cookie = "%s"
		}
	`, sessionNameCookie)
}
