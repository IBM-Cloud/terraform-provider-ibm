// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMNetworkVlan_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
				),
			},

			{
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
						"ibm_network_vlan.test_vlan", "tags.#", "1"),
					CheckStringSet(
						"ibm_network_vlan.test_vlan",
						"tags", []string{tags1},
					),
				),
			},

			{
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMNetworkVlanConfigMultipleSubnets(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.test_vlan", "name", "tfuat_mult_subnet"),
				),
			},
		},
	})
}

func TestAccIBMNetworkVlan_with_vm(t *testing.T) {

	hostname := acctest.RandString(16)
	domain := "vlan.tfmvmuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
						"ibm_network_vlan.pub", "name", "tfuat_pub_subnet"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "datacenter", "lon02"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "type", "PUBLIC"),
					resource.TestCheckResourceAttr(
						"ibm_network_vlan.pub", "router_hostname", "fcr01a.lon02"),
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
   router_hostname = "fcr01a.lon02"
}`

const testAccCheckIBMNetworkVlanConfig_name_update = `
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan_update"
   datacenter = "lon02"
   type = "PUBLIC"
   router_hostname = "fcr01a.lon02"
}`

func testAccCheckIBMNetworkVlanConfigWithTag(tag1 string) string {
	return fmt.Sprintf(`
		resource "ibm_network_vlan" "test_vlan" {
			name = "test_vlan"
			datacenter = "lon02"
			type = "PUBLIC"
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
		router_hostname = "bcr01a.lon02"
	  }

	  resource "ibm_network_vlan" "pub" {
		name            = "tfuat_pub_subnet"
		datacenter      = "lon02"
		type            = "PUBLIC"
		router_hostname = "fcr01a.lon02"
	  }
	 ` +
		fmt.Sprintf(`
		resource "ibm_compute_vm_instance" "vm" {
			hostname = "%s"
			domain = "%s"
			os_reference_code = "DEBIAN_9_64"
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
