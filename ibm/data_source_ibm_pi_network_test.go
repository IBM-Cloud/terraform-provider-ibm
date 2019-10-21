package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMPINetworkDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPINetworkDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_network.testacc_ds_network", "pi_network_name", pi_network_name),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_network" "testacc_ds_network" {
    pi_network_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_network_name, pi_cloud_instance_id)

}
