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

func testAccCheckIBMISInstanceProfileDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_instance_profile" "test1" {
	name = "%s"
}`, acc.InstanceProfileName)
}
