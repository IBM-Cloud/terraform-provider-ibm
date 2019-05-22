package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMStorageEvault_Basic(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "terraformuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageEvaultConfigBasic(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMStorageEvaultExists("ibm_storage_evault.evault"),
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "capacity", "20"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "service_resource_name"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "username"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMStorageEvaultConfigUpdate(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "capacity", "30"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "service_resource_name"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "username"),
				),
			},
		},
	})
}

func TestAccIBMStorageEvault_Import(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "terraformuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageEvaultConfigImport(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMStorageEvaultExists("ibm_storage_evault.evault"),
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "datacenter", "dal05"),
					resource.TestCheckResourceAttr(
						"ibm_storage_evault.evault", "capacity", "20"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "service_resource_name"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "username"),
					resource.TestCheckResourceAttrSet("ibm_storage_evault.evault", "password"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_storage_evault.evault",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMStorageEvaultExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		evaultID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		service := services.GetNetworkStorageBackupEvaultService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundEvault, err := service.Id(evaultID).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundEvault.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testAccCheckIBMStorageEvaultConfigBasic(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "evaultvm1" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal05"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}
resource "ibm_storage_evault" "evault" {
	datacenter          = "${ibm_compute_vm_instance.evaultvm1.datacenter}"
	capacity            = "20"
	virtual_instance_id = "${ibm_compute_vm_instance.evaultvm1.id}"
  }
  `, hostname, domain)
}

func testAccCheckIBMStorageEvaultConfigUpdate(hostname, domain string) string {
	return fmt.Sprintf(`
  resource "ibm_compute_vm_instance" "evaultvm1" {
	  hostname = "%s"
	  domain = "%s"
	  os_reference_code = "DEBIAN_8_64"
	  datacenter = "dal05"
	  network_speed = 100
	  hourly_billing = true
	  private_network_only = false
	  cores = 1
	  memory = 1024
	  disks = [25]
	  local_disk = false
  }
  resource "ibm_storage_evault" "evault" {
	  datacenter          = "${ibm_compute_vm_instance.evaultvm1.datacenter}"
	  capacity            = "30"
	  virtual_instance_id = "${ibm_compute_vm_instance.evaultvm1.id}"
	}
	`, hostname, domain)
}

func testAccCheckIBMStorageEvaultConfigImport(hostname, domain string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "evaultvm1" {
		hostname = "%s"
		domain = "%s"
		os_reference_code = "DEBIAN_8_64"
		datacenter = "dal05"
		network_speed = 100
		hourly_billing = false
		private_network_only = false
		cores = 1
		memory = 1024
		disks = [25]
		local_disk = false
	}
	resource "ibm_storage_evault" "evault" {
		datacenter          = "${ibm_compute_vm_instance.evaultvm1.datacenter}"
		capacity            = "20"
		virtual_instance_id = "${ibm_compute_vm_instance.evaultvm1.id}"
	  }
	  `, hostname, domain)
}
