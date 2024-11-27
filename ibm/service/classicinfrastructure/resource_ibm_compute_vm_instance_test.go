// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeVMInstance_basic(t *testing.T) {
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

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccIBMComputeVMInstanceConfigBasic(hostname, domain, networkSpeed1, cores1, memory1, userMetadata1, tags1),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
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
					resource.TestCheckResourceAttr(
						configInstance, "notes", "VM notes"),
				),
			},

			{
				Config:  testAccIBMComputeVMInstanceConfigBasic(hostname, domain, networkSpeed1, cores1, memory1, userMetadata1, tags2),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "user_metadata", userMetadata1Unquoted),
					CheckStringSet(
						configInstance,
						"tags", []string{tags2},
					),
				),
			},

			{
				Config: testAccIBMComputeVMInstanceConfigBasic(hostname, domain, networkSpeed2, cores2, memory2, userMetadata1, tags2),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "cores", cores2),
					resource.TestCheckResourceAttr(
						configInstance, "memory", memory2),
					resource.TestCheckResourceAttr(
						configInstance, "network_speed", networkSpeed2),
				),
			},

			{
				Config:  testAccIBMComputeVMInstanceConfigUpdate(hostname, domain, networkSpeed2, cores2, memory2, userMetadata1, tags2),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "25"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.2", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.3", "20"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_bulkvms(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname1 := acctest.RandString(16)
	hostname2 := acctest.RandString(16)
	domain := "terraformvmuat.ibm.com"
	networkSpeed1 := "10"
	cores1 := "1"
	memory1 := "1024"
	tags1 := "collectd"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	userMetadata1Unquoted, _ := strconv.Unquote(`"` + userMetadata1 + `"`)

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccIBMComputeVMInstanceConfigBulkVMs(hostname1, hostname2, domain, networkSpeed1, cores1, memory1, userMetadata1, tags1),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
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
					resource.TestCheckResourceAttr(
						configInstance, "notes", "VM notes"),
					resource.TestCheckResourceAttr(
						configInstance, "bulk_vms.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstanceWithFlavor(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "terraformvmuat.ibm.com"
	networkSpeed1 := "10"
	cores1 := "1"
	memory1 := "2048"
	tags1 := "collectd"
	flavor := "B1_1X2X25"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	userMetadata1Unquoted, _ := strconv.Unquote(`"` + userMetadata1 + `"`)
	updatedFlavor := "B1_4X8X25"
	networkSpeed2 := "100"
	cores2 := "4"
	memory2 := "8192"

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigFlavor(hostname, domain, networkSpeed1, flavor, userMetadata1, tags1),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
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
						configInstance, "flavor_key_name", flavor),
					resource.TestCheckResourceAttr(
						configInstance, "cores", cores1),
					resource.TestCheckResourceAttr(
						configInstance, "memory", memory1),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "20"),
					resource.TestCheckResourceAttr(
						configInstance, "user_metadata", userMetadata1Unquoted),
					resource.TestCheckResourceAttr(
						configInstance, "local_disk", "false"),
					resource.TestCheckResourceAttr(
						configInstance, "dedicated_acct_host_only", "false"),
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
					resource.TestCheckResourceAttr(
						configInstance, "notes", "VM notes"),
				),
			},
			{
				Config: testAccIBMComputeVMInstanceConfigFlavorUpdate(hostname, domain, networkSpeed2, updatedFlavor, userMetadata1, tags1),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "20"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.2", "20"),
					resource.TestCheckResourceAttr(
						configInstance, "flavor_key_name", updatedFlavor),
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

func TestAccIBMComputeVMInstance_With_SSH_Keys(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfsshkeyvmuat.ibm.com"
	label := fmt.Sprintf("terraformsshuat_create_step_label_%d", acctest.RandIntRange(10, 100))
	notes := fmt.Sprintf("terraformsshuat_update_step_notes_%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	configInstance := "ibm_compute_vm_instance.terraform-ssh-key"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithSSHKey(label, notes, publicKey, hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "ssh_key_ids.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_basic_import(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "tfsshkeyvmuat.ibm.com"
	label := fmt.Sprintf("terraformsshuat_create_step_label_%d", acctest.RandIntRange(10, 100))
	notes := fmt.Sprintf("terraformsshuat_update_step_notes_%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resourceName := "ibm_compute_vm_instance.terraform-ssh-key"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testComputeInstanceWithSSHKey(label, notes, publicKey, hostname, domain),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes",
					"public_bandwidth_unlimited",
				},
			},
		},
	})
}

