package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMPIInstanceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_instance.testacc_ds_instance", "pi_cloud_instance_id", pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_instance" "testacc_ds_instance" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
