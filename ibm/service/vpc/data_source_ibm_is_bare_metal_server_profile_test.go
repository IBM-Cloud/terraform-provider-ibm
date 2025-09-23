// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBMSProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISBMSProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "disks.#"),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "href"),
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "memory.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "resource_type"),
					resource.TestCheckResourceAttrSet(resName, "supported_trusted_platform_module_modes.#"),
					resource.TestCheckResourceAttrSet(resName, "reservation_terms.#"),
					resource.TestCheckResourceAttrSet(resName, "reservation_terms.0.type"),
					resource.TestCheckResourceAttrSet(resName, "reservation_terms.0.values"),
				),
			},
		},
	})
}
func TestAccIBMISBMSProfileDataSource_EmptyResourceType(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISBMSProfileDataSourceResourceTypeConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "disks.#"),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "href"),
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "memory.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "supported_trusted_platform_module_modes.#"),
				),
			},
		},
	})
}
func TestAccIBMISBMSProfileDataSource_vni(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISBMSProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth.0.default"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.#"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.type"),
					resource.TestCheckResourceAttrSet(resName, "console_types.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_core_count.#"),
					resource.TestCheckResourceAttrSet(resName, "cpu_socket_count.#"),
					resource.TestCheckResourceAttrSet(resName, "disks.#"),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "href"),
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "memory.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.#"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "network_interface_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "network_attachment_count.#"),
					resource.TestCheckResourceAttrSet(resName, "network_attachment_count.0.max"),
					resource.TestCheckResourceAttrSet(resName, "network_attachment_count.0.min"),
					resource.TestCheckResourceAttrSet(resName, "network_attachment_count.0.type"),
					resource.TestCheckResourceAttrSet(resName, "virtual_network_interfaces_supported.#"),
					resource.TestCheckResourceAttrSet(resName, "virtual_network_interfaces_supported.0.type"),
					resource.TestCheckResourceAttrSet(resName, "virtual_network_interfaces_supported.0.value"),
					resource.TestCheckResourceAttrSet(resName, "os_architecture.#"),
					resource.TestCheckResourceAttrSet(resName, "resource_type"),
					resource.TestCheckResourceAttrSet(resName, "supported_trusted_platform_module_modes.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSProfileDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_profiles" "testbmsps" {
		}

		data "ibm_is_bare_metal_server_profile" "test1" {
			name = data.ibm_is_bare_metal_server_profiles.testbmsps.profiles.0.name
		}`)
}
func testAccCheckIBMISBMSProfileDataSourceResourceTypeConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_profile" "test1" {
			name = "cx2d-metal-96x192"
		}`)
}