func TestAccIBMComputeVMInstance_basic_import_WithFlavor(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "terraformuat.ibm.com"
	tags1 := "collectd"
	flavor := "B1_1X2X25"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	networkSpeed1 := "10"

	resourceName := "ibm_compute_vm_instance.terraform-acceptance-test-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigFlavor(hostname, domain, networkSpeed1, flavor, userMetadata1, tags1),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes",
					"public_bandwidth_unlimited",
				},
			},
		},
	})
}

func TestAccIBMComputeVMInstance_InvalidNotes(t *testing.T) {
	hostname := acctest.RandString(16)
	domain := "terraformvmuat.ibm.com"
	networkSpeed1 := "10"
	cores1 := "1"
	memory1 := "1024"
	tags1 := "collectd"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMComputeVMInstanceInvalidNotes(hostname, domain, networkSpeed1, cores1, memory1, userMetadata1, tags1),
				ExpectError: regexp.MustCompile("should not exceed 1000 characters"),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_BlockDeviceTemplateGroup(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "bdtg.terraformvmuat.ibm.com"
	flavor := "B1_1X2X25"
	networkSpeed := "10"
	// Image Id of RightImage_Ubuntu_10.04_x64_v5.7.24
	imageID := 15789
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigBlockDeviceTemplateGroup(hostname, domain, networkSpeed, flavor, imageID),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVMInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-BDTGroup", &guest),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_CustomImageMultipleDisks(t *testing.T) {
	var guest datatypes.Virtual_Guest
	hostname := acctest.RandString(16)
	domain := "mdisk.terraformvmuat.ibm.com"
	// Image Id of RightImage_Ubuntu_10.04_x64_v5.7.24
	imageID := 15789

	configInstance := "ibm_compute_vm_instance.terraform-acceptance-test-disks"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigCustomImageMultipleDisks(hostname, domain, imageID),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
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

func TestAccIBMComputeVMInstance_PostInstallScriptUri(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "pis.terraformvmuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigPostInstallScriptURI(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVMInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-pISU", &guest),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_WINDOWS_PostInstallScriptUri(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(14)
	domain := "terraformuat.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceConfigWindowsPostInstallScriptURI(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					// image_id value is hardcoded. If it's valid then virtual guest will be created well
					testAccIBMComputeVMInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-pISU", &guest),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Network_Storage_Access(t *testing.T) {
	var guest datatypes.Virtual_Guest
	hostname := acctest.RandString(16)
	domain := "storage.tfmvmuat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-vsi-storage-access"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccessToStoragesBasic(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists("ibm_compute_vm_instance.terraform-vsi-storage-access", &guest),
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
					testAccIBMComputeVMInstanceExists("ibm_compute_vm_instance.terraform-vsi-storage-access", &guest),
					resource.TestCheckResourceAttr(
						configInstance, "file_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "block_storage_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Public_Bandwidth_Limited(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmbandwidthuat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-public-bandwidth"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithPublicBandWidth(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "public_bandwidth_limited", "1000"),
				),
			},
			{
				Config:  testComputeInstanceWithPublicBandWidthDefault(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
				),
			},
		},
	})
}
func TestAccIBMComputeVMInstance_With_Public_Bandwidth_Unlimited(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmbandwidthuat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-public-bandwidth"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithPublicBandwidthUnlimited(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_DedicatedHost_Name(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmdedicateduat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-vm-dedicatedhost"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithDedicatdHostName(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_DedicatedHost_ID(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmdedicateduat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-vm-dedicatedhost"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithDedicatdHostID(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Security_Groups(t *testing.T) {
	var guest datatypes.Virtual_Guest
	var pubsg datatypes.Network_SecurityGroup
	var pvtsg datatypes.Network_SecurityGroup
	sgName1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	sgDesc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))
	sgName2 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandIntRange(10, 100))
	sgDesc2 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandIntRange(10, 100))

	hostname := acctest.RandString(16)

	configInstance := "ibm_compute_vm_instance.tfuatvmwithgroups"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccIBMComputeVMInstanceConfigWithSecurityGroups(sgName1, sgDesc1, sgName2, sgDesc2, hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMSecurityGroupExists("ibm_security_group.pubsg", &pubsg),
					testAccCheckIBMSecurityGroupExists("ibm_security_group.pvtsg", &pvtsg),
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "public_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "private_security_group_ids.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Evault(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmevaultuat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-evault"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithEvault(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "evault", "20"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Retry(t *testing.T) {
	var guest datatypes.Virtual_Guest

	hostname := acctest.RandString(16)
	domain := "tfvmretry.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-retry"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithRetry(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "datacenter", "dal06"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Placement_group(t *testing.T) {
	var guest datatypes.Virtual_Guest
	placementGroup := "tf-placement-group" + acctest.RandString(16)
	hostname := acctest.RandString(16)
	domain := "tfvmpguat.ibm.com"

	configInstance := "ibm_compute_vm_instance.terraform-pgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testComputeInstanceWithPlacementGroup(hostname, domain, placementGroup),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "datacenter", "dal05"),
				),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_With_Invalid_Retry(t *testing.T) {

	hostname := acctest.RandString(16)
	domain := "tfvmretry.ibm.com"
	var errMsg = "\"test\" Invalid values are provided in `datacenter_choice`"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testComputeInstanceWithRetryInvalid(hostname, domain),
				ExpectError: regexp.MustCompile(errMsg),
			},
		},
	})
}

