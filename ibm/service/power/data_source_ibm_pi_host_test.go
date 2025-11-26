// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPIHostDataSourceBasic(t *testing.T) {
	hostResData := "data.ibm_pi_host.pi_host_instance"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIHostDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hostResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIHostDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_host" "pi_host_instance" {
			pi_cloud_instance_id = "%s"
			pi_host_id = "%s"
		}
	`, acc.Pi_cloud_instance_id, acc.Pi_host_id)
}
