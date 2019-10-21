package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMPIVolumesDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIVolumesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_volumes.testacc_ds_volumes", "pi_volumes_name", pi_volume_name),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumesDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_volumes" "testacc_ds_volumes" {
    pi_volume_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_volume_name, pi_cloud_instance_id)

}