func TestAccIBMComputeVMInstance_Transient(t *testing.T) {
	var guest datatypes.Virtual_Guest
	hostname := acctest.RandString(16)
	domain := "terraformuat.ibm.com"
	tags1 := "collectd"
	flavor := "B1_1X2X25"
	userMetadata1 := "{\\\"value\\\":\\\"newvalue\\\"}"
	networkSpeed1 := "10"
	cores1 := "1"
	memory1 := "2048"
	userMetadata1Unquoted, _ := strconv.Unquote(`"` + userMetadata1 + `"`)

	configInstance := "ibm_compute_vm_instance.terraform-transient"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMComputeVMInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMComputeVMInstanceTransient(hostname, domain, networkSpeed1, flavor, userMetadata1, tags1),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVMInstanceExists(configInstance, &guest),
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
						configInstance, "flavor_key_name", flavor),
					resource.TestCheckResourceAttr(
						configInstance, "cores", cores1),
					resource.TestCheckResourceAttr(
						configInstance, "memory", memory1),
					resource.TestCheckResourceAttr(
						configInstance, "disks.0", "10"),
					resource.TestCheckResourceAttr(
						configInstance, "disks.1", "20"),
					resource.TestCheckResourceAttr(
						configInstance, "user_metadata", userMetadata1Unquoted),
					resource.TestCheckResourceAttr(
						configInstance, "local_disk", "false"),
					resource.TestCheckResourceAttr(
						configInstance, "dedicated_acct_host_only", "false"),
					resource.TestCheckResourceAttr(
						configInstance, "transient", "true"),
					CheckStringSet(
						configInstance,
						"tags", []string{tags1},
					),
				),
			},
			{
				ResourceName:      configInstance,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes",
					"public_bandwidth_unlimited",
				},
			},
		},
	})
}

func testAccIBMComputeVMInstanceDestroy(s *terraform.State) error {
	service := services.GetVirtualGuestService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_vm_instance" {
			continue
		}
		parts, err := flex.VmIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		for _, str := range parts {
			guestID, _ := strconv.Atoi(str)

			// Try to find the guest
			_, err := service.Id(guestID).GetObject()

			// Wait

			if err != nil && !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error waiting for virtual guest (%s) to be destroyed: %s",
					rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccIBMComputeVMInstanceExists(n string, guest *datatypes.Virtual_Guest) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No virtual guest ID is set")
		}
		parts, err := flex.VmIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		for _, str := range parts {
			id, err := strconv.Atoi(str)

			if err != nil {
				return err
			}

			service := services.GetVirtualGuestService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
			retrieveVirtGuest, err := service.Id(id).GetObject()

			if err != nil {
				return err
			}

			fmt.Printf("The ID is %d\n", id)

			if *retrieveVirtGuest.Id != id {
				return errors.New("Virtual guest not found")
			}

			*guest = retrieveVirtGuest
		}

		return nil
	}
}

