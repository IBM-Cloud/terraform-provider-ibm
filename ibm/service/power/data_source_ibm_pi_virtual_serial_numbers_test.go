// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVirtualSerialNumbers(t *testing.T) {
	vsnData := "data.ibm_pi_virtual_serial_numbers.testacc_virtual_serial_numbers"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVirtualSerialNumbersConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(vsnData, "id"),
					resource.TestCheckResourceAttrSet(vsnData, "virtual_serial_numbers.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVirtualSerialNumbersConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_virtual_serial_numbers" "testacc_virtual_serial_numbers" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
