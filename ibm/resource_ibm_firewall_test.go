package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
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
    os_reference_code = "DEBIAN_7_64"
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
