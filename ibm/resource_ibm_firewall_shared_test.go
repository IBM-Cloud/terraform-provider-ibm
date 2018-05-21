package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMFirewallShared_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewallShared_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_firewall_shared.test_firewall", "firewall_type", "10MBPS_HARDWARE_FIREWALL"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_shared.test_firewall", "guest_id", "1234567"),
					resource.TestCheckResourceAttr(
						"ibm_firewall_shared.test_firewall", "guest_type", "virtual machine"),
				),
			},
		},
	})
}

const testAccCheckIBMFirewallShared_basic = `
resource "ibm_compute_vm_instance" "fwvm1" {
    hostname = "testing"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_7_64"
    datacenter = "sjc01"
    network_speed = 10
    hourly_billing = false
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_firewall_shared" "test_firewall" {
	firewall_type = "HARDWARE_FIREWALL_DEDICATED"
	guest_id = "${ibm_compute_vm_instance.fwvm1.id}"
	guest_type="virtual machine"}`
