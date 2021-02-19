/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIInstanceIPDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIInstanceIPDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_network_name", pi_network_name),
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_instance_name", pi_instance_name),
					resource.TestCheckResourceAttr("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "pi_cloud_instance_id", pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceIPDataSourceConfig() string {
	return fmt.Sprintf(`
	resource "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[3]s"
		pi_network_name      = "%[1]s"
		pi_network_type      = "pub-vlan"
	}
data "ibm_pi_instance_ip" "testacc_ds_instance_ip" {
    pi_network_name = "%[1]s"
    pi_instance_name = "%[2]s"
    pi_cloud_instance_id = "%[3]s"
}`, pi_network_name, pi_instance_name, pi_cloud_instance_id)

}
