package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMComputeVmInstanceDataSource_basic(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "ds.terraform.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeVmInstanceDataSourceConfigBasic(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_compute_vm_instance.tf-vg-ds-acc-test", "power_state", "RUNNING"),
					resource.TestCheckResourceAttr("data.ibm_compute_vm_instance.tf-vg-ds-acc-test", "status", "ACTIVE"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeVmInstanceDataSourceConfigBasic(hostname, domain string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "tf-vg-acc-test" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    tags = ["data-source-test"]
    dedicated_acct_host_only = true
    local_disk = false
}
data "ibm_compute_vm_instance" "tf-vg-ds-acc-test" {
    hostname = "${ibm_compute_vm_instance.tf-vg-acc-test.hostname}"
	domain = "${ibm_compute_vm_instance.tf-vg-acc-test.domain}"
}`, hostname, domain)
}
