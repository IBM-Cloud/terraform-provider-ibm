// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccInstanceSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccInstanceSettingsDataSourceConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_instance_settings.scc_instance_settings_tf", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_instance_settings.scc_instance_settings_tf", "event_notifications.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_instance_settings.scc_instance_settings_tf", "object_storage.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccInstanceSettingsDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_instance_settings" "scc_instance_settings_tf" {
			instance_id = "%s"
		}
	`, instanceID)
}
