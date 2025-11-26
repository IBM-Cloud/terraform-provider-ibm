// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkSecurityGroupActionBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupActionConfigBasic(power.Enable),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_security_group_action.network_security_group_action", "pi_action", power.Enable),
				),
			},
			{
				Config: testAccCheckIBMPINetworkSecurityGroupActionConfigBasic(power.Disable),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_security_group_action.network_security_group_action", "pi_action", power.Disable),
				),
			},
		},
	})
}

func TestAccIBMPINetworkSecurityGroupActionBasicReverse(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupActionConfigBasic(power.Disable),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_security_group_action.network_security_group_action", "pi_action", power.Disable),
				),
			},
			{
				Config: testAccCheckIBMPINetworkSecurityGroupActionConfigBasic(power.Enable),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_security_group_action.network_security_group_action", "pi_action", power.Enable),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkSecurityGroupActionConfigBasic(action string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group_action" "network_security_group_action" {
			pi_cloud_instance_id = "%[2]s"
			pi_action            = "%[1]s"
		}`, action, acc.Pi_cloud_instance_id)
}