func CheckStringSet(n string, name string, set []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
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
			return fmt.Errorf("[ERROR] Could  not find %s.%s", n, name)
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

func testAccIBMComputeVMInstanceConfigBasic(hostname, domain, networkSpeed, cores, memory, userMetadata, tags string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_9_64"
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
    notes = "VM notes"
}`, hostname, domain, networkSpeed, cores, memory, userMetadata, tags)
}

func testAccIBMComputeVMInstanceConfigBulkVMs(hostname1, hostname2, domain, networkSpeed, cores, memory, userMetadata, tags string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
	 bulk_vms  {
	    hostname = "%s"
	
	    domain = "%s"
	  }
	  bulk_vms {
	    hostname = "%s"
	
	    domain = "%s"
	  }

    os_reference_code = "DEBIAN_9_64"
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
    notes = "VM notes"
}`, hostname1, domain, hostname2, domain, networkSpeed, cores, memory, userMetadata, tags)
}

func testAccIBMComputeVMInstanceConfigUpdate(hostname, domain, networkSpeed, cores, memory, userMetadata, tags string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_9_64"
    datacenter = "wdc04"
    network_speed = %s
    hourly_billing = true
    private_network_only = false
    cores = %s
    memory = %s
    disks = [25, 10, 10, 20]
    user_metadata = "%s"
    tags = ["%s"]
    dedicated_acct_host_only = true
    local_disk = false
    ipv6_enabled = true
    secondary_ip_count = 4
    notes = "VM notes"
}`, hostname, domain, networkSpeed, cores, memory, userMetadata, tags)
}

func testAccCheckIBMComputeVMInstanceInvalidNotes(hostname, domain, networkSpeed, cores, memory, userMetadata, tags string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_9_64"
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
    notes = "This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very longThis notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long This notes is very long"
}`, hostname, domain, networkSpeed, cores, memory, userMetadata, tags)
}

func testAccIBMComputeVMInstanceConfigPostInstallScriptURI(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-pISU" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "DEBIAN_9_64"
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

func testAccIBMComputeVMInstanceConfigWindowsPostInstallScriptURI(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-acceptance-test-pISU" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "WIN_2016-STD_64"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = true
	private_network_only = false
    cores = 1
    memory = 2048
    disks = [100]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
    post_install_script_uri = "https://www.google.com"
}`, hostname, domain)
}

func testAccIBMComputeVMInstanceConfigBlockDeviceTemplateGroup(hostname, domain, networkSpeed, flavor string, imageID int) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-transient" {
	hostname = "%s"
	domain = "%s"
	image_id = "%d"
	datacenter = "wdc04"
	network_speed = "%s"
	hourly_billing = true
	private_network_only = false
	flavor_key_name = "%s"
	disks = [10 ,20]
	local_disk = false
}
`, hostname, domain, imageID, networkSpeed, flavor)
}

func testAccIBMComputeVMInstanceConfigCustomImageMultipleDisks(hostname, domain string, imageID int) string {
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
    image_id = "%d"
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
    os_reference_code = "DEBIAN_9_64"
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
    os_reference_code = "DEBIAN_9_64"
    disks = [25, 10]
}

%s

`, hostname, domain, fsConfig2)

}

func testComputeInstanceWithSSHKey(sshLabel, sshNotes, sshPublicKey, hostname, domain string) (config string) {
	config = testAccCheckIBMComputeSSHKeyConfig(sshLabel, sshNotes, sshPublicKey) + fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-ssh-key" {
    hostname = "%s"
    domain = "%s"
    datacenter = "wdc04"
    network_speed = 10
    hourly_billing = true
    ssh_key_ids = ["${ibm_compute_ssh_key.testacc_ssh_key.id}"]
    cores = 1
    memory = 1024
    local_disk = false
    os_reference_code = "DEBIAN_9_64"
    disks = [25]
}
`, hostname, domain)
	return
}

