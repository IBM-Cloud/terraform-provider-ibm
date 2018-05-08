package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMNetworkVlanDataSource_Basic(t *testing.T) {

	name := fmt.Sprintf("terraformuat_vlan_%s", acctest.RandString(2))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNetworkVlanDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResources("data.ibm_network_vlan.tfacc_vlan", "number",
						"ibm_network_vlan.test_vlan_private", "vlan_number"),
					//resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "number", number),
					resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "name", name),
					resource.TestMatchResourceAttr("data.ibm_network_vlan.tfacc_vlan", "id", regexp.MustCompile("^[0-9]+$")),
					resource.TestCheckResourceAttr("data.ibm_network_vlan.tfacc_vlan", "subnets.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkVlanDataSourceConfig(name string) string {
	return fmt.Sprintf(`
    resource "ibm_network_vlan" "test_vlan_private" {
    name            = "%s"
    datacenter      = "dal06"
    type            = "PRIVATE"
    subnet_size     = 8
    
}
data "ibm_network_vlan" "tfacc_vlan" {
    number = "${ibm_network_vlan.test_vlan_private.vlan_number}"
    name = "${ibm_network_vlan.test_vlan_private.name}"
}`, name)
}
