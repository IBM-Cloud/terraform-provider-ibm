package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMPIVolumeDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIVolumeDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_volume.testacc_ds_volume", "pi_volume_name", pi_volume_name),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_volume" "testacc_ds_volume" {
    pi_volume_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_volume_name, pi_cloud_instance_id)

}
