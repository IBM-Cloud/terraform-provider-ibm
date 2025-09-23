// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewaysDataSource_basic(t *testing.T) {
	t.Skip()
	resName := "data.ibm_is_virtual_endpoint_gateways.test1"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.name"),
				),
			},
		},
	})
}
func TestAccIBMISVirtualEndpointGatewaysDataSource_AllowDnsResolutionBinding(t *testing.T) {
	// t.Skip()
	resName := "data.ibm_is_virtual_endpoint_gateways.test1"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	enable_hub := false
	allowbinding := true
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewaysAllowDnsResolutionBindingDataSourceConfig(vpcname1, name1, enable_hub, allowbinding),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.allow_dns_resolution_binding"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateways" "test1" {
		depends_on = [ibm_is_virtual_endpoint_gateway.endpoint_gateway]
	}`)
}
func testAccCheckIBMISVirtualEndpointGatewaysAllowDnsResolutionBindingDataSourceConfig(vpcname1, name1 string, enable_hub, allowbinding bool) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigAllowDnsResolutionBinding(vpcname1, name1, enable_hub, allowbinding) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateways" "test1" {
		depends_on = [ibm_is_virtual_endpoint_gateway.endpoint_gateway]
	}`)
}

// service endpoints

func TestAccIBMISVirtualEndpointGatewaysDataSource_service_endpoints(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateways.test1"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewaysDataSourceServiceEndpointsConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.service_endpoints.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewaysDataSourceServiceEndpointsConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigServiceEndpoints(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateways" "test1" {
		depends_on = [ibm_is_virtual_endpoint_gateway.endpoint_gateway]
	}`)
}
