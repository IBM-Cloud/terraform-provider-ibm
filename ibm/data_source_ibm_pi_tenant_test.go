package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMPITenantDataSource_basic(t *testing.T) {

	name := "Trial Tenant"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPITenantDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_tenant.testacc_ds_network", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPITenantDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
data "ibm_pi_tenant" "testacc_ds_tenant" {
    name = "%s"
}`, name)

}
