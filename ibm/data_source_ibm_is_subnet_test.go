package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISSubnetDatasource_basic(t *testing.T) {
	vpcname := fmt.Sprintf("terraformsubnetuat_vpc_%d", acctest.RandInt())
	name := fmt.Sprintf("terraformvpcuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSubnetConfig(vpcname, name, ISZoneName, ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_subnet.ds_subnet", "name", name),
				),
			},
		},
	})
}

func testDSCheckIBMISSubnetConfig(vpcname, name, zone, cidr string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	zone = "%s"
	ipv4_cidr_block = "%s"
}
data "ibm_is_subnet" "ds_subnet" {
	identifier = "${ibm_is_subnet.testacc_subnet.id}"
}`, vpcname, name, zone, cidr)
}
