package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewaysDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateways.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic() + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateways" "test1" {
		depends_on = [ibm_is_virtual_endpoint_gateway.endpoint_gateway]
	}`)
}
