package ibm

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func init() {
	imageID := os.Getenv("IBM_COMPUTE_VM_INSTANCE_IMAGE_ID")
	if imageID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_COMPUTE_VM_INSTANCE_IMAGE_ID for testing " +
			"the ibm_compute_vm_instance resource. The image should be replicated in the Washington 4 datacenter. Some tests for that resource will fail if this is not set correctly")
	}
}

func TestAccIBMComputeVmInstance_basic(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "terraformvmuat.ibm.com"
	networkSpeed1 := "10"
	networkSpeed2 := "100"
	cores1 := "1"
	cores2 := "2"
	memory1 := "1024"
	memory2 := "2048"
	tags1 := "collectd"
	tags2 := "mesos-master"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	userMetadata1Unquoted, _ := strconv.Unquote(`"` + userMetadata1 + `"`)
	userMetadata2 := "updated"

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccIBMComputeVmInstanceConfigBasic(hostname, domain, networkSpeed1, cores1, memory1, userMetadata1, tags1),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "datacenter", "wdc04"),
					resource.TestCheckResourceAttr(
						configInstance, "network_speed", networkSpeed1),
					resource.TestCheckResourceAttr(
						configInstance, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configInstance, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configInstance, "cores", cores1),
					resource.TestCheckResourceAttr(
						configInstance, "memory", memory1),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "25"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.2", "20"),
					resource.TestCheckResourceAttr(
						configInstance, "user_metadata", userMetadata1Unquoted),
					resource.TestCheckResourceAttr(
						configInstance, "local_disk", "false"),
					resource.TestCheckResourceAttr(
						configInstance, "dedicated_acct_host_only", "true"),
					CheckStringSet(
						configInstance,
						"tags", []string{tags1},
					),
					resource.TestCheckResourceAttrSet(
						configInstance, "ipv6_enabled"),
					resource.TestCheckResourceAttrSet(
						configInstance, "ipv6_address"),
					resource.TestCheckResourceAttrSet(
						configInstance, "ipv6_address_id"),
					resource.TestCheckResourceAttrSet(
						configInstance, "public_ipv6_subnet"),
					resource.TestCheckResourceAttr(
						configInstance, "secondary_ip_count", "4"),
					resource.TestCheckResourceAttrSet(
						configInstance, "secondary_ip_addresses.3"),
				),
			},

			{
				Config:  testAccIBMComputeVmInstanceConfigBasic(hostname, domain, networkSpeed1, cores1, memory1, userMetadata2, tags2),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "user_metadata", userMetadata2),
					CheckStringSet(
						configInstance,
						"tags", []string{tags2},
					),
				),
			},

			{
				Config: testAccIBMComputeVmInstanceConfigBasic(hostname, domain, networkSpeed2, cores2, memory2, userMetadata2, tags2),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "cores", cores2),
					resource.TestCheckResourceAttr(
						configInstance, "memory", memory2),
					resource.TestCheckResourceAttr(
						configInstance, "network_speed", networkSpeed2),
				),
			},
		},
	})
}

func TestAccIBMComputeVmInstance_BlockDeviceTemplateGroup(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "bdtg.terraformvmuat.ibm.com"

	imageID := os.Getenv("IBM_COMPUTE_VM_INSTANCE_IMAGE_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVmInstanceConfigBlockDeviceTemplateGroup(hostname, domain, imageID),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVmInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-BDTGroup", &guest),
				),
			},
		},
	})
}

func TestAccIBMComputeVmInstance_CustomImageMultipleDisks(t *testing.T) {
	var guest datatypes.Virtual_Guest
	hostname := acctest.RandString(16)
	domain := "mdisk.terraformvmuat.ibm.com"

	imageID := os.Getenv("IBM_COMPUTE_VM_INSTANCE_IMAGE_ID")

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-disks"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVmInstanceConfigCustomImageMultipleDisks(hostname, domain, imageID),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVmInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "25"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
				),
			},
		},
	})
}

func TestAccIBMComputeVmInstance_PostInstallScriptUri(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "pis.terraformvmuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVmInstanceConfigPostInstallScriptURI(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVmInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-pISU", &guest),
				),
			},
		},
	})
}

func TestAccIBMComputeVmInstance_With_Network_Storage_Access(t *testing.T) {
	var guest datatypes.Virtual_Guest
	hostname := acctest.RandString(16)
	domain := "storage.tfmvmuat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-vsi-storage-access"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessToStoragesBasic(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists("ibm_compute_vm_instance.terraform-vsi-storage-access", &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "datacenter", "wdc04"),
					resource.TestCheckResourceAttr(
						configInstance, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configInstance, "file_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "block_storage_ids.#", "1"),
				),
			},
			{
				Config: testAccessToStoragesUpdate(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists("ibm_compute_vm_instance.terraform-vsi-storage-access", &guest),
					resource.TestCheckResourceAttr(
						configInstance, "file_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "block_storage_ids.#", "0"),
				),
			},
		},
	})
}

