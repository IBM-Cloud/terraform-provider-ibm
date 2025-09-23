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

func TestAccIBMIsVPCDnsResolutionBindingDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVPCDnsResolutionBindingDataSourceConfigBasic(vpcname1, vpcname2, bindingname, enable_hub1, enable_hub2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "vpc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPCDnsResolutionBindingDataSourceConfigBasic(vpcname1, vpcname2, bindingname string, enable_hub1, enable_hub2 bool) string {
	return testAccCheckIBMIsVPCDnsResolutionBindingResourceConfigBasic(vpcname1, vpcname2, bindingname, enable_hub1, enable_hub2) + fmt.Sprintf(`
		data "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding" {
			vpc_id = ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding.vpc_id
			identifier = split("/", ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding.id)[1]
		}
	`)
}
