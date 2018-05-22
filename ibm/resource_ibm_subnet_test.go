package ibm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccibmSubnet_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSubnetConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					// Check portable IPv4
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "type", "Portable"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "network_type", "private"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "ip_version", "4"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "capacity", "4"),
					testAccCheckIBMResources("ibm_subnet.portable_subnet", "vlan_id",
						"ibm_compute_vm_instance.subnetvm1", "private_vlan_id"),
					resource.TestMatchResourceAttr("ibm_subnet.portable_subnet", "subnet_cidr",
						regexp.MustCompile(`^(([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))\.){3}([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))`)),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "notes", "portable_subnet"),

					// Check static IPv4
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "type", "Static"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "network_type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "ip_version", "4"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "capacity", "4"),
					testAccCheckIBMResources("ibm_subnet.static_subnet", "endpoint_ip",
						"ibm_compute_vm_instance.subnetvm1", "ipv4_address"),
					resource.TestMatchResourceAttr("ibm_subnet.static_subnet", "subnet_cidr",
						regexp.MustCompile(`^(([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))\.){3}([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))`)),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "notes", "static_subnet"),

					// Check portable IPv6
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet_v6", "type", "Portable"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet_v6", "network_type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet_v6", "ip_version", "6"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet_v6", "capacity", "64"),
					testAccCheckIBMResources("ibm_subnet.portable_subnet_v6", "vlan_id",
						"ibm_compute_vm_instance.subnetvm1", "public_vlan_id"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet_v6", "notes", "portable_subnet"),
					// Check static IPv6
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet_v6", "type", "Static"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet_v6", "network_type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet_v6", "ip_version", "6"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet_v6", "capacity", "64"),
					testAccCheckIBMResources("ibm_subnet.static_subnet_v6", "endpoint_ip",
						"ibm_compute_vm_instance.subnetvm1", "ipv6_address"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet_v6", "notes", "static_subnet"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMSubnetConfigNotesUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "notes", "portable_subnet_updated"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.static_subnet", "notes", "static_subnet_updated"),
				),
			},
		},
	})
}

func TestAccibmSubnet_With_Tag(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSubnetConfigWithTag,
				Check: resource.ComposeTestCheckFunc(
					// Check portable IPv4
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "type", "Portable"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "network_type", "public"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "ip_version", "4"),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "capacity", "4"),
					testAccCheckIBMResources("ibm_subnet.portable_subnet", "vlan_id",
						"ibm_compute_vm_instance.subnetvm1", "private_vlan_id"),
					resource.TestMatchResourceAttr("ibm_subnet.portable_subnet", "subnet_cidr",
						regexp.MustCompile(`^(([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))\.){3}([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))`)),
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "notes", "portable_subnet"),
					resource.TestCheckResourceAttr("ibm_subnet.portable_subnet", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMSubnetConfigWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_subnet.portable_subnet", "notes", "portable_subnet"),
					resource.TestCheckResourceAttr("ibm_subnet.portable_subnet", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMSubnetDestroy(s *terraform.State) error {
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()
	service := services.GetNetworkSubnetService(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_subnet" {
			continue
		}

		subnetID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
		}

		// Try to find the key
		_, err = service.Id(subnetID).GetObject()

		if err == nil {
			return fmt.Errorf("Subnet (%s) to be destroyed still exists", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for subnet (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

const testAccCheckIBMSubnetConfigBasic = `
resource "ibm_compute_vm_instance" "subnetvm1" {
    hostname = "subnetvm1"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "wdc04"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true

    lifecycle {
        ignore_changes = ["ipv6_static_enabled"]
    }
}

resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  network_type = "private"
  ip_version = 4
  capacity = 4
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.private_vlan_id}"
  notes = "portable_subnet"
}

resource "ibm_subnet" "static_subnet" {
  type = "Static"
  network_type = "public"
  ip_version = 4
  capacity = 4
  endpoint_ip="${ibm_compute_vm_instance.subnetvm1.ipv4_address}"
  notes = "static_subnet"
}

resource "ibm_subnet" "portable_subnet_v6" {
  type = "Portable"
  network_type = "public"
  ip_version = 6
  capacity = 64
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.public_vlan_id}"
  notes = "portable_subnet"
}

resource "ibm_subnet" "static_subnet_v6" {
  type = "Static"
  network_type = "public"
  ip_version = 6
  capacity = 64
  endpoint_ip="${ibm_compute_vm_instance.subnetvm1.ipv6_address}"
  notes = "static_subnet"
}
`

const testAccCheckIBMSubnetConfigNotesUpdate = `
resource "ibm_compute_vm_instance" "subnetvm1" {
    hostname = "subnetvm1"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "wdc04"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
        
    lifecycle {
        ignore_changes = ["ipv6_static_enabled"] 
    }
}

resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  network_type = "private"
  ip_version = 4
  capacity = 4
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.private_vlan_id}"
  notes = "portable_subnet_updated"
}

resource "ibm_subnet" "static_subnet" {
  type = "Static"
  network_type = "public"
  ip_version = 4
  capacity = 4
  endpoint_ip="${ibm_compute_vm_instance.subnetvm1.ipv4_address}"
  notes = "static_subnet_updated"
}

resource "ibm_subnet" "portable_subnet_v6" {
  type = "Portable"
  network_type = "public"
  ip_version = 6
  capacity = 64
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.public_vlan_id}"
  notes = "portable_subnet"
}

resource "ibm_subnet" "static_subnet_v6" {
  type = "Static"
  network_type = "public"
  ip_version = 6
  capacity = 64
  endpoint_ip="${ibm_compute_vm_instance.subnetvm1.ipv6_address}"
  notes = "static_subnet"
}
`

const testAccCheckIBMSubnetConfigWithTag = `
resource "ibm_compute_vm_instance" "subnetvm1" {
    hostname = "subnetvm1"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "wdc04"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  network_type = "private"
  ip_version = 4
  capacity = 4
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.private_vlan_id}"
  notes = "portable_subnet"
  tags = ["one", "two"]
}
`

const testAccCheckIBMSubnetConfigWithUpdatedTag = `
resource "ibm_compute_vm_instance" "subnetvm1" {
    hostname = "subnetvm1"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "wdc04"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  network_type = "private"
  ip_version = 4
  capacity = 4
  vlan_id = "${ibm_compute_vm_instance.subnetvm1.private_vlan_id}"
  notes = "portable_subnet"
  tags = ["one", "two", "three"]
}
`
