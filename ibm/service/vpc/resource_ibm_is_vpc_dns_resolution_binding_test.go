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

func TestAccIBMIsVPCDnsResolutionBindingResourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVPCDnsResolutionBindingResourceConfigBasic(vpcname1, vpcname2, bindingname, enable_hub1, enable_hub2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "vpc_id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpc_dns_resolution_binding.is_vpc_dns_resolution_binding", "vpc.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPCDnsResolutionBindingResourceConfigBasic(vpcname1, vpcname2, bindingname string, enablehub1, enablehub2 bool) string {
	return fmt.Sprintf(`
	resource ibm_is_vpc testacc_vpc1 {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	resource ibm_is_vpc testacc_vpc2 {
		name = "%s"
		dns {
			enable_hub = %t
		}
	}
	resource ibm_is_vpc_dns_resolution_binding is_vpc_dns_resolution_binding {
		name = "%s"
		vpc_id=  ibm_is_vpc.testacc_vpc2.id
		vpc {
			id = ibm_is_vpc.testacc_vpc1.id
		}
	}
	`, vpcname1, enablehub1, vpcname2, enablehub2, bindingname)
}
