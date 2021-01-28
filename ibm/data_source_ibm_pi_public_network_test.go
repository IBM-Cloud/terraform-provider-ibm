package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIPublicNetworkDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIPublicNetworkDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_public_network.testacc_ds_public_network", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIPublicNetworkDataSourceConfig(name string) string {
	return testAccCheckIBMPINetworkConfig(name) + fmt.Sprintf(`
data "ibm_pi_public_network" "testacc_ds_public_network" {
    pi_network_name = ibm_pi_network.power_networks.network_id
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
