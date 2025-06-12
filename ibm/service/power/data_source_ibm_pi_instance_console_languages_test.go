// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceConsoleLanguages(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceConsoleLanguagesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_console_languages.example", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_console_languages.example", "console_languages.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceConsoleLanguagesConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_console_languages" "example" {
			pi_cloud_instance_id = "%s"
			pi_instance_name     = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_name)
}
