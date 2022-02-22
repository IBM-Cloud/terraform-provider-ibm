// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSccAccountNotificationSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccAccountNotificationSettingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_account_notification_settings.scc_account_notification_settings", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccAccountNotificationSettingsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_account_notification_settings" "scc_account_notification_settings" {
		}
	`)
}
