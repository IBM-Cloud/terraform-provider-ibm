// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPICloudConnectionNetworkAttachBasic(t *testing.T) {
	name := fmt.Sprintf("tf-ccnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionNetworkAttachConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.cloud_connection"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("data.ibm_pi_cloud_connection.cloud_connection",
						"networks.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionNetworkAttachConfig(name string) string {
	return fmt.Sprintf(`
	  resource "ibm_pi_cloud_connection" "cloud_connection" {
		pi_cloud_instance_id				= "%[1]s"
		pi_cloud_connection_name			= "%[2]s"
		pi_cloud_connection_speed			= 100
	  }
	  resource "ibm_pi_network" "network1" {
		pi_cloud_instance_id	= "%[1]s"
		pi_network_name			= "%[2]s"
		pi_network_type         = "vlan"
		pi_cidr         		= "192.152.61.0/24"
	  }
	  resource "ibm_pi_cloud_connection_network_attach" "example" {
		pi_cloud_instance_id	= "%[1]s"
		pi_cloud_connection_id	= ibm_pi_cloud_connection.cloud_connection.cloud_connection_id
		pi_network_id			= ibm_pi_network.network1.network_id
	  }
	  data "ibm_pi_cloud_connection" "cloud_connection" {
		depends_on					= [ibm_pi_cloud_connection_network_attach.example]
		pi_cloud_instance_id		= "%[1]s"
		pi_cloud_connection_name	= "%[2]s"
	  }
	`, acc.Pi_cloud_instance_id, name)
}