func testComputeInstanceWithPublicBandWidthDefault(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-public-bandwidth" {
	hostname = "%s"
	domain = "%s"
	datacenter = "wdc04"
	network_speed = 10
	hourly_billing = false
	cores = 1
	memory = 1024
	local_disk = false
	os_reference_code = "DEBIAN_9_64"
	disks = [25]
}
`, hostname, domain)
}

func testComputeInstanceWithPublicBandWidth(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-public-bandwidth" {
	hostname = "%s"
	domain = "%s"
	datacenter = "wdc04"
	network_speed = 10
	hourly_billing = false
	cores = 1
	memory = 1024
	local_disk = false
	os_reference_code = "DEBIAN_9_64"
	disks = [25]
	public_bandwidth_limited = 1000
}
`, hostname, domain)
}

func testComputeInstanceWithPublicBandwidthUnlimited(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-public-bandwidth" {
	hostname = "%s"
	domain = "%s"
	datacenter = "wdc04"
	network_speed = 100
	hourly_billing = false
	cores = 1
	memory = 1024
	local_disk = false
	os_reference_code = "DEBIAN_9_64"
	disks = [25]
	public_bandwidth_unlimited = true
}
`, hostname, domain)
}

func testComputeInstanceWithDedicatdHostName(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-vm-dedicatedhost" {
	hostname = "%s"
	domain = "%s"
	hourly_billing = true
	datacenter = "dal10"
	network_speed = 100
	cores = 1
	memory = 1024
	os_reference_code = "DEBIAN_9_64"
	disks                = [25, 25, 100]
	dedicated_host_name  = "%s"
}
`, hostname, domain, acc.DedicatedHostName)
}

func testComputeInstanceWithDedicatdHostID(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-vm-dedicatedhost" {
	hostname = "%s"
	domain = "%s"
	hourly_billing = true
	datacenter = "dal10"
	network_speed = 100
	cores = 1
	memory = 1024
	os_reference_code = "DEBIAN_9_64"
	disks                = [25, 100, 25]
	dedicated_host_id  = "%s"
}
`, hostname, domain, acc.DedicatedHostID)
}

func testAccIBMComputeVMInstanceConfigWithSecurityGroups(sgName1, sgDesc1, sgName2, sgDesc2, hostname string) string {
	v := fmt.Sprintf(`
		resource "ibm_security_group" "pubsg" {
			name        = "%s"
			description = "%s"
		  } 
		  resource "ibm_security_group_rule" "pubsgrule" {
			direction         = "ingress"
			port_range_min    = 80
			port_range_max    = 8080
			protocol          = "udp"
			security_group_id = "${ibm_security_group.pubsg.id}"
		  }
		  resource "ibm_security_group" "pvtsg" {
			name        = "%s"
			description = "%s"
		  }
		  resource "ibm_security_group_rule" "pvtsgrule" {
			direction         = "ingress"
			port_range_min    = 80
			port_range_max    = 8085
			protocol          = "tcp"
			security_group_id = "${ibm_security_group.pvtsg.id}"
		  }
		  resource "ibm_compute_vm_instance" "tfuatvmwithgroups" {
			hostname                   = "%s"
			domain                     = "tfvmuatsg.com"
			os_reference_code          = "DEBIAN_9_64"
			datacenter                 = "wdc07"
			network_speed              = 10
			hourly_billing             = true
			private_network_only       = false
			cores                      = 1
			memory                     = 1024
			disks                      = [25, 10, 20]
			dedicated_acct_host_only   = true
			local_disk                 = false
			ipv6_enabled               = true
			secondary_ip_count         = 4
			notes                      = "VM notes"
			public_security_group_ids  = ["${ibm_security_group.pubsg.id}"]
			private_security_group_ids = ["${ibm_security_group.pvtsg.id}"]
		  }`, sgName1, sgDesc1, sgName2, sgDesc2, hostname)
	return v
}

func testAccIBMComputeVMInstanceConfigFlavor(hostname, domain, networkSpeed, flavor, userMetadata, tags string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
	    hostname = "%s"
	    domain = "%s"
	    os_reference_code = "DEBIAN_9_64"
	    datacenter = "wdc04"
	    network_speed = %s
	    hourly_billing = true
	    private_network_only = false
	    flavor_key_name = "%s"
	    user_metadata = "%s"
		tags = ["%s"]
		disks = [10 ,20]
	    local_disk = false
	    ipv6_enabled = true
	    secondary_ip_count = 4
	    notes = "VM notes"
	}`, hostname, domain, networkSpeed, flavor, userMetadata, tags)
}

