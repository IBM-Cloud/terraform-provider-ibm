// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVPCDnsResolutionBindingsDataSourceBasic(t *testing.T) {
	vpcname1 := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	enable_hub1 := true
	vpcname2 := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	bindingname := fmt.Sprintf("tf-vpc-dns-binding-%d", acctest.RandIntRange(10, 100))
	enable_hub2 := false
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPCDnsResolutionBindingsDataSourceConfigBasic(vpcname1, vpcname2, bindingname, enable_hub1, enable_hub2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "vpc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_bindings.is_vpc_dns_resolution_bindings", "dns_resolution_bindings.0.vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPCDnsResolutionBindingsDataSourceConfigBasic(vpcname1, vpcname2, bindingname string, enable_hub1, enable_hub2 bool) string {
	return testAccCheckIBMIsVPCDnsResolutionBindingResourceConfigBasic(vpcname1, vpcname2, bindingname, enable_hub1, enable_hub2) + fmt.Sprintf(`
		data "ibm_is_vpc_dns_resolution_bindings" "is_vpc_dns_resolution_bindings" {
			vpc_id = ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding.vpc_id
		}
	`)
}
