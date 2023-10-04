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

func TestAccIBMIsVirtualNetworkInterfaceIpDataSourceBasic(t *testing.T) {
	virtualNetworkInterfaceIpSubnetID := fmt.Sprintf("tf_vni_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIpDataSourceConfigBasic(virtualNetworkInterfaceIpSubnetID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "reserved_ip_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "owner"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "resource_type"),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceIpDataSourceAllArgs(t *testing.T) {
	virtualNetworkInterfaceIpSubnetID := fmt.Sprintf("tf_subnet_id_%d", acctest.RandIntRange(10, 100))
	virtualNetworkInterfaceIpAddress := fmt.Sprintf("tf_address_%d", acctest.RandIntRange(10, 100))
	virtualNetworkInterfaceIpAutoDelete := "false"
	virtualNetworkInterfaceIpName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIpDataSourceConfig(virtualNetworkInterfaceIpSubnetID, virtualNetworkInterfaceIpAddress, virtualNetworkInterfaceIpAutoDelete, virtualNetworkInterfaceIpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "reserved_ip_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "owner"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ip.is_reserved_ip", "target.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceIpDataSourceConfigBasic(virtualNetworkInterfaceIpSubnetID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_reserved_ip" "is_reserved_ip_instance" {
			subnet_id = "%s"
		}

		data "ibm_is_reserved_ip" "is_reserved_ip_instance" {
			subnet_id = ibm_is_reserved_ip.is_reserved_ip.subnet_id
			id = "id"
		}
	`, virtualNetworkInterfaceIpSubnetID)
}

func testAccCheckIBMIsVirtualNetworkInterfaceIpDataSourceConfig(virtualNetworkInterfaceIpSubnetID string, virtualNetworkInterfaceIpAddress string, virtualNetworkInterfaceIpAutoDelete string, virtualNetworkInterfaceIpName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_reserved_ip" "is_reserved_ip_instance" {
			subnet_id = "%s"
			address = "%s"
			auto_delete = %s
			name = "%s"
			target {
				crn = "crn:v1:bluemix:public:is:us-south:a/123456::endpoint-gateway:r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/endpoint_gateways/r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				id = "r134-d7cc5196-9864-48c4-82d8-3f30da41fcc5"
				name = "my-endpoint-gateway"
				resource_type = "endpoint_gateway"
			}
		}

		data "ibm_is_reserved_ip" "is_reserved_ip_instance" {
			subnet_id = ibm_is_reserved_ip.is_reserved_ip.subnet_id
			id = "id"
		}
	`, virtualNetworkInterfaceIpSubnetID, virtualNetworkInterfaceIpAddress, virtualNetworkInterfaceIpAutoDelete, virtualNetworkInterfaceIpName)
}