func testAccIBMComputeVMInstanceConfigFlavorUpdate(hostname, domain, networkSpeed, flavor, userMetadata, tags string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
	    hostname = "%s"
	    domain = "%s"
	    os_reference_code = "DEBIAN_9_64"
	    datacenter = "wdc04"
	    network_speed = %s
	    hourly_billing = true
	    private_network_only = false
	    flavor_key_name = "%s"
	    user_metadata = "%s"
		tags = ["%s"]
		disks = [10 ,20, 20]
	    local_disk = false
	    ipv6_enabled = true
	    secondary_ip_count = 4
	    notes = "VM notes"
	}`, hostname, domain, networkSpeed, flavor, userMetadata, tags)
}

func testComputeInstanceWithEvault(hostname, domain string) (config string) {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "terraform-evault" {
	hostname = "%s"
	domain = "%s"
	datacenter = "syd01"
	network_speed = 10
	hourly_billing = false
	cores = 1
	memory = 1024
	local_disk = false
	os_reference_code = "DEBIAN_9_64"
	disks = [25]
	evault = 20
}
`, hostname, domain)
}

func testComputeInstanceWithRetry(hostname, domain string) (config string) {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "terraform-retry" {
		hostname          = "%s"
		domain            = "%s"
		network_speed     = 100
		hourly_billing    = true
		cores             = 1
		memory            = 1024
		local_disk        = false
		os_reference_code = "DEBIAN_9_64"
		disks             = [25]
	  
		datacenter_choice = [
		  {
			datacenter      = "dal09"
			public_vlan_id  = 123245
			private_vlan_id = 123255
		  },
		  {
			datacenter = "wdc54"
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 153345
			private_vlan_id = 123255
		  },
		  {
			datacenter = "dal06"
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 123245
			private_vlan_id = 123255
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 1232454
			private_vlan_id = 1234567
		  },
		]
	  }		
`, hostname, domain)
}

func testComputeInstanceWithRetryInvalid(hostname, domain string) (config string) {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "terraform-retry" {
		hostname          = "%s"
		domain            = "%s"
		network_speed     = 100
		hourly_billing    = true
		cores             = 1
		memory            = 1024
		local_disk        = false
		os_reference_code = "DEBIAN_7_64"
		disks             = [25]
	  
		datacenter_choice = [
		  {
			datacenter      = "dal09"
			public_vlan_id  = 123245
			private_vlan_id = 123255
		  },
		  {
			datacenter = "wdc54"
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 153345
			private_vlan_id = 123255
			test = "key"
		  },
		  {
			datacenter = "dal06"
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 123245
			private_vlan_id = 123255
		  },
		  {
			datacenter      = "dal09"
			public_vlan_id  = 1232454
			private_vlan_id = 1234567
		  },
		]
	  }		
`, hostname, domain)
}

func testComputeInstanceWithPlacementGroup(hostname, domain, placementGroup string) (config string) {
	return fmt.Sprintf(`
	resource "ibm_compute_placement_group" "placementGroup" {
		name = "%s"
		datacenter = "dal05"
		pod = "pod01"
	}
resource "ibm_compute_vm_instance" "terraform-pgroup" {
	hostname = "%s"
	domain = "%s"
	network_speed = 10
	hourly_billing = true
	datacenter = "dal05"
	cores = 1
	memory = 1024
	local_disk = false
	os_reference_code = "DEBIAN_9_64"
	disks = [25]
	placement_group_name = ibm_compute_placement_group.placementGroup.name
}
`, placementGroup, hostname, domain)
}

func testAccIBMComputeVMInstanceTransient(hostname, domain, networkSpeed, flavor, userMetadata, tags string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_vm_instance" "terraform-transient" {
	    hostname = "%s"
	    domain = "%s"
	    os_reference_code = "DEBIAN_9_64"
	    datacenter = "wdc04"
	    network_speed = %s
	    hourly_billing = true
	    private_network_only = false
	    flavor_key_name = "%s"
	    user_metadata = "%s"
		tags = ["%s"]
		disks = [10 ,20]
		local_disk = false
		transient = true
	}`, hostname, domain, networkSpeed, flavor, userMetadata, tags)
}
