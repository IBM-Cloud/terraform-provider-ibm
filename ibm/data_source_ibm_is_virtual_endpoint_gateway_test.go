package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewayDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateway.data_test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", "my-endpoint-gateway-1"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic() + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateway" "data_test" {
        name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
	}`)
}
