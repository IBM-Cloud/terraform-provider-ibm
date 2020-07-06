package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPICloudInstanceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPICloudInstanceDataSourceConfig(),
				Check:  resource.ComposeTestCheckFunc(
				//resource.TestCheckResourceAttr("data.ibm_pi_image.testacc_ds_image", "pi_image_name", pi_image),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_cloud_instances" "testacc_ds_cloud_instance" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
