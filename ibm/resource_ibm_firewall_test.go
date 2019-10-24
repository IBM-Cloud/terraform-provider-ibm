package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMFirewall_Basic(t *testing.T) {
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewall_basic(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfw", "ha_enabled", "false"),
					testAccCheckIBMResources("ibm_firewall.accfw", "public_vlan_id",
						"ibm_compute_vm_instance.fwvm1", "public_vlan_id"),
				),
			},
		},
	})
}

func testAccCheckIBMFirewall_basic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "sjc01"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_firewall" "accfw" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm1.public_vlan_id}"
}`, hostname)
}

func TestAccIBMFirewall_FSA(t *testing.T) {
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewall_FSA(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfwfsa", "ha_enabled", "false"),
					resource.TestCheckResourceAttrSet("ibm_firewall.accfwfsa", "username"),
					resource.TestCheckResourceAttrSet("ibm_firewall.accfwfsa", "password"),
					resource.TestCheckResourceAttrSet("ibm_firewall.accfwfsa", "primary_ip"),
					resource.TestCheckResourceAttrSet("ibm_firewall.accfwfsa", "location"),
					testAccCheckIBMResources("ibm_firewall.accfwfsa", "public_vlan_id",
						"ibm_compute_vm_instance.fwfsavm1", "public_vlan_id"),
				),
			},
		},
	})
}

func testAccCheckIBMFirewall_FSA(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwfsavm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "sjc01"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_firewall" "accfwfsa" {
  firewall_type = "FORTIGATE_SECURITY_APPLIANCE"
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwfsavm1.public_vlan_id}"
}`, hostname)
}

func TestAccIBMFirewall_Tag(t *testing.T) {
	hostname := acctest.RandString(16)
	tags1 := "collectd"
	tags2 := "mesos-master"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewallTag(hostname, tags1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfw", "ha_enabled", "false"),
					testAccCheckIBMResources("ibm_firewall.accfw", "public_vlan_id",
						"ibm_compute_vm_instance.fwvm1", "public_vlan_id"),
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfw", "tags.#", "1"),
					CheckStringSet(
						"ibm_firewall.accfw",
						"tags", []string{tags1},
					),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMFirewallUpdateTag(hostname, tags1, tags2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfw", "ha_enabled", "false"),
					testAccCheckIBMResources("ibm_firewall.accfw", "public_vlan_id",
						"ibm_compute_vm_instance.fwvm1", "public_vlan_id"),
					resource.TestCheckResourceAttr(
						"ibm_firewall.accfw", "tags.#", "2"),
					CheckStringSet(
						"ibm_firewall.accfw",
						"tags", []string{tags1, tags2},
					),
				),
			},
		},
	})
}

func testAccCheckIBMFirewallTag(hostname, tag1 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "sjc01"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_firewall" "accfw" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm1.public_vlan_id}"
  tags = ["%s"]
}`, hostname, tag1)
}

func testAccCheckIBMFirewallUpdateTag(hostname, tag1, tag2 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "sjc01"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_firewall" "accfw" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm1.public_vlan_id}"
  tags = ["%s", "%s"]
}`, hostname, tag1, tag2)
}
