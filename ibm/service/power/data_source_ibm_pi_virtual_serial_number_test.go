// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVirtualSerialNumber(t *testing.T) {
	vsnData := "data.ibm_pi_virtual_serial_number.testacc_virtual_serial_number"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVirtualSerialNumberConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(vsnData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVirtualSerialNumberConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_virtual_serial_number" "testacc_virtual_serial_number" {
			pi_cloud_instance_id = "%s"
			pi_serial            = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_virtual_serial_number)
}
