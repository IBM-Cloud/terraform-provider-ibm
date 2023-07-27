// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "network_interface_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_network_interface_floating_ips.is_floating_ips.0", "zone.#"),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceAllArgs(t *testing.T) {
	floatingIPName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfig(floatingIPName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "network_interface_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_floating_ip.is_floating_ip", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_is_floating_ip" "is_floating_ip_instance" {
		}

		data "ibm_is_floating_ip" "is_floating_ip_instance" {
			instance_id = "instance_id"
			network_interface_id = "network_interface_id"
			id = "id"
		}
	`)
}

func testAccCheckIBMIsVirtualNetworkInterfaceFloatingIPsDataSourceConfig(floatingIPName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_floating_ip" "is_floating_ip_instance" {
			name = "%s"
			resource_group {
				href = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
				id = "fee82deba12e4c0fb69c3b09d1f12345"
				name = "my-resource-group"
			}
			target {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/instances/1e09281b-f177-46fb-baf1-bc152b2e391a/network_interfaces/10c02d81-0ecb-4dc5-897d-28392913b81e"
				id = "10c02d81-0ecb-4dc5-897d-28392913b81e"
				name = "my-instance-network-interface"
				primary_ip {
					address = "192.168.3.4"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
					name = "my-reserved-ip"
					resource_type = "subnet_reserved_ip"
				}
				resource_type = "network_interface"
			}
			zone {
				href = "https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1"
				name = "us-south-1"
			}
		}

		data "ibm_is_floating_ip" "is_floating_ip_instance" {
			instance_id = "instance_id"
			network_interface_id = "network_interface_id"
			id = "id"
		}
	`, floatingIPName)
}
