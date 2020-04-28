package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISSubnetsDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_subnets.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "subnets.0.name"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.zone"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.crn"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.network_acl"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.vpc"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_subnets" "test1" {
      }`)
}
