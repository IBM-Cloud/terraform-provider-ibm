package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMLbaasServerInstanceAttachment_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasServerInstanceAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasServerInstanceAttachmentConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_server_instance_attachment.lbaas_member1", "weight", "20"),
					resource.TestCheckResourceAttrSet("ibm_lbaas_server_instance_attachment.lbaas_member1", "uuid"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_server_instance_attachment.lbaas_member2", "weight", "20"),
					resource.TestCheckResourceAttrSet("ibm_lbaas_server_instance_attachment.lbaas_member2", "uuid"),
				),
			},
			{
				Config: testAccCheckIBMLbaasServerInstanceAttachmentConfig_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas_server_instance_attachment.lbaas_member1", "weight", "40"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas_server_instance_attachment.lbaas_member2", "weight", "40"),
				),
			},
		},
	})
}

func TestAccIBMLbaasServerInstanceAttachment_Dynamic_SI_Attachment(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasServerInstanceAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association_attach(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association_dettach(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMLbaasServerInstanceAttachment_InvalidWeight(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasServerInstanceAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasServerInstanceAttachment_InvalidWeight,
				ExpectError: regexp.MustCompile("must be between 1 and 100"),
			},
		},
	})
}

func TestAccIBMLbaasServerInstanceAttachment_InvalidIPAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasServerInstanceAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMLbaasServerInstanceAttachment_IPAddress,
				ExpectError: regexp.MustCompile("must be a valid ip address"),
			},
		},
	})
}

func testAccCheckIBMLbaasServerInstanceAttachmentDestroy(s *terraform.State) error {
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()
	service := services.GetNetworkLBaaSLoadBalancerService(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_lbaas" {
			continue
		}

		// Try to find the key
		_, err := service.GetLoadBalancer(sl.String(rs.Primary.ID))

		if err == nil {
			return fmt.Errorf("load balancer (%s) to be destroyed still exists", rs.Primary.ID)
		} else if apiErr, ok := err.(sl.Error); ok && apiErr.Exception != NOT_FOUND {
			return fmt.Errorf("Error waiting for load balancer (%s) to be destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}

const testAccCheckIBMLbaasServerInstanceAttachment_IPAddress = `
resource "ibm_lbaas_server_instance_attachment" "lbaas" {
  private_ip_address = "10.9.1726.3452"
  weight             = 10
  lbaas_id           = "90528c28-1516-4e71-8612-42d1602eb006"
}
`

const testAccCheckIBMLbaasServerInstanceAttachment_InvalidWeight = `
resource "ibm_lbaas_server_instance_attachment" "lbaas" {
  private_ip_address = "10.9.1726.3452"
  weight             = 120
  lbaas_id           = "90528c28-1516-4e71-8612-42d1602eb006"
}
`

func testAccCheckIBMLbaasServerInstanceAttachmentConfig_basic(name string) string {
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
resource "ibm_compute_vm_instance" "vm2" {
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
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member1" {
  private_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address_private}"
  weight             = 20
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member2" {
  private_ip_address = "${ibm_compute_vm_instance.vm2.ipv4_address_private}"
  weight             = 20
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
`, lbaasDatacenter, lbaasDatacenter, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
	count = 1
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
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count = 1
  private_ip_address = "${element(ibm_compute_vm_instance.vm1.*.ipv4_address_private,count.index)}"
  weight             = 20
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
`, lbaasDatacenter, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association_attach(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
	count = 2
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
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count = 2
  private_ip_address = "${element(ibm_compute_vm_instance.vm1.*.ipv4_address_private,count.index)}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
`, lbaasDatacenter, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasServerInstanceAttachmentConfig_lbaas_dynamic_association_dettach(name string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
	count = 2
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
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count = 1
  private_ip_address = "${element(ibm_compute_vm_instance.vm1.*.ipv4_address_private,count.index)}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
`, lbaasDatacenter, name, lbaasSubnetId)
}

func testAccCheckIBMLbaasServerInstanceAttachmentConfig_update(name string) string {
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
resource "ibm_compute_vm_instance" "vm2" {
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
  description = "desc-used for terraform uat"
  subnets     = ["%s"]
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member1" {
  private_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address_private}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
resource "ibm_lbaas_server_instance_attachment" "lbaas_member2" {
  private_ip_address = "${ibm_compute_vm_instance.vm2.ipv4_address_private}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}
`, lbaasDatacenter, lbaasDatacenter, name, lbaasSubnetId)
}
