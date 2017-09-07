package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMNetworkVlan_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "type", "PUBLIC"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "softlayer_managed", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "router_hostname", "fcr01a.lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "subnet_size", "8"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfig_name_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan_update"),
				),
			},
		},
	})
}

func TestAccIBMNetworkVlan_With_Tag(t *testing.T) {
	fmt.Println("*******")
	tags1 := "collectd"
	tags2 := "mesos-master"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfigWithTag(tags1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "type", "PUBLIC"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "softlayer_managed", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "router_hostname", "fcr01a.lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "subnet_size", "8"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "tags.#", "1"),
					CheckStringSet(
						"ibm_network_vlan.test_vlan",
						"tags", []string{tags1},
					),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfigTagUpdate(tags1, tags2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "test_vlan"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "tags.#", "2"),
					CheckStringSet(
						"ibm_network_vlan.test_vlan",
						"tags", []string{tags1, tags2},
					),
				),
			},
		},
	})
}

const testAccCheckIBMNetworkVlanConfig_basic = `
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan"
   datacenter = "lon02"
   type = "PUBLIC"
   subnet_size = 8
   router_hostname = "fcr01a.lon02"
}`

const testAccCheckIBMNetworkVlanConfig_name_update = `
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan_update"
   datacenter = "lon02"
   type = "PUBLIC"
   subnet_size = 8
   router_hostname = "fcr01a.lon02"
}`

func testAccCheckIBMNetworkVlanConfigWithTag(tag1 string) string {
	return fmt.Sprintf(`
		resource "ibm_network_vlan" "test_vlan" {
			name = "test_vlan"
			datacenter = "lon02"
			type = "PUBLIC"
			subnet_size = 8
			router_hostname = "fcr01a.lon02"
			tags = ["%s"]
		 }`, tag1)
}

func testAccCheckIBMNetworkVlanConfigTagUpdate(tag1, tag2 string) string {
	return fmt.Sprintf(`
	resource "ibm_network_vlan" "test_vlan" {
		name = "test_vlan"
		datacenter = "lon02"
		type = "PUBLIC"
		subnet_size = 8
		router_hostname = "fcr01a.lon02"
		tags = ["%s", "%s"]
	 }`, tag1, tag2)

}
