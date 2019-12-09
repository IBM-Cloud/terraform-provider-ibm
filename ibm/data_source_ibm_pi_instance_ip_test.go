package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
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
	
data "ibm_pi_instance_ip" "testacc_ds_instance_ip" {
    pi_network_name = "%s"
    pi_instance_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_network_name, pi_instance_name, pi_cloud_instance_id)

}
