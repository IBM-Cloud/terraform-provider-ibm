// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewayDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateway.data_test"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name1),
				),
			},
		},
	})
}
func TestAccIBMISVirtualEndpointGatewayDataSource_AllowDnsResolutionBinding(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateway.data_test"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	enable_hub := false
	allowbinding := true
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayAllowDnsResolutionBindingDataSourceConfig(vpcname1, name1, enable_hub, allowbinding),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name1),
					resource.TestCheckResourceAttr(
						resName, "allow_dns_resolution_binding", fmt.Sprintf("%t", allowbinding)),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateway" "data_test" {
        name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
	}`)
}
func testAccCheckIBMISVirtualEndpointGatewayAllowDnsResolutionBindingDataSourceConfig(vpcname1, name1 string, enable_hub, allowbinding bool) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigAllowDnsResolutionBinding(vpcname1, name1, enable_hub, allowbinding) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateway" "data_test" {
        name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
	}`)
}

// service endpoints
func TestAccIBMISVirtualEndpointGatewayDataSource_service_endpoints(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateway.data_test"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayDataSourceServiceEndpointsConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name1),
					resource.TestCheckResourceAttrSet(
						resName, "service_endpoints.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewayDataSourceServiceEndpointsConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigServiceEndpoints(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateway" "data_test" {
        name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
	}`)
}
