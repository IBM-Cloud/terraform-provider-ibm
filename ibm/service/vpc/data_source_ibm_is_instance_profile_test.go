// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "reservation_terms.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "reservation_terms.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "reservation_terms.0.values"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfileDataSource_QoSMode(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "volume_bandwidth_qos_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "zones.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfileDataSource_cluster(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "cluster_network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "cluster_network_attachment_count.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "disks.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "reservation_terms.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "supported_cluster_network_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "supported_cluster_network_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "supported_cluster_network_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "supported_cluster_network_profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "total_volume_bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceProfileDataSource_concom(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.0.values.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceProfileDataSource_sharedcore(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "confidential_compute_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "secure_boot_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_burst_limit.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_burst_limit.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_burst_limit.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_percentage.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_percentage.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_percentage.0.values.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfileDataSourceNetworkBandwidth(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_manufacturer.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_bandwidth_mode.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "network_bandwidth_mode.0.value"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfileDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_instance_profile" "test1" {
	name = "%s"
}

data "ibm_is_instance_profiles" "test1" {

}

`, acc.InstanceProfileName)
}

func testAccCheckIBMISInstanceProfileEnumDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_is_instance_profile" "enum_profile" {
	name = "%s"
}
`, acc.ISInstanceProfileName)
}

func TestAccIBMISInstanceProfileDataSource_AvailabilityClass(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "availability_class.#"),
					resource.TestCheckResourceAttrSet(resName, "availability_class.0.type"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceProfileDataSource_ThreadsPerCore(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.InstanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "threads_per_core.#"),
					resource.TestCheckResourceAttrSet(resName, "threads_per_core.0.type"),
					resource.TestCheckResourceAttrSet(resName, "threads_per_core.0.default"),
					resource.TestCheckResourceAttrSet(resName, "threads_per_core.0.values.#"),
				),
			},
		},
	})
}

// TestAccIBMISInstanceProfileDataSource_SupportedVcpuCountBackfill verifies that
// supported_vcpu_count is correctly backfilled from vcpu_count when type == "enum".
// It requires a profile whose vcpu_count.type is "enum" (e.g. a GPU profile).
func TestAccIBMISInstanceProfileDataSource_SupportedVcpuCountBackfill(t *testing.T) {
	resName := "data.ibm_is_instance_profile.enum_profile"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceProfileEnumDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.ISInstanceProfileName),
					// vcpu_count sub-attributes populated by new SDK response shape
					resource.TestCheckResourceAttr(resName, "vcpu_count.0.type", "enum"),
					resource.TestCheckResourceAttrSet(resName, "vcpu_count.0.default"),
					resource.TestCheckResourceAttrSet(resName, "vcpu_count.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "vcpu_count.0.values.0"),
					// vcpu_count.value is backfilled from vcpu_count.default
					resource.TestCheckResourceAttrSet(resName, "vcpu_count.0.value"),
					resource.TestCheckResourceAttrPair(resName, "vcpu_count.0.value", resName, "vcpu_count.0.default"),
					// supported_vcpu_count is backfilled from vcpu_count — all sub-attributes present
					resource.TestCheckResourceAttr(resName, "supported_vcpu_count.#", "1"),
					resource.TestCheckResourceAttr(resName, "supported_vcpu_count.0.type", "enum"),
					resource.TestCheckResourceAttrSet(resName, "supported_vcpu_count.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "supported_vcpu_count.0.values.0"),
					// backfilled values are identical to the source vcpu_count values
					resource.TestCheckResourceAttrPair(resName, "supported_vcpu_count.0.type", resName, "vcpu_count.0.type"),
					resource.TestCheckResourceAttrPair(resName, "supported_vcpu_count.0.values.#", resName, "vcpu_count.0.values.#"),
					resource.TestCheckResourceAttrPair(resName, "supported_vcpu_count.0.values.0", resName, "vcpu_count.0.values.0"),
				),
			},
		},
	})
}
