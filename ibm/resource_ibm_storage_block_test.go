package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMStorageBlock_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageBlockConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "iops", "0.25"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "snapshot_capacity", "10"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "os_format_type", "Linux"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "notes", "endurance notes"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "hourly_billing", "false"),
					testAccCheckIBMResources("ibm_storage_block.bs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm2", "datacenter"),
					// Performance Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "type", "Performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "iops", "100"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "os_format_type", "Linux"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "hourly_billing", "false"),
					testAccCheckIBMResources("ibm_storage_block.bs_performance", "datacenter",
						"ibm_compute_vm_instance.storagevm2", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_performance", "notes", "performance notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageBlockConfig_update,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_virtual_guest_ids.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_ip_addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_host_info.#", "2"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "notes", "updated endurance notes"),
					// Performance Storage
					resource.TestCheckResourceAttr("ibm_storage_block.bs_performance", "allowed_virtual_guest_ids.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_performance", "allowed_ip_addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_performance", "allowed_host_info.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMStorageBlock_Portable(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageBlockConfigPortable,
				Check: resource.ComposeTestCheckFunc(
					// Portable Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_portable"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_portable", "type", "Portable"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_portable", "type", "Portable"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_portable", "disk_description", "Checking Portable Storage"),
				),
			},
		},
	})
}

func TestAccIBMStorageBlockwithTag(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageBlockWithTag,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "iops", "0.25"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "snapshot_capacity", "10"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "os_format_type", "Linux"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "notes", "endurance notes"),
					testAccCheckIBMResources("ibm_storage_block.bs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm2", "datacenter"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageBlockWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "notes", "endurance notes"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMStorageBlock_hourly(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageBlockConfig_hourly,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "capacity", "40"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "iops", "2"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "snapshot_capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "os_format_type", "Linux"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "notes", "endurance notes"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_endurance", "hourly_billing", "true"),
					testAccCheckIBMResources("ibm_storage_block.bs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm2", "datacenter"),
					// Performance Storage
					testAccCheckIBMStorageBlockExists("ibm_storage_block.bs_performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "type", "Performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "capacity", "40"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "iops", "100"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "os_format_type", "Linux"),
					resource.TestCheckResourceAttr(
						"ibm_storage_block.bs_performance", "hourly_billing", "true"),
					testAccCheckIBMResources("ibm_storage_block.bs_performance", "datacenter",
						"ibm_compute_vm_instance.storagevm2", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_performance", "notes", "performance notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageBlockConfig_hourly_update,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_virtual_guest_ids.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_ip_addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "allowed_host_info.#", "2"),
					resource.TestCheckResourceAttr("ibm_storage_block.bs_endurance", "notes", "updated endurance notes"),
				),
			},
		},
	})
}

func testAccCheckIBMStorageBlockExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		storageId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetNetworkStorageService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundStorage, err := service.Id(storageId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundStorage.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

const testAccCheckIBMStorageBlockConfig_basic = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
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
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        os_format_type = "Linux"
        notes = "endurance notes"
}
resource "ibm_storage_block" "bs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 100
        os_format_type = "Linux"
        notes = "performance notes"
}
`
const testAccCheckIBMStorageBlockConfigPortable = `resource "ibm_storage_block" "bs_portable" {
	datacenter = "dal13"
	type = "Portable"
	capacity = 10
	disk_description = "Checking Portable Storage"
   }
   `
const testAccCheckIBMStorageBlockConfig_update = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
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
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 0.25
        os_format_type = "Linux"
        allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm2.id}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm2.ipv4_address_private}" ]
        snapshot_capacity = 10
        notes = "updated endurance notes"
}
resource "ibm_storage_block" "bs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 100
        os_format_type = "Linux"
        allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm2.id}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm2.ipv4_address_private}" ]
}
`

const testAccCheckIBMStorageBlockWithTag = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
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
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        os_format_type = "Linux"
		notes = "endurance notes"
		tags = ["one", "two"]
}
`

const testAccCheckIBMStorageBlockWithUpdatedTag = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
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
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        os_format_type = "Linux"
		notes = "endurance notes"
		tags = ["one", "two", "thre"]
}
`

const testAccCheckIBMStorageBlockConfig_hourly = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal10"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 40
        iops = 2
        snapshot_capacity = 20
        os_format_type = "Linux"
		notes = "endurance notes"
		hourly_billing = true
}
resource "ibm_storage_block" "bs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 40
        iops = 100
        os_format_type = "Linux"
		notes = "performance notes"
		hourly_billing = true
		allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm2.id}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm2.ipv4_address_private}" ]
}
`

const testAccCheckIBMStorageBlockConfig_hourly_update = `
resource "ibm_compute_vm_instance" "storagevm2" {
    hostname = "storagevm2"
    domain = "example.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal10"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}
resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm2.datacenter}"
        capacity = 40
        iops = 2
        os_format_type = "Linux"
        allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm2.id}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm2.ipv4_address_private}" ]
        snapshot_capacity = 20
		notes = "updated endurance notes"
		hourly_billing = true
}
`
