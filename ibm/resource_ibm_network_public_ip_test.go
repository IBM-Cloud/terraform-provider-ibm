package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMNetworkPublicIp_Basic(t *testing.T) {
	hostname1 := acctest.RandString(16)
	hostname2 := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpConfig_basic(hostname1, hostname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkPublicIpExists("ibm_network_public_ip.test-global-ip"),
					resource.TestMatchResourceAttr("ibm_network_public_ip.test-global-ip", "ip_address",
						regexp.MustCompile(`^(([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))\.){3}([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))$`)),
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip", "routes_to",
						"ibm_compute_vm_instance.vm1", "ipv4_address"),
					resource.TestCheckResourceAttr("ibm_network_public_ip.test-global-ip", "notes", "public ip notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpConfig_updated(hostname1, hostname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip", "routes_to",
						"ibm_compute_vm_instance.vm2", "ipv4_address"),
					resource.TestCheckResourceAttr("ibm_network_public_ip.test-global-ip", "notes", "updated public ip notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpConfig_Ipv6Basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkPublicIpExists("ibm_network_public_ip.test-global-ip-3"),
					resource.TestMatchResourceAttr("ibm_network_public_ip.test-global-ip-3", "ip_address",
						regexp.MustCompile(`^(([[:xdigit:]]{4}:){7})([[:xdigit:]]{4})$`)),
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip-3", "routes_to",
						"ibm_compute_vm_instance.vm3", "ipv6_address"),
					resource.TestCheckResourceAttr("ibm_network_public_ip.test-global-ip-3", "notes", "public ip notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpConfig_Ipv6Updated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip-3", "routes_to",
						"ibm_compute_vm_instance.vm4", "ipv6_address"),
					resource.TestCheckResourceAttr("ibm_network_public_ip.test-global-ip-3", "notes", "updated public ip notes"),
				),
			},
		},
	})
}

func TestAccIBMNetworkPublicIpWitTag(t *testing.T) {
	hostname1 := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpWithTag(hostname1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkPublicIpExists("ibm_network_public_ip.test-global-ip"),
					resource.TestMatchResourceAttr("ibm_network_public_ip.test-global-ip", "ip_address",
						regexp.MustCompile(`^(([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))\.){3}([01]?[0-9]?[0-9]|2([0-4][0-9]|5[0-5]))$`)),
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip", "routes_to",
						"ibm_compute_vm_instance.vm1", "ipv4_address"),
					resource.TestCheckResourceAttr(
						"ibm_network_public_ip.test-global-ip", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkPublicIpWithUpdatedTag(hostname1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResources("ibm_network_public_ip.test-global-ip", "routes_to",
						"ibm_compute_vm_instance.vm1", "ipv4_address"),
					resource.TestCheckResourceAttr(
						"ibm_network_public_ip.test-global-ip", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkPublicIpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		globalIpId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetNetworkSubnetIpAddressGlobalService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundGlobalIp, err := service.Id(globalIpId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundGlobalIp.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testAccCheckIBMResources(srcResource, srcKey, tgtResource, tgtKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sourceResource, ok := s.RootModule().Resources[srcResource]
		if !ok {
			return fmt.Errorf("Not found: %s", srcResource)
		}

		targetResource, ok := s.RootModule().Resources[tgtResource]
		if !ok {
			return fmt.Errorf("Not found: %s", tgtResource)
		}

		if sourceResource.Primary.Attributes[srcKey] != targetResource.Primary.Attributes[tgtKey] {
			return fmt.Errorf("Different values : Source : %s %s %s , Target : %s %s %s",
				srcResource, srcKey, sourceResource.Primary.Attributes[srcKey],
				tgtResource, tgtKey, targetResource.Primary.Attributes[tgtKey])
		}

		return nil
	}
}

func testAccCheckIBMNetworkPublicIpConfig_basic(hostname1 string, hostname2 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_compute_vm_instance" "vm2" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "tor01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_network_public_ip" "test-global-ip" {
    routes_to = "${ibm_compute_vm_instance.vm1.ipv4_address}"
    notes = "public ip notes"
}
`, hostname1, hostname2)
}

func testAccCheckIBMNetworkPublicIpConfig_updated(hostname1 string, hostname2 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_compute_vm_instance" "vm2" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "tor01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_network_public_ip" "test-global-ip" {
    routes_to = "${ibm_compute_vm_instance.vm2.ipv4_address}"
     notes = "updated public ip notes"
}
`, hostname1, hostname2)
}

const testAccCheckIBMNetworkPublicIpConfig_Ipv6Basic = `
resource "ibm_compute_vm_instance" "vm3" {
    hostname = "vm3"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "che01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_compute_vm_instance" "vm4" {
    hostname = "vm4"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "che01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_network_public_ip" "test-global-ip-3" {
    routes_to = "${ibm_compute_vm_instance.vm3.ipv6_address}"
    notes = "public ip notes"
}`

const testAccCheckIBMNetworkPublicIpConfig_Ipv6Updated = `
resource "ibm_compute_vm_instance" "vm3" {
    hostname = "vm3"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "che01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_compute_vm_instance" "vm4" {
    hostname = "vm4"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "che01"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
    ipv6_enabled = true
}

resource "ibm_network_public_ip" "test-global-ip-3" {
    routes_to = "${ibm_compute_vm_instance.vm4.ipv6_address}"
     notes = "updated public ip notes"
}`

func testAccCheckIBMNetworkPublicIpWithTag(hostname1 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_network_public_ip" "test-global-ip" {
    routes_to = "${ibm_compute_vm_instance.vm1.ipv4_address}"
    tags = ["one", "two"]
}
`, hostname1)
}

func testAccCheckIBMNetworkPublicIpWithUpdatedTag(hostname1 string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "vm1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal06"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_network_public_ip" "test-global-ip" {
    routes_to = "${ibm_compute_vm_instance.vm1.ipv4_address}"
    tags = ["one", "two", "three"]
}
`, hostname1)
}
