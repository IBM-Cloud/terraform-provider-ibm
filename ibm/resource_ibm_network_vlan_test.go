package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
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

func TestAccIBMNetworkVlan_With_Multipe_Subnets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfigMultipleSubnets(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "tfuat_mult_subnet"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "subnet_size", "8"),
				),
			},
		},
	})
}

func TestAccIBMNetworkVlan_with_vm(t *testing.T) {

	hostname := acctest.RandString(16)
	domain := "vlan.tfmvmuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkVlanConfigWithVM(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pvt", "name", "tfuat_pvt_subnet"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pvt", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pvt", "type", "PRIVATE"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pvt", "router_hostname", "bcr01a.lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pvt", "subnet_size", "8"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "name", "tfuat_pub_subnet"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "type", "PUBLIC"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "router_hostname", "fcr01a.lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "subnet_size", "8"),
					resource.TestCheckResourceAttr(
						"ibm_compute_vm_instance.vm", "hostname", hostname),
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

func testAccCheckIBMNetworkVlanConfigMultipleSubnets() (config string) {
	return `
	resource "ibm_network_vlan" "test_vlan" {
		name            = "tfuat_mult_subnet"
		datacenter      = "lon02"
		type            = "PRIVATE"
		subnet_size     = 8
		router_hostname = "bcr01a.lon02"
	  }
	  
	  resource "ibm_subnet" "portable_subnet" {
		type       = "Portable"
		private    = true
		ip_version = 4
		capacity   = 4
		vlan_id    = "${ibm_network_vlan.test_vlan.id}"
		notes      = "portable_tfuat"
	  }
	 `
}

func testAccCheckIBMNetworkVlanConfigWithVM(hostname, domain string) (config string) {
	return `
	resource "ibm_network_vlan" "pvt" {
		name            = "tfuat_pvt_subnet"
		datacenter      = "lon02"
		type            = "PRIVATE"
		subnet_size     = 8
		router_hostname = "bcr01a.lon02"
	  }

	  resource "ibm_network_vlan" "pub" {
		name            = "tfuat_pub_subnet"
		datacenter      = "lon02"
		type            = "PUBLIC"
		subnet_size     = 8
		router_hostname = "fcr01a.lon02"
	  }
	 ` +
		fmt.Sprintf(`
		resource "ibm_compute_vm_instance" "vm" {
			hostname = "%s"
			domain = "%s"
			os_reference_code = "DEBIAN_8_64"
			datacenter = "lon02"
			network_speed = 10
			hourly_billing = true
			private_vlan_id = "${ibm_network_vlan.pvt.id}"
			public_vlan_id  = "${ibm_network_vlan.pub.id}"
			private_network_only = false
			cores = 1
			memory = 1024
			disks = [25, 10, 20]
			local_disk = false
			notes = "VM notes"
		}`, hostname, domain)
}
