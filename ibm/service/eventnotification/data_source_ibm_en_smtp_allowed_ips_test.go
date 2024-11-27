// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMEnSMTPAllowedIpsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPAllowedIpsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_allowed_ips.en_smtp_allowed_ips_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_allowed_ips.en_smtp_allowed_ips_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_allowed_ips.en_smtp_allowed_ips_instance", "en_smtp_allowed_ips_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_allowed_ips.en_smtp_allowed_ips_instance", "subnets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_allowed_ips.en_smtp_allowed_ips_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSMTPAllowedIpsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_en_smtp_allowed_ips" "en_smtp_allowed_ips_instance" {
			instance_id = "instance_id"
			id = "id"
		}
	`)
}
