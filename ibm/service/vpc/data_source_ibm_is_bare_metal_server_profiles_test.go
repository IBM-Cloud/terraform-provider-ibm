// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBMSProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.id"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.disks.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.href"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.supported_trusted_platform_module_modes.#"),
				),
			},
		},
	})
}
func TestAccIBMISBMSProfilesDataSource_ResourceTypeNull(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.id"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.disks.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.href"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.supported_trusted_platform_module_modes.#"),
				),
			},
		},
	})
}
func TestAccIBMISBMSProfilesDataSource_vni(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.id"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.0.default"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.bandwidth.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.disks.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.href"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_attachment_count.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_attachment_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_attachment_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.virtual_network_interfaces_supported.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.virtual_network_interfaces_supported.0.type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.virtual_network_interfaces_supported.0.value"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.supported_trusted_platform_module_modes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_bare_metal_server_profiles" "test1" {
      }`)
}
