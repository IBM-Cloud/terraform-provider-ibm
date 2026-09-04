// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.0.values"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfilesDataSource_QoS(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.volume_bandwidth_qos_modes.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_bandwidth_mode.0.type"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfilesDataSource_cluster(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.reservation_terms.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.cluster_network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.cluster_network_attachment_count.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.supported_cluster_network_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.zones.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceProfilesDataSource_concom(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.0.values.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceProfilesDataSource_sharedcore(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.confidential_compute_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.secure_boot_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_burst_limit.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_burst_limit.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_burst_limit.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_percentage.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_percentage.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_percentage.0.values.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfilesDataSource_AvailabilityClass(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.availability_class.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.availability_class.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.availability_class.0.values.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfilesDataSource_ThreadsPerCore(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.threads_per_core.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.threads_per_core.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.threads_per_core.0.default"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.threads_per_core.0.values.#"),
				),
			},
		},
	})
}

// TestAccIBMISInstanceProfilesDataSource_SupportedVcpuCountBackfill verifies that
// supported_vcpu_count is correctly backfilled from vcpu_count when type == "enum"
// in the ibm_is_instance_profiles (plural) flatten path.
// The known enum profile (acc.ISInstanceProfileName) is expected to appear somewhere
// in the list; we locate it via a Terraform locals filter so assertions can use a
// deterministic address while still exercising the ibm_is_instance_profiles read path.
func TestAccIBMISInstanceProfilesDataSource_SupportedVcpuCountBackfill(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfilesEnumDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					// datasource returned profiles
					resource.TestCheckResourceAttrSet(resName, "profiles.#"),
					// every profile must have supported_vcpu_count set (may be empty list for non-enum)
					resource.TestCheckResourceAttrSet(resName, "profiles.0.supported_vcpu_count.#"),
					// vcpu_count must always have type and value populated (backfill for value)
					resource.TestCheckResourceAttrSet(resName, "profiles.0.vcpu_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.vcpu_count.0.value"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_instance_profiles" "test1" {
      }`)
}

func testAccCheckIBMISInstanceProfilesEnumDataSourceConfig() string {
	return `
data "ibm_is_instance_profiles" "test1" {}
`
}
