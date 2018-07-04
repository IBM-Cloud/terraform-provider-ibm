package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMLbaasDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "name", name),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "protocols.0.backend_port", "80"),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "protocols.0.backend_protocol", "HTTP"),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "protocols.0.frontend_protocol", "HTTP"),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "protocols.0.load_balancing_method", "round_robin"),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "protocols.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_lbaas.tfacc_lbaas", "server_instances.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_server_instance_attachment.lbaas_member", "weight", "20"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckIBMLbaasDataSourceConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
    hostname = "lbass-test"
    os_reference_code = "CENTOS_7_64"
    domain = "terraform.com"
    datacenter = "%s"
    network_speed = "10"
    hourly_billing = true
    private_network_only = false
    cores = "1"
    memory = "1024"
    disks = ["25"]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
}
resource "ibm_lbaas" "lbaas" {
  name        = "%s"
  description = "updated desc-used for terraform uat"
  subnets     = ["%s"]

  protocols = [{

    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80

    "load_balancing_method" = "round_robin"
  }]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  private_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address_private}"
  weight             = 20
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
data "ibm_lbaas" "tfacc_lbaas" {
    name = "${ibm_lbaas.lbaas.name}"
    depends_on = ["ibm_lbaas_server_instance_attachment.lbaas_member"]
}
`, lbaasDatacenter, name, lbaasSubnetId)
}
