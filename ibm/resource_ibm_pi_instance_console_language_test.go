// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceConsoleLanguage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
		pi_cloud_instance_id	= "%s"
		pi_instance_name		= "%s"
		pi_language_code		= "%s"
	}
	`, pi_cloud_instance_id, pi_instance_name, code)
}
