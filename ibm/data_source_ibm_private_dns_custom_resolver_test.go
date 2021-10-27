// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_custom_resolvers.test-cr"
	crname := fmt.Sprintf("tf-pdns-custom-resolver-%d", acctest.RandIntRange(100, 200))
	crdescription := fmt.Sprintf("tf-pdns-custom-resolver-tf-test%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.name"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.description"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.enabled"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.subnet_crn"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription string) string {
	return fmt.Sprintf(`
	resource "ibm_dns_custom_resolver" "test" {
		name		= "%s"
		instance_id	= "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
		description	= "%s"
		high_availability = false
		locations{
			subnet_crn	= "crn:v1:bluemix:public:is:us-south-1:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0717-4f53a236-cd7a-4688-9347-066bb5058a5c"
			enabled		= true
		}
	}
	data "ibm_dns_custom_resolvers" "test-cr" {
		depends_on  = [ibm_dns_custom_resolver.test]
		instance_id	= ibm_dns_custom_resolver.test.instance_id
	}`, crname, crdescription)
}
