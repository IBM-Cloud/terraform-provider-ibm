package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMOrgDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org.testacc_ds_org", "org", cfOrganization),
				),
			},
		},
	})
}

func testAccCheckIBMOrgDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}`, cfOrganization)

}
