package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMDLPortsDataSource_basic(t *testing.T) {
	name := "dl_ports"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLPortsDataSourceConfig(name),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccCheckIBMDLPortsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_ports" "test_%s" {
	   }
	  `, name)
}
