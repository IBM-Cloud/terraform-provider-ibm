// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceConsoleLanguage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceConsoleLanguageConfig("037"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_console_language.example", "pi_language_code", "037"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceConsoleLanguageConfig("e1399"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_console_language.example", "pi_language_code", "e1399"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceConsoleLanguageConfig(code string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_console_language" "example" {
			pi_cloud_instance_id	= "%[1]s"
			pi_instance_name		= "%[2]s"
			pi_language_code		= "%[3]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_name, code)
}
