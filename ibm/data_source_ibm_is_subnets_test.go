package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISSubnetsDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_subnets.test1"
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetsDataSourceConfig(vpcname, name, ISZoneName, ISCIDR),
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

func testAccCheckIBMISSubnetsDataSourceConfig(vpcname, name, zone, cidr string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
	name            = "%s"
	vpc             = "${ibm_is_vpc.testacc_vpc.id}"
	zone            = "%s"
	ipv4_cidr_block = "%s"
	}

	data "ibm_is_subnets" "test1" {
	}`, vpcname, name, zone, cidr)
}
