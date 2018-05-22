package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMStorageFile_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageFileConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageFileExists("ibm_storage_file.fs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "capacity", "20"),
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "iops", "0.25"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "snapshot_capacity", "10"),
					testAccCheckIBMResources("ibm_storage_file.fs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "notes", "endurance notes"),
					// Performance Storage
					testAccCheckIBMStorageFileExists("ibm_storage_file.fs_performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "type", "Performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "iops", "200"),
					testAccCheckIBMResources("ibm_storage_file.fs_performance", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_performance", "notes", "performance notes"),
					// NAS
					testAccCheckIBMStorageFileExists("ibm_storage_file.nas"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.nas", "type", "NAS/FTP"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.nas", "capacity", "20"),
					testAccCheckIBMResources("ibm_storage_file.nas", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.nas", "notes", "nas notes"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageFileConfig_update,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "allowed_virtual_guest_ids.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "allowed_subnets.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "allowed_ip_addresses.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "notes", "updated endurance notes"),
					// Performance Storage
					resource.TestCheckResourceAttr("ibm_storage_file.fs_performance", "allowed_virtual_guest_ids.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_performance", "allowed_subnets.#", "1"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_performance", "allowed_ip_addresses.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageFileConfig_enablesnapshot,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "snapshot_schedule.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMStorageFileConfig_updatesnapshot,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "snapshot_schedule.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMStorageFile_With_Hourly(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageFileConfigWithHourlyBilling,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageFileExists("ibm_storage_file.fs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "capacity", "20"),
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "iops", "0.25"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "snapshot_capacity", "10"),
					testAccCheckIBMResources("ibm_storage_file.fs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "notes", "endurance for hourly billing"),
					// Performance Storage
					testAccCheckIBMStorageFileExists("ibm_storage_file.fs_performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "type", "Performance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "capacity", "20"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_performance", "iops", "200"),
					testAccCheckIBMResources("ibm_storage_file.fs_performance", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_performance", "notes", "performance for hourly billing"),
				),
			},
		},
	})
}

func TestAccIBMStorageFileWithTag(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMStorageFileWithTag,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					testAccCheckIBMStorageFileExists("ibm_storage_file.fs_endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "type", "Endurance"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "capacity", "20"),
					resource.TestCheckResourceAttrSet("ibm_storage_file.fs_endurance", "mountpoint"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "iops", "0.25"),
					resource.TestCheckResourceAttr(
						"ibm_storage_file.fs_endurance", "snapshot_capacity", "10"),
					testAccCheckIBMResources("ibm_storage_file.fs_endurance", "datacenter",
						"ibm_compute_vm_instance.storagevm1", "datacenter"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "notes", "endurance notes"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMStorageFileWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					// Endurance Storage
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "notes", "endurance notes"),
					resource.TestCheckResourceAttr("ibm_storage_file.fs_endurance", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMStorageFileExists(n string) resource.TestCheckFunc {
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

const testAccCheckIBMStorageFileConfig_basic = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        notes = "endurance notes"
}
resource "ibm_storage_file" "fs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 200
        notes = "performance notes"
}
resource "ibm_storage_file" "nas" {
	type = "NAS/FTP"
	datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
	capacity = 20
	notes = "nas notes"
}
`
const testAccCheckIBMStorageFileConfig_update = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm1.id}" ]
        allowed_subnets = [ "${ibm_compute_vm_instance.storagevm1.private_subnet}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm1.ipv4_address_private}" ]
        snapshot_capacity = 10
        notes = "updated endurance notes"
}
resource "ibm_storage_file" "fs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 100
        allowed_virtual_guest_ids = [ "${ibm_compute_vm_instance.storagevm1.id}" ]
        allowed_subnets = [ "${ibm_compute_vm_instance.storagevm1.private_subnet}" ]
        allowed_ip_addresses = [ "${ibm_compute_vm_instance.storagevm1.ipv4_address_private}" ]
}
resource "ibm_storage_file" "nas" {
		type = "NAS/FTP"
		datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
		capacity = 20
		notes = "nas notes"
}
`

const testAccCheckIBMStorageFileConfig_enablesnapshot = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        snapshot_schedule = [
  		{
			schedule_type="WEEKLY",
			retention_count= 5,
			minute= 2,
			hour= 13,
			day_of_week= "SUNDAY",
			enable= true
		},
		{
			schedule_type="HOURLY",
			retention_count= 5,
			minute= 30,
			enable= true
		},
		
		{
			schedule_type="DAILY",
			retention_count= 6,
			minute= 2,
			hour= 15
			enable= true
		},
 		]
}
`
const testAccCheckIBMStorageFileConfig_updatesnapshot = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        snapshot_schedule = [
  		{
			schedule_type="WEEKLY",
			retention_count= 2,
			minute= 2,
			hour= 13,
			day_of_week= "MONDAY",
			enable= true
		},
		{
			schedule_type="HOURLY",
			retention_count= 3,
			minute= 40,
			enable= true
		},
		
		{
			schedule_type="DAILY",
			retention_count= 5,
			minute= 2,
			hour= 15
			enable= false
		},
 		]
}
`

const testAccCheckIBMStorageFileWithTag = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
		notes = "endurance notes"
		tags = ["one", "two"]
}
`

const testAccCheckIBMStorageFileWithUpdatedTag = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
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
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
		notes = "endurance notes"
		tags = ["one", "two", "three"]
}
`

const testAccCheckIBMStorageFileConfigWithHourlyBilling = `
resource "ibm_compute_vm_instance" "storagevm1" {
    hostname = "storagevm1"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        notes = "endurance for hourly billing"
		hourly_billing = true
}
resource "ibm_storage_file" "fs_performance" {
        type = "Performance"
        datacenter = "${ibm_compute_vm_instance.storagevm1.datacenter}"
        capacity = 20
        iops = 200
        notes = "performance for hourly billing"
		hourly_billing = true
}
`
