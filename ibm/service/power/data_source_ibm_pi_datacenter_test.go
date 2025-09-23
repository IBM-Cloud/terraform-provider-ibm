// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIDatacenterDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDatacenterDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_datacenter.test", "id"),
				),
			},
		},
	})
}

func TestAccIBMPIDatacenterDataSourcePrivate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDatacenterDataSourcePrivateConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_datacenter.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDatacenterDataSourceConfig() string {
	return `
		data "ibm_pi_datacenter" "test" {
			pi_datacenter_zone = "dal12"
		}`
}

func testAccCheckIBMPIDatacenterDataSourcePrivateConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_datacenter" "test" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
