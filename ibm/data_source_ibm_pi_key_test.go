package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIKeyDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPIKeyDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_key.testacc_ds_key", "pi_key_name", pi_key_name),
				),
			},
		},
	})
}

func testAccCheckIBMPIKeyDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_key" "testacc_ds_key" {
    pi_key_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_key_name, pi_cloud_instance_id)

}