func testAccIBMComputeVmInstanceDestroy(s *terraform.State) error {
	service := services.GetVirtualGuestService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_vm_instance" {
			continue
		}

		guestID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the guest
		_, err := service.Id(guestID).GetObject()

		// Wait

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf(
				"Error waiting for virtual guest (%s) to be destroyed: %s",
				rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccIBMComputeVmInstanceExists(n string, guest *datatypes.Virtual_Guest) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No virtual guest ID is set")
		}

		id, err := strconv.Atoi(rs.Primary.ID)

		if err != nil {
			return err
		}

		service := services.GetVirtualGuestService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		retrieveVirtGuest, err := service.Id(id).GetObject()

		if err != nil {
			return err
		}

		fmt.Printf("The ID is %d\n", id)

		if *retrieveVirtGuest.Id != id {
			return errors.New("Virtual guest not found")
		}

		*guest = retrieveVirtGuest

		return nil
	}
}

func CheckStringSet(n string, name string, set []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		values := []string{}
		setLengthKey := fmt.Sprintf("%s.#", name)
		prefix := fmt.Sprintf("%s.", name)
		for k, v := range rs.Primary.Attributes {
			if k != setLengthKey && strings.HasPrefix(k, prefix) {
				values = append(values, v)
			}
		}

		if len(values) == 0 {
			return fmt.Errorf("Could not find %s.%s", n, name)
		}

		for _, s := range set {
			found := false
			for _, v := range values {
				if s == v {
					found = true
					break
				}
			}

			if !found {
				return fmt.Errorf("%s was not found in the set %s", s, name)
			}
		}

		return nil
	}
}

func testAccIBMComputeVmInstanceConfigBasic(hostname, domain, networkSpeed, cores, memory, userMetadata, tags string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_7_64"
    datacenter = "wdc04"
    network_speed = %s
    hourly_billing = true
    private_network_only = false
    cores = %s
    memory = %s
    disks = [25, 10, 20]
    user_metadata = "%s"
    tags = ["%s"]
    dedicated_acct_host_only = true
    local_disk = false
    ipv6_enabled = true
    secondary_ip_count = 4
}`, hostname, domain, networkSpeed, cores, memory, userMetadata, tags)
}

func testAccIBMComputeVmInstanceConfigPostInstallScriptURI(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-pISU" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_7_64"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = true
	private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
    post_install_script_uri = "https://www.google.com"
}`, hostname, domain)
}

func testAccIBMComputeVmInstanceConfigBlockDeviceTemplateGroup(hostname, domain, imageID string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-BDTGroup" {
    hostname = "%s"
    domain = "%s"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = false
    cores = 1
    memory = 1024
    local_disk = false
    image_id = %s
}`, hostname, domain, imageID)
}

func testAccIBMComputeVmInstanceConfigCustomImageMultipleDisks(hostname, domain, imageID string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-disks" {
    hostname = "%s"
    domain = "%s"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = false
    cores = 1
    memory = 1024
    local_disk = false
    image_id = %s
    disks = [25, 10]
}`, hostname, domain, imageID)
}

const fsConfig1 = `
resource "ibm_storage_file" "fs1" {
  type              = "Endurance"
  datacenter        = "wdc04"
  capacity          = 20
  iops              = 0.25
  snapshot_capacity = 10
}
`

const bsConfig1 = `resource "ibm_storage_block" "bs" {
  type              = "Endurance"
  datacenter        = "wdc04"
  capacity          = 20
  iops              = 0.25
  snapshot_capacity = 10
  os_format_type    = "Linux"
}
`

const fsConfig2 = `resource "ibm_storage_file" "fs2" {
  type              = "Endurance"
  datacenter        = "wdc04"
  capacity          = 20
  iops              = 0.25
  snapshot_capacity = 10
}

`

func testAccessToStoragesBasic(hostname, domain string) string {
	config := fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-vsi-storage-access" {
    hostname = "%s"
    domain = "%s"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = true
	file_storage_ids = ["${ibm_storage_file.fs1.id}"]
	block_storage_ids = ["${ibm_storage_block.bs.id}"]

    cores = 1
    memory = 1024
    local_disk = false
    os_reference_code = "DEBIAN_7_64"
    disks = [25, 10]
}
%s
%s

`, hostname, domain, fsConfig1, bsConfig1)
	return config
}

func testAccessToStoragesUpdate(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-vsi-storage-access" {
    hostname = "%s"
    domain = "%s"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = true
	file_storage_ids = ["${ibm_storage_file.fs2.id}"]
	block_storage_ids = []
    cores = 1
    memory = 1024
    local_disk = false
    os_reference_code = "DEBIAN_7_64"
    disks = [25, 10]
}

%s

`, hostname, domain, fsConfig2)

}
