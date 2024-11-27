// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPIAvailableHostsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIAvailableHostsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_available_hosts.pi_available_hosts_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIAvailableHostsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_available_hosts" "pi_available_hosts_instance" {
			pi_cloud_instance_id = "%s"
		}
	`, acc.Pi_cloud_instance_id)
}
