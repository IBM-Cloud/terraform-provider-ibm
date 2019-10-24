package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMFirewallPolicy_Basic(t *testing.T) {
	hostname := acctest.RandString(16)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewallPolicy_basic(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.action", "deny"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.src_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_ip_address", "any"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_start", "1"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_end", "65535"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.notes", "Deny all"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.protocol", "tcp"),

					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.action", "permit"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.src_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.dst_ip_address", "any"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.dst_port_range_start", "22"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.dst_port_range_end", "22"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.notes", "Allow SSH"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.action", "permit"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.src_ip_address",
						"0000:0000:0000:0000:0000:0000:0000:0000"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.dst_ip_address", "any"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.dst_port_range_start", "22"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.dst_port_range_end", "22"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.notes", "Allow SSH"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.2.protocol", "tcp"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMFirewallPolicy_update(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.action", "permit"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.src_ip_address", "10.1.1.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_start", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_end", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.notes", "Permit from 10.1.1.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.action", "deny"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.src_ip_address", "2401:c900:1501:0032:0000:0000:0000:0000"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.dst_port_range_start", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.dst_port_range_end", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.notes", "Deny for IPv6"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.1.protocol", "udp"),
				),
			},
		},
	})
}

func TestAccIBMFirewallPolicyWithTag(t *testing.T) {
	hostname := acctest.RandString(16)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewallPolicyWithTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.action", "deny"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.src_ip_address", "0.0.0.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_ip_address", "any"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_start", "1"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_end", "65535"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.notes", "Deny all"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMFirewallPolicyWithUpdatedTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.action", "permit"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.src_ip_address", "10.1.1.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_start", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.dst_port_range_end", "80"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.notes", "Permit from 10.1.1.0"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "rules.0.protocol", "udp"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_policy.rules", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMFirewallPolicy_basic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm2" {
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

resource "ibm_firewall" "accfw2" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm2.public_vlan_id}"
}

resource "ibm_firewall_policy" "rules" {
 firewall_id = "${ibm_firewall.accfw2.id}"
 rules = {
      "action" = "deny"
      "src_ip_address"= "0.0.0.0"
      "src_ip_cidr"= 0
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 1
      "dst_port_range_end"= 65535
      "notes"= "Deny all"
      "protocol"= "tcp"
 }
 rules = {
      "action" = "permit"
      "src_ip_address"= "0.0.0.0"
      "src_ip_cidr"= 0
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 22
      "dst_port_range_end"= 22
      "notes"= "Allow SSH"
      "protocol"= "tcp"
 }
 rules = {
      "action" = "permit"
      "src_ip_address"= "0::"
      "src_ip_cidr"= 0
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 128
      "dst_port_range_start"= 22
      "dst_port_range_end"= 22
      "notes"= "Allow SSH"
      "protocol"= "tcp"
 }
}

`, hostname)
}

func testAccCheckIBMFirewallPolicy_update(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm2" {
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

resource "ibm_firewall" "accfw2" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm2.public_vlan_id}"
}

resource "ibm_firewall_policy" "rules" {
 firewall_id = "${ibm_firewall.accfw2.id}"
 rules = {
      "action" = "permit"
      "src_ip_address"= "10.1.1.0"
      "src_ip_cidr"= 24
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 80
      "dst_port_range_end"= 80
      "notes"= "Permit from 10.1.1.0"
      "protocol"= "udp"
 }
 rules = {
      "action" = "deny"
      "src_ip_address"= "2401:c900:1501:0032:0000:0000:0000:0000"
      "src_ip_cidr"= 64
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 128
      "dst_port_range_start"= 80
      "dst_port_range_end"= 80
      "notes"= "Deny for IPv6"
      "protocol"= "udp"
 }
 
}

`, hostname)
}

func testAccCheckIBMFirewallPolicyWithTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm2" {
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

resource "ibm_firewall" "accfw2" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm2.public_vlan_id}"
}

resource "ibm_firewall_policy" "rules" {
 firewall_id = "${ibm_firewall.accfw2.id}"
 rules = {
      "action" = "deny"
      "src_ip_address"= "0.0.0.0"
      "src_ip_cidr"= 0
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 1
      "dst_port_range_end"= 65535
      "notes"= "Deny all"
      "protocol"= "tcp"
 }
 tags = ["one", "two"]
}

`, hostname)
}

func testAccCheckIBMFirewallPolicyWithUpdatedTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "fwvm2" {
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

resource "ibm_firewall" "accfw2" {
  ha_enabled = false
  public_vlan_id = "${ibm_compute_vm_instance.fwvm2.public_vlan_id}"
}

resource "ibm_firewall_policy" "rules" {
 firewall_id = "${ibm_firewall.accfw2.id}"
 rules = {
      "action" = "permit"
      "src_ip_address"= "10.1.1.0"
      "src_ip_cidr"= 24
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 80
      "dst_port_range_end"= 80
      "notes"= "Permit from 10.1.1.0"
      "protocol"= "udp"
 }
 tags = ["one", "two", "three"]
}

`, hostname)
}
