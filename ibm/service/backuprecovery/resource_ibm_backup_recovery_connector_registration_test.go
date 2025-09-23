// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorRegistrationBasic(t *testing.T) {
	username := "admin"
	password := "newpass1"
	registrationToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbHVzdGVyX2VuZHBvaW50IjoiY2NhZmRiZDktMWY2Mi00MjUwLWI0MWYtMjY3N2ZmZTM0NmU4LnByaXZhdGUudXMtZWFzdC5iYWNrdXAtcmVjb3ZlcnktdGVzdHMuY2xvdWQuaWJtLmNvbSIsImNvbm5faWQiOjU1MzgxNjY2ODU3MTg5MTQwNDgsImV4cCI6MTczODIzNTAyMH0.wg5mRavfnM-t7P_sdNv7mqASdnHixDPQy1-UkMnzW7_Idi-eK2rtfc4Yn-9Gr8C35AGDQgkHflWMXSzef3xoWXSxp0JAW8eREz87Ux7TIur_UviCptIBBSAk17atUvE58HZbB1reqz2yheEVd58aw_Sy3p28sLV9SCzFPTZpO057hrQ9JiwN5Rp0ZWw9qqe6g7wqk093B70OAXmdrd09XlVBq-u1_krymPRpnPt2gw5XR3Ybb8zowYJYy5hMC9y281jNdenncRX2DZeEJ5ySqzjDLNUqJiag2DB6RFwJpPYqADX8SdQAWIuXk34Yis1YatSn7b7vtf0pLZ8N62sJHQ"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConnectorRegistrationConfigBasic(username, password, registrationToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "registration_token", registrationToken),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorRegistrationConfigBasic(username, password, registrationToken string) string {
	return fmt.Sprintf(`

		resource "ibm_backup_recovery_connector_access_token" "backup_recovery_connector_access_token_instance" {
			username = "%s"
			password = "%s"
		}
		resource "ibm_backup_recovery_connector_registration" "backup_recovery_connector_registration_instance" {
			access_token = resource.ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance.access_token
			registration_token = "%s"
		}
	`, username, password, registrationToken)
}
