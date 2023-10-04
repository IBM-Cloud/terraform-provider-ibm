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

func TestAccIBMIsVirtualNetworkInterfaceIpsDataSourceBasic(t *testing.T) {
	virtualNetworkInterfaceIpSubnetID := fmt.Sprintf("tf_subnet_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfigBasic(virtualNetworkInterfaceIpSubnetID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "total_count"),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceIpsDataSourceAllArgs(t *testing.T) {
	virtualNetworkInterfaceIpSubnetID := fmt.Sprintf("tf_subnet_id_%d", acctest.RandIntRange(10, 100))
	virtualNetworkInterfaceIpAddress := fmt.Sprintf("tf_address_%d", acctest.RandIntRange(10, 100))
	virtualNetworkInterfaceIpAutoDelete := "false"
	virtualNetworkInterfaceIpName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfig(virtualNetworkInterfaceIpSubnetID, virtualNetworkInterfaceIpAddress, virtualNetworkInterfaceIpAutoDelete, virtualNetworkInterfaceIpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "subnet_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "sort"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "target_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "target_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "target_name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "target_resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "next.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.#"),
					resource.TestCheckResourceAttr("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.address", virtualNetworkInterfaceIpAddress),
					resource.TestCheckResourceAttr("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.auto_delete", virtualNetworkInterfaceIpAutoDelete),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.lifecycle_state"),
					resource.TestCheckResourceAttr("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.name", virtualNetworkInterfaceIpName),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.owner"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "reserved_ips.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reserved_ips.is_reserved_ips", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfigBasic(virtualNetworkInterfaceIpSubnetID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_reserved_ip" "is_reserved_ip_instance" {
			subnet_id = "%s"
		}

		data "ibm_is_reserved_ips" "is_reserved_ips_instance" {
			subnet_id = ibm_is_reserved_ip.is_reserved_ip.subnet_id
			sort = "name"
			target.id = "target.id"
			target.crn = "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727"
			target.name = "my-resource"
			target.resource_type = "target.resource_type"
		}
	`, virtualNetworkInterfaceIpSubnetID)
}

func testAccCheckIBMIsVirtualNetworkInterfaceIpsDataSourceConfig(virtualNetworkInterfaceIpSubnetID string, virtualNetworkInterfaceIpAddress string, virtualNetworkInterfaceIpAutoDelete string, virtualNetworkInterfaceIpName string) string {
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

		data "ibm_is_reserved_ips" "is_reserved_ips_instance" {
			subnet_id = ibm_is_reserved_ip.is_reserved_ip.subnet_id
			sort = "name"
			target.id = "target.id"
			target.crn = "crn:v1:bluemix:public:is:us-south:a/123456::load-balancer:dd754295-e9e0-4c9d-bf6c-58fbc59e5727"
			target.name = "my-resource"
			target.resource_type = "target.resource_type"
		}
	`, virtualNetworkInterfaceIpSubnetID, virtualNetworkInterfaceIpAddress, virtualNetworkInterfaceIpAutoDelete, virtualNetworkInterfaceIpName)
}
